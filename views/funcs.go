package views

import (
	"html/template"

	"strings"
)

// FuncMap passes functions to templates
var FuncMap = template.FuncMap{
	"title": Title,
}

// Title uses first character to capital letter
func Title(s string) string {
	return strings.Title(s)
}
