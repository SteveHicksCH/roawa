package roawa

// FindPageVariables edit page variable
type FindPageVariables struct {
	Title     string
	CompanyID string
	Error     string
}

// NewFindPageVariables creates a new FindPageVariables struct with default values
func NewFindPageVariables() FindPageVariables {
	f := FindPageVariables{}
	f.Title = "Find Company"
	return f
}

// EditPageVariables edit page variable
type EditPageVariables struct {
	Title        string
	Premises     string
	Postcode     string
	AddressLine1 string
	AddressLine2 string
	Town         string
	County       string
	Countries    map[string]string
	Country      string
	POBox        string
}

// NewEditPageVariables creates a new EditPageVariables struct with default values
func NewEditPageVariables() EditPageVariables {
	e := EditPageVariables{}
	e.Title = "Change of Registered office address"
	e.Countries = *GetCountries()
	return e
}

// UpdateEditPageVariables updates the receiver values with that in the RoaAddress variable
func (e *EditPageVariables) UpdateEditPageVariables(r RoaAddress) {
	e.Premises = r.Premises
	e.Postcode = r.Postcode
	e.AddressLine1 = r.AddressLine1
	e.AddressLine2 = r.AddressLine2
	e.Town = r.Locality
	e.County = r.Region
	e.Country = r.Country
	e.POBox = r.POBox
}

// ConfirmPageVariables save page variable
type ConfirmPageVariables struct {
	Title        string
	Premises     string
	Postcode     string
	AddressLine1 string
	AddressLine2 string
	Town         string
	County       string
	Countries    map[string]string
	Country      string
	POBox        string
	Reference    string
}

// NewConfirmPageVariables creates a new ConfirmPageVariables struct with default values
func NewConfirmPageVariables() ConfirmPageVariables {
	e := ConfirmPageVariables{}
	e.Title = "Confirmation of new Registered office address"
	e.Countries = *GetCountries()
	return e
}
