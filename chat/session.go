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

type Session struct {
	svr *Server
	UID string
	Req *Request
}

func NewSession(svr *Server, uid string) *Session {
	session := &Session{
		svr: svr,
		UID: uid,
	}

	session.Req = &Request{
		Header: &RequestHeader{
			AppID: session.svr.appID,
			UID:   session.UID,
		},
		Parameter: &RequestParameter{
			Chat: &SystemSetting{
				Domain: "general",
			},
		},
		Payload: &RequestPayload{
			Message: &RequestMessage{},
		},
	}

	return session
}

func (s *Session) Send(question string) (string, error) {
	cmd := NewCommand(s)
	answer, respErr := cmd.execute(s.Req)

	if respErr != nil {
		return "", respErr
	}

	saveChatHistory(s.Req, question, answer)

	return answer, nil
}

func saveChatHistory(req *Request, question string, answer string) {
	history := req.Payload.Message.Text

	tokenCount := 0
	for _, text := range history {
		tokenCount += len(text.Content)
	}

	tokenCount += len(question) + len(answer)

	if tokenCount > MaxTokenSize {
		cutIndex := len(history)

		for index, text := range history {
			tokenCount -= len(text.Content)
			if tokenCount <= MaxTokenSize {
				cutIndex = index + 1
				break
			}
		}

		copy(history[:len(history)-cutIndex], history[cutIndex:])
	}

	history = append(history, &RequestText{Role: RoleUser, Content: question})
	history = append(history, &RequestText{Role: RoleAssistant, Content: answer})

	req.Payload.Message.Text = history
}
