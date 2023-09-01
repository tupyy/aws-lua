package twt

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/tupyy/aws-lua/internal/lua"
	"golang.org/x/crypto/blake2b"
)

var (
	InvalidSignatureError        = errors.New("invalid signature")
	InvalidTwtError              = errors.New("invalid twt")
	MismatchedObjectContentError = errors.New("object changed")
)

type TwtProvider struct {
	SecretKey string
}

func New(secretKey string) *TwtProvider {
	return &TwtProvider{SecretKey: secretKey}
}

// Create a twt
func (t *TwtProvider) Create(ctx context.Context, o lua.Object) (string, error) {
	jsonData, err := json.Marshal(o)
	if err != nil {
		return "", err
	}

	header := createHeader()
	tag, err := hash(jsonData)
	if err != nil {
		return "", err
	}

	signData := bytes.NewBufferString(fmt.Sprintf("%s%x", header, tag))
	signature, err := sign(signData.Bytes(), t.SecretKey)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s.%x.%x", header, tag, signature), nil
}

// Verify the twt
// Input:
// - spec: the objet to be verified
// - twt: the twt
// - error: the error if something is wrong with either twt or the object
func (t *TwtProvider) Verify(ctx context.Context, o lua.Object, twt string) error {
	jsonData, err := json.Marshal(o)
	if err != nil {
		return err
	}

	parts := strings.Split(twt, ".")
	if len(parts) != 3 {
		return InvalidTwtError
	}

	// verify the signature of the twt
	sigData := bytes.NewBufferString(fmt.Sprintf("%s%s", parts[0], parts[1]))
	newSignature, err := sign(sigData.Bytes(), t.SecretKey)
	if err != nil {
		return err
	}

	if fmt.Sprintf("%x", newSignature) != parts[2] {
		return InvalidSignatureError
	}

	tag, err := hash(jsonData)
	if err != nil {
		return err
	}

	if fmt.Sprintf("%x", tag) != parts[1] {
		return MismatchedObjectContentError
	}

	return nil
}

func hash(data []byte) ([]byte, error) {
	hash, err := blake2b.New(32, []byte{})
	if err != nil {
		return []byte{}, err
	}

	_, err = hash.Write(data)
	if err != nil {
		return []byte{}, err
	}

	return hash.Sum(nil), nil
}

func createHeader() string {
	// set version
	type version struct {
		Version string `json:"ver"`
		Alg     string `json:"alg"`
	}
	v := version{Version: "1", Alg: "blake2b"}
	vdata, _ := json.Marshal(v)

	// base64
	return base64.RawStdEncoding.EncodeToString(vdata)
}

func sign(hash []byte, secretKey string) ([]byte, error) {
	signMac, err := blake2b.New(32, bytes.NewBufferString(secretKey).Bytes())
	if err != nil {
		return []byte{}, err
	}

	if _, err := signMac.Write(hash); err != nil {
		return []byte{}, err
	}

	return signMac.Sum(nil), nil
}
