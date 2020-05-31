package util

import (
	"github.com/gogf/gf/frame/g"
	"testing"
)

func TestCreateToken(t *testing.T) {
	token := CreateMagicToken(g.Map{
		"sss": "ddd",
	})
	t.Log(token)
}

func TestParseToken(t *testing.T) {
	token, _ := ParseMagicToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzc3MiOiJkZGQifQ.QFhXxqgxQxD1drBE_kCdZS9H75eZ4gSPFM-fq3ARywE")
	t.Log(token)
}
