package Greedy

import (
	"fmt"
	"testing"
)

func TestNumWaterBottles(t *testing.T) {

	//numBottles := 9
	//numExchange := 3

	//numBottles := 15
	//numExchange := 4

	numBottles := 15
	numExchange := 8

	nums := numWaterBottles(numBottles, numExchange)
	fmt.Println(nums)
}
