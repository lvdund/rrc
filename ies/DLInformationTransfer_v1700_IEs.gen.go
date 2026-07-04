// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: DLInformationTransfer-v1700-IEs (line 357).

var dLInformationTransferV1700IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "dedicatedInfoF1c-r17", Optional: true},
		{Name: "rxTxTimeDiff-gNB-r17", Optional: true},
		{Name: "ta-PDC-r17", Optional: true},
		{Name: "sib9Fallback-r17", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

const (
	DLInformationTransfer_v1700_IEs_Ta_PDC_r17_Activate   = 0
	DLInformationTransfer_v1700_IEs_Ta_PDC_r17_Deactivate = 1
)

var dLInformationTransferV1700IEsTaPDCR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DLInformationTransfer_v1700_IEs_Ta_PDC_r17_Activate, DLInformationTransfer_v1700_IEs_Ta_PDC_r17_Deactivate},
}

const (
	DLInformationTransfer_v1700_IEs_Sib9Fallback_r17_True = 0
)

var dLInformationTransferV1700IEsSib9FallbackR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DLInformationTransfer_v1700_IEs_Sib9Fallback_r17_True},
}

type DLInformationTransfer_v1700_IEs struct {
	DedicatedInfoF1c_r17 *DedicatedInfoF1c_r17
	RxTxTimeDiff_GNB_r17 *RxTxTimeDiff_r17
	Ta_PDC_r17           *int64
	Sib9Fallback_r17     *int64
	NonCriticalExtension *DLInformationTransfer_v1800_IEs
}

func (ie *DLInformationTransfer_v1700_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(dLInformationTransferV1700IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.DedicatedInfoF1c_r17 != nil, ie.RxTxTimeDiff_GNB_r17 != nil, ie.Ta_PDC_r17 != nil, ie.Sib9Fallback_r17 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. dedicatedInfoF1c-r17: ref
	{
		if ie.DedicatedInfoF1c_r17 != nil {
			if err := ie.DedicatedInfoF1c_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. rxTxTimeDiff-gNB-r17: ref
	{
		if ie.RxTxTimeDiff_GNB_r17 != nil {
			if err := ie.RxTxTimeDiff_GNB_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. ta-PDC-r17: enumerated
	{
		if ie.Ta_PDC_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Ta_PDC_r17, dLInformationTransferV1700IEsTaPDCR17Constraints); err != nil {
				return err
			}
		}
	}

	// 5. sib9Fallback-r17: enumerated
	{
		if ie.Sib9Fallback_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Sib9Fallback_r17, dLInformationTransferV1700IEsSib9FallbackR17Constraints); err != nil {
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

func (ie *DLInformationTransfer_v1700_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(dLInformationTransferV1700IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. dedicatedInfoF1c-r17: ref
	{
		if seq.IsComponentPresent(0) {
			ie.DedicatedInfoF1c_r17 = new(DedicatedInfoF1c_r17)
			if err := ie.DedicatedInfoF1c_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. rxTxTimeDiff-gNB-r17: ref
	{
		if seq.IsComponentPresent(1) {
			ie.RxTxTimeDiff_GNB_r17 = new(RxTxTimeDiff_r17)
			if err := ie.RxTxTimeDiff_GNB_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. ta-PDC-r17: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(dLInformationTransferV1700IEsTaPDCR17Constraints)
			if err != nil {
				return err
			}
			ie.Ta_PDC_r17 = &idx
		}
	}

	// 5. sib9Fallback-r17: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(dLInformationTransferV1700IEsSib9FallbackR17Constraints)
			if err != nil {
				return err
			}
			ie.Sib9Fallback_r17 = &idx
		}
	}

	// 6. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(4) {
			ie.NonCriticalExtension = new(DLInformationTransfer_v1800_IEs)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
