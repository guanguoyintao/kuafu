package sftpool

import (
	"io"
	"os"

	"github.com/pkg/sftp"
)

type SftpFile struct {
  filePath string
	fp *sftp.File
}

// getfilepath
func (fp *SftpFile) GetFilePath() string {
  return fp.filePath
}

// get filestat
func (fp *SftpFile) Stat() (os.FileInfo, error) {
  return fp.fp.Stat()
}

// seek
func (fp *SftpFile) Seek(offset int64, whence int) (int64, error) {
  return fp.fp.Seek(offset, whence)
}

// ----------------------------------------------------  read

// read
func (fp *SftpFile) Read(p []byte) (n int, err error) {
  return fp.fp.Read(p)
}

// ReadAt
func (fp *SftpFile) ReadAt(p []byte, off int64) (n int, err error) {
  return fp.fp.ReadAt(p, off)
}

// WriteTo
func (f *SftpFile) WriteTo(w io.Writer) (n int64, err error) {
  return f.fp.WriteTo(w)
}

// ----------------------------------------------------  write

// Write
func (fp *SftpFile) Write(p []byte) (n int, err error) {
  return fp.fp.Write(p)
}

// WriteAt
func (fp *SftpFile) WriteAt(p []byte, off int64) (n int, err error) {
  return fp.fp.WriteAt(p, off)
}

// ReadFrom
func (fp *SftpFile) ReadFrom(r io.Reader) (n int64, err error) {
  return fp.fp.ReadFrom(r)
}

func (fp *SftpFile) close() error {
	return fp.fp.Close()
}

