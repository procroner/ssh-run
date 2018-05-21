package server

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
)

func RunCommandWithPass(user, host, pass, cmd string) (string, error) {
	client, session, err := connectWithPass(user, host, pass)
	if err != nil {
		return "", err
	}
	defer client.Close()
	return runCommand(session, cmd)
}

func RunCommandAskPass(user, host, cmd string) (string, error) {
	client, session, err := connectAskPass(user, host)
	if err != nil {
		return "", err
	}
	defer client.Close()
	return runCommand(session, cmd)
}

func RunCommandWithKey(user, host, privateKeyPath, cmd string) (string, error) {
	client, session, err := connectWithKey(user, host, privateKeyPath)
	if err != nil {
		return "", err
	}
	defer client.Close()
	return runCommand(session, cmd)
}

func runCommand(session *ssh.Session, cmd string) (string, error) {
	output, err := session.CombinedOutput(cmd)
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func connectWithKey(user, host, privateKeyPath string) (*ssh.Client, *ssh.Session, error) {
	buffer, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		return nil, nil, err
	}
	key, err := ssh.ParsePrivateKey(buffer)
	if err != nil {
		return nil, nil, err
	}

	sshConfig := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{ssh.PublicKeys(key)},
	}
	return connect(sshConfig, host)
}

func connectAskPass(user, host string) (*ssh.Client, *ssh.Session, error) {
	var pass string
	fmt.Print("Password: ")
	fmt.Scanf("%s\n", &pass)
	return connectWithPass(user, host, pass)
}

func connectWithPass(user, host, pass string) (*ssh.Client, *ssh.Session, error) {
	sshConfig := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{ssh.Password(pass)},
	}
	return connect(sshConfig, host)
}

func connect(sshConfig *ssh.ClientConfig, host string) (*ssh.Client, *ssh.Session, error) {
	sshConfig.HostKeyCallback = ssh.InsecureIgnoreHostKey()
	client, err := ssh.Dial("tcp", host, sshConfig)
	if err != nil {
		return nil, nil, err
	}

	session, err := client.NewSession()
	if err != nil {
		client.Close()
		return nil, nil, err
	}

	return client, session, nil
}
