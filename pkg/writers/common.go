package writers

import (
	"os"
	"path/filepath"

	"github.com/xander-jones/bugsnag-exporter/pkg/common"
)

func CreateNewOutputFile(projectId string, filename string) *os.File {
	// TODO: Warn user if the file is going to overwrite an old file
	var outputDirpath string = filepath.Join(common.OutputDir, projectId)
	common.PrintVerbose("Setting up directory '" + outputDirpath + "' to download data to")
	// TODO: Check MkdirAll is not destructive where not wanted!
	err := os.MkdirAll(outputDirpath, os.ModePerm)
	if err != nil {
		common.ExitWithError(1, err)
	}
	var handle *os.File
	var outputFilepath string
	if common.UseCsv {
		outputFilepath = filepath.Join(common.OutputDir, projectId, filename+".csv")
	} else {
		outputFilepath = filepath.Join(common.OutputDir, projectId, filename+".json")
	}
	common.PrintVerbose("Creating file " + outputFilepath)
	handle, err = os.Create(outputFilepath)
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
		// TODO: implement this
		//writeObjectToCsvFile(handle, element)
	} else {
		writeObjectToJsonFile(handle, element)
	}
}

func WriteArrayToFile(handle *os.File, elements []map[string]interface{}) {
	common.PrintVerbose("Writing array to output file '" + handle.Name() + "'")
	if common.UseCsv {
		// TODO: implement this
		//writeArrayToCsvFile(handle, elements)
	} else {
		writeArrayToJsonFile(handle, elements)
	}
}

func CloseOutputFile(handle *os.File) {
	common.PrintVerbose("Closing output file '" + handle.Name() + "'")
	handle.Close()
}
