// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SL-PosTxResourceReq-r18 (line 2262).

var sLPosTxResourceReqR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-PosDestinationIdentity-r18"},
		{Name: "sl-PosCastType-r18"},
		{Name: "sl-PosTxInterestedFreqList-r18", Optional: true},
		{Name: "sl-PosTypeTxSyncList-r18", Optional: true},
		{Name: "sl-PosQoS-InfoList-r18", Optional: true},
		{Name: "sl-CapabilityInformationSidelink-r18", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

const (
	SL_PosTxResourceReq_r18_Sl_PosCastType_r18_Broadcast = 0
	SL_PosTxResourceReq_r18_Sl_PosCastType_r18_Groupcast = 1
	SL_PosTxResourceReq_r18_Sl_PosCastType_r18_Unicast   = 2
	SL_PosTxResourceReq_r18_Sl_PosCastType_r18_Spare1    = 3
)

var sLPosTxResourceReqR18SlPosCastTypeR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_PosTxResourceReq_r18_Sl_PosCastType_r18_Broadcast, SL_PosTxResourceReq_r18_Sl_PosCastType_r18_Groupcast, SL_PosTxResourceReq_r18_Sl_PosCastType_r18_Unicast, SL_PosTxResourceReq_r18_Sl_PosCastType_r18_Spare1},
}

var sLPosTxResourceReqR18SlPosTypeTxSyncListR18Constraints = per.SizeRange(1, common.MaxNrofFreqSL_r16)

var sLPosTxResourceReqR18SlPosQoSInfoListR18Constraints = per.SizeRange(1, common.MaxNrofSL_PRS_PerDest_r18)

var sLPosTxResourceReqR18SlCapabilityInformationSidelinkR18Constraints = per.SizeConstraints{}

type SL_PosTxResourceReq_r18 struct {
	Sl_PosDestinationIdentity_r18        SL_DestinationIdentity_r16
	Sl_PosCastType_r18                   int64
	Sl_PosTxInterestedFreqList_r18       *SL_TxInterestedFreqList_r16
	Sl_PosTypeTxSyncList_r18             []SL_TypeTxSync_r16
	Sl_PosQoS_InfoList_r18               []SL_PRS_QoS_Info_r18
	Sl_CapabilityInformationSidelink_r18 []byte
	Sl_PosTxInterestedFreqList2_r18      *SL_TxInterestedFreqList_r16
}

func (ie *SL_PosTxResourceReq_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLPosTxResourceReqR18Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Sl_PosTxInterestedFreqList2_r18 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_PosTxInterestedFreqList_r18 != nil, ie.Sl_PosTypeTxSyncList_r18 != nil, ie.Sl_PosQoS_InfoList_r18 != nil, ie.Sl_CapabilityInformationSidelink_r18 != nil}); err != nil {
		return err
	}

	// 3. sl-PosDestinationIdentity-r18: ref
	{
		if err := ie.Sl_PosDestinationIdentity_r18.Encode(e); err != nil {
			return err
		}
	}

	// 4. sl-PosCastType-r18: enumerated
	{
		if err := e.EncodeEnumerated(ie.Sl_PosCastType_r18, sLPosTxResourceReqR18SlPosCastTypeR18Constraints); err != nil {
			return err
		}
	}

	// 5. sl-PosTxInterestedFreqList-r18: ref
	{
		if ie.Sl_PosTxInterestedFreqList_r18 != nil {
			if err := ie.Sl_PosTxInterestedFreqList_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. sl-PosTypeTxSyncList-r18: sequence-of
	{
		if ie.Sl_PosTypeTxSyncList_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(sLPosTxResourceReqR18SlPosTypeTxSyncListR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_PosTypeTxSyncList_r18))); err != nil {
				return err
			}
			for i := range ie.Sl_PosTypeTxSyncList_r18 {
				if err := ie.Sl_PosTypeTxSyncList_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 7. sl-PosQoS-InfoList-r18: sequence-of
	{
		if ie.Sl_PosQoS_InfoList_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(sLPosTxResourceReqR18SlPosQoSInfoListR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_PosQoS_InfoList_r18))); err != nil {
				return err
			}
			for i := range ie.Sl_PosQoS_InfoList_r18 {
				if err := ie.Sl_PosQoS_InfoList_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 8. sl-CapabilityInformationSidelink-r18: octet-string
	{
		if ie.Sl_CapabilityInformationSidelink_r18 != nil {
			if err := e.EncodeOctetString(ie.Sl_CapabilityInformationSidelink_r18, sLPosTxResourceReqR18SlCapabilityInformationSidelinkR18Constraints); err != nil {
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
					{Name: "sl-PosTxInterestedFreqList2-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Sl_PosTxInterestedFreqList2_r18 != nil}); err != nil {
				return err
			}

			if ie.Sl_PosTxInterestedFreqList2_r18 != nil {
				if err := ie.Sl_PosTxInterestedFreqList2_r18.Encode(ex); err != nil {
					return err
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

func (ie *SL_PosTxResourceReq_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLPosTxResourceReqR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. sl-PosDestinationIdentity-r18: ref
	{
		if err := ie.Sl_PosDestinationIdentity_r18.Decode(d); err != nil {
			return err
		}
	}

	// 4. sl-PosCastType-r18: enumerated
	{
		v1, err := d.DecodeEnumerated(sLPosTxResourceReqR18SlPosCastTypeR18Constraints)
		if err != nil {
			return err
		}
		ie.Sl_PosCastType_r18 = v1
	}

	// 5. sl-PosTxInterestedFreqList-r18: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Sl_PosTxInterestedFreqList_r18 = new(SL_TxInterestedFreqList_r16)
			if err := ie.Sl_PosTxInterestedFreqList_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. sl-PosTypeTxSyncList-r18: sequence-of
	{
		if seq.IsComponentPresent(3) {
			seqOf := d.NewSequenceOfDecoder(sLPosTxResourceReqR18SlPosTypeTxSyncListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_PosTypeTxSyncList_r18 = make([]SL_TypeTxSync_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_PosTypeTxSyncList_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 7. sl-PosQoS-InfoList-r18: sequence-of
	{
		if seq.IsComponentPresent(4) {
			seqOf := d.NewSequenceOfDecoder(sLPosTxResourceReqR18SlPosQoSInfoListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_PosQoS_InfoList_r18 = make([]SL_PRS_QoS_Info_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_PosQoS_InfoList_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 8. sl-CapabilityInformationSidelink-r18: octet-string
	{
		if seq.IsComponentPresent(5) {
			v, err := d.DecodeOctetString(sLPosTxResourceReqR18SlCapabilityInformationSidelinkR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_CapabilityInformationSidelink_r18 = v
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
				{Name: "sl-PosTxInterestedFreqList2-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Sl_PosTxInterestedFreqList2_r18 = new(SL_TxInterestedFreqList_r16)
			if err := ie.Sl_PosTxInterestedFreqList2_r18.Decode(dx); err != nil {
				return err
			}
		}
	}

	return nil
}
