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

type ProductoResponse struct {
	ClaseCuenta     string                `json:"claseCuenta,omitempty"`
	CodigoProducto  string                `json:"codigoProducto,omitempty"`
	Descripcion     string                `json:"descripcion,omitempty"`
	Email           string                `json:"email,omitempty"`
	Estado          string                `json:"estado,omitempty"`
	Familia         string                `json:"familia,omitempty"`
	FechaInhibicion *XmlGregorianCalendar `json:"fechaInhibicion,omitempty"`
	FlagBloqueo     string                `json:"flagBloqueo,omitempty"`
	FlagInhibicion  string                `json:"flagInhibicion,omitempty"`
	Marca           string                `json:"marca,omitempty"`
	Moneda          string                `json:"moneda,omitempty"`
	NumProducto     string                `json:"numProducto,omitempty"`
	// Cuenta encriptada asociada al cliente
	NumProductoEncriptado string `json:"numProductoEncriptado,omitempty"`
	RedComercialTarjeta   string `json:"redComercialTarjeta,omitempty"`
	RowIdEmail            string `json:"rowIdEmail,omitempty"`
	RowIdProducto         string `json:"rowIdProducto,omitempty"`
	Tipo                  string `json:"tipo,omitempty"`
	TipoPersona           string `json:"tipoPersona,omitempty"`
	TipoProducto          string `json:"tipoProducto,omitempty"`
	TipoTarjeta           string `json:"tipoTarjeta,omitempty"`
	Titularidad           string `json:"titularidad,omitempty"`
}
