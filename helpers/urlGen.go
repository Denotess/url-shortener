package helpers

import "github.com/blakewilliams/go-base36"

func GenerateShortUrl(id int64) (string, error) {
	encodedId := base36.StdEncoding.Encode(id)
	return encodedId, nil
}
