package main

import (
	"cmp"
	"fmt"
	"math"

	"golang.org/x/exp/constraints"
)
func main(){
	myNanValue := math.NaN()
	var myFloatValue float64 = -20000

	// we can't perform comparations with NaN values
	if myNanValue < myFloatValue {
		fmt.Println("comparation with NaN")
	}
	if myNanValue > myFloatValue {
		fmt.Println("comparation with NaN")
	}
	if myNanValue == myFloatValue {
		fmt.Println("comparation with NaN")
	}

	// we can perform comparations with NaN values with Compare & Less methods
	fmt.Printf("Numbers (%f, %f) - Compare: %d\n", myNanValue,myFloatValue, cmp.Compare(myNanValue, myFloatValue) )

	fmt.Printf("Numbers (%f, %f) - Less: %t\n", myNanValue,myFloatValue, cmp.Less(myNanValue, myFloatValue) )

	fmt.Println()

	oldOrderedValues(2, 4)
	orderedValues(2,4)

	fmt.Println(cmp.Compare(3,2))
	fmt.Println(cmp.Less(1,2))
}


func oldOrderedValues[N constraints.Ordered](v1, v2 N){
	fmt.Println(v1, v2)
	fmt.Println(v1 !=v2)
	fmt.Println(v1 < v2)
	fmt.Println(v1 > v2)
}

func orderedValues[N cmp.Ordered](v1, v2 N){
	fmt.Println(v1, v2)
	fmt.Println(v1 !=v2)
	fmt.Println(v1 < v2)
	fmt.Println(v1 > v2)
}