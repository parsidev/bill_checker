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
	billIdValid := checkBillID(splitLastDigit(i.GetBillID()))

	if !billIdValid {
		return errors.New("شناسه قبض وارد شده صحیح نمی باشد")
	}

	payIdValid := checkPayID(splitLastDigit(i.GetPayID()))

	if !payIdValid {
		return errors.New("شناسه پرداخت وارد شده صحیح نمی باشد")
	}

	billValid := checkBill(i.GetBillID(), i.GetPayID())

	if !billValid {
		return errors.New("اطلاعات قبض وارد شده صحیح نمی‌ باشد")
	}

	return nil
}
