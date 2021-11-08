/**
 * Copyright 2017 InsideSales.com Inc.
 * All Rights Reserved.
 *
 * NOTICE: All information contained herein is the property of InsideSales.com, Inc. and its suppliers, if
 * any. The intellectual and technical concepts contained herein are proprietary and are protected by
 * trade secret or copyright law, and may be covered by U.S. and foreign patents and patents pending.
 * Dissemination of this information or reproduction of this material is strictly forbidden without prior
 * written permission from InsideSales.com Inc.
 *
 * Requests for permission should be addressed to the Legal Department, InsideSales.com,
 * 1712 South East Bay Blvd. Provo, UT 84606.
 *
 * The software and any accompanying documentation are provided "as is" with no warranty.
 * InsideSales.com, Inc. shall not be liable for direct, indirect, special, incidental, consequential, or other
 * damages, under any theory of liability.
 */

package common

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
	"sync"
)

var config_map_lock sync.RWMutex
var json_data = map[string]string{}
var config_loaded = false
var defaults_loaded = false

/*
	This class is used for loading configuration values. It supports a three tier setup with precedences.
	See this.MustGetConfigVar for more information.
*/
type IConfigGetter interface {
	MustGetConfigVar(variableName string) string
	SafeGetConfigVar(variableName string) string
}

//Implements the IConfigGetter interface
type configGetter struct {
	/*
		config_file_path is the file path to the .json file to load the configurations from.
		There should also be a default configuration file at <config_file_path>.dist

		An example for this would be "config/config.json" and "config/config.json.dist".
		The files should be in the following json format:
			{
				"key1" : "value1",
				"key2" : "value2",
				"key3" : "value3"
			}
	*/
	config_file_path string
}

/*
	GetConfigGetter is the factory method to create a configGetter with injected information.
	@params
		config_file_path string This is the local file path to the .json file that holds your config json.
*/
func GetConfigGetter(config_file_path string) IConfigGetter {
	return &configGetter{
		config_file_path: config_file_path,
	}
}

/*
	GetConfigVar Loads a configuration value. If a value is not found, a panic is thrown
	Config precedence, highest to lowest:
	- Environment variables
	- Values of <config_file_path>
	- Values of <config_file_path>.dist
	Empty strings are considered to be unset. Do not use them as permitted values.

	@params
		variableName string The string value in the json file or environment variable name of the config to load
	@returns
		string The value found for the variableName of highest precedence found
*/
func (this configGetter) MustGetConfigVar(variableName string) string {
	config := os.Getenv(variableName)
	if config == "" {
		this.loadDefaults()
		this.loadConfig()
		ok := false
		config_map_lock.Lock()
		config, ok = json_data[variableName]
		config_map_lock.Unlock()
		if !ok {
			panic("FATAL ERROR COULD NOT LOAD VAR : " + variableName)
		}
	}
	return config
}

/*
	Same as this.MustGetConfigVar, but returns an empty string if the config is not found instead of panicking
*/
func (this configGetter) SafeGetConfigVar(variableName string) string {
	config := os.Getenv(variableName)
	if config == "" {
		this.loadDefaults()
		this.loadConfig()
		ok := false
		config_map_lock.Lock()
		config, ok = json_data[variableName]
		config_map_lock.Unlock()
		if !ok {
			return ""
		}
	}
	return config
}

// loadConfig Loads the overridden configuration values, if not yet done
func (this configGetter) loadConfig() {
	if !config_loaded {
		loadConfigFile(this.config_file_path)
		config_loaded = true
	}
}

// loadDefaults Loads the default configuration values, if not yet done
func (this configGetter) loadDefaults() {
	if !defaults_loaded {
		loadConfigFile(this.config_file_path + ".dist")
		defaults_loaded = true
	}
}

// loadConfigFile Loads json contents of a config file into json_data
// If the file doesn't exist, it is ignored. The file contents must be a JSON object with only string properties.
// Comment lines are ones that begin with '##', and they will be ignored along with empty lines
// config_file: The path to the JSON file to load
func loadConfigFile(config_file string) {
	if _, err := os.Stat(config_file); err == nil {
		reader_result, read_err := ioutil.ReadFile(config_file)
		if read_err != nil {
			panic("FATAL ERROR COULD NOT LOAD CONFIG FILE " + config_file + " Err: " + read_err.Error())
		}

		//Omit any line that begins with '##'
		all_lines := strings.Split(string(reader_result), "\n")
		cleaned_file_string := ""
		for _, line := range all_lines {
			if !strings.HasPrefix(strings.TrimSpace(line), "##") {
				cleaned_file_string += line + "\n"
			}
		}

		config_map_lock.Lock()
		json_err := json.Unmarshal([]byte(cleaned_file_string), &json_data)
		config_map_lock.Unlock()
		if json_err != nil {
			panic("FATAL ERROR COULD NOT UNMARSHAL CONFIG FILE " + config_file + " Err: " + json_err.Error())
		}
	}
}
