// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: InterFreqCarrierFreqInfo-v1800 (line 4035).

var interFreqCarrierFreqInfoV1800Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "dl-CarrierFreq-r18", Optional: true},
		{Name: "frequencyBandList-r18", Optional: true},
		{Name: "frequencyBandListAerial-r18", Optional: true},
		{Name: "mobileIAB-CellList-r18", Optional: true},
		{Name: "mobileIAB-Freq-r18", Optional: true},
		{Name: "eRedCapAccessAllowed-r18", Optional: true},
		{Name: "tn-AreaIdList-r18", Optional: true},
		{Name: "accessAllowed2RxXR-r18", Optional: true},
	},
}

const (
	InterFreqCarrierFreqInfo_v1800_MobileIAB_Freq_r18_True = 0
)

var interFreqCarrierFreqInfoV1800MobileIABFreqR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{InterFreqCarrierFreqInfo_v1800_MobileIAB_Freq_r18_True},
}

const (
	InterFreqCarrierFreqInfo_v1800_ERedCapAccessAllowed_r18_True = 0
)

var interFreqCarrierFreqInfoV1800ERedCapAccessAllowedR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{InterFreqCarrierFreqInfo_v1800_ERedCapAccessAllowed_r18_True},
}

var interFreqCarrierFreqInfoV1800TnAreaIdListR18Constraints = per.SizeRange(1, common.MaxTN_AreaInfo_r18)

const (
	InterFreqCarrierFreqInfo_v1800_AccessAllowed2RxXR_r18_True = 0
)

var interFreqCarrierFreqInfoV1800AccessAllowed2RxXRR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{InterFreqCarrierFreqInfo_v1800_AccessAllowed2RxXR_r18_True},
}

type InterFreqCarrierFreqInfo_v1800 struct {
	Dl_CarrierFreq_r18          *ARFCN_ValueNR
	FrequencyBandList_r18       *MultiFrequencyBandListNR_SIB
	FrequencyBandListAerial_r18 *MultiFrequencyBandListNR_Aerial_SIB_r18
	MobileIAB_CellList_r18      *PCI_Range
	MobileIAB_Freq_r18          *int64
	ERedCapAccessAllowed_r18    *int64
	Tn_AreaIdList_r18           []TN_AreaId_r18
	AccessAllowed2RxXR_r18      *int64
}

func (ie *InterFreqCarrierFreqInfo_v1800) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(interFreqCarrierFreqInfoV1800Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Dl_CarrierFreq_r18 != nil, ie.FrequencyBandList_r18 != nil, ie.FrequencyBandListAerial_r18 != nil, ie.MobileIAB_CellList_r18 != nil, ie.MobileIAB_Freq_r18 != nil, ie.ERedCapAccessAllowed_r18 != nil, ie.Tn_AreaIdList_r18 != nil, ie.AccessAllowed2RxXR_r18 != nil}); err != nil {
		return err
	}

	// 2. dl-CarrierFreq-r18: ref
	{
		if ie.Dl_CarrierFreq_r18 != nil {
			if err := ie.Dl_CarrierFreq_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. frequencyBandList-r18: ref
	{
		if ie.FrequencyBandList_r18 != nil {
			if err := ie.FrequencyBandList_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. frequencyBandListAerial-r18: ref
	{
		if ie.FrequencyBandListAerial_r18 != nil {
			if err := ie.FrequencyBandListAerial_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. mobileIAB-CellList-r18: ref
	{
		if ie.MobileIAB_CellList_r18 != nil {
			if err := ie.MobileIAB_CellList_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. mobileIAB-Freq-r18: enumerated
	{
		if ie.MobileIAB_Freq_r18 != nil {
			if err := e.EncodeEnumerated(*ie.MobileIAB_Freq_r18, interFreqCarrierFreqInfoV1800MobileIABFreqR18Constraints); err != nil {
				return err
			}
		}
	}

	// 7. eRedCapAccessAllowed-r18: enumerated
	{
		if ie.ERedCapAccessAllowed_r18 != nil {
			if err := e.EncodeEnumerated(*ie.ERedCapAccessAllowed_r18, interFreqCarrierFreqInfoV1800ERedCapAccessAllowedR18Constraints); err != nil {
				return err
			}
		}
	}

	// 8. tn-AreaIdList-r18: sequence-of
	{
		if ie.Tn_AreaIdList_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(interFreqCarrierFreqInfoV1800TnAreaIdListR18Constraints)
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

	// 9. accessAllowed2RxXR-r18: enumerated
	{
		if ie.AccessAllowed2RxXR_r18 != nil {
			if err := e.EncodeEnumerated(*ie.AccessAllowed2RxXR_r18, interFreqCarrierFreqInfoV1800AccessAllowed2RxXRR18Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *InterFreqCarrierFreqInfo_v1800) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(interFreqCarrierFreqInfoV1800Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. dl-CarrierFreq-r18: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Dl_CarrierFreq_r18 = new(ARFCN_ValueNR)
			if err := ie.Dl_CarrierFreq_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. frequencyBandList-r18: ref
	{
		if seq.IsComponentPresent(1) {
			ie.FrequencyBandList_r18 = new(MultiFrequencyBandListNR_SIB)
			if err := ie.FrequencyBandList_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. frequencyBandListAerial-r18: ref
	{
		if seq.IsComponentPresent(2) {
			ie.FrequencyBandListAerial_r18 = new(MultiFrequencyBandListNR_Aerial_SIB_r18)
			if err := ie.FrequencyBandListAerial_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. mobileIAB-CellList-r18: ref
	{
		if seq.IsComponentPresent(3) {
			ie.MobileIAB_CellList_r18 = new(PCI_Range)
			if err := ie.MobileIAB_CellList_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. mobileIAB-Freq-r18: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(interFreqCarrierFreqInfoV1800MobileIABFreqR18Constraints)
			if err != nil {
				return err
			}
			ie.MobileIAB_Freq_r18 = &idx
		}
	}

	// 7. eRedCapAccessAllowed-r18: enumerated
	{
		if seq.IsComponentPresent(5) {
			idx, err := d.DecodeEnumerated(interFreqCarrierFreqInfoV1800ERedCapAccessAllowedR18Constraints)
			if err != nil {
				return err
			}
			ie.ERedCapAccessAllowed_r18 = &idx
		}
	}

	// 8. tn-AreaIdList-r18: sequence-of
	{
		if seq.IsComponentPresent(6) {
			seqOf := d.NewSequenceOfDecoder(interFreqCarrierFreqInfoV1800TnAreaIdListR18Constraints)
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

	// 9. accessAllowed2RxXR-r18: enumerated
	{
		if seq.IsComponentPresent(7) {
			idx, err := d.DecodeEnumerated(interFreqCarrierFreqInfoV1800AccessAllowed2RxXRR18Constraints)
			if err != nil {
				return err
			}
			ie.AccessAllowed2RxXR_r18 = &idx
		}
	}

	return nil
}
