package cashe

type Cashe interface {
	Get(string) string
	Set(string, string)
	Delete(string)
}

type MemoryCashe struct {
	cashe map[string]string
}

func (m MemoryCashe) Get(key string) string {
	return m.cashe[key]
}

func (m MemoryCashe) Set(key string, value string) {
	m.cashe[key] = value
}

func (m MemoryCashe) Delete(key string) {
	delete(m.cashe, key)
}

func New() Cashe {
	return MemoryCashe{cashe: make(map[string]string)}
}
