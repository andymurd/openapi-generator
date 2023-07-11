package v3

import (
	"fmt"
	
	"github.com/andymurd/openapi-generator/pkg/schema"
	v3base "github.com/pb33f/libopenapi/datamodel/high/base"
)

type ExternalDocs struct {
	v3base.ExternalDoc
	schema.ObjectBase
}

func (this *ExternalDocs) ObjectTypeName() schema.ObjectTypeName {
	return "external-docs"
}

func (this *ExternalDocs) Wrap(parent schema.ObjectInterface) {
	fmt.Printf("Wrapping %s %p\n", this.ObjectTypeName(), this)
	
	this.ObjectBase.SetParent(parent)
	var children []schema.ObjectInterface

	// FIXME

	this.ObjectBase.SetChildren(children)
}
