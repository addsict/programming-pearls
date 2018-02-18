package bitsort

import (
	"bufio"
	"io"
	"strconv"
)

func BitSort(input io.Reader, output io.Writer) error {
	bitflags := make([]uint64, 256)
	reader := bufio.NewReaderSize(input, 1024)
	for {
		// read
		read, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}
		numberStr := string(read[:len(read)-1]) // strip '\n'
		number, err := strconv.ParseUint(numberStr, 10, 32)
		if err != nil {
			return err
		}

		// sort
		sliceIdx := number / 64
		bitIdx := number % 64
		bitflags[sliceIdx] = bitflags[sliceIdx] | (1 << bitIdx)
	}

	// write
	var i uint64
	for i = 0; i < 64*256; i++ {
		sliceIdx := i / 64
		bitIdx := i % 64
		if ((bitflags[sliceIdx] >> bitIdx) & 1) == 1 {
			_, err := output.Write([]byte(strconv.FormatUint(i, 10) + "\n"))
			if err != nil {
				return err
			}
		}
	}

	return nil
}
