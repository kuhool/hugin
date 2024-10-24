package common

// 是否选项
type Is int

const (
	// 否
	IsNo Is = iota + 0
	// 是
	IsYes
)

// 转换成整型
func (is Is) ToInt() int {
	return int(is)
}

// 转换成布尔型
func (is Is) ToBool() bool {
	return is == IsYes
}

type User struct {
	IDInt uint64 `json:"id_int"`
	Name  string `json:"name"`
	ID    string `json:"id"`
}

// 一维map
type OneDimensionValue map[string]uint

func (v *OneDimensionValue) Incr(k string, n uint) {
	rv := *v
	rv[k] += n
}

func (v OneDimensionValue) ToNameValue() []*NameValue {
	r := []*NameValue{}
	for v1, k1 := range v {
		r = append(r, &NameValue{
			Name:  v1,
			Value: k1,
		})
	}
	return r
}

// 二维map
type TwoDimensionValue map[string]map[string]uint

func (v *TwoDimensionValue) Incr(k1 string, k2 string, n uint) {
	rv := *v
	if _, ok := rv[k1]; ok {
		rv[k1][k2] += n
		return
	}
	rv[k1] = map[string]uint{k2: n}
}

func (v TwoDimensionValue) TopValue(k string) map[string]uint {
	if val, ok := v[k]; ok {
		return val
	}
	return map[string]uint{}
}

func (v *TwoDimensionValue) SetTopValue(k string, val map[string]uint) {
	vv := *v
	vv[k] = val
}

func (v TwoDimensionValue) ToNameValue() map[string][]*NameValue {
	r := map[string][]*NameValue{}
	for v1, k1 := range v {
		for v2, k2 := range k1 {
			r[v1] = append(r[v1], &NameValue{
				Name:  v2,
				Value: k2,
			})
		}
	}
	return r
}

type NameValue struct {
	Name  string `json:"name"`
	Value uint   `json:"value"`
}
