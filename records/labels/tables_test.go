package labels

import (
	"testing"
)

func BenchmarkAsDNS952(b *testing.B) {
	const (
		original = "89f.gsf---g_7-fgs--d7fddg-"
		expected = "f-gsf---g-7-fgs--d7fddg"
	)

	asDNS952 := HostNameSpec952.Mangler().Mangle
	// run the asDNS952 func b.N times
	for n := 0; n < b.N; n++ {
		if actual := asDNS952(original); actual != expected {
			b.Fatalf("expected %q instead of %q", expected, actual)
		}
	}
}

func TestAsDNS952(t *testing.T) {
	tests := map[string]string{
		"":                                "",
		"a":                               "a",
		"-":                               "",
		"a---":                            "a",
		"---a---":                         "a",
		"---a---b":                        "a---b",
		"a.b.c.d.e":                       "a-b-c-d-e",
		"a.c.d_de.":                       "a-c-d-de",
		"abc123":                          "abc123",
		"4abc123":                         "abc123",
		"-abc123":                         "abc123",
		"abc123-":                         "abc123",
		"abc-123":                         "abc-123",
		"abc--123":                        "abc--123",
		"fd%gsf---gs7-f$gs--d7fddg-123":   "fdgsf---gs7-fgs--d7fddg1",
		"89fdgsf---gs7-fgs--d7fddg-123":   "fdgsf---gs7-fgs--d7fddg1",
		"89fdgsf---gs7-fgs--d7fddg---123": "fdgsf---gs7-fgs--d7fddg1",
		"89fdgsf---gs7-fgs--d7fddg-":      "fdgsf---gs7-fgs--d7fddg",
		"r29f.dev.angrypigs":              "r29f-dev-angrypigs",
	}
	asDNS952 := HostNameSpec952.Mangler().Mangle
	for untrusted, expected := range tests {
		actual := asDNS952(untrusted)
		if actual != expected {
			t.Fatalf("expected %q instead of %q after converting %q", expected, actual, untrusted)
		}
	}
}

func BenchmarkAsDNS1123(b *testing.B) {
	const (
		original = "##fdgsf---gs7-fgs--d7fddg123456789012345678901234567890123456789-"
		expected = "fdgsf---gs7-fgs--d7fddg123456789012345678901234567890123456789"
	)

	asDNS1123 := HostNameSpec1123.Mangler().Mangle
	// run the asDNS1123 func b.N times
	for n := 0; n < b.N; n++ {
		if actual := asDNS1123(original); actual != expected {
			b.Fatalf("expected %q instead of %q", expected, actual)
		}
	}
}

func TestAsDNS1123(t *testing.T) {
	tests := map[string]string{
		"":                                "",
		"a":                               "a",
		"-":                               "",
		"a---":                            "a",
		"---a---":                         "a",
		"---a---b":                        "a---b",
		"a.b.c.d.e":                       "a-b-c-d-e",
		"a.c.d_de.":                       "a-c-d-de",
		"abc123":                          "abc123",
		"4abc123":                         "4abc123",
		"-abc123":                         "abc123",
		"abc123-":                         "abc123",
		"abc-123":                         "abc-123",
		"abc--123":                        "abc--123",
		"fd%gsf---gs7-f$gs--d7fddg-123":   "fdgsf---gs7-fgs--d7fddg-123",
		"89fdgsf---gs7-fgs--d7fddg-123":   "89fdgsf---gs7-fgs--d7fddg-123",
		"89fdgsf---gs7-fgs--d7fddg---123": "89fdgsf---gs7-fgs--d7fddg---123",
		"89fdgsf---gs7-fgs--d7fddg-":      "89fdgsf---gs7-fgs--d7fddg",

		"fd%gsf---gs7-f$gs--d7fddg123456789012345678901234567890123456789-123":   "fdgsf---gs7-fgs--d7fddg1234567890123456789012345678901234567891",
		"$$fdgsf---gs7-fgs--d7fddg123456789012345678901234567890123456789-123":   "fdgsf---gs7-fgs--d7fddg1234567890123456789012345678901234567891",
		"%%fdgsf---gs7-fgs--d7fddg123456789012345678901234567890123456789---123": "fdgsf---gs7-fgs--d7fddg1234567890123456789012345678901234567891",
		"##fdgsf---gs7-fgs--d7fddg123456789012345678901234567890123456789-":      "fdgsf---gs7-fgs--d7fddg123456789012345678901234567890123456789",

		"r29f.dev.angrypigs": "r29f-dev-angrypigs",
	}
	asDNS1123 := HostNameSpec1123.Mangler().Mangle
	for untrusted, expected := range tests {
		actual := asDNS1123(untrusted)
		if actual != expected {
			t.Fatalf("expected %q instead of %q after converting %q", expected, actual, untrusted)
		}
	}
}
