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

package main

import (
	_ "github.com/polarismesh/polaris-server/apiserver/eurekaserver"
	_ "github.com/polarismesh/polaris-server/apiserver/grpcserver/config"
	_ "github.com/polarismesh/polaris-server/apiserver/grpcserver/discover"
	_ "github.com/polarismesh/polaris-server/apiserver/httpserver"
	_ "github.com/polarismesh/polaris-server/apiserver/l5pbserver"
	_ "github.com/polarismesh/polaris-server/apiserver/prometheussd"
	_ "github.com/polarismesh/polaris-server/apiserver/xdsserverv3"
	_ "github.com/polarismesh/polaris-server/auth/defaultauth"
	_ "github.com/polarismesh/polaris-server/cache"
	_ "github.com/polarismesh/polaris-server/plugin/auth/defaultauth"
	_ "github.com/polarismesh/polaris-server/plugin/auth/platform"
	_ "github.com/polarismesh/polaris-server/plugin/cmdb/memory"
	_ "github.com/polarismesh/polaris-server/plugin/discoverevent/local"
	_ "github.com/polarismesh/polaris-server/plugin/discoverstat/discoverlocal"
	_ "github.com/polarismesh/polaris-server/plugin/example"
	_ "github.com/polarismesh/polaris-server/plugin/healthchecker/heartbeatmemory"
	_ "github.com/polarismesh/polaris-server/plugin/healthchecker/heartbeatredis"
	_ "github.com/polarismesh/polaris-server/plugin/history/logger"
	_ "github.com/polarismesh/polaris-server/plugin/password"
	_ "github.com/polarismesh/polaris-server/plugin/ratelimit/lrurate"
	_ "github.com/polarismesh/polaris-server/plugin/ratelimit/token"
	_ "github.com/polarismesh/polaris-server/plugin/statis/local"
	_ "github.com/polarismesh/polaris-server/store/boltdb"
	_ "github.com/polarismesh/polaris-server/store/sqldb"
)
