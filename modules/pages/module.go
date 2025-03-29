// modules/pages/module.go
package pages

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gopher/core"
)

// PagesModule handles public pages stored in the database.
type PagesModule struct{}

func (m *PagesModule) Key() string        { return "pages" }
func (m *PagesModule) Name() string       { return "Pages" }
func (m *PagesModule) Description() string { return "Provides support for public pages from the database." }

func (m *PagesModule) Init(services *core.AppServices) error {
	return nil // Nothing to do at init
}

func (m *PagesModule) RegisterRoutes(e *echo.Echo) {
	core.Route.Register("GET", "/:slug", m.HandlePageView, "pages.view", "pages")
}

func (m *PagesModule) Migrate() error {
	return core.Services().Permissions.SetRules("pages", core.RuleSet{
		List:   ptr(""), // Allow all public access
		View:   ptr(""),
		Create: ptr("@request.auth.id != ''"),
		Update: ptr("@request.auth.id != ''"),
		Delete: ptr("@request.auth.id != ''"),
	})
}

func (m *PagesModule) RegisterJobs(j *core.JobScheduler) {}

func (m *PagesModule) Update() error { return nil }

func (m *PagesModule) HandlePageView(c echo.Context) error {
	slug := c.Param("slug")

	record, err := core.Services().Data.FindFirst("pages", map[string]interface{}{
		"slug": slug,
	})
	if err != nil || record == nil {
		return c.String(http.StatusNotFound, "Page not found")
	}

	page := record.GetString("content")
	title := record.GetString("title")

	return core.RenderPage(c, "pages/view", map[string]interface{}{
		"Title":   title,
		"Content": page,
	})
}

func ptr(s string) *string {
	return &s
}

func init() {
	core.Enable(&PagesModule{})
}
