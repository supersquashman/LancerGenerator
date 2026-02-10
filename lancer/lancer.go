package main

import (
	"mechs"
	"dataload"
	"encoding/json"
	"fmt"
	"net/http"
	// "github.com/gin-gonic/gin"
	"log"
	"github.com/rs/cors"
)

func newMech() string{
	mech := mechs.GenerateMech()
	jsonMech, _ := json.Marshal(mech)
	// return mechs.ToString(mech)
	// return mech
	return string(jsonMech)
}

type testMech struct{}
func (t *testMech) ServeHTTP(w http.ResponseWriter, r *http.Request){
	w.Write([]byte(newMech()))
}

func main(){
	dataload.Load_All()
	
	mech := mechs.GenerateMech()
	fmt.Println(mechs.ToString(mech))

	router := http.NewServeMux()
	router.Handle("/mech/new", &testMech{})

	// config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"*"}
	// config.AllowedMethods = []string{"GET","POST","PUT","DELETE","OPTIONS"}
	// config.AllowedHeaders = []string{"*"}
	// config.AllowedCredentials = true

	config := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET","POST","PUT","DELETE","OPTIONS","OPTION"},
		AllowedHeaders: []string{"*"},
		AllowCredentials: true})

	handler := config.Handler(router)
	// router := httprouter.New()
	// router.GET("/mech/new",newMech)

	log.Fatal(http.ListenAndServe(":9090",handler))

}