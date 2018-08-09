package roawa

// RoaAddress used for data returned from CH Registered Office Address API
type RoaAddress struct {
	Premises     string `json:"premises"`
	Postcode     string `json:"postal_code"`
	AddressLine1 string `json:"address_line_1"`
	AddressLine2 string `json:"address_line_2"`
	Locality     string `json:"locality"`
	Region       string `json:"region"`
	Country      string `json:"country"`
	POBox        string `json:"po_box"`
}
