package main

import (
	"log"
	"os"
	"path/filepath"
	"text/template"

	"github.com/pkg/errors"
)

// ErrorDefinition describes an error that will be generated
type ErrorDefinition struct {
	// Name of the error, generate "{{.Name}}Error"
	Name string
	// Comment of creator functions, generate "is used {{.UseCase}}"
	UseCase string
	// HTTP code returned by error handlers, generate "http.{{.HTTPCode}}"
	HTTPCode string
	// Output filename
	Filename string
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err.Error())
	}
}

func run() error {
	basePath, err := getBasePath()
	if err != nil {
		return err
	}

	tpl, err := template.ParseFiles(basePath + "/generator/error_template.tmpl")
	if err != nil {
		return errors.Wrap(err, "could not parse template file")
	}

	testTpl, err := template.ParseFiles(basePath + "/generator/test_template.tmpl")
	if err != nil {
		return errors.Wrap(err, "could not parse template test file")
	}

	defs := []ErrorDefinition{
		{
			Name:     "ForbiddenResource",
			UseCase:  "when the resource cannot be accessed by authenticated user",
			HTTPCode: "StatusForbidden",
			Filename: "forbidden_resource_error.go",
		},
		{
			Name:     "InvalidArgument",
			UseCase:  "when the provided argument is incorrect",
			HTTPCode: "StatusUnprocessableEntity",
			Filename: "invalid_argument_error.go",
		},
		{
			Name:     "InvalidData",
			UseCase:  "when an invalid data is detected during an action",
			HTTPCode: "StatusUnprocessableEntity",
			Filename: "invalid_data_error.go",
		},
		{
			Name:     "NotFound",
			UseCase:  "when we cannot find a specified resource",
			HTTPCode: "StatusNotFound",
			Filename: "not_found_error.go",
		},
		{
			Name:     "OutdatedResource",
			UseCase:  "when the state of a resource does not allow to apply a specific action",
			HTTPCode: "StatusConflict",
			Filename: "outdated_resource_error.go",
		},
	}

	for _, def := range defs {
		// error file
		f, err := os.Create(basePath + "/" + def.Filename)
		if err != nil {
			return errors.Wrapf(err, "could not create file %s", def.Filename)
		}

		if err := tpl.Execute(f, def); err != nil {
			return errors.Wrapf(err, "could not create execute template for %s", def.Filename)
		}

		if err := f.Close(); err != nil {
			return errors.Wrapf(err, "could not close file %s", def.Filename)
		}

		// test file
		testFilename := def.Filename[:len(def.Filename)-3] + "_test.go"
		f, err = os.Create(basePath + "/" + testFilename)
		if err != nil {
			return errors.Wrapf(err, "could not create test file %s", testFilename)
		}

		if err := testTpl.Execute(f, def); err != nil {
			return errors.Wrapf(err, "could not create execute template for %s", testFilename)
		}

		if err := f.Close(); err != nil {
			return errors.Wrapf(err, "could not close file %s", testFilename)
		}
	}

	return nil
}

func getBasePath() (string, error) {
	basePath, err := os.Getwd()
	if err != nil {
		return "", errors.Wrap(err, "error retrieving working directory")
	}

	if filepath.Base(basePath) == "generator" {
		basePath = filepath.Dir(basePath)
	}

	if filepath.Base(basePath) != "errorz" {
		return "", errors.Wrap(err, "script needs to be run from errorz directory")
	}

	return basePath, nil
}
