package providers

type WeightLevel int

const (
	ErrorLevel WeightLevel = iota - 1
	LowLevel
	CriticalLevel
)

func (w WeightLevel) ToString() string {
	switch w {
	case ErrorLevel:
		return "ErrorWeight"
	case LowLevel:
		return "Under"
	case CriticalLevel:
		return "Normal"
	default:
		return "UnknownWeight"
	}
}

type WeightEngine interface {
	Name() string                                    // name of engine
	SearchWeight(domain string) (WeightLevel, error) // main function to get weight
}
