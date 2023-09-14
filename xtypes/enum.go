package xtypes

type enumValue struct {
	Number int    `json:"number,omitempty"`
	Desc   string `json:"desc,omitempty"`
}

// Enum custom enum type
type Enum map[string]*enumValue

// NewEnum new enum
func NewEnum() Enum {
	return make(Enum)
}

func (e Enum) NewMember(key string, number int, desc string) {
	e[key] = &enumValue{
		Number: number,
		Desc:   desc,
	}
}

func (e Enum) Number(key string) int {
	value, ok := e[key]
	if !ok {
		return 0
	}
	return value.Number
}

func (e Enum) Desc(key string) string {
	value, ok := e[key]
	if !ok {
		return ""
	}
	return value.Desc
}
