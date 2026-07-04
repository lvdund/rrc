// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: DLInformationTransferMRDC-r16-IEs (line 387).

var dLInformationTransferMRDCR16IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "dl-DCCH-MessageNR-r16", Optional: true},
		{Name: "dl-DCCH-MessageEUTRA-r16", Optional: true},
		{Name: "lateNonCriticalExtension", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

var dLInformationTransferMRDCR16IEsDlDCCHMessageNRR16Constraints = per.SizeConstraints{}

var dLInformationTransferMRDCR16IEsDlDCCHMessageEUTRAR16Constraints = per.SizeConstraints{}

var dLInformationTransferMRDCR16IEsLateNonCriticalExtensionConstraints = per.SizeConstraints{}

type DLInformationTransferMRDC_r16_IEs struct {
	Dl_DCCH_MessageNR_r16    []byte
	Dl_DCCH_MessageEUTRA_r16 []byte
	LateNonCriticalExtension []byte
	NonCriticalExtension     *struct{}
}

func (ie *DLInformationTransferMRDC_r16_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(dLInformationTransferMRDCR16IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Dl_DCCH_MessageNR_r16 != nil, ie.Dl_DCCH_MessageEUTRA_r16 != nil, ie.LateNonCriticalExtension != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. dl-DCCH-MessageNR-r16: octet-string
	{
		if ie.Dl_DCCH_MessageNR_r16 != nil {
			if err := e.EncodeOctetString(ie.Dl_DCCH_MessageNR_r16, dLInformationTransferMRDCR16IEsDlDCCHMessageNRR16Constraints); err != nil {
				return err
			}
		}
	}

	// 3. dl-DCCH-MessageEUTRA-r16: octet-string
	{
		if ie.Dl_DCCH_MessageEUTRA_r16 != nil {
			if err := e.EncodeOctetString(ie.Dl_DCCH_MessageEUTRA_r16, dLInformationTransferMRDCR16IEsDlDCCHMessageEUTRAR16Constraints); err != nil {
				return err
			}
		}
	}

	// 4. lateNonCriticalExtension: octet-string
	{
		if ie.LateNonCriticalExtension != nil {
			if err := e.EncodeOctetString(ie.LateNonCriticalExtension, dLInformationTransferMRDCR16IEsLateNonCriticalExtensionConstraints); err != nil {
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

func (ie *DLInformationTransferMRDC_r16_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(dLInformationTransferMRDCR16IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. dl-DCCH-MessageNR-r16: octet-string
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeOctetString(dLInformationTransferMRDCR16IEsDlDCCHMessageNRR16Constraints)
			if err != nil {
				return err
			}
			ie.Dl_DCCH_MessageNR_r16 = v
		}
	}

	// 3. dl-DCCH-MessageEUTRA-r16: octet-string
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeOctetString(dLInformationTransferMRDCR16IEsDlDCCHMessageEUTRAR16Constraints)
			if err != nil {
				return err
			}
			ie.Dl_DCCH_MessageEUTRA_r16 = v
		}
	}

	// 4. lateNonCriticalExtension: octet-string
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeOctetString(dLInformationTransferMRDCR16IEsLateNonCriticalExtensionConstraints)
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
