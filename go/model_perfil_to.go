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



type PerfilTo struct {

	// alias asociado al usuario
	Alias string `json:"alias"`

	// Par código/texto con el banca
	Banca *Entrystringstring `json:"banca,omitempty"`

	// Email del cliente
	EmailParticularCliente string `json:"emailParticularCliente,omitempty"`

	// Par código/texto con el marca
	Marca *Entrystringstring `json:"marca,omitempty"`

	// nombre completo obtenido desde ficha
	Nombre string `json:"nombre,omitempty"`

	// rut del cliente en formato 1236123-K
	Rut string `json:"rut"`

	// Par código/texto con el segmento
	Segmento *Entrystringstring `json:"segmento,omitempty"`

	// Sexo del cliente
	Sexo string `json:"sexo,omitempty"`
}
