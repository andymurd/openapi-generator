package spec

import (
	"testing"

	"github.com/andymurd/openapi-generator/pkg/schema"
	"github.com/andymurd/openapi-generator/pkg/schema/v3"

	"github.com/smartystreets/goconvey/convey"
)

func TestSpec(t *testing.T) {
	convey.Convey("Load the petstore spec", t, func() {
		result, err := LoadSpecFromFile("../../test/openapi-specs/petstore-v3.yml")
		convey.So(err, convey.ShouldBeNil)
		convey.So(result, convey.ShouldNotBeNil)

		doc := v3.Document{Document: result.OpenAPI3Model.Model, ObjectBase: schema.ObjectBase{}}
		doc.Wrap(nil)
	})

	convey.Convey("Load the stripe spec", t, func() {
		result, err := LoadSpecFromFile("../../test/openapi-specs/stripe-v3.json")
		convey.So(err, convey.ShouldBeNil)
		convey.So(result, convey.ShouldNotBeNil)

		doc := v3.Document{Document: result.OpenAPI3Model.Model, ObjectBase: schema.ObjectBase{}}
		doc.Wrap(nil)
	})

	convey.Convey("Fail to load a file that is not JSON or YAML", t, func() {
		_, err := LoadSpecFromFile("../../test/openapi-specs/not-good.jpg")
		convey.So(err, convey.ShouldNotBeNil)
	})

	convey.Convey("Fail to load a JSON file that is not an OpenAPI spec", t, func() {
		_, err := LoadSpecFromFile("../../test/openapi-specs/not-spec.json")
		convey.So(err, convey.ShouldNotBeNil)
	})

	convey.Convey("Fail to load a YAML file that is not an OpenAPI spec", t, func() {
		_, err := LoadSpecFromFile("../../test/openapi-specs/not-spec.yaml")
		convey.So(err, convey.ShouldNotBeNil)
	})
}
