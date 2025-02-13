package text2emoji

import (
	"fmt"
	"testing"
)

func TestEncrypt(t *testing.T) {
	s1 := "hello world ABC 123 你好"
	es1 := Encrypt(s1)
	fmt.Println(es1)
	s2 := Decrypt(es1)
	if s1 != s2 {
		t.Errorf("expected %s, got %s", s1, s2)
	}
	fmt.Println(s2)
}
