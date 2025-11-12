package exp

func Ptr[T any](val T) *T {
	return &val
}

func SlicePtr[T any](v []T) []*T {
	out := make([]*T, len(v))
	for i := range v {
		out[i] = &v[i]
	}
	return out
}

func IsNil(val any) bool {
	return val == nil
}

func In[T comparable](val T, slice []T) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}
