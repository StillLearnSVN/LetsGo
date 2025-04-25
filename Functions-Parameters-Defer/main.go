package main

import "fmt"

// The `do` function takes a pointer to a map of type `map[int]int`.
// This allows the function to modify the original map passed from `main`.
func do(m1 *map[int]int) {
	// Dereference the pointer to access the map and set the key `3` to value `0`.
	(*m1)[3] = 0

	// Replace the original map with a new, empty map.
	// This effectively discards all previous key-value pairs in the map.
	*m1 = make(map[int]int)

	// Add a new key-value pair to the newly created map: key `4` with value `4`.
	(*m1)[4] = 4

	// Print the current state of the map (after modifications) to the console.
	fmt.Println("m1:", *m1)
}

func main() {
	// Initialize a map with three key-value pairs.
	m := map[int]int{4: 1, 7: 2, 8: 3}

	// Print the initial state of the map to the console.
	fmt.Println(m)

	// Call the `do` function, passing the address of the map `m`.
	// This allows the `do` function to modify the original map.
	do(&m)

	// Print the final state of the map after the `do` function has modified it.
	fmt.Println(m)
}

// Explanation of Defer:
// The `defer` statement in Go is used to ensure that a function call is performed later in the program's execution
// after the surrounding function returns. It is often used for cleanup tasks, such as closing files or releasing resources.
