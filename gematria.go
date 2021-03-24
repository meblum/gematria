// Package gematria implements encoding of hebrew string into its gematria value.
//
// The exact mapping between letters and numbers is described in the
// documentation for the Value() function.
package gematria

import (
	"fmt"
)

var runeValues = map[rune]int{
	rune(1488): 1,
	rune(1489): 2,
	rune(1490): 3,
	rune(1491): 4,
	rune(1492): 5,
	rune(1493): 6,
	rune(1494): 7,
	rune(1495): 8,
	rune(1496): 9,
	rune(1497): 10,
	rune(1499): 20,
	rune(1500): 30,
	rune(1502): 40,
	rune(1504): 50,
	rune(1505): 60,
	rune(1506): 70,
	rune(1508): 80,
	rune(1510): 90,
	rune(1511): 100,
	rune(1512): 200,
	rune(1513): 300,
	rune(1514): 400,

	rune(1478): 50, // Nun Hafukha

	rune(1519): 30, // Yod Triangle
	rune(1520): 12, // Double Vav
	rune(1521): 16, // Vav Yod
	rune(1522): 20, // Double Yod

	rune(1498): 20, // Final Kaf
	rune(1501): 40, // Final Mem
	rune(1503): 50, // Final Nun
	rune(1507): 80, // Final Pe
	rune(1509): 90, // Final Tsadi

	rune(64285): 10,
	rune(64287): 20, // Double Yod
	rune(64288): 70,
	rune(64289): 1,
	rune(64290): 4,
	rune(64291): 5,
	rune(64292): 20,
	rune(64293): 30,
	rune(64294): 40, // Final Mem
	rune(64295): 200,
	rune(64296): 400,
	rune(64298): 300,
	rune(64299): 300,
	rune(64300): 300,
	rune(64301): 300,
	rune(64302): 1,
	rune(64303): 1,
	rune(64304): 1,
	rune(64305): 2,
	rune(64306): 3,
	rune(64307): 4,
	rune(64308): 5,
	rune(64309): 6,
	rune(64310): 7,
	rune(64312): 9,
	rune(64313): 10,
	rune(64314): 20, // Final Kaf
	rune(64315): 20,
	rune(64316): 30,
	rune(64318): 40,
	rune(64320): 50,
	rune(64321): 60,
	rune(64323): 80, // Final Pe
	rune(64324): 80,
	rune(64326): 90,
	rune(64327): 100,
	rune(64328): 200,
	rune(64329): 300,
	rune(64330): 400,
	rune(64331): 6,
	rune(64332): 2,
	rune(64333): 20,
	rune(64334): 80,
	rune(64335): 31, // Alef Lamed
}

// Value returns the gematria value of str, and an error if result has overflowed.
// Value is calculated using the standard encoding,
// assigning the values 1–9, 10–90, 100–400 to the 22 Hebrew letters in order.
// Final letters have the value as the non-final. Non Hebrew characters are ignored.
func Value(str string) (int, error) {
	var sum int
	for _, r := range str {

		result, ok := add(sum, int(runeValues[r]))
		sum = result

		if !ok {
			return result, fmt.Errorf("string is too long")
		}
	}
	return sum, nil
}

// add returns the sum of its arguments and a boolean flag which is false if the result overflows an int.
func add(a, b int) (value int, ok bool) {
	result := a + b
	//Overflow if both arguments have the opposite sign of the result.
	if ((a ^ result) & (b ^ result)) < 0 {
		return result, false
	}

	return result, true
}
