package processor

import (
	"encoding/csv"
	"io"
	"sync"
	"time"
)

const (
	firstNameColumn   = 0
	lastNameColumn    = 1
	amountColumn      = 5
	ratetColumn       = 8
	goldCurrency      = "GGM"
	toCurrencyColumn  = 7
	cardSpend         = "CARD SPEND"
	descriptionColumn = 3
	dateFormat        = "02/01/2006 15:04"
	dateColumn        = 9
)

type Reader interface {
	Read() (record []string, err error)
	GetEndDate() time.Time
	GetStartDate() time.Time
}

func NewReader(r io.Reader, endDate time.Time) Reader {
	return &reader{
		r:         csv.NewReader(r),
		startDate: endDate.AddDate(0, -6, 0),
		endDate:   endDate,
		once:      &sync.Once{},
	}
}

type reader struct {
	r         *csv.Reader
	startDate time.Time
	endDate   time.Time
	once      *sync.Once
}

func (r reader) GetEndDate() time.Time {
	return r.endDate
}

func (r reader) GetStartDate() time.Time {
	return r.startDate
}

func (r reader) Read() (record []string, err error) {
	r.once.Do(func() {
		_, _ = r.r.Read()
	})
	record, err = r.r.Read()
	if err != nil {
		return nil, err
	}

	if record[descriptionColumn] != cardSpend {
		return nil, nil
	}

	if record[toCurrencyColumn] != goldCurrency {
		return nil, nil
	}

	t, _ := time.Parse(dateFormat, record[dateColumn])
	if t.After(r.endDate) {
		return nil, nil
	}
	if t.Before(r.startDate) {
		return nil, nil
	}
	return
}
