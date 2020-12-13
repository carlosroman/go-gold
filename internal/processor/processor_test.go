package processor

import (
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func testStores() []stubStore {
	return []stubStore{{}, {}, {}, {}, {}, {}}
}

func TestProcessor_Run(t *testing.T) {
	dir, err := ioutil.TempDir("", "")
	require.NoError(t, err)
	defer os.RemoveAll(dir)

	testStores := testStores()
	p := New([]Store{
		&testStores[0],
		&testStores[1],
		&testStores[2],
		&testStores[3],
		&testStores[4],
		&testStores[5],
	})
	finalCSV := path.Join(dir, "test.csv")
	endDate, err := time.Parse(layoutISO, "2020-05-20")
	require.NoError(t, err)
	sr := &stubReader{
		records: [][]string{
			getRecord("one", "12/05/2020 08:22"),
			getRecord("two", "12/04/2020 08:22"),
			getRecord("three", "12/03/2020 08:22"),
			getRecord("four", "12/02/2020 08:22"),
			getRecord("five", "12/01/2020 08:22"),
			getRecord("six", "12/12/2019 08:22"),
		},
		endDate: endDate,
		err:     io.EOF,
	}
	for i := range testStores {
		testStores[i].On("Save", mock.Anything)
		testStores[i].On("Flush").Return(strings.NewReader("flush\n"))
	}
	err = p.Run(sr, finalCSV)
	require.NoError(t, err)
	_, err = os.Stat(finalCSV)
	assert.NoError(t, err)
	for i := range testStores {
		testStores[i].AssertCalled(t, "Save", sr.records[i])
	}
	content, err := ioutil.ReadFile(finalCSV)
	assert.NoError(t, err)
	assert.Equal(t, "date,first_name,last_name,total\nflush\nflush\nflush\nflush\nflush\nflush\n", string(content))
}

func getRecord(name, date string) []string {
	return []string{name, "Test", "name.test@mailinator.com", "CARD SPEND", "5462", "682.28", "GBP", "GGM", "1", date}
}

type stubReader struct {
	currentRecord int
	records       [][]string
	err           error
	startDate     time.Time
	endDate       time.Time
}

func (r *stubReader) Read() (record []string, err error) {
	if r.currentRecord < len(r.records) {
		defer func() {
			r.currentRecord++
		}()
		return r.records[r.currentRecord], nil
	}
	return nil, r.err
}

func (r stubReader) GetEndDate() time.Time {
	return r.endDate
}

func (r stubReader) GetStartDate() time.Time {
	return r.startDate
}

type stubStore struct {
	mock.Mock
}

func (ss *stubStore) Save(read []string) {
	ss.Called(read)
}

func (ss *stubStore) Flush() io.Reader {
	return ss.Called().Get(0).(io.Reader)
}
