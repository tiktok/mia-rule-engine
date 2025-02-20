/*
 * Copyright 2022 TikTok Pte. Ltd.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package benchmark

import (
	"testing"

	"github.com/tiktok/mia-rule-engine/inferencer/rete"
)

func BenchmarkInference(b *testing.B) {
	network := rete.BuildNetwork(POLICY)
	inputs := make([]*rete.Input, b.N)
	for i := 0; i < b.N; i++ {
		inputs[i] = rete.NewInput([]interface{}{
			"$subject.age", 13, "$subject.geo", "KR",
		})
	}
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		network.Infer(inputs[i])
	}
}
