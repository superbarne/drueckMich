package main

import (
	"fmt"
	"net/http"
	"github.com/superbarne/drueckMich/webui"
	"github.com/superbarne/drueckMich/api"
)

func main() {
	mux := http.NewServeMux()

	webui.Configure(mux)
	api.Configure(mux)

	fmt.Println("Listening...")
	http.ListenAndServe(":3000", mux)
}
