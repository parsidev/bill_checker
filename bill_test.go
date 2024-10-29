package billing

import (
	"reflect"
	"testing"
)

func TestInquiry(t *testing.T) {
	type args struct {
		req *InquiryRequest
	}
	tests := []struct {
		name    string
		args    args
		wantRes *InquiryResponse
		wantErr bool
	}{
		{
			name: "Bill Inquiry Test",
			args: args{
				req: &InquiryRequest{
					BillID:    "5152574430154",
					PaymentID: "82832737",
				},
			},
			wantRes: &InquiryResponse{
				BillID:      "5152574430154",
				PaymentID:   "82832737",
				Amount:      828000,
				BillType:    BillTypeMobile,
				BillTypeStr: BillTypeMobile.Value(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := Inquiry(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Inquiry() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Inquiry() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
