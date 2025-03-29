package pages

import (
	"html/template"
	"log"
	"path/filepath"

	"metascribe/templates/models"

	"github.com/labstack/echo/v4"
)

var TemplateDir string

// HomeHandler for Echo
func HomeHandler(c echo.Context) error {
	// Load data
	data := struct {
		Title        string
		Features     []models.Feature
		Plans        []models.Plan
		Testimonials []models.Testimonial
	}{
		Title: "Home | MetaScribe",
		Features: []models.Feature{
			f.UniversalConnectivity,
			f.SmartFormatting,
			f.BidirectionalSync,
			f.VersionControl,
			f.EnterpriseSecurity,
			f.AutomationReady,
		},
		Plans: []models.Plan{
			p.Free,
			p.Professional,
			p.Team,
		},
		Testimonials: []models.Testimonial{
			t.Jane,
			t.Mark,
			t.Alex,
		},
	}

	// Compute absolute path to the templates directory
	var err error
	TemplateDir, err = filepath.Abs("templates")
	if err != nil {
		log.Fatalf("Error getting absolute template path: %v", err)
	}

	log.Printf("Using templates from: %s", TemplateDir)

	// Parse templates
	tmpl, err := template.ParseFiles(
		filepath.Join(TemplateDir, "base.html"),
		filepath.Join(TemplateDir, "header.html"),
		filepath.Join(TemplateDir, "footer.html"),
		filepath.Join(TemplateDir, "home.html"),
		filepath.Join(TemplateDir, "components.html"),
	)
	if err != nil {
		c.Logger().Errorf("Template parsing error: %v", err)
		return c.String(500, "Error loading templates: "+err.Error())
	}

	// 🛠️ Execute template and send response
	err = tmpl.ExecuteTemplate(c.Response(), "base.html", data)
	if err != nil {
		c.Logger().Errorf("Template execution error: %v", err)
		return c.String(500, "Error rendering template: "+err.Error())
	}

	return nil
}

// ✅ Explicitly reference `models.Feature`, `models.Plan`, `models.Testimonial`

// Features
var f = struct {
	UniversalConnectivity models.Feature
	SmartFormatting       models.Feature
	BidirectionalSync     models.Feature
	VersionControl        models.Feature
	EnterpriseSecurity    models.Feature
	AutomationReady       models.Feature
}{
	UniversalConnectivity: models.Feature{
		Icon:        "🔗",
		Title:       "Universal Connectivity",
		Description: "Seamlessly integrate with Google Docs, Notion, Evernote, Word, and more with our one-click connectors.",
	},
	SmartFormatting: models.Feature{
		Icon:        "✨",
		Title:       "Smart Formatting",
		Description: "Our AI-powered engine automatically maintains your brand guidelines and formatting requirements.",
	},
	BidirectionalSync: models.Feature{
		Icon:        "🔄",
		Title:       "Bi-directional Sync",
		Description: "Changes made in your CMS can be reflected back in your source documents, keeping everything in sync.",
	},
	VersionControl: models.Feature{
		Icon:        "📊",
		Title:       "Version Control",
		Description: "Track changes, compare versions, and restore previous content with our powerful versioning system.",
	},
	EnterpriseSecurity: models.Feature{
		Icon:        "🔐",
		Title:       "Enterprise Security",
		Description: "SOC 2 compliant with end-to-end encryption to keep your content secure at every step.",
	},
	AutomationReady: models.Feature{
		Icon:        "🤖",
		Title:       "Automation Ready",
		Description: "Create custom workflows with our API to automate repetitive tasks and content processing.",
	},
}

// Pricing Plans
var p = struct {
	Free         models.Plan
	Professional models.Plan
	Team         models.Plan
}{
	Free: models.Plan{
		PlanName: "Free",
		Price:    "$0",
		Features: []string{
			"Basic Google Docs integration",
			"Up to 10 content transfers/month",
			"Standard formatting options",
			"Community support",
		},
		CTAText: "Get Started",
	},
	Professional: models.Plan{
		PlanName: "Professional",
		Price:    "$29",
		Popular:  true,
		Features: []string{
			"All Free features plus:",
			"Unlimited content transfers",
			"Advanced formatting controls",
			"Connect to any major CMS",
			"Version history & rollbacks",
			"Priority email support",
		},
		CTAText: "Start Free Trial",
	},
	Team: models.Plan{
		PlanName: "Team",
		Price:    "$99",
		Features: []string{
			"All Professional features plus:",
			"Team collaboration tools",
			"Advanced workflow automation",
			"Custom formatting templates",
			"SSO & advanced security",
			"Dedicated account manager",
		},
		CTAText: "Contact Sales",
	},
}

// Testimonials
var t = struct {
	Jane models.Testimonial
	Mark models.Testimonial
	Alex models.Testimonial
}{
	Jane: models.Testimonial{
		Text:   "MetaScribe has completely transformed our content workflow. We've reduced our publishing time by 65% and eliminated countless formatting errors.",
		Author: "Jane Smith, Content Director, TechCorp",
	},
	Mark: models.Testimonial{
		Text:   "As a remote content team working across three time zones, MetaScribe has been a game-changer for our collaboration and consistency.",
		Author: "Mark Stevens, Managing Editor, Global Media",
	},
	Alex: models.Testimonial{
		Text:   "The ROI on MetaScribe was immediate. What used to take our team hours now happens automatically, and with better results than manual formatting.",
		Author: "Alex Lee, Operations Manager, eCommerce Brand",
	},
}
