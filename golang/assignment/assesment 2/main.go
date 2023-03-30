package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	// Load dataset from CSV file
	dataset, err := loadDataset("students_scores.csv")
	if err != nil {
		log.Fatal(err)
	}

	// Split dataset into training and testing sets
	trainingSet, testingSet := splitDataset(dataset, 0.7)

	// Train the model
	weights := train(trainingSet)

	// Test the model
	accuracy := test(testingSet, weights)
	fmt.Printf("Accuracy: %.2f%%\n", accuracy*100)
}

// loadDataset loads a CSV file and returns a slice of samples.
func loadDataset(filename string) ([][]float64, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)

	var dataset [][]float64

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		var sample []float64
		for _, value := range record {
			x, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return nil, err
			}
			sample = append(sample, x)
		}

		dataset = append(dataset, sample)
	}

	return dataset, nil
}

// splitDataset splits a dataset into training and testing sets.
func splitDataset(dataset [][]float64, splitRatio float64) ([][]float64, [][]float64) {
	trainSize := int(math.Round(float64(len(dataset)) * splitRatio))
	trainSet := make([][]float64, trainSize)
	copy(trainSet, dataset)

	testSet := make([][]float64, len(dataset)-trainSize)
	copy(testSet, dataset[trainSize:])

	return trainSet, testSet
}

// train trains a linear regression model using gradient descent.
func train(dataset [][]float64) []float64 {
	// Initialize weights to 0
	weights := make([]float64, len(dataset[0])-1)

	// Set hyperparameters
	alpha := 0.01
	numIterations := 1000

	// Train the model using gradient descent
	for i := 0; i < numIterations; i++ {
		for _, sample := range dataset {
			// Make a prediction using the current weights
			prediction := predict(sample[:len(sample)-1], weights)

			// Compute the error
			error := sample[len(sample)-1] - prediction

			// Update the weights
			for j := 0; j < len(weights); j++ {
				weights[j] += alpha * error * sample[j]
			}
		}
	}

	return weights
}

// test tests the accuracy of the linear regression model.
func test(dataset [][]float64, weights []float64) float64 {
	numCorrect := 0

	for _, sample := range dataset {
		// Make a prediction using the trained weights
		prediction := predict(sample[:len(sample)-1], weights)

		// Classify the sample based on the prediction
		if prediction >= 0.5 {
			if sample[len(sample)-1] == 1 {
				numCorrect++
			}
		} else {
			if sample[len(sample)-1] == 0 {
				numCorrect++
			}
		}
	

		