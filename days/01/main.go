package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func moduleFuel(moduleMass int64) int64 {
	return int64(float64(moduleMass) / 3.0 - 2)
}

func totalFuel(mass int64) int64 {
	var total int64 = mass

	for mass >= 0 {
		mass = moduleFuel(mass)
		if mass > 0 {
			total += mass
		}
	}

	return total
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var _moduleFuel int64 = 0
	var _totalFuel int64 = 0

	for scanner.Scan() {
		s := scanner.Text()
		i, _ := strconv.ParseInt(s, 10, 64)
		_fuel := moduleFuel(i)
		_moduleFuel += _fuel
		_totalFuel += totalFuel(_fuel)
	}

	fmt.Printf("just modules: %d\n", _moduleFuel)
	fmt.Printf("total (2) _moduleFuel: %d\n", _totalFuel)
}
