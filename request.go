package billing

import (
	"errors"
	"math/big"
	"strings"
)

type InquiryRequest struct {
	BillID    string `json:"bill_id"`
	PaymentID string `json:"payment_id"`
}

type InquiryResponse struct {
	BillID      string   `json:"bill_id"`
	PaymentID   string   `json:"payment_id"`
	Amount      uint64   `json:"amount"`
	BillType    BillType `json:"type"`
	BillTypeStr string   `json:"bill_type"`
}

func (i *InquiryRequest) GetBillID() *big.Int {
	r := new(big.Int)
	r.SetString(i.BillID, 10)

	return r
}

func (i *InquiryRequest) GetPayID() *big.Int {
	r := new(big.Int)
	r.SetString(i.PaymentID, 10)

	return r
}

func (i *InquiryRequest) GetType() BillType {
	billID, _ := splitLastDigit(i.GetBillID())
	_, billType := splitLastDigit(billID)

	return BillType(billType)
}

func (i *InquiryRequest) GetAmount() uint64 {
	payID := strings.Join([]string{strings.Repeat("0", 13-len(i.PaymentID)), i.PaymentID}, "")

	priceStr := payID[:8]

	price := new(big.Int)
	price.SetString(priceStr, 10)

	return price.Uint64() * 1000
}

func (i *InquiryRequest) Validate() error {
	billID := i.GetBillID()
	payID := i.GetPayID()

	if !checkBillID(splitLastDigit(billID)) {
		return errors.New("شناسه قبض وارد شده صحیح نمی باشد")
	}

	splitPayID, _ := splitLastDigit(payID)
	if !checkPayID(splitLastDigit(splitPayID)) {
		return errors.New("شناسه پرداخت وارد شده صحیح نمی باشد")
	}

	if !checkBill(billID, payID) {
		return errors.New("اطلاعات قبض وارد شده صحیح نمی‌ باشد")
	}

	return nil
}
