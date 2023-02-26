package ports

type ProgressReporter interface {
	ReportProgress(float32, string)
}
