package writers

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/xander-jones/bugsnag-exporter/pkg/common"
)

func writeToJsonFile(handle *os.File, data []map[string]interface{}) {
	file, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		common.ExitWithError(1, err)
	} else {
		err := ioutil.WriteFile(handle.Name(), file, 0644)
		if err != nil {
			common.ExitWithError(1, err)
		}
	}
}
