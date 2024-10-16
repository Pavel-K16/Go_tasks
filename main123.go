package main 

import (
"fmt"
)

type Human_ struct{
age int32 
weight float32
name string
}


func (s *Human_) set_val(){
s.age = 0
s.weight = 3.5
s.name = "child" 
} 

func (s *Human_) Set_val(){
	s.age = 24
	s.weight = 85
	s.name = "Pavel" 
	}


func main(){
	var a = Human_ {23,85, "Pasha"}
	b := Human_{weight: 90, age: 24, name: "Pavel"}
	fmt.Println(a, b)

mas_h:= make([]Human_,10,10)

for i:=0; i < len(mas_h); i++{
	if (i%2 == 0){
	a.set_val()
	} else {
		a.Set_val()
	}

	mas_h[i] = a  
}

 str:="Всем привет! Меня зовут Паша!"
 rs:=[]rune(str)
fmt.Print(str, "\n")
fmt.Print(str[0]," ", len(str), len(rs),"\n")
fmt.Print(rs, "\n")

for _, val:= range str{
	fmt.Print(val, " ")
} 

var age,j,k  = 12,13,"wfwe" 

fmt.Println("werfwsefwse" , age, "\n")
var happiness = 0.04 
var isHappy = true 
if happiness >= 0.5 || isHappy { 
 fmt.Println("Happy ") 
}
fmt.Println(age,j,k)
const pi = 3.13
}