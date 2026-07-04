// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MUSIM-BandEntryIndex-r18 (line 2669).
// MUSIM-BandEntryIndex-r18 ::=            INTEGER(1.. maxCandidateBandIndex-r18)

var mUSIMBandEntryIndexR18Constraints = per.Constrained(1, common.MaxCandidateBandIndex_r18)

type MUSIM_BandEntryIndex_r18 struct {
	Value int64
}

func (v *MUSIM_BandEntryIndex_r18) Encode(e *per.Encoder) error {
	return e.EncodeInteger(v.Value, mUSIMBandEntryIndexR18Constraints)
}

func (v *MUSIM_BandEntryIndex_r18) Decode(d *per.Decoder) error {
	x, err := d.DecodeInteger(mUSIMBandEntryIndexR18Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
