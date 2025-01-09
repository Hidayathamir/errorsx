package errorsx

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetCode(t *testing.T) {
	type args struct {
		err  error
		code int
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "error nil",
			args: args{
				err:  nil,
				code: 200,
			},
			wantErr: errors.New("--200--"),
		},
		{
			name: "error not nil",
			args: args{
				err:  errors.New("haha:: hehe"),
				code: 200,
			},
			wantErr: errors.New("--200-- haha:: hehe"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := SetCode(tt.args.err, tt.args.code)
			assert.Equal(t, tt.wantErr.Error(), err.Error())
		})
	}
}

func TestGetCode(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "error nil",
			args: args{
				err: nil,
			},
			want: http.StatusInternalServerError,
		},
		{
			name: "error not contain code",
			args: args{
				err: errors.New("haha:: hehe"),
			},
			want: http.StatusInternalServerError,
		},
		{
			name: "error contain invalid code",
			args: args{
				err: errors.New("--asdf-- haha:: hehe"),
			},
			want: http.StatusInternalServerError,
		},
		{
			name: "error contain code",
			args: args{
				err: errors.New("--404-- haha:: hehe"),
			},
			want: 404,
		},
		{
			name: "error contain multiple code",
			args: args{
				err: errors.New("--404-- haha:: --400-- hehe"),
			},
			want: 404,
		},
		{
			name: "",
			args: args{
				err: errors.New("transporthttp.(*Quote).Checkout:: usecase.(*Quote).Checkout:: extapi.(*B2bcartGRPC).GetCartsToCheckout:: --400-- !!fsefs!! error get list cart checkout:: rpc error: code = NotFound desc = transportgrpc.(*Cart).GetCartsToCheckout:: usecase.(*Cart).GetCartByCustIDForCheckout:: repo.(*Cart).GetCartListByIDListAndDealerCodePreloadAll:: ~~error get cart~~:: record not found"),
			},
			want: 400,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetCode(tt.args.err)
			assert.Equal(t, tt.want, got)
		})
	}
}
