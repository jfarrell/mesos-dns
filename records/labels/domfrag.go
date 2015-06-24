package labels

// mangles the given name in order to produce a valid domain fragment.
// a valid domain fragment will consist of one or more host name labels
// concatenated by a SeparatorChar char.
func AsDomainFrag(name string, spec Mangler) string {
	if name == "" {
		return ""
	}
	sz := len(name)
	frag := make([]byte, sz, sz)
	ll := 0  // overall fragment length so far
	li := -1 // last fragment we found ended here
	for i, c := range name {
		if c == SeparatorChar {
			if f := spec.Mangle(name[li+1 : i]); f != "" {
				if li > -1 {
					frag[ll] = SeparatorChar
					ll++
				}
				// len(f) is <= len(slice-of-name)
				copy(frag[ll:], f)
				ll += len(f)
				li = i
			}
		}
	}
	// final frag
	if f := spec.Mangle(name[li+1:]); f != "" {
		if li > -1 {
			frag[ll] = SeparatorChar
			ll++
		}
		copy(frag[ll:], f)
		ll += len(f)
	}
	if ll > 0 {
		return string(frag[:ll])
	}
	return ""
}
