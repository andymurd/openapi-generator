package schema

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
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
	convey.Convey("A single schema object", t, func() {
		out := &ObjectBase{}

		convey.So(out.Root(), convey.ShouldEqual, out)
		convey.So(len(out.Children()), convey.ShouldEqual, 0)
		convey.So(len(out.ChildrenByTypeName("test1")), convey.ShouldEqual, 0)

		convey.So(out.Parent(), convey.ShouldBeNil)
		convey.So(len(out.AllParents()), convey.ShouldEqual, 0)
		convey.So(len(out.AllParentsByTypeName("test1")), convey.ShouldEqual, 0)
		convey.So(out.NearestParentByTypeName("test1"), convey.ShouldBeNil)
	})

	convey.Convey("A schema object with no parents, one child", t, func() {
		child := &TestSchema1{}
		out := &ObjectBase{
			parent:   nil,
			children: []ObjectInterface{child},
		}

		convey.So(out.Root(), convey.ShouldEqual, out)
		convey.So(len(out.Children()), convey.ShouldEqual, 1)
		convey.So(out.Children()[0], convey.ShouldEqual, child)
		convey.So(len(out.ChildrenByTypeName("test1")), convey.ShouldEqual, 1)
		convey.So(out.ChildrenByTypeName("test1")[0], convey.ShouldEqual, child)
		convey.So(len(out.ChildrenByTypeName("not-test")), convey.ShouldEqual, 0)

		convey.So(out.Parent(), convey.ShouldBeNil)
		convey.So(len(out.AllParents()), convey.ShouldEqual, 0)
		convey.So(len(out.AllParentsByTypeName("test1")), convey.ShouldEqual, 0)
		convey.So(out.NearestParentByTypeName("test1"), convey.ShouldBeNil)
	})

	convey.Convey("A schema object with no parents, two children", t, func() {
		child1 := &TestSchema1{}
		child2 := &TestSchema2{}
		out := &ObjectBase{
			parent:   nil,
			children: []ObjectInterface{child1, child2},
		}

		convey.So(out.Root(), convey.ShouldEqual, out)
		convey.So(len(out.Children()), convey.ShouldEqual, 2)
		convey.So(out.Children()[0], convey.ShouldEqual, child1)
		convey.So(out.Children()[1], convey.ShouldEqual, child2)
		convey.So(len(out.ChildrenByTypeName("test1")), convey.ShouldEqual, 1)
		convey.So(out.ChildrenByTypeName("test1")[0], convey.ShouldEqual, child1)
		convey.So(len(out.ChildrenByTypeName("test2")), convey.ShouldEqual, 1)
		convey.So(out.ChildrenByTypeName("test2")[0], convey.ShouldEqual, child2)
		convey.So(len(out.ChildrenByTypeName("not-test")), convey.ShouldEqual, 0)

		convey.So(out.Parent(), convey.ShouldBeNil)
		convey.So(len(out.AllParents()), convey.ShouldEqual, 0)
		convey.So(len(out.AllParentsByTypeName("test1")), convey.ShouldEqual, 0)
		convey.So(out.NearestParentByTypeName("test1"), convey.ShouldBeNil)
	})

	convey.Convey("A schema object with one parent, no children", t, func() {
		parent := &TestSchema1{}
		out := &ObjectBase{
			parent: parent,
		}

		convey.So(out.Root(), convey.ShouldEqual, parent)
		convey.So(len(out.Children()), convey.ShouldEqual, 0)

		convey.So(out.Parent(), convey.ShouldEqual, parent)
		convey.So(len(out.AllParents()), convey.ShouldEqual, 1)
		convey.So(out.AllParents()[0], convey.ShouldEqual, parent)
		convey.So(len(out.AllParentsByTypeName("test1")), convey.ShouldEqual, 1)
		convey.So(out.AllParentsByTypeName("test1")[0], convey.ShouldEqual, parent)
		convey.So(len(out.AllParentsByTypeName("not-test")), convey.ShouldEqual, 0)
		convey.So(out.NearestParentByTypeName("test1"), convey.ShouldEqual, parent)
	})

	convey.Convey("A schema object with parent and grandparent, no children", t, func() {
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

		convey.So(out.Root(), convey.ShouldEqual, grandparent)
		convey.So(len(out.Children()), convey.ShouldEqual, 0)

		convey.So(out.Parent(), convey.ShouldEqual, parent)
		convey.So(len(out.AllParents()), convey.ShouldEqual, 2)
		convey.So(out.AllParents()[0], convey.ShouldEqual, parent)
		convey.So(out.AllParents()[1], convey.ShouldEqual, grandparent)
		convey.So(len(out.AllParentsByTypeName("test1")), convey.ShouldEqual, 1)
		convey.So(out.AllParentsByTypeName("test1")[0], convey.ShouldEqual, parent)
		convey.So(len(out.AllParentsByTypeName("test2")), convey.ShouldEqual, 1)
		convey.So(out.AllParentsByTypeName("test2")[0], convey.ShouldEqual, grandparent)
		convey.So(len(out.AllParentsByTypeName("not-test")), convey.ShouldEqual, 0)
		convey.So(out.NearestParentByTypeName("test1"), convey.ShouldEqual, parent)
		convey.So(out.NearestParentByTypeName("test2"), convey.ShouldEqual, grandparent)
	})
}
