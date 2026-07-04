// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RRCReconfigurationComplete-v1700-IEs (line 1173).

var rRCReconfigurationCompleteV1700IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "needForGapNCSG-InfoNR-r17", Optional: true},
		{Name: "needForGapNCSG-InfoEUTRA-r17", Optional: true},
		{Name: "selectedCondRRCReconfig-r17", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

type RRCReconfigurationComplete_v1700_IEs struct {
	NeedForGapNCSG_InfoNR_r17    *NeedForGapNCSG_InfoNR_r17
	NeedForGapNCSG_InfoEUTRA_r17 *NeedForGapNCSG_InfoEUTRA_r17
	SelectedCondRRCReconfig_r17  *CondReconfigId_r16
	NonCriticalExtension         *RRCReconfigurationComplete_v1720_IEs
}

func (ie *RRCReconfigurationComplete_v1700_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rRCReconfigurationCompleteV1700IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.NeedForGapNCSG_InfoNR_r17 != nil, ie.NeedForGapNCSG_InfoEUTRA_r17 != nil, ie.SelectedCondRRCReconfig_r17 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. needForGapNCSG-InfoNR-r17: ref
	{
		if ie.NeedForGapNCSG_InfoNR_r17 != nil {
			if err := ie.NeedForGapNCSG_InfoNR_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. needForGapNCSG-InfoEUTRA-r17: ref
	{
		if ie.NeedForGapNCSG_InfoEUTRA_r17 != nil {
			if err := ie.NeedForGapNCSG_InfoEUTRA_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. selectedCondRRCReconfig-r17: ref
	{
		if ie.SelectedCondRRCReconfig_r17 != nil {
			if err := ie.SelectedCondRRCReconfig_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. nonCriticalExtension: ref
	{
		if ie.NonCriticalExtension != nil {
			if err := ie.NonCriticalExtension.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *RRCReconfigurationComplete_v1700_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rRCReconfigurationCompleteV1700IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. needForGapNCSG-InfoNR-r17: ref
	{
		if seq.IsComponentPresent(0) {
			ie.NeedForGapNCSG_InfoNR_r17 = new(NeedForGapNCSG_InfoNR_r17)
			if err := ie.NeedForGapNCSG_InfoNR_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. needForGapNCSG-InfoEUTRA-r17: ref
	{
		if seq.IsComponentPresent(1) {
			ie.NeedForGapNCSG_InfoEUTRA_r17 = new(NeedForGapNCSG_InfoEUTRA_r17)
			if err := ie.NeedForGapNCSG_InfoEUTRA_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. selectedCondRRCReconfig-r17: ref
	{
		if seq.IsComponentPresent(2) {
			ie.SelectedCondRRCReconfig_r17 = new(CondReconfigId_r16)
			if err := ie.SelectedCondRRCReconfig_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(3) {
			ie.NonCriticalExtension = new(RRCReconfigurationComplete_v1720_IEs)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
