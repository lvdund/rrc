// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RRCReconfigurationComplete-v1720-IEs (line 1180).

var rRCReconfigurationCompleteV1720IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "uplinkTxDirectCurrentMoreCarrierList-r17", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

type RRCReconfigurationComplete_v1720_IEs struct {
	UplinkTxDirectCurrentMoreCarrierList_r17 *UplinkTxDirectCurrentMoreCarrierList_r17
	NonCriticalExtension                     *RRCReconfigurationComplete_v1800_IEs
}

func (ie *RRCReconfigurationComplete_v1720_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rRCReconfigurationCompleteV1720IEsConstraints)

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

func (ie *RRCReconfigurationComplete_v1720_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rRCReconfigurationCompleteV1720IEsConstraints)

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
			ie.NonCriticalExtension = new(RRCReconfigurationComplete_v1800_IEs)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
