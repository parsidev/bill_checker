package billing

import (
	"math/big"
	"strings"
)

type InquiryRequest struct {
	BillID string `json:"bill_id"`
	PayID  string `json:"pay_id"`
}

type InquiryResponse struct {
	BillID      string   `json:"bill_id"`
	PayID       string   `json:"pay_id"`
	Amount      uint64   `json:"amount"`
	BillType    BillType `json:"bill_type"`
	BillTypeStr string   `json:"type"`
}

func (i *InquiryRequest) GetBillID() *big.Int {
	r := new(big.Int)
	r.SetString(i.BillID, 10)

	return r
}

func (i *InquiryRequest) GetPayID() *big.Int {
	r := new(big.Int)
	r.SetString(i.PayID, 10)

	return r
}

func (i *InquiryRequest) GetType() BillType {
	billID, _ := splitLastDigit(i.GetBillID())
	_, billType := splitLastDigit(billID)

	return BillType(billType)
}

func (i *InquiryRequest) GetAmount() uint64 {
	payID := strings.Join([]string{strings.Repeat("0", 13-len(i.PayID)), i.PayID}, "")

	priceStr := payID[:8]

	price := new(big.Int)
	price.SetString(priceStr, 10)

	return price.Uint64() * 1000
}

func (i *InquiryRequest) IsValid() bool {
	return checkBill(i.GetBillID(), i.GetPayID())
}
