package questions

import (
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

const catImageAPI = "https://cataas.com/cat"

func main2() {
	http.HandleFunc("/greeting", greeting)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("ListenAndServe Error: %v", err)
	}
}

func greeting(w http.ResponseWriter, r *http.Request) {
	if currentUTCHour := time.Now().UTC().Hour(); isCatImageTime(currentUTCHour) {
		responseWithCatImage(w)
		return
	}
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("failed to parse request: %v", err.Error())
		return
	}
	responseWithGreetingText(w, r.Form.Get("name"))
}

func isCatImageTime(currentHour int) bool {
	if currentHour >= 14 && currentHour < 16 {
		return true
	}
	return false
}

func responseWithGreetingText(w http.ResponseWriter, name string) {
	greetingText := getGreetingText(name)
	if _, err := w.Write(greetingText); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("greeting text err", err.Error())
		return
	}
	return
}

func getGreetingText(name string) []byte {
	noSpacedName := strings.TrimSpace(name)
	if noSpacedName != "" {
		return []byte("hello " + noSpacedName)
	}
	return []byte("hello world")
}

func responseWithCatImage(w http.ResponseWriter) {
	catResponse, err := http.Get(catImageAPI)
	if err != nil || catResponse.StatusCode != http.StatusOK {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("get random cat err", err.Error())
		return
	}
	if catResponse == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	body := catResponse.Body
	defer body.Close()
	bytesArr, err := io.ReadAll(body)
	if _, err = w.Write(bytesArr); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("random cat err", err.Error())
	}
}
