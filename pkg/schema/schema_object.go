package schema

import ()

// Used to uniquely identify schema object types - corresponds to the OpenAPI Spec ObjectInterface
// name, e.g. "Schema" or "Parameter"
type ObjectTypeName string

type ObjectInterface interface {
	// The type name of this object - corresponds to the OpenAPI Spec ObjectInterface
	// name, e.g. "Schema" or "Parameter"
	ObjectTypeName() ObjectTypeName

	// Get all the immediate child schema objects, regardless of their types.
	Children() []ObjectInterface

	// Get all the immediate child schema objects, filtered by their type name.
	ChildrenByTypeName(ObjectTypeName) []ObjectInterface

	// Get the root schema object.
	Root() ObjectInterface

	// Get the immediate parent schema object.
	Parent() ObjectInterface

	// Get all the parents all the way up to the root object, immediate
	// parent first in the array, then its parent, and so on.
	AllParents() []ObjectInterface

	// Get all the parents filtered by type name, all the way up to the root
	AllParentsByTypeName(ObjectTypeName) []ObjectInterface

	// Get the nearest parent that matches the supplied type name
	NearestParentByTypeName(ObjectTypeName) ObjectInterface
}

type ObjectBase struct {
	// The base type of all schema objects, it implements all the methods
	// required by the ObjectInterface interface
	parent   ObjectInterface
	children []ObjectInterface
}

func (this *ObjectBase) ObjectTypeName() ObjectTypeName {
	panic("Missing ObjectTypeName method")
}

// Get the root schema object.
func (this *ObjectBase) Root() ObjectInterface {
	candidate := ObjectInterface(this)
	for candidate.Parent() != nil {
		candidate = candidate.Parent()
	}
	return candidate
}

// Get the immediate parent schema object.
func (this *ObjectBase) Parent() ObjectInterface {
	return this.parent
}

// Get all the immediate child schema objects, regardless of their types.
func (this *ObjectBase) Children() []ObjectInterface {
	retval := make([]ObjectInterface, len(this.children))
	for index, child := range this.children {
		retval[index] = child
	}
	return retval
}

// Get all the immediate child schema objects, filtered by their type name.
func (this *ObjectBase) ChildrenByTypeName(typeName ObjectTypeName) []ObjectInterface {
	retval := []ObjectInterface{}
	for _, child := range this.Children() {
		if child.ObjectTypeName() == typeName {
			retval = append(retval, child)
		}
	}
	return retval
}

// Get all the parents all the way up to the root object, immediate
// parent first in the array, then its parent, and so on.
func (this *ObjectBase) AllParents() []ObjectInterface {
	retval := []ObjectInterface{}
	if this.parent != nil {
		retval = append(retval, this.parent)
		retval = append(retval, this.parent.AllParents()...)
	}
	return retval
}

// Get all the parents filtered by type name, all the way up to the root
func (this *ObjectBase) AllParentsByTypeName(typeName ObjectTypeName) []ObjectInterface {
	retval := []ObjectInterface{}
	for _, child := range this.AllParents() {
		if child.ObjectTypeName() == typeName {
			retval = append(retval, child)
		}
	}
	return retval
}

// Get the nearest parent that matches the supplied type name
func (this *ObjectBase) NearestParentByTypeName(typeName ObjectTypeName) ObjectInterface {
	parent := this.Parent()
	for parent != nil {
		if parent.ObjectTypeName() == typeName {
			return parent
		}
		parent = parent.Parent()
	}
	return nil
}
