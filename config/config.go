// Copyright 2021 Northern.tech AS
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

package config

import (
	"github.com/mendersoftware/go-lib-micro/config"
)

const (
	// SettingListen is the config key for the listen address
	SettingListen = "listen"
	// SettingListenDefault is the default value for the listen address
	SettingListenDefault = ":8080"

	// SettingElasticsearchAddresses is the config key for the elasticsearch addresses
	SettingElasticsearchAddresses = "elasticsearch_addresses"
	// SettingElasticsearchAddressesDefault is the default value for the elasticsearch addresses
	SettingElasticsearchAddressesDefault = "http://localhost:9200"

	SettingInventoryAddr        = "inventory_addr"
	SettingInventoryAddrDefault = "http://mender-inventory:8080/"

	// SettingDebugLog is the config key for the truning on the debug log
	SettingDebugLog = "debug_log"
	// SettingDebugLogDefault is the default value for the debug log enabling
	SettingDebugLogDefault = false
)

var (
	// Defaults are the default configuration settings
	Defaults = []config.Default{
		{Key: SettingListen, Value: SettingListenDefault},
		{Key: SettingElasticsearchAddresses, Value: SettingElasticsearchAddressesDefault},
		{Key: SettingDebugLog, Value: SettingDebugLogDefault},
		{Key: SettingInventoryAddr, Value: SettingInventoryAddrDefault},
	}
)
