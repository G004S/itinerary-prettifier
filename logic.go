package main
import (
		/* "os" */
		"time"
		"strings"
		/* "fmt" */
	)

func Parse(input string, lookup [][]string) (string, map[string]int){
	var result string
	runeInput := []rune(input)
	var stats = map[string]int{
		"AirportsTotal":0,
		"IATA":0,
		"ICAO":0,
		"Dates":0,
		"T12":0,
		"T24":0,
	}

	for i:=0 ; i<len(runeInput); i++ {

		//potential airport code
		if runeInput[i] == '#'{
			var citymark bool
			if runeInput[i-1] == '*'{
					citymark = true
				}
			//potential ICAO code
			if runeInput[i+1] == '#'{
				airportCode := runeInput[i+2:i+6]
			 	airportName := GetAirportName(string(airportCode), lookup, citymark)
				if airportName == ""{
					airportName = "##"+string(airportCode)
				}
				if citymark{
					result = result[:len(result)-1]
				}
				result += G+string(airportName)+D
				stats["AirportsTotal"]++
				stats["ICAO"]++
				i+=5
				continue
			}
			//IATA code
			airportCode := runeInput[i+1:i+4]
		 	airportName := GetAirportName(string(airportCode), lookup, citymark)
			if citymark{
				result = result[:len(result)-1]
			}
			result += G+string(airportName)+D
			stats["AirportsTotal"]++
			stats["IATA"]++
			i+= 3
			continue
		}

		//potential date
		if runeInput[i] == 'D' && runeInput[i+1] == '('{
			date:= runeInput[i:i+19]
			parsedDate:= ParseDate(date)
			result += parsedDate
			stats["Dates"]++
			i+=19
			continue
		}
		//potential T12
		if runeInput[i] == 'T' && runeInput [i+1] == '1' && runeInput [i+2] == '2'{
			var time12 []rune
			for j:=0; j<len(runeInput)-i ;j++{
				if runeInput[i+j] == ')'{
					time12= append(time12,runeInput[i+j])
					break
				}
				time12= append(time12,runeInput[i+j])
			}
			parsedT12:= ParseT12(time12)
			if parsedT12[0]!='T'{
				stats["T12"]++
			}
			result += parsedT12
			i+= len(time12)-1
			continue
		}
		//potential T24
		if runeInput[i] == 'T' && runeInput [i+1] == '2' && runeInput [i+2] == '4'{
			var time24 []rune
			for k:=0; k<len(runeInput)-i ;k++{
				if runeInput[i+k] == ')'{
					time24= append(time24,runeInput[i+k])
					break
				}
				time24= append(time24,runeInput[i+k])
			}
			parsedT24:= ParseT24(time24)
			if parsedT24[0]!='T'{
				stats["T24"]++
			}
			result += parsedT24
			i+= len(time24)-1
			continue

		}
		//Convert line-break charachters
		if runeInput[i] == '\v' || runeInput[i] == '\f' || runeInput[i] == '\r'{
			result += "\n"
			i++
			continue
		}

		//remove blank space
		if runeInput[i] == '\n'{
			for l:=0; l<len(runeInput)-i; l++{
				if runeInput[i+l] != '\n'{
					if l>1{
						i+=l-2
						break
					}
				}else if i+l==len(runeInput)-1{
					return result, stats
				}
			}
		}

		result += string(runeInput[i])
	}
	return result, stats
}



func GetAirportName(airportCode string, lookup [][]string, citymark bool) string{
	var match string
	for _,row := range lookup {
		for _,cell:= range row{
			if airportCode == cell{
				match = row[0]
				if citymark{
					if len(match)>24{
						if match[len(match)-21:] == "International Airport"{   //Check if it's "Airport" or "International Airport" and trim it
						match = match[:len(match)-22]
					}else if match[len(match)-7:] == "Airport"{
						match = match[:len(match)-8]
					}
					}else{
						if match[len(match)-7:] == "Airport"{
						match = match[:len(match)-8]
						}
					}	
				}
			}
		}
	}
	return match
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

	res+= R+string(date[10:12])+" "+m+" "+string(date[2:6])+D
	return res
	
}

func ParseT12(time12 []rune) string{

	raw:= strings.TrimPrefix(string(time12), "T12(")
	raw= strings.TrimSuffix(raw, ")")
	if rune(raw[len(raw)-1]) == 'Z'{
		parsedT12,err:= time.Parse("2006-01-02T15:04Z", raw)
		if err != nil{
			return string(time12)
		}
		formatedT12:= parsedT12.Format("03:04PM (-07:00)")
		formatedT12 = formatedT12[:8]+Y+formatedT12[8:]+D
		return formatedT12
	}else{
		parsedT12,err:= time.Parse("2006-01-02T15:04-07:00", raw)
		if err != nil{
			return string(time12)
		}

		formatedT12:= parsedT12.Format("03:04PM (-07:00)")
		formatedT12 = formatedT12[:8]+Y+formatedT12[8:]+D
		return formatedT12
	}
}

func ParseT24(time24 []rune) string{
	raw:= strings.TrimPrefix(string(time24), "T24(")
	raw= strings.TrimSuffix(raw, ")")
	if rune(raw[len(raw)-1]) == 'Z'{
		parsedT24,err:= time.Parse("2006-01-02T15:04Z", raw)
		if err != nil{
			return string(time24)
		}
		formatedT24:= parsedT24.Format("15:04 (-07:00)")
		formatedT24 = formatedT24[:6]+Y+formatedT24[6:]+D
		return formatedT24
	}else{
		parsedT24,err:= time.Parse("2006-01-02T15:04-07:00", raw)
		if err != nil{
			return string(time24)
		}

		formatedT24:= parsedT24.Format("15:04 (-07:00)")
		formatedT24 = formatedT24[:6]+Y+formatedT24[6:]+D
		return formatedT24
	}
}
