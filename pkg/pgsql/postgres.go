package pgsql

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	defaultMaxPoolSize  = 10
	defaultConnAttempts = 3
	defaultConnTimeout  = 60 * time.Second
)

type Postgres struct {
	*pgxpool.Pool
	maxPoolSize  int
	connAttempts int
	connTimeout  time.Duration
	shardLabel   string
	serviceName  string
}

func New(connString string, opts ...Option) (*Postgres, error) {
	pg := &Postgres{
		maxPoolSize:  defaultMaxPoolSize,
		connAttempts: defaultConnAttempts,
		connTimeout:  defaultConnTimeout,
	}

	for _, opt := range opts {
		opt(pg)
	}

	poolConfig, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("postgres - NewPostgres - pgxpool.ParseConfig: %w", err)
	}

	poolConfig.MaxConns = int32(pg.maxPoolSize)
	for pg.connAttempts > 0 {
		pool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
		if err == nil {
			pg.Pool = pool
			break
		}

		log.Printf("Postgres is trying to connect, attempts left: %d", pg.connAttempts)
		time.Sleep(pg.connTimeout)
		pg.connAttempts--
	}

	if err != nil {
		return nil, fmt.Errorf("postgres - NewPostgres - connAttempts == 0: %w", err)
	}

	return pg, nil
}

func (p *Postgres) RegisterCollector() error {
	fn := func() PgxStat { return p.Pool.Stat() }
	labels := map[string]string{"shard": p.shardLabel, "service": p.serviceName}
	collector := NewCollector(fn, labels)
	return prometheus.Register(collector)
}

func (p *Postgres) Close() {
	if p.Pool != nil {
		p.Pool.Close()
	}
}

func (p *Postgres) PoolPing() error {
	return p.Pool.Ping(context.Background())
}
