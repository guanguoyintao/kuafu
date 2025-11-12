package sftpool

import (
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"sync/atomic"
	"time"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

type Conn struct {
	connName string
	usedAt   int64 // atomic

	// hostAddr string
	// sshSessions map[string]*ssh.Session
	conn    net.Conn
	sshConn *ssh.Client

	sshClientConfig *ssh.ClientConfig
	sessionNames    []string
	// sftpSessions map[string]*SftpSession

	idelSftpSessionsChan chan *SftpSession
	// sftpSessions chan *SftpSession // this channel is used for MaxSessions limiting

	// Inited      bool
	maxSessions int
	pooled      bool
	createdAt   time.Time

	atomicOpeningFiles int32
}

// func NewConn(addr string, config *ssh.ClientConfig, maxSessions int) (*Conn, error) {
func NewConn(conn net.Conn, connName string, config *ssh.ClientConfig) (*Conn, error) {
	// sshConn , err := ssh.Dial("tcp", addr, config)
	// if err != nil {
	//   return nil, err
	// }

	// log.Printf("conn: +%v", conn)
	// log.Printf("conn.RemoteAddr(): +%v\n", conn.RemoteAddr())
	if conn.RemoteAddr() == nil || config == nil {
		log.Println("conn.RemoteAddr().String() is empty")
		return &Conn{
			connName: connName,
			// hostAddr: addr,
			// sshSessions: make(map[string]*ssh.Session),
			conn:            conn,
			sshConn:         nil,
			sshClientConfig: config,
			sessionNames:    nil,
			// sftpSessions: nil,
			idelSftpSessionsChan: nil,
			maxSessions:          0,
			createdAt:            time.Now(),
		}, nil
	}

	c, chans, reqs, err := ssh.NewClientConn(conn, conn.RemoteAddr().String(), config)
	if err != nil {
		return nil, err
	}

	sshConn := ssh.NewClient(c, chans, reqs)

	maxSessions := 5
	// sftpSessions := make(map[string]*SftpSession, maxSessions)
	sessionNames := make([]string, maxSessions)
	sftpSessionsChan := make(chan *SftpSession, maxSessions)

	i := 0
	for i < maxSessions {
		sftpClient, err := sftp.NewClient(sshConn, sftp.UseFstat(true), sftp.UseConcurrentWrites(true))
		if err != nil {
			continue
		}
		name := connName + "-" + fmt.Sprint(i)
		sftpSessionsChan <- &SftpSession{
			name: name,
			// sshSession: sshConn,
			client: sftpClient,
		}
		sessionNames[i] = name
		i++
	}

	return &Conn{
		connName: connName,
		// hostAddr: addr,
		// sshSessions: make(map[string]*ssh.Session),
		conn:            conn,
		sshConn:         sshConn,
		sshClientConfig: config,
		sessionNames:    sessionNames,
		// sftpSessions: sftpSessions,
		idelSftpSessionsChan: sftpSessionsChan,
		maxSessions:          maxSessions,
		createdAt:            time.Now(),
	}, nil
}

func (cn *Conn) UsedAt() time.Time {
	unix := atomic.LoadInt64(&cn.usedAt)
	return time.Unix(unix, 0)
}

func (cn *Conn) SetUsedAt(tm time.Time) {
	atomic.StoreInt64(&cn.usedAt, tm.Unix())
}

func (cn *Conn) Close() error {
	if cn.sshConn == nil {
		return nil
	}
	return cn.sshConn.Close()
}

func (cn *Conn) GetSftpSession(timeout time.Duration) (*SftpSession, error) {
	timer := timers.Get().(*time.Timer)
	timer.Reset(timeout)
	select {
	case sftpSession := <-cn.idelSftpSessionsChan:
		cn.idelSftpSessionsChan <- sftpSession
		timers.Put(timer)
		return sftpSession, nil
	case <-timer.C:
		timers.Put(timer)
		return nil, errors.New("timeout")
	}
}

func (cn *Conn) PutSftpSession(sftpSession *SftpSession) {
	cn.idelSftpSessionsChan <- sftpSession
}

func (cn *Conn) addOpeningFileNumBy(num int32) {
	atomic.AddInt32(&cn.atomicOpeningFiles, num)
}

// open file for read
func (cn *Conn) OpenFileForRead(path string) (*SftpFile, error) {
	return cn.OpenFile(path, os.O_RDONLY)
}

// open file for write
func (cn *Conn) OpenFileForWrite(path string) (*SftpFile, error) {
	return cn.OpenFile(path, os.O_WRONLY|os.O_CREATE)
}

func (cn *Conn) OpenFile(path string, osFlag int) (*SftpFile, error) {
	// todo hzx, fine grained session management
	// get ramdom sessionNames
	// r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// sessionName := cn.sessionNames[r.Intn(len(cn.sessionNames))]

	// timer := timers.Get().(*time.Timer)
	// timer.Reset(timeout)

	sftpSession := <-cn.idelSftpSessionsChan
	cn.idelSftpSessionsChan <- sftpSession
	// select {
	// case sftpSession = <- cn.idelSftpSessionsChan:
	// 	cn.idelSftpSessionsChan <- sftpSession
	// 	timers.Put(timer)
	// 	break
	// case <- timer.C:
	// 	timers.Put(timer)
	// 	return nil, errors.New("get sftp session timeout(s) " + timeout.String())
	// }

	log.Println("use session: ", sftpSession.name)

	sftpFp, err := sftpSession.OpenFile(path, osFlag)
	if err != nil {
		return nil, err
	}
	// atomic + 1
	cn.addOpeningFileNumBy(1)
	return sftpFp, nil
}

func (cn *Conn) CloseFile(fp *SftpFile) error {
	cn.addOpeningFileNumBy(-1)
	return fp.close()
}

func (cn *Conn) LstatFile(path string) (os.FileInfo, error) {
	sftpSession := <-cn.idelSftpSessionsChan
	cn.idelSftpSessionsChan <- sftpSession

	return sftpSession.LstatFile(path)
}

func (cn *Conn) RemoveFile(filePath string) error {
	sftpSession := <-cn.idelSftpSessionsChan
	cn.idelSftpSessionsChan <- sftpSession

	return sftpSession.RemoveFile(filePath)
}

func (cn *Conn) GetOpeningFileNum() int32 {
	return atomic.LoadInt32(&cn.atomicOpeningFiles)
}

func (cn *Conn) IsLightLoaded(openedFilesThreshold int32) bool {
	return cn.GetOpeningFileNum() < openedFilesThreshold
}
