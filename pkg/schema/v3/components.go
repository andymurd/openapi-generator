package v3

import (
	"github.com/andymurd/openapi-generator/pkg/schema"
	v3high "github.com/pb33f/libopenapi/datamodel/high/v3"
)

type Components struct {
	v3high.Components
	schema.ObjectBase
}

func (this *Components) ObjectTypeName() schema.ObjectTypeName {
	return "components"
}

func (this *Components) Wrap(parent schema.ObjectInterface) {
	this.ObjectBase.SetParent(parent)
	var children []schema.ObjectInterface

	// FIXME

	this.ObjectBase.SetChildren(children)
}
