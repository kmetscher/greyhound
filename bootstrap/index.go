package bootstrap 

import (
    "html/template"
)

func ParseIndex() *template.Template {
    return template.Must(template.ParseFiles("index.html"))
}

