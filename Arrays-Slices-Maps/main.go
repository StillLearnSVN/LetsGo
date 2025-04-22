package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {

	// 1. Reading Input and Splitting into Words
	scan := bufio.NewScanner(os.Stdin)
	words := make(map[string]int)

	scan.Split(bufio.ScanWords)

	/*
		bufio.NewScanner(os.Stdin): Creates a scanner to read input from the terminal or a piped file.
		scan.Split(bufio.ScanWords): Configures the scanner to split the input into words (default is splitting by lines).
		words := make(map[string]int): Initializes a map to store each word as a key and its frequency as the value.
	*/

	// 2. Counting Word Frequencies
	for scan.Scan() {
		words[scan.Text()]++
	}

	/*
		scan.Scan(): Reads the next word from the input.
		scan.Text(): Retrieves the current word as a string.
		words[scan.Text()]++: Increments the count of the word in the words map. If the word doesn't exist in the map, it is added with an initial count of 1.
	*/

	// 3. Printing the Number of Unique Words
	fmt.Println(len(words), "unique words")

	/*
		len(words): Counts the number of unique keys (words) in the map.
		This prints the total number of unique words found in the input.
	*/

	// 4. Sorting Words by Frequency
	type kv struct {
		key string
		val int
	}

	var ss []kv

	for k, v := range words {
		ss = append(ss, kv{k, v})
	}

	/*
		type kv: Defines a struct with two fields: key (the word) and val (its frequency).
		var ss []kv: Declares a slice of kv structs to store the word-frequency pairs.
		for k, v := range words: Iterates over the words map.
		ss = append(ss, kv{k, v}): Converts the map into a slice of kv structs for sorting.
	*/

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].val > ss[j].val
	})

	/*
		sort.Slice: Sorts the ss slice in-place.
		func(i, j int) bool: A custom comparison function that sorts the slice in descending order of val (frequency). If ss[i].val > ss[j].val, ss[i] comes before ss[j].
	*/

	// 5. Printing Words and Their Frequencies
	for _, s := range ss {
		fmt.Println(s.key, "appears", s.val, "times")
	}

	/*
		for _, s := range ss: Iterates over the sorted slice of kv structs.
		fmt.Println(s.key, "appears", s.val, "times"): Prints each word and its frequency in descending order.
	*/

	// 6. Key Concepts
	/* Key Concepts:
	Maps: Used to store word frequencies efficiently.
	Slices: Used to sort the word-frequency pairs.
	Sorting: Custom sorting with sort.Slice allows flexible ordering.
	bufio.Scanner: A powerful tool for reading and processing input.
	*/

}
