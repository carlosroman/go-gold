package processor

type Processor interface {
	Run(csvFile, outputDir string) error
}

func New() Processor {
	return &processor{}
}

type processor struct {
}

func (p processor) Run(csvFile, outputDir string) (err error) {
	return err
}
