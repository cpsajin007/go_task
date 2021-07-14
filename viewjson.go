package main
import (
	"fmt"
	"log"
	"net/http"
)
type Book struct{
	Title string
	Author string
}
func handler(w http.ResponseWriter,r *http.Request){
	fmt.Println("hello sample http")
	fmt.Fprintf(w, "hello\n")
	book := Book{Title: "Learning Concurreny in Python", Author: "Elliot Forbes",}
	fmt.Printf("%+v\n",book)
	fmt.Fprintf(w,"%+v\n", book)
}
func main(){
	// book := Book{Title: "Learning Concurreny in Python", Author: "Elliot Forbes",}
	// fmt.Printf("%+v\n",book)
	http.HandleFunc("/sajin",handler)

	log.Fatal(http.ListenAndServe(":8080",nil))
}
