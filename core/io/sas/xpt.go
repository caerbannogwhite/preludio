package sasio

import (
	"bufio"
	"fmt"
	"os"
)

const FIRST_HEADER = "HEADER RECORD*******LIBRARY HEADER RECORD!!!!!!!000000000000000000000000000000"

// This functions writes a SAS XPT file (versions 5/6).
func ReadXPTv56(path string) error {

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	stats, err := file.Stat()
	if err != nil {
		return err
	}

	var size int64 = stats.Size()
	bytecode := make([]byte, size)

	bufr := bufio.NewReader(file)
	_, err = bufr.Read(bytecode)
	if err != nil {
		return err
	}

	err = Unpackv56(bytecode)
	if err != nil {
		return err
	}

	return nil
}

// This functions writes a SAS XPT file (versions 5/6).
func WriteXPTv56(path string) error {
	return nil
}

// This functions writes a SAS XPT file (versions 7/8).
func ReadXPTv78(path string) error {
	return nil
}

// This functions writes a SAS XPT file (versions 7/8).
func WriteXPTv78(path string) error {
	return nil
}

func Unpackv56(bytecode []byte) error {

	fmt.Println(string(bytecode))

	return nil
}
