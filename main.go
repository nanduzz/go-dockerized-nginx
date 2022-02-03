package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func readFile() (string, error) {
	data, err := ioutil.ReadFile("static/number.txt")

	if err != nil {
		return "", err
	}

	return string(data), nil

}
func main() {

	http.HandleFunc("/", handleCall)

	http.ListenAndServe(":80", nil)
	fmt.Println(readFile())
}

func handleCall(w http.ResponseWriter, r *http.Request) {
	number, err := readFile()
	if err != nil {
		log.Fatal(err)
	}
	name := os.Getenv("NAME")

	w.Write([]byte(fmt.Sprintf("Number: %s, name: %s", number, name)))
}
