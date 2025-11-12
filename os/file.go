package eos

import (
	"bytes"
	"io"
	"os"
)

func CheckFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func GetFile(filePath string) (fd *os.File, err error) {
	if CheckFileIsExist(filePath) {
		fd, err = os.OpenFile(filePath, os.O_APPEND|os.O_RDWR, os.ModeAppend)
		if err != nil {
			return fd, err
		}
	} else {
		fd, err = os.Create(filePath)
		if err != nil {
			return fd, err
		}
	}

	return fd, err
}

func Mkdir(dirPath string) (err error) {
	if CheckFileIsExist(dirPath) {
		return nil
	}
	err = os.MkdirAll(dirPath, 0766)
	if err != nil {
		return err
	}

	return nil
}

func LineCounter(r io.Reader) (int, error) {
	var readSize int
	var err error
	var count int
	buf := make([]byte, 1024)
	for {
		readSize, err = r.Read(buf)
		if err != nil {
			break
		}
		var buffPosition int
		for {
			i := bytes.IndexByte(buf[buffPosition:], '\n')
			if i == -1 || readSize == buffPosition {
				break
			}
			buffPosition += i + 1
			count++
		}
	}
	if readSize > 0 && count == 0 || count > 0 {
		count++
	}
	if err == io.EOF {
		return count, nil
	}

	return count, err
}
