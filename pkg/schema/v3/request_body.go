package v3

import (
	"fmt"
	
	"github.com/andymurd/openapi-generator/pkg/schema"
	v3high "github.com/pb33f/libopenapi/datamodel/high/v3"
)

type RequestBody struct {
	v3high.RequestBody
	schema.ObjectBase
}

func (this *RequestBody) ObjectTypeName() schema.ObjectTypeName {
	return "request-body"
}

func (this *RequestBody) Wrap(parent schema.ObjectInterface) {
	fmt.Printf("Wrapping %s %p\n", this.ObjectTypeName(), this)
	
	this.ObjectBase.SetParent(parent)
	var children []schema.ObjectInterface

	for _, oasMediaType := range this.RequestBody.Content {
		mediaType := MediaType{*oasMediaType, schema.ObjectBase{}}
		mediaType.Wrap(this)
		children = append(children, &mediaType)
	}

	this.ObjectBase.SetChildren(children)
}
