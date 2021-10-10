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
		return nil
	} else {
		return handle
	}
}

func WriteToFile(handle *os.File, data []map[string]interface{}) {
	if common.UseCsv {
		// TODO: Implement CSV handling. This should take the top level JSON keys as columns,
		//       and then store for each element in the top level array the value as a flat string
		//writeToCsv(handle, data)
	} else {
		writeToJsonFile(handle, data)
	}
}

func CloseOutputFile(handle *os.File) {
	handle.Close()
}
