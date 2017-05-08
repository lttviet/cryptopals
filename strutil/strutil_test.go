package strutil

import (
	"bytes"
	"testing"
)

func TestHexToBase64(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"", ""},
		{"49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d",
			"SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"},
	}

	for _, c := range cases {
		got := HexToBase64(c.in)
		if got != c.want {
			t.Errorf("HexToBase64(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestPKCS7Padding(t *testing.T) {
	cases := []struct {
		in        []byte
		blocksize int
		want      []byte
	}{
		{[]byte("YELLOW SUBMARINE"), 20, []byte("YELLOW SUBMARINE\x04\x04\x04\x04")},
		{[]byte("HELLO"), 8, []byte("HELLO\x03\x03\x03")},
	}

	for _, c := range cases {
		got := PKCS7Padding(c.in, c.blocksize)
		if !bytes.Equal(got, c.want) {
			t.Errorf("PKCS7Padding(%q, %q) == %q, want %q", c.in, c.blocksize, got, c.want)
		}
	}
}
