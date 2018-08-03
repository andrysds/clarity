package clarity

// MapJoin joins two variables type map into one. Only support map[string]string at the moment.
func MapJoin(mapA, mapB map[string]string) map[string]string {
	result := mapA
	for key, val := range mapB {
		mapA[key] = val
	}
	return result
}
