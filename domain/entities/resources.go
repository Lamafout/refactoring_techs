package entities

type Resources struct {
	ID                        int     `json:"id" db:"id"`
	Energy                    float64 `json:"energy" db:"energy"`
	Water                     float64 `json:"water" db:"water"`
	NeutralizationAndDisposal float64 `json:"neutralization_and_disposal" db:"neutralization_and_disposal"`
}
