package bitsort

import (
	"bufio"
	"io"
	"strconv"
)

const (
	MAX_NUMBER        = 10000
	BITFLAG_SIZE      = (MAX_NUMBER / 64) + 1 // 64bit
	WORKING_MEM_BYTES = 1024
)

func BitSort(input io.Reader, output io.Writer) error {
	bitflags := make([]uint64, BITFLAG_SIZE)
	reader := bufio.NewReaderSize(input, WORKING_MEM_BYTES)
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
	for i = 0; i < MAX_NUMBER; i++ {
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
