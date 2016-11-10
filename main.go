package main

import (
	"fmt"
	"log"
	"os"

	"github.com/regattebzh/gribParser/grib"
)

func main() {
	path := "gfs.t00z.pgrb2.1p00.f000"

	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Error while opening file", err)
	}

	defer file.Close()

	fmt.Printf("%s opened\n", path)

	if isGrib, err := grib.CheckFileSignature(file); !isGrib || err != nil {
		return
	}

	if section0, err := grib.ReadSection0(file); err == nil {
		fmt.Printf("Section0:\n%+v\n", section0)
	}

	data, section, err := grib.ReadSection(file)
	fmt.Printf("Found section%d\n%+v\nLength:%d\n", section.SectionType, section, len(data))

	switch section.SectionType {
	case 0:
		fmt.Printf("Too late to read such section\n")
	case 1:
		currentSection := grib.Section1{}
		if _, err := currentSection.Write(data); err != nil {
			log.Fatal("Error on section1\n", err)
		}
		fmt.Printf("%+v\n", currentSection)
	default:
		fmt.Printf("Section%d reader not implemented\n", section.SectionType)
	}

}
