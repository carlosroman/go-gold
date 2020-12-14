package processor_test

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/carlosroman/go-gold/internal/processor"
)

func TestNewStores(t *testing.T) {
	assert.Len(t, processor.NewStores(), 6)
}

func TestStore(t *testing.T) {
	store := processor.NewStores()[0]
	require.NotNil(t, store)
	store.Save(getRecord("Bob", "2971.96", "47.7555", "12/05/2020 08:22"))
	store.Save(getRecord("Bob", "1485.98", "47.7555", "12/05/2020 08:22"))
	content, err := ioutil.ReadAll(store.Flush())
	require.NoError(t, err)
	assert.Equal(t, "May 2020,Bob,Test,93.35\n", string(content))
}

func TestStore_multiple(t *testing.T) {
	store := processor.NewStores()[0]
	require.NotNil(t, store)
	store.Save(getRecord("Dave", "85.98", "47.7555", "12/05/2020 08:22"))
	store.Save(getRecord("Bob", "1485.98", "47.7555", "12/05/2020 08:22"))
	store.Save(getRecord("Alice", "2971.96", "47.7555", "12/05/2020 08:22"))
	store.Save(getRecord("John", "1.96", "47.7555", "12/05/2020 08:22"))
	content, err := ioutil.ReadAll(store.Flush())
	require.NoError(t, err)
	assert.Equal(t, "May 2020,Alice,Test,62.23\nMay 2020,Bob,Test,31.12\nMay 2020,Dave,Test,1.80\n", string(content))
}

func getRecord(name, amount, rate, date string) []string {
	return []string{name, "Test", name + "@mailinator.com", "CARD SPEND", "5462", amount, "GBP", "GGM", rate, date}
}
