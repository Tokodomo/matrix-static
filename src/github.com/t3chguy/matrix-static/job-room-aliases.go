// Copyright 2017 Michael Telatynski <7t3chguy@gmail.com>
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

package main

import (
	"github.com/t3chguy/matrix-static/mxclient"
	"github.com/t3chguy/matrix-static/utils"
)

type RoomAliasesResp struct {
	RoomInfo    mxclient.RoomInfo
	RoomAliases mxclient.RoomAliases
	PageSize    int
	Page        int
}

type RoomAliasesJob struct {
	roomID   string
	page     int
	pageSize int
}

func (job RoomAliasesJob) Work(w *Worker) {
	room := w.rooms[job.roomID]
	aliases := room.GetState().Aliases

	start, end := utils.CalcPaginationStartEnd(job.page, job.pageSize, len(aliases))

	w.Output <- RoomAliasesResp{
		room.RoomInfo(),
		aliases[start:end],
		job.pageSize,
		job.page,
	}
}
