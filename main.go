package main
import (
	"fmt"
)

func main(){
	err:= ArgsNumberValidation()   // Validate the number of Args or provide a use example if "-h"
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
	
	input:= GetInput()
 
	res:= Parse(input, lookup)
	fmt.Println(res)
 
	

}





