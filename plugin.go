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
	"github.com/mailgun/vulcand/Godeps/_workspace/src/github.com/codegangsta/cli"
	"github.com/mailgun/vulcand/Godeps/_workspace/src/github.com/mailgun/vulcan/middleware"
	"github.com/mailgun/vulcand/plugin"
)

type SwiftCache struct{}

func NewSwiftCache() *SwiftCache {
	return &SwiftCache{}
}

func (sc *SwiftCache) NewMiddleware() (middleware.Middleware, error) {
	return NewSwiftCacher()
}

func FromOther(a SwiftCache) (plugin.Middleware, error) {
	return NewSwiftCache(), nil
}

// Constructs the middleware from the command line
func FromCli(c *cli.Context) (plugin.Middleware, error) {
	return NewSwiftCache(), nil
}

func cliFlags() []cli.Flag {
	return []cli.Flag{}
}

func GetSpec() *plugin.MiddlewareSpec {
	return &plugin.MiddlewareSpec{
		Type:      "swiftcache",
		FromOther: FromOther,
		FromCli:   FromCli,
		CliFlags:  cliFlags(),
	}
}
