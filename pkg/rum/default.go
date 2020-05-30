package rum

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
	"net/http"
)

func DefaultHash(r *http.Request) string {
	h := sha1.New()
	io.WriteString(h, r.Method)
	io.WriteString(h, r.RequestURI)
	return hex.EncodeToString(h.Sum(nil))
}
