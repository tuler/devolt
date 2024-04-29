package entity

import (
	"gonum.org/v1/gonum/stat"
	"log"
	"math"
	"math/rand"
	"time"
)

type Payload struct {
	Station_ID string  `json:"station_id"`
	Battery    float64 `json:"battery"`
	Percentage float64 `json:"percentage"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
}

func EntropyWithConfidenceInterval(min float64, max float64, z float64) float64 {

	// https://en.wikipedia.org/wiki/Confidence_interval
	// 𝑥̄ ± 𝑧 × 𝑠/√𝑛
	//
	// Where:
	//   𝑥̄ is the sample mean.
	//   𝑧 is the critical value from the standard normal distribution for the desired confidence level.
	//   𝑠 is the sample standard deviation.
	//   𝑛 is the sample size.

	interval := make([]float64, int(max-min)+1)
	for i := range interval {
		interval[i] = float64(min) + float64(i)
	}
	mean, stdDev := stat.MeanStdDev(interval, nil)
	literal := stdDev / math.Sqrt(float64(len(interval)))
	a := mean - z*literal
	b := mean + z*literal
	rand.NewSource(time.Now().UnixNano())
	return math.Round(rand.Float64()*(a-b) + b)
}

func NewPayload(id string, params map[string]interface{}, latitude float64, longitude float64) (*Payload, error) {
	min, ok := params["min"].(float64)
	if !ok {
		log.Fatalf("min value not found or not a float64: %v", params["min"])
	}
	max, ok := params["max"].(float64)
	if !ok {
		log.Fatalf("max value not found or not a float64: %v", params["max"])
	}
	batteryValue := EntropyWithConfidenceInterval(min, max, 1.96) // 95% confidence interval with z = 1.96 (https://en.wikipedia.org/wiki/Standard_normal_table)
	batteryPercentage := (batteryValue - min) / (max - min) * 100
	return &Payload{
		Station_ID: id,
		Battery:    batteryValue,
		Percentage: batteryPercentage,
		Latitude:   latitude,
		Longitude:  longitude,
	}, nil
}
