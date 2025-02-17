package slices_test

import (
	"reflect"
	"sort"
	"strings"
	"testing"

	"github.com/twharmon/slices"
)

func TestPush(t *testing.T) {
	s := []int{1}
	s = slices.Append(s, 5)
	want := []int{1, 5}
	if !reflect.DeepEqual(s, want) {
		t.Fatalf("want %v; got %v", want, s)
	}
}

func TestContainsTrue(t *testing.T) {
	s := []int{1, 2, 3}
	got := slices.Contains(s, 2)
	want := true
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want %v; got %v", want, got)
	}
}

func TestContainsFalse(t *testing.T) {
	s := []int{1, 2, 3}
	got := slices.Contains(s, 4)
	want := false
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want %v; got %v", want, got)
	}
}

func TestSpliceNoInserts(t *testing.T) {
	s := []string{"foo", "bar", "baz"}
	got := slices.Splice(s, 1, 1)
	want := []string{"foo", "baz"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %v; got %v", want, got)
	}
}

func TestSpliceWithInserts(t *testing.T) {
	s := []string{"foo", "bar", "baz"}
	got := slices.Splice(s, 1, 1, "boo")
	want := []string{"foo", "boo", "baz"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %v; got %v", want, got)
	}
}

func TestReverseEven(t *testing.T) {
	s := []string{"foo", "bar"}
	got := slices.Reverse(s)
	want := []string{"bar", "foo"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %v; got %v", want, got)
	}
}

func TestReverseOdd(t *testing.T) {
	s := []string{"foo", "bar", "baz"}
	got := slices.Reverse(s)
	want := []string{"baz", "bar", "foo"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %v; got %v", want, got)
	}
}

func TestUnshift(t *testing.T) {
	s := []string{"foo"}
	got := slices.Unshift(s, "bar")
	want := []string{"bar", "foo"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %v; got %v", want, got)
	}
}

func TestFindFound(t *testing.T) {
	s := []string{"foo"}
	found := slices.Find(s, func(item string) bool {
		return item == "foo"
	})
	if found != "foo" {
		t.Fatalf("wrong find %s", found)
	}
}

func TestFindNotFound(t *testing.T) {
	s := []string{"foo"}
	found := slices.Find(s, func(item string) bool {
		return item == "bar"
	})
	if found != "" {
		t.Fatalf("wrong find %s", found)
	}
}

func TestFilter(t *testing.T) {
	s := []string{"foo", "bar"}
	ns := slices.Filter(s, func(item string) bool {
		return strings.HasPrefix(item, "f")
	})
	if len(ns) != 1 {
		t.Fatalf("wrong len %d", len(ns))
	}
}

func TestMaxZeroValue(t *testing.T) {
	s := []string{}
	got := slices.Max(s)
	want := ""
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %v; got %v", want, got)
	}
}
func TestMinZeroValue(t *testing.T) {
	s := []string{}
	got := slices.Min(s)
	want := ""
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %v; got %v", want, got)
	}
}

func TestMax(t *testing.T) {
	s := []int{2, 6, 1, 4, 3}
	got := slices.Max(s)
	want := 6
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %v; got %v", want, got)
	}
}

func TestMin(t *testing.T) {
	s := []int{2, 6, 1, 4, 3}
	got := slices.Min(s)
	want := 1
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %v; got %v", want, got)
	}
}

func TestMaxFuncZeroValue(t *testing.T) {
	s := []int{}
	got := slices.MaxFunc(s, func(a, b int) bool { return a < b })
	want := 0
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %v; got %v", want, got)
	}
}

func TestMinFuncZeroValue(t *testing.T) {
	s := []int{}
	got := slices.MinFunc(s, func(a, b int) bool { return a < b })
	want := 0
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %v; got %v", want, got)
	}
}

func TestMaxFunc(t *testing.T) {
	s := []int{2, 6, 1, 4, 3}
	got := slices.MaxFunc(s, func(a, b int) bool { return a < b })
	want := 6
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %v; got %v", want, got)
	}
}

func TestMinFunc(t *testing.T) {
	s := []int{2, 6, 1, 4, 3}
	got := slices.MinFunc(s, func(a, b int) bool { return a < b })
	want := 1
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %v; got %v", want, got)
	}
}

func TestSomeTrue(t *testing.T) {
	s := []string{"foo", "bar", "baz"}
	if !slices.Some(s, func(s string) bool { return s == "bar" }) {
		t.Fatalf("some was false; expected true")
	}
}

func TestSomeFalse(t *testing.T) {
	s := []string{"foo", "bar", "baz"}
	if slices.Some(s, func(s string) bool { return s == "x" }) {
		t.Fatalf("some was true; expected false")
	}
}

func TestEveryFalse(t *testing.T) {
	s := []string{"foo", "bar", "baz"}
	if slices.Every(s, func(s string) bool { return len(s) < 2 }) {
		t.Fatalf("every was true; expected false")
	}
}

func TestEveryTrue(t *testing.T) {
	s := []string{"foo", "bar", "baz"}
	if !slices.Every(s, func(s string) bool { return len(s) == 3 }) {
		t.Fatalf("every was true; expected false")
	}
}

func TestMap(t *testing.T) {
	s := []string{"f", "ba", "baz"}
	got := slices.Map(s, func(i string) int {
		return len(i)
	})
	want := []int{1, 2, 3}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %v; got %v", want, got)
	}
}

func TestReduce(t *testing.T) {
	s := []string{"f", "ba", "baz"}
	cnt := slices.Reduce(s, func(cnt int, i string) int {
		return cnt + len(i)
	})
	if cnt != 6 {
		t.Fatalf("%d != %d", cnt, 6)
	}
}

var unsorted = []string{"foo", "bar", "baz", "lorem", "ipsum", "donor", "sit", "alpha", "beta", "delta", "gamma", "omega", "epsilon", "mu", "nu", "lambda", "upsilon", "zeta", "eta", "rho", "psi", "iota", "apple", "banana", "pomegranate", "orange", "kiwi", "carrot", "broccoli"}

func TestSortFunc(t *testing.T) {
	s := slices.Clone(unsorted)
	got := slices.SortFunc(s, func(a string, b string) bool {
		return a < b
	})
	want := slices.Clone(s)
	sort.Slice(want, func(i, j int) bool {
		return want[i] < want[j]
	})
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %v; got %v", want, got)
	}
}

func TestSort(t *testing.T) {
	s := slices.Clone(unsorted)
	got := slices.Sort(s)
	want := slices.Clone(s)
	sort.Strings(want)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %v; got %v", want, got)
	}
}

func TestIntersection(t *testing.T) {
	a := []string{"foo", "bar"}
	b := []string{"bar", "baz"}
	want := []string{"bar"}
	got := slices.Intersection(a, b)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %v; got %v", want, got)
	}
}

func TestUnion(t *testing.T) {
	a := []string{"foo", "bar"}
	b := []string{"bar", "baz"}
	want := slices.Sort([]string{"foo", "bar", "baz"})
	got := slices.Sort(slices.Union(a, b))
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %v; got %v", want, got)
	}
}

func BenchmarkStdLibSortFunc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := slices.Clone(unsorted)
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})
	}
}

func BenchmarkSortFuncShort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := slices.Clone(unsorted)
		_ = slices.SortFunc(s, func(a string, b string) bool {
			return a < b
		})
	}
}

func BenchmarkStdLibSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := slices.Clone(unsorted)
		sort.Strings(s)
	}
}

func BenchmarkSortShort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := slices.Clone(unsorted)
		_ = slices.Sort(s)
	}
}
