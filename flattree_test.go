package flattree_test

import (
	"reflect"
	"testing"

	"github.com/bcomnes/flattree"
)

type testpair struct {
	in  []uint64
	out uint64
}

type testpairRange struct {
	in  []uint64
	out flattree.Range
}

type testpairArray struct {
	in  uint64
	out []uint64
}

func TestBaseBlocks(t *testing.T) {
	var tests = []testpair{
		{[]uint64{0, 0}, 0},
		{[]uint64{0, 1}, 2},
		{[]uint64{0, 2}, 4},
	}

	for _, pair := range tests {
		v := flattree.Index(pair.in[0], pair.in[1])
		if v != pair.out {
			t.Error("For", pair.in, "expected", pair.out, "got", v)
		}
	}
}

func TestParents(t *testing.T) {
	var indexTests = []testpair{
		{[]uint64{1, 0}, 1},
		{[]uint64{1, 1}, 5},
		{[]uint64{2, 0}, 3},
	}

	for _, pair := range indexTests {
		v := flattree.Index(pair.in[0], pair.in[1])
		if v != pair.out {
			t.Error("For", pair.in, "expected", pair.out, "got", v)
		}
	}

	var parentTests = []testpair{
		{[]uint64{0, 0}, 1},
		{[]uint64{2, 0}, 1},
		{[]uint64{1, 0}, 3},
	}

	for _, pair := range parentTests {
		v := flattree.Parent(pair.in[0], pair.in[1])
		if v != pair.out {
			t.Error("For", pair.in, "expected", pair.out, "got", v)
		}
	}
}

func TestChildren(t *testing.T) {
	var childrenTests = []testpairRange{
		{[]uint64{0, 0}, flattree.Range{}},
		{[]uint64{1, 0}, flattree.Range{0, 2}},
		{[]uint64{3, 0}, flattree.Range{1, 5}},
	}

	for _, pair := range childrenTests {
		v := flattree.Children(pair.in[0], pair.in[1])
		if !reflect.DeepEqual(v, pair.out) {
			t.Error("For", pair.in, "expected", pair.out, "got", v)
		}
	}
}

type childTest struct {
	in  uint64
	out uint64
	err bool
}

func TestLeftChild(t *testing.T) {
	var tests = []childTest{
		{0, 0, true},
		{1, 0, false},
		{3, 1, false},
	}

	for _, test := range tests {
		v, err := flattree.LeftChild(test.in, 0)
		if err != nil {
			if !test.err {
				t.Error(err)
			}
		}
		if v != test.out {
			t.Error("For", test.in, "expected", test.out, "got", v)
		}
	}
}

func TestRightChild(t *testing.T) {
	var tests = []childTest{
		{0, 0, true},
		{1, 2, false},
		{3, 5, false},
	}

	for _, test := range tests {
		v, err := flattree.RightChild(test.in, 0)
		if err != nil {
			if !test.err {
				t.Error(err)
			}
		}
		if v != test.out {
			t.Error("For", test.in, "expected", test.out, "got", v)
		}
	}
}

func TestSibling(t *testing.T) {
	var tests = []testpair{
		{[]uint64{0, 0}, 2},
		{[]uint64{2, 0}, 0},
		{[]uint64{1, 0}, 5},
		{[]uint64{5, 0}, 1},
	}

	for _, pair := range tests {
		v := flattree.Sibling(pair.in[0], pair.in[1])
		if v != pair.out {
			t.Error("For", pair.in, "expected", pair.out, "got", v)
		}
	}
}

func TestFullRoots(t *testing.T) {
	var tests = []testpairArray{
		{0, []uint64{}},
		{2, []uint64{0}},
		{8, []uint64{3}},
		{20, []uint64{7, 17}},
		{18, []uint64{7, 16}},
		{16, []uint64{7}},
	}

	for _, pair := range tests {
		v, err := flattree.FullRoots(pair.in, []uint64{})
		if err != nil {
			t.Error(err)
		}
		if !reflect.DeepEqual(v, pair.out) {
			t.Error("For", pair.in, "expected", pair.out, "got", v)
		}
	}
}

func TestDepths(t *testing.T) {
	var tests = [][]uint64{
		{0, 0},
		{1, 1},
		{2, 0},
		{3, 2},
		{4, 0},
	}

	for _, pair := range tests {
		v := flattree.Depth(pair[0])
		if v != pair[1] {
			t.Error("For", pair[0], "expected", pair[1], "got", v)
		}
	}
}

func TestOffsets(t *testing.T) {
	var tests = [][]uint64{
		{0, 0},
		{1, 0},
		{2, 1},
		{3, 0},
		{4, 2},
	}

	for _, pair := range tests {
		v := flattree.Offset(pair[0], 0)
		if v != pair[1] {
			t.Error("For", pair[0], "expected", pair[1], "got", v)
		}
	}
}

