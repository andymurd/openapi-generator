package spec

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSpec(t *testing.T) {
	Convey("Load the petstore spec", t, func() {
		result, err := LoadSpecFromFile("../../test/openapi-specs/petstore-v3.yml")
		So(err, ShouldBeNil)
		So(result, ShouldNotBeNil)
	})

	Convey("Load the stripe spec", t, func() {
		result, err := LoadSpecFromFile("../../test/openapi-specs/stripe-v3.json")
		So(err, ShouldBeNil)
		So(result, ShouldNotBeNil)
	})

	Convey("Fail to load a file that is not JSON or YAML", t, func() {
		_, err := LoadSpecFromFile("../../test/openapi-specs/not-good.jpg")
		So(err, ShouldNotBeNil)
	})

	Convey("Fail to load a JSON file that is not an OpenAPI spec", t, func() {
		_, err := LoadSpecFromFile("../../test/openapi-specs/not-spec.json")
		So(err, ShouldNotBeNil)
	})

	Convey("Fail to load a YAML file that is not an OpenAPI spec", t, func() {
		_, err := LoadSpecFromFile("../../test/openapi-specs/not-spec.yaml")
		So(err, ShouldNotBeNil)
	})
}
