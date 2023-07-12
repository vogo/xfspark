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

type RequestHeader struct {
	AppID string `json:"app_id,omitempty"`
	UID   string `json:"uid,omitempty"`
}

type SystemSetting struct {
	Domain      string `json:"domain,omitempty"`
	Temperature string `json:"temperature,omitempty"`
	MaxToken    int    `json:"max_token,omitempty"`
	TopK        int    `json:"top_k,omitempty"`
	ChatID      string `json:"chat_id,omitempty"`
}

type RequestParameter struct {
	Chat *SystemSetting `json:"chat,omitempty"`
}
type RequestText struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type RequestMessage struct {
	Text []*RequestText `json:"text,omitempty"`
}

type RequestPayload struct {
	Message *RequestMessage `json:"message,omitempty"`
}

type Request struct {
	Header    *RequestHeader    `json:"header,omitempty"`
	Parameter *RequestParameter `json:"parameter,omitempty"`
	Payload   *RequestPayload   `json:"payload,omitempty"`
}

type ResponseHeader struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Sid     string `json:"sid,omitempty"`
	Status  int    `json:"status,omitempty"`
}

type ResponseChoicesText struct {
	Content string `json:"content"`
}

type ResponseChoices struct {
	Text   []ResponseChoicesText `json:"text"`
	Seq    int                   `json:"seq,omitempty"`
	Status int                   `json:"status,omitempty"`
}

type ResponsePayload struct {
	Choices *ResponseChoices `json:"choices,omitempty"`
}

type Response struct {
	Header  *ResponseHeader  `json:"header,omitempty"`
	Payload *ResponsePayload `json:"payload,omitempty"`
}
