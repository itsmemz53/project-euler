package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindSumOfEventFibonacciSeriesWhereN400000AssertEqual(t *testing.T) {
	assert.Equal(t, findSumOrEvenFibonacciSeries(4000000), 4613732)
}

func TestFindSumOfEventFibonacciSeriesWhereN60GreaterMatchAssertEqual(t *testing.T) {
	assert.Equal(t, findSumOrEvenFibonacciSeries(60), 44)
}

func TestFindSumOfEventFibonacciSeriesWhereN34ExactMatchAssertEqual(t *testing.T) {
	assert.Equal(t, findSumOrEvenFibonacciSeries(34), 44)
}

func TestFindSumOfEventFibonacciSeriesWhereN0ExactMatchAssertEqual(t *testing.T) {
	assert.Equal(t, findSumOrEvenFibonacciSeries(0), 0)
}

func TestFindSumOfEventFibonacciSeriesWhereN1ExactMatchAssertEqual(t *testing.T) {
	assert.Equal(t, findSumOrEvenFibonacciSeries(1), 0)
}

func TestFindSumOfEventFibonacciSeriesWhereN2ExactMatchAssertEqual(t *testing.T) {
	assert.Equal(t, findSumOrEvenFibonacciSeries(2), 2)
}
