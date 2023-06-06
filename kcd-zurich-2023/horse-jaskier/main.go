package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var doorState = ""
var brokerURL = ""

func main() {
	log.SetOutput(os.Stdout)
	brokerURL = os.Getenv("BROKER_URL")

	http.Handle("/", http.FileServer(http.Dir(os.Getenv("KO_DATA_PATH"))))
	http.HandleFunc("/knock-on-door", knockOnDoor)
	http.HandleFunc("/cloudevents", receiveCloudEvent)
	http.HandleFunc("/door-state", getDoorState)

	log.Println("Running web server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func knockOnDoor(w http.ResponseWriter, r *http.Request) {
	if brokerURL == "" {
		log.Fatal("Please specify env `BROKER_URL`")
	}

	log.Printf("server: sending `Knocking on Door` to %s", brokerURL)

	jsonBody := []byte(`Knock on the Door`)
	bodyReader := bytes.NewReader(jsonBody)

	req, err := http.NewRequest(http.MethodPost, brokerURL, bodyReader)
	if err != nil {
		log.Printf("server: failed to create http request: %s\n", err.Error())
	}

	req.Header.Add("ce-specversion", "1.0")
	req.Header.Add("ce-id", "demo")
	req.Header.Add("ce-type", "knock-on-door")
	req.Header.Add("ce-source", "horse-jaskier")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Printf("server: failed to send http request: %s\n", err.Error())
	}

	log.Printf("server: received statusCode: %v from broker", res.StatusCode)
}

func receiveCloudEvent(w http.ResponseWriter, r *http.Request) {
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("server: could not read request body: %s\n", err)
	}
	log.Printf("server: request body: %s\n", reqBody)

	doorState = string(reqBody)
}

func getDoorState(w http.ResponseWriter, r *http.Request) {
	log.Println("Sending back door state: ", doorState)
	fmt.Fprintf(w, doorState)
}
