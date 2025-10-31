package ies

// FeatureSetDownlink-timeDurationForQCL ::= SEQUENCE
type FeaturesetdownlinkTimedurationforqcl struct {
	Scs60khz  *FeaturesetdownlinkTimedurationforqclScs60khz
	Scs120khz *FeaturesetdownlinkTimedurationforqclScs120khz
}
