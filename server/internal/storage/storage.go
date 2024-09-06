package storage

type ElementInfo struct {
	Host         string `json:"host"`
	ElementQuery string `json:"element-qeury"`
}

type Storage interface {
	Save(host string, bindings []string) error
	SaveElement(elementInfo ElementInfo) (int, error)
	Read(host string) ([]string, error)
}
