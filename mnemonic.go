// Package mnemonic encodes data into menemonic word list.
// encoding is implements the BIP-0039 specification using big.Int.
// This implementations allows for the use of arbitrary data sizes.
// This package also provides optional number obfuscation.
// Obfuscation uses Modular multiplicative inverse to convert the provided number to a seemingly random number before creating the mnemonic word list.
package mnemonic

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/c2h5oh/hide"
	"golang.org/x/exp/constraints"
)

const (
	// ObfuscationPrime is the prime number used to obfuscate the data.
	ObfuscationPrime = 2127482879
	hexBase          = 16
	bitsPerHexChar   = 4
)

var (
	bigPrime = new(big.Int).SetUint64(ObfuscationPrime)

	// Hider used for optional data obfuscattion.
	Hider = hide.Hide{}

	// ErrInvalidData is returned when the data is invalid.
	ErrInvalidData = errors.New("invalid data")
)

func init() {
	err := Hider.SetUint32(bigPrime)
	if err != nil {
		panic(fmt.Sprintf("Error setting uint32 prime: %v", err))
	}
	err = Hider.SetUint64(bigPrime)
	if err != nil {
		panic(fmt.Sprintf("Error setting uint64 prime: %v", err))
	}
	err = Hider.SetInt32(bigPrime)
	if err != nil {
		panic(fmt.Sprintf("Error setting int32 prime: %v", err))
	}
	err = Hider.SetInt64(bigPrime)
	if err != nil {
		panic(fmt.Sprintf("Error setting int64 prime: %v", err))
	}
}

// FromInt converts a signed integer to a list of mnemonic words.
func FromInt[T constraints.Signed](data T) []string {
	return FromBigInt(big.NewInt(int64(data)))
}

// FromUint converts an unsigned integer to a list of mnemonic words.
func FromUint[T constraints.Unsigned](data T) []string {
	return FromBigInt(new(big.Int).SetUint64(uint64(data)))
}

// FromHex converts a hex string to a list of mnemonic words.
func FromHex(data string) ([]string, error) {
	entBin, ok := big.NewInt(0).SetString(data, hexBase)
	if !ok {
		return nil, fmt.Errorf("%w: invalid hex %q", ErrInvalidData, data)
	}
	// encode any leading 0 in the hex string
	bitSize := NextNumberDivisibleBy32(len(data) * bitsPerHexChar)
	return FromBigIntFixed(entBin, bitSize)
}

// FromBytes converts a byte slice to a list of mnemonic words.
func FromBytes(data []byte) []string {
	return FromBigInt(big.NewInt(0).SetBytes(data))
}

// FromUint32WithObfuscation behaves the same as FromUint, but the provided data is obfuscated first.
func FromUint32WithObfuscation(data uint32) []string {
	obfuscatedData := Hider.Uint32Obfuscate(data)
	return FromUint(obfuscatedData)
}

// FromInt32WithObfuscation behaves the same as FromInt, but the provided data is obfuscated first.
func FromInt32WithObfuscation(data int32) []string {
	obfuscatedData := Hider.Int32Obfuscate(data)
	return FromInt(obfuscatedData)
}

// FromUint64WithObfuscation behaves the same as FromUint, but the provided data is obfuscated first.
func FromUint64WithObfuscation(data uint64) []string {
	obfuscatedData := Hider.Uint64Obfuscate(data)
	return FromUint(obfuscatedData)
}

// FromInt64WithObfuscation behaves the same as FromInt, but the provided data is obfuscated first.
func FromInt64WithObfuscation(data int64) []string {
	obfuscatedData := Hider.Int64Obfuscate(data)
	return FromInt(obfuscatedData)
}

// ToInt converts a list of mnemonic words to an int64.
func ToInt(words []string) (int64, error) {
	entBin, err := ToBigInt(words)
	if err != nil {
		return 0, err
	}
	return entBin.Int64(), nil
}

// ToUInt converts a list of mnemonic words to a uint64.
func ToUInt(words []string) (uint64, error) {
	entBin, err := ToBigInt(words)
	if err != nil {
		return 0, err
	}
	return entBin.Uint64(), nil
}

// ToHex converts a list of mnemonic words to a hex string.
func ToHex(words []string) (string, error) {
	entBin, err := ToBigInt(words)
	if err != nil {
		return "", err
	}
	return entBin.Text(hexBase), nil
}

// ToBytes converts a list of mnemonic words to a byte slice.
func ToBytes(words []string) ([]byte, error) {
	entBin, err := ToBigInt(words)
	if err != nil {
		return nil, err
	}
	return entBin.Bytes(), nil
}

// ToUint32WithDeobfuscation behaves the same as ToUInt, but the result is deobfuscated.
func ToUint32WithDeobfuscation(words []string) (uint32, error) {
	decodedData, err := ToUInt(words)
	if err != nil {
		return 0, err
	}
	return Hider.Uint32Deobfuscate(uint32(decodedData)), nil
}

// ToInt32WithDeobfuscation behaves the same as ToInt, but the result is deobfuscated.
func ToInt32WithDeobfuscation(words []string) (int32, error) {
	decodedData, err := ToInt(words)
	if err != nil {
		return 0, err
	}
	return Hider.Int32Deobfuscate(int32(decodedData)), nil
}

// ToUint64WithDeobfuscation behaves the same as ToUInt, but the result is deobfuscated.
func ToUint64WithDeobfuscation(words []string) (uint64, error) {
	decodedData, err := ToUInt(words)
	if err != nil {
		return 0, err
	}
	return Hider.Uint64Deobfuscate(decodedData), nil
}

// ToInt64WithDeobfuscation behaves the same as ToInt, but the result is deobfuscated.
func ToInt64WithDeobfuscation(words []string) (int64, error) {
	decodedData, err := ToInt(words)
	if err != nil {
		return 0, err
	}
	return Hider.Int64Deobfuscate(decodedData), nil
}
