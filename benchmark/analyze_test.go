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
	"fmt"
	"testing"

	"github.com/antlr4-go/antlr/v4"
	"github.com/tiktok/mia-rule-engine/analyzer"
	cmpl "github.com/tiktok/mia-rule-engine/parser"
)

func BenchmarkAnalyze(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if err := analyze(POLICY); err != nil {
			b.Fatalf("Parsing failed: %v", err)
		}
	}
}

func analyze(policy string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("\n[ERR]: %v\n", r)
		}
	}()

	input := antlr.NewInputStream(policy)
	lexer := cmpl.NewcmplLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	parser := cmpl.NewcmplParser(stream)
	parser.BuildParseTrees = true
	parser.RemoveErrorListeners()
	parser.AddErrorListener(analyzer.NewErrorListener())
	parser.Policy()
	return nil
}
