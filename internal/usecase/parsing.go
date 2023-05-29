package usecase

import (
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/agnivade/levenshtein"
	"github.com/kljensen/snowball"
	"github.com/pemistahl/lingua-go"
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/html"
)

func GetTextRelatedToTag(url, tag, word string) (map[string]string, error) {
	texts := make(map[string]string)
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}

	var text string

	doc.Find("p").Each(func(i int, s *goquery.Selection) {
		text += s.Text() + " "
	})

	text = minifyText(text)

	words := strings.Fields(text)

	var nominativeWord string
	if word != "" {
		wordLang, err := DetectLanguage(word)
		if err != nil {
			return nil, err
		}

		nominativeWord, err = snowball.Stem(strings.ToLower(word), strings.ToLower(wordLang), false)
		if err != nil {
			fmt.Println(err)
		}
	}

	var nominativeWords []string
	languageMap := make(map[string]string)
	for _, w := range words {
		lang, err := DetectLanguage(w)
		if err != nil {
			if !IsNumber(w) {
				languageMap[w] = "else"
			} else {
				languageMap[w] = "digit"
			}
		}
		languageMap[w] = strings.ToLower(lang)
		sklonWord, _ := snowball.Stem(strings.ToLower(w), strings.ToLower(languageMap[w]), false)

		nominativeWords = append(nominativeWords, sklonWord)
	}
	closestWords := findClosestWords(nominativeWords, nominativeWord)

	if len(closestWords) < 3 {
		for _, w := range closestWords {
			if w == word {
				texts[url] = text
			}
		}
	} else {
		texts[url] = text
	}
	if word != "" {
		texts = checkHrefs(doc, word, texts)

	}

	return texts, nil
}

func checkHrefs(doc *goquery.Document, word string, texts map[string]string) map[string]string {

	doc.Find("a").Each(func(index int, element *goquery.Selection) {
		href, exists := element.Attr("href")
		if exists {
			response, err := http.Get(href)
			if err != nil {
				return
			}
			defer response.Body.Close()

			// Создание объекта goquery.Document из полученного HTML-кода
			docSite, err := goquery.NewDocumentFromReader(response.Body)
			if err != nil {
				return
			}

			var text string
			docSite.Find("p").Each(func(i int, s *goquery.Selection) {
				text += s.Text() + " "
			})

			// Очистка текста от HTML-тегов
			text = minifyText(text)

			// Разбиение текста на слова
			words := strings.Fields(text)
			languageMap := make(map[string]string)

			var nominativeWord string
			if word != "" {
				wordLang, err := DetectLanguage(word)
				if err != nil {
					return
				}

				nominativeWord, err = snowball.Stem(strings.ToLower(word), strings.ToLower(wordLang), false)
				if err != nil {
					fmt.Println(err)
				}
			}

			var nominativeWords []string

			for _, w := range words {
				lang, err := DetectLanguage(w)
				if err != nil {
					if !IsNumber(w) {
						languageMap[w] = "else"
					} else {
						languageMap[w] = "digit"
					}
				}
				languageMap[w] = strings.ToLower(lang)

				sklonWord, _ := snowball.Stem(strings.ToLower(w), strings.ToLower(languageMap[w]), false)

				nominativeWords = append(nominativeWords, RemoveExtraCharacters(sklonWord))
			}
			closestWords := findClosestWords(nominativeWords, nominativeWord)

			if len(closestWords) < 3 {
				for _, w := range closestWords {
					if w == word {
						texts[href] = text
					}
				}
			} else {
				texts[href] = text
			}

		}
	})
	return texts
}

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
		// fmt.Println(dist, w)
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

func IsNumber(str string) bool {
	_, err := strconv.ParseFloat(str, 64)
	return err == nil
}

func DetectLanguage(word string) (string, error) {
	languages := []lingua.Language{
		lingua.English,
		lingua.Russian,
	}

	detector := lingua.NewLanguageDetectorBuilder().
		FromLanguages(languages...).
		Build()

	language, exists := detector.DetectLanguageOf(word)
	if exists {
		return language.String(), nil
	}
	return "", fmt.Errorf("Я не знаю такого языка")
}

func RemoveExtraCharacters(word string) string {
	// Определяем набор символов, которые нужно удалить
	charsToRemove := ",.\"'«»()"

	// Заменяем каждый символ из набора на пустую строку
	for _, char := range charsToRemove {
		word = strings.ReplaceAll(word, string(char), "")
	}

	return word
}
