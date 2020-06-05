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



type AvatarTo struct {

	// Content type asociado al avatar
	ContentType string `json:"contentType"`

	// contenido del avatar en base64
	Image string `json:"image"`

	// rut asociado al avatar
	Rut string `json:"rut"`
}
