package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	dir, _ := os.Getwd()
	// fmt.Println("current path :" + dir)
	public_dir := fmt.Sprintf("%s/public", dir)

	fs := http.FileServer(http.Dir(public_dir))
	log.Print(fs)
	http.Handle("/", fs)

	log.Print("Listening on :3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
