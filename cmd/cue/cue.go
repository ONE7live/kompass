package main

import (
	"cuelang.org/go/cue"
	"fmt"
	"os"

	app "github.com/ONE7live/kompass/pkg/cue"
)

type Parameter struct {
	Name     string      `json:"name"`
	Short    string      `json:"short,omitempty"`
	Required bool        `json:"required,omitempty"`
	Default  interface{} `json:"default,omitempty"`
	Usage    string      `json:"usage,omitempty"`
	Ignore   bool        `json:"ignore,omitempty"`
	Type     cue.Kind    `json:"type,omitempty"`
	Alias    string      `json:"alias,omitempty"`
	JSONType string      `json:"jsonType,omitempty"`
}

func main() {
	content, err := os.ReadFile("D:\\code\\go\\src\\kosmos\\self\\kompass\\cmd\\cue\\pod.cue")
	if err != nil {
		fmt.Errorf("the cue script is invalid:%w", err)
	}

	value, err := app.ParseToTemplateValue(content)
	if err != nil {
		fmt.Errorf("the cue script is invalid:%w", err)
	}

	templateValue, err := value.LookupValue("template")
	if err != nil {
		fmt.Errorf("the cue script is invalid:%w", err)
	}

	paramVal, err := templateValue.LookupValue("parameter")
	if err != nil || !paramVal.CueValue().Exists() {
		fmt.Errorf("the cue script is invalid:%w", err)
	}

	iter, err := paramVal.CueValue().Fields(cue.Definitions(true), cue.Hidden(true), cue.All())
	if err != nil {
		fmt.Errorf("the cue script is invalid:%w", err)
	}
	// parse each fields in the parameter fields
	var params []Parameter
	for iter.Next() {
		if iter.Selector().IsDefinition() {
			continue
		}
		var param = Parameter{
			Name:     iter.Label(),
			Required: !iter.IsOptional(),
		}
		val := iter.Value()
		param.Type = val.IncompleteKind()
		if def, ok := val.Default(); ok && def.IsConcrete() {
			param.Required = false
			param.Type = def.Kind()
		}

		params = append(params, param)
	}

	fmt.Println("【TemplateValue】:", templateValue)
	fmt.Println("【Result】:", "")
}
