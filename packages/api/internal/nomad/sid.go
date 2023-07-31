package nomad

import (
	"math/rand"
)

var sessionAlphabet = []rune("abcdefghijklmnopqrstuvwxyz1234567890")

func genRandomSession(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = sessionAlphabet[rand.Intn(len(sessionAlphabet))]
	}
	return string(b)
}

var idAlphabet = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func GenRandomID(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = idAlphabet[rand.Intn(len(idAlphabet))]
	}
	return string(b)
}