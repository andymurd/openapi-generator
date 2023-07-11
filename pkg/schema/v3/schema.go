package v3

import (
	"fmt"

	"github.com/andymurd/openapi-generator/pkg/schema"
	v3base "github.com/pb33f/libopenapi/datamodel/high/base"
)

type Schema struct {
	v3base.Schema
	schema.ObjectBase
}

func (this *Schema) ObjectTypeName() schema.ObjectTypeName {
	return "schema"
}

func (this *Schema) Wrap(parent schema.ObjectInterface) {
	fmt.Printf("Wrapping %s %p\n", this.ObjectTypeName(), this)
	
	this.ObjectBase.SetParent(parent)

	// No children
}
