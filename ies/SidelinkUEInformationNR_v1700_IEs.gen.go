// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SidelinkUEInformationNR-v1700-IEs (line 2135).

var sidelinkUEInformationNRV1700IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-TxResourceReqList-v1700", Optional: true},
		{Name: "sl-RxDRX-ReportList-v1700", Optional: true},
		{Name: "sl-RxInterestedGC-BC-DestList-r17", Optional: true},
		{Name: "sl-RxInterestedFreqListDisc-r17", Optional: true},
		{Name: "sl-TxResourceReqListDisc-r17", Optional: true},
		{Name: "sl-TxResourceReqListCommRelay-r17", Optional: true},
		{Name: "ue-Type-r17", Optional: true},
		{Name: "sl-SourceIdentityRemoteUE-r17", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

const (
	SidelinkUEInformationNR_v1700_IEs_Ue_Type_r17_RelayUE  = 0
	SidelinkUEInformationNR_v1700_IEs_Ue_Type_r17_RemoteUE = 1
)

var sidelinkUEInformationNRV1700IEsUeTypeR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SidelinkUEInformationNR_v1700_IEs_Ue_Type_r17_RelayUE, SidelinkUEInformationNR_v1700_IEs_Ue_Type_r17_RemoteUE},
}

type SidelinkUEInformationNR_v1700_IEs struct {
	Sl_TxResourceReqList_v1700        *SL_TxResourceReqList_v1700
	Sl_RxDRX_ReportList_v1700         *SL_RxDRX_ReportList_v1700
	Sl_RxInterestedGC_BC_DestList_r17 *SL_RxInterestedGC_BC_DestList_r17
	Sl_RxInterestedFreqListDisc_r17   *SL_InterestedFreqList_r16
	Sl_TxResourceReqListDisc_r17      *SL_TxResourceReqListDisc_r17
	Sl_TxResourceReqListCommRelay_r17 *SL_TxResourceReqListCommRelay_r17
	Ue_Type_r17                       *int64
	Sl_SourceIdentityRemoteUE_r17     *SL_SourceIdentity_r17
	NonCriticalExtension              *SidelinkUEInformationNR_v1800_IEs
}

func (ie *SidelinkUEInformationNR_v1700_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sidelinkUEInformationNRV1700IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_TxResourceReqList_v1700 != nil, ie.Sl_RxDRX_ReportList_v1700 != nil, ie.Sl_RxInterestedGC_BC_DestList_r17 != nil, ie.Sl_RxInterestedFreqListDisc_r17 != nil, ie.Sl_TxResourceReqListDisc_r17 != nil, ie.Sl_TxResourceReqListCommRelay_r17 != nil, ie.Ue_Type_r17 != nil, ie.Sl_SourceIdentityRemoteUE_r17 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. sl-TxResourceReqList-v1700: ref
	{
		if ie.Sl_TxResourceReqList_v1700 != nil {
			if err := ie.Sl_TxResourceReqList_v1700.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. sl-RxDRX-ReportList-v1700: ref
	{
		if ie.Sl_RxDRX_ReportList_v1700 != nil {
			if err := ie.Sl_RxDRX_ReportList_v1700.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. sl-RxInterestedGC-BC-DestList-r17: ref
	{
		if ie.Sl_RxInterestedGC_BC_DestList_r17 != nil {
			if err := ie.Sl_RxInterestedGC_BC_DestList_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. sl-RxInterestedFreqListDisc-r17: ref
	{
		if ie.Sl_RxInterestedFreqListDisc_r17 != nil {
			if err := ie.Sl_RxInterestedFreqListDisc_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. sl-TxResourceReqListDisc-r17: ref
	{
		if ie.Sl_TxResourceReqListDisc_r17 != nil {
			if err := ie.Sl_TxResourceReqListDisc_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 7. sl-TxResourceReqListCommRelay-r17: ref
	{
		if ie.Sl_TxResourceReqListCommRelay_r17 != nil {
			if err := ie.Sl_TxResourceReqListCommRelay_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 8. ue-Type-r17: enumerated
	{
		if ie.Ue_Type_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Ue_Type_r17, sidelinkUEInformationNRV1700IEsUeTypeR17Constraints); err != nil {
				return err
			}
		}
	}

	// 9. sl-SourceIdentityRemoteUE-r17: ref
	{
		if ie.Sl_SourceIdentityRemoteUE_r17 != nil {
			if err := ie.Sl_SourceIdentityRemoteUE_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 10. nonCriticalExtension: ref
	{
		if ie.NonCriticalExtension != nil {
			if err := ie.NonCriticalExtension.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SidelinkUEInformationNR_v1700_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sidelinkUEInformationNRV1700IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sl-TxResourceReqList-v1700: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Sl_TxResourceReqList_v1700 = new(SL_TxResourceReqList_v1700)
			if err := ie.Sl_TxResourceReqList_v1700.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. sl-RxDRX-ReportList-v1700: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Sl_RxDRX_ReportList_v1700 = new(SL_RxDRX_ReportList_v1700)
			if err := ie.Sl_RxDRX_ReportList_v1700.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. sl-RxInterestedGC-BC-DestList-r17: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Sl_RxInterestedGC_BC_DestList_r17 = new(SL_RxInterestedGC_BC_DestList_r17)
			if err := ie.Sl_RxInterestedGC_BC_DestList_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. sl-RxInterestedFreqListDisc-r17: ref
	{
		if seq.IsComponentPresent(3) {
			ie.Sl_RxInterestedFreqListDisc_r17 = new(SL_InterestedFreqList_r16)
			if err := ie.Sl_RxInterestedFreqListDisc_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. sl-TxResourceReqListDisc-r17: ref
	{
		if seq.IsComponentPresent(4) {
			ie.Sl_TxResourceReqListDisc_r17 = new(SL_TxResourceReqListDisc_r17)
			if err := ie.Sl_TxResourceReqListDisc_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 7. sl-TxResourceReqListCommRelay-r17: ref
	{
		if seq.IsComponentPresent(5) {
			ie.Sl_TxResourceReqListCommRelay_r17 = new(SL_TxResourceReqListCommRelay_r17)
			if err := ie.Sl_TxResourceReqListCommRelay_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 8. ue-Type-r17: enumerated
	{
		if seq.IsComponentPresent(6) {
			idx, err := d.DecodeEnumerated(sidelinkUEInformationNRV1700IEsUeTypeR17Constraints)
			if err != nil {
				return err
			}
			ie.Ue_Type_r17 = &idx
		}
	}

	// 9. sl-SourceIdentityRemoteUE-r17: ref
	{
		if seq.IsComponentPresent(7) {
			ie.Sl_SourceIdentityRemoteUE_r17 = new(SL_SourceIdentity_r17)
			if err := ie.Sl_SourceIdentityRemoteUE_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 10. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(8) {
			ie.NonCriticalExtension = new(SidelinkUEInformationNR_v1800_IEs)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
