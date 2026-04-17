package main
import (
	"fmt"
	"os"
	"errors"
	"strings"
	"encoding/csv"
	"io"
)


func ArgsNumberValidation() (bool, error) {
	if len(os.Args) > 1 && os.Args[1] == "-h" {
			fmt.Println("itinerary usage:\ngo run . ./input.txt ./output.txt ./airport-lookup.csv")
			os.Exit(0)
	}
	if len(os.Args) == 5 && os.Args[4] == "--stats"{
		return true, nil
	}
	if len(os.Args) != 4{
		return false, errors.New("Invalid number of arguments")
	}
	return false, nil
}

func ArgsPathValidation() error {
	inputPath,_ := os.Stat(os.Args[1])
	outputPath,_ := os.Stat(os.Args[3]) 
	if inputPath == nil{
		return errors.New("input not found")
	}
	if outputPath == nil{
		return errors.New("Airport lookup not found")
	}
	return nil
}

func LookupFileValidation() ([][]string, error) {
	var lookupTable [][]string
	lookup,err := os.Open(os.Args[3])
	if err != nil {
		return nil, errors.New("Couldn't open the airport lookup file")
	}
	defer lookup.Close()

	reader:= csv.NewReader(lookup)
	header, err:= reader.Read()
	if err != nil{
		return nil, err
	}
	
	expectedHeader:= []string{"name","iso_country","municipality","icao_code","iata_code","coordinates"}
	for i,v := range header{
		if v != expectedHeader[i]{
			return nil, errors.New("Airport lookup malformed")
		}
	}


	for{
		row,err:= reader.Read()
		if err == io.EOF{
			break
		}
		if err != nil{
			return nil, errors.New("Airport lookup malformed")
		}
		if len(row) != 6 {
			return nil, errors.New("Airport lookup malformed")
		}

		for _,v := range row{
			if strings.TrimSpace(v) == ""{
				return nil, errors.New("Airport lookup malformed")
			}
		}
		
		lookupTable = append(lookupTable, row)
	}

	return lookupTable, nil
}

