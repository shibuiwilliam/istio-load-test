package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

var (
	num int
)

func x2y2(x, y float64) float64 {
	return (x * x) + (y * y)
}

func random() float64 {
	x := rand.Float64()
	xi := rand.Intn(2)
	if xi == 0 {
		x = -x
	}
	return x
}

func pi(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(num)
	o := 0
	for i := 0; i < n; i++ {
		x := random()
		y := random()
		xy := x2y2(x, y)
		if xy <= 1 {
			o++
		}
	}
	p := 4 * float64(o) / float64(n)
	s := strconv.FormatFloat(p, 'f', 4, 64)
	fmt.Fprintf(w, s)
}

func startWebServer() error {
	fmt.Println("Rest API with Mux Routers")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/pi", pi).Methods("GET")

	return http.ListenAndServe(fmt.Sprintf(":%d", 8080), router)
}

func main() {
	flag.IntVar(&num, "num", 1000000, "num")
	flag.Parse()
	startWebServer()
}
