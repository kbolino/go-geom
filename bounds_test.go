package geom

import (
	"reflect"
	"testing"
)

func TestBoundsExtend(t *testing.T) {
	for i, testData := range []struct {
		initial, expected Bounds
		expandBy          T
	}{
		{
			initial:  Bounds{layout: XY, min: Coord{0, 0}, max: Coord{0, 0}},
			expected: Bounds{layout: XY, min: Coord{0, -10}, max: Coord{10, 0}},
			expandBy: NewPointFlat(XY, []float64{10, -10}),
		},
		{
			initial:  Bounds{layout: XY, min: Coord{-100, -100}, max: Coord{100, 100}},
			expected: Bounds{layout: XY, min: Coord{-100, -100}, max: Coord{100, 100}},
			expandBy: NewPointFlat(XY, []float64{10, -10}),
		},
		{
			initial:  Bounds{layout: XYZ, min: Coord{0, 0, -1}, max: Coord{10, 10, 1}},
			expected: Bounds{layout: XYZ, min: Coord{0, -10, -1}, max: Coord{10, 10, 1}},
			expandBy: NewPointFlat(XY, []float64{5, -10}),
		},
		{
			initial:  Bounds{layout: XYZ, min: Coord{0, 0, 0}, max: Coord{10, 10, 10}},
			expected: Bounds{layout: XYZ, min: Coord{0, -10, 0}, max: Coord{10, 10, 10}},
			expandBy: NewPolygonFlat(XYZ, []float64{5, -10, 3}, []int{2}),
		},
	} {
		extended := Bounds{layout: testData.initial.layout, min: testData.initial.min, max: testData.initial.max}
		extended.Extend(testData.expandBy)

		if !reflect.DeepEqual(extended, testData.expected) {
			t.Errorf("Test %v Failed.  Expected: \n%v but got: \n%v", i+1, testData.expected, extended)
		}
	}
}

func TestBoundsIsEmpty(t *testing.T) {
	for i, testData := range []struct {
		bounds  Bounds
		isEmpty bool
	}{
		{
			bounds:  Bounds{layout: XY, min: Coord{0, 0}, max: Coord{-1, -1}},
			isEmpty: true,
		}, {
			bounds:  Bounds{layout: XY, min: Coord{0, 0}, max: Coord{0, 0}},
			isEmpty: false,
		},
		{
			bounds:  Bounds{layout: XY, min: Coord{-100, -100}, max: Coord{100, 100}},
			isEmpty: false,
		},
	} {
		copy := Bounds{layout: testData.bounds.layout, min: testData.bounds.min, max: testData.bounds.max}

		for j := 0; j < 10; j++ {
			// do multiple checks to verify no obvious side effects are caused
			isEmpty := copy.IsEmpty()
			if isEmpty != testData.isEmpty {
				t.Errorf("Test %v Failed.  Expected: %v but got: %v", i+1, testData.isEmpty, isEmpty)
				break
			}
			if !reflect.DeepEqual(copy, testData.bounds) {
				t.Errorf("Test %v Failed.  Function IsEmpty modified internal state of bounds.  Before: \n%v After: \n%v", i+1, testData.bounds, copy)
				break
			}
		}

	}
}

func TestBoundsOverlaps(t *testing.T) {
	for i, testData := range []struct {
		bounds, other Bounds
		overlaps      bool
	}{
		{
			bounds:   Bounds{layout: XY, min: Coord{0, 0}, max: Coord{0, 0}},
			other:    Bounds{layout: XY, min: Coord{-10, 0}, max: Coord{-5, 10}},
			overlaps: false,
		},
		{
			bounds:   Bounds{layout: XY, min: Coord{-100, -100}, max: Coord{100, 100}},
			other:    Bounds{layout: XY, min: Coord{-10, 0}, max: Coord{-5, 10}},
			overlaps: true,
		},
		{
			bounds:   Bounds{layout: XY, min: Coord{1, 1}, max: Coord{5, 5}},
			other:    Bounds{layout: XY, min: Coord{-5, -5}, max: Coord{-1, -1}},
			overlaps: false,
		},
		{
			bounds:   Bounds{layout: XYZ, min: Coord{-100, -100, -100}, max: Coord{100, 100, 100}},
			other:    Bounds{layout: XYZ, min: Coord{-10, 0, 0}, max: Coord{-5, 10, 10}},
			overlaps: true,
		},
		{
			bounds:   Bounds{layout: XYZ, min: Coord{0, 0, 0}, max: Coord{100, 100, 100}},
			other:    Bounds{layout: XYZ, min: Coord{5, 5, -10}, max: Coord{10, 10, -5}},
			overlaps: false,
		},
		{
			bounds:   Bounds{layout: XY, min: Coord{0, 0}, max: Coord{0, 0}},
			other:    Bounds{layout: XY, min: Coord{-10, -10}, max: Coord{-0.000000000000000000000000000001, 0}},
			overlaps: false,
		},
	} {
		copy := Bounds{layout: testData.bounds.layout, min: testData.bounds.min, max: testData.bounds.max}

		for j := 0; j < 10; j++ {
			// do multiple checks to verify no obvious side effects are caused
			overlaps := copy.Overlaps(testData.bounds.layout, &testData.other)
			if overlaps != testData.overlaps {
				t.Errorf("Test %v Failed.  Expected: %v but got: %v", i+1, testData.overlaps, overlaps)
				break
			}
			if !reflect.DeepEqual(copy, testData.bounds) {
				t.Errorf("Test %v Failed.  Function Overlaps modified internal state of bounds.  Before: \n%v After: \n%v", i+1, testData.bounds, copy)
				break
			}
		}

	}
}

