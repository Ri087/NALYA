package nalya

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func DataHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(body))
	w.Write([]byte("online"))
}
