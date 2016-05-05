package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"

	"golang.org/x/net/websocket"
)

var eightBallResponse = []string{
	"It is certain",
	"It is decidedly so",
	"Without a doubt",
	"Yes definitely",
	"You may rely on it",
	"As I see it yes",
	"Most likely",
	"Outlook good",
	"Yes",
	"Signs point to yes",
	"Reply hazy try again",
	"Ask again later",
	"Better not tell you now",
	"Cannot predict now",
	"Concentrate and ask again",
	"Don't count on it",
	"My reply is no",
	"My sources say no",
	"Outlook not so good",
	"Very doubtful",
}

func echoHandler(ws *websocket.Conn) {
	io.Copy(ws, ws)
}

func homeHandler(w http.ResponseWriter, req *http.Request) {
	//fmt.Printf("%+v\n", req)
	fmt.Printf("%v %v %v\nRemoteAddr:%v\nUser-Agent:%v\nAccept:%v\n", req.Method, req.URL, req.Proto, req.RemoteAddr, req.Header["User-Agent"], req.Header["Accept"])
	fmt.Fprint(w, "Welcome to my Go API.")
}

func md5Handler(w http.ResponseWriter, req *http.Request) {
	parms := req.URL.Query()
	_, ok := parms["v"]

	if ok {
		data := []byte(req.URL.Query()["v"][0])
		fmt.Fprintf(w, "%x", md5.Sum(data))
	} else {
		fmt.Fprintln(w, "Add query string '?v=hashThis' to url.")
	}
}

func sha1Handler(w http.ResponseWriter, req *http.Request) {
	//data := []byte(req.URL.Query()["v"][0])
	parms := req.URL.Query()
	_, ok := parms["v"]

	if ok {
		h := sha1.New()
		io.WriteString(h, req.URL.Query()["v"][0])
		fmt.Fprintf(w, "%x", h.Sum(nil))
	} else {
		fmt.Fprintln(w, "Add query string '?v=hashThis' to url.")
	}
}

func eightBallHandler(w http.ResponseWriter, req *http.Request) {
	l := len(eightBallResponse)
	rand.Seed(time.Now().UnixNano())
	fmt.Fprintln(w, eightBallResponse[rand.Intn(l)])
}

func testerHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "You've hit the test endpoint.")
}

func fakeSoap(w http.ResponseWriter, req *http.Request) {
	parms := req.URL.Query()
	_, ok := parms["wsdl"]
	if ok {
		fmt.Fprintln(w, "This would be the actual WSDL XML")
	} else {
		fmt.Fprintln(w, "This would be that SOAP usage page.")
	}
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.Handle("/echo", websocket.Handler(echoHandler))
	http.HandleFunc("/md5", md5Handler)
	http.HandleFunc("/sha1", sha1Handler)
	http.HandleFunc("/eightball", eightBallHandler)
	http.HandleFunc("/tester", testerHandler)
	http.HandleFunc("/SomeWare/SomeService.svc", fakeSoap)

	log.Println("Attempting to listen on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
