package fabric

import (
	"fmt"

	mspclient "github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

func ConTest() {
	cfgProvider1 := config.FromFile("/root/teamate/BS22_class-examples/teamate/application/ccp/connection-org1.yaml")

	sdk1, err := fabsdk.New(cfgProvider1)
	if err != nil {
		fmt.Print(err)
		return
	}

	org1MspClient, err := mspclient.New(sdk1.Context(), mspclient.WithOrg("Org1"))
	if err != nil {
		return
	}
	fmt.Print(org1MspClient)
	cfgProvider2 := config.FromFile("/root/teamate/BS22_class-examples/teamate/application/ccp/connection-org2.yaml")

	sdk2, err := fabsdk.New(cfgProvider2)
	if err != nil {
		fmt.Print(err)
		return
	}
	org2MspClient, err := mspclient.New(sdk2.Context(), mspclient.WithOrg("Org2"))
	if err != nil {
		return
	}
	fmt.Print(org2MspClient)

}