func TestBoundsOverlapsPoint(t *testing.T) {
	for i, testData := range []struct {
		bounds   Bounds
		point    Coord
		overlaps bool
	}{
		{
			bounds:   Bounds{layout: XY, min: Coord{0, 0}, max: Coord{0, 0}},
			point:    Coord{-10, 0},
			overlaps: false,
		},
		{
			bounds:   Bounds{layout: XY, min: Coord{-100, -100}, max: Coord{100, 100}},
			point:    Coord{-10, 0},
			overlaps: true,
		},
		{
			bounds:   Bounds{layout: XYZ, min: Coord{-100, -100, -100}, max: Coord{100, 100, 100}},
			point:    Coord{-5, 10, 10},
			overlaps: true,
		},
		{
			bounds:   Bounds{layout: XYZ, min: Coord{0, 0, 0}, max: Coord{100, 100, 100}},
			point:    Coord{5, 5, -10},
			overlaps: false,
		},
		{
			bounds:   Bounds{layout: XY, min: Coord{0, 0}, max: Coord{10, 10}},
			point:    Coord{-0.000000000000000000000000000001, 0},
			overlaps: false,
		},
	} {
		copy := Bounds{layout: testData.bounds.layout, min: testData.bounds.min, max: testData.bounds.max}

		for j := 0; j < 10; j++ {
			// do multiple checks to verify no obvious side effects are caused
			overlaps := copy.OverlapsPoint(testData.bounds.layout, testData.point)
			if overlaps != testData.overlaps {
				t.Errorf("Test %v Failed.  Expected: %v but got: %v", i+1, testData.overlaps, overlaps)
				break
			}
			if !reflect.DeepEqual(copy, testData.bounds) {
				t.Errorf("Test %v Failed.  Function Overlaps modified internal state of bounds.  Before: \n%v After: \n%v", i+1, testData.bounds, copy)
				break
			}
		}
	}
}

func TestBoundsSet(t *testing.T) {
	bounds := Bounds{layout: XY, min: Coord{0, 0}, max: Coord{10, 10}}
	bounds.Set(0, 0, 20, 20)
	expected := Bounds{layout: XY, min: Coord{0, 0}, max: Coord{20, 20}}
	if !reflect.DeepEqual(expected, bounds) {
		t.Errorf("Expected %v but got %v", expected, bounds)
	}

	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Error("Expected a panic but didn't get it as expected")
			} else if !reflect.DeepEqual(expected, bounds) {
				t.Errorf("Set modified bounds even though error was thrown.", expected, bounds)
			}
		}()

		bounds.Set(2, 2, 2, 2, 2)
	}()
}

func TestBoundsSetCoords(t *testing.T) {
	bounds := &Bounds{layout: XY, min: Coord{0, 0}, max: Coord{10, 10}}
	bounds.SetCoords(Coord{0, 0}, Coord{20, 20})
	expected := Bounds{layout: XY, min: Coord{0, 0}, max: Coord{20, 20}}
	if !reflect.DeepEqual(expected, *bounds) {
		t.Errorf("Expected %v but got %v", expected, *bounds)
	}

	bounds = NewBounds(XY)
	bounds.SetCoords(Coord{0, 0}, Coord{20, 20})
	if !reflect.DeepEqual(expected, *bounds) {
		t.Errorf("Expected %v but got %v", expected, *bounds)
	}

	bounds = NewBounds(XY)
	bounds.SetCoords(Coord{20, 0}, Coord{0, 20}) // set coords should ensure valid min / max
	if !reflect.DeepEqual(expected, *bounds) {
		t.Errorf("Expected %v but got %v", expected, *bounds)
	}
}
