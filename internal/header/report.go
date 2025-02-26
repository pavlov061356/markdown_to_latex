package header

import "fmt"

const headerType = "report"

// Пример
var (
	_ = Report{
		fontSize:  14,
		paperSize: "a4",
	}
)

type Report struct {
	fontSize  int
	paperSize string
}

func (r *Report) Generate() string {
	return fmt.Sprintf("\\documentclass[%vpt,%vpaper,%v]{ncc}", r.fontSize, r.paperSize, headerType)
}
