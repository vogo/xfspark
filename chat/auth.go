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
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/url"
	"time"
)

// getAuthorizationURL 获取授权URL
// see https://www.xfyun.cn/doc/spark/general_url_authentication.html.
func getAuthorizationURL(hostURL, apiKey, apiSecret string) (*url.URL, error) {
	urlAddr, err := url.Parse(hostURL)
	if err != nil {
		return nil, err
	}

	date := time.Now().UTC().Format("Mon, 02 Jan 2006 15:04:05 MST")

	// Get signature_origin
	signatureOrigin := fmt.Sprintf("host: %s\ndate: %s\nGET %s HTTP/1.1", urlAddr.Host, date, urlAddr.Path)

	// Calculate signature
	key := []byte(apiSecret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(signatureOrigin))
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))

	// Calculate authorization_origin
	authorizationOrigin := fmt.Sprintf("api_key=%q,algorithm=%q,headers=%q,signature=%q",
		apiKey, "hmac-sha256", "host date request-line", signature)

	// Encode authorization_origin
	authorization := base64.StdEncoding.EncodeToString([]byte(authorizationOrigin))

	// Construct authorization URL
	q := urlAddr.Query()
	q.Set("authorization", authorization)
	q.Set("date", date)
	q.Set("host", urlAddr.Host)
	urlAddr.RawQuery = q.Encode()

	return urlAddr, nil
}
