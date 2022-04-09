package fabric

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

var ccproot string = "/root/teamate/BS22_class-examples/teamate/application/ccp/"

func ConTest() {
	wallet, err := gateway.NewFileSystemWallet("wallet")
	if err != nil {
		fmt.Printf("Failed to create wallet: %s\n", err)
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
		gateway.WithIdentity(wallet, "appUser"),
	)
	if err != nil {
		fmt.Printf("Failed to connect to gateway: %s\n", err)
		os.Exit(1)
	}
	defer gw.Close()

}
