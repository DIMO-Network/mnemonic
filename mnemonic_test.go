package mnemonic_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/KevinJoiner/mnemonic"
)

var testVectors = trezorVectors{}

type trezorVectors struct {
	English [][]string `json:"english"`
}

func TestMain(m *testing.M) {
	err := json.Unmarshal(trezorTestVectors, &testVectors)
	if err != nil {
		panic(fmt.Sprintf("Error unmarshalling test vectors: %v", err))
	}
	m.Run()
}

func TestFromHex(t *testing.T) {
	t.Parallel()
	hex := "123456789abcdef"
	words, err := mnemonic.FromHex(hex)
	if err != nil {
		t.Errorf("Error converting hex to mnemonic: %v", err)
		return
	}
	decodedHex, err := mnemonic.ToHex(words)
	if err != nil {
		t.Errorf("Error converting mnemonic to hex: %v", err)
		return
	}
	if hex != decodedHex {
		t.Errorf("Original hex: %v, Decoded hex: %v", hex, decodedHex)
	}
}

func TestFromBytes(t *testing.T) {
	t.Parallel()
	bytes := []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef}
	words := mnemonic.FromBytes(bytes)
	decodedBytes, err := mnemonic.ToBytes(words)
	if err != nil {
		t.Errorf("Error converting mnemonic to bytes: %v", err)
		return
	}
	if string(bytes) != string(decodedBytes) {
		t.Errorf("Original bytes: %v, Decoded bytes: %v", bytes, decodedBytes)
	}
}

func TestInteger(t *testing.T) {
	t.Parallel()
	num := uint64(0x01234567ab)
	words := mnemonic.FromUint(num)
	decodedNum, err := mnemonic.ToUInt(words)
	if err != nil {
		t.Errorf("Error converting mnemonic to uint: %v", err)
		return
	}
	if num != decodedNum {
		t.Errorf("Original number: %v, Decoded uint64 number: %v", num, decodedNum)
	}

	num2 := int64(0x01234567ab)
	words2 := mnemonic.FromInt(num2)
	decodedNum2, err := mnemonic.ToInt(words2)
	if err != nil {
		t.Errorf("Error converting mnemonic to int: %v", err)
		return
	}
	if num2 != decodedNum2 {
		t.Errorf("Original number: %v, Decoded in64 number: %v", num2, decodedNum2)
	}
}

func TestUint64Obfuscation(t *testing.T) {
	t.Parallel()
	num := uint64(0x01234567ab)
	words := mnemonic.FromUint64WithObfuscation(num)
	decodedNum, err := mnemonic.ToUInt(words)
	if err != nil {
		t.Errorf("Error converting mnemonic to uint: %v", err)
		return
	}
	if num == decodedNum {
		t.Errorf("Original number was not obfuscated, original number equals decoded number: %v", num)
	}
	decodedNum, err = mnemonic.ToUint64WithDeobfuscation(words)
	if err != nil {
		t.Errorf("Error converting mnemonic to uint: %v", err)
		return
	}
	if num != decodedNum {
		t.Errorf("Original number: %v, Decoded uint64 number: %v", num, decodedNum)
	}
}

func TestInt64Obfuscation(t *testing.T) {
	t.Parallel()
	num := int64(0x01234567ab)
	words := mnemonic.FromInt64WithObfuscation(num)
	decodedNum, err := mnemonic.ToInt(words)
	if err != nil {
		t.Errorf("Error converting mnemonic to int: %v", err)
		return
	}
	if num == decodedNum {
		t.Errorf("Original number was not obfuscated, original number equals decoded number: %v", num)
	}
	decodedNum, err = mnemonic.ToInt64WithDeobfuscation(words)
	if err != nil {
		t.Errorf("Error converting mnemonic to int: %v", err)
		return
	}
	if num != decodedNum {
		t.Errorf("Original number: %v, Decoded int64 number: %v", num, decodedNum)
	}
}

func TestUint32Obfuscation(t *testing.T) {
	t.Parallel()
	num := uint32(0x01234567)
	words := mnemonic.FromUint32WithObfuscation(num)
	decodedNum, err := mnemonic.ToUInt(words)
	if err != nil {
		t.Errorf("Error converting mnemonic to uint: %v", err)
		return
	}
	if num == uint32(decodedNum) {
		t.Errorf("Original number was not obfuscated, original number equals decoded number: %v", num)
	}
	decodedNum2, err := mnemonic.ToUint32WithDeobfuscation(words)
	if err != nil {
		t.Errorf("Error converting mnemonic to uint: %v", err)
		return
	}
	if num != decodedNum2 {
		t.Errorf("Original number: %v, Decoded uint32 number: %v", num, decodedNum)
	}
}

func TestInt32Obfuscation(t *testing.T) {
	t.Parallel()
	num := int32(0x01234567)
	words := mnemonic.FromInt32WithObfuscation(num)
	decodedNum, err := mnemonic.ToInt(words)
	if err != nil {
		t.Errorf("Error converting mnemonic to int: %v", err)
		return
	}
	if num == int32(decodedNum) {
		t.Errorf("Original number was not obfuscated, original number equals decoded number: %v", num)
	}
	decodedNum2, err := mnemonic.ToInt32WithDeobfuscation(words)
	if err != nil {
		t.Errorf("Error converting mnemonic to int: %v", err)
		return
	}
	if num != decodedNum2 {
		t.Errorf("Original number: %v, Decoded int32 number: %v", num, decodedNum)
	}
}
