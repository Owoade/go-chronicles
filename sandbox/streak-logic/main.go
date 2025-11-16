package main

import (
	"fmt"

	"sandbox.go/utils"
)

type Prediction struct {
	Direction        string // "bullish" or "bearish"
	PredictionResult string // "hit" or "miss"
}

func main() {

	predictions := []Prediction{
		{"bullish", "hit"}, // bull = 1
		{"bullish", "hit"}, // bull = 2
		{"bearish", "hit"}, // BUG: saves bullishStreak(0), then bearish = 1
		{"bearish", "hit"}, // bear = 2
		{"bullish", "hit"}, // BUG: saves bearishStreak(0), then bull = 1
		{"bullish", "hit"}, // bull = 2
	}

	getStreak(predictions)

}

func getStreak(p []Prediction) {
	bearishStreaks := []int{}
	bullishStreaks := []int{}
	bearishStreak := 0
	bullishStreak := 0

	for _, prediction := range p {
		if prediction.Direction == "bullish" && prediction.PredictionResult == "hit" {
			bullishStreak++
			if !utils.Contains(bearishStreaks, bearishStreak) {
				bearishStreaks = append(bearishStreaks, bearishStreak)
			}
			bearishStreak = 0
		} else if prediction.Direction == "bearish" && prediction.PredictionResult == "hit" {
			bearishStreak++
			if !utils.Contains(bullishStreaks, bullishStreak) {
				bullishStreaks = append(bullishStreaks, bullishStreak)
			}
			bullishStreak = 0
		} else {
			if !utils.Contains(bearishStreaks, bearishStreak) {
				bearishStreaks = append(bearishStreaks, bearishStreak)
			}
			bearishStreak = 0
			if !utils.Contains(bullishStreaks, bullishStreak) {
				bullishStreaks = append(bullishStreaks, bullishStreak)
			}
			bullishStreak = 0
		}

	}

	if bullishStreak != 0 {
		if !utils.Contains(bullishStreaks, bullishStreak) {
			bullishStreaks = append(bullishStreaks, bullishStreak)
		}
	}

	if bearishStreak != 0 {
		if !utils.Contains(bearishStreaks, bearishStreak) {
			bearishStreaks = append(bearishStreaks, bearishStreak)
		}
	}

	longestBullishStreak := utils.MaxElementInSlice(bullishStreaks)
	longestBearishStreak := utils.MaxElementInSlice(bearishStreaks)

	fmt.Println("Longest bearish streak", longestBearishStreak)
	fmt.Println("Longest bullish streak", longestBullishStreak)

}
