package service

import (
	"log"
	"poker/model"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func Judge(input model.Input) string {
	cards := input.Cards
	eachCard := strings.Split(cards, " ")

	markArray := []string{}
	numArray := []int{}

	rMark := regexp.MustCompile(`[CDHS]`)
	rNum := regexp.MustCompile(`[CDHS]{1}(1[0-3]\b|[1-9]\b)`)

	for _, c := range eachCard {
		num := rNum.FindString(c)
		mark := rMark.FindString(c)
		if num != "" {
			num = rMark.ReplaceAllString(num, "")
			intVar, err := strconv.Atoi(num)
			if err != nil {
				log.Fatal(err)
			}
			numArray = append(numArray, intVar)
		}
		if mark != "" {
			markArray = append(markArray, mark)
		}
	}
	response := judgeSuit(numArray, markArray)
	return response
}

// return suit
func judgeSuit(numArray []int, markArray []string) string {
	var outcome string

	sort.Ints(numArray)
	sortNums := numArray
	variation := sortNums[4] - sortNums[0]

	numCountsOverlap := numArrayToSize(sortNums)
	uniqNumsSize := len(uniqInt(sortNums))

	markArrayUniq := uniq(markArray)
	countMarks := countFactor(markArrayUniq)

	config := model.ConfigJudge{SortNums: sortNums,
		CountMarks:       countMarks,
		Variation:        variation,
		NumCountsOverlap: numCountsOverlap,
		UniqNumsSize:     uniqNumsSize}

	outcome = conditionalJudgement(config)
	return outcome
}

// make markArray unique
func uniq(slice []string) []string {
	m := make(map[string]bool)
	uniqSlice := make([]string, 0)

	for _, v := range slice {
		if _, ok := m[v]; !ok {
			m[v] = true
			uniqSlice = append(uniqSlice, v)
		}
	}
	return uniqSlice
}

func uniqInt(slice []int) []int {
	m := make(map[int]bool)
	uniqSlice := make([]int, 0)

	for _, v := range slice {
		if _, ok := m[v]; !ok {
			m[v] = true
			uniqSlice = append(uniqSlice, v)
		}
	}
	return uniqSlice
}

// count markArray's each mark
func countFactor(markArray []string) int {
	var count int
	sizeMap := map[string]int{"C": 0, "D": 0, "H": 0, "S": 0}
	// [H, H, S, D, S]
	for _, v := range markArray {
		sizeMap[v]++
	}
	for _, v := range sizeMap {
		count = count + v
	}
	return count
}

// slice comparison
func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// count numArray's overlapping
func numArrayToSize(numArray []int) []int {
	nSizeMap := map[int]int{0: 0}
	delete(nSizeMap, 0)
	var responseSlice []int
	for _, v := range numArray {
		nSizeMap[v] = nSizeMap[v] + 1
	}
	for _, v := range nSizeMap {
		responseSlice = append(responseSlice, v)
	}
	sort.Ints(responseSlice)
	return responseSlice
}

func conditionalJudgement(con model.ConfigJudge) string {
	var outcome string

	var (
		sortNums         = con.SortNums
		variation        = con.Variation
		countMarks       = con.CountMarks
		numCountsOverlap = con.NumCountsOverlap
		uniqNumsSize     = con.UniqNumsSize
	)

	if comparison := []int{1, 10, 11, 12, 13}; Equal(sortNums, comparison) && countMarks == 1 {
		outcome = "ロイヤルストレートフラッシュ"
	} else if countMarks == 1 && variation == 4 {
		outcome = "ストレートフラッシュ"
	} else if comparison := []int{1, 4}; Equal(numCountsOverlap, comparison) {
		outcome = "フォーカード"
	} else if comparison := []int{2, 3}; Equal(numCountsOverlap, comparison) {
		outcome = "フルハウス"
	} else if uniqNumsSize == 5 && variation == 4 {
		outcome = "ストレート"
	} else if countMarks == 1 && uniqNumsSize == 5 {
		outcome = "フラッシュ"
	} else if comparison := []int{1, 1, 3}; Equal(numCountsOverlap, comparison) {
		outcome = "スリーカード"
	} else if comparison := []int{1, 2, 2}; Equal(numCountsOverlap, comparison) {
		outcome = "ツーペア"
	} else if comparison := []int{1, 1, 1, 2}; Equal(numCountsOverlap, comparison) {
		outcome = "ワンペア"
	} else {
		outcome = "ハイカード"
	}
	return outcome
}
