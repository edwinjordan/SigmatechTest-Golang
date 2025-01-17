package helpers

import (
	"encoding/base64"
	"errors"
	"os"
	"strings"
	// "bitbucket.org/yaayakk/rolling_glory/config"
)

/*
	fileExt : png,jpg,pdf
	extAllowed : png/ png|jpg
*/

func FileUploadFormat(fileExt string, extAllowed string) (bool, error) {

	allExt := ""
	formatAllowed := strings.Split(extAllowed, "|")

	for i, v := range formatAllowed {
		allExt += "." + v
		if i != len(formatAllowed)-1 {
			allExt += ","

		}

		if fileExt == v {
			return true, nil
		}
	}
	return false, errors.New("Hanya masukkan format " + allExt)
}

/*
	example

err := helpers.SaveFileFromBase64(file["FileName"].(string), file["Base64"].(string), "./uploaded_files/")
*/
func SaveFileFromBase64(fName string, data string, dir string) error {
	index := strings.Index(data, "base64,")
	dec, err := base64.StdEncoding.DecodeString(data[index+1:])
	PanicIfError(err)
	f, err := os.Create(dir + fName)
	PanicIfError(err)

	defer f.Close()
	if _, err := f.Write(dec); err != nil {
		PanicIfError(err)
	}
	f.Sync()
	return nil
}

func RemoveFile(fName string, dir string) error {
	os.Remove(dir + fName)
	return nil
}
