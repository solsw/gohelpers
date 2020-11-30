package chanhelper

import (
	"reflect"
	"sort"
	"testing"
)

func TestPeekInt(t *testing.T) {
	ch0 := make(chan int)
	defer close(ch0)
	ch1 := make(chan int, 1)
	ch1 <- 1
	defer close(ch1)
	ch2 := make(chan int, 2)
	ch2 <- 1
	ch2 <- 2
	close(ch2)

	type args struct {
		ch <-chan int
	}
	tests := []struct {
		name     string
		args     args
		wantV    int
		wantOk   bool
		wantOpen bool
	}{
		{name: "0", args: args{ch: ch0}, wantV: 0, wantOk: false, wantOpen: true},
		{name: "11", args: args{ch: ch1}, wantV: 1, wantOk: true, wantOpen: true},
		{name: "12", args: args{ch: ch1}, wantV: 0, wantOk: false, wantOpen: true},
		{name: "21", args: args{ch: ch2}, wantV: 1, wantOk: true, wantOpen: true},
		{name: "22", args: args{ch: ch2}, wantV: 2, wantOk: true, wantOpen: true},
		{name: "23", args: args{ch: ch2}, wantV: 0, wantOk: false, wantOpen: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotV, gotOk, gotOpen := PeekInt(tt.args.ch)
			if gotV != tt.wantV {
				t.Errorf("PeekInt() gotV = %v, want %v", gotV, tt.wantV)
			}
			if gotOk != tt.wantOk {
				t.Errorf("PeekInt() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
			if gotOpen != tt.wantOpen {
				t.Errorf("PeekInt() gotOpen = %v, want %v", gotOpen, tt.wantOpen)
			}
		})
	}
}

func TestMerge2Int(t *testing.T) {
	in1 := make(chan int)
	in2 := make(chan int)
	go func() {
		for i := 0; i < 4; i++ {
			in1 <- 1
		}
		close(in1)
	}()
	go func() {
		for i := 0; i < 4; i++ {
			in2 <- 2
		}
		close(in2)
	}()
	var out []int
	for v := range Merge2Int(in1, in2) {
		out = append(out, v)
	}
	sort.Ints(out)
	want := []int{1, 1, 1, 1, 2, 2, 2, 2}
	if !reflect.DeepEqual(out, want) {
		t.Errorf("Merge2Int() = %v, want %v", out, want)
	}
}

func TestMergeInt0(t *testing.T) {
	_, ok := <-MergeInt()
	want := false
	if ok != want {
		t.Errorf("MergeInt0() = %v, want %v", ok, want)
	}
}

func TestMergeInt1(t *testing.T) {
	in1 := make(chan int)
	go func() {
		for i := 0; i < 4; i++ {
			in1 <- i
		}
		close(in1)
	}()
	var out []int
	for v := range MergeInt(in1) {
		out = append(out, v)
	}
	sort.Ints(out)
	want := []int{0, 1, 2, 3}
	if !reflect.DeepEqual(out, want) {
		t.Errorf("MergeInt1() = %v, want %v", out, want)
	}
}

func TestMergeInt(t *testing.T) {
	type args struct {
		ins []chan int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "2",
			args: args{ins: func() []chan int {
				cc := []chan int{make(chan int), make(chan int)}
				for i, c := range cc {
					go func(i int, c chan<- int) {
						for n := 0; n < 4; n++ {
							c <- i
						}
						close(c)
					}(i, c)
				}
				return cc
			}(),
			},
			want: []int{0, 0, 0, 0, 1, 1, 1, 1},
		},
		{name: "3",
			args: args{ins: func() []chan int {
				cc := []chan int{make(chan int), make(chan int), make(chan int)}
				for i, c := range cc {
					go func(i int, c chan<- int) {
						for n := 0; n < 2; n++ {
							c <- i
						}
						close(c)
					}(i, c)
				}
				return cc
			}(),
			},
			want: []int{0, 0, 1, 1, 2, 2},
		},
		{name: "4",
			args: args{ins: func() []chan int {
				cc := []chan int{make(chan int), make(chan int), make(chan int), make(chan int)}
				for i, c := range cc {
					go func(i int, c chan<- int) {
						for n := 0; n < 2; n++ {
							c <- i
						}
						close(c)
					}(i, c)
				}
				return cc
			}(),
			},
			want: []int{0, 0, 1, 1, 2, 2, 3, 3},
		},
		{name: "5",
			args: args{ins: func() []chan int {
				cc := []chan int{make(chan int), make(chan int), make(chan int), make(chan int), make(chan int)}
				for i, c := range cc {
					go func(i int, c chan<- int) {
						for n := 0; n < 2; n++ {
							c <- i * i
						}
						close(c)
					}(i, c)
				}
				return cc
			}(),
			},
			want: []int{0, 0, 1, 1, 4, 4, 9, 9, 16, 16},
		},
		{name: "23",
			args: args{ins: func() []chan int {
				var cc []chan int
				for i := 0; i < 23; i++ {
					cc = append(cc, make(chan int))
				}
				for i, c := range cc {
					go func(i int, c chan<- int) {
						c <- i
						close(c)
					}(i, c)
				}
				return cc
			}(),
			},
			want: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []int
			for v := range MergeInt(tt.args.ins...) {
				got = append(got, v)
			}
			sort.Ints(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeBufInt0(t *testing.T) {
	cc := []chan int{}
	out := make([]int, 0)
	for v := range MergeBufInt(cc...) {
		out = append(out, v)
	}
	want := make([]int, 0)
	if !reflect.DeepEqual(out, want) {
		t.Errorf("MergeBufInt0() = %v, want %v", out, want)
	}
}

func TestMergeBufInt(t *testing.T) {
	cc := []chan int{make(chan int), make(chan int), make(chan int), make(chan int)}
	for i, c := range cc {
		go func(i int, c chan<- int) {
			for n := 0; n < 2; n++ {
				c <- i
			}
			close(c)
		}(i, c)
	}
	var out []int
	for v := range MergeBufInt(cc...) {
		out = append(out, v)
	}
	sort.Ints(out)
	want := []int{0, 0, 1, 1, 2, 2, 3, 3}
	if !reflect.DeepEqual(out, want) {
		t.Errorf("MergeBufInt() = %v, want %v", out, want)
	}
}
