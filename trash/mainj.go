package main

import ("fmt";"encoding/json";"os")
type (
student struct{

	Rating []int
}
Group struct{
Students []student
}
Average struct{
 Average float32
}
)

func main(){
var data Group
var avg Average
jinf,_:=os.ReadFile("json.txt")
//fmt.Println(string(jinf))
json.Unmarshal(jinf,&data)
//fmt.Printf()
var mark float32 = 0
var num float32 = 0
for _,val:=range data.Students{
  num++
	for range val.Rating{
       mark++
	}
	
}
avg.Average = mark/num

//fmt.Print(avr.average)
str,_ := json.MarshalIndent(avg,"","    ")
fmt.Printf("%s\n",str)
fmt.Println(avg)
}