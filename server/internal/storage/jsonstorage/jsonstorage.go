package jsonstorage

import (
	"browser-remote-server/internal/storage"
	"browser-remote-server/lib/e"
	"encoding/json"
	"os"
)

type JsonStorage struct {
	Path string
}

type StorageMap = map[string][]string

const defaultPerm = 0774

func New(path string) *JsonStorage {
	return &JsonStorage{
		Path: path,
	}
}

func (s *JsonStorage) Save(host string, bindings []string) (err error) {
	defer func() { e.WrapIfErr("jsonstorage.Save", err) }()

	storageMap, err := readStorageMap(s.Path)

	if err != nil {
		return err
	}

	storageMap[host] = bindings

	data, err := json.Marshal(storageMap)

	if err != nil {
		return err
	}

	file, err := os.Create(s.Path)

	if err != nil {
		return err
	}

	if _, err = file.Write(data); err != nil {
		return err
	}

	err = file.Close()

	return err
}

func (s *JsonStorage) SaveElement(elementInfo storage.ElementInfo) (id int, err error) {
	defer func() { e.WrapIfErr("jsonstorage.Read", err) }()

	bindings, err := s.Read(elementInfo.Host)

	if err != nil {
		return -1, err
	}

	bindings = append(bindings, elementInfo.ElementQuery)

	err = s.Save(elementInfo.Host, bindings)
}

func (s *JsonStorage) Read(host string) (bindings []string, err error) {
	defer func() { e.WrapIfErr("jsonstorage.Read", err) }()

	storageMap, err := readStorageMap(s.Path)

	if err != nil {
		return nil, err
	}

	bidings, ok := storageMap[host]

	if ok {
		return bidings, nil
	}

	return nil, nil
}

func readStorageMap(filePath string) (sotargeMap StorageMap, err error) {
	defer func() { e.WrapIfErr("jsonstorage.readStorageMap", err) }()

	if err := os.MkdirAll(filePath, defaultPerm); err != nil {
		return nil, err
	}

	exists := true

	_, err = os.Stat(filePath)

	if err != nil {
		exists = false

		if !os.IsNotExist(err) {
			return nil, err
		}
	}

	if exists {
		var storageMap StorageMap

		data, err := os.ReadFile(filePath)

		if err != nil {
			return nil, err
		}

		if err = json.Unmarshal(data, &storageMap); err != nil {
			return nil, err
		}

		return storageMap, nil

	} else {
		return make(StorageMap), nil
	}
}
