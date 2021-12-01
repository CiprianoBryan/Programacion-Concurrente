package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// estructura
type Alumno struct {
	Codigo string `json:"code"`
	Nombre string `json:"name"`
	Dni    int    `json:"dni"`
}

var listaAlumnos []Alumno

func cargarDatosAlumnos() {
	listaAlumnos = []Alumno{
		{"2021486", "Juan Garcia", 81234567},
		{"2021485", "Miguel Campos", 41234565},
		{"2021484", "Maria Perez", 31234562}}
}

func mostrarHome(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "text/html")
	io.WriteString(response, `
		<html>
			<head></head>
			<body><h2>API de Alumnos</h2</body>
		</html>
	`)
}

func listarAlumnos(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	jsonBytes, _ := json.MarshalIndent(listaAlumnos, "", " ")
	io.WriteString(response, string(jsonBytes))
}

func buscarAlumno(response http.ResponseWriter, request *http.Request) {
	log.Println("Se ingresó a buscar Alumno")
	cod := request.FormValue("codigo")
	response.Header().Set("Content-Type", "application/json")

	for _, alumno := range listaAlumnos {
		if alumno.Codigo == cod {
			jsonBytes, _ := json.MarshalIndent(alumno, "", " ")
			io.WriteString(response, string(jsonBytes))
		}
	}
}

func agregarAlumno(response http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		if request.Header.Get("Content-type") == "application/json" {
			log.Println("Accede al método agregar alumno")
			jsonBytes, err := ioutil.ReadAll(request.Body)
			if err != nil {
				log.Fatal(err)
				http.Error(response, "Error al leer el cuerpo del request", http.StatusInternalServerError)
			}
			var alumno Alumno
			json.Unmarshal(jsonBytes, &alumno)
			listaAlumnos = append(listaAlumnos, alumno)
			// MENSAJE DE RESPUESTA
			response.Header().Set("Content-Type", "application/json")
			io.WriteString(response, `
			{
				"respuesta": "Alumno creado satisfactoriamente!!"
			}`)
		} else {
			http.Error(response, "Contenido no válido", http.StatusBadRequest)
		}
	} else {
		http.Error(response, "Método no válido", http.StatusMethodNotAllowed)
	}
}

func manejadorSolicitudes() {
	// enrutador
	mux := http.NewServeMux()
	// endpoints
	mux.HandleFunc("/home", mostrarHome)
	mux.HandleFunc("/listar", listarAlumnos)
	mux.HandleFunc("/buscar", buscarAlumno)
	mux.HandleFunc("/agregar", agregarAlumno)
	http.ListenAndServe(":9000", mux)
}

func main() {
	cargarDatosAlumnos()
	manejadorSolicitudes()
}
