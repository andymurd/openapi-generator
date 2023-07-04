package template

import (
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/andymurd/openapi-generator/pkg/schema"
)

type InputOutputFilePair struct {
	// The template filename (without the path) that this object can use to
	// create some output.
	TemplateFileName string

	// If a TemplateFileName() is available, this can be provided to specify
	// the output filename. If no value is returned, no separate file will be
	// output.
	OutputFileNameFn func(schema.ObjectInterface) (string, error)
}

type Generator struct {
	// A map of object types to sets of input and output files. It OK for an
	// object type name to not exist in a generator,
	// that file will simply not be output. The content of this object may be
	// used in a different template, usually that of a parent, or a child
	// collection.
	ObjectTypeTemplates map[schema.ObjectTypeName][]InputOutputFilePair
}

// Output file(s) for the supplied schema object and its children
func (this *Generator) Output(schemaObject schema.ObjectInterface, outputDirectory string) error {
	// Lookup templates for the schema object
	ioFilePairs, exists := this.ObjectTypeTemplates[schemaObject.ObjectTypeName()]
	if exists {
		for _, ioFilePair := range ioFilePairs {
			err := this.OutputFile(schemaObject, outputDirectory, ioFilePair)
			if err != nil {
				return err
			}
		}
	}

	// Recurse to the children
	for _, child := range schemaObject.Children() {
		err := this.Output(child, outputDirectory)
		if err != nil {
			return err
		}
	}
	return nil
}

// Output a file for the supplied schema object
func (this *Generator) OutputFile(schemaObject schema.ObjectInterface, outputDirectory string, ioFilePair InputOutputFilePair) error {
	outputFileName, err := ioFilePair.OutputFileNameFn(schemaObject)
	if err != nil {
		return err
	}

	if outputFileName != "" {
		// Open the file
		fullPath := filepath.Join(outputDirectory, outputFileName)
		fh, err := os.Create(fullPath)
		if err != nil {
			return err
		}
		defer fh.Close()

		// Build the content
		content, err := this.ExpandTemplate(schemaObject, ioFilePair.TemplateFileName)
		if err != nil {
			return err
		}

		// Write to the file
		_, err = fh.WriteString(content)
		if err != nil {
			return err
		}
	}

	return nil
}

func (this *Generator) ExpandTemplate(schemaObject schema.ObjectInterface, templateFileName string) (string, error) {
	// Parse the template
	parsedTemplate, err := template.ParseFiles(templateFileName)
	if err != nil {
		return "", err
	}

	// Apply the schema object to the template
	buffer := new(strings.Builder)
	err = parsedTemplate.Execute(buffer, schemaObject)
	if err != nil {
		return "", err
	}

	return buffer.String(), nil
}
