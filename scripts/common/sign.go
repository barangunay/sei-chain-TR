package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	clienttx "github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	xauthsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

const NODE_URI = "tcp://localhost:26657"

type AccountInfo struct {
	Address  string `json:"address"`
	Mnemonic string `json:"mnemonic"`
}

func GetKey(keyname string) cryptotypes.PrivKey {
	userHomeDir, _ := os.UserHomeDir()
	accountKeyFilePath := filepath.Join(userHomeDir, "test_accounts", fmt.Sprintf("%s.json", keyname))
	jsonFile, err := os.Open(accountKeyFilePath)
	if err != nil {
		panic(err)
	}
	var accountInfo AccountInfo
	byteVal, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		panic(err)
	}
	jsonFile.Close()
	json.Unmarshal(byteVal, &accountInfo)
	kr, _ := keyring.New(sdk.KeyringServiceName(), "os", filepath.Join(userHomeDir, ".sei-chain"), os.Stdin)
	keyringAlgos, _ := kr.SupportedAlgorithms()
	algoStr := string(hd.Secp256k1Type)
	algo, _ := keyring.NewSigningAlgoFromString(algoStr, keyringAlgos)
	hdpath := hd.CreateHDPath(sdk.GetConfig().GetCoinType(), 0, 0).String()
	derivedPriv, _ := algo.Derive()(accountInfo.Mnemonic, "", hdpath)
	return algo.Generate()(derivedPriv)
}

func SignTx(txBuilder *client.TxBuilder, privKey cryptotypes.PrivKey, sequenceNum uint64) {
	var sigsV2 []signing.SignatureV2
	sigV2 := signing.SignatureV2{
		PubKey: privKey.PubKey(),
		Data: &signing.SingleSignatureData{
			SignMode:  TEST_CONFIG.TxConfig.SignModeHandler().DefaultMode(),
			Signature: nil,
		},
		Sequence: sequenceNum,
	}
	sigsV2 = append(sigsV2, sigV2)
	_ = (*txBuilder).SetSignatures(sigsV2...)
	sigsV2 = []signing.SignatureV2{}
	signerData := xauthsigning.SignerData{
		ChainID:       CHAIN_ID,
		AccountNumber: GetAccountNumber(privKey),
		Sequence:      sequenceNum,
	}
	sigV2, _ = clienttx.SignWithPrivKey(
		TEST_CONFIG.TxConfig.SignModeHandler().DefaultMode(),
		signerData,
		*txBuilder,
		privKey,
		TEST_CONFIG.TxConfig,
		sequenceNum,
	)
	sigsV2 = append(sigsV2, sigV2)
	_ = (*txBuilder).SetSignatures(sigsV2...)
}

func GetAccountNumber(privKey cryptotypes.PrivKey) uint64 {
	hexAccount := privKey.PubKey().Address()
	address, err := sdk.AccAddressFromHex(hexAccount.String())
	if err != nil {
		panic(err)
	}
	accountRetriever := authtypes.AccountRetriever{}
	cl, err := client.NewClientFromNode(NODE_URI)
	if err != nil {
		panic(err)
	}
	context := client.Context{}
	context = context.WithNodeURI(NODE_URI)
	context = context.WithClient(cl)
	context = context.WithInterfaceRegistry(TEST_CONFIG.InterfaceRegistry)
	account, err := accountRetriever.GetAccount(context, address)
	if err != nil {
		time.Sleep(5 * time.Second)
		// retry once after 5 seconds
		account, err = accountRetriever.GetAccount(context, address)
		if err != nil {
			panic(err)
		}
	}
	return account.GetAccountNumber()
}