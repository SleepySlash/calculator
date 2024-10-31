package calculate

import (
	"fmt"
	"log"
	"testing"
)

func TestingCalculator(t *testing.T) {

	log.Println("Unit testing for the calculator app")


	t.Run("Conversion of input expression",func (t *testing.T){
		exp1 := makeNewInputExpression("0000","55","+")
		exp2 := makeNewInputExpression("   ","55","x")
		exp3 := makeNewInputExpression("  ","   ","  ")
		exp4 := makeNewInputExpression("9223372036854775807","-9223372036854775808","-")

		var mathExp expression
		
		log.Println("expression one, with multiple zeros")
		exp,err := convertInputExpression(exp1)
		if err!=nil{
			log.Fatal(err)
		}
		mathExp = exp.(expression)
		log.Printf("%v",mathExp)

		log.Println("expression two, with one empty input")
		exp,err = convertInputExpression(exp2)
		if err!=nil{
			log.Fatal(err)
		}
		mathExp = exp.(expression)
		log.Printf("%v",mathExp)
		
		log.Println("expression three, with all inputs being empty")
		exp,err = convertInputExpression(exp3)
		if err!=nil{
			log.Fatal(err)
		}
		mathExp = exp.(expression)
		log.Printf("%v",mathExp)

		log.Println("expression four, with largest possible operand values of int64")
		exp,err = convertInputExpression(exp4)
		if err!=nil{
			log.Fatal(err)
		}
		mathExp = exp.(expression)
		log.Printf("%v",mathExp)
	})


	t.Run("Addition of 2 numbers including postive and negative integers",func(t *testing.T) {
		
	})
}

func makeNewInputExpression(op1,op2,operator string) InputExpression{
	return InputExpression{
		Oparend1: op1,
		Oparend2: op2,
		Operator: operator,
	}
}