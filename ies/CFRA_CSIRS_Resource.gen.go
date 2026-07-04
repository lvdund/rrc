// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CFRA-CSIRS-Resource (line 12983).

var cFRACSIRSResourceConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "csi-RS"},
		{Name: "ra-OccasionList"},
		{Name: "ra-PreambleIndex"},
	},
}

var cFRACSIRSResourceRaOccasionListConstraints = per.SizeRange(1, common.MaxRA_OccasionsPerCSIRS)

var cFRACSIRSResourceRaPreambleIndexConstraints = per.Constrained(0, 63)

type CFRA_CSIRS_Resource struct {
	Csi_RS           CSI_RS_Index
	Ra_OccasionList  []int64
	Ra_PreambleIndex int64
}

func (ie *CFRA_CSIRS_Resource) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cFRACSIRSResourceConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. csi-RS: ref
	{
		if err := ie.Csi_RS.Encode(e); err != nil {
			return err
		}
	}

	// 3. ra-OccasionList: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(cFRACSIRSResourceRaOccasionListConstraints)
		if err := seqOf.EncodeLength(int64(len(ie.Ra_OccasionList))); err != nil {
			return err
		}
		for i := range ie.Ra_OccasionList {
			if err := e.EncodeInteger(ie.Ra_OccasionList[i], per.Constrained(0, common.MaxRA_Occasions_1)); err != nil {
				return err
			}
		}
	}

	// 4. ra-PreambleIndex: integer
	{
		if err := e.EncodeInteger(ie.Ra_PreambleIndex, cFRACSIRSResourceRaPreambleIndexConstraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *CFRA_CSIRS_Resource) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cFRACSIRSResourceConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. csi-RS: ref
	{
		if err := ie.Csi_RS.Decode(d); err != nil {
			return err
		}
	}

	// 3. ra-OccasionList: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(cFRACSIRSResourceRaOccasionListConstraints)
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

	// 4. ra-PreambleIndex: integer
	{
		v2, err := d.DecodeInteger(cFRACSIRSResourceRaPreambleIndexConstraints)
		if err != nil {
			return err
		}
		ie.Ra_PreambleIndex = v2
	}

	return nil
}
