// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RRCResumeComplete-v1720-IEs (line 1631).

var rRCResumeCompleteV1720IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "uplinkTxDirectCurrentMoreCarrierList-r17", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

type RRCResumeComplete_v1720_IEs struct {
	UplinkTxDirectCurrentMoreCarrierList_r17 *UplinkTxDirectCurrentMoreCarrierList_r17
	NonCriticalExtension                     *RRCResumeComplete_v1800_IEs
}

func (ie *RRCResumeComplete_v1720_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rRCResumeCompleteV1720IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.UplinkTxDirectCurrentMoreCarrierList_r17 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. uplinkTxDirectCurrentMoreCarrierList-r17: ref
	{
		if ie.UplinkTxDirectCurrentMoreCarrierList_r17 != nil {
			if err := ie.UplinkTxDirectCurrentMoreCarrierList_r17.Encode(e); err != nil {
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

func (ie *RRCResumeComplete_v1720_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rRCResumeCompleteV1720IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. uplinkTxDirectCurrentMoreCarrierList-r17: ref
	{
		if seq.IsComponentPresent(0) {
			ie.UplinkTxDirectCurrentMoreCarrierList_r17 = new(UplinkTxDirectCurrentMoreCarrierList_r17)
			if err := ie.UplinkTxDirectCurrentMoreCarrierList_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(1) {
			ie.NonCriticalExtension = new(RRCResumeComplete_v1800_IEs)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
