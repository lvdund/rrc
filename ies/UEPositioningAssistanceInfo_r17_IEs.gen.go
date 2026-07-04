// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UEPositioningAssistanceInfo-r17-IEs (line 3580).

var uEPositioningAssistanceInfoR17IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ue-TxTEG-AssociationList-r17", Optional: true},
		{Name: "lateNonCriticalExtension", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

var uEPositioningAssistanceInfoR17IEsLateNonCriticalExtensionConstraints = per.SizeConstraints{}

type UEPositioningAssistanceInfo_r17_IEs struct {
	Ue_TxTEG_AssociationList_r17 *UE_TxTEG_AssociationList_r17
	LateNonCriticalExtension     []byte
	NonCriticalExtension         *UEPositioningAssistanceInfo_v1720_IEs
}

func (ie *UEPositioningAssistanceInfo_r17_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uEPositioningAssistanceInfoR17IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ue_TxTEG_AssociationList_r17 != nil, ie.LateNonCriticalExtension != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. ue-TxTEG-AssociationList-r17: ref
	{
		if ie.Ue_TxTEG_AssociationList_r17 != nil {
			if err := ie.Ue_TxTEG_AssociationList_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. lateNonCriticalExtension: octet-string
	{
		if ie.LateNonCriticalExtension != nil {
			if err := e.EncodeOctetString(ie.LateNonCriticalExtension, uEPositioningAssistanceInfoR17IEsLateNonCriticalExtensionConstraints); err != nil {
				return err
			}
		}
	}

	// 4. nonCriticalExtension: ref
	{
		if ie.NonCriticalExtension != nil {
			if err := ie.NonCriticalExtension.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *UEPositioningAssistanceInfo_r17_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uEPositioningAssistanceInfoR17IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. ue-TxTEG-AssociationList-r17: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Ue_TxTEG_AssociationList_r17 = new(UE_TxTEG_AssociationList_r17)
			if err := ie.Ue_TxTEG_AssociationList_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. lateNonCriticalExtension: octet-string
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeOctetString(uEPositioningAssistanceInfoR17IEsLateNonCriticalExtensionConstraints)
			if err != nil {
				return err
			}
			ie.LateNonCriticalExtension = v
		}
	}

	// 4. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(2) {
			ie.NonCriticalExtension = new(UEPositioningAssistanceInfo_v1720_IEs)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
