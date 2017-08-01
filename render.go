package jet

import (
	"github.com/CloudyKit/jet"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin/render"
)

type JetRender struct {
	Set *jet.Set
}

func NewJetRender(set *jet.Set) *JetRender {
	return &JetRender{
		Set: set,
	}
}

func (r *JetRender) Instance(name string, data interface{}) render.Render {
	template, err := r.Set.GetTemplate(name)
	if err != nil {
		panic(err)
	}

	if vars, ok := data.(jet.VarMap); ok {
		return NewJetTemplate(template, vars)
	}

	varMap, ok := data.(map[string]interface{})
	if !ok {
		varMap = structs.Map(data)
	}

	vars := make(jet.VarMap)

	for key, value := range varMap {
		vars.Set(key, value)
	}

	return NewJetTemplate(template, vars)
}
