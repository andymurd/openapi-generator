package v3

import (
	"fmt"
	
	"github.com/andymurd/openapi-generator/pkg/schema"
	v3base "github.com/pb33f/libopenapi/datamodel/high/base"
)

type Tag struct {
	v3base.Tag
	schema.ObjectBase
}

func (this *Tag) ObjectTypeName() schema.ObjectTypeName {
	return "tag"
}

func (this *Tag) Wrap(parent schema.ObjectInterface) {
	fmt.Printf("Wrapping %s %p\n", this.ObjectTypeName(), this)
	
	this.ObjectBase.SetParent(parent)
	var children []schema.ObjectInterface

	// FIXME

	this.ObjectBase.SetChildren(children)
}
