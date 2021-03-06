package ixo

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	ed25519 "github.com/tendermint/ed25519"
	crypto "github.com/tendermint/go-crypto"
)

func SignIxoMessage(msg sdk.Msg, did string, privKey [64]byte) IxoSignature {
	fmt.Println("*******SIGNING_MSG******* \n", string(msg.GetSignBytes()))

	signatureBytes := ed25519.Sign(&privKey, msg.GetSignBytes())
	fmt.Println(string(msg.GetSignBytes()))
	signature := crypto.SignatureEd25519(*signatureBytes).Wrap()

	return NewSignature(time.Now(), signature)
}

func VerifySignature(msg sdk.Msg, publicKey [32]byte, sig sdk.StdSignature) bool {
	fmt.Println("*******VERIFY_MSG******* \n", string(msg.GetSignBytes()))

	// First we unwrap the crypto.Signature to the crypto.SignatureEd25519 then we cast it to bytes
	innerSignature := sig.Signature.Unwrap().(crypto.SignatureEd25519)
	signatureBytes := [64]byte(innerSignature)
	result := ed25519.Verify(&publicKey, msg.GetSignBytes(), &signatureBytes)

	return result
}
