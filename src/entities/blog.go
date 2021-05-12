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
	Author      string
}

type BlogMeta struct {
	Title       string
	PublishDate string
	Tags        []string
	SkimTime    string
	Href        string
}
