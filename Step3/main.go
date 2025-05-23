package main

import (
	"fmt"
	"sort"
	"strings"
)

func cleanText(text string) string {
	return strings.Map(func(r rune) rune {
		if strings.ContainsRune(".,!?:;\"'-()[]{}", r) {
			return -1
		}
		return r
	}, text)
}

func getTopWords(wordMap map[string]int, n int) []string {
	// Собираем слова и их счётчики
	type pair struct {
		word  string
		count int
		index int 
	}

	// Создаем слайс пар (слово, счёт, индекс)
	// Порядок добавления в map важен, поэтому заранее запомним его в AnalyzeText и сохраним его в map[word]index
	indexMap := map[string]int{}
	i := 0
	for word := range wordMap {
		indexMap[word] = i
		i++
	}

	pairs := []pair{}
	for word, count := range wordMap {
		pairs = append(pairs, pair{word, count, indexMap[word]})
	}

	// Сортируем: сначала по count убыв., потом по index возр.
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].count == pairs[j].count {
			return pairs[i].index < pairs[j].index
		}
		return pairs[i].count > pairs[j].count
	})

	// Формируем результат
	top := []string{}
	for i := 0; i < n && i < len(pairs); i++ {
		top = append(top, pairs[i].word)
	}
	return top
}

func AnalyzeText(text string) {
	text = strings.ToLower(text)
	text = cleanText(text)

	words := strings.Fields(text)
	fmt.Println("Количество слов:", len(words))

	wordCount := map[string]int{}
	wordOrder := []string{}
	seen := map[string]bool{}
	for _, word := range words {
		wordCount[word]++
		if !seen[word] {
			wordOrder = append(wordOrder, word)
			seen[word] = true
		}
	}

	// Создаем новую карту с тем же содержимым, но в порядке wordOrder
	orderedWordMap := map[string]int{}
	for _, word := range wordOrder {
		orderedWordMap[word] = wordCount[word]
	}

	fmt.Println("Количество уникальных слов:", len(wordCount))
	topWords := getTopWords(orderedWordMap, 5)
	if len(topWords) > 0 {
		fmt.Printf("Самое часто встречающееся слово: \"%s\" (встречается %d раз)\n", topWords[0], wordCount[topWords[0]])
	}
	fmt.Println("Топ-5 самых часто встречающихся слов:")
	for _, word := range topWords {
		fmt.Printf("\"%s\": %d раз\n", word, wordCount[word])
	}
}
