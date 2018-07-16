// Copyright 2017, OpenCensus Authors
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
//

package stats

import "sync/atomic"

// Float64Measure is a measure for float64 values.
type Float64Measure struct {
	md *measureDescriptor
}

// Name returns the name of the measure.
func (m *Float64Measure) Name() string {
	return m.md.name
}

// Description returns the description of the measure.
func (m *Float64Measure) Description() string {
	return m.md.description
}

// Unit returns the unit of the measure.
func (m *Float64Measure) Unit() string {
	return m.md.unit
}

// M creates a new float64 measurement.
// Use Record to record measurements.
func (m *Float64Measure) M(v float64) Measurement {
	if n := atomic.LoadInt32(&registrySize); !m.md.subscribed() && n == 0 {
		return Measurement{}
	}
	return Measurement{m: m, v: v}
}

// Float64 creates a new measure for float64 values.
//
// See the documentation for interface Measure for more guidance on the
// parameters of this function.
func Float64(name, description, unit string) *Float64Measure {
	mi := registerMeasureHandle(name, description, unit)
	return &Float64Measure{mi}
}
