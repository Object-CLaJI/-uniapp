package main

import (
"fmt"
"log"
"net/http"
_"net/http/pprof"
)

func main() {
fmt.Println("123")
go func() {
 log.Println(http.ListenAndServe(":9999",nil))
}()

for {
Add("pprof")

}
}

func Add(str string) string {
data := []byte(str)
sData := string(data)
var sum = 0
for i := 0; i < 1000; i++ {

sum += i
}
return sData
}


