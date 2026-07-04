// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SL-TxResourceReqL2-U2U-r18 (line 2241).

var sLTxResourceReqL2U2UR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-DestinationIdentityL2-U2U-r18", Optional: true},
		{Name: "sl-TxInterestedFreqListL2-U2U-r18"},
		{Name: "sl-TypeTxSyncListL2-U2U-r18"},
		{Name: "sl-CapabilityInformationSidelink-r18", Optional: true},
		{Name: "sl-U2U-InfoList-r18", Optional: true},
		{Name: "sl-RLC-ModeIndicationListL2-U2U-r18", Optional: true},
	},
}

var sLTxResourceReqL2U2UR18SlTypeTxSyncListL2U2UR18Constraints = per.SizeRange(1, common.MaxNrofFreqSL_r16)

var sLTxResourceReqL2U2UR18SlCapabilityInformationSidelinkR18Constraints = per.SizeConstraints{}

var sLTxResourceReqL2U2UR18SlU2UInfoListR18Constraints = per.SizeRange(1, common.MaxNrofRemoteUE_r17)

var sLTxResourceReqL2U2UR18SlRLCModeIndicationListL2U2UR18Constraints = per.SizeRange(1, common.MaxNrofSLRB_r16)

type SL_TxResourceReqL2_U2U_r18 struct {
	Sl_DestinationIdentityL2_U2U_r18     *SL_DestinationIdentity_r16
	Sl_TxInterestedFreqListL2_U2U_r18    SL_TxInterestedFreqList_r16
	Sl_TypeTxSyncListL2_U2U_r18          []SL_TypeTxSync_r16
	Sl_CapabilityInformationSidelink_r18 []byte
	Sl_U2U_InfoList_r18                  []SL_U2U_Info_r18
	Sl_RLC_ModeIndicationListL2_U2U_r18  []SL_RLC_Mode_r18
}

func (ie *SL_TxResourceReqL2_U2U_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLTxResourceReqL2U2UR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_DestinationIdentityL2_U2U_r18 != nil, ie.Sl_CapabilityInformationSidelink_r18 != nil, ie.Sl_U2U_InfoList_r18 != nil, ie.Sl_RLC_ModeIndicationListL2_U2U_r18 != nil}); err != nil {
		return err
	}

	// 3. sl-DestinationIdentityL2-U2U-r18: ref
	{
		if ie.Sl_DestinationIdentityL2_U2U_r18 != nil {
			if err := ie.Sl_DestinationIdentityL2_U2U_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. sl-TxInterestedFreqListL2-U2U-r18: ref
	{
		if err := ie.Sl_TxInterestedFreqListL2_U2U_r18.Encode(e); err != nil {
			return err
		}
	}

	// 5. sl-TypeTxSyncListL2-U2U-r18: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(sLTxResourceReqL2U2UR18SlTypeTxSyncListL2U2UR18Constraints)
		if err := seqOf.EncodeLength(int64(len(ie.Sl_TypeTxSyncListL2_U2U_r18))); err != nil {
			return err
		}
		for i := range ie.Sl_TypeTxSyncListL2_U2U_r18 {
			if err := ie.Sl_TypeTxSyncListL2_U2U_r18[i].Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. sl-CapabilityInformationSidelink-r18: octet-string
	{
		if ie.Sl_CapabilityInformationSidelink_r18 != nil {
			if err := e.EncodeOctetString(ie.Sl_CapabilityInformationSidelink_r18, sLTxResourceReqL2U2UR18SlCapabilityInformationSidelinkR18Constraints); err != nil {
				return err
			}
		}
	}

	// 7. sl-U2U-InfoList-r18: sequence-of
	{
		if ie.Sl_U2U_InfoList_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(sLTxResourceReqL2U2UR18SlU2UInfoListR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_U2U_InfoList_r18))); err != nil {
				return err
			}
			for i := range ie.Sl_U2U_InfoList_r18 {
				if err := ie.Sl_U2U_InfoList_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 8. sl-RLC-ModeIndicationListL2-U2U-r18: sequence-of
	{
		if ie.Sl_RLC_ModeIndicationListL2_U2U_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(sLTxResourceReqL2U2UR18SlRLCModeIndicationListL2U2UR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_RLC_ModeIndicationListL2_U2U_r18))); err != nil {
				return err
			}
			for i := range ie.Sl_RLC_ModeIndicationListL2_U2U_r18 {
				if err := ie.Sl_RLC_ModeIndicationListL2_U2U_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *SL_TxResourceReqL2_U2U_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLTxResourceReqL2U2UR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. sl-DestinationIdentityL2-U2U-r18: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Sl_DestinationIdentityL2_U2U_r18 = new(SL_DestinationIdentity_r16)
			if err := ie.Sl_DestinationIdentityL2_U2U_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. sl-TxInterestedFreqListL2-U2U-r18: ref
	{
		if err := ie.Sl_TxInterestedFreqListL2_U2U_r18.Decode(d); err != nil {
			return err
		}
	}

	// 5. sl-TypeTxSyncListL2-U2U-r18: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(sLTxResourceReqL2U2UR18SlTypeTxSyncListL2U2UR18Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.Sl_TypeTxSyncListL2_U2U_r18 = make([]SL_TypeTxSync_r16, n)
		for i := int64(0); i < n; i++ {
			if err := ie.Sl_TypeTxSyncListL2_U2U_r18[i].Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. sl-CapabilityInformationSidelink-r18: octet-string
	{
		if seq.IsComponentPresent(3) {
			v, err := d.DecodeOctetString(sLTxResourceReqL2U2UR18SlCapabilityInformationSidelinkR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_CapabilityInformationSidelink_r18 = v
		}
	}

	// 7. sl-U2U-InfoList-r18: sequence-of
	{
		if seq.IsComponentPresent(4) {
			seqOf := d.NewSequenceOfDecoder(sLTxResourceReqL2U2UR18SlU2UInfoListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_U2U_InfoList_r18 = make([]SL_U2U_Info_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_U2U_InfoList_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 8. sl-RLC-ModeIndicationListL2-U2U-r18: sequence-of
	{
		if seq.IsComponentPresent(5) {
			seqOf := d.NewSequenceOfDecoder(sLTxResourceReqL2U2UR18SlRLCModeIndicationListL2U2UR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_RLC_ModeIndicationListL2_U2U_r18 = make([]SL_RLC_Mode_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_RLC_ModeIndicationListL2_U2U_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
