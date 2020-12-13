package processor

import (
	"io"
	"os"
	"time"
)

type Processor interface {
	Run(csvFile Reader, outputDir string) error
}

func New(stores []Store) Processor {
	return &processor{
		stores: stores,
	}
}

type processor struct {
	stores []Store
}

func (p processor) Run(csvFile Reader, outputFile string) (err error) {
	outCSV, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	stores := make(map[time.Month]Store)
	for i := 0; i < 6; i++ {
		key := csvFile.GetEndDate().AddDate(0, -i, 0)
		stores[key.Month()] = p.stores[i]
	}
	defer func() { _ = outCSV.Close() }()
	for {
		record, err := csvFile.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		t, _ := time.Parse(dateFormat, record[dateColumn])
		stores[t.Month()].Save(record)
	}
	_, err = outCSV.WriteString("date,first_name,last_name,total\n")
	if err != nil {
		return err
	}
	for i := 0; i < 6; i++ {
		_, err = io.Copy(outCSV, p.stores[i].Flush())
		if err != nil {
			return err
		}
	}
	return err
}
