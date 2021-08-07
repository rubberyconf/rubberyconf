package feature

type QueryParam struct {
	Key   string   `yaml:"key"`
	Value []string `yaml:"value"`
}

func queryString(q QueryParam, vars map[string]string) bool {

	value, ok := vars[q.Key]
	if !ok {
		return false
	}

	for _, candidate := range q.Value {
		if candidate == value {
			return true
		}

	}

	return false
}
