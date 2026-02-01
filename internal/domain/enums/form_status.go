package enums

type FormStatus int

const (
	FormStatusInactive FormStatus = iota // 0
	FormStatusActive                     // 1
)

func (fs FormStatus) IsValid() bool {
	switch fs {
	case FormStatusInactive, FormStatusActive:
		return true
	default:
		return false
	}
}
