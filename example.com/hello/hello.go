package main

import (
	"example.com/greetings"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

import "rsc.io/quote"


func main() {
    fmt.Println(quote.Go())
    message := greetings.Hello("Gladys")
    fmt.Println(message)

	server := &http.Server{
		Addr:         ":4000",
		WriteTimeout: time.Second * 2, // 4
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	mux := http.NewServeMux()
	mux.Handle("/", new(myHandler))
	mux.HandleFunc("/bye", sayBye)

	go func() {
		<-quit

		if err := server.Close(); err != nil {
			log.Fatal("Close server:", err)
		}
	}()

	server.Handler = mux
	log.Print("Starting server... v3")
	err := server.ListenAndServe()
	if err != nil {
		if err == http.ErrServerClosed {
			log.Print("Server closed under request")
		} else {
			log.Fatal("Server closed unexpected")
		}
	}
}

func sayBye(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Bye bye this is version 1!"))
}

type myHandler struct {}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("Hello v2, the request URL is: " + r.URL.String()))
}

var nC int
var groupC []int
var profitC []int
var minProfitC int
var length int
var res int = 0

func profitableSchemes(n int, minProfit int, group []int, profit []int) int {
	nC = n
	groupC = group
	profitC = profit
	minProfitC = minProfit
	length = len(group)
	helper(0, 0, 0)
	return res
}

func helper(peopleCount int, curPro int, index int) {
	if curPro >= minProfitC {
		res += 1
	}
	for i := index; i < length; i++ {
		if groupC[i] + peopleCount > nC {
			continue
		} else {
			helper(groupC[i]+peopleCount, curPro+profitC[i], i)
		}
	}
}