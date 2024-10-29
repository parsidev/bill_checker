package billing

type BillType uint8

const (
	BillTypeWater BillType = iota + 1
	BillTypeElectronic
	BillTypeGas
	BillTypeLandLine
	BillTypeMobile
	BillTypeMunicipalFees
	_
	BillTypeTax
	BillTypeTrafficFines
)

func (b BillType) Value() string {
	return [...]string{
		"قبض آب",
		"قبض برق",
		"قبض گاز",
		"قبض تلفن ثابت",
		"قبض تلفن همراه",
		"عوارض شهرداری",
		"",
		"سازمان مالیات",
		"جرایم راهنمایی و رانندگی",
	}[b-1]
}
