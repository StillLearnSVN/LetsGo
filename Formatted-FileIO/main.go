package main

import (
	"bufio"
	"fmt"
	// "io/ioutil"
	"os"
	"strings"
)

func main() {

	// -- Formatted output
	
	// a, b := 12, 345
	// c, d := 1.2, 3.45

	// fmt.Printf("Value of a: %d b: %d\n", a, b)
	// fmt.Printf("Value of c: %.2f d: %f\n", c, d)
	// fmt.Printf("The data type of a: %Tand datatype of b: %T\n", a, b)
	// fmt.Printf("The data type of c: %T and d: %T", c, d)


	// fmt.Println("Formatted output:")

	// fmt.Printf("Fomatted value of |a: %6d| |b: %6d|\n", a, b)
	// fmt.Printf("Fomatted value of |a: %06d| |b: %06d|\n", a, b)
	// fmt.Printf("Fomatted value of |a: %-6d| |b: %-6d|\n", a, b)

	// s := []int{1, 2, 3}
	// a := [3]rune{'a', 'b', 'c'}
	// m := map[string]int{"a": 1, "b": 2, "c": 3}

	// fmt.Printf("%T\n", s) // This will print the type of s
	// fmt.Printf("%v\n", s) // This will print the value of s
	// fmt.Printf("%#v\n", s) // This will print the combination of type and value

	// fmt.Printf("\n%T\n", a) // This will print the type of a
	// fmt.Printf("%v\n", a) // This will print the value of a
	// fmt.Printf("%#v\n", a) // This will print the combination of type and value

	// fmt.Printf("\n%T\n", m) // This will print the type of m
	// fmt.Printf("%v\n", m) // This will print the value of m
	// fmt.Printf("%#v\n", m) // This will print looks very much like what i type in my code

	// -- File I/O

	// for _, fname := range os.Args[1:] {
	// 	file, err := os.Open(fname)

	// 	if err != nil {
	// 		fmt.Fprintln(os.Stderr, err)
	// 		continue
	// 	}

	// 	if _, err :=io.Copy(os.Stdout, file); err != nil {
	// 		fmt.Fprintln(os.Stderr, err)
	// 		continue
	// 	}
		
	// 	file.Close()
	// }
	// When we run go run . *.txt > d.txt, it will copy the contents of all the text files in the current directory to d.txt

	// Now we'll move to another program that will calculate the file size, we'll use the same code as above
	// for _, fname := range os.Args[1:] {
	// 	file, err := os.Open(fname)

	// 	if err != nil {
	// 		fmt.Fprintln(os.Stderr, err)
	// 		continue
	// 	}

	// 	data, err :=ioutil.ReadAll(file)

	// 	if err != nil {
	// 		fmt.Fprintln(os.Stderr, err)
	// 		continue
	// 	}

	// 	fmt.Println("The file has", len(data), "bytes")
		
	// 	file.Close()
	// }

	// ReadAll reads the entire file into memory, this is not a good idea for large files
		// but for small files, it's fine. If the file is too large, it will cause a memory overflow
		// and the program will crash.
		// So, we need to use a buffer to read the file in chunks
		// buf := make([]byte, 1024) // 1KB buffer
		// for {
		// 	n, err := file.Read(buf)
		// 	if err == io.EOF {
		// 		break
		// 	}
		// 	if err != nil {
		// 		fmt.Fprintln(os.Stderr, err)
		// 		break
		// 	}
		// 	fmt.Print(string(buf[:n]))
		// }

	// Now we'll move to another program that will calculate the number of lines, words and characters in a file
	// The program will take the file name as a command line argument
	// and will print the number of lines, words and characters in the file, also the file name

	// So this program is similar to UNIX WC Utility, that is used to count the number of lines, words and characters in a file
	// wc -l filename.txt will give the number of lines in the file
	// wc -w filename.txt will give the number of words in the file
	// wc -c filename.txt will give the number of characters in the file
	// wc filename.txt will give the number of lines, words and characters in the file and also the total number of lines, words and characters in all the files
	for _, fname := range os.Args[1:] {

		var lc, wc, cc int

		file, err := os.Open(fname)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		scan := bufio.NewScanner(file)

		for scan.Scan() {
			s := scan.Text()

			wc += len(strings.Fields(s))
			cc +=len(s)
			lc++
		}

		fmt.Printf("%7d lines %7d words %7d bytes Filename:%s\n", lc, wc, cc, fname)

		file.Close()
	}

	// So we'll modify the code to combine all text file, one more which is the total of all files
	// we'll make it similar to the UNIX wc command

	fmt.Printf("\n\n\n")
	fmt.Println("The total number of lines, words and characters in all the files is:")
	var totalLines, totalWords, totalChars int

	for _, fname := range os.Args[1:] {
		var lc, wc, cc int

		file, err := os.Open(fname)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		scan := bufio.NewScanner(file)
		for scan.Scan() {
			s := scan.Text()
			wc += len(strings.Fields(s))
			cc += len(s)
			lc++
		}

		fmt.Printf("%7d lines %7d words %7d bytes Filename:%s\n", lc, wc, cc, fname)

		totalLines += lc
		totalWords += wc
		totalChars += cc

		file.Close()
	}

	fmt.Printf("\n%7d lines %7d words %7d bytes Total\n", totalLines, totalWords, totalChars)
}