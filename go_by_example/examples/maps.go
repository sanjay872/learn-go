package examples

import (
	"fmt"
	"maps"
)

func Maps() {
	m := make(map[string]int)

	m["1"] = 1

	m["2"] = 2

	fmt.Println(m)

	v1 := m["1"]
	v2 := m["2"]

	fmt.Println(v1, v2)

	delete(m, "1")

	fmt.Println(m)

	clear(m)

	fmt.Println(m)

	m["1"] = 1
	m["2"] = 2

	_, val := m["3"] // to check if the key exist, val returns true or false

	fmt.Println(val)

	n1 := map[string]int{"1": 1, "2": 2}
	n2 := map[string]int{"1": 1, "2": 2}

	if maps.Equal(n1, n2) {
		fmt.Println("they are equal")
	}
}
