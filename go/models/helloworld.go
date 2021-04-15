package models

type HowDoYouSayHello struct {
	Name    string
	Country *Country
}

type Country struct {
	Name string
}
