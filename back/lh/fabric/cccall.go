package fabric

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

const org1 = "org1"
const org2 = "org2"
const channel_right = "mychannl1"
const channel_contract = "mychannl2"
const CC = "CC_NAME"

func GetCC(org, iden_label, channel, CCName string) *gateway.Contract {
	ccpPath := filepath.Join(
		ccproot,
		fmt.Sprintf("connection-%s.yaml", org),
	)
	gw, err := gateway.Connect(
		gateway.WithConfig(config.FromFile(filepath.Clean(ccpPath))),
		gateway.WithIdentity(wallet, iden_label),
	)
	if err != nil {
		fmt.Printf("Failed to connect to gateway: %s\n", err)
		os.Exit(1)
	}
	defer gw.Close()
	fmt.Print(gw)
	network, err := gw.GetNetwork(channel)
	if err != nil {
		fmt.Printf("Failed to get network: %s\n", err)
		os.Exit(1)
	}

	contract := network.GetContract(CCName)

	return contract
}

type Ret struct {
	Result string `json:"result"`
}

func InitUser(key, cert1, cert2 string) error {
	var Init_ret Ret
	result, err := GetCC(org1, cert1, channel_right, CC).SubmitTransaction("AddUser", key)
	if err != nil {
		return fmt.Errorf("CC fail[%s]", err.Error())
	}
	err = json.Unmarshal(result, &Init_ret)
	if Init_ret.Result != "succes" || err != nil {
		return fmt.Errorf("AddUser fail[%s]", err.Error())
	}
	result, err = GetCC(org1, cert1, channel_contract, CC).SubmitTransaction("AddUser", key)
	if err != nil {
		return fmt.Errorf("CC fail[%s]", err.Error())
	}
	err = json.Unmarshal(result, &Init_ret)
	if Init_ret.Result != "succes" || err != nil {
		return fmt.Errorf("AddUser fail[%s]", err.Error())
	}
	result, err = GetCC(org2, cert2, channel_right, CC).SubmitTransaction("AddUser", key)
	if err != nil {
		return fmt.Errorf("CC fail[%s]", err.Error())
	}
	err = json.Unmarshal(result, &Init_ret)
	if Init_ret.Result != "succes" || err != nil {
		return fmt.Errorf("AddUser fail[%s]", err.Error())
	}
	result, err = GetCC(org2, cert2, channel_contract, CC).SubmitTransaction("AddUser", key)
	if err != nil {
		return fmt.Errorf("CC fail[%s]", err.Error())
	}
	err = json.Unmarshal(result, &Init_ret)
	if Init_ret.Result != "succes" || err != nil {
		return fmt.Errorf("AddUser fail[%s]", err.Error())
	}

	return nil
}

func UpdateRightDoc(key, name, docu_name, docu_id, hash, cert string) error {
	var Update_ret Ret
	result, err := GetCC(org1, cert, channel_right, CC).SubmitTransaction("UpdateContract", key, name, docu_id, docu_name, hash)
	if err != nil {
		return fmt.Errorf("CC fail[%s]", err.Error())
	}
	err = json.Unmarshal(result, &Update_ret)
	if Update_ret.Result != "succes" || err != nil {
		return fmt.Errorf("AddUser fail[%s]", err.Error())
	}

	return nil
}
func UpdateRightState(key, status, cert string) error {
	var Update_ret Ret
	result, err := GetCC(org2, cert, channel_right, CC).SubmitTransaction("UpdateState", key, status)
	if err != nil {
		return fmt.Errorf("CC fail[%s]", err.Error())
	}
	err = json.Unmarshal(result, &Update_ret)
	if Update_ret.Result != "succes" || err != nil {
		return fmt.Errorf("AddUser fail[%s]", err.Error())
	}

	return nil
}

func GetRightState(key, cert string) (string, error) {
	var Get_ret Ret
	result, err := GetCC(org1, cert, channel_right, CC).EvaluateTransaction("RequestState", key)
	if err != nil {
		return "", fmt.Errorf("CC fail[%s]", err.Error())
	}
	err = json.Unmarshal(result, &Get_ret)
	if err != nil {
		return "", fmt.Errorf("AddUser fail[%s]", err.Error())
	}

	return Get_ret.Result, nil
}

func GetRightDoc(key, cert string) (string, error) {
	/*
		var Get_ret Ret
		result, err := GetCC(org1, cert, channel_right, CC).EvaluateTransaction("ReadContract", key)
		if err != nil {
			return "", fmt.Errorf("CC fail[%s]", err.Error())
		}

			err = json.Unmarshal(result, &Get_ret)
			if err != nil {
				return "", fmt.Errorf("AddUser fail[%s]", err.Error())
			}
	*/
	return "", nil
}

func UpdateContactDoc(key, name, docu_name, contract_id, hash, cert string) error {
	var Update_ret Ret
	result, err := GetCC(org1, cert, channel_contract, CC).SubmitTransaction("UpdateContract", key, name, contract_id, docu_name, hash)
	if err != nil {
		return fmt.Errorf("CC fail[%s]", err.Error())
	}
	err = json.Unmarshal(result, &Update_ret)
	if Update_ret.Result != "succes" || err != nil {
		return fmt.Errorf("AddUser fail[%s]", err.Error())
	}

	return nil
}

func GetContractDoc(key, name, docu_name, contract_id, hash, cert string) (string, error) {
	/*
		var Get_ret Ret
		result, err := GetCC(org1, cert, channel_right, CC).EvaluateTransaction("ReadContract", key)
		if err != nil {
			return "", fmt.Errorf("CC fail[%s]", err.Error())
		}

			err = json.Unmarshal(result, &Get_ret)
			if err != nil {
				return "", fmt.Errorf("AddUser fail[%s]", err.Error())
			}
	*/
	return "", nil
}