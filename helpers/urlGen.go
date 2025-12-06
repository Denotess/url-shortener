package helpers

import "encoding/base32"

func GenerateShortUrl(oldUrl string) (string, error) {
	url := []byte(oldUrl)
	encodedUrl := base32.StdEncoding.EncodeToString(url)

	return encodedUrl, nil
}
