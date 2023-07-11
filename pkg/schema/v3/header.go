package v3

import (
	"fmt"
	"github.com/andymurd/openapi-generator/pkg/schema"
	v3high "github.com/pb33f/libopenapi/datamodel/high/v3"
)

type Header struct {
	v3high.Header
	schema.ObjectBase
}

func (this *Header) ObjectTypeName() schema.ObjectTypeName {
	return "header"
}

func (this *Header) Wrap(parent schema.ObjectInterface) {
	fmt.Printf("Wrapping %s %p\n", this.ObjectTypeName(), this)
	
	this.ObjectBase.SetParent(parent)
	var children []schema.ObjectInterface

	if this.Header.Schema != nil {
		oasSchema := Schema{*this.Header.Schema.Schema(), schema.ObjectBase{}}
		oasSchema.Wrap(this)
		children = append(children, &oasSchema)
	}

	for _, oasMediaType := range this.Header.Content {
		mediaType := MediaType{*oasMediaType, schema.ObjectBase{}}
		mediaType.Wrap(this)
		children = append(children, &mediaType)
	}

	this.ObjectBase.SetChildren(children)
}
