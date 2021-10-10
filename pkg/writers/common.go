package writers

import (
	"os"
	"path/filepath"

	"github.com/xander-jones/bugsnag-exporter/pkg/common"
)

func CreateNewOutputFile(filename string) *os.File {
	var handle *os.File
	var outputFilepath string
	if common.UseCsv {
		outputFilepath = filepath.Join(common.OutputDir, filename+".csv")
	} else {
		outputFilepath = filepath.Join(common.OutputDir, filename+".json")
	}
	common.PrintVerbose("Creating file " + outputFilepath)
	handle, err := os.Create(outputFilepath)
	if err != nil {
		common.ExitWithError(1, err)
	} else {
		return handle
	}
	return nil
}

func WriteObjectToFile(handle *os.File, element map[string]interface{}) {
	common.PrintVerbose("Writing object to output file '" + handle.Name() + "'")
	if common.UseCsv {
		//writeObjectToCsvFile(handle, element)
	} else {
		writeObjectToJsonFile(handle, element)
	}
}

func WriteArrayToFile(handle *os.File, elements []map[string]interface{}) {
	common.PrintVerbose("Writing array to output file '" + handle.Name() + "'")
	if common.UseCsv {
		//writeArrayToCsvFile(handle, elements)
	} else {
		writeArrayToJsonFile(handle, elements)
	}
}

func CloseOutputFile(handle *os.File) {
	common.PrintVerbose("Closing output file '" + handle.Name() + "'")
	handle.Close()
}
