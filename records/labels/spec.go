package labels

import (
	"fmt"
)

const (
	// LabelSeparator delimits concatenated labels.
	SeparatorChar = '.'

	DNS952MaxLength  int = 24
	DNS1123MaxLength int = 63
)

// Mangler is a label mangling interface
type Mangler interface {
	// Mangle transforms an input string to produce an output string that is
	// compliant with the specification. If the input string is already
	// compliant then it is returned unchanged.
	Mangle(string) string
}

// ManglerFunc implements Mangler
type ManglerFunc func(string) string

func (sf ManglerFunc) Mangle(input string) string {
	return sf(input)
}

type HostNameSpec int

const (
	// see http://www.rfc-base.org/txt/rfc-952.txt
	HostNameSpec952 HostNameSpec = iota

	// see http://www.rfc-base.org/txt/rfc-1123.txt
	HostNameSpec1123
)

func (spec HostNameSpec) Mangler() Mangler {
	switch spec {
	case HostNameSpec952:
		return ManglerFunc(func(s string) string { return dns952table.toLabel(s, DNS952MaxLength) })
	case HostNameSpec1123:
		return ManglerFunc(func(s string) string { return dns1123table.toLabel(s, DNS1123MaxLength) })
	}
	panic(fmt.Sprintf("bad HostNameSpec: %#v", spec))
}

func (spec HostNameSpec) AsDomainFrag(s string) string {
	return AsDomainFrag(s, spec.Mangler())
}
