package analys

import (
	"math"
)

func getTf(documents ...[]string) map[string]float64 {
	var counter = map[string]int{}
	var l = 0
	var tf = map[string]float64{}

	for _, document := range documents {
		for _, word := range document {
			counter[word]++
			l++
		}
	}

	for word, val := range counter {
		tf[word] = float64(val) / float64(l)
	}

	return tf
}

func getIdf(documents ...[]string) map[string]int {
	var counter = map[string]int{}

	for _, document := range documents {
		seen := make(map[string]int, len(document))
		for _, word := range document {
			seen[word]++
		}

		for key := range seen {
			counter[key]++
		}
	}

	return counter
}

func GetTfIdf(documents ...[]string) (map[string]float64, map[string]float64) {
	tf := getTf(documents...)
	wordInDocuments := getIdf(documents...)

	idf := make(map[string]float64, len(wordInDocuments))

	for word, val := range wordInDocuments {
		idf[word] = math.Log(float64(len(documents)) / float64(val))
	}

	return tf, idf
}