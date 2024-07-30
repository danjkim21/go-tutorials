package main

import ( "fmt" ) 

// Structs and Types 
type gasEngine struct {
	mph uint8
	gallons uint8
  company
}

type electricEngine struct {
	mpkwh uint8
	kwh uint8
	company
}

type company struct {
	name string
}

// Methods
func (engine gasEngine) milesLeft() uint8 {
	return engine.mph * engine.gallons
}

func (engine electricEngine) milesLeft() uint8 {
	return engine.mpkwh * engine.kwh
}
	

func main() {

	var myEngine gasEngine
	var myOtherEngine gasEngine = gasEngine{mph: 10, gallons: 10, company: company{name: "Chevy"}}
	var myEvEngine electricEngine = electricEngine{mpkwh: 60, kwh: 6, company: company{name: "Ford"}}

	fmt.Println(myEngine, myEvEngine)
	fmt.Printf("miles left: %v\n", myOtherEngine.milesLeft())
	fmt.Printf("miles left: %v\n", myEvEngine.milesLeft())

}