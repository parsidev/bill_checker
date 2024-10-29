package billing

import "testing"

func TestInquiryRequest_GetType(t *testing.T) {
	type fields struct {
		BillID string
		PayID  string
	}
	tests := []struct {
		name   string
		fields fields
		want   BillType
	}{
		{
			name: "Bill Type",
			fields: fields{
				BillID: "5152574430154",
				PayID:  "82832737",
			},
			want: BillTypeMobile,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &InquiryRequest{
				BillID: tt.fields.BillID,
				PayID:  tt.fields.PayID,
			}
			if got := i.GetType(); got != tt.want {
				t.Errorf("GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInquiryRequest_GetAmount(t *testing.T) {
	type fields struct {
		BillID string
		PayID  string
	}
	tests := []struct {
		name   string
		fields fields
		want   uint64
	}{
		{
			name: "Bill Amount",
			fields: fields{
				BillID: "5152574430154",
				PayID:  "82832737",
			},
			want: 828000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &InquiryRequest{
				BillID: tt.fields.BillID,
				PayID:  tt.fields.PayID,
			}
			if got := i.GetAmount(); got != tt.want {
				t.Errorf("GetAmount() = %v, want %v", got, tt.want)
			}
		})
	}
}
