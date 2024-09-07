package storage

import (
	"encoding/json"
	"io"

	"github.com/snaztoz/granary/internal/crypto"
	"github.com/snaztoz/granary/internal/data"
)

func Init(out io.Writer, passphrase string) (storage *Storage, err error) {
	key, keyString := crypto.DeriveKey(passphrase)
	jsonData, _ := json.Marshal(make(data.T))

	ciphertext, err := crypto.Encrypt(jsonData, key)
	if err != nil {
		return nil, err
	}

	content := []byte(toFileContent(keyString, ciphertext))
	if _, err := out.Write(content); err != nil {
		return nil, err
	}

	return &Storage{
		content:   content,
		keyString: keyString,
		key:       key,
	}, nil
}

func Open(in io.Reader, passphrase string) (storage *Storage, err error) {
	content, err := io.ReadAll(in)
	if err != nil {
		return nil, err
	}

	keyString, _, err := toKeyStringAndData(string(content))
	if err != nil {
		return nil, err
	}

	key, err := crypto.MatchPassphrase(passphrase, keyString)
	if err != nil {
		return nil, err
	}

	return &Storage{
		content:   content,
		keyString: keyString,
		key:       key,
	}, nil
}

type Storage struct {
	content   []byte
	keyString string
	key       []byte
}

func (s *Storage) ReadData() (data data.T, err error) {
	_, ciphertext, _ := toKeyStringAndData(string(s.content))

	plaintext, err := crypto.Decrypt(ciphertext, s.key)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(plaintext, &data); err != nil {
		return nil, err
	}

	return data, nil
}

func (s *Storage) WriteData(data data.T) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	ciphertext, err := crypto.Encrypt(jsonData, s.key)
	if err != nil {
		return err
	}

	s.content = []byte(toFileContent(s.keyString, ciphertext))

	return nil
}

func (s *Storage) Persist(out io.Writer) error {
	if _, err := out.Write(s.content); err != nil {
		return err
	}
	return nil
}
