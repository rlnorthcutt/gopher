package core

import (
	"html/template"
	"io"
	"net/http"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

type TemplateEngine struct {
	templates *template.Template
}

// NewTemplateEngine loads templates from disk.
func NewTemplateEngine(pattern string) (*TemplateEngine, error) {
	tmpls, err := template.New("").
		Funcs(DefaultFuncMap()).
		ParseGlob(pattern)
	if err != nil {
		return nil, err
	}
	return &TemplateEngine{templates: tmpls}, nil
}

// Render implements echo's Renderer interface.
func (t *TemplateEngine) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// DefaultFuncMap returns helpers usable inside templates.
func DefaultFuncMap() template.FuncMap {
	return template.FuncMap{
		"safe": func(s string) template.HTML {
			return template.HTML(s)
		},
		"upper": func(s string) string {
			return strings.ToUpper(s)
		},
		// Add more helpers as needed
	}
}

func RenderPage(c echo.Context, templateName string, data interface{}) error {
	return renderPageWithCacheFlag(c, templateName, data, true)
}

func RenderPageNoCache(c echo.Context, templateName string, data interface{}) error {
	return renderPageWithCacheFlag(c, templateName, data, false)
}

func renderPageWithCacheFlag(c echo.Context, templateName string, data interface{}, cache bool) error {
	cacheKey := "page:" + c.Path() + ":" + templateName

	if cache {
		if html, found := globalCache.Get(cacheKey); found {
			return c.HTML(http.StatusOK, html.(string))
		}
	}

	rec := c.Echo().AcquireResponse()
	defer c.Echo().ReleaseResponse(rec)
	c.Response().Writer = rec

	isHTMX := c.Request().Header.Get("HX-Request") == "true"
	var err error
	if isHTMX {
		err = c.Render(http.StatusOK, templateName, data)
	} else {
		err = c.Render(http.StatusOK, "base.html", echo.Map{
			"Content": templateName,
			"Data":    data,
		})
	}
	if err != nil {
		return err
	}

	rendered := rec.Body.String()

	if cache {
		globalCache.SetWithTTL(cacheKey, rendered, 1, 5*time.Minute)
		globalCache.AddTags(cacheKey, []string{"tag:" + templateName})
	}

	return c.HTML(http.StatusOK, rendered)
}


