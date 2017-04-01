package main

import (
	"net/http"
	"fmt"
	"time"
	"github.com/hughbrien/godemo/stringutils"
	"github.com/instana/golang-sensor"
	ot "github.com/opentracing/opentracing-go"
	"golang.org/x/net/context"


)

const (
	Service = "goes-eleven-demo"
	Entry   = "http://localhost:9060/golang/entry"
	Exit    = "http://localhost:9060/golang/exit"
)

var demoOptions instana.Options

var FirstName string
var LastName string
var MiddleName string
var DateOfBirth time.Time
var DemoContext context.Context




func main() {

	FirstName = "Hugh"
	MiddleName = "Plunkett"
	LastName = "Brien"
	DateOfBirth = time.Now()


	thisBannerFunc := func (stringValue string) string {
	return " # # # # # # # # # " +
		stringValue +
		"# # # # # # # # # # "
	}




	fmt.Println(thisBannerFunc("Starting COR"))


	opts := instana.Options {
		Service: Service,
		LogLevel: instana.Info,
	}

	ot.InitGlobalTracer(instana.NewTracerWithOptions(&opts))


	DemoContext = context.WithValue(context.Background(), "Foo", "Bar")

	fmt.Println(DemoContext)

	//Span span  = ot.StartSpanFromContext(c)
	fmt.Println(stringutils.Banner("Starting on port 8080"))


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
	fmt.Println(stringutils.Banner(numstring))
	fmt.Println(w)
	fmt.Println("\n")
	fmt.Println(stringutils.Banner("HTTP Request"))
	fmt.Print(req)
	fmt.Println("\n")
	printPage(w)

}


func printPage( w http.ResponseWriter){

	DateOfBirth = time.Now()
	extraResult := FirstName + " " + LastName
	w.Write([]byte ("<h1>"))
	w.Write([]byte (extraResult))
	w.Write([]byte ("This is the first Text from the the Application Server</h1>") )


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


