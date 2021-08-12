package rules

import "container/list"

type RuleVersion struct {
}

func (me *RuleVersion) CheckRule(r FeatureRule, vars map[string]string, matches *list.List) (bool, bool) {
	if len(r.Version) > 0 && len(vars["version"]) > 0 {
		sentByclient := vars["version"]
		ok := me.versionCheck(r.Version, sentByclient)
		if ok {
			matches.PushBack("querystring")
		}
		return ok, false
	}
	return false, true
}

func (me *RuleVersion) versionOrdinal(version string) string {
	// ISO/IEC 14651:2011
	const maxByte = 1<<8 - 1
	vo := make([]byte, 0, len(version)+8)
	j := -1
	for i := 0; i < len(version); i++ {
		b := version[i]
		if '0' > b || b > '9' {
			vo = append(vo, b)
			j = -1
			continue
		}
		if j == -1 {
			vo = append(vo, 0x00)
			j = len(vo) - 1
		}
		if vo[j] == 1 && vo[j+1] == '0' {
			vo[j+1] = b
			continue
		}
		if vo[j]+1 > maxByte {
			panic("VersionOrdinal: invalid version")
		}
		vo = append(vo, b)
		vo[j]++
	}
	return string(vo)
}

func (me *RuleVersion) versionCheck(objective []string, sentBy string) bool {

	for _, obj := range objective {

		fist := string(obj[0])

		if fist == ">" {
			o := len(obj)
			newObj := obj[1:o]
			a, b := me.versionOrdinal(newObj), me.versionOrdinal(sentBy)
			if b > a {
				return true
			} else {
				return false
			}
		}
		a, b := me.versionOrdinal(obj), me.versionOrdinal(sentBy)
		if a == b {
			return true
		}
		/*
			switch {
			case a > b:
				return false
			case a < b:
				return false
			case a == b:
				return true
			}*/
	}

	return false
}
