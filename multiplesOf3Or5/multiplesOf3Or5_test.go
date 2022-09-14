package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMultiplesOf3Or5WithN10AssertEqual(t *testing.T) {
	assert.Equal(t, findSumOfMultiplesOf3Or5(3, 5, 10), 23)
}

func TestMultiplesOf3Or5WithN100AssertEqual(t *testing.T) {
	assert.Equal(t, findSumOfMultiplesOf3Or5(3, 5, 100), 2318)
}

func TestMultiplesOf3Or5WithN1000AssertEqual(t *testing.T) {
	assert.Equal(t, findSumOfMultiplesOf3Or5(3, 5, 1000), 233168)
}

func TestMultiplesOf3Or5WithN0AssertEqual(t *testing.T) {
	assert.Equal(t, findSumOfMultiplesOf3Or5(3, 5, 0), 0)
}
