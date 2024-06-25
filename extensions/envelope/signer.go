package envelope

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/btcsuite/btcutil/base58"
)

type (
	Signer interface {
		CreateNonce() string
		Hash(payload []byte, nonce, channel, chaincode, method, deadline string, publicKey []byte) []byte
		Sign(payload []byte, nonce, channel, chaincode, method, deadline string, privateKey []byte) ([]byte, error)
		CheckSignature(payload []byte, nonce, channel, chaincode, method, deadline string, publicKey []byte, sig []byte) error
	}

	DefaultSigner struct {
		Crypto Crypto
	}
)

func NewSigner(crypto Crypto) *DefaultSigner {
	return &DefaultSigner{Crypto: crypto}
}

func removeSpacesBetweenCommaAndQuotes(s []byte) []byte {
	removed := strings.ReplaceAll(string(s), `", "`, `","`)
	removed = strings.ReplaceAll(removed, `"}, {"`, `"},{"`)
	return []byte(strings.ReplaceAll(removed, `], "`, `],"`))
}

func (b *DefaultSigner) CreateNonce() string {
	return strconv.Itoa(int(time.Now().Unix()))
}

func (b *DefaultSigner) Hash(payload []byte, nonce, channel, chaincode, method, deadline string, pubkey []byte) []byte {
	bb := append(removeSpacesBetweenCommaAndQuotes(payload), nonce...) // resolve the unclear json serialization behavior in protojson package
	bb = append(bb, channel...)
	bb = append(bb, chaincode...)
	bb = append(bb, method...)
	bb = append(bb, deadline...)
	b58Pubkey := base58.Encode(pubkey)
	bb = append(bb, b58Pubkey...)
	return b.Crypto.Hash(bb)
}

func (b *DefaultSigner) Sign(
	payload []byte, nonce, channel, chaincode, method, deadline string, privateKey []byte) ([]byte, error) {
	pubKey, err := b.Crypto.PublicKey(privateKey)
	if err != nil {
		return nil, fmt.Errorf(`extract public key: %w`, err)
	}

	hashed := b.Hash(payload, nonce, channel, chaincode, method, deadline, pubKey)
	return b.Crypto.Sign(hashed, privateKey)
}

func (b *DefaultSigner) CheckSignature(payload []byte, nonce, channel, chaincode, method, deadline string, pubKey []byte, sig []byte) error {
	hashed := b.Hash(payload, nonce, channel, chaincode, method, deadline, pubKey)
	if !b.Crypto.Verify(pubKey, hashed, sig) {
		return ErrCheckSignatureFailed
	}
	return nil
}
