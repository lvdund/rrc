// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SSB-MTC2 (line 15863).

var sSBMTC2Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "pci-List", Optional: true},
		{Name: "periodicity"},
	},
}

var sSBMTC2PciListConstraints = per.SizeRange(1, common.MaxNrofPCIsPerSMTC)

const (
	SSB_MTC2_Periodicity_Sf5    = 0
	SSB_MTC2_Periodicity_Sf10   = 1
	SSB_MTC2_Periodicity_Sf20   = 2
	SSB_MTC2_Periodicity_Sf40   = 3
	SSB_MTC2_Periodicity_Sf80   = 4
	SSB_MTC2_Periodicity_Spare3 = 5
	SSB_MTC2_Periodicity_Spare2 = 6
	SSB_MTC2_Periodicity_Spare1 = 7
)

var sSBMTC2PeriodicityConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SSB_MTC2_Periodicity_Sf5, SSB_MTC2_Periodicity_Sf10, SSB_MTC2_Periodicity_Sf20, SSB_MTC2_Periodicity_Sf40, SSB_MTC2_Periodicity_Sf80, SSB_MTC2_Periodicity_Spare3, SSB_MTC2_Periodicity_Spare2, SSB_MTC2_Periodicity_Spare1},
}

type SSB_MTC2 struct {
	Pci_List    []PhysCellId
	Periodicity int64
}

func (ie *SSB_MTC2) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sSBMTC2Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Pci_List != nil}); err != nil {
		return err
	}

	// 2. pci-List: sequence-of
	{
		if ie.Pci_List != nil {
			seqOf := e.NewSequenceOfEncoder(sSBMTC2PciListConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.Pci_List))); err != nil {
				return err
			}
			for i := range ie.Pci_List {
				if err := ie.Pci_List[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 3. periodicity: enumerated
	{
		if err := e.EncodeEnumerated(ie.Periodicity, sSBMTC2PeriodicityConstraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SSB_MTC2) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sSBMTC2Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. pci-List: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(sSBMTC2PciListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Pci_List = make([]PhysCellId, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Pci_List[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 3. periodicity: enumerated
	{
		v1, err := d.DecodeEnumerated(sSBMTC2PeriodicityConstraints)
		if err != nil {
			return err
		}
		ie.Periodicity = v1
	}

	return nil
}
