package models

type Plan struct {
	Popular  bool
	PlanName string
	Price    string
	Features []string
	CTAText  string
}

type Feature struct {
	Icon        string
	Title       string
	Description string
}

type Testimonial struct {
	Text   string
	Author string
}
