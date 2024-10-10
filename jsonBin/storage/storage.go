package storage

import (
	"encoding/json"
	"fmt"
	"jsonBin/bins"
	"os"
)

type StorageService interface {
	SaveBinListFile(filePath string, binList *bins.BinList) error
	LoadBinListBytes(data []byte) (*bins.BinList, error)
	LoadBinListFile(filePath string) (*bins.BinList, error)
}

type FileStorage struct {
}

func (s *FileStorage) SaveBinListFile(filePath string, binList *bins.BinList) error {
	data, err := json.Marshal(binList)
	if err != nil {
		return fmt.Errorf("error for marhsal JSON: %v", err)
	}

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error in the creating file: %v", err)
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("error in the writing file: %v", err)
	}
	return nil
}

func (s *FileStorage) LoadBinListBytes(data []byte) (*bins.BinList, error) {
	var binList bins.BinList
	err := json.Unmarshal(data, &binList)
	if err != nil {
		return nil, fmt.Errorf("error in the unmarshal JSON: %v", err)
	}
	return &binList, nil
}

func (s *FileStorage) LoadBinListFile(filePath string) (*bins.BinList, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error in the reading file: %v", err)
	}
	return s.LoadBinListBytes(data)
}
