package grib

import (
	"bytes"
	"encoding/binary"
	"log"
	"os"
)

// ReadNextBytes read bytes from file
func ReadNextBytes(file *os.File, number int) (bytes []byte, err error) {
	bytes = make([]byte, number)

	_, err = file.Read(bytes)
	if err != nil {
		log.Fatal(err)
	}

	return
}

func readSection(file *os.File, number int, section interface{}) (err error) {
	data, err := ReadNextBytes(file, number)

	buffer := bytes.NewBuffer(data)
	err = binary.Read(buffer, binary.LittleEndian, section)
	if err != nil {
		log.Fatal("binary.Read failed", err)
	}

	return
}
