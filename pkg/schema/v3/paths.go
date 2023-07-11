package v3

import (
	"fmt"
	
	"github.com/andymurd/openapi-generator/pkg/schema"
	v3high "github.com/pb33f/libopenapi/datamodel/high/v3"
)

type Paths struct {
	v3high.Paths
	schema.ObjectBase
}

func (this *Paths) ObjectTypeName() schema.ObjectTypeName {
	return "paths"
}

func (this *Paths) Wrap(parent schema.ObjectInterface) {
	fmt.Printf("Wrapping %s %p\n", this.ObjectTypeName(), this)
	
	this.ObjectBase.SetParent(parent)
	var children []schema.ObjectInterface

	for _, oasPathItem := range this.Paths.PathItems {
		pathItem := PathItem{*oasPathItem, schema.ObjectBase{}}
		pathItem.Wrap(this)
		children = append(children, &pathItem)
	}

	this.ObjectBase.SetChildren(children)
}
