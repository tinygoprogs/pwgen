package pwgen

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"io"
	"math/big"
	"strings"
)

type GenPasswordConfig struct {
	RandomReader             io.Reader
	Alphabet                 *string
	MustExclude, MustInclude string
	Len                      int
}

const DefaultAlphabet = `abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!?@#$%^&*()_+-/\<>`

func randIntMax255(rnd io.Reader) int {
	b := make([]byte, 1)
	n, err := rnd.Read(b)
	if err != nil || n != 1 {
		panic(fmt.Sprintf("err=%v, n=%d", err, n))
	}
	return int(b[0])
}

func choice(alphabet string, rnd io.Reader) byte {
	max := big.NewInt(int64(len(alphabet)-1))
	index, err := rand.Int(rnd, max)
	if err != nil {
		panic(err)
	}
	return alphabet[index.Int64()]
}

func modifyAlphabet(in, exclude, include string) string {
	out := in
	for _, char := range exclude {
		out = strings.ReplaceAll(out, string(char), "")
	}
	for _, char := range include {
		if !strings.Contains(out, string(char)) {
			out += string(char)
		}
	}
	return out
}

func containsAll(pw []byte, includes string) bool {
	for _, b := range includes {
		if !bytes.ContainsRune(pw, b) {
			return false
		}
	}
	return true
}

func replaceRandom(in []byte, target byte, rnd io.Reader) {
	in[randIntMax255(rnd)%len(in)] = target
}

func insertEachAtRandomPos(pw []byte, includes string, rnd io.Reader) []byte {
	maxtries := 100
	for {
		for _, char := range includes {
			if !bytes.ContainsRune(pw, char) {
				replaceRandom(pw, byte(string(char)[0]), rnd)
			}
		}
		if containsAll(pw, includes) {
			break
		}
		maxtries--
		if maxtries == 0 {
			panic("failed to generate password: too many includes / too small len")
		}
	}
	return pw
}

func GenPassword(c *GenPasswordConfig) string {
	if c.RandomReader == nil {
		c.RandomReader = rand.Reader
	}
	if c.Alphabet == nil {
		copyofdefault := DefaultAlphabet
		c.Alphabet = &copyofdefault
	}
	pw := []byte{}
	alphabet := modifyAlphabet(*c.Alphabet, c.MustExclude, c.MustInclude)
	for i := 0; i < c.Len; i++ {
		pw = append(pw, choice(alphabet, c.RandomReader))
	}
	pw = insertEachAtRandomPos(pw, c.MustInclude, c.RandomReader)
	return string(pw)
}
