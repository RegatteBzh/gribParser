package grib

import (
	"encoding/binary"
	"fmt"
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
	Length              uint32
	SectionNumber       uint8
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
	err = readSection(file, binary.Size(section), &section)
	section.Length = reverse64(section.Length)
	return
}

// ReadSection1 read section1
func ReadSection1(file *os.File) (section Section1, err error) {
	section = Section1{}
	err = readSection(file, binary.Size(section), &section)
	section.GeneratingCenter = reverse16(section.GeneratingCenter)
	section.GeneratingSubCenter = reverse16(section.GeneratingSubCenter)
	section.Year = reverse16(section.Year)
	section.Length = reverse32(section.Length)
	offset := int64(section.Length) - int64(binary.Size(section))
	if offset > 0 {
		file.Seek(offset, 1)
	}
	return
}

// ReadSection2 read section2
func ReadSection2(file *os.File) (section Section2, err error) {
	section = Section2{}
	err = readSection(file, binary.Size(section), &section)
	section.Length = reverse32(section.Length)
	offset := int64(section.Length) - int64(binary.Size(section))
	if offset > 0 {
		file.Seek(offset, 1)
	}
	return
}

// CheckFileSignature checks if the file is a grib2
func CheckFileSignature(file *os.File) (result bool, err error) {
	formatName, err := ReadNextBytes(file, 4)
	fmt.Printf("Parsed format: %s\n", formatName)

	result = (string(formatName) == "GRIB")
	return
}
