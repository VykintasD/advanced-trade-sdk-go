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

package public

import (
	"context"
	"fmt"
	"time"

	"github.com/VykintasD/advanced-trade-sdk-go/client"
	"github.com/VykintasD/core-go"
)

type GetServerTimeRequest struct{}

type GetServerTimeResponse struct {
	Iso          time.Time             `json:"iso"`
	EpochSeconds string                `json:"epochSeconds"`
	EpochMillis  string                `json:"epochMillis"`
	Request      *GetServerTimeRequest `json:"request"`
}

func (s publicServiceImpl) GetServerTime(
	ctx context.Context,
	request *GetServerTimeRequest,
) (*GetServerTimeResponse, error) {

	path := fmt.Sprint("/brokerage/time")

	response := &GetServerTimeResponse{Request: request}

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
