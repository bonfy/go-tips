package main

import (
	"net/http"

	"github.com/unrolled/render"
)

var rd = render.New()

// Person Struct
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	rd.JSON(w, http.StatusOK, map[string]interface{}{"status": "ok", "result": "hello world"})
}

func handleObject(w http.ResponseWriter, r *http.Request) {
	p := Person{Name: "ben", Age: 16}
	rd.JSON(w, http.StatusOK, map[string]interface{}{"status": "ok", "result": p})
}

func handleList(w http.ResponseWriter, r *http.Request) {
	people := []Person{
		Person{Name: "ben", Age: 16},
		Person{Name: "ken", Age: 18},
	}
	rd.JSON(w, http.StatusOK, map[string]interface{}{"status": "ok", "result": people})
}

func main() {
	http.HandleFunc("/", handleHello)
	http.HandleFunc("/obj", handleObject)
	http.HandleFunc("/list", handleList)

	http.ListenAndServe(":8888", nil)
}

/*

r := render.New(render.Options{
  DisableHTTPErrorRendering: true,
})

//...

err := r.HTML(w, http.StatusOK, "example", "World")
if err != nil{
  http.Redirect(w, r, "/my-custom-500", http.StatusFound)
}

*/
