package aghash

import (
	"crypto/md5" // #nosec
	"encoding/base64"
	"encoding/hex"

	"github.com/davecgh/go-spew/spew"
	"github.com/pkg/errors"
)

// Hash computes a hash of the provided object
func Hash(i interface{}) ([]byte, error) {
	h := md5.New() // #nosec
	printer := spew.ConfigState{
		Indent:         " ",
		SortKeys:       true,
		DisableMethods: true,
		SpewKeys:       true,
	}
	_, err := printer.Fprintf(h, "%#v", i)
	return h.Sum(nil), errors.Wrap(err, "can't take hash")
}

// HashSet computes a hash of the provided set of objects.
// Hash will have the same value regardless of objects order.
func HashSet(elems ...interface{}) ([]byte, error) {
	result := make([]byte, md5.Size)
	for _, elem := range elems {
		h, err := Hash(elem)
		if err != nil {
			return result, err
		}
		for i := range result {
			result[i] = result[i] ^ h[i]
		}
	}
	return result, nil
}

// HashBase64 computes a string hash of the provided object
func HashBase64(i interface{}) (string, error) {
	bs, err := Hash(i)
	return base64.StdEncoding.EncodeToString(bs), err
}

// HashSetBase64 computes a string hash of the provided set of objects
func HashSetBase64(elems ...interface{}) (string, error) {
	bs, err := HashSet(elems...)
	return base64.StdEncoding.EncodeToString(bs), err
}

// HashHex takes a hash and encodes hash into hex
func HashHex(content []byte) (string, error) {
	hash, err := Hash(content)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(hash), nil
}
