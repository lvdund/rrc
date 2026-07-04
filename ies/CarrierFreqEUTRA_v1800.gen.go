// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CarrierFreqEUTRA-v1800 (line 4161).

var carrierFreqEUTRAV1800Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "eutra-MultiBandInfoListAerial-r18", Optional: true},
		{Name: "tn-AreaIdList-r18", Optional: true},
	},
}

var carrierFreqEUTRAV1800TnAreaIdListR18Constraints = per.SizeRange(1, common.MaxTN_AreaInfo_r18)

type CarrierFreqEUTRA_v1800 struct {
	Eutra_MultiBandInfoListAerial_r18 *EUTRA_MultiBandInfoListAerial_r18
	Tn_AreaIdList_r18                 []TN_AreaId_r18
}

func (ie *CarrierFreqEUTRA_v1800) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(carrierFreqEUTRAV1800Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Eutra_MultiBandInfoListAerial_r18 != nil, ie.Tn_AreaIdList_r18 != nil}); err != nil {
		return err
	}

	// 2. eutra-MultiBandInfoListAerial-r18: ref
	{
		if ie.Eutra_MultiBandInfoListAerial_r18 != nil {
			if err := ie.Eutra_MultiBandInfoListAerial_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. tn-AreaIdList-r18: sequence-of
	{
		if ie.Tn_AreaIdList_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(carrierFreqEUTRAV1800TnAreaIdListR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Tn_AreaIdList_r18))); err != nil {
				return err
			}
			for i := range ie.Tn_AreaIdList_r18 {
				if err := ie.Tn_AreaIdList_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *CarrierFreqEUTRA_v1800) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(carrierFreqEUTRAV1800Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. eutra-MultiBandInfoListAerial-r18: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Eutra_MultiBandInfoListAerial_r18 = new(EUTRA_MultiBandInfoListAerial_r18)
			if err := ie.Eutra_MultiBandInfoListAerial_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. tn-AreaIdList-r18: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(carrierFreqEUTRAV1800TnAreaIdListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Tn_AreaIdList_r18 = make([]TN_AreaId_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Tn_AreaIdList_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
