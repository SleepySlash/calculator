package calculate

import (
	"testing"
)

func TestCalculator(t *testing.T) {

	t.Log("Unit testing for the calculator app")


	t.Run("Conversion of input expression",func (t *testing.T){
		exp1 := makeNewInputExpression("0000","55","+")
		exp2 := makeNewInputExpression("   ","55","x")
		exp3 := makeNewInputExpression("  ","   ","  ")
		exp4 := makeNewInputExpression("9223372036854775807","-9223372036854775808","-")
		
		var mathExp expression
		
		t.Log("expression one, with multiple zeros")
		exp,err := convertInputExpression(exp1)
		if err!=nil{
			t.Error(err)
			}else{
				mathExp = exp.(expression)
			t.Logf("%v",mathExp)
		}
		
		t.Log("expression four, with largest possible operand values of int64")
		exp,err = convertInputExpression(exp4)
		if err!=nil{
			t.Error(err)
		}else{
			mathExp = exp.(expression)
			t.Logf("%v",mathExp)
		}
		t.Log("expression three, with all inputs being empty")
		exp,err = convertInputExpression(exp3)
		if err!=nil{
			t.Error(err)
		}else{
			mathExp = exp.(expression)
			t.Logf("%v",mathExp)
		}

		
		t.Log("expression two, with one empty input")
		exp,err = convertInputExpression(exp2)
		if err!=nil{
			t.Error(err)
		}else{
			mathExp = exp.(expression)
			t.Logf("%v",mathExp)
		}
		
	})

	test1 := createNewMathExpression(-9223372036854775808,9223372036854775807,"+");
	test2 := createNewMathExpression(1,9223372036854775807,"+");
	test3 := createNewMathExpression(0,0,"/");
	test4 := createNewMathExpression(1000000000,999,"x");
	test5 := createNewMathExpression(9223372036854775807,9223372036854775807,"-")
	test6 := createNewMathExpression(-4611686018427387904,2,"*")


	t.Run("Addition of 2 numbers including postive and negative integers",func(t *testing.T) {
		result,err := theMather(test1.opr1,test1.opr2,test1.operator)
		if err!=nil{
			t.Error("the error in adding the limits in range of int64 : ",err)
		}
		t.Log("the result of adding the limits in range of int64 : ",result)
	})

	t.Run("Addition of 2 numbers which results a number out of 64 range ",func(t *testing.T) {
		result,err := theMather(test2.opr1,test2.opr2,test2.operator)
		if err!=nil{
			t.Error("the error in addition where result is a number out of int64 range : ",err)
		}
		t.Log("the result of addition where result is a number out of int64 range : ",result)
	})
	
	t.Run("Zero division error ",func(t *testing.T) {
		result,err := theMather(test3.opr1,test3.opr2,test3.operator)
		if err!=nil{
			t.Error("the error in dividing with zero : ",err)
		}
		t.Log("the result of zero division : ",result)
	})
	
	t.Run("Muliplication of 2 numbers which results a number out of 64 range ",func(t *testing.T) {
		result,err := theMather(test4.opr1,test4.opr2,test4.operator)
		if err!=nil{
			t.Error("the error in multiplication where result is a number out of int64 range : ",err)
		}
		t.Log("the result of multiplication where result is a number out of int64 range : ",result)
	})
	
	t.Run("Substraction of 2 numbers which are int min and int max ",func(t *testing.T) {
		result,err := theMather(test5.opr1,test5.opr2,test5.operator)
		if err!=nil{
			t.Error("the error in substraction of int max and int min : ",err)
		}
		t.Log("the result of substraction of int max and int min : ",result)
	})
	
	t.Run("Mulitplication of half of int min with 2",func(t *testing.T) {
		result,err := theMather(test6.opr1,test6.opr2,test6.operator)
		if err!=nil{
			t.Error("the error in mulitplication of half of int min with 2 : ",err)
		}
		t.Log("the result of mulitplication of half of int min with 2 : ",result)
	})
}

func makeNewInputExpression(op1,op2,operator string) InputExpression{
	return InputExpression{
		Oparend1: op1,
		Oparend2: op2,
		Operator: operator,
	}
}