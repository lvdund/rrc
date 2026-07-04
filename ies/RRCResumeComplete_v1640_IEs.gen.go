// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RRCResumeComplete-v1640-IEs (line 1620).

var rRCResumeCompleteV1640IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "uplinkTxDirectCurrentTwoCarrierList-r16", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

type RRCResumeComplete_v1640_IEs struct {
	UplinkTxDirectCurrentTwoCarrierList_r16 *UplinkTxDirectCurrentTwoCarrierList_r16
	NonCriticalExtension                    *RRCResumeComplete_v1700_IEs
}

func (ie *RRCResumeComplete_v1640_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rRCResumeCompleteV1640IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.UplinkTxDirectCurrentTwoCarrierList_r16 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. uplinkTxDirectCurrentTwoCarrierList-r16: ref
	{
		if ie.UplinkTxDirectCurrentTwoCarrierList_r16 != nil {
			if err := ie.UplinkTxDirectCurrentTwoCarrierList_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. nonCriticalExtension: ref
	{
		if ie.NonCriticalExtension != nil {
			if err := ie.NonCriticalExtension.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *RRCResumeComplete_v1640_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rRCResumeCompleteV1640IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. uplinkTxDirectCurrentTwoCarrierList-r16: ref
	{
		if seq.IsComponentPresent(0) {
			ie.UplinkTxDirectCurrentTwoCarrierList_r16 = new(UplinkTxDirectCurrentTwoCarrierList_r16)
			if err := ie.UplinkTxDirectCurrentTwoCarrierList_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(1) {
			ie.NonCriticalExtension = new(RRCResumeComplete_v1700_IEs)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
