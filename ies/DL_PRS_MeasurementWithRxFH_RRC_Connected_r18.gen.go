// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: DL-PRS-MeasurementWithRxFH-RRC-Connected-r18 (line 19395).

var dLPRSMeasurementWithRxFHRRCConnectedR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "maximumPRS-BandwidthAcrossAllHopsFR1-r18", Optional: true},
		{Name: "maximumPRS-BandwidthAcrossAllHopsFR2-r18", Optional: true},
		{Name: "maximumFH-Hops-r18", Optional: true},
		{Name: "processingDuration-r18", Optional: true},
		{Name: "rf-RxRetuneTimeFR1-r18", Optional: true},
		{Name: "rf-RxRetuneTimeFR2-r18", Optional: true},
		{Name: "numOfOverlappingPRB-r18", Optional: true},
	},
}

const (
	DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_MaximumPRS_BandwidthAcrossAllHopsFR1_r18_Mhz40  = 0
	DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_MaximumPRS_BandwidthAcrossAllHopsFR1_r18_Mhz50  = 1
	DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_MaximumPRS_BandwidthAcrossAllHopsFR1_r18_Mhz80  = 2
	DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_MaximumPRS_BandwidthAcrossAllHopsFR1_r18_Mhz100 = 3
)

var dLPRSMeasurementWithRxFHRRCConnectedR18MaximumPRSBandwidthAcrossAllHopsFR1R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_MaximumPRS_BandwidthAcrossAllHopsFR1_r18_Mhz40, DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_MaximumPRS_BandwidthAcrossAllHopsFR1_r18_Mhz50, DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_MaximumPRS_BandwidthAcrossAllHopsFR1_r18_Mhz80, DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_MaximumPRS_BandwidthAcrossAllHopsFR1_r18_Mhz100},
}

const (
	DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_MaximumPRS_BandwidthAcrossAllHopsFR2_r18_Mhz100 = 0
	DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_MaximumPRS_BandwidthAcrossAllHopsFR2_r18_Mhz200 = 1
	DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_MaximumPRS_BandwidthAcrossAllHopsFR2_r18_Mhz400 = 2
)

var dLPRSMeasurementWithRxFHRRCConnectedR18MaximumPRSBandwidthAcrossAllHopsFR2R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_MaximumPRS_BandwidthAcrossAllHopsFR2_r18_Mhz100, DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_MaximumPRS_BandwidthAcrossAllHopsFR2_r18_Mhz200, DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_MaximumPRS_BandwidthAcrossAllHopsFR2_r18_Mhz400},
}

const (
	DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_MaximumFH_Hops_r18_N2 = 0
	DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_MaximumFH_Hops_r18_N3 = 1
	DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_MaximumFH_Hops_r18_N4 = 2
	DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_MaximumFH_Hops_r18_N5 = 3
	DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_MaximumFH_Hops_r18_N6 = 4
)

var dLPRSMeasurementWithRxFHRRCConnectedR18MaximumFHHopsR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_MaximumFH_Hops_r18_N2, DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_MaximumFH_Hops_r18_N3, DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_MaximumFH_Hops_r18_N4, DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_MaximumFH_Hops_r18_N5, DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_MaximumFH_Hops_r18_N6},
}

const (
	DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_Rf_RxRetuneTimeFR1_r18_N70  = 0
	DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_Rf_RxRetuneTimeFR1_r18_N140 = 1
	DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_Rf_RxRetuneTimeFR1_r18_N210 = 2
)

var dLPRSMeasurementWithRxFHRRCConnectedR18RfRxRetuneTimeFR1R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_Rf_RxRetuneTimeFR1_r18_N70, DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_Rf_RxRetuneTimeFR1_r18_N140, DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_Rf_RxRetuneTimeFR1_r18_N210},
}

const (
	DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_Rf_RxRetuneTimeFR2_r18_N35  = 0
	DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_Rf_RxRetuneTimeFR2_r18_N70  = 1
	DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_Rf_RxRetuneTimeFR2_r18_N140 = 2
)

var dLPRSMeasurementWithRxFHRRCConnectedR18RfRxRetuneTimeFR2R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_Rf_RxRetuneTimeFR2_r18_N35, DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_Rf_RxRetuneTimeFR2_r18_N70, DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_Rf_RxRetuneTimeFR2_r18_N140},
}

const (
	DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_NumOfOverlappingPRB_r18_N0 = 0
	DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_NumOfOverlappingPRB_r18_N1 = 1
	DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_NumOfOverlappingPRB_r18_N2 = 2
	DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_NumOfOverlappingPRB_r18_N4 = 3
)

