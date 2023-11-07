package main

import (
	"github.com/Nav1Cr0ss/algorithms-lab4/internal/service"
)

func main() {

	populationSize := 100
	generations := 100

	capacity := 250
	numItems := 100

	service.Run(populationSize, generations, capacity, numItems)

}
