package filelisting

import (
	"net/http"
	"os"
	"io/ioutil"
	"strings"
)

const prefix  = "/list/"

type userError string

func (u userError)Error() string {
	return u.Message()
}

func (u userError)Message() string {
	return string(u)
}

func HanleFileList(writer http.ResponseWriter,
	request *http.Request) error {
	if strings.Index(request.URL.Path, prefix) != 0 {
		return userError("path must start " +
			"with " + prefix)
	}

	path := request.URL.Path[len(prefix):]

	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return err
	}

	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	writer.Write(all)

	return nil
}