var dLPRSMeasurementWithRxFHRRCConnectedR18NumOfOverlappingPRBR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_NumOfOverlappingPRB_r18_N0, DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_NumOfOverlappingPRB_r18_N1, DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_NumOfOverlappingPRB_r18_N2, DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_NumOfOverlappingPRB_r18_N4},
}

const (
	DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_ProcessingDuration_r18_ProcessingPRS_SymbolsDurationN3_r18_MsDot125 = 0
	DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_ProcessingDuration_r18_ProcessingPRS_SymbolsDurationN3_r18_MsDot25  = 1
	DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_ProcessingDuration_r18_ProcessingPRS_SymbolsDurationN3_r18_MsDot5   = 2
	DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_ProcessingDuration_r18_ProcessingPRS_SymbolsDurationN3_r18_Ms1      = 3
	DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_ProcessingDuration_r18_ProcessingPRS_SymbolsDurationN3_r18_Ms2      = 4
	DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_ProcessingDuration_r18_ProcessingPRS_SymbolsDurationN3_r18_Ms4      = 5
	DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_ProcessingDuration_r18_ProcessingPRS_SymbolsDurationN3_r18_Ms6      = 6
	DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_ProcessingDuration_r18_ProcessingPRS_SymbolsDurationN3_r18_Ms8      = 7
	DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_ProcessingDuration_r18_ProcessingPRS_SymbolsDurationN3_r18_Ms12     = 8
	DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_ProcessingDuration_r18_ProcessingPRS_SymbolsDurationN3_r18_Ms16     = 9
	DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_ProcessingDuration_r18_ProcessingPRS_SymbolsDurationN3_r18_Ms20     = 10
	DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_ProcessingDuration_r18_ProcessingPRS_SymbolsDurationN3_r18_Ms25     = 11
	DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_ProcessingDuration_r18_ProcessingPRS_SymbolsDurationN3_r18_Ms30     = 12
	DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_ProcessingDuration_r18_ProcessingPRS_SymbolsDurationN3_r18_Ms32     = 13
	DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_ProcessingDuration_r18_ProcessingPRS_SymbolsDurationN3_r18_Ms35     = 14
	DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_ProcessingDuration_r18_ProcessingPRS_SymbolsDurationN3_r18_Ms40     = 15
	DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_ProcessingDuration_r18_ProcessingPRS_SymbolsDurationN3_r18_Ms45     = 16
	DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_ProcessingDuration_r18_ProcessingPRS_SymbolsDurationN3_r18_Ms50     = 17
)

var dLPRSMeasurementWithRxFHRRCConnectedR18ProcessingDurationR18ProcessingPRSSymbolsDurationN3R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_ProcessingDuration_r18_ProcessingPRS_SymbolsDurationN3_r18_MsDot125, DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_ProcessingDuration_r18_ProcessingPRS_SymbolsDurationN3_r18_MsDot25, DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_ProcessingDuration_r18_ProcessingPRS_SymbolsDurationN3_r18_MsDot5, DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_ProcessingDuration_r18_ProcessingPRS_SymbolsDurationN3_r18_Ms1, DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_ProcessingDuration_r18_ProcessingPRS_SymbolsDurationN3_r18_Ms2, DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_ProcessingDuration_r18_ProcessingPRS_SymbolsDurationN3_r18_Ms4, DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_ProcessingDuration_r18_ProcessingPRS_SymbolsDurationN3_r18_Ms6, DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_ProcessingDuration_r18_ProcessingPRS_SymbolsDurationN3_r18_Ms8, DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_ProcessingDuration_r18_ProcessingPRS_SymbolsDurationN3_r18_Ms12, DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_ProcessingDuration_r18_ProcessingPRS_SymbolsDurationN3_r18_Ms16, DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_ProcessingDuration_r18_ProcessingPRS_SymbolsDurationN3_r18_Ms20, DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_ProcessingDuration_r18_ProcessingPRS_SymbolsDurationN3_r18_Ms25, DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_ProcessingDuration_r18_ProcessingPRS_SymbolsDurationN3_r18_Ms30, DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_ProcessingDuration_r18_ProcessingPRS_SymbolsDurationN3_r18_Ms32, DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_ProcessingDuration_r18_ProcessingPRS_SymbolsDurationN3_r18_Ms35, DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_ProcessingDuration_r18_ProcessingPRS_SymbolsDurationN3_r18_Ms40, DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_ProcessingDuration_r18_ProcessingPRS_SymbolsDurationN3_r18_Ms45, DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_ProcessingDuration_r18_ProcessingPRS_SymbolsDurationN3_r18_Ms50},
}

