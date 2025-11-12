package ebuffer

import "os"

// TempFile 是一个临时文件，实现了 io.Closer 接口，可以在 Close 方法中删除文件。
type TempFile struct {
	*os.File
	path string
}

// NewTempFile 创建一个新的临时文件。
func NewTempFile(tempFileName string) (*TempFile, error) {
	//file, err := os.CreateTemp(dir, pattern)
	//if err != nil {
	//	return nil, err
	//}
	file, err := os.OpenFile(tempFileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		os.Remove(tempFileName)
		return nil, err
	}
	return &TempFile{File: file, path: file.Name()}, nil
}

// Close 关闭临时文件并删除它。
func (t *TempFile) Close() error {
	err := t.File.Close()
	if err != nil {
		return err
	}
	return os.Remove(t.path)
}

// Len 返回临时文件的长度。
func (t *TempFile) Len() (int64, error) {
	fileInfo, err := t.File.Stat()
	if err != nil {
		return 0, err
	}
	return fileInfo.Size(), nil
}
