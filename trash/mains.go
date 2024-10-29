package main

import ("fmt"; "encoding/csv"; "os")
func main(){
//	k:=true
File, _:=os.Open("task.data")
defer File.Close()
rd:=csv.NewReader(File)
rd.Comma = ';'
for records, err:=rd.Read(); err == nil; records, err = rd.Read() {
for i,r:=range records{ 
	n_:=len(r)
	if (n_ == 1 &&  r[0] == '0'){
      fmt.Print(i+1)
	  break
	}
}
}
}