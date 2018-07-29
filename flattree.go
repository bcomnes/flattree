package flattree

import "errors"

// Range type is an inclusive tree range that is in the form of [start, finish]
type Range []uint64

// FullRoots returns a list of all the full roots (subtrees where all nodes have either 2 or 0 children) `<` index.
// For example `fullRoots(8)` returns `[3]` since the subtree rooted at `3` spans `0 -> 6` and the tree
// rooted at `7` has a child located at `9` which is `>= 8`.
func FullRoots(index uint64) []uint64 {
	var result []uint64

	if (index & 1) != 0 {
		return result
	}

	index /= 2
	var offset uint64
	var factor uint64 = 1

	for true {
		if index == 0 {
			return result, nil
		}
		for (factor * 2) <= index {
			factor *= 2
		}
		result = append(result, offset+factor-1)
		offset = offset + 2*factor
		index -= factor
		factor = 1
	}
	return nil, nil
}

// Depth returns the depth of an element.
func Depth(index uint64) uint64 {
	var depth uint64

	index++
	for (index & 1) == 0 {
		depth++
		index = rightShift(index)
	}

	return depth
}

// Sibling returns the index of this elements sibling.
func Sibling(index, depth uint64) uint64 {
	if depth == 0 {
		depth = Depth(index)
	}
	offset := Offset(index, depth)
	if offset&1 == 1 {
		return Index(depth, offset-1)
	}
	return Index(depth, offset+1)
}

// Parent returns the index of the parent element in tree.
func Parent(index, depth uint64) uint64 {
	if depth == 0 {
		depth = Depth(index)
	}
	offset := Offset(index, depth)

	return Index(depth+1, rightShift(offset))
}

// LeftChild returns the left most child at index and depth.
func LeftChild(index, depth uint64) (uint64, error) {
	if (index & 1) == 0 {
		return 0, errors.New("No more left children")
	}
	if depth == 0 {
		depth = Depth(index)
	}
	return Index(depth-1, Offset(index, depth)*2), nil
}

// RightChild returns the right most child at index and depth.
func RightChild(index, depth uint64) (uint64, error) {
	if (index & 1) == 0 {
		return 0, errors.New("No more right children")
	}
	if depth == 0 {
		depth = Depth(index)
	}
	return Index(depth-1, 1+Offset(index, depth)*2), nil
}

// Children returns a Range slice `[leftChild, rightChild]` with indexes of this elements children.
// If this element does not have any children it returns an empty slice.
func Children(index, depth uint64) Range {
	if (index & 1) == 0 {
		return Range{}
	}

	if depth == 0 {
		depth = Depth(index)
	}
	offset := Offset(index, depth)

	return Range{Index(depth-1, offset), Index(depth-1, offset+1)}
}

// LeftSpan returns the left spanning in index in the tree `index` spans.
func LeftSpan(index, depth uint64) uint64 {
	if (index & 1) == 0 {
		return index
	}
	if depth == 0 {
		depth = Depth(index)
	}
	return Offset(index, depth) * twoPow(depth+1)
}

// RightSpan returns the right spanning in index in the tree `index` spans.
func RightSpan(index, depth uint64) uint64 {
	if (index & 1) == 0 {
		return index
	}
	if depth == 0 {
		depth = Depth(index)
	}
	return (Offset(index, depth)+1)*twoPow(depth+1) - 2
}

// Count returns how many nodes (including the parent nodes) a tree contains.
func Count(index, depth uint64) uint64 {
	if (index & 1) == 0 {
		return 1
	}
	if depth == 0 {
		depth = Depth(index)
	}
	return twoPow(depth+1) - 1
}

// Spans returns the Range (inclusive) the tree root at `index` spans.
// For example `Spans(3)` would return `[]Range{0, 6}`
func Spans(index, depth uint64) Range {
	if (index & 1) == 0 {
		return Range{index, index}
	}
	if depth == 0 {
		depth = Depth(index)
	}

	offset := Offset(index, depth)
	width := twoPow(depth + 1)
	return Range{offset * width, (offset+1)*width - 2}
}

