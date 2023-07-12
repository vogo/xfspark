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

package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/vogo/xfspark/chat"
)

func main() {
	appID := chat.EnvString("APP_ID")
	apiKey := chat.EnvString("API_KEY")
	apiSecret := chat.EnvString("API_SECRET")

	if appID == "" || apiKey == "" || apiSecret == "" {
		panic("APP_ID, API_KEY, API_SECRET are required")
	}

	s := chat.NewServer(appID, apiKey, apiSecret)
	session, sessionErr := s.GetSession("123456789")

	if sessionErr != nil {
		panic(sessionErr)
	}

	answer := ""
	var err error
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("You: ")
		question, _ := reader.ReadString('\n')
		answer, err = session.Send(question)

		if err != nil {
			panic(err)
		}

		fmt.Println("AI: ", answer)
	}
}
