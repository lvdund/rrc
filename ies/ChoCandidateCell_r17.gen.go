// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: ChoCandidateCell-r17 (line 3465).

var choCandidateCellR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "cellGlobalId-r17"},
		{Name: "pci-arfcn-r17"},
	},
}

const (
	ChoCandidateCell_r17_CellGlobalId_r17 = 0
	ChoCandidateCell_r17_Pci_Arfcn_r17    = 1
)

type ChoCandidateCell_r17 struct {
	Choice           int
	CellGlobalId_r17 *CGI_Info_Logging_r16
	Pci_Arfcn_r17    *PCI_ARFCN_NR_r16
}

func (ie *ChoCandidateCell_r17) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(choCandidateCellR17Constraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case ChoCandidateCell_r17_CellGlobalId_r17:
		if err := ie.CellGlobalId_r17.Encode(e); err != nil {
			return err
		}
	case ChoCandidateCell_r17_Pci_Arfcn_r17:
		if err := ie.Pci_Arfcn_r17.Encode(e); err != nil {
			return err
		}
	default:
		// Choice index out of root range.
		return &per.ConstraintViolationError{
			TypeName:   "ChoCandidateCell-r17",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *ChoCandidateCell_r17) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(choCandidateCellR17Constraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "ChoCandidateCell-r17", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case ChoCandidateCell_r17_CellGlobalId_r17:
		ie.CellGlobalId_r17 = new(CGI_Info_Logging_r16)
		if err := ie.CellGlobalId_r17.Decode(d); err != nil {
			return err
		}
	case ChoCandidateCell_r17_Pci_Arfcn_r17:
		ie.Pci_Arfcn_r17 = new(PCI_ARFCN_NR_r16)
		if err := ie.Pci_Arfcn_r17.Decode(d); err != nil {
			return err
		}
	default:
		return &per.DecodeError{TypeName: "ChoCandidateCell-r17", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
