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

// UpdateRoaAddress updates the receiver values from what the user entered on the web form
func (r *RoaAddress) UpdateRoaAddress(c ConfirmPageVariables) {
	r.Premises = c.Premises
	r.Postcode = c.Postcode
	r.AddressLine1 = c.AddressLine1
	r.AddressLine2 = c.AddressLine2
	r.Locality = c.Town
	r.Region = c.County
	r.Country = c.Country
	r.POBox = c.POBox
}
