// LianDi - 链滴笔记，连接点滴
// Copyright (c) 2020-present, b3log.org
//
// LianDi is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package main

import "encoding/json"

type setmd struct {
	*BaseCmd
}

func (cmd *setmd) Exec() {
	ret := NewCmdResult(cmd.Name(), cmd.id)

	param, err := json.Marshal(cmd.param)
	if nil != err {
		ret.Code = -1
		ret.Msg = err.Error()
		Push(ret.Bytes())
		return
	}

	md := &Markdown{}
	if err = json.Unmarshal(param, md); nil != err {
		ret.Code = -1
		ret.Msg = err.Error()
		Push(ret.Bytes())
		return
	}

	Conf.Markdown = md
	Conf.Save()

	ret.Data = Conf.Markdown
	Push(ret.Bytes())
}

func (cmd *setmd) Name() string {
	return "setmd"
}
