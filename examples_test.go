package mnemonic_test

import (
	"fmt"

	"github.com/DIMO-Network/mnemonic"
)

func ExampleFromInt() {
	// Generate a mnemonic from an integer
	words := mnemonic.FromInt(1234567890)

	fmt.Println(words)
	// Output: [end quote region]
}

func ExampleToInt() {
	// Generate an integer from a mnemonic
	words := []string{"end", "quote", "region"}
	number, err := mnemonic.ToInt(words)
	if err != nil {
		panic(err)
	}

	fmt.Println(number)
	// Output: 1234567890
}

func ExampleFromUint() {
	// Generate a mnemonic from an unsigned integer
	unsignedInt := uint(1234567890)
	words := mnemonic.FromUint(unsignedInt)

	fmt.Println(words)
	// Output: [end quote region]
}

func ExampleToUint() {
	// Generate an unsigned integer from a mnemonic
	words := []string{"end", "quote", "region"}
	unsignedInt, err := mnemonic.ToUint(words)
	if err != nil {
		panic(err)
	}

	_, _ = fmt.Printf("Value: %d Type: %T\n", unsignedInt, unsignedInt)
	// Output: Value: 1234567890 Type: uint64
}

func ExampleFromHex() {
	// Generate a mnemonic from a hex string
	hex := "023456789abcdef0"
	words, err := mnemonic.FromHex(hex)
	if err != nil {
		panic(err)
	}

	fmt.Println(words)
	// Output: [acquire pencil owner cube social journey]
}

func ExampleToHex() {
	// Generate a hex string from a mnemonic
	words := []string{"acquire", "pencil", "owner", "cube", "social", "journey"}
	hex, err := mnemonic.ToHex(words)
	if err != nil {
		panic(err)
	}

	fmt.Println(hex)
	// Output: 023456789abcdef0
}

func ExampleFromBytes() {
	// Generate a mnemonic from a byte slice
	bytes := []byte{'z', 35, 67, 0xff, 0x89, 0, 0xcd, 0xef}
	words, err := mnemonic.FromBytes(bytes)
	if err != nil {
		panic(err)
	}

	fmt.Println(words)
	// Output: [kick borrow zoo bamboo art wasp]
}

func ExampleToBytes() {
	// Generate a byte slice from a mnemonic
	words := []string{"kick", "borrow", "zoo", "bamboo", "art", "wasp"}
	bytes, err := mnemonic.ToBytes(words)
	if err != nil {
		panic(err)
	}

	fmt.Println(bytes)
	// Output: [122 35 67 255 137 0 205 239]
}

func ExampleFromUint32WithObfuscation() {
	// Generate a mnemonic from an obfuscated unsigned integer
	obfuscatedUint := uint32(1)
	words := mnemonic.FromUint(obfuscatedUint)
	obfuscatedWords := mnemonic.FromUint32WithObfuscation(obfuscatedUint)

	fmt.Println(words)
	fmt.Println(obfuscatedWords)
	// Output: [abandon abandon about]
	// [learn island zoo]
}

func ExampleToUint32WithDeobfuscation() {
	// Generate an unsigned integer from an obfuscated mnemonic
	obfuscatedWords := []string{"learn", "island", "zoo"}
	deobfuscatedUint, err := mnemonic.ToUint32WithDeobfuscation(obfuscatedWords)
	if err != nil {
		panic(err)
	}

	fmt.Println(deobfuscatedUint)
	// Output: 1
}

func ExampleFromInt32WithObfuscation() {
	// Generate a mnemonic from an obfuscated integer
	obfuscatedInt := int32(1)
	words := mnemonic.FromInt(obfuscatedInt)
	obfuscatedWords := mnemonic.FromInt32WithObfuscation(obfuscatedInt)
	fmt.Println(words)
	fmt.Println(obfuscatedWords)
	// Output: [abandon abandon about]
	// [learn island zoo]
}

func ExampleToInt32WithDeobfuscation() {
	// Generate an integer from an obfuscated mnemonic
	obfuscatedWords := []string{"learn", "island", "zoo"}
	deobfuscatedInt, err := mnemonic.ToInt32WithDeobfuscation(obfuscatedWords)
	if err != nil {
		panic(err)
	}

	fmt.Println(deobfuscatedInt)
	// Output: 1
}

func ExampleFromUint64WithObfuscation() {
	// Generate a mnemonic from an obfuscated unsigned integer
	obfuscatedUint := uint64(1)
	words := mnemonic.FromUint(obfuscatedUint)
	obfuscatedWords := mnemonic.FromUint64WithObfuscation(obfuscatedUint)

	fmt.Println(words)
	fmt.Println(obfuscatedWords)
	// Output: [abandon abandon about]
	// [learn island zoo]
}

func ExampleToUint64WithDeobfuscation() {
	// Generate an unsigned integer from an obfuscated mnemonic
	obfuscatedWords := []string{"learn", "island", "zoo"}
	deobfuscatedUint, err := mnemonic.ToUint64WithDeobfuscation(obfuscatedWords)
	if err != nil {
		panic(err)
	}

	fmt.Println(deobfuscatedUint)
	// Output: 1
}

func ExampleFromInt64WithObfuscation() {
	// Generate a mnemonic from an obfuscated integer
	obfuscatedInt := int64(1)
	words := mnemonic.FromInt(obfuscatedInt)
	obfuscatedWords := mnemonic.FromInt64WithObfuscation(obfuscatedInt)

	fmt.Println(words)
	fmt.Println(obfuscatedWords)
	// Output: [abandon abandon about]
	// [learn island zoo]
}

func ExampleToInt64WithDeobfuscation() {
	// Generate an integer from an obfuscated mnemonic
	obfuscatedWords := []string{"learn", "island", "zoo"}
	deobfuscatedInt, err := mnemonic.ToInt64WithDeobfuscation(obfuscatedWords)
	if err != nil {
		panic(err)
	}

	fmt.Println(deobfuscatedInt)
	// Output: 1
}
