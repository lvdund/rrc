// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CSI-LogMeasInfoCell-r19 (line 3493).

var cSILogMeasInfoCellR19Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "cellId-r19"},
		{Name: "csi-LogMeasInfoConfigList-r19"},
	},
}

var cSI_LogMeasInfoCell_r19CellIdR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "cellGlobalId-r19"},
		{Name: "pci-arfcn-r19"},
	},
}

const (
	CSI_LogMeasInfoCell_r19_CellId_r19_CellGlobalId_r19 = 0
	CSI_LogMeasInfoCell_r19_CellId_r19_Pci_Arfcn_r19    = 1
)

var cSILogMeasInfoCellR19CsiLogMeasInfoConfigListR19Constraints = per.SizeRange(1, common.MaxNrofLoggedMeasurementConfigurations_r19)

type CSI_LogMeasInfoCell_r19 struct {
	CellId_r19 struct {
		Choice           int
		CellGlobalId_r19 *CGI_Info_Logging_r16
		Pci_Arfcn_r19    *PCI_ARFCN_NR_r16
	}
	Csi_LogMeasInfoConfigList_r19 []CSI_LogMeasInfoConfig_r19
}

func (ie *CSI_LogMeasInfoCell_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cSILogMeasInfoCellR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. cellId-r19: choice
	{
		choiceEnc := e.NewChoiceEncoder(cSI_LogMeasInfoCell_r19CellIdR19Constraints)
		if err := choiceEnc.EncodeChoice(int64(ie.CellId_r19.Choice), false, nil); err != nil {
			return err
		}
		switch ie.CellId_r19.Choice {
		case CSI_LogMeasInfoCell_r19_CellId_r19_CellGlobalId_r19:
			if err := ie.CellId_r19.CellGlobalId_r19.Encode(e); err != nil {
				return err
			}
		case CSI_LogMeasInfoCell_r19_CellId_r19_Pci_Arfcn_r19:
			if err := ie.CellId_r19.Pci_Arfcn_r19.Encode(e); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.CellId_r19.Choice), Constraint: "index out of range"}
		}
	}

	// 3. csi-LogMeasInfoConfigList-r19: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(cSILogMeasInfoCellR19CsiLogMeasInfoConfigListR19Constraints)
		if err := seqOf.EncodeLength(int64(len(ie.Csi_LogMeasInfoConfigList_r19))); err != nil {
			return err
		}
		for i := range ie.Csi_LogMeasInfoConfigList_r19 {
			if err := ie.Csi_LogMeasInfoConfigList_r19[i].Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *CSI_LogMeasInfoCell_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cSILogMeasInfoCellR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. cellId-r19: choice
	{
		choiceDec := d.NewChoiceDecoder(cSI_LogMeasInfoCell_r19CellIdR19Constraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.CellId_r19.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case CSI_LogMeasInfoCell_r19_CellId_r19_CellGlobalId_r19:
			ie.CellId_r19.CellGlobalId_r19 = new(CGI_Info_Logging_r16)
			if err := ie.CellId_r19.CellGlobalId_r19.Decode(d); err != nil {
				return err
			}
		case CSI_LogMeasInfoCell_r19_CellId_r19_Pci_Arfcn_r19:
			ie.CellId_r19.Pci_Arfcn_r19 = new(PCI_ARFCN_NR_r16)
			if err := ie.CellId_r19.Pci_Arfcn_r19.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. csi-LogMeasInfoConfigList-r19: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(cSILogMeasInfoCellR19CsiLogMeasInfoConfigListR19Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.Csi_LogMeasInfoConfigList_r19 = make([]CSI_LogMeasInfoConfig_r19, n)
		for i := int64(0); i < n; i++ {
			if err := ie.Csi_LogMeasInfoConfigList_r19[i].Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
