package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_processFile(t *testing.T) {
	type args struct {
		csvFile   string
		outputDir string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "happy",
			args: args{
				csvFile:   "csvFile",
				outputDir: "outputDir",
			},
		},
		{
			name: "sad",
			args: args{
				csvFile:   "csvFile",
				outputDir: "outputDir",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &mockProcessor{
				wantErr: tt.wantErr,
			}
			err := processFile(p, tt.args.csvFile, tt.args.outputDir)
			assert.True(t, p.called)
			assert.Equal(t, tt.args.csvFile, p.csvFile)
			assert.Equal(t, tt.args.outputDir, p.outputDir)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}

type mockProcessor struct {
	called    bool
	wantErr   bool
	csvFile   string
	outputDir string
}

func (m *mockProcessor) Run(csvFile, outputDir string) (err error) {
	m.called = true
	m.csvFile = csvFile
	m.outputDir = outputDir
	if m.wantErr {
		err = errors.New("something went wrong")
	}
	return err
}
