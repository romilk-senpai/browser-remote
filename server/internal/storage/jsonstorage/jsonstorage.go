package jsonstorage

import (
	"browser-remote-server/internal/storage"
	"encoding/json"
	"errors"
	"os"
	"path"
)

var (
	ErrNoData = errors.New("no data on host")
)

const (
	StorageFilename = "storage.json"
)

type StorageMap = map[string]storage.Host

const defaultPerm = 0774

type JsonStorage struct {
	Path string
}

func New(path string) *JsonStorage {
	return &JsonStorage{
		Path: path,
	}
}

func (s *JsonStorage) Init() error {
	if err := os.MkdirAll(s.Path, os.ModePerm); err != nil {
		return err
	}

	dbfilepath := s.dbfilepath()

	_, err := os.Stat(dbfilepath)

	if err != nil {
		storageMap := make(StorageMap)

		data, err := json.Marshal(storageMap)
		if err != nil {
			return err
		}

		err = os.WriteFile(dbfilepath, data, defaultPerm)

		return err
	}

	return nil
}

func (s *JsonStorage) Save(host storage.Host) error {
	data, err := os.ReadFile(s.dbfilepath())

	if err != nil {
		return err
	}

	var storageMap StorageMap

	if err = json.Unmarshal(data, &storageMap); err != nil {
		return err
	}

	storageMap[host.Url] = host

	if data, err = json.Marshal(storageMap); err != nil {
		return err
	}

	if err = os.WriteFile(s.dbfilepath(), data, defaultPerm); err != nil {
		return err
	}

	return nil
}

func (s *JsonStorage) SaveElement(url string, name string, query string) (storage.Element, error) {
	host, err := s.Read(url)

	if err != nil {
		if errors.Is(err, ErrNoData) {
			host = storage.Host{
				Url:      url,
				Bindings: make([]storage.Element, 0),
			}
		} else {
			return storage.Element{}, err
		}
	}

	id := 0

	if len(host.Bindings) > 0 {
		id = host.Bindings[len(host.Bindings)-1].Id + 1
	}

	el := storage.Element{
		Id:    id,
		Name:  name,
		Query: query,
	}

	host.Bindings = append(host.Bindings, el)

	err = s.Save(host)

	if err != nil {
		return storage.Element{}, err
	}

	return el, nil
}

func (s *JsonStorage) Read(url string) (storage.Host, error) {
	data, err := os.ReadFile(s.dbfilepath())

	if err != nil {
		return storage.Host{}, err
	}

	var storageMap StorageMap

	if err = json.Unmarshal(data, &storageMap); err != nil {
		return storage.Host{}, err
	}

	host, ok := storageMap[url]

	if ok {
		return host, nil
	}

	return storage.Host{}, ErrNoData
}

func (s *JsonStorage) GetElementById(url string, id int) (storage.Element, error) {
	host, err := s.Read(url)

	if err != nil {
		return storage.Element{}, err
	}

	for _, binding := range host.Bindings {
		if binding.Id == id {
			return binding, nil
		}
	}

	return storage.Element{}, ErrNoData
}

func (s *JsonStorage) Delete(url string) error {
	data, err := os.ReadFile(s.dbfilepath())

	if err != nil {
		return err
	}

	var storageMap StorageMap

	if err = json.Unmarshal(data, &storageMap); err != nil {
		return err
	}

	delete(storageMap, url)

	return nil
}

func (s *JsonStorage) DeleteElement(url string, id int) error {
	host, err := s.Read(url)

	if err != nil {
		return err
	}

	for i := 0; i < len(host.Bindings); i++ {
		if host.Bindings[i].Id == id {
			host.Bindings = append(host.Bindings[:i], host.Bindings[i+1:]...)
		}
	}

	err = s.Save(host)

	return err
}

func (s *JsonStorage) dbfilepath() string {
	return path.Join(s.Path, StorageFilename)
}
