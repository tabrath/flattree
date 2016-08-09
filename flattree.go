/*
Package flattree implements a series of functions to map a binary tree to a list
*/
package flattree

import "errors"

// FullRoots returns a list of all the full roots (subtrees where all nodes have either 2 or 0 children) < index
func FullRoots(index uint) (result []uint, err error) {
	if index&1 != 0 {
		return nil, errors.New("You can only look up roots for depth(0) blocks")
	}

	result = make([]uint, 0)
	index /= 2

	var offset uint
	var factor uint = 1

	for index > 0 {
		for factor*2 <= index {
			factor *= 2
		}
		result = append(result, offset+factor-1)
		offset = offset + 2*factor
		index -= factor
		factor = 1
	}

	return
}

func rightShift(n uint) uint {
	return (n - (n & 1)) / 2
}

func twoPow(n uint) uint {
	switch {
	case n < 31:
		return 1 << n
	default:
		return ((1 << 30) * (1 << (n - 30)))
	}
}

// Depth returns the depth of an element
func Depth(index uint) (depth uint) {
	index++
	for (index & 1) == 0 {
		depth++
		index = rightShift(index)
	}
	return
}

// Sibling returns the index of this elements sibling
func Sibling(index uint) uint {
	depth := Depth(index)
	offset := Offset(index, depth)

	if offset&1 != 0 {
		return Index(depth, offset-1)
	}

	return Index(depth, offset+1)
}

// Index returns an array index for the tree element at the given depth and offset
func Index(depth, offset uint) uint {
	return (1+2*offset)*twoPow(depth) - 1
}

// Offset returns the relative offset of an element
func Offset(index, depth uint) uint {
	if index&1 == 0 {
		return index / 2
	}
	if depth == 0 {
		depth = Depth(index)
	}

	return ((index+1)/twoPow(depth) - 1) / 2
}

// Parent returns the index of the parent element in tree
func Parent(index uint) uint {
	depth := Depth(index)
	offset := Offset(index, depth)

	return Index(depth+1, rightShift(offset))
}

// LeftChild returns the index of the left child, or error if no child
func LeftChild(index, depth uint) (uint, error) {
	if index&1 == 0 {
		return 0, errors.New("No left child")
	}
	if depth == 0 {
		depth = Depth(index)
	}

	return Index(depth-1, Offset(index, depth)*2), nil
}

// RightChild returns the index of the right child, or error if no child
func RightChild(index, depth uint) (uint, error) {
	if index&1 == 0 {
		return 0, errors.New("No right child")
	}
	if depth == 0 {
		depth = Depth(index)
	}

	return Index(depth-1, 1+(Offset(index, depth)*2)), nil
}

// Children returns an array [leftChild, rightChild] with the indexes of this elements children. If this element does not have any children it returns nil and an error
func Children(index, depth uint) ([]uint, error) {
	if index&1 == 0 {
		return nil, errors.New("No children")
	}

	if depth == 0 {
		depth = Depth(index)
	}
	offset := Offset(index, depth) * 2

	return []uint{Index(depth-1, offset), Index(depth-1, offset+1)}, nil
}

// LeftSpan returns the range the index spans to the left
func LeftSpan(index, depth uint) uint {
	if index&1 == 0 {
		return index
	}
	if depth == 0 {
		depth = Depth(index)
	}

	return Offset(index, depth) * twoPow(depth+1)
}

// RightSpan returns the range the index spans to the right
func RightSpan(index, depth uint) uint {
	if index&1 == 0 {
		return index
	}
	if depth == 0 {
		depth = Depth(index)
	}

	return (Offset(index, depth)+1)*twoPow(depth+1) - 2
}

// Count returns how many nodes (including parent nodes) a tree contains
func Count(index, depth uint) uint {
	if index&1 == 0 {
		return 1
	}
	if depth == 0 {
		depth = Depth(index)
	}

	return twoPow(depth+1) - 1
}

// Spans returns the range (inclusive) the tree root at index spans
func Spans(index, depth uint) []uint {
	if index&1 == 0 {
		return []uint{index, index}
	}
	if depth == 0 {
		depth = Depth(index)
	}

	offset := Offset(index, depth)
	width := twoPow(depth + 1)

	return []uint{offset * width, (offset+1)*width - 2}
}
