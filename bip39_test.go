package mnemonic_test

import (
	_ "embed"
	"errors"
	"fmt"
	"math"
	"math/big"
	"strings"
	"testing"

	"github.com/DIMO-Network/mnemonic"
)

//go:embed trezorTestVectors.json
var trezorTestVectors []byte

func TestToAndFromMnemonic(t *testing.T) {
	t.Parallel()
	for originalNum := 0; originalNum < math.MaxUint16; originalNum++ {
		originalNum := originalNum
		t.Run(fmt.Sprintf("Number:0x%x", originalNum), func(t *testing.T) {
			t.Parallel()
			// Convert the number to a mnemonic
			words := mnemonic.FromBigInt(big.NewInt(int64(originalNum)))
			// Convert the mnemonic back to a number
			decodedNum, err := mnemonic.ToBigInt(words)
			if err != nil {
				t.Errorf("Error converting mnemonic to int for original number: %v, error: %v", originalNum, err)
				return
			}
			// If the original number and the decoded number are not the same, then the test fails
			if originalNum != int(decodedNum.Int64()) {
				t.Errorf("Original number: %v, Decoded number: %v", originalNum, decodedNum)
				return
			}
		})
	}
}

func TestTrezorEnglishVectors(t *testing.T) {
	t.Parallel()
	for _, vector := range testVectors.English {
		vector := vector
		t.Run(vector[0], func(t *testing.T) {
			t.Parallel()
			hex, ok := big.NewInt(0).SetString(vector[0], 16)
			if !ok {
				t.Errorf("Invalid hex: %v", vector[0])
				return
			}

			// Convert the mnemonic to a number
			words, err := mnemonic.FromBigIntFixed(hex, len(vector[0])*4)
			if err != nil {
				t.Errorf("Failed to create mnemonic from hex: %v", err)
				return
			}

			if strings.Join(words, " ") != vector[1] {
				t.Errorf("Expected mnemonic: %v\n Got mnemonic: %v", vector[1], words)
				return
			}
			decodedNum, err := mnemonic.ToBigInt(strings.Split(vector[1], " "))
			if err != nil {
				t.Errorf("Error converting mnemonic to integer: %v", err)
				return
			}
			hex, _ = big.NewInt(0).SetString(vector[0], 16)
			if decodedNum.Cmp(hex) != 0 {
				t.Errorf("Expected hex: %v, Got hex: %v", hex.Text(16), decodedNum.Text(16))
				return
			}
		})
	}
}

func TestFixedTruncation(t *testing.T) {
	t.Parallel()
	const test64BitHex = 0xEEEEEEEEFFFFFFFF
	const test32BitTruncatedHex = 0xEEEEEEEE
	hex := big.NewInt(0).SetUint64(test64BitHex)
	words, err := mnemonic.FromBigIntFixed(hex, 32)
	if err != nil {
		t.Errorf("Failed to create mnemonic from hex: %v", err)
		return
	}

	// 32 bit hex with right most bits truncated
	hexTruncaated := big.NewInt(0).SetUint64(test32BitTruncatedHex)
	truncatedWords, err := mnemonic.FromBigIntFixed(hexTruncaated, 32)
	if err != nil {
		t.Errorf("Failed to create mnemonic from hex: %v", err)
		return
	}
	for i, word := range words {
		if word != truncatedWords[i] {
			t.Errorf("manually truncated mnemonic does not match truncated mnemonic expected: %v\ngot:%v", words, truncatedWords)
			return
		}
	}
}

func TestNextNumberDivisibleBy32(t *testing.T) {
	t.Parallel()
	tests := []struct {
		num  int
		want int
	}{
		{0, 32},
		{1, 32},
		{31, 32},
		{32, 32},
		{33, 64},
		{63, 64},
		{64, 64},
		{65, 96},
		{127, 128},
		{128, 128},
		{129, 160},
		{255, 256},
		{256, 256},
	}
	for _, test := range tests {
		test := test
		t.Run(fmt.Sprintf("Number: %v", test.num), func(t *testing.T) {
			t.Parallel()
			got := mnemonic.NextNumberDivisibleBy32(test.num)
			if got != test.want {
				t.Errorf("Got: %v, Want: %v", got, test.want)
				return
			}
		})
	}
}

func TestInvalidWord(t *testing.T) {
	t.Parallel()
	words := []string{"DIMO", "world", "abandon"}
	_, err := mnemonic.ToBigInt(words)
	if !errors.Is(err, mnemonic.ErrInvalidWord) {
		t.Errorf("Expected error for invalid word, got %v", err)
		return
	}
}

func TestInvalidWordLength(t *testing.T) {
	t.Parallel()
	words := []string{"zoo"}
	_, err := mnemonic.ToBigInt(words)
	if !errors.Is(err, mnemonic.ErrInvalidWord) {
		t.Errorf("Expected error for invalid word, got %v", err)
		return
	}
}

// TestInvalidChecksum tests the case where the checksum is invalid
func TestInvalidChecksum(t *testing.T) {
	t.Parallel()
	words := []string{"sort", "labor", "armor", "various", "gown", "method", "public", "round", "deer", "leaf", "zoo", "zoo"}
	_, err := mnemonic.ToBigInt(words)
	if !errors.Is(err, mnemonic.ErrInvalidChecksum) {
		t.Errorf("Expected error for invalid checksum, got %v", err)
		return
	}
}

func TestInvalidBitSize(t *testing.T) {
	t.Parallel()
	ent := big.NewInt(0).SetUint64(0x1234567890)
	_, err := mnemonic.FromBigIntFixed(ent, 0)
	if !errors.Is(err, mnemonic.ErrInvalidBitSize) {
		t.Errorf("Expected error for invalid bit size, got %v", err)
		return
	}
	_, err = mnemonic.FromBigIntFixed(ent, 50)
	if !errors.Is(err, mnemonic.ErrInvalidBitSize) {
		t.Errorf("Expected error for invalid bit size, got %v", err)
		return
	}
}

func TestHugeNumber(t *testing.T) {
	t.Parallel()
	ent := big.NewInt(1)
	ent.Lsh(ent, 32*1000)
	ent.Sub(ent, big.NewInt(1))
	orig := new(big.Int).Set(ent)

	words := mnemonic.FromBigInt(ent)
	decoded, err := mnemonic.ToBigInt(words)
	if err != nil {
		t.Errorf("Error converting mnemonic to integer: %v", err)
		return
	}
	if decoded.Cmp(orig) != 0 {
		t.Errorf("Expected: %v, Got: %v", orig, decoded)
		return
	}
}
