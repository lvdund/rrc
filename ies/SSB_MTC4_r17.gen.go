// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SSB-MTC4-r17 (line 15890).

var sSBMTC4R17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "pci-List-r17", Optional: true},
		{Name: "offset-r17"},
	},
}

var sSBMTC4R17PciListR17Constraints = per.SizeRange(1, common.MaxNrofPCIsPerSMTC)

var sSBMTC4R17OffsetR17Constraints = per.Constrained(0, 159)

type SSB_MTC4_r17 struct {
	Pci_List_r17 []PhysCellId
	Offset_r17   int64
}

func (ie *SSB_MTC4_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sSBMTC4R17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Pci_List_r17 != nil}); err != nil {
		return err
	}

	// 2. pci-List-r17: sequence-of
	{
		if ie.Pci_List_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(sSBMTC4R17PciListR17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Pci_List_r17))); err != nil {
				return err
			}
			for i := range ie.Pci_List_r17 {
				if err := ie.Pci_List_r17[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 3. offset-r17: integer
	{
		if err := e.EncodeInteger(ie.Offset_r17, sSBMTC4R17OffsetR17Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SSB_MTC4_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sSBMTC4R17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. pci-List-r17: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(sSBMTC4R17PciListR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Pci_List_r17 = make([]PhysCellId, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Pci_List_r17[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 3. offset-r17: integer
	{
		v1, err := d.DecodeInteger(sSBMTC4R17OffsetR17Constraints)
		if err != nil {
			return err
		}
		ie.Offset_r17 = v1
	}

	return nil
}
