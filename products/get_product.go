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

package products

import (
	"context"
	"fmt"

	"github.com/VykintasD/advanced-trade-sdk-go/client"
	"github.com/VykintasD/advanced-trade-sdk-go/model"
	"github.com/VykintasD/core-go"
)

type GetProductRequest struct {
	ProductId string `json:"product_id"`
}

type GetProductResponse struct {
	ProductId                 string                     `json:"product_id"`
	Price                     string                     `json:"price"`
	PricePercentageChange24h  string                     `json:"price_percentage_change_24h"`
	Volume24h                 string                     `json:"volume_24h"`
	VolumePercentageChange24h string                     `json:"volume_percentage_change_24h"`
	BaseIncrement             string                     `json:"base_increment"`
	QuoteIncrement            string                     `json:"quote_increment"`
	QuoteMinSize              string                     `json:"quote_min_size"`
	QuoteMaxSize              string                     `json:"quote_max_size"`
	BaseMinSize               string                     `json:"base_min_size"`
	BaseMaxSize               string                     `json:"base_max_size"`
	BaseName                  string                     `json:"base_name"`
	QuoteName                 string                     `json:"quote_name"`
	Watched                   bool                       `json:"watched"`
	IsDisabled                bool                       `json:"is_disabled"`
	New                       bool                       `json:"new"`
	Status                    string                     `json:"status"`
	CancelOnly                bool                       `json:"cancel_only"`
	LimitOnly                 bool                       `json:"limit_only"`
	PostOnly                  bool                       `json:"post_only"`
	TradingDisabled           bool                       `json:"trading_disabled"`
	AuctionMode               bool                       `json:"auction_mode"`
	ProductType               string                     `json:"product_type"`
	QuoteCurrencyId           string                     `json:"quote_currency_id"`
	BaseCurrencyId            string                     `json:"base_currency_id"`
	FCMSessionDetails         model.SessionDetails       `json:"fcm_trading_session_details"`
	MidMarketPrice            string                     `json:"mid_market_price"`
	Alias                     string                     `json:"alias"`
	AliasTo                   []string                   `json:"alias_to"`
	BaseDisplaySymbol         string                     `json:"base_display_symbol"`
	QuoteDisplaySymbol        string                     `json:"quote_display_symbol"`
	ViewOnly                  bool                       `json:"view_only"`
	PriceIncrement            string                     `json:"price_increment"`
	FutureProductDetails      model.FutureProductDetails `json:"future_product_details"`
	Request                   *GetProductRequest         `json:"request"`
}

func (s productsServiceImpl) GetProduct(
	ctx context.Context,
	request *GetProductRequest,
) (*GetProductResponse, error) {

	path := fmt.Sprintf("/brokerage/products/%s", request.ProductId)

	response := &GetProductResponse{Request: request}

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
