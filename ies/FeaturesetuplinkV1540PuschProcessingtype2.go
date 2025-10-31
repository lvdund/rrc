package ies

// FeatureSetUplink-v1540-pusch-ProcessingType2 ::= SEQUENCE
type FeaturesetuplinkV1540PuschProcessingtype2 struct {
	Scs15khz *Processingparameters
	Scs30khz *Processingparameters
	Scs60khz *Processingparameters
}
