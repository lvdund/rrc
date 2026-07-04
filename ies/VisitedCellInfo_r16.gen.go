// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: VisitedCellInfo-r16 (line 26699).

var visitedCellInfoR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "visitedCellId-r16", Optional: true},
		{Name: "timeSpent-r16"},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

var visitedCellInfo_r16VisitedCellIdR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "nr-CellId-r16"},
		{Name: "eutra-CellId-r16"},
	},
}

const (
	VisitedCellInfo_r16_VisitedCellId_r16_Nr_CellId_r16    = 0
	VisitedCellInfo_r16_VisitedCellId_r16_Eutra_CellId_r16 = 1
)

var visitedCellInfoR16TimeSpentR16Constraints = per.Constrained(0, 4095)

var visitedCellInfoR16VisitedCellIdR16NrCellIdR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "cgi-Info"},
		{Name: "pci-arfcn-r16"},
	},
}

const (
	VisitedCellInfo_r16_VisitedCellId_r16_Nr_CellId_r16_Cgi_Info      = 0
	VisitedCellInfo_r16_VisitedCellId_r16_Nr_CellId_r16_Pci_Arfcn_r16 = 1
)

var visitedCellInfoR16VisitedCellIdR16EutraCellIdR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "cellGlobalId-r16"},
		{Name: "pci-arfcn-r16"},
	},
}

const (
	VisitedCellInfo_r16_VisitedCellId_r16_Eutra_CellId_r16_CellGlobalId_r16 = 0
	VisitedCellInfo_r16_VisitedCellId_r16_Eutra_CellId_r16_Pci_Arfcn_r16    = 1
)

type VisitedCellInfo_r16 struct {
	VisitedCellId_r16 *struct {
		Choice        int
		Nr_CellId_r16 *struct {
			Choice        int
			Cgi_Info      *CGI_Info_Logging_r16
			Pci_Arfcn_r16 *PCI_ARFCN_NR_r16
		}
		Eutra_CellId_r16 *struct {
			Choice           int
			CellGlobalId_r16 *CGI_InfoEUTRA
			Pci_Arfcn_r16    *PCI_ARFCN_EUTRA_r16
		}
	}
	TimeSpent_r16                   int64
	VisitedPSCellInfoListReport_r17 *VisitedPSCellInfoList_r17
}

