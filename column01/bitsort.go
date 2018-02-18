package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("bitsort_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// read and sort
	bitflags := make([]uint64, 256)
	reader := bufio.NewReaderSize(file, 1024)
	for {
		read, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		}
		numberStr := string(read[:len(read)-1]) // strip '\n'
		number, err := strconv.ParseUint(numberStr, 10, 32)
		if err != nil {
			log.Fatal(err)
		}

		sliceIdx := number / 64
		bitIdx := number % 64
		bitflags[sliceIdx] = bitflags[sliceIdx] | (1 << bitIdx)
	}

	var i uint64
	for i = 0; i < 64*256; i++ {
		sliceIdx := i / 64
		bitIdx := i % 64
		if ((bitflags[sliceIdx] >> bitIdx) & 1) == 1 {
			fmt.Printf("%d\n", i)
		}
	}
}
