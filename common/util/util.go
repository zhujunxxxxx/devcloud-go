/*
 * Copyright (c) Huawei Technologies Co., Ltd. 2021.
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License.  You may obtain a copy of the
 * License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed
 * under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
 * CONDITIONS OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 *
 * Package util provides some util function, such as ValidateHostPort.
 */

package util

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

// ValidateHostPort validate that hostPort is correct.
func ValidateHostPort(hostPort string) error {
	host, port, err := net.SplitHostPort(hostPort)
	if err != nil {
		return err
	}
	if ip := net.ParseIP(host); ip == nil {
		return fmt.Errorf("bad IP address: %s", host)
	}

	p, err := strconv.Atoi(port)
	if err != nil {
		return err
	}
	if p < 1 || p > 65535 {
		return fmt.Errorf("bad port number %s", port)
	}
	return nil
}

// ConvertAddressStrToSlice convert address like "127.0.0.1:2379,127.0.0.1:2380" to endpoints like ["127.0.0.1:2379", "127.0.0.1:2380"]
func ConvertAddressStrToSlice(addressStr string) []string {
	addressSlice := strings.Split(addressStr, ",")
	var res []string
	for _, address := range addressSlice {
		address = strings.TrimSpace(address)
		if len(address) == 0 {
			continue
		}
		if err := ValidateHostPort(address); err != nil {
			log.Printf("ERROR: hostPort '%s' is invalid, %v", address, err)
			continue
		}
		res = append(res, address)
	}
	return res
}
