package sdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/dollarkillerx/fireworks/internal/response"
	"github.com/dollarkillerx/urllib"
)

type FireworksSDK struct {
	url string
}

func NewFireworksSDK(url string) *FireworksSDK {
	return &FireworksSDK{
		url: url,
	}
}

func (f *FireworksSDK) InitConfig(token string, pat string) error {
	os.MkdirAll(pat, 00766)

	var resp response.Configurations
	httpcode, respbody, err := urllib.Get(fmt.Sprintf("%s/api/v1/configurations", f.url)).Queries("configuration_token", token).Byte()
	if err != nil {
		log.Println(err)
		return err
	}
	if httpcode != 200 {
		return errors.New(string(respbody))
	}

	err = json.Unmarshal(respbody, &resp)
	if err != nil {
		log.Println(err)
		return err
	}

	for _, v := range resp.Configs {
		newPath := path.Join(pat, v.Filename)
		err := ioutil.WriteFile(newPath, []byte(v.Body), 00766)
		if err != nil {
			return err
		}
	}

	return nil
}
