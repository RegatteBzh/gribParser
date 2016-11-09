package main

import (
	"fmt"
	"log"
	"os"

	"github.com/RegatteBzh/gribParser/grib"
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

	if section1, err := grib.ReadSection1(file); err == nil {
		fmt.Printf("Section1:\n%+v\n", section1)
	}

}
