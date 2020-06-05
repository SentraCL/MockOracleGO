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



type QName struct {

	LocalPart string `json:"localPart,omitempty"`

	NamespaceURI string `json:"namespaceURI,omitempty"`

	Prefix string `json:"prefix,omitempty"`
}
