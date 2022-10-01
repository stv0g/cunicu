// Package test implements universal helpers for unit and integration tests
package test

import (
	"bytes"
	"crypto/rand"
	"math"
	"math/big"

	"github.com/stv0g/cunicu/pkg/crypto"
	"github.com/stv0g/cunicu/pkg/signaling"

	protoepdisc "github.com/stv0g/cunicu/pkg/proto/feature/epdisc"
)

func GenerateKeyPairs() (*crypto.KeyPair, *crypto.KeyPair, error) {
	ourKey, err := crypto.GeneratePrivateKey()
	if err != nil {
		return nil, nil, err
	}

	theirKey, err := crypto.GeneratePrivateKey()
	if err != nil {
		return nil, nil, err
	}

	return &crypto.KeyPair{
			Ours:   ourKey,
			Theirs: theirKey.PublicKey(),
		}, &crypto.KeyPair{
			Ours:   theirKey,
			Theirs: ourKey.PublicKey(),
		}, nil
}

func GenerateSignalingMessage() *signaling.Message {
	r, err := rand.Int(rand.Reader, big.NewInt(math.MaxInt64))
	if err != nil {
		panic(err)
	}

	return &signaling.Message{
		Candidate: &protoepdisc.Candidate{
			Port: int32(r.Int64()),
		},
	}
}

func Entropy(data []byte) float64 {
	if len(data) == 0 {
		return 0
	}

	var length = float64(len(data))
	var entropy = 0.0

	for i := 0; i < 256; i++ {
		if p := float64(bytes.Count(data, []byte{byte(i)})) / length; p > 0 {
			entropy += -p * math.Log2(p)
		}
	}

	return entropy
}
