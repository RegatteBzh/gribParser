package grib

import (
	"bytes"
	"encoding/binary"
	"io"
	"log"
)

//Section is the Template of all sections (except 0 and 8)
type Section struct {
	Length      uint32
	SectionType uint8
}

// ReadNextBytes read bytes from file
func ReadNextBytes(file io.Reader, number int) (bytes []byte, err error) {
	bytes = make([]byte, number)

	_, err = file.Read(bytes)
	if err != nil {
		log.Fatal(err)
	}

	return
}

func (section *Section) Write(p []byte) (size int, err error) {
	buffer := bytes.NewBuffer(p)

	if err = binary.Read(buffer, binary.LittleEndian, section); err != nil {
		log.Fatal("binary.Read failed (Section.Write)\n", err)
	}
	section.Length = reverse32(section.Length)
	size = len(p)

	return
}

// ReadSection read a section in the GRIB2 file
func ReadSection(file io.Reader) (data []byte, section Section, err error) {
	head := make([]byte, 5)
	data = make([]byte, 0)
	section = Section{}

	if _, err = file.Read(head); err != nil {
		log.Fatal("file.Read failed (ReadSection)\n", err)
		return
	}

	if _, err = section.Write(head); err != nil {
		log.Fatal("section.Write failed (ReadSection)\n", err)
		return
	}

	if section.Length > 5 {
		data = make([]byte, section.Length-5)
		if _, err = file.Read(data); err != nil {
			log.Fatal("file.Read failed (ReadSection)\n", err)
		}
	}

	return
}
