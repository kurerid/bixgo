package bixgo

// Params представляет собой список параметров для запросов к API Bitrix24.
type Params map[string]any

func (p Params) Set(key string, value any) {
	p[key] = value
}

func (p Params) Get(key string) (any, bool) {
	value, found := p[key]
	return value, found
}
