package queue

import (
	"reflect"
	"testing"
)

func TestArray_Put(t *testing.T) {
	type fields struct {
		front int
		rear  int
		size  int
		queue []interface{}
	}
	type args struct {
		v interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "empty",
			fields: fields{
				size:  10,
				queue: make([]interface{}, 10),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Array{
				front: tt.fields.front,
				rear:  tt.fields.rear,
				size:  tt.fields.size,
				queue: tt.fields.queue,
			}
			if got := a.Put(tt.args.v); got != tt.want {
				t.Errorf("Array.Put() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArray_Poll(t *testing.T) {
	type fields struct {
		front int
		rear  int
		size  int
		queue []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
		want1  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Array{
				front: tt.fields.front,
				rear:  tt.fields.rear,
				size:  tt.fields.size,
				queue: tt.fields.queue,
			}
			got, got1 := a.Poll()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Array.Poll() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Array.Poll() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
