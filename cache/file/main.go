// Copyright 2020 beego-dev
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"time"

	"github.com/beego/beego/v2/client/cache"
	"github.com/beego/beego/v2/core/logs"
)

func main() {
	bm, err := cache.NewCache("file", `{"CachePath":"./cache","FileSuffix":".cache", "EmbedExpiry": "120"}`)
	if err != nil {
		logs.Error(err)
	}

	// put
	isPut := bm.Put(context.Background(), "name", "beego", time.Second*10)
	logs.Info(isPut)

	isPut = bm.Put(context.Background(), "hello", "world", time.Second*10)
	logs.Info(isPut)

	// get
	result, _ := bm.Get(context.Background(), "name")
	logs.Info(result)

	multiResult, _ := bm.GetMulti(context.Background(), []string{"name", "hello"})
	for i := range multiResult {
		logs.Info(multiResult[i])
	}

	// isExist
	isExist, _ := bm.IsExist(context.Background(), "name")
	logs.Info(isExist)

	// delete
	isDelete := bm.Delete(context.Background(), "name")
	logs.Info(isDelete)
}
