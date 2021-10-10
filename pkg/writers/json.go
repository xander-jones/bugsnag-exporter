package writers

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/xander-jones/bugsnag-exporter/pkg/common"
)

func writeArrayToJsonFile(handle *os.File, elements []map[string]interface{}) {
	data, err := json.MarshalIndent(elements, "", "    ")
	if err != nil {
		common.ExitWithError(1, err)
	} else {
		err := ioutil.WriteFile(handle.Name(), data, 0644)
		if err != nil {
			common.ExitWithError(1, err)
		}
	}
}

func writeObjectToJsonFile(handle *os.File, element map[string]interface{}) {
	data, err := json.MarshalIndent(element, "", "    ")
	if err != nil {
		common.ExitWithError(1, err)
	} else {
		err := ioutil.WriteFile(handle.Name(), data, 0644)
		if err != nil {
			common.ExitWithError(1, err)
		}
	}
}
