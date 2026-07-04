// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SL-TxResourceReq-r16 (line 2166).

var sLTxResourceReqR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-DestinationIdentity-r16"},
		{Name: "sl-CastType-r16"},
		{Name: "sl-RLC-ModeIndicationList-r16", Optional: true},
		{Name: "sl-QoS-InfoList-r16", Optional: true},
		{Name: "sl-TypeTxSyncList-r16", Optional: true},
		{Name: "sl-TxInterestedFreqList-r16", Optional: true},
		{Name: "sl-CapabilityInformationSidelink-r16", Optional: true},
	},
}

const (
	SL_TxResourceReq_r16_Sl_CastType_r16_Broadcast = 0
	SL_TxResourceReq_r16_Sl_CastType_r16_Groupcast = 1
	SL_TxResourceReq_r16_Sl_CastType_r16_Unicast   = 2
	SL_TxResourceReq_r16_Sl_CastType_r16_Spare1    = 3
)

var sLTxResourceReqR16SlCastTypeR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_TxResourceReq_r16_Sl_CastType_r16_Broadcast, SL_TxResourceReq_r16_Sl_CastType_r16_Groupcast, SL_TxResourceReq_r16_Sl_CastType_r16_Unicast, SL_TxResourceReq_r16_Sl_CastType_r16_Spare1},
}

var sLTxResourceReqR16SlRLCModeIndicationListR16Constraints = per.SizeRange(1, common.MaxNrofSLRB_r16)

var sLTxResourceReqR16SlQoSInfoListR16Constraints = per.SizeRange(1, common.MaxNrofSL_QFIsPerDest_r16)

var sLTxResourceReqR16SlTypeTxSyncListR16Constraints = per.SizeRange(1, common.MaxNrofFreqSL_r16)

var sLTxResourceReqR16SlCapabilityInformationSidelinkR16Constraints = per.SizeConstraints{}

type SL_TxResourceReq_r16 struct {
	Sl_DestinationIdentity_r16           SL_DestinationIdentity_r16
	Sl_CastType_r16                      int64
	Sl_RLC_ModeIndicationList_r16        []SL_RLC_ModeIndication_r16
	Sl_QoS_InfoList_r16                  []SL_QoS_Info_r16
	Sl_TypeTxSyncList_r16                []SL_TypeTxSync_r16
	Sl_TxInterestedFreqList_r16          *SL_TxInterestedFreqList_r16
	Sl_CapabilityInformationSidelink_r16 []byte
}

func (ie *SL_TxResourceReq_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLTxResourceReqR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_RLC_ModeIndicationList_r16 != nil, ie.Sl_QoS_InfoList_r16 != nil, ie.Sl_TypeTxSyncList_r16 != nil, ie.Sl_TxInterestedFreqList_r16 != nil, ie.Sl_CapabilityInformationSidelink_r16 != nil}); err != nil {
		return err
	}

	// 2. sl-DestinationIdentity-r16: ref
	{
		if err := ie.Sl_DestinationIdentity_r16.Encode(e); err != nil {
			return err
		}
	}

	// 3. sl-CastType-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.Sl_CastType_r16, sLTxResourceReqR16SlCastTypeR16Constraints); err != nil {
			return err
		}
	}

	// 4. sl-RLC-ModeIndicationList-r16: sequence-of
	{
		if ie.Sl_RLC_ModeIndicationList_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(sLTxResourceReqR16SlRLCModeIndicationListR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_RLC_ModeIndicationList_r16))); err != nil {
				return err
			}
			for i := range ie.Sl_RLC_ModeIndicationList_r16 {
				if err := ie.Sl_RLC_ModeIndicationList_r16[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 5. sl-QoS-InfoList-r16: sequence-of
	{
		if ie.Sl_QoS_InfoList_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(sLTxResourceReqR16SlQoSInfoListR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_QoS_InfoList_r16))); err != nil {
				return err
			}
			for i := range ie.Sl_QoS_InfoList_r16 {
				if err := ie.Sl_QoS_InfoList_r16[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 6. sl-TypeTxSyncList-r16: sequence-of
	{
		if ie.Sl_TypeTxSyncList_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(sLTxResourceReqR16SlTypeTxSyncListR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_TypeTxSyncList_r16))); err != nil {
				return err
			}
			for i := range ie.Sl_TypeTxSyncList_r16 {
				if err := ie.Sl_TypeTxSyncList_r16[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 7. sl-TxInterestedFreqList-r16: ref
	{
		if ie.Sl_TxInterestedFreqList_r16 != nil {
			if err := ie.Sl_TxInterestedFreqList_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 8. sl-CapabilityInformationSidelink-r16: octet-string
	{
		if ie.Sl_CapabilityInformationSidelink_r16 != nil {
			if err := e.EncodeOctetString(ie.Sl_CapabilityInformationSidelink_r16, sLTxResourceReqR16SlCapabilityInformationSidelinkR16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SL_TxResourceReq_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLTxResourceReqR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sl-DestinationIdentity-r16: ref
	{
		if err := ie.Sl_DestinationIdentity_r16.Decode(d); err != nil {
			return err
		}
	}

	// 3. sl-CastType-r16: enumerated
	{
		v1, err := d.DecodeEnumerated(sLTxResourceReqR16SlCastTypeR16Constraints)
		if err != nil {
			return err
		}
		ie.Sl_CastType_r16 = v1
	}

	// 4. sl-RLC-ModeIndicationList-r16: sequence-of
	{
		if seq.IsComponentPresent(2) {
			seqOf := d.NewSequenceOfDecoder(sLTxResourceReqR16SlRLCModeIndicationListR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_RLC_ModeIndicationList_r16 = make([]SL_RLC_ModeIndication_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_RLC_ModeIndicationList_r16[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 5. sl-QoS-InfoList-r16: sequence-of
	{
		if seq.IsComponentPresent(3) {
			seqOf := d.NewSequenceOfDecoder(sLTxResourceReqR16SlQoSInfoListR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_QoS_InfoList_r16 = make([]SL_QoS_Info_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_QoS_InfoList_r16[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 6. sl-TypeTxSyncList-r16: sequence-of
	{
		if seq.IsComponentPresent(4) {
			seqOf := d.NewSequenceOfDecoder(sLTxResourceReqR16SlTypeTxSyncListR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_TypeTxSyncList_r16 = make([]SL_TypeTxSync_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_TypeTxSyncList_r16[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 7. sl-TxInterestedFreqList-r16: ref
	{
		if seq.IsComponentPresent(5) {
			ie.Sl_TxInterestedFreqList_r16 = new(SL_TxInterestedFreqList_r16)
			if err := ie.Sl_TxInterestedFreqList_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 8. sl-CapabilityInformationSidelink-r16: octet-string
	{
		if seq.IsComponentPresent(6) {
			v, err := d.DecodeOctetString(sLTxResourceReqR16SlCapabilityInformationSidelinkR16Constraints)
			if err != nil {
				return err
			}
			ie.Sl_CapabilityInformationSidelink_r16 = v
		}
	}

	return nil
}
