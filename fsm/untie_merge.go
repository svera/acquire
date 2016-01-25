package fsm

type UntieMerge struct {
	ErrorState
}

func (s *UntieMerge) Name() string {
	return UntieMergeStateName
}

func (s *UntieMerge) ToSellTrade() State {
	return &SellTrade{}
}
