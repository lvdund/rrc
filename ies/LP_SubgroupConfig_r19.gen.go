// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: LP-SubgroupConfig-r19 (line 8041).

var lPSubgroupConfigR19Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "lp-SubgroupsNumPerPO-r19"},
		{Name: "lp-SubgroupsNumForUEID-r19", Optional: true},
	},
}

var lPSubgroupConfigR19LpSubgroupsNumPerPOR19Constraints = per.Constrained(1, common.MaxNrofPagingSubgroupsLP_r19)

var lPSubgroupConfigR19LpSubgroupsNumForUEIDR19Constraints = per.Constrained(1, common.MaxNrofPagingSubgroupsLP_r19)

type LP_SubgroupConfig_r19 struct {
	Lp_SubgroupsNumPerPO_r19   int64
	Lp_SubgroupsNumForUEID_r19 *int64
}

func (ie *LP_SubgroupConfig_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(lPSubgroupConfigR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Lp_SubgroupsNumForUEID_r19 != nil}); err != nil {
		return err
	}

	// 3. lp-SubgroupsNumPerPO-r19: integer
	{
		if err := e.EncodeInteger(ie.Lp_SubgroupsNumPerPO_r19, lPSubgroupConfigR19LpSubgroupsNumPerPOR19Constraints); err != nil {
			return err
		}
	}

	// 4. lp-SubgroupsNumForUEID-r19: integer
	{
		if ie.Lp_SubgroupsNumForUEID_r19 != nil {
			if err := e.EncodeInteger(*ie.Lp_SubgroupsNumForUEID_r19, lPSubgroupConfigR19LpSubgroupsNumForUEIDR19Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *LP_SubgroupConfig_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(lPSubgroupConfigR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. lp-SubgroupsNumPerPO-r19: integer
	{
		v0, err := d.DecodeInteger(lPSubgroupConfigR19LpSubgroupsNumPerPOR19Constraints)
		if err != nil {
			return err
		}
		ie.Lp_SubgroupsNumPerPO_r19 = v0
	}

	// 4. lp-SubgroupsNumForUEID-r19: integer
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeInteger(lPSubgroupConfigR19LpSubgroupsNumForUEIDR19Constraints)
			if err != nil {
				return err
			}
			ie.Lp_SubgroupsNumForUEID_r19 = &v
		}
	}

	return nil
}
