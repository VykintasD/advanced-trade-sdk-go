/**
 * Copyright 2024-present Coinbase Global, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package futures

import (
	"context"

	"github.com/VykintasD/advanced-trade-sdk-go/client"
	"github.com/VykintasD/advanced-trade-sdk-go/model"
	"github.com/VykintasD/core-go"
)

type ListFuturesSweepsRequest struct{}

type ListFuturesSweepsResponse struct {
	Sweeps  *model.Sweep              `json:"sweeps"`
	Request *ListFuturesSweepsRequest `json:"request"`
}

func (s futuresServiceImpl) ListFuturesSweeps(
	ctx context.Context,
	request *ListFuturesSweepsRequest,
) (*ListFuturesSweepsResponse, error) {

	path := "/brokerage/cfm/sweeps"

	response := &ListFuturesSweepsResponse{Request: request}

	if err := core.HttpGet(
		ctx,
		s.client,
		path,
		core.EmptyQueryParams,
		client.DefaultSuccessHttpStatusCodes,
		request,
		response,
		s.client.HeadersFunc(),
	); err != nil {
		return nil, err
	}

	return response, nil
}
