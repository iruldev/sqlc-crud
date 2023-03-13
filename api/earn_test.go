package api

import (
	"fmt"
	"math"
	"testing"
)

func CalculateDIR(balance float64, yieldGain float64, durationInSeconds int64) float64 {
	yieldPerSecond := yieldGain / float64(durationInSeconds)
	dir := yieldPerSecond * 86400 / balance
	return dir
}

func CalculateAPRFromDIR(dailyRate float64) float64 {
	apr := math.Pow(1+dailyRate, 365) - 1
	return apr
}

func CalculateAPRWithAPYAndCompounding(APY float64, compounding string) float64 {
	var n float64
	switch compounding {
	case "daily":
		n = 365
	case "weekly":
		n = 52
	case "bi-weekly":
		n = 26
	case "monthly":
		n = 12
	case "quarterly":
		n = 4
	case "half-yearly":
		n = 2
	}
	apr := math.Pow(1+(APY/100)/n, n) - 1
	return apr
}

func CalculateAPYFromDIR(dailyRate float64) float64 {
	apy := math.Pow(1+dailyRate, 365) - 1
	return apy
}

func CalculateAPYFromAPR(apr float64) float64 {
	apy := math.Pow(1+apr/365, 365) - 1
	return apy
}

func TestCalculateDIR(t *testing.T) {
	testCase := []struct {
		name              string
		balance           float64
		yieldGain         float64
		durationInSeconds int64
	}{
		{
			name:              "1.a",
			balance:           100.0,
			yieldGain:         0.0253,
			durationInSeconds: 86400,
		},
		{
			name:              "1.b",
			balance:           72,
			yieldGain:         0.0621,
			durationInSeconds: 259200,
		},
		{
			name:              "1.c",
			balance:           68.5,
			yieldGain:         0.0038,
			durationInSeconds: 12600,
		},
		{
			name:              "1.d",
			balance:           24.54,
			yieldGain:         0.0002,
			durationInSeconds: 2700,
		},
	}

	for i := range testCase {
		tc := testCase[i]
		dailyInterestRate := CalculateDIR(tc.balance, tc.yieldGain, tc.durationInSeconds)
		fmt.Println(tc.name, dailyInterestRate)
	}
}

func TestCalculateAPRFromDIR(t *testing.T) {
	testCase := []struct {
		name              string
		balance           float64
		yieldGain         float64
		durationInSeconds int64
	}{
		{
			name:              "2.a",
			balance:           78,
			yieldGain:         0.0253,
			durationInSeconds: 86400,
		},
		{
			name:              "2.b",
			balance:           92,
			yieldGain:         0.0521,
			durationInSeconds: 432000,
		},
		{
			name:              "2.c",
			balance:           43,
			yieldGain:         0.00038,
			durationInSeconds: 5100,
		},
		{
			name:              "2.d",
			balance:           100,
			yieldGain:         0.0002,
			durationInSeconds: 2100,
		},
	}

	for i := range testCase {
		tc := testCase[i]
		dailyInterestRate := CalculateDIR(tc.balance, tc.yieldGain, tc.durationInSeconds)
		apr := CalculateAPRFromDIR(dailyInterestRate)
		fmt.Println(tc.name, apr)
	}
}

func TestCalculateAPRWithAPYAndCompounding(t *testing.T) {
	testCase := []struct {
		name        string
		APY         float64
		compounding string
	}{
		{
			name:        "1.g",
			APY:         2.4,
			compounding: "daily",
		},
		{
			name:        "1.h",
			APY:         4.8,
			compounding: "weekly",
		},
		{
			name:        "1.i",
			APY:         5.2,
			compounding: "bi-weekly",
		},
		{
			name:        "1.j",
			APY:         3.5,
			compounding: "monthly",
		},
		{
			name:        "1.k",
			APY:         8.5,
			compounding: "quarterly",
		},
		{
			name:        "1.l",
			APY:         7,
			compounding: "half-yearly",
		},
	}

	for i := range testCase {
		tc := testCase[i]
		apr := CalculateAPRWithAPYAndCompounding(tc.APY, tc.compounding)
		fmt.Println(tc.name, apr)
	}
}

func TestCalculateAPYFromAPR(t *testing.T) {
	testCase := []struct {
		name string
		apr  func() (dir float64, apr float64)
	}{
		{
			name: "3.a",
			apr: func() (dir float64, apr float64) {
				dir = CalculateDIR(55, 0.00432, 86400)
				apr = CalculateAPRFromDIR(dir)
				return dir, apr
			},
		},
		{
			name: "3.b",
			apr: func() (dir float64, apr float64) {
				dir = CalculateDIR(80, 0.0285, 432000)
				apr = CalculateAPRFromDIR(dir)
				return dir, apr
			},
		},
		{
			name: "3.c",
			apr: func() (dir float64, apr float64) {
				dir = CalculateDIR(76, 0.000128, 2520)
				apr = CalculateAPRFromDIR(dir)
				return dir, apr
			},
		},
		{
			name: "3.d",
			apr: func() (dir float64, apr float64) {
				dir = CalculateDIR(45.224, 0.000008, 180)
				apr = CalculateAPRFromDIR(dir)
				return dir, apr
			},
		},
		{
			name: "3.e",
			apr: func() (dir float64, apr float64) {
				dir = 0.032 / 100
				apr = 0
				return dir, apr
			},
		},
		{
			name: "3.f",
			apr: func() (dir float64, apr float64) {
				dir = 0.024 / 100
				apr = 0
				return dir, apr
			},
		},
		{
			name: "3.g",
			apr: func() (dir float64, apr float64) {
				dir = 0
				apr = 4.9 / 100
				return dir, apr
			},
		},
		{
			name: "3.h",
			apr: func() (dir float64, apr float64) {
				dir = 0
				apr = 8.1 / 100
				return dir, apr
			},
		},
	}

	for i := range testCase {
		tc := testCase[i]
		var apy float64
		dir, apr := tc.apr()
		if dir > 0 {
			apy = CalculateAPYFromDIR(dir)
		} else {
			apy = CalculateAPYFromAPR(apr)
		}
		fmt.Println(tc.name, apy)
	}
}
