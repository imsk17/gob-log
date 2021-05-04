package entities

type Blog struct {
	Title       string
	Summary     string
	PublishDate string
	Tags        []interface{}
	Content     string
	Image       string
	Href        string
	SkimTime    string
}
