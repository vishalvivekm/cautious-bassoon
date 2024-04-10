package main

//
//import (
//	"encoding/json"
//	log "github.com/sirupsen/logrus"
//	"net/http"
//	"os"
//)
//
//func main() {
//	log.Info("starting server..")
//	http.HandleFunc("/api", Handler)
//
//	log.Fatalln(http.ListenAndServe(":8080", nil))
//}
//
//func Handler(writer http.ResponseWriter, request *http.Request) {
//	response, err := http.Get("http://hn.algolia.com/api/v1/search?query=foo&tags=story")
//	if err != nil {
//		log.Fatalf("err geting response form api, %v", err)
//	}
//	var rd map[string]interface{}
//	err = json.NewDecoder(response.Body).Decode(&rd)
//	if err != nil {
//		log.Fatalf("error decoding response, %v", err)
//	}
//	bf, err := json.MarshalIndent(rd, "", " ")
//	if err != nil {
//		log.Fatalf("error marshaling, %v", err)
//	}
//
//	file, err := os.OpenFile("./tempt.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer file.Close()
//
//	_, err = file.WriteString(string(bf))
//	if err != nil {
//		log.Fatalln(err)
//	}
//	_, err = writer.Write(bf)
//	if err != nil {
//		log.Fatalf("error writing, %v", err)
//		return
//	}
//}
