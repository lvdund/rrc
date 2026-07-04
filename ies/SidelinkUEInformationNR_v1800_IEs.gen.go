// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SidelinkUEInformationNR-v1800-IEs (line 2147).

var sidelinkUEInformationNRV1800IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-CarrierFailureList-r18", Optional: true},
		{Name: "sl-TxResourceReqListL2-U2U-r18", Optional: true},
		{Name: "sl-PosRxInterestedFreqList-r18", Optional: true},
		{Name: "sl-PosTxResourceReqList-r18", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

var sidelinkUEInformationNRV1800IEsSlTxResourceReqListL2U2UR18Constraints = per.SizeRange(1, common.MaxNrofSL_Dest_r16)

type SidelinkUEInformationNR_v1800_IEs struct {
	Sl_CarrierFailureList_r18      *SL_CarrierFailureList_r18
	Sl_TxResourceReqListL2_U2U_r18 []SL_TxResourceReqL2_U2U_r18
	Sl_PosRxInterestedFreqList_r18 *SL_InterestedFreqList_r16
	Sl_PosTxResourceReqList_r18    *SL_PosTxResourceReqList_r18
	NonCriticalExtension           *SidelinkUEInformationNR_v1840_IEs
}

func (ie *SidelinkUEInformationNR_v1800_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sidelinkUEInformationNRV1800IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_CarrierFailureList_r18 != nil, ie.Sl_TxResourceReqListL2_U2U_r18 != nil, ie.Sl_PosRxInterestedFreqList_r18 != nil, ie.Sl_PosTxResourceReqList_r18 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. sl-CarrierFailureList-r18: ref
	{
		if ie.Sl_CarrierFailureList_r18 != nil {
			if err := ie.Sl_CarrierFailureList_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. sl-TxResourceReqListL2-U2U-r18: sequence-of
	{
		if ie.Sl_TxResourceReqListL2_U2U_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(sidelinkUEInformationNRV1800IEsSlTxResourceReqListL2U2UR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_TxResourceReqListL2_U2U_r18))); err != nil {
				return err
			}
			for i := range ie.Sl_TxResourceReqListL2_U2U_r18 {
				if err := ie.Sl_TxResourceReqListL2_U2U_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 4. sl-PosRxInterestedFreqList-r18: ref
	{
		if ie.Sl_PosRxInterestedFreqList_r18 != nil {
			if err := ie.Sl_PosRxInterestedFreqList_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. sl-PosTxResourceReqList-r18: ref
	{
		if ie.Sl_PosTxResourceReqList_r18 != nil {
			if err := ie.Sl_PosTxResourceReqList_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. nonCriticalExtension: ref
	{
		if ie.NonCriticalExtension != nil {
			if err := ie.NonCriticalExtension.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SidelinkUEInformationNR_v1800_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sidelinkUEInformationNRV1800IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sl-CarrierFailureList-r18: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Sl_CarrierFailureList_r18 = new(SL_CarrierFailureList_r18)
			if err := ie.Sl_CarrierFailureList_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. sl-TxResourceReqListL2-U2U-r18: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(sidelinkUEInformationNRV1800IEsSlTxResourceReqListL2U2UR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_TxResourceReqListL2_U2U_r18 = make([]SL_TxResourceReqL2_U2U_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_TxResourceReqListL2_U2U_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 4. sl-PosRxInterestedFreqList-r18: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Sl_PosRxInterestedFreqList_r18 = new(SL_InterestedFreqList_r16)
			if err := ie.Sl_PosRxInterestedFreqList_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. sl-PosTxResourceReqList-r18: ref
	{
		if seq.IsComponentPresent(3) {
			ie.Sl_PosTxResourceReqList_r18 = new(SL_PosTxResourceReqList_r18)
			if err := ie.Sl_PosTxResourceReqList_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(4) {
			ie.NonCriticalExtension = new(SidelinkUEInformationNR_v1840_IEs)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
