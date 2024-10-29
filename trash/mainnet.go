package main
// данные пакеты нужны для системы проверки
import (
    "fmt"
	"io"
	"net/http"
	"net/url"
	//"time"
)

func main() {
    
    var  name, age string
    fmt.Scan(&name,&age)
    URL,_:= url.Parse("http://127.0.0.1:8080/hello")
    params :=url.Values{}
    params.Add("name", name)
    params.Add("age", age)
    URL.RawQuery = params.Encode()

    resp,_:=http.Get(URL.String())
    defer resp.Body.Close()
    data,_:=io.ReadAll(resp.Body)
    fmt.Printf("%s",data)
}