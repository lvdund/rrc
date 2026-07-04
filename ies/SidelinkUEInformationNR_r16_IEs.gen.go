// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SidelinkUEInformationNR-r16-IEs (line 2127).

var sidelinkUEInformationNRR16IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-RxInterestedFreqList-r16", Optional: true},
		{Name: "sl-TxResourceReqList-r16", Optional: true},
		{Name: "sl-FailureList-r16", Optional: true},
		{Name: "lateNonCriticalExtension", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

var sidelinkUEInformationNRR16IEsLateNonCriticalExtensionConstraints = per.SizeConstraints{}

type SidelinkUEInformationNR_r16_IEs struct {
	Sl_RxInterestedFreqList_r16 *SL_InterestedFreqList_r16
	Sl_TxResourceReqList_r16    *SL_TxResourceReqList_r16
	Sl_FailureList_r16          *SL_FailureList_r16
	LateNonCriticalExtension    []byte
	NonCriticalExtension        *SidelinkUEInformationNR_v1700_IEs
}

func (ie *SidelinkUEInformationNR_r16_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sidelinkUEInformationNRR16IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_RxInterestedFreqList_r16 != nil, ie.Sl_TxResourceReqList_r16 != nil, ie.Sl_FailureList_r16 != nil, ie.LateNonCriticalExtension != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. sl-RxInterestedFreqList-r16: ref
	{
		if ie.Sl_RxInterestedFreqList_r16 != nil {
			if err := ie.Sl_RxInterestedFreqList_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. sl-TxResourceReqList-r16: ref
	{
		if ie.Sl_TxResourceReqList_r16 != nil {
			if err := ie.Sl_TxResourceReqList_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. sl-FailureList-r16: ref
	{
		if ie.Sl_FailureList_r16 != nil {
			if err := ie.Sl_FailureList_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. lateNonCriticalExtension: octet-string
	{
		if ie.LateNonCriticalExtension != nil {
			if err := e.EncodeOctetString(ie.LateNonCriticalExtension, sidelinkUEInformationNRR16IEsLateNonCriticalExtensionConstraints); err != nil {
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

func (ie *SidelinkUEInformationNR_r16_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sidelinkUEInformationNRR16IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sl-RxInterestedFreqList-r16: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Sl_RxInterestedFreqList_r16 = new(SL_InterestedFreqList_r16)
			if err := ie.Sl_RxInterestedFreqList_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. sl-TxResourceReqList-r16: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Sl_TxResourceReqList_r16 = new(SL_TxResourceReqList_r16)
			if err := ie.Sl_TxResourceReqList_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. sl-FailureList-r16: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Sl_FailureList_r16 = new(SL_FailureList_r16)
			if err := ie.Sl_FailureList_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. lateNonCriticalExtension: octet-string
	{
		if seq.IsComponentPresent(3) {
			v, err := d.DecodeOctetString(sidelinkUEInformationNRR16IEsLateNonCriticalExtensionConstraints)
			if err != nil {
				return err
			}
			ie.LateNonCriticalExtension = v
		}
	}

	// 6. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(4) {
			ie.NonCriticalExtension = new(SidelinkUEInformationNR_v1700_IEs)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
