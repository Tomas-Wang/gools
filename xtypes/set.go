package xtypes

type Set map[interface{}]struct{}

func NewSet() Set {
	return make(Set)
}

func NewSetWithSlice(data []interface{}) Set {
	s := make(Set)
	for _, item := range data {
		s.Add(item)
	}
	return s
}

func (s Set) Add(data interface{}) {
	s[data] = struct{}{}
}

func (s Set) All() []interface{} {
	result := make([]interface{}, 0)
	for k := range s {
		result = append(result, k)
	}
	return result
}
