package def

import (
	"fmt"

	v1 "github.com/puppetlabs/nebula-sdk/pkg/container/types/v1"
)

type UnknownFileSourceError struct {
	Got v1.FileSource
}

func (e *UnknownFileSourceError) Error() string {
	return fmt.Sprintf("def: unknown file source: %q", e.Got)
}

type UnknownSDKVersionError struct {
	Got string
}

func (e *UnknownSDKVersionError) Error() string {
	return fmt.Sprintf("def: unknown SDK version %q", e.Got)
}

type MissingSettingValueError struct {
	Name string
}

func (e *MissingSettingValueError) Error() string {
	return fmt.Sprintf("def: setting %q has no value", e.Name)
}

type TemplateError struct {
	FileRef *FileRef
	Cause   error
}

func (e *TemplateError) Error() string {
	return fmt.Sprintf("def: error in parent template %q: %+v", e.FileRef, e.Cause)
}
