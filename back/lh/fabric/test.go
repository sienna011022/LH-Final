package fabric

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

func populateWallet(wallet *gateway.Wallet) error {
	credPath := filepath.Join(

		"../",
		"msp",
		"org1.example.com",
		"users",
	)

	certPath := filepath.Join(credPath, "signcerts", "User1@org1.example.com-cert.pem")
	// read the certificate pem
	cert, err := ioutil.ReadFile(filepath.Clean(certPath))
	if err != nil {
		return err
	}

	keyDir := filepath.Join("../msp/org1.example.com/users/keystore/")
	// there's a single file in this dir containing the private key
	files, err := ioutil.ReadDir(keyDir)
	if err != nil {
		return err
	}
	if len(files) != 1 {
		return errors.New("keystore folder should have contain one file")
	}
	keyPath := filepath.Join(keyDir, files[0].Name())
	key, err := ioutil.ReadFile(filepath.Clean(keyPath))
	if err != nil {
		return err
	}

	identity := gateway.NewX509Identity("Org1MSP", string(cert), string(key))

	err = wallet.Put("lh", identity)
	if err != nil {
		return err
	}
	return nil
}

var ccproot string = "../ccp/"

func ConTest() {

	wallet, err := gateway.NewFileSystemWallet("wallet")
	if err != nil {
		fmt.Printf("Failed to create wallet: %s\n", err)
		os.Exit(1)
	}
	err = populateWallet(wallet)
	if err != nil {
		fmt.Printf("Failed to connect to gateway: %s\n", err)
		os.Exit(1)
	}
	// Path to the network config (CCP) file
	ccpPath := filepath.Join(
		ccproot,
		"connection-org1.yaml",
	)

	// Connect to the gateway peer(s) using the network config and identity in the wallet
	fmt.Printf("ccp : %s\n11", filepath.Clean(ccpPath))
	gw, err := gateway.Connect(
		gateway.WithConfig(config.FromFile(filepath.Clean(ccpPath))),
		gateway.WithIdentity(wallet, "lh"),
	)
	if err != nil {
		fmt.Printf("Failed to connect to gateway: %s\n", err)
		os.Exit(1)
	} else {
		fmt.Printf(" gateway: %v\n", gw)
	}
	defer gw.Close()

}
