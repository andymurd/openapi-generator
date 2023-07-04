package schema

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type TestSchema1 struct {
	ObjectBase
}

func (this *TestSchema1) ObjectTypeName() ObjectTypeName {
	return "test1"
}

type TestSchema2 struct {
	ObjectBase
}

func (this *TestSchema2) ObjectTypeName() ObjectTypeName {
	return "test2"
}

func TestSpec(t *testing.T) {
	Convey("A single schema object", t, func() {
		out := &ObjectBase{}

		So(out.Root(), ShouldEqual, out)
		So(len(out.Children()), ShouldEqual, 0)
		So(len(out.ChildrenByTypeName("test1")), ShouldEqual, 0)

		So(out.Parent(), ShouldBeNil)
		So(len(out.AllParents()), ShouldEqual, 0)
		So(len(out.AllParentsByTypeName("test1")), ShouldEqual, 0)
		So(out.NearestParentByTypeName("test1"), ShouldBeNil)
	})

	Convey("A schema object with no parents, one child", t, func() {
		child := &TestSchema1{}
		out := &ObjectBase{
			parent:   nil,
			children: []ObjectInterface{child},
		}

		So(out.Root(), ShouldEqual, out)
		So(len(out.Children()), ShouldEqual, 1)
		So(out.Children()[0], ShouldEqual, child)
		So(len(out.ChildrenByTypeName("test1")), ShouldEqual, 1)
		So(out.ChildrenByTypeName("test1")[0], ShouldEqual, child)
		So(len(out.ChildrenByTypeName("not-test")), ShouldEqual, 0)

		So(out.Parent(), ShouldBeNil)
		So(len(out.AllParents()), ShouldEqual, 0)
		So(len(out.AllParentsByTypeName("test1")), ShouldEqual, 0)
		So(out.NearestParentByTypeName("test1"), ShouldBeNil)
	})

	Convey("A schema object with no parents, two children", t, func() {
		child1 := &TestSchema1{}
		child2 := &TestSchema2{}
		out := &ObjectBase{
			parent:   nil,
			children: []ObjectInterface{child1, child2},
		}

		So(out.Root(), ShouldEqual, out)
		So(len(out.Children()), ShouldEqual, 2)
		So(out.Children()[0], ShouldEqual, child1)
		So(out.Children()[1], ShouldEqual, child2)
		So(len(out.ChildrenByTypeName("test1")), ShouldEqual, 1)
		So(out.ChildrenByTypeName("test1")[0], ShouldEqual, child1)
		So(len(out.ChildrenByTypeName("test2")), ShouldEqual, 1)
		So(out.ChildrenByTypeName("test2")[0], ShouldEqual, child2)
		So(len(out.ChildrenByTypeName("not-test")), ShouldEqual, 0)

		So(out.Parent(), ShouldBeNil)
		So(len(out.AllParents()), ShouldEqual, 0)
		So(len(out.AllParentsByTypeName("test1")), ShouldEqual, 0)
		So(out.NearestParentByTypeName("test1"), ShouldBeNil)
	})

	Convey("A schema object with one parent, no children", t, func() {
		parent := &TestSchema1{}
		out := &ObjectBase{
			parent: parent,
		}

		So(out.Root(), ShouldEqual, parent)
		So(len(out.Children()), ShouldEqual, 0)

		So(out.Parent(), ShouldEqual, parent)
		So(len(out.AllParents()), ShouldEqual, 1)
		So(out.AllParents()[0], ShouldEqual, parent)
		So(len(out.AllParentsByTypeName("test1")), ShouldEqual, 1)
		So(out.AllParentsByTypeName("test1")[0], ShouldEqual, parent)
		So(len(out.AllParentsByTypeName("not-test")), ShouldEqual, 0)
		So(out.NearestParentByTypeName("test1"), ShouldEqual, parent)
	})

	Convey("A schema object with parent and grandparent, no children", t, func() {
		grandparent := &TestSchema2{}
		parent := &TestSchema1{
			ObjectBase{
				parent: grandparent,
			},
		}
		grandparent.children = []ObjectInterface{parent}
		out := &ObjectBase{
			parent: parent,
		}

		So(out.Root(), ShouldEqual, grandparent)
		So(len(out.Children()), ShouldEqual, 0)

		So(out.Parent(), ShouldEqual, parent)
		So(len(out.AllParents()), ShouldEqual, 2)
		So(out.AllParents()[0], ShouldEqual, parent)
		So(out.AllParents()[1], ShouldEqual, grandparent)
		So(len(out.AllParentsByTypeName("test1")), ShouldEqual, 1)
		So(out.AllParentsByTypeName("test1")[0], ShouldEqual, parent)
		So(len(out.AllParentsByTypeName("test2")), ShouldEqual, 1)
		So(out.AllParentsByTypeName("test2")[0], ShouldEqual, grandparent)
		So(len(out.AllParentsByTypeName("not-test")), ShouldEqual, 0)
		So(out.NearestParentByTypeName("test1"), ShouldEqual, parent)
		So(out.NearestParentByTypeName("test2"), ShouldEqual, grandparent)
	})
}
