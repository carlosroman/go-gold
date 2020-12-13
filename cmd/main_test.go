package main

import (
	"errors"
	"io"
	"testing"
	"time"

	"github.com/carlosroman/go-gold/internal/processor"
	"github.com/stretchr/testify/assert"
)

func Test_processFile(t *testing.T) {
	now := time.Now()
	type args struct {
		csvFile   io.Reader
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
				csvFile:   &stubReader{},
				outputDir: "outputDir",
			},
		},
		{
			name: "sad",
			args: args{
				csvFile:   &stubReader{},
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
			err := processFile(p, tt.args.csvFile, tt.args.outputDir, now)
			assert.True(t, p.called)
			assert.Equal(t, processor.NewReader(tt.args.csvFile, now), p.csvFile)
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
	csvFile   processor.Reader
	outputDir string
}

func (m *mockProcessor) Run(csvFile processor.Reader, outputDir string) (err error) {
	m.called = true
	m.csvFile = csvFile
	m.outputDir = outputDir
	if m.wantErr {
		err = errors.New("something went wrong")
	}
	return err
}

type stubReader struct {
}

func (s stubReader) Read(p []byte) (n int, err error) {
	return
}
