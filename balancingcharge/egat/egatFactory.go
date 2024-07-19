package egat

import "hello/balancingcharge"

type EgatFactory struct {
	balancingcharge.BalancingChargeFactory
	Load, Plan, Rate float64
	Type             string
}

func (f *EgatFactory) Consumer() {
	f.BalancingChargeFactory.ChargeType = f.Type
	f.BalancingChargeFactory.ImbalanceState = 1
	f.BalancingChargeFactory.ImbalanceValue = 2

}

func (f *EgatFactory) Generation() {
	f.BalancingChargeFactory.ChargeType = f.Type
	f.BalancingChargeFactory.ImbalanceState = 1
	f.BalancingChargeFactory.ImbalanceValue = 2
}

func (f *EgatFactory) GetBalancingCharge() balancingcharge.Balancing {

	if f.Type == "consumer" {
		f.Consumer()
	}

	if f.Type == "generation" {
		f.Consumer()
	}

	return balancingcharge.Balancing{
		Type:            f.ChargeType,
		ImbalanceEnergy: 1,
		ImbalanceDesc:   f.GetImbalanceDescription(),
	}
}
