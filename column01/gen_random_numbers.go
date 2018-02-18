package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	var n int
	var k int
	var out string

	flag.IntVar(&n, "n", 1000, "max number")
	flag.IntVar(&k, "k", 1000, "number of items")
	flag.StringVar(&out, "out", "", "output filename")
	flag.Parse()

	fmt.Printf("generate n=%d, k=%d\n", n, k)

	// generate n items
	numbers := make([]uint32, n)
	for i := 0; i < n; i++ {
		numbers[i] = uint32(i)
	}
	for i := 0; i < n; i++ {
		idx1 := rand.Intn(n)
		idx2 := rand.Intn(n)

		// swap
		tmp := numbers[idx1]
		numbers[idx1] = numbers[idx2]
		numbers[idx2] = tmp
	}

	// write first k items
	file, err := os.OpenFile(out, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	writer := bufio.NewWriter(file)
	for i := 0; i < k; i++ {
		writer.WriteString(strconv.FormatUint(uint64(numbers[i]), 10) + "\n")
	}
	err = writer.Flush()
	if err != nil {
		log.Fatal(err)
	}

	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
}
