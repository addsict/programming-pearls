package bitsort

import (
	"bytes"
	"io"
	"strconv"
	"strings"
	"testing"
)

func TestBitSort(t *testing.T) {
	inputNumbers := []uint32{
		9999,
		1234,
		8765,
		1,
		3,
		5,
		14,
		103,
		199,
		83,
		23,
		1900,
		2876,
		158,
		688,
		9998,
		7654,
		331,
		15,
		16,
	}
	expectedNumbers := []uint32{
		1,
		3,
		5,
		14,
		15,
		16,
		23,
		83,
		103,
		158,
		199,
		331,
		688,
		1234,
		1900,
		2876,
		7654,
		8765,
		9998,
		9999,
	}

	var inputStr string
	for _, num := range inputNumbers {
		inputStr += (strconv.FormatUint(uint64(num), 10) + "\n")
	}
	input := strings.NewReader(inputStr)

	outputBytes := make([]byte, 0, 1024)
	output := bytes.NewBuffer(outputBytes)

	err := BitSort(input, output)
	if err != nil {
		t.Fatal(err)
	}

	gotNumbers := make([]uint32, 0, 1024)
	for {
		read, err := output.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				t.Fatal(err)
			}
		}
		numberStr := string(read[:len(read)-1]) // strip '\n'
		number, err := strconv.ParseUint(numberStr, 10, 32)
		if err != nil {
			t.Fatal(err)
		}
		gotNumbers = append(gotNumbers, uint32(number))
	}

	if len(gotNumbers) != len(expectedNumbers) {
		t.Errorf("invalid output size: expected=%d, got=%d", len(expectedNumbers), len(gotNumbers))
	}

	for i := 0; i < len(gotNumbers); i++ {
		if gotNumbers[i] != expectedNumbers[i] {
			t.Errorf("invalid sort at %d: expected=%d, got=%d", i, expectedNumbers[i], gotNumbers[i])
		}
	}
}
