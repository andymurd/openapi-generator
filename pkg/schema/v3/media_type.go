package v3

import (
	"fmt"

	"github.com/andymurd/openapi-generator/pkg/schema"
	v3high "github.com/pb33f/libopenapi/datamodel/high/v3"
)

type MediaType struct {
	v3high.MediaType
	schema.ObjectBase
}

func (this *MediaType) ObjectTypeName() schema.ObjectTypeName {
	return "media-type"
}

func (this *MediaType) Wrap(parent schema.ObjectInterface) {
	fmt.Printf("Wrapping %s %p\n", this.ObjectTypeName(), this)
	
	this.ObjectBase.SetParent(parent)
	var children []schema.ObjectInterface

	fmt.Printf("\tMediaType.Schema %v\n", this.MediaType.Schema)
	if this.MediaType.Schema != nil {
		oasSchema := Schema{*this.MediaType.Schema.Schema(), schema.ObjectBase{}}
		oasSchema.Wrap(this)
		children = append(children, &oasSchema)
	}

	for _, oasEncoding := range this.MediaType.Encoding {
		encoding := Encoding{*oasEncoding, schema.ObjectBase{}}
		encoding.Wrap(this)
		children = append(children, &encoding)
	}

	this.ObjectBase.SetChildren(children)
}
