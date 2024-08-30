package storage

import (
	"encoding/json"
	"os"

	"github.com/snaztoz/granary/internal/crypto"
	"github.com/snaztoz/granary/internal/data"
)

func New(path, password string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	key, keyString := crypto.DeriveKey(password)
	jsonData, _ := json.Marshal(make(data.T))

	ciphertext, err := crypto.Encrypt(jsonData, key)
	if err != nil {
		return err
	}

	content := toFileContent(keyString, ciphertext)
	if _, err := f.Write([]byte(content)); err != nil {
		return err
	}

	return nil
}

func Open(path, password string) (storage *Storage, err error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	keyString, _, err := toKeyStringAndData(string(content))
	if err != nil {
		return nil, err
	}

	key, err := crypto.MatchPassword(password, keyString)
	if err != nil {
		return nil, err
	}

	return &Storage{
		path:      path,
		content:   string(content),
		keyString: keyString,
		key:       key,
	}, nil
}

type Storage struct {
	path      string
	content   string
	keyString string
	key       []byte
}

func (s *Storage) ReadFile() (data data.T, err error) {
	_, ciphertext, _ := toKeyStringAndData(s.content)

	plaintext, err := crypto.Decrypt(ciphertext, s.key)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(plaintext, &data); err != nil {
		return nil, err
	}

	return data, nil
}

func (s *Storage) WriteFile(data data.T) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	ciphertext, err := crypto.Encrypt(jsonData, s.key)
	if err != nil {
		return err
	}

	s.content = toFileContent(s.keyString, ciphertext)

	if err := os.WriteFile(s.path, []byte(s.content), 0644); err != nil {
		return err
	}

	return nil
}
