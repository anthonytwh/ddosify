/*
*
*	Ddosify - Load testing tool for any web system.
*   Copyright (C) 2021  Ddosify (https://ddosify.com)
*
*   This program is free software: you can redistribute it and/or modify
*   it under the terms of the GNU Affero General Public License as published
*   by the Free Software Foundation, either version 3 of the License, or
*   (at your option) any later version.
*
*   This program is distributed in the hope that it will be useful,
*   but WITHOUT ANY WARRANTY; without even the implied warranty of
*   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
*   GNU Affero General Public License for more details.
*
*   You should have received a copy of the GNU Affero General Public License
*   along with this program.  If not, see <https://www.gnu.org/licenses/>.
*
 */

package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"go.ddosify.com/ddosify/core/types"
	"go.ddosify.com/ddosify/core/util"
)

// Costants of the Supported Config Types
const (
	ConfigTypeJson    = "jsonReader"
)

var availableConfigTypes = []string{ConfigTypeJson}

// ConfigReader is the interface that abstracts different config reader implementations.
type ConfigReader interface {
	init([]byte) error
	CreateHammer() (types.Hammer, error)
}

// NewConfigReader is the factory method of the ConfigReader.
func NewConfigReader(path string, configType string) (reader ConfigReader, err error) {
	if !util.StringInSlice(configType, availableConfigTypes) {
		return nil, fmt.Errorf("unsupported config type %s", configType)
	}

	if strings.EqualFold(configType, ConfigTypeJson) {
		reader = &jsonReader{}
	}
	jsonFile, err := os.Open(path)
	if err != nil {
		return
	}

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return
	}

	if !json.Valid(byteValue) {
		return nil, fmt.Errorf("json is invalid")
	}

	err = reader.init(byteValue)

	return
}