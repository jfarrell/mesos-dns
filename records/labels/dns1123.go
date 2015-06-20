package labels

const DNS1123MaxLength int = 63

// mangle the given name to be compliant as a DNS1123 name "component".
// the returned result should match the regexp:
//    ^[a-z0-9]([-a-z0-9]*?[a-z0-9])?$
// returns "" if the name cannot be mangled.
func AsDNS1123(name string) string {
	if name == "" {
		return ""
	}
	sz := len(name)
	if sz > DNS1123MaxLength {
		sz = DNS1123MaxLength
	}
	last := sz - 1
	label := make([]byte, sz, sz)
	ll := 0
	la := -1 // index of last alphanumeric
	for _, c := range name {
		b := dns952table[uint8(c)]
		switch {
		case b == -int32('-'):
			if ll == 0 || ll == last {
				continue
			}
			b = -b
		case b < 0:
			b = -b
			fallthrough
		case b > 0:
			la = ll
		default:
			continue
		}
		label[ll] = byte(b)
		ll++
		if ll == sz {
			break
		}
	}
	if ll > 0 && label[ll-1] == '-' {
		ll = la + 1
	}
	if ll > 0 {
		return string(label[:ll])
	}
	return ""
}
