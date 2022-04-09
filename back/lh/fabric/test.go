package fabric

import (
	"fmt"

	mspclient "github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

func ConTest() {
	cfgProvider := config.FromFile("/root/teamate/BS22_class-examples/teamate/application/ccp/ccp-template.yaml")

	sdk, err := fabsdk.New(cfgProvider)
	if err != nil {
		fmt.Print(err, "\n")
		return
	}

	org1MspClient, err := mspclient.New(sdk.Context(), mspclient.WithOrg("org1"))
	if err != nil {
		return
	}
	fmt.Print(org1MspClient)

	org2MspClient, err := mspclient.New(sdk.Context(), mspclient.WithOrg("org2"))
	if err != nil {
		return
	}
	fmt.Print(org2MspClient)

}