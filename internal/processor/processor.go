package processor

type Processor interface {
	Run(csvFile Reader, outputDir string) error
}

func New() Processor {
	return &processor{}
}

type processor struct {
}

func (p processor) Run(csvFile Reader, outputDir string) (err error) {
	return err
}
