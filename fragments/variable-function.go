package fragments

import "os"

func dummy()  {
	// define

	createFile := func (path string) *os.File {
		tmpfn, err := os.Create(path)
		if err != nil {
			"Error while creating file: ", path
		}
		return tmpfn
	}

	// invoke file creation

	createFile()
}