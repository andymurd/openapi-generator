package v3

import (
	"fmt"
	
	"github.com/andymurd/openapi-generator/pkg/schema"
	v3high "github.com/pb33f/libopenapi/datamodel/high/v3"
)

type Response struct {
	v3high.Response
	schema.ObjectBase
}

func (this *Response) ObjectTypeName() schema.ObjectTypeName {
	return "response"
}

func (this *Response) Wrap(parent schema.ObjectInterface) {
	fmt.Printf("Wrapping %s %p\n", this.ObjectTypeName(), this)
	
	this.ObjectBase.SetParent(parent)
	var children []schema.ObjectInterface

	for _, oasHeader := range this.Response.Headers {
		header := Header{*oasHeader, schema.ObjectBase{}}
		header.Wrap(this)
		children = append(children, &header)
	}

	for _, oasMediaType := range this.Response.Content {
		mediaType := MediaType{*oasMediaType, schema.ObjectBase{}}
		mediaType.Wrap(this)
		children = append(children, &mediaType)
	}

	this.ObjectBase.SetChildren(children)
}
