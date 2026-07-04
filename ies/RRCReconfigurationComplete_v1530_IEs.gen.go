// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RRCReconfigurationComplete-v1530-IEs (line 1149).

var rRCReconfigurationCompleteV1530IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "uplinkTxDirectCurrentList", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

type RRCReconfigurationComplete_v1530_IEs struct {
	UplinkTxDirectCurrentList *UplinkTxDirectCurrentList
	NonCriticalExtension      *RRCReconfigurationComplete_v1560_IEs
}

func (ie *RRCReconfigurationComplete_v1530_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rRCReconfigurationCompleteV1530IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.UplinkTxDirectCurrentList != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. uplinkTxDirectCurrentList: ref
	{
		if ie.UplinkTxDirectCurrentList != nil {
			if err := ie.UplinkTxDirectCurrentList.Encode(e); err != nil {
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

func (ie *RRCReconfigurationComplete_v1530_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rRCReconfigurationCompleteV1530IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. uplinkTxDirectCurrentList: ref
	{
		if seq.IsComponentPresent(0) {
			ie.UplinkTxDirectCurrentList = new(UplinkTxDirectCurrentList)
			if err := ie.UplinkTxDirectCurrentList.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(1) {
			ie.NonCriticalExtension = new(RRCReconfigurationComplete_v1560_IEs)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
