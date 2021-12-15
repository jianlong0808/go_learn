package sshutils

import (
	"errors"
	"fmt"
	"go_learn/017.utils/logging"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"time"
)

var log, _ = logging.NewLoggerWithRotate()

//GetSSHConnection 获取远程主机的ssh连接, 需要id_rsa私钥文件
func GetSSHConnection(user, host string, port int) (*ssh.Client, error) {
	var (
		addr         string
		clientConfig *ssh.ClientConfig
		client       *ssh.Client
		err          error
	)

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return nil, errors.New("Can not get executable file path ")
	}

	key, err := ioutil.ReadFile(path.Join(dir, "id_rsa"))
	if err != nil {
		return nil, err
	}
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return nil, err
	}

	clientConfig = &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		Timeout:         5 * time.Second,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// connet to ssh
	addr = fmt.Sprintf("%s:%d", host, port)

	if client, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, err
	}
	log.Infof("获取 " + user + "@" + addr + " ssh连接成功")
	return client, nil
}
