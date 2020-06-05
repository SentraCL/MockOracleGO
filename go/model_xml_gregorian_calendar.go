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



type XmlGregorianCalendar struct {

	Day int32 `json:"day,omitempty"`

	Eon int32 `json:"eon,omitempty"`

	EonAndYear int32 `json:"eonAndYear,omitempty"`

	FractionalSecond float32 `json:"fractionalSecond,omitempty"`

	Hour int32 `json:"hour,omitempty"`

	Millisecond int32 `json:"millisecond,omitempty"`

	Minute int32 `json:"minute,omitempty"`

	Month int32 `json:"month,omitempty"`

	Second int32 `json:"second,omitempty"`

	Timezone int32 `json:"timezone,omitempty"`

	Valid bool `json:"valid,omitempty"`

	XmlschemaType *QName `json:"xmlschemaType,omitempty"`

	Year int32 `json:"year,omitempty"`
}
