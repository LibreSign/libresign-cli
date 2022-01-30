package pkg

import (
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
)

type PageInfo struct {
	Rot         int     `json:"rotation"`
	Orientation string  `json:"orientation"`
	Height      float64 `json:"height"`
	Width       float64 `json:"width"`
}

type PDFInfo struct {
	Pages []PageInfo `json:"pages"`
}

func GetInfoFromPDFFile(fileName string) (info PDFInfo, err error) {
	if fileName == "" {
		return info, ErrEmptyFileName
	}

	ctx, err := api.ReadContextFile(fileName)

	if err != nil {
		return info, err
	}

	err = ctx.EnsurePageCount()

	if err != nil {
		return info, err
	}

	pages, err := ctx.PageBoundaries()

	list := make([]PageInfo, len(pages))

	for index, page := range pages {
		list[index] = extractPageInfo(page)
	}

	info.Pages = list

	return info, err
}

func extractPageInfo(pb pdfcpu.PageBoundaries) PageInfo {
	dt := pb.CropBox().Dimensions()

	or := "portrait"
	if dt.Landscape() {
		or = "landscape"
	}

	return PageInfo{
		Rot:         pb.Rot,
		Width:       dt.Width,
		Height:      dt.Height,
		Orientation: or,
	}
}
