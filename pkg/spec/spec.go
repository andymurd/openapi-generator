package spec

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	multierror "github.com/hashicorp/go-multierror"
	"github.com/pb33f/libopenapi"
	"github.com/pb33f/libopenapi/datamodel"
	v3high "github.com/pb33f/libopenapi/datamodel/high/v3"
	"github.com/pb33f/libopenapi/resolver"
)

type Spec struct {
	// The high-level model from libopenapi
	OpenAPI3Model *libopenapi.DocumentModel[v3high.Document]
}

func LoadSpecFromFile(filePath string) (*Spec, error) {
	// Read the file
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	config := datamodel.DocumentConfiguration{
		AllowFileReferences:   true,
		AllowRemoteReferences: true,
		BasePath:              filepath.Dir(filePath),
	}

	// Parse the content
	document, err := libopenapi.NewDocumentWithConfiguration(bytes, &config)
	if err != nil {
		return nil, err
	}

	// Check that we actually have an OpenAPI spec
	specType := document.GetSpecInfo().SpecType
	if specType != "openapi" {
		return nil, fmt.Errorf("Unsupported Spec type %s - only openapi specs allowed", specType)
	}

	// Use the spec info to discover what version of OpenAPI spec we have
	specVersion := document.GetSpecInfo().Version
	matched, err := regexp.MatchString(`^3\.[0-1](\.[0-9])?$`, specVersion)
	if err != nil {
		return nil, err
	}
	if matched {
		v3Model, errs := document.BuildV3Model()

		var err error
		var errorsFound = false
		for _, parseError := range errs {
			// Check if this is a resolving error.
			if err2, ok := parseError.(*resolver.ResolvingError); ok {
				// Check if there is a circular reference attached.
				if err2.CircularReference != nil {
					// Output a warning but don't propagate the error to the caller
					fmt.Printf("Circular reference found: %s\n", err2.CircularReference.GenerateJourneyPath())
					continue
				}
			}

			// Add the error to the response
			err = multierror.Append(err, parseError)
			errorsFound = true
		}

		// Any errors in the spec?
		if errorsFound {
			return nil, err
		}

		// Build the model
		retval := Spec{
			OpenAPI3Model: v3Model,
		}
		return &retval, nil
	} else {
		return nil, fmt.Errorf("Unsupported OpenAPI Spec version %s", specVersion)
	}
}
