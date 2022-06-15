package maps

import (
	"reflect"
	"sync"
	"testing"
)

func newHashMapFromMap[K comparable, V any](m map[K]V) *hashMap[K, V] {
	return &hashMap[K, V]{
		m:  m,
		mu: &sync.RWMutex{},
	}
}

func TestHashMap_Clear(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := newHashMap[string, int64]()
			h.Clear()
		})
	}
}

func TestHashMap_ContainsKey(t *testing.T) {
	tests := []struct {
		name string
		hm   *hashMap[string, int64]
		args string
		want bool
	}{
		{
			name: "test",
			hm:   newHashMap[string, int64](),
			args: "test",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.hm.ContainsKey(tt.args); got != tt.want {
				t.Errorf("ContainsKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashMap_ContainsValue(t *testing.T) {
	tests := []struct {
		name string
		hm   *hashMap[string, string]
		args string
		want bool
	}{
		{
			name: "test",
			hm:   newHashMap[string, string](),
			args: "test",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.hm.ContainsValue(tt.args); got != tt.want {
				t.Errorf("ContainsValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashMap_Get(t *testing.T) {
	type args[K comparable, V any] struct {
		key  K
		val  V
		want V
	}
	tests := []struct {
		name string
		hm   *hashMap[string, int64]
		args args[string, int64]
	}{
		{
			name: "test",
			hm:   newHashMap[string, int64](),
			args: args[string, int64]{
				key:  "test",
				val:  1,
				want: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.hm.Put(tt.args.key, tt.args.val)
			got := tt.hm.Get(tt.args.key)
			if got == nil {
				t.Errorf("Get() = %v, want %v", got, tt.args.want)
				return
			}
			if !reflect.DeepEqual(*got, tt.args.want) {
				t.Errorf("Get() = %v, want %v", *got, tt.args.want)
			}
		})
	}
}

func TestHashMap_IsEmpty(t *testing.T) {
	tests := []struct {
		name string
		hm   *hashMap[string, int64]
		want bool
	}{
		{
			name: "test",
			hm:   newHashMap[string, int64](),
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.hm.IsEmpty(); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashMap_Keys(t *testing.T) {
	type args[K comparable, V any] struct {
		m    map[K]V
		want []K
	}
	tests := []struct {
		name string
		args args[string, int64]
	}{
		{
			name: "test",
			args: args[string, int64]{
				m: map[string]int64{
					"test": 1,
				},
				want: []string{"test"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := newHashMap[string, int64]()
			for k, v := range tt.args.m {
				h.Put(k, v)
			}

			got := h.Keys()
			if !reflect.DeepEqual(got.Values(), tt.args.want) {
				t.Errorf("Keys() = %v, want %v", got.Values(), tt.args.want)
			}
		})
	}
}

func TestHashMap_Merge(t *testing.T) {
	type fields[K comparable, V any] struct {
		ms []map[K]V
	}

	type args[M comparable, N any] struct {
		m map[M]N
	}

	tests := []struct {
		name   string
		fields fields[string, int64]
		args   args[string, int64]
		want   bool
	}{
		{
			name: "test",
			fields: fields[string, int64]{
				ms: []map[string]int64{
					{
						"test": 1,
					},
					{
						"test2": 2,
					},
				},
			},
			args: args[string, int64]{
				m: map[string]int64{
					"test":  1,
					"test2": 2,
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := newHashMap[string, int64]()
			ms := make([]Map[string, int64], 0)
			for _, m := range tt.fields.ms {
				ms = append(ms, newHashMapFromMap[string, int64](m))
			}
			if got := h.Merge(ms...); got != tt.want {
				t.Errorf("Merge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashMap_Put(t *testing.T) {
	type args[K comparable, V any] struct {
		key K
		val V
	}
	tests := []struct {
		name string
		args args[string, int64]
		want bool
	}{
		{
			name: "test",
			args: args[string, int64]{
				key: "test",
				val: 1,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := newHashMap[string, int64]()
			if got := h.Put(tt.args.key, tt.args.val); got != tt.want {
				t.Errorf("Put() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashMap_Range(t *testing.T) {
	type args[K comparable, V any] struct {
		k K
		v V
		f func(K, V) bool
	}
	tests := []struct {
		name string
		args args[string, int64]
	}{
		{
			name: "test",
			args: args[string, int64]{
				k: "test",
				v: 1,
				f: func(string, int64) bool {
					return true
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := newHashMap[string, int64]()
			h.Put(tt.args.k, tt.args.v)
			h.Range(tt.args.f)
		})
	}
}

func TestHashMap_Remove(t *testing.T) {
	type args[K comparable, V any] struct {
		key K
		val V
	}
	tests := []struct {
		name string
		args args[string, int64]
		want bool
	}{
		{
			name: "test",
			args: args[string, int64]{
				key: "test",
				val: 1,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := newHashMap[string, int64]()
			h.Put(tt.args.key, tt.args.val)
			if got := h.Remove(tt.args.key); got != tt.want {
				t.Errorf("Remove() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashMap_Size(t *testing.T) {
	type args[K comparable, V any] map[K]V
	tests := []struct {
		name string
		args args[string, int64]
		want int
	}{
		{
			name: "test",
			args: args[string, int64]{
				"test":  1,
				"test2": 2,
				"test3": 3,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := newHashMap[string, int64]()
			for k, v := range tt.args {
				h.Put(k, v)
			}
			if got := h.Size(); got != tt.want {
				t.Errorf("Size() = %v, want %v", got, tt.want)
			}
		})
	}
}
