package datasource

func enableFeature(keys map[string]string) (Feature, bool) {
	fe1 := Feature{Key: "", Value: nil}
	key := keys[keyFeature]
	if key == "" {
		return fe1, false
	}
	fe1.Key = key
	return fe1, true
}