const (
	DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_ProcessingDuration_r18_ProcessingDurationT3_r18_Ms8    = 0
	DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_ProcessingDuration_r18_ProcessingDurationT3_r18_Ms16   = 1
	DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_ProcessingDuration_r18_ProcessingDurationT3_r18_Ms20   = 2
	DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_ProcessingDuration_r18_ProcessingDurationT3_r18_Ms30   = 3
	DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_ProcessingDuration_r18_ProcessingDurationT3_r18_Ms40   = 4
	DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_ProcessingDuration_r18_ProcessingDurationT3_r18_Ms80   = 5
	DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_ProcessingDuration_r18_ProcessingDurationT3_r18_Ms160  = 6
	DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_ProcessingDuration_r18_ProcessingDurationT3_r18_Ms320  = 7
	DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_ProcessingDuration_r18_ProcessingDurationT3_r18_Ms640  = 8
	DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_ProcessingDuration_r18_ProcessingDurationT3_r18_Ms1280 = 9
)

var dLPRSMeasurementWithRxFHRRCConnectedR18ProcessingDurationR18ProcessingDurationT3R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_ProcessingDuration_r18_ProcessingDurationT3_r18_Ms8, DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_ProcessingDuration_r18_ProcessingDurationT3_r18_Ms16, DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_ProcessingDuration_r18_ProcessingDurationT3_r18_Ms20, DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_ProcessingDuration_r18_ProcessingDurationT3_r18_Ms30, DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_ProcessingDuration_r18_ProcessingDurationT3_r18_Ms40, DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_ProcessingDuration_r18_ProcessingDurationT3_r18_Ms80, DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_ProcessingDuration_r18_ProcessingDurationT3_r18_Ms160, DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_ProcessingDuration_r18_ProcessingDurationT3_r18_Ms320, DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_ProcessingDuration_r18_ProcessingDurationT3_r18_Ms640, DL_PRS_MeasurementWithRxFH_RRC_Connected_r18_ProcessingDuration_r18_ProcessingDurationT3_r18_Ms1280},
}

type DL_PRS_MeasurementWithRxFH_RRC_Connected_r18 struct {
	MaximumPRS_BandwidthAcrossAllHopsFR1_r18 *int64
	MaximumPRS_BandwidthAcrossAllHopsFR2_r18 *int64
	MaximumFH_Hops_r18                       *int64
	ProcessingDuration_r18                   *struct {
		ProcessingPRS_SymbolsDurationN3_r18 int64
		ProcessingDurationT3_r18            int64
	}
	Rf_RxRetuneTimeFR1_r18  *int64
	Rf_RxRetuneTimeFR2_r18  *int64
	NumOfOverlappingPRB_r18 *int64
}

