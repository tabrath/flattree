package flattree

import (
	"reflect"
	"testing"
)

func TestIndex(t *testing.T) {
	if index := Index(0, 0); index != 0 {
		t.Errorf("Index failed; expected %d, got %d", 0, index)
	}
	if index := Index(0, 1); index != 2 {
		t.Errorf("Index failed; expected %d, got %d", 2, index)
	}
	if index := Index(0, 2); index != 4 {
		t.Errorf("Index failed; expected %d, got %d", 4, index)
	}
}

func TestParents(t *testing.T) {
	if index := Index(1, 0); index != 1 {
		t.Errorf("Index failed; expected %d, got %d", 1, index)
	}
	if index := Index(1, 1); index != 5 {
		t.Errorf("Index failed; expected %d, got %d", 5, index)
	}
	if index := Index(2, 0); index != 3 {
		t.Errorf("Index failed; expected %d, got %d", 3, index)
	}

	if parent := Parent(0); parent != 1 {
		t.Errorf("Parent failed; expected %d, got %d", 1, parent)
	}
	if parent := Parent(2); parent != 1 {
		t.Errorf("Parent failed; expected %d, got %d", 1, parent)
	}
	if parent := Parent(1); parent != 3 {
		t.Errorf("Parent failed; expected %d, got %d", 3, parent)
	}
}

func TestChildren(t *testing.T) {
	if children, _ := Children(0, 0); len(children) > 0 {
		t.Errorf("Children failed; expected nil, got %x", children)
	}
	if children, _ := Children(1, 0); !reflect.DeepEqual(children, []uint{0, 2}) {
		t.Errorf("Children failed; expected {0,2}, got %x", children)
	}
	if children, _ := Children(3, 0); !reflect.DeepEqual(children, []uint{1, 5}) {
		t.Errorf("Children failed; expected {1,5}, got %x", children)
	}
}

func TestLeftChild(t *testing.T) {
	if _, err := LeftChild(0, 0); err == nil {
		t.Errorf("LeftChild failed; expected error, got nil")
	}
	if child, _ := LeftChild(1, 0); child != 0 {
		t.Errorf("LeftChild failed; expected 0, got %x", child)
	}
	if child, _ := LeftChild(3, 0); child != 1 {
		t.Errorf("LeftChild failed; expected 1, got %x", child)
	}
}

func TestRightChild(t *testing.T) {
	if _, err := RightChild(0, 0); err == nil {
		t.Errorf("RightChild failed; expected error, got nil")
	}
	if child, _ := RightChild(1, 0); child != 2 {
		t.Errorf("RightChild failed; expected 2, got %x", child)
	}
	if child, _ := RightChild(3, 0); child != 5 {
		t.Errorf("RightChild failed; expected 5, got %x", child)
	}
}

func TestSiblings(t *testing.T) {
	if sibling := Sibling(0); sibling != 2 {
		t.Errorf("Sibling failed; expected 2, got %x", sibling)
	}
	if sibling := Sibling(2); sibling != 0 {
		t.Errorf("Sibling failed; expected 0, got %x", sibling)
	}
	if sibling := Sibling(1); sibling != 5 {
		t.Errorf("Sibling failed; expected 5, got %x", sibling)
	}
	if sibling := Sibling(5); sibling != 1 {
		t.Errorf("Sibling failed; expected 1, got %x", sibling)
	}
}

func TestFullRoots(t *testing.T) {
	if roots, _ := FullRoots(0); len(roots) > 0 {
		t.Errorf("FullRoots failed; expected nil, got %x", roots)
	}
	if roots, _ := FullRoots(2); !reflect.DeepEqual(roots, []uint{0}) {
		t.Errorf("FullRoots failed; expected {0}, got %x", roots)
	}
	if roots, _ := FullRoots(8); !reflect.DeepEqual(roots, []uint{3}) {
		t.Errorf("FullRoots failed; expected {3}, got %x", roots)
	}
	if roots, _ := FullRoots(20); !reflect.DeepEqual(roots, []uint{7, 17}) {
		t.Errorf("FullRoots failed; expected {7,17}, got %x", roots)
	}
	if roots, _ := FullRoots(18); !reflect.DeepEqual(roots, []uint{7, 16}) {
		t.Errorf("FullRoots failed; expected {7,16}, got %x", roots)
	}
	if roots, _ := FullRoots(16); !reflect.DeepEqual(roots, []uint{7}) {
		t.Errorf("FullRoots failed; expected {7}, got %x", roots)
	}
}

func TestDepth(t *testing.T) {
	if depth := Depth(0); depth != 0 {
		t.Errorf("Depth failed; expected 0, got %x", depth)
	}
	if depth := Depth(1); depth != 1 {
		t.Errorf("Depth failed; expected 1, got %x", depth)
	}
	if depth := Depth(2); depth != 0 {
		t.Errorf("Depth failed; expected 0, got %x", depth)
	}
	if depth := Depth(3); depth != 2 {
		t.Errorf("Depth failed; expected 2, got %x", depth)
	}
	if depth := Depth(4); depth != 0 {
		t.Errorf("Depth failed; expected 0, got %x", depth)
	}
}

