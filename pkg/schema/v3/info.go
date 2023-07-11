package v3

import (
	"fmt"
	
	"github.com/andymurd/openapi-generator/pkg/schema"
	v3base "github.com/pb33f/libopenapi/datamodel/high/base"
)

type Info struct {
	v3base.Info
	schema.ObjectBase
}

func (this *Info) ObjectTypeName() schema.ObjectTypeName {
	return "info"
}

func (this *Info) Wrap(parent schema.ObjectInterface) {
	fmt.Printf("Wrapping %s %p\n", this.ObjectTypeName(), this)
	
	this.ObjectBase.SetParent(parent)
	var children []schema.ObjectInterface

	// FIXME

	this.ObjectBase.SetChildren(children)
}
