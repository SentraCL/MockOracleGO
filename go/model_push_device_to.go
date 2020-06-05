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



type PushDeviceTo struct {

	// identificador de plataforma: android (G), ios (A)
	Platform string `json:"platform"`

	// token de push FCM
	PushToken string `json:"pushToken"`
}
