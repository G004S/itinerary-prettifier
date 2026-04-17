package main
import (
	"fmt"
)

func main(){
	statflag, err:= ArgsNumberValidation()   // Validate the number of Args or provide a use example if "-h"
	if err != nil{
		fmt.Println("itinerary usage:\ngo run . ./input.txt ./output.txt ./airport-lookup.csv")
		return
	}

	err = ArgsPathValidation()   // Validate the existance of the required files
	if err != nil{
		fmt.Println(err)
		return
	}


	lookup,err := LookupFileValidation() // Validate columns in the lookup file
	if err != nil{
		fmt.Println(err)
		return
	}
	
	input:= GetInput() //Read the data out of the input file

	res, stats:= Parse(input, lookup) 	//Execute parsing (Airport names, correct data and time layouts)
	fmt.Println(res)
	
	if statflag{
		ShowStats(stats)
	}
	WriteOutput(res) //Write parsed data into the output.txt file
	

}





