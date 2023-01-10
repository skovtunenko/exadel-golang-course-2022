package stack

import (
	is2 "github.com/matryer/is"
	"testing"
)

func TestSliceStack_End2End(t *testing.T) {
	t.Run("popping_from_empty_stack", func(t *testing.T) {
		is := is2.New(t)

		st := NewSliceStack[int]()
		got, gotOk := st.Pop()
		is.True(!gotOk)
		is.Equal(0, got)
	})
	t.Run("popping_from_non-empty_stack", func(t *testing.T) {
		is := is2.New(t)

		const four = 4
		st := NewSliceStack(1, 2, 3)
		st.Push(four)
		gotFour, ok := st.Top()
		is.True(ok)
		is.Equal(four, gotFour)

		gotFour, ok = st.Pop()
		is.True(ok)
		is.Equal(four, gotFour)

		gotThree, ok := st.Pop()
		is.True(ok)
		is.Equal(3, gotThree)

		gotTwo, ok := st.Pop()
		is.True(ok)
		is.Equal(2, gotTwo)
	})
}

func TestNewSliceStack(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
		want *SliceStack[int]
	}{
		{
			name: "no_params",
			args: args{
				data: nil,
			},
			want: &SliceStack[int]{},
		},
		{
			name: "some_params",
			args: args{
				data: []int{1, 2, 3},
			},
			want: &SliceStack[int]{
				data: []int{1, 2, 3},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			is := is2.New(t)
			got := NewSliceStack(tt.args.data...)
			is.Equal(*tt.want, *got)
		})
	}
}

func TestSliceStack_IsEmpty(t *testing.T) {
	type fields struct {
		data []int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "nil stack",
			fields: fields{
				data: nil,
			},
			want: true,
		},
		{
			name: "stack_without_elements",
			fields: fields{
				data: []int{},
			},
			want: true,
		},
		{
			name: "stack_with_elements",
			fields: fields{
				data: []int{1, 2, 3},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			is := is2.New(t)

			ss := NewSliceStack(tt.fields.data...)
			got := ss.IsEmpty()
			is.Equal(tt.want, got)
		})
	}
}

func TestSliceStack_Pop(t *testing.T) {
	type fields struct {
		data []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
		wantOK bool
	}{
		{
			name: "no_initial_values",
			fields: fields{
				data: nil,
			},
			want:   0,
			wantOK: false,
		},
		{
			name: "one_initial_value",
			fields: fields{
				data: []int{1},
			},
			want:   1,
			wantOK: true,
		},
		{
			name: "many_initial_values",
			fields: fields{
				data: []int{1, 2, 3},
			},
			want:   3,
			wantOK: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			is := is2.New(t)

			ss := NewSliceStack(tt.fields.data...)
			got, gotOk := ss.Pop()
			is.Equal(tt.want, got)
			is.Equal(tt.wantOK, gotOk)
		})
	}
}

func TestSliceStack_Push(t *testing.T) {
	type fields struct {
		data []int
	}
	type args struct {
		val int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int
	}{
		{
			name: "pushing_new_value_to_nil_stack",
			fields: fields{
				data: nil,
			},
			args: args{
				val: 777,
			},
			want: []int{777},
		},
		{
			name: "pushing_new_value_to_empty_stack",
			fields: fields{
				data: []int{},
			},
			args: args{
				val: 777,
			},
			want: []int{777},
		},
		{
			name: "pushing_new_value_to_existing_ones",
			fields: fields{
				data: []int{1, 2, 3},
			},
			args: args{
				val: 777,
			},
			want: []int{1, 2, 3, 777},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			is := is2.New(t)
			ss := NewSliceStack(tt.fields.data...)
			ss.Push(tt.args.val)
			is.Equal(tt.want, ss.data)
		})
	}
}

func TestSliceStack_Top(t *testing.T) {
	type fields struct {
		data []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
		wantOK bool
	}{
		{
			name: "nil_data",
			fields: fields{
				data: nil,
			},
			want:   0,
			wantOK: false,
		},
		{
			name: "one_element",
			fields: fields{
				data: []int{1},
			},
			want:   1,
			wantOK: true,
		},
		{
			name:   "couple_elements",
			fields: fields{data: []int{1, 2, 3}},
			want:   3,
			wantOK: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			is := is2.New(t)
			ss := NewSliceStack(tt.fields.data...)
			got, gotOK := ss.Top()
			is.Equal(tt.want, got)
			is.Equal(tt.wantOK, gotOK)
		})
	}
}
