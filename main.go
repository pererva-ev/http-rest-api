package main

import (
	"log"

	"github.com/ev-pererva/http-rest-api/internal/app/apiserver"
)

func main() {
	s := apiserver.New()
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}

}
