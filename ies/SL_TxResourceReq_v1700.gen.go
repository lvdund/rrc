// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SL-TxResourceReq-v1700 (line 2180).

var sLTxResourceReqV1700Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-DRX-InfoFromRxList-r17", Optional: true},
		{Name: "sl-DRX-Indication-r17", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

var sLTxResourceReqV1700SlDRXInfoFromRxListR17Constraints = per.SizeRange(1, common.MaxNrofSL_RxInfoSet_r17)

const (
	SL_TxResourceReq_v1700_Sl_DRX_Indication_r17_On  = 0
	SL_TxResourceReq_v1700_Sl_DRX_Indication_r17_Off = 1
)

var sLTxResourceReqV1700SlDRXIndicationR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_TxResourceReq_v1700_Sl_DRX_Indication_r17_On, SL_TxResourceReq_v1700_Sl_DRX_Indication_r17_Off},
}

var sLTxResourceReqV1700ExtSlQoSInfoListV1800Constraints = per.SizeRange(1, common.MaxNrofSL_QFIsPerDest_r16)

type SL_TxResourceReq_v1700 struct {
	Sl_DRX_InfoFromRxList_r17 []SL_DRX_ConfigUC_SemiStatic_r17
	Sl_DRX_Indication_r17     *int64
	Sl_QoS_InfoList_v1800     []SL_QoS_Info_v1800
}

func (ie *SL_TxResourceReq_v1700) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLTxResourceReqV1700Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Sl_QoS_InfoList_v1800 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_DRX_InfoFromRxList_r17 != nil, ie.Sl_DRX_Indication_r17 != nil}); err != nil {
		return err
	}

	// 3. sl-DRX-InfoFromRxList-r17: sequence-of
	{
		if ie.Sl_DRX_InfoFromRxList_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(sLTxResourceReqV1700SlDRXInfoFromRxListR17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_DRX_InfoFromRxList_r17))); err != nil {
				return err
			}
			for i := range ie.Sl_DRX_InfoFromRxList_r17 {
				if err := ie.Sl_DRX_InfoFromRxList_r17[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 4. sl-DRX-Indication-r17: enumerated
	{
		if ie.Sl_DRX_Indication_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_DRX_Indication_r17, sLTxResourceReqV1700SlDRXIndicationR17Constraints); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "sl-QoS-InfoList-v1800", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Sl_QoS_InfoList_v1800 != nil}); err != nil {
				return err
			}

			if ie.Sl_QoS_InfoList_v1800 != nil {
				seqOf := ex.NewSequenceOfEncoder(sLTxResourceReqV1700ExtSlQoSInfoListV1800Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.Sl_QoS_InfoList_v1800))); err != nil {
					return err
				}
				for i := range ie.Sl_QoS_InfoList_v1800 {
					if err := ie.Sl_QoS_InfoList_v1800[i].Encode(ex); err != nil {
						return err
					}
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		if err := seq.EncodeExtensionAdditions(extPresence, extValues); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SL_TxResourceReq_v1700) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLTxResourceReqV1700Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. sl-DRX-InfoFromRxList-r17: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(sLTxResourceReqV1700SlDRXInfoFromRxListR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_DRX_InfoFromRxList_r17 = make([]SL_DRX_ConfigUC_SemiStatic_r17, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_DRX_InfoFromRxList_r17[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 4. sl-DRX-Indication-r17: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(sLTxResourceReqV1700SlDRXIndicationR17Constraints)
			if err != nil {
				return err
			}
			ie.Sl_DRX_Indication_r17 = &idx
		}
	}

	// Extension additions as open types.
	extValues, err := seq.DecodeExtensionAdditions()
	if err != nil {
		return err
	}
	extValueIndex := 0

	// Extension group 0:
	if seq.IsExtensionPresent(0) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "sl-QoS-InfoList-v1800", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			seqOf := dx.NewSequenceOfDecoder(sLTxResourceReqV1700ExtSlQoSInfoListV1800Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_QoS_InfoList_v1800 = make([]SL_QoS_Info_v1800, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_QoS_InfoList_v1800[i].Decode(dx); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