func (ie *VisitedCellInfo_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(visitedCellInfoR16Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.VisitedPSCellInfoListReport_r17 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.VisitedCellId_r16 != nil}); err != nil {
		return err
	}

	// 3. visitedCellId-r16: choice
	{
		if ie.VisitedCellId_r16 != nil {
			choiceEnc := e.NewChoiceEncoder(visitedCellInfo_r16VisitedCellIdR16Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.VisitedCellId_r16).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.VisitedCellId_r16).Choice {
			case VisitedCellInfo_r16_VisitedCellId_r16_Nr_CellId_r16:
				choiceEnc := e.NewChoiceEncoder(visitedCellInfoR16VisitedCellIdR16NrCellIdR16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*(*ie.VisitedCellId_r16).Nr_CellId_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*(*ie.VisitedCellId_r16).Nr_CellId_r16).Choice {
				case VisitedCellInfo_r16_VisitedCellId_r16_Nr_CellId_r16_Cgi_Info:
					if err := (*(*ie.VisitedCellId_r16).Nr_CellId_r16).Cgi_Info.Encode(e); err != nil {
						return err
					}
				case VisitedCellInfo_r16_VisitedCellId_r16_Nr_CellId_r16_Pci_Arfcn_r16:
					if err := (*(*ie.VisitedCellId_r16).Nr_CellId_r16).Pci_Arfcn_r16.Encode(e); err != nil {
						return err
					}
				}
			case VisitedCellInfo_r16_VisitedCellId_r16_Eutra_CellId_r16:
				choiceEnc := e.NewChoiceEncoder(visitedCellInfoR16VisitedCellIdR16EutraCellIdR16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*(*ie.VisitedCellId_r16).Eutra_CellId_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*(*ie.VisitedCellId_r16).Eutra_CellId_r16).Choice {
				case VisitedCellInfo_r16_VisitedCellId_r16_Eutra_CellId_r16_CellGlobalId_r16:
					if err := (*(*ie.VisitedCellId_r16).Eutra_CellId_r16).CellGlobalId_r16.Encode(e); err != nil {
						return err
					}
				case VisitedCellInfo_r16_VisitedCellId_r16_Eutra_CellId_r16_Pci_Arfcn_r16:
					if err := (*(*ie.VisitedCellId_r16).Eutra_CellId_r16).Pci_Arfcn_r16.Encode(e); err != nil {
						return err
					}
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.VisitedCellId_r16).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 4. timeSpent-r16: integer
	{
		if err := e.EncodeInteger(ie.TimeSpent_r16, visitedCellInfoR16TimeSpentR16Constraints); err != nil {
			return err
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "visitedPSCellInfoListReport-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.VisitedPSCellInfoListReport_r17 != nil}); err != nil {
				return err
			}

			if ie.VisitedPSCellInfoListReport_r17 != nil {
				if err := ie.VisitedPSCellInfoListReport_r17.Encode(ex); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		if err := seq.EncodeExtensionAdditions(extPresence, extValues); err != nil {
			return err
		}
	}

	return nil
}

func (ie *VisitedCellInfo_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(visitedCellInfoR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. visitedCellId-r16: choice
	{
		if seq.IsComponentPresent(0) {
			ie.VisitedCellId_r16 = &struct {
				Choice        int
				Nr_CellId_r16 *struct {
					Choice        int
					Cgi_Info      *CGI_Info_Logging_r16
					Pci_Arfcn_r16 *PCI_ARFCN_NR_r16
				}
				Eutra_CellId_r16 *struct {
					Choice           int
					CellGlobalId_r16 *CGI_InfoEUTRA
					Pci_Arfcn_r16    *PCI_ARFCN_EUTRA_r16
				}
			}{}
			choiceDec := d.NewChoiceDecoder(visitedCellInfo_r16VisitedCellIdR16Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.VisitedCellId_r16).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case VisitedCellInfo_r16_VisitedCellId_r16_Nr_CellId_r16:
				(*ie.VisitedCellId_r16).Nr_CellId_r16 = &struct {
					Choice        int
					Cgi_Info      *CGI_Info_Logging_r16
					Pci_Arfcn_r16 *PCI_ARFCN_NR_r16
				}{}
				choiceDec := d.NewChoiceDecoder(visitedCellInfoR16VisitedCellIdR16NrCellIdR16Constraints)
				index, _, _, err := choiceDec.DecodeChoice()
				if err != nil {
					return err
				}
				(*(*ie.VisitedCellId_r16).Nr_CellId_r16).Choice = int(index)
				switch index {
				case VisitedCellInfo_r16_VisitedCellId_r16_Nr_CellId_r16_Cgi_Info:
					(*(*ie.VisitedCellId_r16).Nr_CellId_r16).Cgi_Info = new(CGI_Info_Logging_r16)
					if err := (*(*ie.VisitedCellId_r16).Nr_CellId_r16).Cgi_Info.Decode(d); err != nil {
						return err
					}
				case VisitedCellInfo_r16_VisitedCellId_r16_Nr_CellId_r16_Pci_Arfcn_r16:
					(*(*ie.VisitedCellId_r16).Nr_CellId_r16).Pci_Arfcn_r16 = new(PCI_ARFCN_NR_r16)
					if err := (*(*ie.VisitedCellId_r16).Nr_CellId_r16).Pci_Arfcn_r16.Decode(d); err != nil {
						return err
					}
				}
			case VisitedCellInfo_r16_VisitedCellId_r16_Eutra_CellId_r16:
				(*ie.VisitedCellId_r16).Eutra_CellId_r16 = &struct {
					Choice           int
					CellGlobalId_r16 *CGI_InfoEUTRA
					Pci_Arfcn_r16    *PCI_ARFCN_EUTRA_r16
				}{}
				choiceDec := d.NewChoiceDecoder(visitedCellInfoR16VisitedCellIdR16EutraCellIdR16Constraints)
				index, _, _, err := choiceDec.DecodeChoice()
				if err != nil {
					return err
				}
				(*(*ie.VisitedCellId_r16).Eutra_CellId_r16).Choice = int(index)
				switch index {
				case VisitedCellInfo_r16_VisitedCellId_r16_Eutra_CellId_r16_CellGlobalId_r16:
					(*(*ie.VisitedCellId_r16).Eutra_CellId_r16).CellGlobalId_r16 = new(CGI_InfoEUTRA)
					if err := (*(*ie.VisitedCellId_r16).Eutra_CellId_r16).CellGlobalId_r16.Decode(d); err != nil {
						return err
					}
				case VisitedCellInfo_r16_VisitedCellId_r16_Eutra_CellId_r16_Pci_Arfcn_r16:
					(*(*ie.VisitedCellId_r16).Eutra_CellId_r16).Pci_Arfcn_r16 = new(PCI_ARFCN_EUTRA_r16)
					if err := (*(*ie.VisitedCellId_r16).Eutra_CellId_r16).Pci_Arfcn_r16.Decode(d); err != nil {
						return err
					}
				}
			}
		}
	}

	// 4. timeSpent-r16: integer
	{
		v1, err := d.DecodeInteger(visitedCellInfoR16TimeSpentR16Constraints)
		if err != nil {
			return err
		}
		ie.TimeSpent_r16 = v1
	}

	// Extension additions as open types.
	extValues, err := seq.DecodeExtensionAdditions()
	if err != nil {
		return err
	}
	extValueIndex := 0

	// Extension group 0:
	if seq.IsExtensionPresent(0) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "visitedPSCellInfoListReport-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.VisitedPSCellInfoListReport_r17 = new(VisitedPSCellInfoList_r17)
			if err := ie.VisitedPSCellInfoListReport_r17.Decode(dx); err != nil {
				return err
			}
		}
	}

	return nil
}
