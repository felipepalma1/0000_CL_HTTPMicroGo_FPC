package main

import ( 
	"net/http"
	"fmt"
	"log"
	"io/ioutil"
)

var mensajeUno string = `Este mensaje, que es una cadena de caracteres, viaja por los interiores de la máquina y regresa para imprimirse por la pantalla.`
var mensajeDos string = `Adios!`

func main(){
	http.HandleFunc("/", func(rw http.ResponseWriter, r*http.Request){
		log.Println(mensajeUno)
		d, err := ioutil.ReadAll(r.Body)

		if err != nil {

			http.Error(rw, "Oh No", http.StatusBadRequest)
			return
			
		}

		log.Printf("Datos %s \n", d)

		fmt.Fprintf(rw, "Hola %s", d)


	})

	http.HandleFunc("/exit", func(http.ResponseWriter, *http.Request){
		log.Println(mensajeDos)
	})

	http.ListenAndServe(":9090", nil)
}