// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: ULInformationTransferIRAT-r16-IEs (line 3664).

var uLInformationTransferIRATR16IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ul-DCCH-MessageEUTRA-r16", Optional: true},
		{Name: "lateNonCriticalExtension", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

var uLInformationTransferIRATR16IEsUlDCCHMessageEUTRAR16Constraints = per.SizeConstraints{}

var uLInformationTransferIRATR16IEsLateNonCriticalExtensionConstraints = per.SizeConstraints{}

type ULInformationTransferIRAT_r16_IEs struct {
	Ul_DCCH_MessageEUTRA_r16 []byte
	LateNonCriticalExtension []byte
	NonCriticalExtension     *struct{}
}

func (ie *ULInformationTransferIRAT_r16_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uLInformationTransferIRATR16IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ul_DCCH_MessageEUTRA_r16 != nil, ie.LateNonCriticalExtension != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. ul-DCCH-MessageEUTRA-r16: octet-string
	{
		if ie.Ul_DCCH_MessageEUTRA_r16 != nil {
			if err := e.EncodeOctetString(ie.Ul_DCCH_MessageEUTRA_r16, uLInformationTransferIRATR16IEsUlDCCHMessageEUTRAR16Constraints); err != nil {
				return err
			}
		}
	}

	// 3. lateNonCriticalExtension: octet-string
	{
		if ie.LateNonCriticalExtension != nil {
			if err := e.EncodeOctetString(ie.LateNonCriticalExtension, uLInformationTransferIRATR16IEsLateNonCriticalExtensionConstraints); err != nil {
				return err
			}
		}
	}

	// 4. nonCriticalExtension: sequence
	{
		if ie.NonCriticalExtension != nil {
			// empty SEQUENCE {}: no content to encode.
		}
	}

	return nil
}

func (ie *ULInformationTransferIRAT_r16_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uLInformationTransferIRATR16IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. ul-DCCH-MessageEUTRA-r16: octet-string
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeOctetString(uLInformationTransferIRATR16IEsUlDCCHMessageEUTRAR16Constraints)
			if err != nil {
				return err
			}
			ie.Ul_DCCH_MessageEUTRA_r16 = v
		}
	}

	// 3. lateNonCriticalExtension: octet-string
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeOctetString(uLInformationTransferIRATR16IEsLateNonCriticalExtensionConstraints)
			if err != nil {
				return err
			}
			ie.LateNonCriticalExtension = v
		}
	}

	// 4. nonCriticalExtension: sequence
	{
		if seq.IsComponentPresent(2) {
			ie.NonCriticalExtension = &struct{}{}
		}
	}

	return nil
}
