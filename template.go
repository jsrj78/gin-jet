package jet

import (
	"github.com/CloudyKit/jet"
	"net/http"
)

type JetTemplate struct {
	Template *jet.Template
	Date     jet.VarMap
}

func NewJetTemplate(template *jet.Template, data jet.VarMap) *JetTemplate {
	return &JetTemplate{
		Template: template,
		Date:     data,
	}
}

func (template *JetTemplate) Render(w http.ResponseWriter) error {
	return template.Template.Execute(w, template.Date, nil)
}

func (template *JetTemplate) WriteContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html")
}
