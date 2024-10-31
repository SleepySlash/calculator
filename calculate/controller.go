package calculate

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
)

type InputExpression struct{
	Oparend1 string    `json:"oparend_1"`
	Oparend2 string    `json:"oparend_2"`
	Operator string `json:"operator"`
}

type expression struct{
	opr1 int64;
	opr2 int64;
	operator string;
}

func Calculator(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type","Application-json")
	w.WriteHeader(http.StatusContinue)
	w.Write([]byte("This is calculator app"))
}

func Calculate(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","Application-json")
	
	var exp InputExpression
	err := json.NewDecoder(r.Body).Decode(&exp)
	if(err!=nil){
		log.Fatal("error in the calucalte function",err)
	}

	operandsInterface, err := convertInputExpression(exp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	operands := operandsInterface.(expression)
	afterMath,err := theMather(operands.opr1,operands.opr2,exp.Operator)

	if err!=nil{
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
	}else{
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(afterMath)
	}
}


// converting the input expression into a mathematically accurate expression
func convertInputExpression(exp InputExpression) (interface{},error){

	okay,err:= checkInputExpression(&exp)
	if err!=nil || !okay{
		return nil,err
	}

	int1,err := strconv.ParseInt(exp.Oparend1,10,64)
	if err!=nil{
		return nil,err
	}
	int2,err := strconv.ParseInt(exp.Oparend2,10,64)
	if err!=nil{
		return nil,err
	}
	operands :=createNewMathExpression(int1,int2,exp.Operator)

	return operands,nil
}

// checking the correctness of the expression requested by the user
func checkInputExpression(exp *InputExpression) (bool,error){
	exp.Oparend1=strings.ReplaceAll(exp.Oparend1," ","")
	exp.Oparend2=strings.ReplaceAll(exp.Oparend2," ","")
	exp.Operator=strings.ReplaceAll(exp.Operator," ","")

	if exp.Oparend1=="" || exp.Oparend2 =="" || len(exp.Operator)!=1 || exp.Operator==""{
		return false,errors.New("BAD EXPRESSION")
	}
	return true,nil
}

// creating a new expression with integer type operands
func createNewMathExpression(a, b int64, c string) *expression{
	return &expression{
		opr1 : a,
		opr2 : b,
		operator : c,
	}
}

// the function that carries out the math operations 
func theMather(op1,op2 int64,opr string) (string,error){
	switch opr {
	case "+":
		if (op1 > 0 && op2 > 0 && op1+op2 < 0) || (op1 < 0 && op2 < 0 && op1+op2 > 0) {
			res := fmt.Sprintf("%.5f", float64(op1)+float64(op2))
			return res, nil
		}
		res := fmt.Sprintf("%d", op1+op2)
		return res, nil

	case "-":
		if (op2 > 0 && op1 < math.MinInt64+op2) || (op2 < 0 && op1 > math.MaxInt64+op2) {
			res := fmt.Sprintf("%.5f", float64(op1)-float64(op2))
			return res, nil
		}
		res := fmt.Sprintf("%d", op1-op2)
		return res, nil

	case "*", "x":
		if op1 != 0 && op2 != 0 && (op1 > math.MaxInt64/op2 || op1 < math.MinInt64/op2) {
			res := fmt.Sprintf("%.5f", float64(op1)*float64(op2))
			return res, nil
		}
		res := fmt.Sprintf("%d", op1*op2)
		return res, nil

	case "/":
		if op2 == 0 {
			return "", errors.New("cannot divide by zero")
		}
		res := fmt.Sprintf("%.5f", float64(op1)/float64(op2))
		return res, nil

	default:
		return "", errors.New("operation does not exist or is unsupported")
	}
}