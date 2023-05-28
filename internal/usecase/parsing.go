package usecase

import (
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/agnivade/levenshtein"
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/html"
)

func GetTextRelatedToTag(url, tag, word string) (string, error) {
	// Загрузка HTML-страницы
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	var text string
	doc.Find("p").Each(func(i int, s *goquery.Selection) {
		text += s.Text() + " "
	})

	fmt.Println(text)
	// Очистка текста от HTML-тегов
	text = minifyText(text)

	// Разбиение текста на слова
	words := strings.Fields(text)

	// Поиск наиболее близких слов по смыслу
	closestWords := findClosestWords(words, word)

	fmt.Println(closestWords)
	if len(closestWords) < 3 {
		return "Текст не релевантен", nil
	}

	return text, nil
}

// Очистка текста от HTML-тегов
func minifyText(text string) string {
	m := minify.New()
	m.AddFunc("text/html", html.Minify)
	minified, err := m.String("text/html", text)
	if err != nil {
		log.Fatal(err)
	}
	return minified
}

func findClosestWords(words []string, word string) []string {
	// Создаем слайс для хранения пар (слово, расстояние Левенштейна)
	type wordDistance struct {
		word     string
		distance int
	}
	var distances []wordDistance

	// Вычисляем расстояние Левенштейна для каждого слова
	for _, w := range words {
		dist := levenshtein.ComputeDistance(word, w)
		fmt.Println(dist, w)
		distances = append(distances, wordDistance{word: w, distance: dist})
	}

	// Сортируем слова по возрастанию расстояния Левенштейна
	sort.Slice(distances, func(i, j int) bool {
		return distances[i].distance < distances[j].distance
	})

	// Создаем слайс для хранения наиболее похожих слов
	var closestWords []string

	// Выбираем первые 5 слов с минимальным расстоянием Левенштейна
	for i := 0; i < 5 && i < len(distances); i++ {
		if distances[i].distance < 3 {
			closestWords = append(closestWords, distances[i].word)
		}
	}

	return closestWords
}