type testSpans struct {
	in  uint64
	out flattree.Range
}

func TestSpans(t *testing.T) {
	var tests = []testSpans{
		{0, flattree.Range{0, 0}},
		{1, flattree.Range{0, 2}},
		{3, flattree.Range{0, 6}},
		{23, flattree.Range{16, 30}},
		{27, flattree.Range{24, 30}},
	}

	for _, pair := range tests {
		v := flattree.Spans(pair.in, 0)
		if !reflect.DeepEqual(v, pair.out) {
			t.Error("For", pair.in, "expected", pair.out, "got", v)
		}
	}
}

type spanTest struct {
	in     uint64
	expect uint64
}

func TestLeftSpan(t *testing.T) {
	var tests = []spanTest{
		{0, 0},
		{1, 0},
		{3, 0},
		{23, 16},
		{27, 24},
	}

	for _, pair := range tests {
		v := flattree.LeftSpan(pair.in, 0)
		if !reflect.DeepEqual(v, pair.expect) {
			t.Error("For", pair.in, "expected", pair.expect, "got", v)
		}
	}
}

func TestRightSpan(t *testing.T) {
	var tests = []spanTest{
		{0, 0},
		{1, 2},
		{3, 6},
		{23, 30},
		{27, 30},
	}

	for _, pair := range tests {
		v := flattree.RightSpan(pair.in, 0)
		if !reflect.DeepEqual(v, pair.expect) {
			t.Error("For", pair.in, "expected", pair.expect, "got", v)
		}
	}
}

type countTest struct {
	in     uint64
	expect uint64
}

func TestCount(t *testing.T) {
	tests := []countTest{
		{0, 1},
		{1, 3},
		{3, 7},
		{5, 3},
		{23, 15},
		{27, 7},
	}

	for _, pair := range tests {
		v := flattree.Count(pair.in, 0)
		if !reflect.DeepEqual(v, pair.expect) {
			t.Error("For", pair.in, "expected", pair.expect, "got", v)
		}
	}
}

type parentTest struct {
	in     uint64
	expect uint64
}

func TestParent(t *testing.T) {
	tests := []parentTest{
		{10000000000, 10000000001},
	}

	for _, pair := range tests {
		v := flattree.Parent(pair.in, 0)
		if !reflect.DeepEqual(v, pair.expect) {
			t.Error("For", pair.in, "expected", pair.expect, "got", v)
		}
	}
}

func TestParent2Child(t *testing.T) {
	var child uint64
	for i := 0; i < 50; i++ {
		child = flattree.Parent(child, 0)
	}
	var expect uint64 = 1125899906842623
	if child != expect {
		t.Error("Expected", expect, "got", child)
	}
	for j := 0; j < 50; j++ {
		child, _ = flattree.LeftChild(child, 0)
	}
	expect = 0
	if child != expect {
		t.Error("Expected", expect, "got", child)
	}
}

func TestIterator(t *testing.T) {
	itr := flattree.NewIterator(0)
	var val, expect uint64

	if val, expect = itr.Index, 0; val != expect {
		t.Error("Expected", 0, "got", val)
	}

	if val, expect = itr.Parent(), 1; val != expect {
		t.Error("Expected", expect, "got", val)
	}

	if val, expect = itr.Parent(), 3; val != expect {
		t.Error("Expected", expect, "got", val)
	}

	if val, expect = itr.Parent(), 7; val != expect {
		t.Error("Expected", expect, "got", val)
	}

	if val, expect = itr.RightChild(), 11; val != expect {
		t.Error("Expected", expect, "got", val)
	}

	if val, expect = itr.LeftChild(), 9; val != expect {
		t.Error("Expected", expect, "got", val)
	}

	if val, expect = itr.Next(), 13; val != expect {
		t.Error("Expected", expect, "got", val)
	}

	if val, expect = itr.LeftSpan(), 12; val != expect {
		t.Error("Expected", expect, "got", val)
	}
}

func TestIteratorNonLeafStart(t *testing.T) {
	itr := flattree.NewIterator(1)
	var val, expect uint64

	if val, expect = itr.Index, 1; val != expect {
		t.Error("Expected", 0, "got", itr.Index)
	}

	if val, expect = itr.Parent(), 3; val != expect {
		t.Error("Expected", expect, "got", val)
	}

	if val, expect = itr.Parent(), 7; val != expect {
		t.Error("Expected", expect, "got", val)
	}

	if val, expect = itr.RightChild(), 11; val != expect {
		t.Error("Expected", expect, "got", val)
	}

	if val, expect = itr.LeftChild(), 9; val != expect {
		t.Error("Expected", expect, "got", val)
	}

	if val, expect = itr.Next(), 13; val != expect {
		t.Error("Expected", expect, "got", val)
	}

	if val, expect = itr.LeftSpan(), 12; val != expect {
		t.Error("Expected", expect, "got", val)
	}
}
