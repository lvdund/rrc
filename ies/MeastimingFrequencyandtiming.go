package ies

// MeasTiming-frequencyAndTiming ::= SEQUENCE
type MeastimingFrequencyandtiming struct {
	Carrierfreq                       ArfcnValuenr
	Ssbsubcarrierspacing              Subcarrierspacing
	SsbMeasurementtimingconfiguration SsbMtc
	SsRssiMeasurement                 *SsRssiMeasurement
}
