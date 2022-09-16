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

package example

import "github.com/polarismesh/polaris-server/plugin"

const (
	PluginName = "example"
)

// init 初始化注册函数
func init() {
	plugin.RegisterPlugin(PluginName, &Example{})
}

// Example 例子
type Example struct{}

// Name 返回插件名字
func (p *Example) Name() string {
	return PluginName
}

// Destroy 销毁插件
func (p *Example) Destroy() error {
	return nil
}

// Initialize 插件初始化
func (p *Example) Initialize(c *plugin.ConfigEntry) error {
	return nil
}

// WhoAreYou 返回名字
func (p *Example) WhoAreYou(name string) string {
	return "f**k" + name
}
