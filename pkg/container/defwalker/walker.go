package defwalker

import (
	"os"
	"path/filepath"

	"github.com/puppetlabs/nebula-sdk/pkg/container/def"
	v1 "github.com/puppetlabs/nebula-sdk/pkg/container/types/v1"
)

func Walk(root string) ([]*def.ResolvedContainer, error) {
	var cs []*def.ResolvedContainer
	var errs WalkErrors

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() || info.Name() != def.DefaultFilename {
			return nil
		}

		c, err := def.NewFromFilePath(path)
		if _, ok := err.(*v1.InvalidVersionKindError); ok {
			// This file is named container.yaml but does not contain a
			// container specification.
			return nil
		} else if err != nil {
			errs = append(errs, &WalkError{
				Path:  path,
				Cause: err,
			})
			return nil
		}

		cs = append(cs, c)
		return nil
	})
	if err != nil {
		return nil, err
	} else if len(errs) > 0 {
		return nil, errs
	}

	return cs, nil
}
