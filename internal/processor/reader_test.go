package processor

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const layoutISO = "2006-01-02"

func TestReader_Read(t *testing.T) {

	csvTemplate := `first_name,last_name,email,description,merchant_code,amount,from_currency,to_currency,rate,date
%s
`
	type fields struct {
		csvData string
		endDate string
	}
	tests := []struct {
		name       string
		fields     fields
		wantRecord []string
		wantErr    bool
	}{
		{
			name: "happy",
			fields: fields{
				csvData: "Niyah,Singleton,niyah.singleton@mailinator.com,CARD SPEND,5462,682.28,GBP,GGM,1,12/05/2020 08:22",
				endDate: "2020-06-01",
			},
			wantRecord: []string{"Niyah", "Singleton", "niyah.singleton@mailinator.com", "CARD SPEND", "5462", "682.28", "GBP", "GGM", "1", "12/05/2020 08:22"},
		},
		{
			name: "filter_too_old",
			fields: fields{
				csvData: "Niyah,Singleton,niyah.singleton@mailinator.com,CARD SPEND,5462,682.28,GBP,GBP,1,12/05/2020 08:22",
				endDate: "2020-11-13",
			},
			wantRecord: nil,
		},
		{
			name: "filter_too_new",
			fields: fields{
				csvData: "Niyah,Singleton,niyah.singleton@mailinator.com,CARD SPEND,5462,682.28,GBP,GBP,1,12/05/2020 08:22",
				endDate: "2020-05-11",
			},
			wantRecord: nil,
		},
		{
			name: "not_card_spend",
			fields: fields{
				csvData: "Niyah,Singleton,niyah.singleton@mailinator.com,SELL GOLD,5462,682.28,GBP,GBP,1,12/05/2020 08:22",
				endDate: "2020-06-01",
			},
			wantRecord: nil,
		},
		{
			name: "not_gold",
			fields: fields{
				csvData: "Niyah,Singleton,niyah.singleton@mailinator.com,CARD SPEND,5462,682.28,GBP,GBP,1,12/05/2020 08:22",
				endDate: "2020-06-01",
			},
			wantRecord: nil,
		},
		{
			name: "error",
			fields: fields{
				endDate: "2020-06-01",
			},
			wantRecord: nil,
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			endDate, err := time.Parse(layoutISO, tt.fields.endDate)
			require.NoError(t, err)
			sr := strings.NewReader(fmt.Sprintf(csvTemplate, tt.fields.csvData))
			r := NewReader(sr, endDate)
			gotRecord, err := r.Read()
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			assert.Len(t, gotRecord, len(tt.wantRecord))
			assert.Equal(t, tt.wantRecord, gotRecord)
		})
	}
}

func TestReader_Read_once(t *testing.T) {
	csv := `first_name,last_name,email,description,merchant_code,amount,from_currency,to_currency,rate,date
Niyah,Singleton,niyah.singleton@mailinator.com,CARD SPEND,5462,682.28,GBP,GGM,1,12/05/2020 08:22
Amanda,Burn,amanda.burn@mailinator.com,CARD SPEND,5013,2906.19,GBP,GGM,47.0892,12/05/2020 03:07
Andreea,Suarez,andreea.suarez@mailinator.com,CARD SPEND,5411,2318.96,GBP,GGM,1,12/05/2020 09:20
Alayna,Sparks,alayna.sparks@mailinator.com,CARD SPEND,5311,2629.16,GBP,GGM,47.0892,12/05/2020 13:28
`
	endDate, err := time.Parse(layoutISO, "2020-06-01")
	require.NoError(t, err)
	r := NewReader(strings.NewReader(csv), endDate)

	assertRead(t, r, "Niyah")
	assertRead(t, r, "Amanda")
	assertRead(t, r, "Andreea")
	assertRead(t, r, "Alayna")
	read, err := r.Read()
	assert.Empty(t, read)
	assert.Error(t, err)
}

func assertRead(t *testing.T, r Reader, expected string) {
	read, err := r.Read()
	assert.NoError(t, err)
	assert.NotEmpty(t, read)
	assert.Equal(t, expected, read[0])
}
