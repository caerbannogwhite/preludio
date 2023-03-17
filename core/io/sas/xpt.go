package sasio

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"time"
)

///////////////////////////////////////     SAS XPT v5/6     ///////////////////////////////////////

// Technical documentation:
// https://support.sas.com/content/dam/SAS/support/en/technical-papers/record-layout-of-a-sas-version-5-or-6-data-set-in-sas-transport-xport-format.pdf
const FIRST_HEADER_V56 = "HEADER RECORD*******LIBRARY HEADER RECORD!!!!!!!000000000000000000000000000000"

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

func Unpackv56(bytecode []byte) error {

	fmt.Println(string(bytecode))

	return nil
}

// This functions writes a SAS XPT file (versions 5/6).
func WriteXPTv56(path string) error {
	buff := make([]byte, 0)

	buff = append(buff, []byte(fmt.Sprintf(
		"%s%8s%8s%8s%8s",
		FIRST_HEADER_V56,       // 1-80 HEADER RECORD
		"SAS", "SAS", "SASLIB", // 81-104
		time.Now().Format("ddMMMyy:hh:mm:ss")), // 105-128
	)...)

	// write buff to file
	os.WriteFile(path, buff, 0644)

	return nil
}

///////////////////////////////////////     SAS XPT v8/9     ///////////////////////////////////////

// Technical documentation:
// https://support.sas.com/content/dam/SAS/support/en/technical-papers/record-layout-of-a-sas-version-8-or-9-data-set-in-sas-transport-format.pdf

const FIRST_HEADER_V89 = "HEADER RECORD*******LIBV8 HEADER RECORD!!!!!!!000000000000000000000000000000"

// This functions writes a SAS XPT file (versions 7/8).
func ReadXPTv89(path string) error {
	return nil
}

func Unpackv89(bytecode []byte) error {

	fmt.Println(string(bytecode))

	return nil
}

// This functions writes a SAS XPT file (versions 7/8).
func WriteXPTv89(path string) error {

	buff := make([]byte, 0)

	buff = append(buff, []byte(fmt.Sprintf(
		"%s%8s%8s%8s%8s%24s%16s%80s",
		FIRST_HEADER_V89,                      // 1-80 		First header record
		"SAS",                                 // 81-88 	SAS
		"SAS",                                 // 89-96 	SAS
		"SASLIB  9.4",                         // 97-104 	SASLIB
		runtime.GOOS,                          // 105-128	OS Name
		"",                                    // 129-152 	24 blanks
		time.Now().Format("ddMMMyy:hh:mm:ss"), // 153-176   Date/time created
		time.Now().Format("ddMMMyy:hh:mm:ss"), // 177-200 	Second header record, date/time modified
	))...)

	// write buff to file
	os.WriteFile(path, buff, 0644)

	return nil
}
