package errorsx

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetMessage(t *testing.T) {
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
			wantErr: errors.New("~~haii~~:: hehe"),
		},
		{
			name: "format with a",
			args: args{
				err:    errors.New("hehe"),
				format: "name=%s age=%d",
				a:      []any{"hidayat", 25},
			},
			wantErr: errors.New("~~name=hidayat age=25~~:: hehe"),
		},
		{
			name: "empty format",
			args: args{
				err:    errors.New("hehe"),
				format: "",
				a:      []any{},
			},
			wantErr: errors.New("~~~~:: hehe"),
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
			err := SetMessage(tt.args.err, tt.args.format, tt.args.a...)
			assert.Equal(t, tt.wantErr.Error(), err.Error())
		})
	}
}

func TestSetMessageE(t *testing.T) {
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
			wantErr: errors.New("~~haha~~:: hehe"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := SetMessageE(tt.args.err1, tt.args.err2)
			if tt.wantErr != nil && err != nil {
				assert.Equal(t, tt.wantErr.Error(), err.Error())
				return
			}
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestGetMessage(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "many match",
			args: args{err: errors.New("grpc.(*ErajolBike).OrderDriver:: usecase.(*ErajolBike).OrderDriver:: ~~invalid request~~:: usecase.(*ErajolBike).validateReqOrderDriver:: ~~request not found~~:: record not found")},
			want: "invalid request",
		},
		{
			name: "",
			args: args{err: errors.New("transporthttp.(*Cart).ATC:: usecase.(*Cart).ATC:: extapi.(*Magento).GetSalesMap:: ~~Account anda belum mempunyai sales brand, silahkan hubungi Sales admin~~:: ~~error get sales map~~:: Cannot find Salesmap data with selected criteria")},
			want: "Account anda belum mempunyai sales brand, silahkan hubungi Sales admin",
		},
		{
			name: "one match",
			args: args{err: errors.New("~~invalid request~~")},
			want: "invalid request",
		},
		{
			name: "no match",
			args: args{err: errors.New("haii there")},
			want: "haii there",
		},
		{
			name: "empty string",
			args: args{err: errors.New("")},
			want: "",
		},
		{
			name: "wraped error",
			args: args{err: errors.New("haha:: hehe")},
			want: "hehe",
		},
		{
			name: "error nil",
			args: args{err: nil},
			want: "",
		},
		{
			name: "",
			args: args{err: errors.New("transporthttp.(*Quote).Checkout:: usecase.(*Quote).Checkout:: extapi.(*B2bcartGRPC).GetCartsToCheckout:: --400-- error get list cart checkout:: rpc error: code = NotFound desc = transportgrpc.(*Cart).GetCartsToCheckout:: usecase.(*Cart).GetCartByCustIDForCheckout:: repo.(*Cart).GetCartListByIDListAndDealerCodePreloadAll:: ~~error get cart~~:: record not found")},
			want: "error get cart",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMessage(tt.args.err); got != tt.want {
				t.Errorf("GetMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
