// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SRS-RSRP-MeasResourceSet-r19 (line 15810).

var sRSRSRPMeasResourceSetR19Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "aperiodicTriggeringOffset-r19", Optional: true},
		{Name: "srs-RSRP-MeasResourceSetId-r19", Optional: true},
		{Name: "srs-RSRP-MeasResourceIdList-r19", Optional: true},
	},
}

var sRSRSRPMeasResourceSetR19AperiodicTriggeringOffsetR19Constraints = per.Constrained(1, 31)

var sRSRSRPMeasResourceSetR19SrsRSRPMeasResourceIdListR19Constraints = per.SizeRange(1, common.MaxNrofSRS_RSRP_MeasResourcesPerSet_r19)

type SRS_RSRP_MeasResourceSet_r19 struct {
	AperiodicTriggeringOffset_r19   *int64
	Srs_RSRP_MeasResourceSetId_r19  *SRS_RSRP_MeasResourceSetId_r19
	Srs_RSRP_MeasResourceIdList_r19 []SRS_RSRP_MeasResourceId_r19
}

func (ie *SRS_RSRP_MeasResourceSet_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sRSRSRPMeasResourceSetR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.AperiodicTriggeringOffset_r19 != nil, ie.Srs_RSRP_MeasResourceSetId_r19 != nil, ie.Srs_RSRP_MeasResourceIdList_r19 != nil}); err != nil {
		return err
	}

	// 3. aperiodicTriggeringOffset-r19: integer
	{
		if ie.AperiodicTriggeringOffset_r19 != nil {
			if err := e.EncodeInteger(*ie.AperiodicTriggeringOffset_r19, sRSRSRPMeasResourceSetR19AperiodicTriggeringOffsetR19Constraints); err != nil {
				return err
			}
		}
	}

	// 4. srs-RSRP-MeasResourceSetId-r19: ref
	{
		if ie.Srs_RSRP_MeasResourceSetId_r19 != nil {
			if err := ie.Srs_RSRP_MeasResourceSetId_r19.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. srs-RSRP-MeasResourceIdList-r19: sequence-of
	{
		if ie.Srs_RSRP_MeasResourceIdList_r19 != nil {
			seqOf := e.NewSequenceOfEncoder(sRSRSRPMeasResourceSetR19SrsRSRPMeasResourceIdListR19Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Srs_RSRP_MeasResourceIdList_r19))); err != nil {
				return err
			}
			for i := range ie.Srs_RSRP_MeasResourceIdList_r19 {
				if err := ie.Srs_RSRP_MeasResourceIdList_r19[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *SRS_RSRP_MeasResourceSet_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sRSRSRPMeasResourceSetR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. aperiodicTriggeringOffset-r19: integer
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeInteger(sRSRSRPMeasResourceSetR19AperiodicTriggeringOffsetR19Constraints)
			if err != nil {
				return err
			}
			ie.AperiodicTriggeringOffset_r19 = &v
		}
	}

	// 4. srs-RSRP-MeasResourceSetId-r19: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Srs_RSRP_MeasResourceSetId_r19 = new(SRS_RSRP_MeasResourceSetId_r19)
			if err := ie.Srs_RSRP_MeasResourceSetId_r19.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. srs-RSRP-MeasResourceIdList-r19: sequence-of
	{
		if seq.IsComponentPresent(2) {
			seqOf := d.NewSequenceOfDecoder(sRSRSRPMeasResourceSetR19SrsRSRPMeasResourceIdListR19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Srs_RSRP_MeasResourceIdList_r19 = make([]SRS_RSRP_MeasResourceId_r19, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Srs_RSRP_MeasResourceIdList_r19[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
