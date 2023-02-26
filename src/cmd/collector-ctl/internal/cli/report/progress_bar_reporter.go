package report

import (
	"github.com/schollz/progressbar/v3"
)

type ProgressBarReporter struct {
	Bar *progressbar.ProgressBar
}

func (p *ProgressBarReporter) Draw() {
	p.Bar.RenderBlank()
}

func (p *ProgressBarReporter) ReportProgress(_ float32, message string) {
	p.Bar.Add(1)
}
