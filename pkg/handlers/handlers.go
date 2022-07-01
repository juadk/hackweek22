package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/juadk/hackweek22/pkg/api"
	"github.com/juadk/hackweek22/pkg/render"
)

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		render.RenderTemplate(w, "index.page.tmpl")
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() failed. err=%v", err)
			return
		}

		name := r.FormValue("name")
		currency := r.FormValue("currency")

		cg := api.NewClient(nil)

		price, err := cg.SimpleSinglePrice(name, currency)
		if err != nil {
			log.Fatal(err)
		}
		render.RenderTemplate(w, "index.page.tmpl")
		fmt.Fprintf(w, "One %v is worth %v %v", price.ID, price.MarketPrice, price.Currency)
	}
}
