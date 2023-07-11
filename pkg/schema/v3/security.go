package v3

import (
	"fmt"
	
	"github.com/andymurd/openapi-generator/pkg/schema"
	v3base "github.com/pb33f/libopenapi/datamodel/high/base"
)

type Security struct {
	v3base.SecurityRequirement
	schema.ObjectBase
}

func (this *Security) ObjectTypeName() schema.ObjectTypeName {
	return "security"
}

func (this *Security) Wrap(parent schema.ObjectInterface) {
	fmt.Printf("Wrapping %s %p\n", this.ObjectTypeName(), this)
	
	this.ObjectBase.SetParent(parent)
	var children []schema.ObjectInterface

	// FIXME

	this.ObjectBase.SetChildren(children)
}
