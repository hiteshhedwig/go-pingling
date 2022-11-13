package pinger

import (
	"fmt"

	sshcl "github.com/helloyi/go-sshclient"
	"golang.org/x/crypto/ssh"
)

func Sshclient(ip string, username string, passwd string) error {
	client, err := sshcl.DialWithPasswd(ip, username, passwd)
	if err != nil {
		return fmt.Errorf("cannot SSH into the provided ip: %v", err)
	}

	defer client.Close()

	// with a terminal config
	config := &sshcl.TerminalConfig{
		Term:   "xterm",
		Height: 40,
		Weight: 80,
		Modes: ssh.TerminalModes{
			ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
			ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
		},
	}

	if err := client.Terminal(config).Start(); err != nil {
		return err
	}

	return nil
}
