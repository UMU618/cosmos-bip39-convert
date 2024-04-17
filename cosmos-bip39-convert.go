package main // import "github.com/UMU618/cosmos-bip39-convert"

import (
	"encoding/hex"
	"flag"
	"fmt"

	"github.com/cosmos/cosmos-sdk/crypto/hd"

	sdk "github.com/cosmos/cosmos-sdk/types"
	bip39 "github.com/cosmos/go-bip39"
	//"github.com/tendermint/tendermint/crypto/secp256k1"
)

func main() {
	inputHex := flag.String("hex", "", "Hex string, [128, 256] bytes")
	inputMnemonic := flag.String("mnemonic", "", "Mnemonic, 12 or 24 words")
	flag.Parse()
	if len(*inputHex) > 0 {
		bytes, err := hex.DecodeString(*inputHex)
		if err != nil {
			fmt.Println(err)
			return
		}
		mnemonic, err := bip39.NewMnemonic(bytes)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(mnemonic)
	} else if len(*inputMnemonic) > 0 {
		mnemonic := *inputMnemonic
		bytes, err := bip39.MnemonicToByteArray(mnemonic)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Bytes     :", fmt.Sprintf("%x", bytes))

		seed, err := bip39.NewSeedWithErrorChecking(mnemonic, "")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Seed      :", fmt.Sprintf("%x", seed))

		masterPriv, ch := hd.ComputeMastersFromSeed(seed)
		derivedPriv, err := hd.DerivePrivateKeyForPath(masterPriv, ch, sdk.FullFundraiserPath)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Privatekey:", fmt.Sprintf("%x", derivedPriv))
		//fmt.Println("Publickey :", secp256k1.PrivKeySecp256k1(derivedPriv).PubKey())
	} else {
		flag.Usage()
		//flag.PrintDefaults()
	}
}
