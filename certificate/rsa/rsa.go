package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

// https://www.sohamkamani.com/golang/rsa-encryption/
func main() {
	rsaEncryption()
}

func rsaSigning() {

}

func rsaEncryption() {
	// the GenerateKey method takes in a reader that returns random bits, and
	// the number of bits
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	// the public key is a part of the *rsa.PrivateKey struct
	publicKey := privateKey.PublicKey

	// use the public and private keys
	// ...

	encryptedBytes, err := rsa.EncryptOAEP(
		// a hashing function, chosen so that even if the input is changed slightly,
		// the output hash changes completely. The SHA256 algorithm is suitable for this
		sha256.New(),
		// a random reader used for generating random bits so that the same input doesn’t give the same output twice
		rand.Reader,
		// the public key generated previously
		&publicKey,
		// the message we want to encrypt
		[]byte("super secret message"),
		// an optional label (which we will omit in this case)
		nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("encrypted bytes: ", encryptedBytes)

	// the first argument is an optional random data generator (the rand.Reader we used before)
	// we can set this value as nil
	// the OAEPOptions in the end signify that we encrypted the data using OAEP, and that we used
	// SHA256 to hash the input.
	decryptedBytes, err := privateKey.Decrypt(
		nil,
		// the encrypted data (called the cipher text)
		encryptedBytes,
		// the hash that we used to encrypt the data
		&rsa.OAEPOptions{Hash: crypto.SHA256})
	if err != nil {
		panic(err)
	}

	// we get back the original information in the form of bytes, which we
	// the cast to a string and print
	fmt.Println("decrypted message: ", string(decryptedBytes))

	// signing And Verification

	msg := []byte("verifiable message")

	// before signing, we need to hash our message
	// the hash is what we actually sign
	msgHash := sha256.New()
	_, err = msgHash.Write(msg)
	if err != nil {
		panic(err)
	}
	msgHashSum := msgHash.Sum(nil)

	// in order to generate the signature, we provide a random number generator,
	// our private key, the hashing algorithm that we used, and the hash sum
	// of our message
	signature, err := rsa.SignPSS(rand.Reader, privateKey, crypto.SHA256, msgHashSum, nil)
	if err != nil {
		panic(err)
	}

	// to verify the signature, we provide the public key, the hashing algorithm
	// the hash sum of our message and the signature we generated previously
	// there is an optional "options" parameter which can omit for now
	err = rsa.VerifyPSS(&publicKey, crypto.SHA256, msgHashSum, signature, nil)
	if err != nil {
		fmt.Println("could not verify signature: ", err)
		return
	}
	// if we don't get any error from the `VerifyPSS` method, that means our
	// signature is valid
	fmt.Println("signature verified")

}
