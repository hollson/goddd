// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package core

type Pages struct {
    NeedAll   bool // 特殊处理，需要所有数据
    OnlyCount bool // 只需要数量，不需要数据
    Page      int
    PageSize  int // 默认必须控制返回数据的数量禁止大于100，禁止全量返回所有数据。如果需要返回所有基础数据等，额外写接口。
}