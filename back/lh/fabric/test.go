package fabric

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/core"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
	"github.com/hyperledger/fabric-sdk-go/pkg/msp"
)

type testFixture struct {
	cryptoSuiteConfig core.CryptoSuiteConfig
	identityConfig    msp.IdentityConfig
}

func populateWallet(wallet *gateway.Wallet) error {
	credPath := filepath.Join(

		"/root/teamate/BS22_class-examples/teamate/application",
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

	keyDir := filepath.Join("/root/teamate/BS22_class-examples/teamate/application/msp/org1.example.com/users/keystore/")
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

	err = wallet.Put("admin", identity)
	if err != nil {
		return err
	}
	return nil
}
func ConcaTest() {

	os.Setenv("DISCOVERY_AS_LOCALHOST", "true")
	wallet, err := gateway.NewFileSystemWallet("wallet")
	if err != nil {
		fmt.Printf("Failed to create wallet: %s\n", err)
		os.Exit(1)
	}

	if !wallet.Exists("appUser") {
		err = populateWallet(wallet)
		if err != nil {
			fmt.Printf("Failed to populate wallet contents: %s\n", err)
			os.Exit(1)
		}
	}

	ccpPath := filepath.Join(
		"..",
		"..",
		"test-network",
		"organizations",
		"peerOrganizations",
		"org1.example.com",
		"connection-org1.yaml",
	)

	gw, err := gateway.Connect(
		gateway.WithConfig(config.FromFile(filepath.Clean(ccpPath))),
		gateway.WithIdentity(wallet, "appUser"),
	)
	if err != nil {
		fmt.Printf("Failed to connect to gateway: %s\n", err)
		os.Exit(1)
	}
	defer gw.Close()

	network, err := gw.GetNetwork("mychannel")
	if err != nil {
		fmt.Printf("Failed to get network: %s\n", err)
		os.Exit(1)
	}

	contract := network.GetContract("basic")

	result, err := contract.EvaluateTransaction("GetAllAssets")
	if err != nil {
		fmt.Printf("Failed to evaluate transaction: %s\n", err)
		os.Exit(1)
	}
	fmt.Println(string(result))
	/*
		result, err = contract.SubmitTransaction("createCar", "CAR10", "VW", "Polo", "Grey", "Mary")
		if err != nil {
			fmt.Printf("Failed to submit transaction: %s\n", err)
			os.Exit(1)
		}
		fmt.Println(string(result))

		result, err = contract.EvaluateTransaction("queryCar", "CAR10")
		if err != nil {
			fmt.Printf("Failed to evaluate transaction: %s\n", err)
			os.Exit(1)
		}
		fmt.Println(string(result))

		_, err = contract.SubmitTransaction("changeCarOwner", "CAR10", "Archie")
		if err != nil {
			fmt.Printf("Failed to submit transaction: %s\n", err)
			os.Exit(1)
		}

		result, err = contract.EvaluateTransaction("queryCar", "CAR10")
		if err != nil {
			fmt.Printf("Failed to evaluate transaction: %s\n", err)
			os.Exit(1)
		}
		fmt.Println(string(result))
	*/
	// Output: enroll user is completed
}

var ccproot string = "/root/teamate/BS22_class-examples/teamate/application/ccp/"

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
		gateway.WithIdentity(wallet, "admin"),
	)
	if err != nil {
		fmt.Printf("Failed to connect to gateway: %s\n", err)
		os.Exit(1)
	} else {
		fmt.Printf(" gateway: %v\n", gw)
	}
	defer gw.Close()

}
