package slices

type StringSlice []string
type IntSlice []int
type UIntSlice []uint

// Unique make string slice items unique
func (s *StringSlice) Unique() {
	check := make(map[string]bool)
	var res []string

	for _, val := range *s {
		if _, ok := check[val]; !ok {
			check[val] = true
			res = append(res, val)
		}
	}
	*s = res
}

// Contains looks for the needle
func (s StringSlice) Contains(needle string) bool {
	for _, v := range s {
		if v == needle {
			return true
		}
	}
	return false
}

func (s *StringSlice) Filter(f func(val string) bool) {
	result := make([]string, 0)
	for _, v := range *s {
		if f(v) {
			result = append(result, v)
		}
	}
	*s = result
}

// Unique make int slice items unique
func (s *IntSlice) Unique() {
	check := make(map[int]bool)
	var res []int

	for _, val := range *s {
		if _, ok := check[val]; !ok {
			check[val] = true
			res = append(res, val)
		}
	}

	*s = res
}

// Contains looks for the needle
func (s IntSlice) Contains(needle int) bool {
	for _, v := range s {
		if v == needle {
			return true
		}
	}
	return false
}

// Unique make uint slice items unique
func (s *UIntSlice) Unique() {
	check := make(map[uint]bool)
	var res []uint

	for _, val := range *s {
		if _, ok := check[val]; !ok {
			check[val] = true
			res = append(res, val)
		}
	}

	*s = res
}

// Contains looks for the needle
func (s UIntSlice) Contains(needle uint) bool {
	for _, v := range s {
		if v == needle {
			return true
		}
	}
	return false
}
