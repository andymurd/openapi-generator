package v3

import (
	"fmt"
	
	"github.com/andymurd/openapi-generator/pkg/schema"
	v3high "github.com/pb33f/libopenapi/datamodel/high/v3"
)

type PathItem struct {
	v3high.PathItem
	schema.ObjectBase
}

func (this *PathItem) ObjectTypeName() schema.ObjectTypeName {
	return "path-item"
}

func (this *PathItem) Wrap(parent schema.ObjectInterface) {
	fmt.Printf("Wrapping %s %p\n", this.ObjectTypeName(), this)
	
	this.ObjectBase.SetParent(parent)
	var children []schema.ObjectInterface

	if this.PathItem.Get != nil {
		operation := Operation{*this.PathItem.Get, schema.ObjectBase{}}
		operation.Wrap(this)
		children = append(children, &operation)
	}

	if this.PathItem.Put != nil {
		operation := Operation{*this.PathItem.Put, schema.ObjectBase{}}
		operation.Wrap(this)
		children = append(children, &operation)
	}

	if this.PathItem.Post != nil {
		operation := Operation{*this.PathItem.Post, schema.ObjectBase{}}
		operation.Wrap(this)
		children = append(children, &operation)
	}

	if this.PathItem.Delete != nil {
		operation := Operation{*this.PathItem.Delete, schema.ObjectBase{}}
		operation.Wrap(this)
		children = append(children, &operation)
	}

	if this.PathItem.Patch != nil {
		operation := Operation{*this.PathItem.Patch, schema.ObjectBase{}}
		operation.Wrap(this)
		children = append(children, &operation)
	}

	if this.PathItem.Head != nil {
		operation := Operation{*this.PathItem.Head, schema.ObjectBase{}}
		operation.Wrap(this)
		children = append(children, &operation)
	}

	if this.PathItem.Options != nil {
		operation := Operation{*this.PathItem.Options, schema.ObjectBase{}}
		operation.Wrap(this)
		children = append(children, &operation)
	}

	if this.PathItem.Trace != nil {
		operation := Operation{*this.PathItem.Trace, schema.ObjectBase{}}
		operation.Wrap(this)
		children = append(children, &operation)
	}

	for _, oasServer := range this.PathItem.Servers {
		server := Server{*oasServer, schema.ObjectBase{}}
		server.Wrap(this)
		children = append(children, &server)
	}

	for _, oasParameter := range this.PathItem.Parameters {
		parameter := Parameter{*oasParameter, schema.ObjectBase{}}
		parameter.Wrap(this)
		children = append(children, &parameter)
	}

	this.ObjectBase.SetChildren(children)
}
