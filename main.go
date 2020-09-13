package rapidcaptcha_server

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Captcha struct {
	GroupID   string `json:"groupId"`
	Image     string `json:"image"`
	IsRequest bool   `json:"isRequest"`
}

var addr = flag.String("addr", ":80", "http service address")
var hub = newHub()

func serveCaptchaByID(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Println(r.Body)
	}
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	log.Println("Path:", r.URL.Path)

	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "home.html")
}

func serveCaptcha(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/captcha" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	groupIDs, ok := r.URL.Query()["groupID"]

	if !ok || len(groupIDs[0]) < 1 {
		log.Println("Url Param 'groupID' is missing")
		http.ServeFile(w, r, "rapidcaptcha.html")
	}

	groupID := groupIDs[0]

	if _, ok := captchas[string(groupID)]; ok {
		b, err := json.Marshal(captchas[string(groupID)])

		cm := map[string]interface{}{"CaptchaJson": string(b)}
		t, err := template.ParseFiles("rapidcaptcha.html")
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		if err := t.Execute(w, cm); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	} else {
		http.ServeFile(w, r, "rapidcaptcha.html")
	}
}

func main() {
	flag.Parse()
	go run()

	http.Handle("/css/", http.FileServer(http.Dir(".")))
	http.Handle("/js/", http.FileServer(http.Dir(".")))
	http.Handle("/img/", http.FileServer(http.Dir(".")))

	http.HandleFunc("/", serveHome)
	http.HandleFunc("/captcha", serveCaptcha)

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
