package models

// swagger:model HowDoYouSayHello
type HowDoYouSayHello struct {
	Name    string
	Country *Country
}

// swagger:model Country
type Country struct {
	Name string
}
