package ies

// CSI-IM-Resource ::= SEQUENCE
// Extensible
type CsiImResource struct {
	CsiImResourceid             CsiImResourceid
	CsiImResourceelementpattern *CsiImResourceCsiImResourceelementpattern
	Freqband                    *CsiFrequencyoccupation
	Periodicityandoffset        *CsiResourceperiodicityandoffset
}
