package main

import (
	"encoding/json" // пакет используется для проверки ответа, не удаляйте его
	"fmt"           // пакет используется для проверки ответа, не удаляйте его
	
	
)

func main() {
	groupCity:= map[int][]string{
		10: []string{"Чернозём", "ЖОПА", "JOPA"},
		100: []string{"1","2","3"},
		1000: []string{"Moscow","St.Petersburg","KRAS"},
	}
   
   //Население городов в тысячах человек:
	 cityPopulation:= map[string]int{
		"1": 101,
		"2": 101,
		"3": 101,
		"4": 101,
		"5": 101,
	 }
 k:= false
 //fmt.Println(cityPopulation)
		for id,_:=range cityPopulation{
			for _, val:= range groupCity[100]{
			if val == id{
                k = true
			}
		}
		if !k{
			delete(cityPopulation,id)
		}
		k = false
	 }
 //  fmt.Print(cityPopulation)
    fmt.Printf("%T",func(fn func(i int) int, i int) func(int) int { return fn }(func(i int) int { return i + 1 }, 5))
	 

	value1, value2, operation := readTask() // исходные данные получаются с помощью этой функции
	// все полученные значения имеют тип пустого интерфейса
	var i []interface{}
i = append(i,value1)
i = append(i,value2)
i = append(i,operation)
val1, ok1:= i[0].(float64)
val2, ok2:= i[1].(float64)
val3, ok3:= i[2].(string)
if (ok1 && ok2 && ok3){
switch val3 {
case "+":
fmt.Printf("%.4f",val1 + val2)
case "-": 
fmt.Printf("%.4f",val1 - val2)
case "*": 
fmt.Printf("%.4f",val1 * val2)
case "/":  
fmt.Printf("%.4f",val1/val2)
default:
fmt.Print("неизвестная операция")
}

} else if ok1 == false{
fmt.Print("value=",val1,": ",i[0])
} else if ok2 == false{
fmt.Print("value=",val2,": ",i[1])
} else if ok3 == false{
fmt.Print("value=",val3,": ",i[2])
}






}


