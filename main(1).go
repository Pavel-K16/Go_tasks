package main

import (
	//"bufio"
	"fmt"
	"strconv"
	//"os"
//	"unicode"
)

type Smth struct {
	On          bool
	Ammo, Power int
}

func (s *Smth) Shoot() bool {
	if s.On == false {
		return false
	} else {
		s.Ammo--
		return true
	}
}
func (s *Smth) Ridebike() bool {
	if s.On == false {
		return false
	} else {
		s.Power--
		return true
	}
}

type SREZ struct{
a int
}
//func ()
func (s SREZ) CHANGE() {

	s.a = 0
}
func main() {
var sr SREZ
var b = 1
sr.a = b
fmt.Println(sr.a)
sr.CHANGE()
fmt.Println(sr.a)
	//var  a = []int{1,2,3}
//	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	//s := []rune(text)
	//if s[len(s)-1-2] == '.' && unicode.IsUpper(s[0]) {
	//	fmt.Print("Right")
	//} else {
//		fmt.Println("Wrong")
	//}
//	fmt.Print(4 % 2)

	//s__:= "123456"
	//RUNA:=[]rune(s__)
	//fmt.Println(s__)
	//fmt.Print(s__[0]," ",s__[1]," ",s__[2]," ",s__[3]," ",s__[4]," ",s__[5])
	//fmt.Println(len(s__))
	
	//fmt.Println((RUNA))
	//fmt.Println(RUNA[0])

    fn:= func(i uint64) uint {
		s:=strconv.Itoa(int(i))
		r:=[]rune(s)
		n:= len(r)
		k:=0
		for j:=0; j < n; j++{
			val,_ := strconv.Atoi(string(r[j]))
			if (val % 2 == 0){
				r[k] = r[j]
				k++
			}
		} 
		r = r[0:k]
	     val,_ := strconv.Atoi(string(r))
		 if k == 0{
			val = 100
			}
			if (k == 1 && val == 0){
				val = 100
			}
		return uint(val)
	}
     fmt.Println(fn(0))

	//fmt.Println(len(RUNA))
}
