// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CSI-LogMeasInfo-r19 (line 3508).

var cSILogMeasInfoR19Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "csi-RS-MeasResultList-r19", Optional: true},
		{Name: "ssb-MeasResultList-r19", Optional: true},
		{Name: "timeGap-r19", Optional: true},
	},
}

var cSILogMeasInfoR19CsiRSMeasResultListR19Constraints = per.SizeRange(1, common.MaxNrofNZP_CSI_RS_Resources)

var cSILogMeasInfoR19SsbMeasResultListR19Constraints = per.SizeRange(1, common.MaxNrofSSBs_r16)

const (
	CSI_LogMeasInfo_r19_TimeGap_r19_True = 0
)

var cSILogMeasInfoR19TimeGapR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CSI_LogMeasInfo_r19_TimeGap_r19_True},
}

type CSI_LogMeasInfo_r19 struct {
	Csi_RS_MeasResultList_r19 []CSI_RS_MeasResult_r19
	Ssb_MeasResultList_r19    []SSB_MeasResult_r19
	TimeGap_r19               *int64
}

func (ie *CSI_LogMeasInfo_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cSILogMeasInfoR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Csi_RS_MeasResultList_r19 != nil, ie.Ssb_MeasResultList_r19 != nil, ie.TimeGap_r19 != nil}); err != nil {
		return err
	}

	// 3. csi-RS-MeasResultList-r19: sequence-of
	{
		if ie.Csi_RS_MeasResultList_r19 != nil {
			seqOf := e.NewSequenceOfEncoder(cSILogMeasInfoR19CsiRSMeasResultListR19Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Csi_RS_MeasResultList_r19))); err != nil {
				return err
			}
			for i := range ie.Csi_RS_MeasResultList_r19 {
				if err := ie.Csi_RS_MeasResultList_r19[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 4. ssb-MeasResultList-r19: sequence-of
	{
		if ie.Ssb_MeasResultList_r19 != nil {
			seqOf := e.NewSequenceOfEncoder(cSILogMeasInfoR19SsbMeasResultListR19Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Ssb_MeasResultList_r19))); err != nil {
				return err
			}
			for i := range ie.Ssb_MeasResultList_r19 {
				if err := ie.Ssb_MeasResultList_r19[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 5. timeGap-r19: enumerated
	{
		if ie.TimeGap_r19 != nil {
			if err := e.EncodeEnumerated(*ie.TimeGap_r19, cSILogMeasInfoR19TimeGapR19Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *CSI_LogMeasInfo_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cSILogMeasInfoR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. csi-RS-MeasResultList-r19: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(cSILogMeasInfoR19CsiRSMeasResultListR19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Csi_RS_MeasResultList_r19 = make([]CSI_RS_MeasResult_r19, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Csi_RS_MeasResultList_r19[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 4. ssb-MeasResultList-r19: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(cSILogMeasInfoR19SsbMeasResultListR19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Ssb_MeasResultList_r19 = make([]SSB_MeasResult_r19, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Ssb_MeasResultList_r19[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 5. timeGap-r19: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(cSILogMeasInfoR19TimeGapR19Constraints)
			if err != nil {
				return err
			}
			ie.TimeGap_r19 = &idx
		}
	}

	return nil
}
