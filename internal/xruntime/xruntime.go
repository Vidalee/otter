// Copyright (c) 2023 Alexey Mayshev and contributors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package xruntime

import (
	"math"
	"math/rand/v2"
	"runtime"
	"time"
	"unsafe"

	"golang.org/x/sys/cpu"
)

const (
	// CacheLineSize is useful for preventing false sharing.
	CacheLineSize = unsafe.Sizeof(cpu.CacheLinePad{})

	MaxDuration = time.Duration(math.MaxInt64)
)

// Parallelism returns the maximum possible number of concurrently running goroutines.
func Parallelism() uint32 {
	//nolint:gosec // there will never be an overflow
	maxProcs := uint32(runtime.GOMAXPROCS(0))
	//nolint:gosec // there will never be an overflow
	numCPU := uint32(runtime.NumCPU())
	if maxProcs < numCPU {
		return maxProcs
	}
	return numCPU
}

func Fastrand() uint32 {
	//nolint:gosec // we don't need a cryptographically secure random number generator
	return rand.Uint32()
}
