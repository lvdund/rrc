// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: ULInformationTransferMRDC-IEs (line 3683).

var uLInformationTransferMRDCIEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ul-DCCH-MessageNR", Optional: true},
		{Name: "ul-DCCH-MessageEUTRA", Optional: true},
		{Name: "lateNonCriticalExtension", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

var uLInformationTransferMRDCIEsUlDCCHMessageNRConstraints = per.SizeConstraints{}

var uLInformationTransferMRDCIEsUlDCCHMessageEUTRAConstraints = per.SizeConstraints{}

var uLInformationTransferMRDCIEsLateNonCriticalExtensionConstraints = per.SizeConstraints{}

type ULInformationTransferMRDC_IEs struct {
	Ul_DCCH_MessageNR        []byte
	Ul_DCCH_MessageEUTRA     []byte
	LateNonCriticalExtension []byte
	NonCriticalExtension     *struct{}
}

func (ie *ULInformationTransferMRDC_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uLInformationTransferMRDCIEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ul_DCCH_MessageNR != nil, ie.Ul_DCCH_MessageEUTRA != nil, ie.LateNonCriticalExtension != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. ul-DCCH-MessageNR: octet-string
	{
		if ie.Ul_DCCH_MessageNR != nil {
			if err := e.EncodeOctetString(ie.Ul_DCCH_MessageNR, uLInformationTransferMRDCIEsUlDCCHMessageNRConstraints); err != nil {
				return err
			}
		}
	}

	// 3. ul-DCCH-MessageEUTRA: octet-string
	{
		if ie.Ul_DCCH_MessageEUTRA != nil {
			if err := e.EncodeOctetString(ie.Ul_DCCH_MessageEUTRA, uLInformationTransferMRDCIEsUlDCCHMessageEUTRAConstraints); err != nil {
				return err
			}
		}
	}

	// 4. lateNonCriticalExtension: octet-string
	{
		if ie.LateNonCriticalExtension != nil {
			if err := e.EncodeOctetString(ie.LateNonCriticalExtension, uLInformationTransferMRDCIEsLateNonCriticalExtensionConstraints); err != nil {
				return err
			}
		}
	}

	// 5. nonCriticalExtension: sequence
	{
		if ie.NonCriticalExtension != nil {
			// empty SEQUENCE {}: no content to encode.
		}
	}

	return nil
}

func (ie *ULInformationTransferMRDC_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uLInformationTransferMRDCIEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. ul-DCCH-MessageNR: octet-string
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeOctetString(uLInformationTransferMRDCIEsUlDCCHMessageNRConstraints)
			if err != nil {
				return err
			}
			ie.Ul_DCCH_MessageNR = v
		}
	}

	// 3. ul-DCCH-MessageEUTRA: octet-string
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeOctetString(uLInformationTransferMRDCIEsUlDCCHMessageEUTRAConstraints)
			if err != nil {
				return err
			}
			ie.Ul_DCCH_MessageEUTRA = v
		}
	}

	// 4. lateNonCriticalExtension: octet-string
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeOctetString(uLInformationTransferMRDCIEsLateNonCriticalExtensionConstraints)
			if err != nil {
				return err
			}
			ie.LateNonCriticalExtension = v
		}
	}

	// 5. nonCriticalExtension: sequence
	{
		if seq.IsComponentPresent(3) {
			ie.NonCriticalExtension = &struct{}{}
		}
	}

	return nil
}
