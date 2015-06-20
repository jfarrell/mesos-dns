package labels

import (
	"fmt"
)

// LabelSeparator delimits concatenated labels.
const LabelSeparator = '.'

// Spec is a label specification interface
type Spec interface {
	// Mangle transforms an input string to produce an output string that is
	// compliant with the specification. If the input string is already
	// compliant then it is returned unchanged.
	Mangle(string) string
}

// SpecFunc implements Spec
type SpecFunc func(string) string

func (sf SpecFunc) Mangle(input string) string {
	return sf(input)
}

type HostNameSpec int

const (
	// see http://www.rfc-base.org/txt/rfc-952.txt
	HostNameSpec952 HostNameSpec = iota

	// see http://www.rfc-base.org/txt/rfc-1123.txt
	HostNameSpec1123
)

func (spec HostNameSpec) Spec() Spec {
	switch spec {
	case HostNameSpec952:
		return SpecFunc(AsDNS952)
	case HostNameSpec1123:
		return SpecFunc(AsDNS1123)
	}
	panic(fmt.Sprintf("bad HostNameSpec: %#v", spec))
}
