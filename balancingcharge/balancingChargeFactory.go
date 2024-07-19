package balancingcharge

type Balancing struct {
	Type            string  `json:"type" example:"consumer"`
	ImbalanceEnergy float64 `json:"imbalanceEnergy" example:"0.09999999999999964"`
	ImbalanceDesc   string  `json:"imbalanceDesc" example:"Imbalance energy at generation supply lower than Energy Plan"`
}

// Polymorphism
type Factory interface {
	Consumer()
	Generation()
	GetBalancingCharge() Balancing
}

type BalancingChargeFactory struct {
	ImbalanceState float64
	ImbalanceValue float64
	ChargeType     string
}

func (f *BalancingChargeFactory) GetImbalanceDescription() string {
	if f.ChargeType == "generation" {
		if f.ImbalanceState >= 0 {
			return "Imbalance energy at generation supply higher than Energy Plan"
		}

		return "Imbalance energy at generation supply lower than Energy Plan"
	}

	if f.ChargeType == "consumer" {
		if f.ImbalanceState >= 0 {
			return "Imbalance energy at consumer / prosumer higher than Energy Plan"
		}

		return "Imbalance energy at consumer / prosumer lower than Energy Plan"
	}

	return ""
}
