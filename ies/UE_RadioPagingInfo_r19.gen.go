// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: UE-RadioPagingInfo-r19 (line 25955).

var uERadioPagingInfoR19Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "lpwus-SupportedBandList-r19", Optional: true},
		{Name: "pagingAdaptation-r19", Optional: true},
		{Name: "pagingAdaptationPEI-SupportBandList-r19", Optional: true},
	},
}

var uERadioPagingInfoR19LpwusSupportedBandListR19Constraints = per.SizeRange(1, common.MaxBands)

const (
	UE_RadioPagingInfo_r19_PagingAdaptation_r19_Supported = 0
)

var uERadioPagingInfoR19PagingAdaptationR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_RadioPagingInfo_r19_PagingAdaptation_r19_Supported},
}

var uERadioPagingInfoR19PagingAdaptationPEISupportBandListR19Constraints = per.SizeRange(1, common.MaxBands)

type UE_RadioPagingInfo_r19 struct {
	Lpwus_SupportedBandList_r19             []LPWUS_SupportedBandInfo_r19
	PagingAdaptation_r19                    *int64
	PagingAdaptationPEI_SupportBandList_r19 []FreqBandIndicatorNR
}

func (ie *UE_RadioPagingInfo_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uERadioPagingInfoR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Lpwus_SupportedBandList_r19 != nil, ie.PagingAdaptation_r19 != nil, ie.PagingAdaptationPEI_SupportBandList_r19 != nil}); err != nil {
		return err
	}

	// 3. lpwus-SupportedBandList-r19: sequence-of
	{
		if ie.Lpwus_SupportedBandList_r19 != nil {
			seqOf := e.NewSequenceOfEncoder(uERadioPagingInfoR19LpwusSupportedBandListR19Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Lpwus_SupportedBandList_r19))); err != nil {
				return err
			}
			for i := range ie.Lpwus_SupportedBandList_r19 {
				if err := ie.Lpwus_SupportedBandList_r19[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 4. pagingAdaptation-r19: enumerated
	{
		if ie.PagingAdaptation_r19 != nil {
			if err := e.EncodeEnumerated(*ie.PagingAdaptation_r19, uERadioPagingInfoR19PagingAdaptationR19Constraints); err != nil {
				return err
			}
		}
	}

	// 5. pagingAdaptationPEI-SupportBandList-r19: sequence-of
	{
		if ie.PagingAdaptationPEI_SupportBandList_r19 != nil {
			seqOf := e.NewSequenceOfEncoder(uERadioPagingInfoR19PagingAdaptationPEISupportBandListR19Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.PagingAdaptationPEI_SupportBandList_r19))); err != nil {
				return err
			}
			for i := range ie.PagingAdaptationPEI_SupportBandList_r19 {
				if err := ie.PagingAdaptationPEI_SupportBandList_r19[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *UE_RadioPagingInfo_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uERadioPagingInfoR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. lpwus-SupportedBandList-r19: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(uERadioPagingInfoR19LpwusSupportedBandListR19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Lpwus_SupportedBandList_r19 = make([]LPWUS_SupportedBandInfo_r19, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Lpwus_SupportedBandList_r19[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 4. pagingAdaptation-r19: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(uERadioPagingInfoR19PagingAdaptationR19Constraints)
			if err != nil {
				return err
			}
			ie.PagingAdaptation_r19 = &idx
		}
	}

	// 5. pagingAdaptationPEI-SupportBandList-r19: sequence-of
	{
		if seq.IsComponentPresent(2) {
			seqOf := d.NewSequenceOfDecoder(uERadioPagingInfoR19PagingAdaptationPEISupportBandListR19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.PagingAdaptationPEI_SupportBandList_r19 = make([]FreqBandIndicatorNR, n)
			for i := int64(0); i < n; i++ {
				if err := ie.PagingAdaptationPEI_SupportBandList_r19[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