func TestOffset(t *testing.T) {
	if offset := Offset(0, 0); offset != 0 {
		t.Errorf("Offset failed; expected 0, got %x", offset)
	}
	if offset := Offset(1, 0); offset != 0 {
		t.Errorf("Offset failed; expected 0, got %x", offset)
	}
	if offset := Offset(2, 0); offset != 1 {
		t.Errorf("Offset failed; expected 1, got %x", offset)
	}
	if offset := Offset(3, 0); offset != 0 {
		t.Errorf("Offset failed; expected 0, got %x", offset)
	}
	if offset := Offset(4, 0); offset != 2 {
		t.Errorf("Offset failed; expected 2, got %x", offset)
	}
}

func TestSpans(t *testing.T) {
	if spans := Spans(0, 0); !reflect.DeepEqual(spans, []uint{0, 0}) {
		t.Errorf("Spans failed; expected {0,0}, got %x", spans)
	}
	if spans := Spans(1, 0); !reflect.DeepEqual(spans, []uint{0, 2}) {
		t.Errorf("Spans failed; expected {0,2}, got %x", spans)
	}
	if spans := Spans(3, 0); !reflect.DeepEqual(spans, []uint{0, 6}) {
		t.Errorf("Spans failed; expected {0,6}, got %x", spans)
	}
	if spans := Spans(23, 0); !reflect.DeepEqual(spans, []uint{16, 30}) {
		t.Errorf("Spans failed; expected {16,30}, got %x", spans)
	}
	if spans := Spans(27, 0); !reflect.DeepEqual(spans, []uint{24, 30}) {
		t.Errorf("Spans failed; expected {24,30}, got %x", spans)
	}
}

func TestLeftSpan(t *testing.T) {
	if span := LeftSpan(0, 0); span != 0 {
		t.Errorf("LeftSpan failed; expected 0, got %x", span)
	}
	if span := LeftSpan(1, 0); span != 0 {
		t.Errorf("LeftSpan failed; expected 0, got %x", span)
	}
	if span := LeftSpan(3, 0); span != 0 {
		t.Errorf("LeftSpan failed; expected 0, got %x", span)
	}
	if span := LeftSpan(23, 0); span != 16 {
		t.Errorf("LeftSpan failed; expected 16, got %x", span)
	}
	if span := LeftSpan(27, 0); span != 24 {
		t.Errorf("LeftSpan failed; expected 24, got %x", span)
	}
}

func TestRightSpan(t *testing.T) {
	if span := RightSpan(0, 0); span != 0 {
		t.Errorf("RightSpan failed; expected 0, got %x", span)
	}
	if span := RightSpan(1, 0); span != 2 {
		t.Errorf("RightSpan failed; expected 2, got %x", span)
	}
	if span := RightSpan(3, 0); span != 6 {
		t.Errorf("RightSpan failed; expected 6, got %x", span)
	}
	if span := RightSpan(23, 0); span != 30 {
		t.Errorf("RightSpan failed; expected 30, got %x", span)
	}
	if span := RightSpan(27, 0); span != 30 {
		t.Errorf("RightSpan failed; expected 30, got %x", span)
	}
}

func TestCount(t *testing.T) {
	if count := Count(0, 0); count != 1 {
		t.Errorf("Count failed; expected 1, got %x", count)
	}
	if count := Count(1, 0); count != 3 {
		t.Errorf("Count failed; expected 3, got %x", count)
	}
	if count := Count(3, 0); count != 7 {
		t.Errorf("Count failed; expected 7, got %x", count)
	}
	if count := Count(5, 0); count != 3 {
		t.Errorf("Count failed; expected 3, got %x", count)
	}
	if count := Count(23, 0); count != 15 {
		t.Errorf("Count failed; expected 15, got %x", count)
	}
	if count := Count(27, 0); count != 7 {
		t.Errorf("Count failed; expected 7, got %x", count)
	}
}

func TestParentGreaterThanInt32(t *testing.T) {
	if parent := Parent(10000000000); parent != 10000000001 {
		t.Errorf("Parent failed; not greater than int32")
	}
}

func TestChildToParentToChild(t *testing.T) {
	var child uint
	for i := 0; i < 50; i++ {
		child = Parent(child)
	}
	if child != 1125899906842623 {
		t.Errorf("ChildToParentToChild failed; expected 1125899906842623, got %x", child)
	}
	for j := 0; j < 50; j++ {
		child, _ = LeftChild(child, 0)
	}
	if child != 0 {
		t.Errorf("ChildToParentToChild failed; expected 0, got %x", child)
	}
}
