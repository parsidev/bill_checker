package billing

import "errors"

func Inquiry(req *InquiryRequest) (res *InquiryResponse, err error) {
	if !req.IsValid() {
		return nil, errors.New("اطلاعات قبض وارد شده صحیح نمی‌ باشد")
	}

	return &InquiryResponse{
		BillID:      req.BillID,
		PaymentID:   req.PaymentID,
		Amount:      req.GetAmount(),
		BillType:    req.GetType(),
		BillTypeStr: req.GetType().Value(),
	}, nil
}
