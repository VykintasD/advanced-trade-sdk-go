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

package converts

import (
	"context"
	"fmt"

	"github.com/VykintasD/advanced-trade-sdk-go/client"
	"github.com/VykintasD/advanced-trade-sdk-go/model"
	"github.com/VykintasD/core-go"
)

type CommitConvertQuoteRequest struct {
	TradeId     string `json:"trade_id"`
	FromAccount string `json:"from_account"`
	ToAccount   string `json:"to_account"`
}

type CommitConvertQuoteResponse struct {
	Trade   *model.Convert             `json:"trade"`
	Request *CommitConvertQuoteRequest `json:"request"`
}

func (s convertsServiceImpl) CommitConvertQuote(
	ctx context.Context,
	request *CommitConvertQuoteRequest,
) (*CommitConvertQuoteResponse, error) {

	path := fmt.Sprintf("/brokerage/convert/trade/%s", request.TradeId)

	response := &CommitConvertQuoteResponse{Request: request}

	if err := core.HttpPost(
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
