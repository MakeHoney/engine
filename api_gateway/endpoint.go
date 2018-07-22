/*
 * Copyright 2018 It-chain
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package api_gateway

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

//This file is based on the following sample.
// https://github.com/marcusolsson/goddd/blob/master/booking/endpoint.go
func makeFindUncommittedTransactionsEndpoint(t TransactionQueryApi) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		txs, err := t.FindUncommittedTransactions()

		if err != nil {
			return nil, err
		}

		return txs, nil
	}
}

/*
 * blockchain
 */

func makeFindCommittedBlocksEndpoint(b BlockQueryApi) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		blocks, err := b.commitedBlockRepository.GetEveryBlock()

		if err != nil {
			return nil, err
		}

		return blocks, nil
	}
}
