package pkg

import (
	"github.com/pdfcpu/pdfcpu/pkg/api"
)

func GetInfoFromPDFFile(fileName string) ([]string, error) {
	if fileName == "" {
		return nil, ErrEmptyFileName
	}

	return api.InfoFile(fileName, nil, nil)
}
