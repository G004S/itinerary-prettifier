package main

import(
	"os"
	"errors"
	"fmt"
)

func GetInput() string{
	input,err := os.ReadFile(os.Args[1])
	
	if err != nil {
		return ""
	}
	return string(input)
}


func WriteOutput(output string) error{
	f,err := os.Create("./output.txt")
	defer f.Close()
	if err != nil{
		return errors.New("ERROR: Couldn't create the output file")
	}

	_,err = f.WriteString(output)
	if err != nil{
		return errors.New("ERROR: Couldn't write to the output file")
	}
	return nil
}

func ShowStats(stats map[string]int){
	strStat:=fmt.Sprintf("-------Parsing statistics-------\n\nAirportsTotal: %d\nIATA codes: %d\nICAO codes: %d\nDates: %d\nT12: %d\nT24: %d", stats["AirportsTotal"],stats["IATA"],stats["ICAO"],stats["Dates"],stats["T12"],stats["T24"])
	fmt.Println(strStat)
}