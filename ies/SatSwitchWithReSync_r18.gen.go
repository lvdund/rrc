// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SatSwitchWithReSync-r18 (line 4548).

var satSwitchWithReSyncR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ntn-Config-r18"},
		{Name: "t-ServiceStart-r18", Optional: true},
		{Name: "ssb-TimeOffset-r18", Optional: true},
	},
}

var satSwitchWithReSyncR18TServiceStartR18Constraints = per.Constrained(0, 549755813887)

var satSwitchWithReSyncR18SsbTimeOffsetR18Constraints = per.Constrained(0, 159)

type SatSwitchWithReSync_r18 struct {
	Ntn_Config_r18     NTN_Config_r17
	T_ServiceStart_r18 *int64
	Ssb_TimeOffset_r18 *int64
}

func (ie *SatSwitchWithReSync_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(satSwitchWithReSyncR18Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.T_ServiceStart_r18 != nil, ie.Ssb_TimeOffset_r18 != nil}); err != nil {
		return err
	}

	// 2. ntn-Config-r18: ref
	{
		if err := ie.Ntn_Config_r18.Encode(e); err != nil {
			return err
		}
	}

	// 3. t-ServiceStart-r18: integer
	{
		if ie.T_ServiceStart_r18 != nil {
			if err := e.EncodeInteger(*ie.T_ServiceStart_r18, satSwitchWithReSyncR18TServiceStartR18Constraints); err != nil {
				return err
			}
		}
	}

	// 4. ssb-TimeOffset-r18: integer
	{
		if ie.Ssb_TimeOffset_r18 != nil {
			if err := e.EncodeInteger(*ie.Ssb_TimeOffset_r18, satSwitchWithReSyncR18SsbTimeOffsetR18Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SatSwitchWithReSync_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(satSwitchWithReSyncR18Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. ntn-Config-r18: ref
	{
		if err := ie.Ntn_Config_r18.Decode(d); err != nil {
			return err
		}
	}

	// 3. t-ServiceStart-r18: integer
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeInteger(satSwitchWithReSyncR18TServiceStartR18Constraints)
			if err != nil {
				return err
			}
			ie.T_ServiceStart_r18 = &v
		}
	}

	// 4. ssb-TimeOffset-r18: integer
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeInteger(satSwitchWithReSyncR18SsbTimeOffsetR18Constraints)
			if err != nil {
				return err
			}
			ie.Ssb_TimeOffset_r18 = &v
		}
	}

	return nil
}
