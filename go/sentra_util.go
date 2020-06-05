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
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

//ResponseJSON : Responde en formato JSON cualquier Objeto
func ResponseJSON(w http.ResponseWriter, object interface{}) {
	fmt.Fprintln(w, StringifyJSON(object))
}

//StringifyJSON , convierte cualquier objeto en una representacion STRING en formato JSON
func StringifyJSON(object interface{}) string {
	b, err := json.Marshal(object)
	if err != nil {
		return "Error"
	}
	return string(b)
}

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func String(length int) string {
	return StringWithCharset(length, charset)
}
