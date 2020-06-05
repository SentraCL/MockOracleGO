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
	"net/http"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

//GetProductosClienteUsingGET ,Devuelve los productos del cliente
//curl -X GET "https://localhost:8080/v1/productos/0101001010" -H "accept: application/json" -H "Authorization: dsf" -H "idApp: sdf"
func GetProductosClienteUsingGET(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	productoResponse := ProductoResponse{}
	productoResponse.ClaseCuenta = "cte"
	productoResponse.CodigoProducto = "0101010"
	productoResponse.Descripcion = ""
	productoResponse.Email = "ficticio@gmail.com" //StringWithCharset(100, charset)
	productoResponse.Estado = "S"
	productoResponse.Familia = "s"
	productoResponse.FlagBloqueo = "s"
	productoResponse.FlagInhibicion = "s"
	productoResponse.Marca = StringWithCharset(10, charset)
	productoResponse.Moneda = "CLP"
	productoResponse.NumProducto = "96794387"
	productoResponse.NumProductoEncriptado = StringWithCharset(32, charset)
	productoResponse.RedComercialTarjeta = ""
	productoResponse.RowIdEmail = "0"
	productoResponse.RowIdProducto = "1"
	productoResponse.Tipo = "P"
	productoResponse.TipoPersona = "Natural"
	productoResponse.TipoProducto = "CUENTA_CORRIENTE"
	productoResponse.TipoTarjeta = "ctecte"
	productoResponse.Titularidad = "OK"
	ResponseJSON(w, productoResponse)
	w.WriteHeader(http.StatusOK)
}
