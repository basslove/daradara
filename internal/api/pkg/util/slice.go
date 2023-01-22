package util

func BuildUniqueStringSlice(l []string) []string {
	stringMap := make(map[string]struct{})
	uniqueStrings := make([]string, 0, len(l))

	for _, s := range l {
		if _, ok := stringMap[s]; !ok {
			stringMap[s] = struct{}{}
			uniqueStrings = append(uniqueStrings, s)
		}
	}

	return uniqueStrings
}
