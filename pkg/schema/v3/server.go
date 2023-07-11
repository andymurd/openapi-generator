package v3

import (
	"fmt"
	
	"github.com/andymurd/openapi-generator/pkg/schema"
	v3high "github.com/pb33f/libopenapi/datamodel/high/v3"
)

type Server struct {
	v3high.Server
	schema.ObjectBase
}

func (this *Server) ObjectTypeName() schema.ObjectTypeName {
	return "server"
}

func (this *Server) Wrap(parent schema.ObjectInterface) {
	fmt.Printf("Wrapping %s %p\n", this.ObjectTypeName(), this)
	
	this.ObjectBase.SetParent(parent)
	var children []schema.ObjectInterface

	// FIXME

	this.ObjectBase.SetChildren(children)
}
