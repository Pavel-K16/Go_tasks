package main 

import "fmt"
type Human struct{
   height,weight int
   health float64  
}
type Num struct{
   num int 
Person Human
}

func (s *Human) Set_zero() (int){
	
  s.health = 0
  s.weight = 0
  s.height = 0
  return 0
}


func (s *Num) Set_zero() (int, float64){
	s.num = 0
  return 0,0.3
}


func main(){
	var mas = [3]int{2,2,2}  
fmt.Println(mas)
s:=make([]int,10,12)
fmt.Println(s[9])
s = append(s, 10,11)
fmt.Println(s[11])
fmt.Println(s)
fmt.Println(cap(s))
s = append(s, 10,11,12)
fmt.Println(cap(s))
//pointer:= fmt.Sprintf("%p", s)
fmt.Println(s)
s = append(s[0:2],s[3:]...)
fmt.Println(s)
s_:=make([]int,22,22)
n:=copy(s,s_)
fmt.Println(s,"  ", n, "  ", len(s))

//var man Human = Human{health: 1,weight: 2,height: 3}
man := Human{1,2,3}
fmt.Println(man)
men := Num {16, man}
fmt.Println(men)
men = Num{
	num: 16,
	Person: Human{
		weight: 12,
		height: 12,
		health: 27.1,
	},
}
fmt.Println(men)
_,v := men.Set_zero()
men.Person.Set_zero()
fmt.Println(men, "  ",v)


var m  = map[string]int{
	"1": 12,
	"2":13,
}
fmt.Println(m, "  ")
m_:= make(map[int]int)
m_[1] = 1
m_[2] = 20

fmt.Println(m_, "  ")
delete(m_,1)
fmt.Println(m_, "  ")
//var val int 
val, key := m_[2]
fmt.Println(val, "  ", key)
slice:=mas[:]
slice = append(slice, 5)
fmt.Println(slice," ", cap(slice)," ", len(slice)," ", cap(mas), len(mas))
}