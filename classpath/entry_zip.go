package classpath

import (
	"archive/zip"
	"errors"
	"io"
	"io/ioutil"
)

type zipEntry struct {
	absPath string
}

func newZipEntry(path string) *zipEntry {
	return &zipEntry{
		absPath: path,
	}
}

func (z *zipEntry) readClass(className string) ([]byte, Entry, error) {
	reader, err := zip.OpenReader(z.absPath)
	if err != nil {
		panic(err)
	}
	defer reader.Close()

	var rc io.ReadCloser
	defer rc.Close()

	for _, file := range reader.File {
		if file.Name == className {
			rc, err = file.Open()
			if err != nil {
				return nil, nil, err
			}

			data, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil, nil, err
			}

			return data, z, nil
		}
	}
	return nil, nil, errors.New("class not found: " + className)
}

func (z *zipEntry) String() string {
	return z.absPath
}
