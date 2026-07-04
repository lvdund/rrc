// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SRS-PosRRC-InactiveAggBW-AdditionalCarriers-r18 (line 1465).

var sRSPosRRCInactiveAggBWAdditionalCarriersR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "aggregatedPosSRS-CarrierList-r18", Optional: true},
	},
}

var sRSPosRRCInactiveAggBWAdditionalCarriersR18AggregatedPosSRSCarrierListR18Constraints = per.SizeRange(1, common.MaxNrOfLinkedSRS_CarriersInactive_1_r18)

type SRS_PosRRC_InactiveAggBW_AdditionalCarriers_r18 struct {
	AggregatedPosSRS_CarrierList_r18 []SRS_PosConfigPerULCarrier_r18
}

func (ie *SRS_PosRRC_InactiveAggBW_AdditionalCarriers_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sRSPosRRCInactiveAggBWAdditionalCarriersR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.AggregatedPosSRS_CarrierList_r18 != nil}); err != nil {
		return err
	}

	// 3. aggregatedPosSRS-CarrierList-r18: sequence-of
	{
		if ie.AggregatedPosSRS_CarrierList_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(sRSPosRRCInactiveAggBWAdditionalCarriersR18AggregatedPosSRSCarrierListR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.AggregatedPosSRS_CarrierList_r18))); err != nil {
				return err
			}
			for i := range ie.AggregatedPosSRS_CarrierList_r18 {
				if err := ie.AggregatedPosSRS_CarrierList_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *SRS_PosRRC_InactiveAggBW_AdditionalCarriers_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sRSPosRRCInactiveAggBWAdditionalCarriersR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. aggregatedPosSRS-CarrierList-r18: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(sRSPosRRCInactiveAggBWAdditionalCarriersR18AggregatedPosSRSCarrierListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.AggregatedPosSRS_CarrierList_r18 = make([]SRS_PosConfigPerULCarrier_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.AggregatedPosSRS_CarrierList_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
