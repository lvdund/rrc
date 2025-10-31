package ies

// FeatureSetDownlink-v1540-pdsch-ProcessingType2 ::= SEQUENCE
type FeaturesetdownlinkV1540PdschProcessingtype2 struct {
	Scs15khz *Processingparameters
	Scs30khz *Processingparameters
	Scs60khz *Processingparameters
}
