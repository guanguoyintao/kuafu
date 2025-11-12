package sftpool

import (
	"io/fs"

	"github.com/pkg/sftp"
)

type SftpSession struct {
	name string
	client *sftp.Client
}

func (s *SftpSession) OpenFile(path string, osFlag int) (*SftpFile, error) {
	sftpFp, err := s.client.OpenFile(path, osFlag)
	if err != nil {
		return nil, err
	}
	return &SftpFile{filePath: path, fp: sftpFp}, nil
}

func (s *SftpSession) LstatFile(path string) (fs.FileInfo, error) {
	return s.client.Lstat(path)
}

func (s *SftpSession) RemoveFile(path string) error {
	return s.client.Remove(path)
}