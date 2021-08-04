package crypto

import (
	"crypto/ecdsa"
	"crypto/rand"
	"deltaboard-server/internal/library/base36"
	"encoding/hex"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto"
)

const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"

// GenPrivKeySecp256k1
func GenPrivateKeySecp256k1() (string, string) {
	p, err := crypto.GenerateKey()
	if err != nil {
		return "", ""
	}
	return hex.EncodeToString(crypto.FromECDSA(p)), hex.EncodeToString(crypto.CompressPubkey(&p.PublicKey))
}

// GetPubKeyFrom
func GetPubKeyFrom(pubkey string) (*ecdsa.PublicKey, error) {
	if d, err := hex.DecodeString(pubkey); err != nil {
		return nil, err
	} else {
		return crypto.DecompressPubkey(d)
	}
}

// Address
func Address(pubkey string) (string, error) {
	if d, err := hex.DecodeString(pubkey); err != nil {
		return "", err
	} else {
		pub, err := crypto.DecompressPubkey(d)
		if err != nil {
			return "", err
		} else {
			return crypto.PubkeyToAddress(*pub).String(), nil
		}
	}
}

func GetPubKeyFromPri(private_key []byte) (pubKey string) {
	p, err := crypto.ToECDSA(private_key)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(crypto.CompressPubkey(&p.PublicKey))
}
func GetPubKeyFromPriString(private_key string) (pubKey string) {
	priBs, err := hex.DecodeString(private_key)
	if err != nil || len(priBs) != 32 {
		return ""
	}
	p, err := crypto.ToECDSA(priBs)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(crypto.CompressPubkey(&p.PublicKey))
}

// GetPriKeyFrom
func GetPriKeyFrom(prikey string) (*ecdsa.PrivateKey, error) {
	if d, err := hex.DecodeString(prikey); err != nil {
		return nil, err
	} else {
		return crypto.ToECDSA(d)
	}
}

// Sign
func Sign(hash, prv []byte) (sig []byte, err error) {
	p, err := crypto.ToECDSA(prv)
	if err != nil {
		return nil, err
	}
	return crypto.Sign(hash, p)
}

func SignString(hash []byte, prv string) (sig []byte, err error) {
	pBs, err := hex.DecodeString(prv)
	if err != nil {
		return nil, err
	}
	p, err := crypto.ToECDSA(pBs)
	if err != nil {
		return nil, err
	}
	return crypto.Sign(hash, p)
}

// VerifySignature
func VerifySignature(pubKey, hash, signature []byte) bool {
	signBs := signature
	if len(signature) == 65 {
		signBs = signature[:len(signature)-1]
	}
	return crypto.VerifySignature(pubKey, hash, signBs) // remove recovery id
}

func VerifySignatureByUnCompressPublicKey(pubKey, hash, signature []byte) bool {
	signBs := signature
	if len(signature) == 65 {
		signBs = signature[:len(signature)-1]
	}
	return crypto.VerifySignature(pubKey[:33], hash, signBs) // remove recovery id
}

// Hasher
func Keccak256(data ...[]byte) []byte {
	return crypto.Keccak256(data...)
}

// GenerateDNA
func GenerateDNA(md_sign string) string {
	d, err := hex.DecodeString(md_sign)
	if err != nil {
		return ""
	}
	digest := base36.EncodeBytes(crypto.Keccak256(d))
	return digest
}
func GenerateRandomString(n int) (string, error) {
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		ret[i] = letters[num.Int64()]
	}

	return string(ret), nil
}
