package v3

import (
	"fmt"

	"github.com/andymurd/openapi-generator/pkg/schema"
	v3high "github.com/pb33f/libopenapi/datamodel/high/v3"
)

type Responses struct {
	v3high.Responses
	schema.ObjectBase
}

func (this *Responses) ObjectTypeName() schema.ObjectTypeName {
	return "responses"
}

func (this *Responses) Wrap(parent schema.ObjectInterface) {
	fmt.Printf("Wrapping %s %p\n", this.ObjectTypeName(), this)
	
	this.ObjectBase.SetParent(parent)
	var children []schema.ObjectInterface

	if this.Responses.Default != nil {
		oasResponse := Response{*this.Responses.Default, schema.ObjectBase{}}
		oasResponse.Wrap(this)
		children = append(children, &oasResponse)
	}

	for _, oasResponse := range this.Responses.Codes {
		response := Response{*oasResponse, schema.ObjectBase{}}
		response.Wrap(this)
		children = append(children, &response)
	}

	this.ObjectBase.SetChildren(children)
}
