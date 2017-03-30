package main

import (
	"net/http"
	"fmt"
	"time"
	"github.com/instana/golang-sensor"
	ot "github.com/opentracing/opentracing-go"
)

const (
	Service = "goes-eleven-demo"
	Entry   = "http://localhost:9060/golang/entry"
	Exit    = "http://localhost:9060/golang/exit"
)

var demoOptions instana.Options


func main() {

	fmt.Print ( " # # # # # # Starting CORE # # # # # # " )

	opts := instana.Options {
		Service: Service,
		LogLevel: instana.Info,
	}

	ot.InitGlobalTracer(instana.NewTracerWithOptions(&opts))

	//Span span  = ot.StartSpanFromContext(c)
	fmt.Println ("Starting on port 8080")
	myMux := http.NewServeMux()
	myMux.HandleFunc("/", catchall)
	myMux.HandleFunc("/service", demoservice)
	myMux.HandleFunc("/products", products)
	myMux.HandleFunc("/demos", demos)

	http.ListenAndServe(":8080", myMux)
}



func catchall (w http.ResponseWriter, req *http.Request) {

	demoOptions = instana.Options{"hostname","",0,0}

	second := time.Second
	miliseconds := int64(second/time.Millisecond)

	fmt.Println ("\n")
	var num int64 = int64(second/time.Millisecond)

	fmt.Print(miliseconds)
	numstring:=fmt.Sprintf("%d",num)

	fmt.Print ( " # # # # # # # # # # # # " )
	fmt.Println(numstring)
	fmt.Print ( "Response Writer # # # # # # # # # # ")
	fmt.Println(w)
	fmt.Println("\n")
	fmt.Println(" # # # # # # # # #  GoLang HTTP Request # # # # # # # # # # ")
	fmt.Print(req)
	fmt.Println("\n")
	printPage(w)

}




func printPage( w http.ResponseWriter){

	w.Write([]byte ("<h1>This is the first Text from the the Application Server</h1>"))

}

func demoservice (w http.ResponseWriter, req *http.Request) {

	w.Write([]byte("Demo services takes 4 seconds "))
	time.Sleep(time.Second * 4)
	w.Write([]byte("\n"))
}

func products (w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("products"))
}

func demos (w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("demos"))
}


