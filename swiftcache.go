/**
 *  Copyright 2014 Paul Querna
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 */

package swiftcache

import (
	"github.com/mailgun/vulcand/Godeps/_workspace/src/github.com/mailgun/vulcan/request"
	"github.com/pquerna/vulcancache"
	"net/http"
)

type SwiftCacher struct {
	cache *vulcancache.HttpCacher
}

func (sc *SwiftCacher) ProcessRequest(r request.Request) (*http.Response, error) {
	return nil, nil
}

func NewSwiftCacher() (*SwiftCacher, error) {
	c := vulcancache.NewHttpCacher(&vulcancache.CacheOptions{})
	return &SwiftCacher{cache: c}, nil
}

func (sc *SwiftCacher) ProcessResponse(r request.Request, a request.Attempt) {}
