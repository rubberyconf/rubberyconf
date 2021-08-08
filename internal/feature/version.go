package feature

func VersionOrdinal(version string) string {
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

func versionCheck(objective []string, sentBy string) bool {

	for _, obj := range objective {

		fist := string(obj[0])

		if fist == ">" {
			o := len(obj)
			newObj := obj[1:o]
			a, b := VersionOrdinal(newObj), VersionOrdinal(sentBy)
			if b > a {
				return true
			} else {
				return false
			}
		}
		a, b := VersionOrdinal(obj), VersionOrdinal(sentBy)
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
