package gloat

import (
	"os"
	"path/filepath"
)

type Source interface {
	Collect() (Migrations, error)
}

type FileSystemSource struct {
	MigrationsFolder string
}

func (s *FileSystemSource) Collect() (Migrations, error) {
	var migrations Migrations

	err := filepath.Walk(s.MigrationsFolder, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			migration, err := FromPath(path)
			if err != nil {
				return err
			}

			migrations = append(migrations, migration)
		}

		return nil
	})

	return migrations, err
}