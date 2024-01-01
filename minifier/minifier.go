package minifier

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"os"

	"github.com/itchyny/base58-go"
)

func sha256Of(url string) []byte {
	algo := sha256.New()
	algo.Write([]byte(url))
	return algo.Sum(nil)
}

func base58Encode(bytes []byte) string {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(bytes)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return string(encoded)
}

func GenreateMinUrl(original_url string, user_id string) string {
	urlHashBytes := sha256Of(original_url + user_id)
	generatedNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()
	final_string := base58Encode([]byte(fmt.Sprintf("%d", generatedNumber)))
	return final_string[:8]
}
