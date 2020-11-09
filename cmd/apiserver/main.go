package main

import (
	"flag"
	"fmt"
	"log"
	"io/ioutil"
	"net/http"
	"github.com/BurntSushi/toml"
	"github.com/pererva-ev/http-rest-api/internal/app/apiserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}




func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	s := apiserver.New(config)
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Start(); err != nil {
		log.Fatal(err)


		var URL = "localhost:8080/hello"

		tr := &http.Transport{DisableKeepAlives: false}
        req, _ := http.NewRequest("GET", URL, nil)
        req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", "Token"))
        req.Close = false
		req.Header.Add("User-Agent", "myClient")
		

		if err != nil {
			fmt.Println(err)
		}


        res, err := tr.RoundTrip(req)
        if err != nil {
            fmt.Println(err)
        }
        body, _ := ioutil.ReadAll(res.Body)
        fmt.Println(string(body))
	}
}