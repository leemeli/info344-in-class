package models

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

type Zip struct {
	Code  string
	City  string
	State string
}

type ZipSlice []*Zip

type ZipIndex map[string]ZipSlice

func LoadZips(fileName string) (ZipSlice, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}

	reader := csv.NewReader(f)
	_, err = reader.Read() // the underscore means "I don't care about this value"
	if err != nil {
		return nil, fmt.Errorf("error reading header row: %v", err)
	}

	zips := make(ZipSlice, 0, 43000) // One memory allocation and that's it by using the make method
	for {
		fields, err := reader.Read()
		if err == io.EOF { // reached end of file
			return zips, nil
		}
		if err != nil {
			return nil, fmt.Errorf("error reading record: %v", err)
		}
		// If a variable has a very short lifespan, use single letters/shorter names
		// That's the go style but dr. stearns doesn't mind either way
		z := &Zip{ // without the "&", you would have to do append(zips, &z) later on
			Code:  fields[0],
			City:  fields[3],
			State: fields[6],
		}
		zips = append(zips, z) // append what we have to zips
	}
}
