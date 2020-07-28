package viewmodel

// Home defines data for the home page
type Home struct {
	Title  string
	Active string
}

// NewHome is the function to load data into the Home struct
func NewHome() Home {
	result := Home{
		Active: "home",
		Title:  "AR CMS",
	}
	return result
}
