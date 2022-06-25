package model

type Input struct {
	Cards string `json:"input" binding:"required"`
}

type ConfigJudge struct {
	SortNums         []int
	CountMarks       int
	Variation        int
	NumCountsOverlap []int
	UniqNumsSize     int
}
