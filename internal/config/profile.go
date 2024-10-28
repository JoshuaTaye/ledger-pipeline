package config

import "encoding/json"

type Profile struct {
    Name            string   `json:"name"`
    DefaultCurrency string   `json:"default_currency"`
    Categories      []string `json:"categories"`
}

func ParseProfile(data []byte) (Profile, error) {
    var p Profile
    err := json.Unmarshal(data, &p)
    return p, err
}
