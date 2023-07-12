/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package chat

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	"github.com/vogo/gorun"
	"github.com/vogo/logger"
)

type Command struct {
	s       *Session
	C       chan string
	runner  *gorun.Runner
	conn    *websocket.Conn
	respBuf bytes.Buffer
}

func NewCommand(s *Session) *Command {
	cmd := &Command{
		s:       s,
		C:       make(chan string),
		runner:  gorun.New(),
		respBuf: bytes.Buffer{},
	}

	return cmd
}

func (c *Command) start() error {
	authURL, err := getAuthorizationURL(chatAPIAddr, c.s.svr.apiKey, c.s.svr.apiSecret)
	if err != nil {
		return err
	}

	dialer := websocket.DefaultDialer
	header := make(http.Header)
	header.Set("Authorization", authURL.Query().Get("authorization"))
	header.Set("Date", authURL.Query().Get("date"))
	dialer.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	c.conn, _, err = dialer.Dial(strings.Replace(authURL.String(), "https://", "wss://", 1), header)

	if err != nil {
		return err
	}

	c.runner.Loop(c.handleConn)

	return nil
}

func (c *Command) Stop() {
	if c.conn != nil {
		_ = c.conn.Close()
		c.conn = nil
	}

	close(c.C)
}

func (c *Command) complete(message string) {
	c.C <- message

	// stop runner, to stop read message from websocket
	c.runner.Stop()
}

func (c *Command) handleConn() {
	_, message, err := c.conn.ReadMessage()
	if err != nil {
		logger.Warnf("read message error: %v", err)

		c.complete(err.Error())

		return
	}

	logger.Debugf("command response: %s", message)

	resp := &Response{}
	err = json.Unmarshal(message, resp)

	if err != nil {
		logger.Warnf("invalid message error: %v, message: %s", err, message)

		c.complete(string(message))

		return
	}

	c.handleResp(resp)
}

func (c *Command) handleResp(resp *Response) {
	if resp.Header.Code != 0 {
		c.complete(fmt.Sprintf("请求错误：%d %s", resp.Header.Code, resp.Header.Message))

		return
	}

	_, err := c.respBuf.WriteString(resp.Payload.Choices.Text[0].Content)
	if err != nil {
		c.complete(fmt.Sprintf("记录结果错误：%v", err))

		return
	}

	if resp.Header.Status == StatusLast {
		content := c.respBuf.String()
		c.complete(content)
	}
}

func (c *Command) execute(req *Request) (string, error) {
	var err error

	var body []byte

	body, err = json.Marshal(req)
	if err != nil {
		return "", err
	}

	if startErr := c.start(); startErr != nil {
		return "", startErr
	}

	// call stop only once.
	defer c.Stop()

	logger.Debugf("command request: %s", body)
	err = c.conn.WriteMessage(websocket.TextMessage, body)

	if err != nil {
		return "", err
	}

	select {
	case answer := <-c.C:
		return answer, nil
	case <-time.After(time.Second * defaultTimeoutSeconds):
		return "", ErrTimeout
	}
}
