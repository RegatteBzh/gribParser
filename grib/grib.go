package grib

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"os"
)

//Section0 is the grib2 Section0
type Section0 struct {
	Reserved      uint16
	Discipline    uint8
	EditionNumber uint8
	Length        uint64
}

//Section1 is the grib2 Section1
type Section1 struct {
	GeneratingCenter    uint16
	GeneratingSubCenter uint16
	MasterTablesVersion uint8
	LocalTablesVersion  uint8
	ReferenceTime       uint8
	Year                uint16
	Month               uint8
	Day                 uint8
	Hour                uint8
	Minute              uint8
	Second              uint8
	ProductionStatus    uint8
	DataType            uint8
}

//Section2 is the grib2 Section2
type Section2 struct {
	Length        uint32
	SectionNumber uint8
}

// ReadSection0 read section0
func ReadSection0(file *os.File) (section Section0, err error) {
	section = Section0{}
	data, err := ReadNextBytes(file, binary.Size(section))

	buffer := bytes.NewBuffer(data)
	err = binary.Read(buffer, binary.LittleEndian, &section)
	if err != nil {
		log.Fatal("binary.Read failed (ReadSection0)", err)
	}

	section.Length = reverse64(section.Length)
	return
}

func (section *Section1) Write(p []byte) (size int, err error) {
	buffer := bytes.NewBuffer(p)

	if err = binary.Read(buffer, binary.LittleEndian, section); err != nil {
		log.Fatal("binary.Read failed (Section1.Write)", err)
	}
	size = len(p)

	section.GeneratingCenter = reverse16(section.GeneratingCenter)
	section.GeneratingSubCenter = reverse16(section.GeneratingSubCenter)
	section.Year = reverse16(section.Year)

	return
}

// CheckFileSignature checks if the file is a grib2
func CheckFileSignature(file *os.File) (result bool, err error) {
	formatName, err := ReadNextBytes(file, 4)
	fmt.Printf("Parsed format: %s\n", formatName)

	result = (string(formatName) == "GRIB")
	return
}
