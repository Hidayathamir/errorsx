package errorsx

import (
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrap(t *testing.T) {
	type args struct {
		err    error
		format string
		a      []any
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "only format",
			args: args{
				err:    errors.New("hehe"),
				format: "haii",
				a:      []any{},
			},
			wantErr: errors.New("haii:: hehe"),
		},
		{
			name: "format with a",
			args: args{
				err:    errors.New("hehe"),
				format: "name=%s age=%d",
				a:      []any{"hidayat", 25},
			},
			wantErr: errors.New("name=hidayat age=25:: hehe"),
		},
		{
			name: "empty format",
			args: args{
				err:    errors.New("hehe"),
				format: "",
				a:      []any{},
			},
			wantErr: errors.New(":: hehe"),
		},
		{
			name: "error nil",
			args: args{
				err:    nil,
				format: "",
				a:      []any{},
			},
			wantErr: errors.New(""),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Wrap(tt.args.err, tt.args.format, tt.args.a...)
			assert.Equal(t, tt.wantErr.Error(), err.Error())
		})
	}
}

func TestWrapE(t *testing.T) {
	type args struct {
		err1 error
		err2 error
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "all nil",
			args: args{
				err1: nil,
				err2: nil,
			},
			wantErr: nil,
		},
		{
			name: "err2 nil",
			args: args{
				err1: assert.AnError,
				err2: nil,
			},
			wantErr: assert.AnError,
		},
		{
			name: "err1 nil",
			args: args{
				err1: nil,
				err2: assert.AnError,
			},
			wantErr: assert.AnError,
		},
		{
			name: "all not nil",
			args: args{
				err1: errors.New("hehe"),
				err2: errors.New("haha"),
			},
			wantErr: errors.New("haha:: hehe"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := WrapE(tt.args.err1, tt.args.err2)
			if tt.wantErr != nil && err != nil {
				assert.Equal(t, tt.wantErr.Error(), err.Error())
				return
			}
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestUnwrapToList(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "error nil",
			args: args{
				err: nil,
			},
			want: nil,
		},
		{
			name: "no split",
			args: args{
				err: errors.New("hehe"),
			},
			want: []string{"hehe"},
		},
		{
			name: "split",
			args: args{
				err: errors.New("haha:: hehe"),
			},
			want: []string{"haha", "hehe"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UnwrapToList(tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnwrapToList() = %v, want %v", got, tt.want)
			}
		})
	}
}
