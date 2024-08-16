package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Calc struct {
	Num1 float64 `json:"num1"`
	Num2 float64 `json:"num2"`
}

func (c *Calc) add() string {
	c.log("+")
	return strconv.FormatFloat(c.Num1+c.Num2, 'f', -1, 64)
}
func (c *Calc) subtract() string {
	c.log("-")
	return strconv.FormatFloat(c.Num1-c.Num2, 'f', -1, 64)
}
func (c *Calc) multiply() string {
	c.log("*")
	return strconv.FormatFloat(c.Num1*c.Num2, 'f', -1, 64)
}
func (c *Calc) divide() string {
	c.log("/")
	return strconv.FormatFloat(c.Num1/c.Num2, 'f', -1, 64)
}

func (c *Calc) log(op string) {
	log.Printf("Calculation used %v %v %v", c.Num1, op, c.Num2)
}

func Add(w http.ResponseWriter, req *http.Request) {
	var rBody Calc
	json.NewDecoder(req.Body).Decode(&rBody)
	fmt.Fprintf(w, "Calculation: %v + %v\nAnswer: %v", rBody.Num1, rBody.Num2, rBody.add())
}

func Subtract(w http.ResponseWriter, req *http.Request) {
	var rBody Calc
	json.NewDecoder(req.Body).Decode(&rBody)
	fmt.Fprintf(w, "Calculation: %v - %v\nAnswer: %v", rBody.Num1, rBody.Num2, rBody.subtract())
}

func Divide(w http.ResponseWriter, req *http.Request) {
	var rBody Calc
	json.NewDecoder(req.Body).Decode(&rBody)
	fmt.Fprintf(w, "Calculation: %v / %v\nAnswer: %v", rBody.Num1, rBody.Num2, rBody.divide())
}

func Multiply(w http.ResponseWriter, req *http.Request) {
	var rBody Calc
	json.NewDecoder(req.Body).Decode(&rBody)
	fmt.Fprintf(w, "Calculation: %v x %v\nAnswer: %v", rBody.Num1, rBody.Num2, rBody.multiply())
}
