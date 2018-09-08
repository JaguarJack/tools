package controllers

import (
	"golang.org/x/crypto/ssh"
	"time"
	"net"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func connect(user, password, host string, port int) (*ssh.Session, error) {
	var (
		auth  []ssh.AuthMethod
		addr  string
		clientConfig *ssh.ClientConfig
		client *ssh.Client
		session *ssh.Session
		err  error
	)
	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))

	clientConfig = &ssh.ClientConfig{
		User: user,
		Auth: auth,
		Timeout: 30 * time.Second,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	// connet to ssh
	addr = fmt.Sprintf("%s:%d", host, port)

	if client, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, err
	}

	// create session
	if session, err = client.NewSession(); err != nil {
		return nil, err
	}

	return session, nil
}

func OutPut(c *gin.Context) {
	session, err := connect("", "", "", 22)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"msg": err.Error(),
			"data": "",
		})
		c.Abort()
		return
	}
	defer session.Close()
	//实时循环读取输出流中的一行内容
	command := c.PostForm("command")
	res, err := session.Output(command)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"msg": err.Error(),
			"data": "",
		})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
		"data": string(res),
	})

}
