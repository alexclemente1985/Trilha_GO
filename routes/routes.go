package routes

import (
	c "loja/controllers"
	"net/http"
)



func CarregaRotas() {
	http.HandleFunc("/", c.Index)
	http.HandleFunc("/new", c.New)
	http.HandleFunc("/insert", c.Insert)
	http.HandleFunc("/edit", c.Edit)
	http.HandleFunc("/update", c.Update)

}