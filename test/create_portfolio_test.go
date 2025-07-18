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

package test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/VykintasD/advanced-trade-sdk-go/portfolios"
)

func TestCreatePortfolio(t *testing.T) {
	client, err := setupClient()
	if err != nil {
		t.Fatalf("Error setting up client: %v", err)
	}

	service := portfolios.NewPortfoliosService(client)

	response, err := service.CreatePortfolio(context.Background(), &portfolios.CreatePortfolioRequest{
		Name: fmt.Sprintf("test_portfolio-%d", time.Now().Unix()),
	})

	if err != nil {
		t.Errorf("Failed to create portfolio: %v", err)
	}

	if response == nil {
		t.Error("Expected non-nil response, got nil")
	}

}
