package main

import (
	"fmt"
	"sort"
)

func isStraight(cards []int) bool {
	// Ordenar las cartas
	sort.Ints(cards)

	// Contar la cantidad de veces que aparece el AS
	aceCount := 0

	// Verificar si hay duplicados
	seen := make(map[int]bool)
	for _, card := range cards {
		if card == 1 {
			aceCount++
		}
		if seen[card] {
			return false
		}
		seen[card] = true
	}

	// Si hay más de un AS, uno de ellos debe ser considerado como 1
	if aceCount > 1 {
		aceCount = 1
	}

	// Verificar si las cartas forman una escalera
	for i := 0; i < len(cards)-1; i++ {
		if cards[i+1]-cards[i] != 1 {
			if aceCount > 0 {
				aceCount--
			} else {
				return false
			}
		}
	}

	return true
}

func main() {
	// Ejemplos
	fmt.Println(isStraight([]int{2, 3, 4, 5, 6}))          // true
	fmt.Println(isStraight([]int{14, 5, 4, 2, 3}))         // true
	fmt.Println(isStraight([]int{7, 7, 12, 11, 3, 4, 14})) // false
	fmt.Println(isStraight([]int{7, 3, 2}))                // false
}
