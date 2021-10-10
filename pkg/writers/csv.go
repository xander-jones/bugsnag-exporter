package writers

import (
	"fmt"
	"os"
	"reflect"

	"github.com/xander-jones/bugsnag-exporter/pkg/common"
)

func writeToCsvFile(handle *os.File, elements []map[string]interface{}) {
	// TODO: Implement CSV handling. This should take the top level JSON keys as columns,
	//       and then store for each element in the top level array the value as a flat string

	// log every header we've written so far so we can append to that column,
	// or create a new one if it does not yet exist
	// var csvHeadersWriten []string
	for _, element := range elements {
		v := reflect.ValueOf(element)
		common.Print(fmt.Sprint(element))
		common.Print(fmt.Sprint(v))
	}
}
