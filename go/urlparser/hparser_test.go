package urlparser

import (
	"testing"
)

func FuzzParseURL(f *testing.F) {
	//Seeds
	f.Add("https://archlinux.org")
	f.Add("http://user:pass@localhost:8080/api?q=1#frag")

	f.Fuzz(func(t *testing.T, raw string) {
		_, _ = ParseURL(raw)
	})
}
