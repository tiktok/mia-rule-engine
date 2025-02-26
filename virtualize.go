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

// Package main
package main

import (
	"log"
	"os"
	"runtime/debug"

	"github.com/antlr4-go/antlr/v4"
	"github.com/tiktok/mia-rule-engine/analyzer"
	"github.com/tiktok/mia-rule-engine/analyzer/rule"
	"github.com/tiktok/mia-rule-engine/inferencer/dag"
	"github.com/tiktok/mia-rule-engine/inferencer/rete"
	cmpl "github.com/tiktok/mia-rule-engine/parser"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("parsing error: %v\n", r)
			debug.PrintStack()
		}
	}()
	policy := `ADD $fact AS $subject.age_status:underage WHEN $subject.age < 14 AND ($subject.geo = "KR" OR $subject.geo = "ID");
ADD $fact AS $subject.age_status:underage WHEN $subject.age < 13 AND $subject.geo = "US";
ADD $fact AS $subject.age_status:adult WHEN $subject.age >= 18;
GRANT $scope AS account:kids_mode WHEN $subject.age_status = "underage" AND $subject.geo = "US" PRIORITY 1;
GRANT $scope AS account:banned WHEN $subject.age_status = "underage" PRIORITY 2;`
	rules := analyze(policy)
	network := build(rules)
	graph := dag.NewGraph(network)
	data, err := graph.Virtualize()
	if err != nil {
		panic(err)
	}
	f, err := os.Create(".assets/rete_network.dot")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Printf("failed to close file: %v\n", err)
		}
	}()

	// Write string to file
	_, err = f.WriteString(data)
	if err != nil {
		panic(err)
	}
}

func analyze(policy string) []rule.Rule {
	input := antlr.NewInputStream(policy)

	lexer := cmpl.NewcmplLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := cmpl.NewcmplParser(stream)

	parser.RemoveErrorListeners()
	parser.AddErrorListener(analyzer.NewErrorListener())

	parser.BuildParseTrees = true // validate the syntax the build the AST
	tree := parser.Policy()
	listener := analyzer.NewParser()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	return listener.Rules()
}

func build(rules []rule.Rule) rete.Network {
	builder := rete.NewBuilder()
	for _, r := range rules {
		builder.AddRule(r)
	}
	return builder.Build()
}
