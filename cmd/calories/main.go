package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"os"
)

// Formula constant for calculating Calories per minute.
const calorieConst float64 = 3.5

const kgConversion float64 = 0.45359237

func burnRate(MET, kg float64) float64 {
	return (MET * calorieConst * kg) / 200
}

func lbTokg(lb float64) float64 {
	return roundWeight(lb * kgConversion)
}

func roundWeight(weight float64) float64 {
	return math.Round(weight*10.0) / 10.0
}

func totalBurn(w io.Writer, kg, MET, time float64) {
	calPerMinute := burnRate(MET, kg)

	calories := calPerMinute * time

	// Print with default width and precision of 2 (for rounding).
	fmt.Fprintf(w, "You burned %.2f Calories\n", calories)
}

func main() {
	var kg, lb, MET, minutes float64
	flag.Float64Var(&kg, "kg", -99, "your weight in KG")
	flag.Float64Var(&lb, "lb", -99, "your weight in pounds")
	flag.Float64Var(&MET, "met", 5.00, "MET for the task")
	flag.Float64Var(&minutes, "time", 20.00, "time you spent in minutes")
	flag.Parse()

	if kg == -99 && lb == -99 {
		log.Fatalln("Must provide either -kg or -lb.")
	}

	if kg >= 1 {
		totalBurn(os.Stdout, kg, MET, minutes)
		return
	}

	kg = lbTokg(lb)

	totalBurn(os.Stdout, kg, MET, minutes)
}
