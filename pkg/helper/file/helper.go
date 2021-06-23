package file

import "os"

// openFile open the log file
func openFile(filePath string, fileName string) (file *os.File, err error) {
	if !isExist(filePath) {
		err = os.MkdirAll(filePath, os.ModePerm)
		if err != nil {
			return
		}
	}
	name := filePath + "/" + fileName
	if isExist(name) {
		file, err = os.OpenFile(name, os.O_APPEND|os.O_RDWR, 0666)
	} else {
		file, err = os.Create(name)
	}
	if err != nil {
		return
	}
	return
}

// isExist the file if exist return true else false
func isExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
