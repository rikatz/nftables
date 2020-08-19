// Copyright 2020 Google LLC. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package expr

import (
	"encoding/binary"

	"github.com/mdlayher/netlink"
)

type Target struct {
	Info []byte
	Name string
}

// Propose a PR later to Golang
const (
	NFTA_TARGET_UNSPEC   = 0x0
	NFTA_TARGET_NAME     = 0x1
	NFTA_TARGET_REV      = 0x2
	NFTA_TARGET_INFO     = 0x3
	NFTA_TARGET_PROTOCOL = 0x4
	NFTA_TARGET_FLAGS    = 0x5
)

// TODO HERE :D
func (e *Target) marshal() ([]byte, error) {
	return []byte{}, nil
}

func (e *Target) unmarshal(data []byte) error {
	ad, err := netlink.NewAttributeDecoder(data)
	if err != nil {
		return err
	}
	ad.ByteOrder = binary.BigEndian
	for ad.Next() {
		switch ad.Type() {
		case NFTA_TARGET_NAME:
			e.Name = ad.String()
		case NFTA_TARGET_INFO:
			e.Info = ad.Bytes()
		}
	}
	return ad.Err()
}
