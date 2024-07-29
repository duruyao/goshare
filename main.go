//  Copyright 2022-2032 Ryan Du <duruyao@gmail.com>
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

var (
	arg  = NewArgument()
	quit = make(chan struct{})
)

func main() {
	if arg.WantHelp() {
		fmt.Println(arg.Usage())
		return
	}
	if arg.WantVersion() {
		fmt.Println(VersionSerial())
		return
	}

	host := arg.Host()
	path := AbsPathMust(arg.Path())
	scheme := arg.Scheme()
	urlPrefix := FixedUrlPrefix(arg.UrlPrefix())

	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		fmt.Println(err.Error())
		fmt.Println(arg.Usage())
		return
	}
	if info.IsDir() {
		go HttpStaticFS(host, path, urlPrefix)
		fmt.Println(RunningStatus(path, host, scheme, urlPrefix, ""))
	} else {
		go HttpStaticFile(host, path, urlPrefix)
		fmt.Println(RunningStatus(filepath.Dir(path), host, scheme, urlPrefix, filepath.Base(path)))
	}
	<-quit
}
