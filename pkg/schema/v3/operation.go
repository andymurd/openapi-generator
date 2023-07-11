package v3

import (
	"fmt"
	
	"github.com/andymurd/openapi-generator/pkg/schema"
	v3high "github.com/pb33f/libopenapi/datamodel/high/v3"
)

type Operation struct {
	v3high.Operation
	schema.ObjectBase
}

func (this *Operation) ObjectTypeName() schema.ObjectTypeName {
	return "operation"
}

func (this *Operation) Wrap(parent schema.ObjectInterface) {
	fmt.Printf("Wrapping %s %p\n", this.ObjectTypeName(), this)
	
	this.ObjectBase.SetParent(parent)
	var children []schema.ObjectInterface

	if this.Operation.RequestBody != nil {
		oasRequestBody := RequestBody{*this.Operation.RequestBody, schema.ObjectBase{}}
		oasRequestBody.Wrap(this)
		children = append(children, &oasRequestBody)
	}

	if this.Operation.Responses != nil {
		oasResponses := Responses{*this.Operation.Responses, schema.ObjectBase{}}
		oasResponses.Wrap(this)
		children = append(children, &oasResponses)
	}

	for _, oasServer := range this.Operation.Servers {
		server := Server{*oasServer, schema.ObjectBase{}}
		server.Wrap(this)
		children = append(children, &server)
	}

	for _, oasParameter := range this.Operation.Parameters {
		parameter := Parameter{*oasParameter, schema.ObjectBase{}}
		parameter.Wrap(this)
		children = append(children, &parameter)
	}

	this.ObjectBase.SetChildren(children)
}
