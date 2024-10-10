package storage

type Element struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Query string `json:"qeury"`
}

type Host struct {
	Url      string    `json:"url"`
	Bindings []Element `json:"bindings"`
}

type Storage interface {
	Save(host Host) error
	SaveElement(url string, element Element) error
	Read(url string) (Host, error)
	Delete(url string) error
	DeleteElement(url string, id int) error
}