func (ie *DL_PRS_MeasurementWithRxFH_RRC_Connected_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(dLPRSMeasurementWithRxFHRRCConnectedR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.MaximumPRS_BandwidthAcrossAllHopsFR1_r18 != nil, ie.MaximumPRS_BandwidthAcrossAllHopsFR2_r18 != nil, ie.MaximumFH_Hops_r18 != nil, ie.ProcessingDuration_r18 != nil, ie.Rf_RxRetuneTimeFR1_r18 != nil, ie.Rf_RxRetuneTimeFR2_r18 != nil, ie.NumOfOverlappingPRB_r18 != nil}); err != nil {
		return err
	}

	// 3. maximumPRS-BandwidthAcrossAllHopsFR1-r18: enumerated
	{
		if ie.MaximumPRS_BandwidthAcrossAllHopsFR1_r18 != nil {
			if err := e.EncodeEnumerated(*ie.MaximumPRS_BandwidthAcrossAllHopsFR1_r18, dLPRSMeasurementWithRxFHRRCConnectedR18MaximumPRSBandwidthAcrossAllHopsFR1R18Constraints); err != nil {
				return err
			}
		}
	}

	// 4. maximumPRS-BandwidthAcrossAllHopsFR2-r18: enumerated
	{
		if ie.MaximumPRS_BandwidthAcrossAllHopsFR2_r18 != nil {
			if err := e.EncodeEnumerated(*ie.MaximumPRS_BandwidthAcrossAllHopsFR2_r18, dLPRSMeasurementWithRxFHRRCConnectedR18MaximumPRSBandwidthAcrossAllHopsFR2R18Constraints); err != nil {
				return err
			}
		}
	}

	// 5. maximumFH-Hops-r18: enumerated
	{
		if ie.MaximumFH_Hops_r18 != nil {
			if err := e.EncodeEnumerated(*ie.MaximumFH_Hops_r18, dLPRSMeasurementWithRxFHRRCConnectedR18MaximumFHHopsR18Constraints); err != nil {
				return err
			}
		}
	}

	// 6. processingDuration-r18: sequence
	{
		if ie.ProcessingDuration_r18 != nil {
			c := ie.ProcessingDuration_r18
			if err := e.EncodeEnumerated(c.ProcessingPRS_SymbolsDurationN3_r18, dLPRSMeasurementWithRxFHRRCConnectedR18ProcessingDurationR18ProcessingPRSSymbolsDurationN3R18Constraints); err != nil {
				return err
			}
			if err := e.EncodeEnumerated(c.ProcessingDurationT3_r18, dLPRSMeasurementWithRxFHRRCConnectedR18ProcessingDurationR18ProcessingDurationT3R18Constraints); err != nil {
				return err
			}
		}
	}

	// 7. rf-RxRetuneTimeFR1-r18: enumerated
	{
		if ie.Rf_RxRetuneTimeFR1_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Rf_RxRetuneTimeFR1_r18, dLPRSMeasurementWithRxFHRRCConnectedR18RfRxRetuneTimeFR1R18Constraints); err != nil {
				return err
			}
		}
	}

	// 8. rf-RxRetuneTimeFR2-r18: enumerated
	{
		if ie.Rf_RxRetuneTimeFR2_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Rf_RxRetuneTimeFR2_r18, dLPRSMeasurementWithRxFHRRCConnectedR18RfRxRetuneTimeFR2R18Constraints); err != nil {
				return err
			}
		}
	}

	// 9. numOfOverlappingPRB-r18: enumerated
	{
		if ie.NumOfOverlappingPRB_r18 != nil {
			if err := e.EncodeEnumerated(*ie.NumOfOverlappingPRB_r18, dLPRSMeasurementWithRxFHRRCConnectedR18NumOfOverlappingPRBR18Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *DL_PRS_MeasurementWithRxFH_RRC_Connected_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(dLPRSMeasurementWithRxFHRRCConnectedR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. maximumPRS-BandwidthAcrossAllHopsFR1-r18: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(dLPRSMeasurementWithRxFHRRCConnectedR18MaximumPRSBandwidthAcrossAllHopsFR1R18Constraints)
			if err != nil {
				return err
			}
			ie.MaximumPRS_BandwidthAcrossAllHopsFR1_r18 = &idx
		}
	}

	// 4. maximumPRS-BandwidthAcrossAllHopsFR2-r18: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(dLPRSMeasurementWithRxFHRRCConnectedR18MaximumPRSBandwidthAcrossAllHopsFR2R18Constraints)
			if err != nil {
				return err
			}
			ie.MaximumPRS_BandwidthAcrossAllHopsFR2_r18 = &idx
		}
	}

	// 5. maximumFH-Hops-r18: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(dLPRSMeasurementWithRxFHRRCConnectedR18MaximumFHHopsR18Constraints)
			if err != nil {
				return err
			}
			ie.MaximumFH_Hops_r18 = &idx
		}
	}

	// 6. processingDuration-r18: sequence
	{
		if seq.IsComponentPresent(3) {
			ie.ProcessingDuration_r18 = &struct {
				ProcessingPRS_SymbolsDurationN3_r18 int64
				ProcessingDurationT3_r18            int64
			}{}
			c := ie.ProcessingDuration_r18
			{
				v, err := d.DecodeEnumerated(dLPRSMeasurementWithRxFHRRCConnectedR18ProcessingDurationR18ProcessingPRSSymbolsDurationN3R18Constraints)
				if err != nil {
					return err
				}
				c.ProcessingPRS_SymbolsDurationN3_r18 = v
			}
			{
				v, err := d.DecodeEnumerated(dLPRSMeasurementWithRxFHRRCConnectedR18ProcessingDurationR18ProcessingDurationT3R18Constraints)
				if err != nil {
					return err
				}
				c.ProcessingDurationT3_r18 = v
			}
		}
	}

	// 7. rf-RxRetuneTimeFR1-r18: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(dLPRSMeasurementWithRxFHRRCConnectedR18RfRxRetuneTimeFR1R18Constraints)
			if err != nil {
				return err
			}
			ie.Rf_RxRetuneTimeFR1_r18 = &idx
		}
	}

	// 8. rf-RxRetuneTimeFR2-r18: enumerated
	{
		if seq.IsComponentPresent(5) {
			idx, err := d.DecodeEnumerated(dLPRSMeasurementWithRxFHRRCConnectedR18RfRxRetuneTimeFR2R18Constraints)
			if err != nil {
				return err
			}
			ie.Rf_RxRetuneTimeFR2_r18 = &idx
		}
	}

	// 9. numOfOverlappingPRB-r18: enumerated
	{
		if seq.IsComponentPresent(6) {
			idx, err := d.DecodeEnumerated(dLPRSMeasurementWithRxFHRRCConnectedR18NumOfOverlappingPRBR18Constraints)
			if err != nil {
				return err
			}
			ie.NumOfOverlappingPRB_r18 = &idx
		}
	}

	return nil
}
