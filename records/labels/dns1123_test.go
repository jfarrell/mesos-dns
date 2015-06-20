package labels

import (
	"testing"
)

func BenchmarkAsDNS1123(b *testing.B) {
	const (
		original = "##fdgsf---gs7-fgs--d7fddg123456789012345678901234567890123456789-"
		expected = "fdgsf---gs7-fgs--d7fddg123456789012345678901234567890123456789"
	)

	// run the AsDNS1123 func b.N times
	for n := 0; n < b.N; n++ {
		if actual := AsDNS1123(original); actual != expected {
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
	for untrusted, expected := range tests {
		actual := AsDNS1123(untrusted)
		if actual != expected {
			t.Fatalf("expected %q instead of %q after converting %q", expected, actual, untrusted)
		}
	}
}
