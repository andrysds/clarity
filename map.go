package clarity

// MapJoin joins two variables type map into one. Only support map[string]string at the moment.
func MapJoin(a, b map[string]string) {
	for key, val := range b {
		a[key] = val
	}
}
