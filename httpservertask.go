package main
import (
	"fmt"
	"log"
	"net/http"
)
func handler(w http.ResponseWriter,r *http.Request){
	fmt.Println("hello sample http")
	fmt.Fprintf(w, "hello\n")
}
func main(){

	http.HandleFunc("/",handler)

	log.Fatal(http.ListenAndServe(":8080",nil))
}


//json write cheyyuka--variabline json akkunnathu nokkuuka goyil