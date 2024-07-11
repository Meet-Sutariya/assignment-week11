package main

<<<<<<< HEAD
import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
=======
func main()
>>>>>>> f0f980aadc13a68e9ec6db5ee95e79e052747f76
