// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CHO-WithCandidateSCGInfo-r19 (line 5969).

var cHOWithCandidateSCGInfoR19Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "firstFulfilledConfig-r19", Optional: true},
		{Name: "timeBetweenFulfillment-r19", Optional: true},
		{Name: "timeBetweenLastFulfillmentAndFailure-r19", Optional: true},
		{Name: "fulfilledConfigWhenChoOnly-r19", Optional: true},
		{Name: "pCellId-r19", Optional: true},
		{Name: "psCellId-r19", Optional: true},
	},
}

const (
	CHO_WithCandidateSCGInfo_r19_FirstFulfilledConfig_r19_Cho  = 0
	CHO_WithCandidateSCGInfo_r19_FirstFulfilledConfig_r19_Cpac = 1
)

var cHOWithCandidateSCGInfoR19FirstFulfilledConfigR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CHO_WithCandidateSCGInfo_r19_FirstFulfilledConfig_r19_Cho, CHO_WithCandidateSCGInfo_r19_FirstFulfilledConfig_r19_Cpac},
}

const (
	CHO_WithCandidateSCGInfo_r19_FulfilledConfigWhenChoOnly_r19_Cho     = 0
	CHO_WithCandidateSCGInfo_r19_FulfilledConfigWhenChoOnly_r19_Cpac    = 1
	CHO_WithCandidateSCGInfo_r19_FulfilledConfigWhenChoOnly_r19_Neither = 2
	CHO_WithCandidateSCGInfo_r19_FulfilledConfigWhenChoOnly_r19_Spare   = 3
)

var cHOWithCandidateSCGInfoR19FulfilledConfigWhenChoOnlyR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CHO_WithCandidateSCGInfo_r19_FulfilledConfigWhenChoOnly_r19_Cho, CHO_WithCandidateSCGInfo_r19_FulfilledConfigWhenChoOnly_r19_Cpac, CHO_WithCandidateSCGInfo_r19_FulfilledConfigWhenChoOnly_r19_Neither, CHO_WithCandidateSCGInfo_r19_FulfilledConfigWhenChoOnly_r19_Spare},
}

var cHO_WithCandidateSCGInfo_r19PCellIdR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "cellGlobalId-r19"},
		{Name: "pci-arfcn-r19"},
	},
}

const (
	CHO_WithCandidateSCGInfo_r19_PCellId_r19_CellGlobalId_r19 = 0
	CHO_WithCandidateSCGInfo_r19_PCellId_r19_Pci_Arfcn_r19    = 1
)

var cHO_WithCandidateSCGInfo_r19PsCellIdR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "cellGlobalId-r19"},
		{Name: "pci-arfcn-r19"},
	},
}

const (
	CHO_WithCandidateSCGInfo_r19_PsCellId_r19_CellGlobalId_r19 = 0
	CHO_WithCandidateSCGInfo_r19_PsCellId_r19_Pci_Arfcn_r19    = 1
)

type CHO_WithCandidateSCGInfo_r19 struct {
	FirstFulfilledConfig_r19                 *int64
	TimeBetweenFulfillment_r19               *TimeBetweenEvent_r17
	TimeBetweenLastFulfillmentAndFailure_r19 *TimeBetweenEvent_r17
	FulfilledConfigWhenChoOnly_r19           *int64
	PCellId_r19                              *struct {
		Choice           int
		CellGlobalId_r19 *CGI_Info_Logging_r16
		Pci_Arfcn_r19    *PCI_ARFCN_NR_r16
	}
	PsCellId_r19 *struct {
		Choice           int
		CellGlobalId_r19 *CGI_Info_Logging_r16
		Pci_Arfcn_r19    *PCI_ARFCN_NR_r16
	}
}

