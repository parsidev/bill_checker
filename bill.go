package billing

func Inquiry(req *InquiryRequest) (res *InquiryResponse, err error) {
	if err = req.Validate(); err != nil {
		return nil, err
	}

	res = &InquiryResponse{
		BillID:      req.BillID,
		PaymentID:   req.PaymentID,
		Amount:      req.GetAmount(),
		BillType:    req.GetType(),
		BillTypeStr: req.GetType().Value(),
	}

	return res, nil
}
