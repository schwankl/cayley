// Copyright 2014 The Cayley Authors. All rights reserved.
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

package iterator

// A quickly mocked version of the TripleStore interface, for use in tests.
// Can better used Mock.Called but will fill in as needed.

import (
	"github.com/stretchrcom/testify/mock"

	"github.com/google/cayley/graph"
)

type TestTripleStore struct {
	mock.Mock
}

func (ts *TestTripleStore) GetIdFor(s string) graph.TSVal {
	args := ts.Mock.Called(s)
	return args.Get(0)
}
func (ts *TestTripleStore) AddTriple(*graph.Triple)             {}
func (ts *TestTripleStore) AddTripleSet([]*graph.Triple)        {}
func (ts *TestTripleStore) GetTriple(graph.TSVal) *graph.Triple { return &graph.Triple{} }
func (ts *TestTripleStore) GetTripleIterator(d graph.Direction, i graph.TSVal) graph.Iterator {
	args := ts.Mock.Called(d, i)
	return args.Get(0).(graph.Iterator)
}
func (ts *TestTripleStore) GetNodesAllIterator() graph.Iterator   { return &Null{} }
func (ts *TestTripleStore) GetTriplesAllIterator() graph.Iterator { return &Null{} }
func (ts *TestTripleStore) GetIteratorByString(string, string, string) graph.Iterator {
	return &Null{}
}
func (ts *TestTripleStore) GetNameFor(v graph.TSVal) string {
	args := ts.Mock.Called(v)
	return args.Get(0).(string)
}
func (ts *TestTripleStore) Size() int64 { return 0 }
func (ts *TestTripleStore) DebugPrint() {}
func (ts *TestTripleStore) OptimizeIterator(it graph.Iterator) (graph.Iterator, bool) {
	return &Null{}, false
}
func (ts *TestTripleStore) FixedIterator() graph.FixedIterator {
	return NewFixedIteratorWithCompare(BasicEquality)
}
func (ts *TestTripleStore) Close()                                                      {}
func (ts *TestTripleStore) GetTripleDirection(graph.TSVal, graph.Direction) graph.TSVal { return 0 }
func (ts *TestTripleStore) RemoveTriple(t *graph.Triple)                                {}