func (ie *CHO_WithCandidateSCGInfo_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cHOWithCandidateSCGInfoR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.FirstFulfilledConfig_r19 != nil, ie.TimeBetweenFulfillment_r19 != nil, ie.TimeBetweenLastFulfillmentAndFailure_r19 != nil, ie.FulfilledConfigWhenChoOnly_r19 != nil, ie.PCellId_r19 != nil, ie.PsCellId_r19 != nil}); err != nil {
		return err
	}

	// 3. firstFulfilledConfig-r19: enumerated
	{
		if ie.FirstFulfilledConfig_r19 != nil {
			if err := e.EncodeEnumerated(*ie.FirstFulfilledConfig_r19, cHOWithCandidateSCGInfoR19FirstFulfilledConfigR19Constraints); err != nil {
				return err
			}
		}
	}

	// 4. timeBetweenFulfillment-r19: ref
	{
		if ie.TimeBetweenFulfillment_r19 != nil {
			if err := ie.TimeBetweenFulfillment_r19.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. timeBetweenLastFulfillmentAndFailure-r19: ref
	{
		if ie.TimeBetweenLastFulfillmentAndFailure_r19 != nil {
			if err := ie.TimeBetweenLastFulfillmentAndFailure_r19.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. fulfilledConfigWhenChoOnly-r19: enumerated
	{
		if ie.FulfilledConfigWhenChoOnly_r19 != nil {
			if err := e.EncodeEnumerated(*ie.FulfilledConfigWhenChoOnly_r19, cHOWithCandidateSCGInfoR19FulfilledConfigWhenChoOnlyR19Constraints); err != nil {
				return err
			}
		}
	}

	// 7. pCellId-r19: choice
	{
		if ie.PCellId_r19 != nil {
			choiceEnc := e.NewChoiceEncoder(cHO_WithCandidateSCGInfo_r19PCellIdR19Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.PCellId_r19).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.PCellId_r19).Choice {
			case CHO_WithCandidateSCGInfo_r19_PCellId_r19_CellGlobalId_r19:
				if err := (*ie.PCellId_r19).CellGlobalId_r19.Encode(e); err != nil {
					return err
				}
			case CHO_WithCandidateSCGInfo_r19_PCellId_r19_Pci_Arfcn_r19:
				if err := (*ie.PCellId_r19).Pci_Arfcn_r19.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.PCellId_r19).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 8. psCellId-r19: choice
	{
		if ie.PsCellId_r19 != nil {
			choiceEnc := e.NewChoiceEncoder(cHO_WithCandidateSCGInfo_r19PsCellIdR19Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.PsCellId_r19).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.PsCellId_r19).Choice {
			case CHO_WithCandidateSCGInfo_r19_PsCellId_r19_CellGlobalId_r19:
				if err := (*ie.PsCellId_r19).CellGlobalId_r19.Encode(e); err != nil {
					return err
				}
			case CHO_WithCandidateSCGInfo_r19_PsCellId_r19_Pci_Arfcn_r19:
				if err := (*ie.PsCellId_r19).Pci_Arfcn_r19.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.PsCellId_r19).Choice), Constraint: "index out of range"}
			}
		}
	}

	return nil
}

func (ie *CHO_WithCandidateSCGInfo_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cHOWithCandidateSCGInfoR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. firstFulfilledConfig-r19: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(cHOWithCandidateSCGInfoR19FirstFulfilledConfigR19Constraints)
			if err != nil {
				return err
			}
			ie.FirstFulfilledConfig_r19 = &idx
		}
	}

	// 4. timeBetweenFulfillment-r19: ref
	{
		if seq.IsComponentPresent(1) {
			ie.TimeBetweenFulfillment_r19 = new(TimeBetweenEvent_r17)
			if err := ie.TimeBetweenFulfillment_r19.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. timeBetweenLastFulfillmentAndFailure-r19: ref
	{
		if seq.IsComponentPresent(2) {
			ie.TimeBetweenLastFulfillmentAndFailure_r19 = new(TimeBetweenEvent_r17)
			if err := ie.TimeBetweenLastFulfillmentAndFailure_r19.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. fulfilledConfigWhenChoOnly-r19: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(cHOWithCandidateSCGInfoR19FulfilledConfigWhenChoOnlyR19Constraints)
			if err != nil {
				return err
			}
			ie.FulfilledConfigWhenChoOnly_r19 = &idx
		}
	}

	// 7. pCellId-r19: choice
	{
		if seq.IsComponentPresent(4) {
			ie.PCellId_r19 = &struct {
				Choice           int
				CellGlobalId_r19 *CGI_Info_Logging_r16
				Pci_Arfcn_r19    *PCI_ARFCN_NR_r16
			}{}
			choiceDec := d.NewChoiceDecoder(cHO_WithCandidateSCGInfo_r19PCellIdR19Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.PCellId_r19).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case CHO_WithCandidateSCGInfo_r19_PCellId_r19_CellGlobalId_r19:
				(*ie.PCellId_r19).CellGlobalId_r19 = new(CGI_Info_Logging_r16)
				if err := (*ie.PCellId_r19).CellGlobalId_r19.Decode(d); err != nil {
					return err
				}
			case CHO_WithCandidateSCGInfo_r19_PCellId_r19_Pci_Arfcn_r19:
				(*ie.PCellId_r19).Pci_Arfcn_r19 = new(PCI_ARFCN_NR_r16)
				if err := (*ie.PCellId_r19).Pci_Arfcn_r19.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 8. psCellId-r19: choice
	{
		if seq.IsComponentPresent(5) {
			ie.PsCellId_r19 = &struct {
				Choice           int
				CellGlobalId_r19 *CGI_Info_Logging_r16
				Pci_Arfcn_r19    *PCI_ARFCN_NR_r16
			}{}
			choiceDec := d.NewChoiceDecoder(cHO_WithCandidateSCGInfo_r19PsCellIdR19Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.PsCellId_r19).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case CHO_WithCandidateSCGInfo_r19_PsCellId_r19_CellGlobalId_r19:
				(*ie.PsCellId_r19).CellGlobalId_r19 = new(CGI_Info_Logging_r16)
				if err := (*ie.PsCellId_r19).CellGlobalId_r19.Decode(d); err != nil {
					return err
				}
			case CHO_WithCandidateSCGInfo_r19_PsCellId_r19_Pci_Arfcn_r19:
				(*ie.PsCellId_r19).Pci_Arfcn_r19 = new(PCI_ARFCN_NR_r16)
				if err := (*ie.PsCellId_r19).Pci_Arfcn_r19.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
