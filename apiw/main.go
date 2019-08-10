package main

import (
	"fmt"
	mw "github.com/wesmota/goworkshop2.0/apiw/middleware"
	"io/ioutil"
	"net/http"
	"strings"
)
var sb strings.Builder

func SendJson(w http.ResponseWriter, r *http.Request){
	//fmt.Println("Recebendo Json")

	b,err := ioutil.ReadAll(r.Body)
	//fmt.Println(b)
	w.Write([]byte(b))
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", string(len(b)))



	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(400)
		sb.Write([]byte(`{"status":"error", "msg:"`))
		sb.Write([]byte(err.Error()))
		sb.Write([]byte(`"}"`))
		fmt.Println(sb.String())
		w.Write([]byte(sb.String()))

		return
	}

}


func Ping(w http.ResponseWriter, r *http.Request){
	fmt.Println("method: ", r.Method)

	if strings.ToUpper(r.Method) == "POST" {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))

		return
	}
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("Error permitido somente POST ...\n"))
}

func main() {

	Ping2 := func(w http.ResponseWriter, r *http.Request){
		w.WriteHeader(http.StatusOK)
		json := `{"status":"OK", "msg":"sucesso"}`
		json += "\n"
		w.Write([]byte(json))
	}

	handlerApiJson := http.HandlerFunc(SendJson)
	http.Handle("/api/v1/sendjson", mw.Use(handlerApiJson,mw.Logger("SendJson")))


	http.HandleFunc("/api/v1/ping1", Ping)
	http.Handle("/api/v1/ping2", http.HandlerFunc((Ping2)))

	http.HandleFunc("/api/v1/ping4", Ping)
	fmt.Println("Run Server: 8085")
	http.ListenAndServe(":8085", nil)



}
