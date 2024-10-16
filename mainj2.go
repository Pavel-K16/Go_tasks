package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type (
	ID struct {
		Id uint64 `json:"global_id"`
	}
	All struct {
		IDs []ID
	}
)

func main() {
	jinf, _ := os.ReadFile("data-20190514T0100.json")
	var (
		ALL []ID
		sum uint64
	)
	json.Unmarshal(jinf, &ALL)
	for _, val := range ALL {
		sum += val.Id
	}

	//fmt.Println(sum)

	jinf, _ = json.MarshalIndent(ALL, "", "    ")
	//fmt.Println(string(jinf))


	s_time:="1986-04-16T05:20:00+06:00"
    resTime,_:=time.Parse(time.RFC3339,string(s_time))
    fmt.Println(string(s_time))
    fmt.Println(resTime.Format(time.UnixDate))


}
