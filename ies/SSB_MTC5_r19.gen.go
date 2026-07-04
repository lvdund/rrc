// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SSB-MTC5-r19 (line 15895).

var sSBMTC5R19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "pci-List-r19", Optional: true},
		{Name: "periodicity-r19", Optional: true},
		{Name: "offset-r19", Optional: true},
	},
}

var sSBMTC5R19PciListR19Constraints = per.SizeRange(1, common.MaxNrofPCIsPerSMTC)

const (
	SSB_MTC5_r19_Periodicity_r19_Sf10   = 0
	SSB_MTC5_r19_Periodicity_r19_Sf20   = 1
	SSB_MTC5_r19_Periodicity_r19_Sf40   = 2
	SSB_MTC5_r19_Periodicity_r19_Sf80   = 3
	SSB_MTC5_r19_Periodicity_r19_Sf160  = 4
	SSB_MTC5_r19_Periodicity_r19_Spare3 = 5
	SSB_MTC5_r19_Periodicity_r19_Spare2 = 6
	SSB_MTC5_r19_Periodicity_r19_Spare1 = 7
)

var sSBMTC5R19PeriodicityR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SSB_MTC5_r19_Periodicity_r19_Sf10, SSB_MTC5_r19_Periodicity_r19_Sf20, SSB_MTC5_r19_Periodicity_r19_Sf40, SSB_MTC5_r19_Periodicity_r19_Sf80, SSB_MTC5_r19_Periodicity_r19_Sf160, SSB_MTC5_r19_Periodicity_r19_Spare3, SSB_MTC5_r19_Periodicity_r19_Spare2, SSB_MTC5_r19_Periodicity_r19_Spare1},
}

var sSBMTC5R19OffsetR19Constraints = per.Constrained(0, 159)

type SSB_MTC5_r19 struct {
	Pci_List_r19    []PhysCellId
	Periodicity_r19 *int64
	Offset_r19      *int64
}

func (ie *SSB_MTC5_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sSBMTC5R19Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Pci_List_r19 != nil, ie.Periodicity_r19 != nil, ie.Offset_r19 != nil}); err != nil {
		return err
	}

	// 2. pci-List-r19: sequence-of
	{
		if ie.Pci_List_r19 != nil {
			seqOf := e.NewSequenceOfEncoder(sSBMTC5R19PciListR19Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Pci_List_r19))); err != nil {
				return err
			}
			for i := range ie.Pci_List_r19 {
				if err := ie.Pci_List_r19[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 3. periodicity-r19: enumerated
	{
		if ie.Periodicity_r19 != nil {
			if err := e.EncodeEnumerated(*ie.Periodicity_r19, sSBMTC5R19PeriodicityR19Constraints); err != nil {
				return err
			}
		}
	}

	// 4. offset-r19: integer
	{
		if ie.Offset_r19 != nil {
			if err := e.EncodeInteger(*ie.Offset_r19, sSBMTC5R19OffsetR19Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SSB_MTC5_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sSBMTC5R19Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. pci-List-r19: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(sSBMTC5R19PciListR19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Pci_List_r19 = make([]PhysCellId, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Pci_List_r19[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 3. periodicity-r19: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(sSBMTC5R19PeriodicityR19Constraints)
			if err != nil {
				return err
			}
			ie.Periodicity_r19 = &idx
		}
	}

	// 4. offset-r19: integer
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeInteger(sSBMTC5R19OffsetR19Constraints)
			if err != nil {
				return err
			}
			ie.Offset_r19 = &v
		}
	}

	return nil
}
