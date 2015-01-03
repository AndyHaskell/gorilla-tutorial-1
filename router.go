package main

import(
    "net/http"
    "fmt"
	"log"
	"html"

    "github.com/gorilla/mux"
)

func initRouter() *mux.Router{
    r := mux.NewRouter()

    r.HandleFunc("/sloths", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Sloths rule!")
    })

    r.PathPrefix("/images/").Handler(http.StripPrefix("/images/",
                                       http.FileServer(
                                         http.Dir("public/images"))))

	r.HandleFunc("/tea/{flavor}", func(w http.ResponseWriter, r *http.Request){
		params := mux.Vars(r)
		tea    := html.EscapeString(params["flavor"]) + " tea"
		html   := `<body><img src="/images/sloth.jpg" /><br />`+
	               `<h2>I could use some ` + tea + `!</h2></body>`
		fmt.Fprintf(w, html)
	})

	r.HandleFunc(`/{drink:(coffee)+}`, func(w http.ResponseWriter, r *http.Request){
		html := `<body><img src="/images/lemur.jpg" /><br />`+
        	     `<h2>Lemurs = sloths that had too much coffee!</h2></body>`
		fmt.Fprintf(w, html)
	})

	r.HandleFunc("/coffee-shop", func(w http.ResponseWriter, r *http.Request){
		html := `<body><form action="order" method="POST">`+
			`Your name <input type="text" name="name"><br />`+
			`Your beverage order <input type="text" name="beverage"><br />`+
			`<input type="submit" value="Submit">`+
			`</form></body>`
		fmt.Fprintf(w, html)
	}).Methods("GET")

	r.HandleFunc("/order", func(w http.ResponseWriter, r *http.Request){
		err := r.ParseForm()
		if err != nil {
			log.Fatal(err.Error())
		}

		name     := html.EscapeString(r.Form.Get("name"))
		beverage := html.EscapeString(r.Form.Get("beverage"))
		html := `<body><h1>One ` + beverage + ` coming right up, ` + name + `!</h1></body>`
		fmt.Fprintf(w, html)
	}).Methods("POST")

    r.PathPrefix("/").HandlerFunc(serveHelloWorld)

    return r
}
