package storage_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/snaztoz/granary/internal/storage"
)

func TestReadWriteStorage(t *testing.T) {
	password := "my-password"
	data := map[string]string{"foo": "bar"}

	st, err := storage.Init(new(bytes.Buffer), password)
	if err != nil {
		t.Fatal(err)
	}

	if err := st.WriteData(data); err != nil {
		t.Fatal(err)
	}

	retrievedData, err := st.ReadData()
	if err != nil {
		t.Fatal(err)
	}

	for k, v := range data {
		rv, exist := retrievedData[k]
		if !exist {
			t.Fatalf("key of value '%s' not exist in retrieved data", k)
		} else if rv != v {
			t.Fatalf("expecting to get value '%s', but received '%s' instead", v, rv)
		}
	}
}

func TestPersistingData(t *testing.T) {
	password := "my-password"
	data := map[string]string{"foo": "bar"}

	st, err := storage.Init(new(bytes.Buffer), password)
	if err != nil {
		t.Fatal(err)
	}

	if err := st.WriteData(data); err != nil {
		t.Fatal(err)
	}

	persistent := new(bytes.Buffer)
	if err := st.Persist(persistent); err != nil {
		t.Fatal(err)
	}

	if !strings.HasPrefix(persistent.String(), storage.SecretFileHeader) {
		t.Fatal("corrupted result")
	}
}
