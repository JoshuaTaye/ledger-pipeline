package config

import "testing"

func TestParseProfile(t *testing.T) {
    p, err := ParseProfile([]byte(`{"name":"home","default_currency":"USD","categories":["Food"]}`))
    if err != nil || p.Name != "home" {
        t.Fatalf("err=%v p=%+v", err, p)
    }
}
