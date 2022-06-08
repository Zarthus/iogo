package internal

func Wrap(s string, maxlength uint) []string {
	staticSlen := uint(len(s))
	slen := staticSlen
	if maxlength > slen {
		return []string{s}
	}

	var lines []string
	pointer := uint(0)

	for slen > maxlength {
		if pointer > staticSlen {
			break
		}

		pterEnd := pointer + maxlength
		if pterEnd > staticSlen {
			pterEnd = staticSlen
		}
		lines = append(lines, s[pointer:pterEnd])

		pointer += maxlength
		slen = maxlength - slen
	}

	return lines
}
