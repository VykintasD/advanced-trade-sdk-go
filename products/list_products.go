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
	"github.com/VykintasD/advanced-trade-sdk-go/utils"
	"github.com/VykintasD/core-go"
)

type ListProductsRequest struct {
	ProductType            string                  `json:"product_type,omitempty"`
	ProductIds             []string                `json:"product_ids,omitempty"`
	ContractExpiryType     string                  `json:"contract_expiry_type,omitempty"`
	ExpiringContractStatus string                  `json:"expiring_contract_status,omitempty"`
	Pagination             *model.PaginationParams `json:"pagination_params,omitempty"`
}

type ListProductsResponse struct {
	Products []*model.Product     `json:"products"`
	Request  *ListProductsRequest `json:"request"`
}

func (s productsServiceImpl) ListProducts(
	ctx context.Context,
	request *ListProductsRequest,
) (*ListProductsResponse, error) {

	path := fmt.Sprintf("/brokerage/products")

	var queryParams string

	for _, p := range request.ProductIds {
		queryParams = core.AppendHttpQueryParam(queryParams, "product_ids", p)
	}

	if len(request.ProductType) > 0 {
		queryParams = core.AppendHttpQueryParam(queryParams, "product_type", request.ProductType)
	}

	if len(request.ContractExpiryType) > 0 {
		queryParams = core.AppendHttpQueryParam(queryParams, "contract_expiry_type", request.ContractExpiryType)
	}

	if len(request.ExpiringContractStatus) > 0 {
		queryParams = core.AppendHttpQueryParam(queryParams, "expiring_contract_status", request.ExpiringContractStatus)
	}

	queryParams = utils.AppendPaginationParams(queryParams, request.Pagination)

	response := &ListProductsResponse{Request: request}

	if err := core.HttpGet(
		ctx,
		s.client,
		path,
		queryParams,
		client.DefaultSuccessHttpStatusCodes,
		request,
		response,
		s.client.HeadersFunc(),
	); err != nil {
		return nil, err
	}

	return response, nil
}
