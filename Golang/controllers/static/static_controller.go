package static

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const templateDir = "./static/templates"

func StaticController(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	path = strings.TrimSuffix(path, "/")
	if path == "" || path == "/" {
		path = "/index"
	}

	file := filepath.Join(templateDir, path+".html")
	clean := filepath.Clean(file)

	// ✅ Windows-safe absolute path security check
	absBase, _ := filepath.Abs(templateDir)
	absFile, _ := filepath.Abs(clean)

	if !strings.HasPrefix(absFile, absBase) {
		http.Error(w, "Access denied", http.StatusForbidden)
		return
	}

	// fallback to index if not found
	if _, err := os.Stat(clean); os.IsNotExist(err) {
		clean = filepath.Join(templateDir, "index.html")
	}

	http.ServeFile(w, r, clean)
}
