/*
Copyright 2024 The west2-online Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package service

import (
	"path/filepath"
	"strings"

	"github.com/west2-online/DomTok/pkg/errno"
)

type CommodityVerifyOps func() error

func (svc *CommodityService) Verify(opts ...CommodityVerifyOps) error {
	for _, opt := range opts {
		if err := opt(); err != nil {
			return err
		}
	}
	return nil
}

func (svc *CommodityService) VerifyFileType(filename string) CommodityVerifyOps {
	return func() error {
		// TODO: 类型需要更多的吗？
		ext := strings.ToLower(filepath.Ext(filename))
		legals := []string{".jpg", ".jpeg", ".png", ".gif", ".bmp"}
		for _, legal := range legals {
			if ext == legal {
				return nil
			}
		}
		return errno.NewErrNo(errno.ParamVerifyErrorCode, "invalid file type")
	}
}
