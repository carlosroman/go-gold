package processor

import (
	"bytes"
	"fmt"
	"io"
	"sort"
	"strconv"
	"sync"
	"time"
)

type Store interface {
	Save([]string)
	Flush() io.Reader
}

func NewStores() []Store {
	return []Store{
		&store{
			m:    make(map[string]float64, 100),
			once: &sync.Once{},
		},
		&store{
			m:    make(map[string]float64, 100),
			once: &sync.Once{},
		},
		&store{
			m:    make(map[string]float64, 100),
			once: &sync.Once{},
		},
		&store{
			m:    make(map[string]float64, 100),
			once: &sync.Once{},
		},
		&store{
			m:    make(map[string]float64, 100),
			once: &sync.Once{},
		},
		&store{
			m:    make(map[string]float64, 100),
			once: &sync.Once{},
		},
		&store{
			m:    make(map[string]float64, 100),
			once: &sync.Once{},
		}}
}

type store struct {
	m    map[string]float64
	once *sync.Once
	date string
}

func (s *store) Save(record []string) {
	key := fmt.Sprintf("%v,%v", record[firstNameColumn], record[lastNameColumn])
	s.once.Do(func() {
		t, _ := time.Parse(dateFormat, record[dateColumn])
		s.date = t.Format("Jan 2006")
	})
	_, ok := s.m[key]
	if !ok {
		s.m[key] = 0
	}
	amount, err := strconv.ParseFloat(record[amountColumn], 8)
	if err != nil {
		return
	}
	rate, err := strconv.ParseFloat(record[ratetColumn], 8)
	if err != nil {
		return
	}
	s.m[key] = s.m[key] + (amount / rate)
}

func (s store) Flush() io.Reader {
	b := new(bytes.Buffer)
	l := make([]struct {
		text string
		val  float64
	}, len(s.m))
	idx := 0
	for k, v := range s.m {
		l[idx] = struct {
			text string
			val  float64
		}{text: k, val: v}
		idx++
	}
	sort.SliceStable(l, func(i, j int) bool {
		return l[i].val > l[j].val
	})

	if len(l) > 0 {
		b.WriteString(fmt.Sprintf("%s,%s,%.2f\n", s.date, l[0].text, l[0].val))
	}
	if len(l) > 1 {
		b.WriteString(fmt.Sprintf("%s,%s,%.2f\n", s.date, l[1].text, l[1].val))
	}
	if len(l) > 2 {
		b.WriteString(fmt.Sprintf("%s,%s,%.2f\n", s.date, l[2].text, l[2].val))
	}
	return b
}
