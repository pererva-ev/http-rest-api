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


func Cors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=ascii")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers","Content-Type,access-control-allow-origin, access-control-allow-headers")
	w.Write([]byte("Hello, World!"))
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
		
		mux := http.NewServeMux()
		mux.HandleFunc("/hello",Cors)
		http.ListenAndServe(":8080", mux)
		
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
// package main

// import (
//     "net/http"
//     "sync/atomic"
// )

// var requests int64 = 0

// // increments the number of requests and returns the new value
// func incRequests() int64 {
//     return atomic.AddInt64(&requests, 1)
// }

// // returns the current value
// func getRequests() int64 {
//     return atomic.LoadInt64(&requests)
// }

// func handler(w http.ResponseWriter, r *http.Request) {

//     incRequests()

//     // handle the request here ...
// }

// func main() {
//     http.HandleFunc("/", handler)
//     log.Fatal(http.ListenAndServe(":8080", nil))
// }

// <script src="https://unpkg.com/axios/dist/axios.min.js"></script>
// <script>
// 	(async () => {
// 		const response = await axios({
// 		url: 'http://localhost:8080/hello',
// 		method: 'get'
// 	})
// 	alert(response)
// 	})()
// </script>