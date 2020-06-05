package mocksentra

/*
 * BFF Mi Pago
 *
 * BFF de Mi Pago
 *
 * API version: 1.0-beta
 * Contact: evillagran@bancochile.cl
 * SEnTRA 2020
 */

import (
	"io"
	"log"
	"net/http"
	"os"
)

type ActualizaPerfilUsingResponse struct {
	Alias string `json:"alias"`
	Banca struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"banca"`
	EmailParticularCliente string `json:"emailParticularCliente"`
	Marca                  struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"marca"`
	Nombre   string `json:"nombre"`
	Rut      string `json:"rut"`
	Segmento struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"segmento"`
	Sexo string `json:"sexo"`
}

//ActualizaPerfilUsingPOST , Creación o actualización del perfil de usuario
//curl -X GET "http://localhost:8080/v1/clientes/persona/perfil/" -H "accept: application/json" -H "Authorization: asddasd" -H "idApp: aasdasd"
func ActualizaPerfilUsingPOST(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	response := ActualizaPerfilUsingResponse{}
	response.Alias = "cdiaz"
	response.EmailParticularCliente = "cdiaz@gmail.com"
	response.Nombre = "Carlos Diaz"
	response.Rut = "15125990-1"
	response.Sexo = "M"
	response.Segmento.Key = "PER"
	response.Segmento.Value = "Persona Natural"
	response.Banca.Key = "CTA"
	response.Banca.Value = "121234"
	ResponseJSON(w, response)
	w.WriteHeader(http.StatusOK)
}

//ActualizarAvatarUsingPOST , Actualiza Foto de Avatar
//curl -X POST "https://localhost:8080/v1/clientes/persona/perfil/avatar" -H "accept: application/json" -H "Authorization: 2132112" -H "idApp: 32234" -H "sucursal: 2342" -H "Content-Type: application/json" -d "{ \"contentType\": \"string\", \"image\": \"string\", \"rut\": \"string\"}"
func ActualizarAvatarUsingPOST(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	w.WriteHeader(http.StatusOK)
}

func ObtenerAvatarUsingGET(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	url := "http://www.thispersondoesnotexist.com/image"
	// don't worry about errors
	response, e := http.Get(url)
	if e != nil {
		log.Fatal(e)
	}
	defer response.Body.Close()

	//open a file for writing
	file, err := os.Create("avatar.tmp.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Use io.Copy to just dump the response body to the file. This supports huge files
	_, err = io.Copy(file, response.Body)
	if err != nil {
		log.Fatal(err)
	}

	img, err := os.Open("avatar.tmp.jpg")
	if err != nil {
		log.Fatal(err) // perhaps handle this nicer
	}
	defer img.Close()
	w.Header().Set("Content-Type", "image/jpeg") // <-- set the content-type header
	io.Copy(w, img)

	w.WriteHeader(http.StatusOK)
}

func ObtenerPerfilUsingGET(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	response := ActualizaPerfilUsingResponse{}
	response.Alias = "cdiaz"
	response.EmailParticularCliente = "cdiaz@gmail.com"
	response.Nombre = "Carlos Diaz"
	response.Rut = "15125990-1"
	response.Sexo = "M"
	response.Segmento.Key = "PER"
	response.Segmento.Value = "Persona Natural"
	response.Banca.Key = "CTA"
	response.Banca.Value = "121234"
	ResponseJSON(w, response)
	w.WriteHeader(http.StatusOK)
}
