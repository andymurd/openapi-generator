package v3

import (
	"fmt"
	
	"github.com/andymurd/openapi-generator/pkg/schema"
	v3high "github.com/pb33f/libopenapi/datamodel/high/v3"
)

type Encoding struct {
	v3high.Encoding
	schema.ObjectBase
}

func (this *Encoding) ObjectTypeName() schema.ObjectTypeName {
	return "encoding"
}

func (this *Encoding) Wrap(parent schema.ObjectInterface) {
	fmt.Printf("Wrapping %s %p\n", this.ObjectTypeName(), this)
	
	this.ObjectBase.SetParent(parent)
	var children []schema.ObjectInterface

	for _, oasHeader := range this.Encoding.Headers {
		header := Header{*oasHeader, schema.ObjectBase{}}
		header.Wrap(this)
		children = append(children, &header)
	}

	this.ObjectBase.SetChildren(children)
}
