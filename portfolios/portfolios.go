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

package portfolios

import (
	"context"

	"github.com/VykintasD/advanced-trade-sdk-go/client"
)

type PortfoliosService interface {
	CreatePortfolio(ctx context.Context, request *CreatePortfolioRequest) (*CreatePortfolioResponse, error)
	DeletePortfolio(ctx context.Context, request *DeletePortfolioRequest) (*DeletePortfolioResponse, error)
	EditPortfolio(ctx context.Context, request *EditPortfolioRequest) (*EditPortfolioResponse, error)
	GetPortfolioBreakdown(ctx context.Context, request *GetPortfolioBreakdownRequest) (*GetPortfolioBreakdownResponse, error)
	ListPortfolios(ctx context.Context, request *ListPortfoliosRequest) (*ListPortfoliosResponse, error)
	MovePortfolioFunds(ctx context.Context, request *MovePortfolioFundsRequest) (*MovePortfolioFundsResponse, error)
}

func NewPortfoliosService(c client.RestClient) PortfoliosService {
	return &portfoliosServiceImpl{client: c}
}

type portfoliosServiceImpl struct {
	client client.RestClient
}
