package server

import (
	"encoding/xml"
	"fmt"
	"github.com/gombio/batory/operators"
	"github.com/sirupsen/logrus"
	"net/http"
)

type SayResponse struct {
	XMLName xml.Name `xml:"Response"`
	Say     string   `xml:",omitempty"`
}

func callHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		twilio := operators.NewTwilio()
		info, error := twilio.ParseIncomingCall(r.PostForm)

		if error != nil {
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
		}
		logrus.Println(fmt.Sprintf("Call from %s", info.CallerNumber))
		res := SayResponse{Say: "Hello Jery what are you looking for"}

		_,e := xml.Marshal(res)
		if e != nil {
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/xml")
		fmt.Fprintf(w, "s")
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
