package router

import (
	"fmt"
	"io"
	"os/exec"
	"path/filepath"
	"runtime"
	"text/template"

	"github.com/labstack/echo/v4"
)

type WebRenderer struct {
	templates *template.Template
}

func NewWebRenderer(templateDir string) *WebRenderer {
	templates := template.Must(template.ParseGlob(filepath.Join(templateDir, "*.html")))
	return &WebRenderer{templates: templates}
}

func (t *WebRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}
	return t.templates.ExecuteTemplate(w, name, data)
}
func openBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}

	if err != nil {
		fmt.Printf("Failed to open browser: %v\n", err)
	}
}
