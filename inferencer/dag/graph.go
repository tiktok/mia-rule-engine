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

package dag

import (
	"errors"

	"github.com/tiktok/mia-rule-engine/inferencer/rete"
	"gonum.org/v1/gonum/graph/encoding/dot"
	"gonum.org/v1/gonum/graph/simple"
)

// ErrGenDotCode the error thrown when fails to generate the dot code for virtualization
var ErrGenDotCode = errors.New("failed to generate dot code for the graph")

// Graph the virtualization of Rete network
type Graph struct {
	rete.Network
}

// NewGraph builds a new graph object
func NewGraph(network rete.Network) Graph {
	return Graph{
		Network: network,
	}
}

// Virtualize executes the logics to generate the network metadata for Gonum
func (gph *Graph) Virtualize() (string, error) {
	g := simple.NewDirectedGraph()

	idx := int64(0)
	root := gph.Network.Root()
	rootG := &RootNode{RootNode: root, id: idx}
	g.AddNode(rootG)
	idx++

	betaMap := make(map[string]*BetaNode)
	for _, typeNode := range root.GetTypes() {
		typeNodeG := &TypeNode{TypeNode: *typeNode, id: idx}
		g.SetEdge(g.NewEdge(rootG, typeNodeG))
		idx++
		for _, alphaNode := range typeNode.AlphaNodes() {
			alphaNodeG := &AlphaNode{AlphaNode: *alphaNode, id: idx}
			g.SetEdge(g.NewEdge(typeNodeG, alphaNodeG))
			idx++
			for _, betaNode := range alphaNode.BetaNodes() {
				betaNodeG, exist := betaMap[betaNode.String()]
				if !exist {
					betaNodeG = &BetaNode{
						BetaNode: *betaNode,
						id:       idx,
					}
					idx++
					betaMap[betaNode.String()] = betaNodeG
				}
				g.SetEdge(g.NewEdge(alphaNodeG, betaNodeG))
				for _, termNode := range betaNode.TermNodes() {
					termNodeG := &TermNode{TermNode: *termNode, id: idx}
					g.SetEdge(g.NewEdge(betaNodeG, termNodeG))
					idx++
				}
			}
			for _, termNode := range alphaNode.TermNodes() {
				termNodeG := &TermNode{TermNode: *termNode, id: idx}
				g.SetEdge(g.NewEdge(alphaNodeG, termNodeG))
				idx++
			}
		}
	}

	data, err := dot.Marshal(g, "Rete Network", "", "")
	if err != nil {
		return "", ErrGenDotCode
	}
	return string(data), nil
}
