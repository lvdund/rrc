package ies

// AppLayerMeasParameters-r17 ::= SEQUENCE
// Extensible
type ApplayermeasparametersR17 struct {
	QoeStreamingMeasreportR17           *ApplayermeasparametersR17QoeStreamingMeasreportR17
	QoeMtsiMeasreportR17                *ApplayermeasparametersR17QoeMtsiMeasreportR17
	QoeVrMeasreportR17                  *ApplayermeasparametersR17QoeVrMeasreportR17
	RanVisibleqoeStreamingMeasreportR17 *ApplayermeasparametersR17RanVisibleqoeStreamingMeasreportR17
	RanVisibleqoeVrMeasreportR17        *ApplayermeasparametersR17RanVisibleqoeVrMeasreportR17
	UlMeasurementreportapplayerSegR17   *ApplayermeasparametersR17UlMeasurementreportapplayerSegR17
}
