package pkg

import (
	"errors"

	"github.com/pdfcpu/pdfcpu/pkg/api"
)

var ErrEmptyFileName = errors.New("EMPTY_FILE_NAME")

func ValidatePDFFile(fileName string) error {
	if fileName == "" {
		return ErrEmptyFileName
	}

	return api.ValidateFile(fileName, nil)
}
