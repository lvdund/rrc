package ies

// SystemInformationBlockType3-cellReselectionServingFreqInfo ::= SEQUENCE
type Systeminformationblocktype3Cellreselectionservingfreqinfo struct {
	SNonintrasearch         *Reselectionthreshold
	Threshservinglow        Reselectionthreshold
	Cellreselectionpriority Cellreselectionpriority
}
