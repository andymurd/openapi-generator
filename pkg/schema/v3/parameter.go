package v3

import (
	"fmt"

	"github.com/andymurd/openapi-generator/pkg/schema"
	v3high "github.com/pb33f/libopenapi/datamodel/high/v3"
)

type Parameter struct {
	v3high.Parameter
	schema.ObjectBase
}

func (this *Parameter) ObjectTypeName() schema.ObjectTypeName {
	return "parameter"
}

func (this *Parameter) Wrap(parent schema.ObjectInterface) {
	fmt.Printf("Wrapping %s %p\n", this.ObjectTypeName(), this)
	
	this.ObjectBase.SetParent(parent)
	var children []schema.ObjectInterface

	if this.Parameter.Schema != nil {
		oasSchema := Schema{*this.Parameter.Schema.Schema(), schema.ObjectBase{}}
		oasSchema.Wrap(this)
		children = append(children, &oasSchema)
	}

	this.ObjectBase.SetChildren(children)
}
