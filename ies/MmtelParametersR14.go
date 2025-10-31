package ies

// MMTEL-Parameters-r14 ::= SEQUENCE
type MmtelParametersR14 struct {
	DelaybudgetreportingR14    *MmtelParametersR14DelaybudgetreportingR14
	PuschEnhancementsR14       *MmtelParametersR14PuschEnhancementsR14
	RecommendedbitrateR14      *MmtelParametersR14RecommendedbitrateR14
	RecommendedbitratequeryR14 *MmtelParametersR14RecommendedbitratequeryR14
}