// Index returns an array index for the tree element at the given depth and offset.
func Index(depth, offset uint64) uint64 {
	return (1+2*offset)*twoPow(depth) - 1
}

// Offset returns the relative offset of an element.
func Offset(index, depth uint64) uint64 {
	if (index & 1) == 0 {
		return index / 2
	}
	if depth == 0 {
		depth = Depth(index)
	}
	return ((index+1)/twoPow(depth) - 1) / 2
}

func twoPow(n uint64) uint64 {
	if n < 31 {
		return 1 << n
	}
	return ((1 << 30) * (1 << (n - 30)))
}

func rightShift(n uint64) uint64 {
	return (n - (n & 1)) / 2
}

// NewIterator returns a stateful tree iterator at a given index.
func NewIterator(index uint64) *Iterator {
	var ite Iterator
	ite.Seek(index)
	return &ite
}

// Iterator is a stateful iterator type.  Use NewIterator to create tree Iterators at a given index.
type Iterator struct {
	Index, Offset, Factor uint64
}

// Seek moves the iterator to this specific tree index.
func (i *Iterator) Seek(index uint64) {
	i.Index = index
	if i.Index&1 == 1 {
		i.Offset = Offset(index, 0)
		i.Factor = twoPow(Depth(index) + 1)
	} else {
		i.Offset = index / 2
		i.Factor = 2
	}
}

// IsLeft tests if the iterator at a left sibling.
func (i *Iterator) IsLeft() bool {
	if i.Offset&1 == 0 {
		return true
	}
	return false
}

// IsRight tests if the iterator is at a right subling.
func (i *Iterator) IsRight() bool {
	return !i.IsLeft()
}

// Prev moves the iterator to the prev item in the tree.
func (i *Iterator) Prev() uint64 {
	if i.Offset == 0 {
		return i.Index
	}
	i.Offset--
	i.Index -= i.Factor
	return i.Index
}

// Next moves the iterator to the next item in the tree.
func (i *Iterator) Next() uint64 {
	i.Offset++
	i.Index += i.Factor
	return i.Index
}

// Sibling moves the iterator to the current sibling.
func (i *Iterator) Sibling() uint64 {
	if i.IsLeft() {
		return i.Next()
	}
	return i.Prev()
}

// Parent moves the iterator to the current parent index.
func (i *Iterator) Parent() uint64 {
	if (i.Offset & 1) == 1 {
		i.Index -= i.Factor / 2
		i.Offset = (i.Offset - 1) / 2
	} else {
		i.Index += i.Factor / 2
		i.Offset /= 2
	}
	i.Factor *= 2
	return i.Index
}

// LeftSpan moves the iterator to the current left span index.
func (i *Iterator) LeftSpan() uint64 {
	i.Index = i.Index - i.Factor/2 + 1
	i.Offset = i.Index / 2
	i.Factor = 2
	return i.Index
}

// RightSpan moves the iterator to the current right span index.
func (i *Iterator) RightSpan() uint64 {
	i.Index = i.Index + i.Factor/2 - 1
	i.Offset = i.Index / 2
	i.Factor = 2
	return i.Index
}

// LeftChild moves the iterator to the current left child index.
func (i *Iterator) LeftChild() uint64 {
	if i.Factor == 2 {
		return i.Index
	}
	i.Factor /= 2
	i.Index -= i.Factor / 2
	i.Offset *= 2
	return i.Index
}

// RightChild moves the iterator to the current right child index.
func (i *Iterator) RightChild() uint64 {
	if i.Factor == 2 {
		return i.Index
	}
	i.Factor /= 2
	i.Index += i.Factor / 2
	i.Offset = 2*i.Offset + 1
	return i.Index
}
