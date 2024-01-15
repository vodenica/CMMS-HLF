/*
 * Copyright 2018 IBM All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the 'License');
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an 'AS IS' BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Copyright the Hyperledger Fabric contributors. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// CustomTransactionContextInterface adds extra methods to basic context
// interface to give access to callData
type CustomTransactionContextInterface interface {
	contractapi.TransactionContextInterface

	SetCallData([]byte)
	GetCallData() []byte
}

// CustomTransactionContext adds extra field to contractapi.TransactionContext
// so that data can be between calls
type CustomTransactionContext struct {
	contractapi.TransactionContext
	callData []byte
}

// SetCallData sets the call data property
func (ctx *CustomTransactionContext) SetCallData(bytes []byte) {
	ctx.callData = bytes
}

// GetCallData gets the call data property
func (ctx *CustomTransactionContext) GetCallData() []byte {
	return ctx.callData
}

// GetWorldState takes a key and sets what is found in the world state for that
// key in the transaction context
func GetWorldState(ctx CustomTransactionContextInterface) error {
	_, params := ctx.GetStub().GetFunctionAndParameters()

	if len(params) < 1 {
		return errors.New("missing key for world state")
	}

	existing, err := ctx.GetStub().GetState(params[0])

	if err != nil {
		return errors.New("unable to interact with world state")
	}

	ctx.SetCallData(existing)

	return nil
}

// UnknownTransactionHandler logs details of a bad transaction request
// and returns a shim error
func UnknownTransactionHandler(ctx CustomTransactionContextInterface) error {
	fcn, args := ctx.GetStub().GetFunctionAndParameters()
	return fmt.Errorf("invalid function %s passed with args [%s]", fcn, strings.Join(args, ", "))
}
