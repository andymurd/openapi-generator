package v3

import (
	"fmt"

	"github.com/andymurd/openapi-generator/pkg/schema"
	v3high "github.com/pb33f/libopenapi/datamodel/high/v3"
)

type Document struct {
	v3high.Document
	schema.ObjectBase
}

func (this *Document) ObjectTypeName() schema.ObjectTypeName {
	return "document"
}

func (this *Document) Wrap(parent schema.ObjectInterface) {
	fmt.Printf("Wrapping %s %p\n", this.ObjectTypeName(), this)
	
	this.ObjectBase.SetParent(parent)
	var children []schema.ObjectInterface

	if this.Document.Info != nil {
		info := Info{*this.Document.Info, schema.ObjectBase{}}
		info.Wrap(this)
		children = append(children, &info)
	}

	for _, oasServer := range this.Document.Servers {
		server := Server{*oasServer, schema.ObjectBase{}}
		server.Wrap(this)
		children = append(children, &server)
	}

	if this.Document.Paths != nil {
		paths := Paths{*this.Document.Paths, schema.ObjectBase{}}
		paths.Wrap(this)
		children = append(children, &paths)
	}

	if this.Document.Components != nil {
		components := Components{*this.Document.Components, schema.ObjectBase{}}
		components.Wrap(this)
		children = append(children, &components)
	}

	for _, oasSecurity := range this.Document.Security {
		security := Security{*oasSecurity, schema.ObjectBase{}}
		security.Wrap(this)
		children = append(children, &security)
	}

	for _, oasTags := range this.Document.Tags {
		tag := Tag{*oasTags, schema.ObjectBase{}}
		tag.Wrap(this)
		children = append(children, &tag)
	}

	if this.Document.ExternalDocs != nil {
		externalDocs := ExternalDocs{*this.Document.ExternalDocs, schema.ObjectBase{}}
		externalDocs.Wrap(this)
		children = append(children, &externalDocs)
	}

	this.ObjectBase.SetChildren(children)
}
