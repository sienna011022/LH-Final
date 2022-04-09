package fabric

import (
	"fmt"

	mspclient "github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

func ConTest() {
	cfgProvider := config.FromFile("/root/teamate/BS22_class-examples/teamte/application/ccp/ccp-template.yaml")

	sdk, err := fabsdk.New(cfgProvider)
	if err != nil {
		return
	}

	org1MspClient, err := mspclient.New(sdk.Context(), mspclient.WithOrg("Org1"))
	if err != nil {
		return
	}
	fmt.Print(org1MspClient)

	org2MspClient, err := mspclient.New(sdk.Context(), mspclient.WithOrg("Org2"))
	if err != nil {
		return
	}
	fmt.Print(org2MspClient)

}
