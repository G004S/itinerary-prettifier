package main
import (
		"os"
		"time"
		"strings"
		"fmt"
	)

func GetInput() string{
	input,err := os.ReadFile(os.Args[1])
	
	if err != nil {
		return ""
	}

	return string(input)
}



func Parse(input string, lookup [][]string) string{
	var result string
	runeInput := []rune(input)

	for i:=0 ; i<len(runeInput); i++ {

		//potential airport code
		if runeInput[i] == '#'{
			//potential ICAO code
			if runeInput[i+1] == '#'{
				airportCode := runeInput[i+2:i+6]
			 	airportName := GetAirportName(string(airportCode), lookup)
				if airportName == ""{
					airportName = "##"+string(airportCode)
				}
				result += string(airportName) 
				i+=5
				continue
			}
			airportCode := runeInput[i+1:i+4]
		 	airportName := GetAirportName(string(airportCode), lookup)
			result += string(airportName) 
			i+= 3
			continue
		}

		//potential time
		if runeInput[i] == 'D' && runeInput[i+1] == '('{
			date:= runeInput[i:i+19]
			parsedDate:= ParseDate(date)
			result += parsedDate
			i+=19
			continue
		}

		if runeInput[i] == 'T' && runeInput [i+1] == '1' && runeInput [i+2] == '2'{
			time12:= runeInput[i:i+25]
			parsedT12:= ParseT12(time12)
			result += parsedT12
			i+=27
			continue
		}

		result += string(runeInput[i])
	}
	return result
}




func GetAirportName(airportCode string, lookup [][]string) string{
	for _,row := range lookup {
		for _,cell:= range row{
			if airportCode == cell{
				return row[0]
			}
		}
	}
	return ""
}


func ParseDate(date []rune) string{
	var res string
	months:= map[string]string{
		"01": "January",
		"02": "February",
		"03": "March",
		"04": "April",
		"05": "May",
		"06": "June",
		"07": "July",
		"08": "August",
		"09": "September",
		"10": "October",
		"11": "November",
		"12": "December",
	}
	m,ok:= months[string(date[7:9])]
	if !ok{
		return string(date)
	}

	res+= string(date[10:12])+" "+m+" "+string(date[2:6])
	
	return res
	
}

func ParseT12(time12 []rune) string{
	fmt.Println(string(time12))
	fmt.Println("")
	fmt.Println("")
	raw:= strings.TrimPrefix(string(time12), "T12(")
	raw= strings.TrimSuffix(raw, ")")
	parsedT12,_:= time.Parse(time.RFC3339, raw)

	return parsedT12.Format("03:04PM (-07:00)")
}