// Copyright 2020 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package tracers

import (
	"encoding/json"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

type Tracer interface {
	// Stop terminates execution of the tracer at the first opportune moment.
	Stop(error)

	// CaptureState is called for each step of the VM with the current VM state.
	CaptureState(env *vm.EVM, pc uint64, op vm.OpCode, gas, cost uint64, memory *vm.Memory, stack *vm.Stack, contract *vm.Contract, depth int, err error) error

	// CaptureStart implements the Tracer interface to initialize the tracing operation.
	CaptureStart(common.Address, common.Address, bool, []byte, uint64, *big.Int) error

	// CaptureFault implements the Tracer interface to trace an execution fault
	// while running an opcode.
	CaptureFault(*vm.EVM, uint64, vm.OpCode, uint64, uint64, *vm.Memory, *vm.Stack, *vm.Contract, int, error) error

	// CaptureEnd is called after the call finishes to finalize the tracing.
	CaptureEnd([]byte, uint64, time.Duration, error) error

	// GetResult calls the Javascript 'result' function and returns its value, or any accumulated error
	GetResult() (json.RawMessage, error)
}

func New(code string) (Tracer, error) {
	return NewJSTracer(code)
}
