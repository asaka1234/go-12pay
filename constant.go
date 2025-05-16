package go_12pay

// 返回值的code枚举

type One2payDepositStatus int

const (
	Success One2payDepositStatus = iota + 200
	Created
)

func (s One2payDepositStatus) Code() int {
	return int(s)
}

func (s One2payDepositStatus) Name() string {
	switch s {
	case Success:
		return "Success"
	case Created:
		return "Created"
	default:
		return "Unknown"
	}
}

func (s One2payDepositStatus) Desc() string {
	switch s {
	case Success:
		return "Success"
	case Created:
		return "Created"
	default:
		return "Unknown"
	}
}

func (s One2payDepositStatus) Eq(value interface{}) bool {
	switch v := value.(type) {
	case int:
		return s.Code() == v
	case One2payDepositStatus:
		return s == v
	default:
		return false
	}
}
