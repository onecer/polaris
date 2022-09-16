/**
 * Tencent is pleased to support the open source community by making Polaris available.
 *
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 *
 * Licensed under the BSD 3-Clause License (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * https://opensource.org/licenses/BSD-3-Clause
 *
 * Unless required by applicable law or agreed to in writing, software distributed
 * under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
 * CONDITIONS OF ANY KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 */

package plugin

import (
	"os"
	"sync"

	"go.uber.org/zap"

	"github.com/polarismesh/polaris-server/common/log"
)

var (
	// 插件初始化原子变量
	exampleOnce = &sync.Once{}
)

// Example 这是一个例子
type Example interface {
	Plugin

	WhoAreYou(name string) string
}

// GetExample 获取例子实例
func GetExample() Example {
	c := &config.Example

	plugin, exist := pluginSet[c.Name]
	if !exist {
		log.Error("[Plugin][Example] not found", zap.String("name", c.Name))
		return nil
	}

	exampleOnce.Do(func() {
		if err := plugin.Initialize(c); err != nil {
			log.Errorf("plugin init err: %s", err.Error())
			os.Exit(-1)
		}
	})

	return plugin.(Example)
}
