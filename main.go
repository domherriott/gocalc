package main

import (
   "fmt"
   "strconv"
)

var mem float64
var op string

func string_to_float(s string) float64 {
   f, err := strconv.ParseFloat(s, 8)
   _ = err
   return f
}

func apply_operator(op string, v1 float64, v2 float64) float64 {
   var v3 float64
	    
   switch op {
      case "+":
         v3 = v1 + v2
      case "-":
         v3 = v1 - v2
      case "*":
         v3 = v1 * v2
      case "/":
         v3 = v1 / v2
   }

   return v3
}

func main() {

   for true {	
      var input string
      fmt.Scan(&input)

      switch input {
	 case "clear":
            op = ""
            mem = 0
         case "=":
	    fmt.Println(mem)
         case "+", "-", "*", "/":
            op = input
    	 default:
       	    if op == "" {
	       mem = string_to_float(input)
            } else {
	       mem = apply_operator(op, mem, string_to_float(input))
	    }
      }	
      	
  }
}
