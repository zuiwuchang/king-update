package update

import (
	"crypto/sha512"
	"encoding/hex"
	"io"
)

// HashReader 計算 sha512
func HashReader(r io.Reader) (hash string, e error) {
	w := sha512.New()
	if _, e = io.Copy(w, r); e != nil {
		return
	}
	hash = hex.EncodeToString(w.Sum(nil))
	return
}
