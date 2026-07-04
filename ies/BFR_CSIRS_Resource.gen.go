// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: BFR-CSIRS-Resource (line 5146).

var bFRCSIRSResourceConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "csi-RS"},
		{Name: "ra-OccasionList", Optional: true},
		{Name: "ra-PreambleIndex", Optional: true},
	},
}

var bFRCSIRSResourceRaOccasionListConstraints = per.SizeRange(1, common.MaxRA_OccasionsPerCSIRS)

var bFRCSIRSResourceRaPreambleIndexConstraints = per.Constrained(0, 63)

type BFR_CSIRS_Resource struct {
	Csi_RS           NZP_CSI_RS_ResourceId
	Ra_OccasionList  []int64
	Ra_PreambleIndex *int64
}

func (ie *BFR_CSIRS_Resource) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bFRCSIRSResourceConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ra_OccasionList != nil, ie.Ra_PreambleIndex != nil}); err != nil {
		return err
	}

	// 3. csi-RS: ref
	{
		if err := ie.Csi_RS.Encode(e); err != nil {
			return err
		}
	}

	// 4. ra-OccasionList: sequence-of
	{
		if ie.Ra_OccasionList != nil {
			seqOf := e.NewSequenceOfEncoder(bFRCSIRSResourceRaOccasionListConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.Ra_OccasionList))); err != nil {
				return err
			}
			for i := range ie.Ra_OccasionList {
				if err := e.EncodeInteger(ie.Ra_OccasionList[i], per.Constrained(0, common.MaxRA_Occasions_1)); err != nil {
					return err
				}
			}
		}
	}

	// 5. ra-PreambleIndex: integer
	{
		if ie.Ra_PreambleIndex != nil {
			if err := e.EncodeInteger(*ie.Ra_PreambleIndex, bFRCSIRSResourceRaPreambleIndexConstraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *BFR_CSIRS_Resource) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bFRCSIRSResourceConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. csi-RS: ref
	{
		if err := ie.Csi_RS.Decode(d); err != nil {
			return err
		}
	}

	// 4. ra-OccasionList: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(bFRCSIRSResourceRaOccasionListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Ra_OccasionList = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxRA_Occasions_1))
				if err != nil {
					return err
				}
				ie.Ra_OccasionList[i] = v
			}
		}
	}

	// 5. ra-PreambleIndex: integer
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeInteger(bFRCSIRSResourceRaPreambleIndexConstraints)
			if err != nil {
				return err
			}
			ie.Ra_PreambleIndex = &v
		}
	}

	return nil
}
