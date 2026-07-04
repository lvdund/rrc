// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RateMatchPatternLTE-CRS (line 13298).

var rateMatchPatternLTECRSConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "carrierFreqDL"},
		{Name: "carrierBandwidthDL"},
		{Name: "mbsfn-SubframeConfigList", Optional: true},
		{Name: "nrofCRS-Ports"},
		{Name: "v-Shift"},
	},
}

var rateMatchPatternLTECRSCarrierFreqDLConstraints = per.Constrained(0, 16383)

const (
	RateMatchPatternLTE_CRS_CarrierBandwidthDL_N6     = 0
	RateMatchPatternLTE_CRS_CarrierBandwidthDL_N15    = 1
	RateMatchPatternLTE_CRS_CarrierBandwidthDL_N25    = 2
	RateMatchPatternLTE_CRS_CarrierBandwidthDL_N50    = 3
	RateMatchPatternLTE_CRS_CarrierBandwidthDL_N75    = 4
	RateMatchPatternLTE_CRS_CarrierBandwidthDL_N100   = 5
	RateMatchPatternLTE_CRS_CarrierBandwidthDL_Spare2 = 6
	RateMatchPatternLTE_CRS_CarrierBandwidthDL_Spare1 = 7
)

var rateMatchPatternLTECRSCarrierBandwidthDLConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RateMatchPatternLTE_CRS_CarrierBandwidthDL_N6, RateMatchPatternLTE_CRS_CarrierBandwidthDL_N15, RateMatchPatternLTE_CRS_CarrierBandwidthDL_N25, RateMatchPatternLTE_CRS_CarrierBandwidthDL_N50, RateMatchPatternLTE_CRS_CarrierBandwidthDL_N75, RateMatchPatternLTE_CRS_CarrierBandwidthDL_N100, RateMatchPatternLTE_CRS_CarrierBandwidthDL_Spare2, RateMatchPatternLTE_CRS_CarrierBandwidthDL_Spare1},
}

const (
	RateMatchPatternLTE_CRS_NrofCRS_Ports_N1 = 0
	RateMatchPatternLTE_CRS_NrofCRS_Ports_N2 = 1
	RateMatchPatternLTE_CRS_NrofCRS_Ports_N4 = 2
)

var rateMatchPatternLTECRSNrofCRSPortsConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RateMatchPatternLTE_CRS_NrofCRS_Ports_N1, RateMatchPatternLTE_CRS_NrofCRS_Ports_N2, RateMatchPatternLTE_CRS_NrofCRS_Ports_N4},
}

const (
	RateMatchPatternLTE_CRS_V_Shift_N0 = 0
	RateMatchPatternLTE_CRS_V_Shift_N1 = 1
	RateMatchPatternLTE_CRS_V_Shift_N2 = 2
	RateMatchPatternLTE_CRS_V_Shift_N3 = 3
	RateMatchPatternLTE_CRS_V_Shift_N4 = 4
	RateMatchPatternLTE_CRS_V_Shift_N5 = 5
)

var rateMatchPatternLTECRSVShiftConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RateMatchPatternLTE_CRS_V_Shift_N0, RateMatchPatternLTE_CRS_V_Shift_N1, RateMatchPatternLTE_CRS_V_Shift_N2, RateMatchPatternLTE_CRS_V_Shift_N3, RateMatchPatternLTE_CRS_V_Shift_N4, RateMatchPatternLTE_CRS_V_Shift_N5},
}

type RateMatchPatternLTE_CRS struct {
	CarrierFreqDL            int64
	CarrierBandwidthDL       int64
	Mbsfn_SubframeConfigList *EUTRA_MBSFN_SubframeConfigList
	NrofCRS_Ports            int64
	V_Shift                  int64
}

func (ie *RateMatchPatternLTE_CRS) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rateMatchPatternLTECRSConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Mbsfn_SubframeConfigList != nil}); err != nil {
		return err
	}

	// 2. carrierFreqDL: integer
	{
		if err := e.EncodeInteger(ie.CarrierFreqDL, rateMatchPatternLTECRSCarrierFreqDLConstraints); err != nil {
			return err
		}
	}

	// 3. carrierBandwidthDL: enumerated
	{
		if err := e.EncodeEnumerated(ie.CarrierBandwidthDL, rateMatchPatternLTECRSCarrierBandwidthDLConstraints); err != nil {
			return err
		}
	}

	// 4. mbsfn-SubframeConfigList: ref
	{
		if ie.Mbsfn_SubframeConfigList != nil {
			if err := ie.Mbsfn_SubframeConfigList.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. nrofCRS-Ports: enumerated
	{
		if err := e.EncodeEnumerated(ie.NrofCRS_Ports, rateMatchPatternLTECRSNrofCRSPortsConstraints); err != nil {
			return err
		}
	}

	// 6. v-Shift: enumerated
	{
		if err := e.EncodeEnumerated(ie.V_Shift, rateMatchPatternLTECRSVShiftConstraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *RateMatchPatternLTE_CRS) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rateMatchPatternLTECRSConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. carrierFreqDL: integer
	{
		v0, err := d.DecodeInteger(rateMatchPatternLTECRSCarrierFreqDLConstraints)
		if err != nil {
			return err
		}
		ie.CarrierFreqDL = v0
	}

	// 3. carrierBandwidthDL: enumerated
	{
		v1, err := d.DecodeEnumerated(rateMatchPatternLTECRSCarrierBandwidthDLConstraints)
		if err != nil {
			return err
		}
		ie.CarrierBandwidthDL = v1
	}

	// 4. mbsfn-SubframeConfigList: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Mbsfn_SubframeConfigList = new(EUTRA_MBSFN_SubframeConfigList)
			if err := ie.Mbsfn_SubframeConfigList.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. nrofCRS-Ports: enumerated
	{
		v3, err := d.DecodeEnumerated(rateMatchPatternLTECRSNrofCRSPortsConstraints)
		if err != nil {
			return err
		}
		ie.NrofCRS_Ports = v3
	}

	// 6. v-Shift: enumerated
	{
		v4, err := d.DecodeEnumerated(rateMatchPatternLTECRSVShiftConstraints)
		if err != nil {
			return err
		}
		ie.V_Shift = v4
	}

	return nil
}
