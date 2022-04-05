package sdk

import "testing"

func TestNewFireworksSDK(t *testing.T) {
	sdk := NewFireworksSDK("http://127.0.0.1:8087")
	err := sdk.InitConfig("C66SVP", "./configs")
	if err != nil {
		panic(err)
	}
}
