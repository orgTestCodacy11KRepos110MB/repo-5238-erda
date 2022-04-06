// Copyright (c) 2021 Terminus, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package utils

import (
	"fmt"
	"strings"

	"github.com/erda-project/erda/modules/core/monitor/storekit/clickhouse/table/loader"
)

func ConvertUnknownField(tableMeta *loader.TableMeta, field string) string {
	if tableMeta == nil {
		return field
	}
	_, ok := tableMeta.Columns[field]
	if ok {
		return field
	}
	splits := strings.SplitN(field, ".", 2)
	if len(splits) != 2 {
		return field
	}
	prefixCol, ok := tableMeta.Columns[splits[0]]
	if !ok {
		return field
	}
	switch prefixCol.Type {
	case loader.MapStringString:
		return fmt.Sprintf("%s['%s']", splits[0], splits[1])
	case loader.String:
		return fmt.Sprintf("JSONExtractString(%s,'%s')", splits[0], splits[1])
	default:
		return field
	}
}
