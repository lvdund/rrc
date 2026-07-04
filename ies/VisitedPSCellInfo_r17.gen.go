// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: VisitedPSCellInfo-r17 (line 26719).

var visitedPSCellInfoR17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "visitedCellId-r17", Optional: true},
		{Name: "timeSpent-r17"},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

var visitedPSCellInfo_r17VisitedCellIdR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "nr-CellId-r17"},
		{Name: "eutra-CellId-r17"},
	},
}

const (
	VisitedPSCellInfo_r17_VisitedCellId_r17_Nr_CellId_r17    = 0
	VisitedPSCellInfo_r17_VisitedCellId_r17_Eutra_CellId_r17 = 1
)

var visitedPSCellInfoR17TimeSpentR17Constraints = per.Constrained(0, 4095)

var visitedPSCellInfoR17ScgActiveDurationR19Constraints = per.Constrained(0, 4095)

var visitedPSCellInfoR17VisitedCellIdR17NrCellIdR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "cgi-Info-r17"},
		{Name: "pci-arfcn-r17"},
	},
}

const (
	VisitedPSCellInfo_r17_VisitedCellId_r17_Nr_CellId_r17_Cgi_Info_r17  = 0
	VisitedPSCellInfo_r17_VisitedCellId_r17_Nr_CellId_r17_Pci_Arfcn_r17 = 1
)

var visitedPSCellInfoR17VisitedCellIdR17EutraCellIdR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "cellGlobalId-r17"},
		{Name: "pci-arfcn-r17"},
	},
}

const (
	VisitedPSCellInfo_r17_VisitedCellId_r17_Eutra_CellId_r17_CellGlobalId_r17 = 0
	VisitedPSCellInfo_r17_VisitedCellId_r17_Eutra_CellId_r17_Pci_Arfcn_r17    = 1
)

type VisitedPSCellInfo_r17 struct {
	VisitedCellId_r17 *struct {
		Choice        int
		Nr_CellId_r17 *struct {
			Choice        int
			Cgi_Info_r17  *CGI_Info_Logging_r16
			Pci_Arfcn_r17 *PCI_ARFCN_NR_r16
		}
		Eutra_CellId_r17 *struct {
			Choice           int
			CellGlobalId_r17 *CGI_InfoEUTRALogging
			Pci_Arfcn_r17    *PCI_ARFCN_EUTRA_r16
		}
	}
	TimeSpent_r17         int64
	ScgActiveDuration_r19 *int64
}

func (ie *VisitedPSCellInfo_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(visitedPSCellInfoR17Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.ScgActiveDuration_r19 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.VisitedCellId_r17 != nil}); err != nil {
		return err
	}

	// 3. visitedCellId-r17: choice
	{
		if ie.VisitedCellId_r17 != nil {
			choiceEnc := e.NewChoiceEncoder(visitedPSCellInfo_r17VisitedCellIdR17Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.VisitedCellId_r17).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.VisitedCellId_r17).Choice {
			case VisitedPSCellInfo_r17_VisitedCellId_r17_Nr_CellId_r17:
				choiceEnc := e.NewChoiceEncoder(visitedPSCellInfoR17VisitedCellIdR17NrCellIdR17Constraints)
				if err := choiceEnc.EncodeChoice(int64((*(*ie.VisitedCellId_r17).Nr_CellId_r17).Choice), false, nil); err != nil {
					return err
				}
				switch (*(*ie.VisitedCellId_r17).Nr_CellId_r17).Choice {
				case VisitedPSCellInfo_r17_VisitedCellId_r17_Nr_CellId_r17_Cgi_Info_r17:
					if err := (*(*ie.VisitedCellId_r17).Nr_CellId_r17).Cgi_Info_r17.Encode(e); err != nil {
						return err
					}
				case VisitedPSCellInfo_r17_VisitedCellId_r17_Nr_CellId_r17_Pci_Arfcn_r17:
					if err := (*(*ie.VisitedCellId_r17).Nr_CellId_r17).Pci_Arfcn_r17.Encode(e); err != nil {
						return err
					}
				}
			case VisitedPSCellInfo_r17_VisitedCellId_r17_Eutra_CellId_r17:
				choiceEnc := e.NewChoiceEncoder(visitedPSCellInfoR17VisitedCellIdR17EutraCellIdR17Constraints)
				if err := choiceEnc.EncodeChoice(int64((*(*ie.VisitedCellId_r17).Eutra_CellId_r17).Choice), false, nil); err != nil {
					return err
				}
				switch (*(*ie.VisitedCellId_r17).Eutra_CellId_r17).Choice {
				case VisitedPSCellInfo_r17_VisitedCellId_r17_Eutra_CellId_r17_CellGlobalId_r17:
					if err := (*(*ie.VisitedCellId_r17).Eutra_CellId_r17).CellGlobalId_r17.Encode(e); err != nil {
						return err
					}
				case VisitedPSCellInfo_r17_VisitedCellId_r17_Eutra_CellId_r17_Pci_Arfcn_r17:
					if err := (*(*ie.VisitedCellId_r17).Eutra_CellId_r17).Pci_Arfcn_r17.Encode(e); err != nil {
						return err
					}
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.VisitedCellId_r17).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 4. timeSpent-r17: integer
	{
		if err := e.EncodeInteger(ie.TimeSpent_r17, visitedPSCellInfoR17TimeSpentR17Constraints); err != nil {
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
					{Name: "scgActiveDuration-r19", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.ScgActiveDuration_r19 != nil}); err != nil {
				return err
			}

			if ie.ScgActiveDuration_r19 != nil {
				if err := ex.EncodeInteger(*ie.ScgActiveDuration_r19, visitedPSCellInfoR17ScgActiveDurationR19Constraints); err != nil {
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

func (ie *VisitedPSCellInfo_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(visitedPSCellInfoR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. visitedCellId-r17: choice
	{
		if seq.IsComponentPresent(0) {
			ie.VisitedCellId_r17 = &struct {
				Choice        int
				Nr_CellId_r17 *struct {
					Choice        int
					Cgi_Info_r17  *CGI_Info_Logging_r16
					Pci_Arfcn_r17 *PCI_ARFCN_NR_r16
				}
				Eutra_CellId_r17 *struct {
					Choice           int
					CellGlobalId_r17 *CGI_InfoEUTRALogging
					Pci_Arfcn_r17    *PCI_ARFCN_EUTRA_r16
				}
			}{}
			choiceDec := d.NewChoiceDecoder(visitedPSCellInfo_r17VisitedCellIdR17Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.VisitedCellId_r17).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case VisitedPSCellInfo_r17_VisitedCellId_r17_Nr_CellId_r17:
				(*ie.VisitedCellId_r17).Nr_CellId_r17 = &struct {
					Choice        int
					Cgi_Info_r17  *CGI_Info_Logging_r16
					Pci_Arfcn_r17 *PCI_ARFCN_NR_r16
				}{}
				choiceDec := d.NewChoiceDecoder(visitedPSCellInfoR17VisitedCellIdR17NrCellIdR17Constraints)
				index, _, _, err := choiceDec.DecodeChoice()
				if err != nil {
					return err
				}
				(*(*ie.VisitedCellId_r17).Nr_CellId_r17).Choice = int(index)
				switch index {
				case VisitedPSCellInfo_r17_VisitedCellId_r17_Nr_CellId_r17_Cgi_Info_r17:
					(*(*ie.VisitedCellId_r17).Nr_CellId_r17).Cgi_Info_r17 = new(CGI_Info_Logging_r16)
					if err := (*(*ie.VisitedCellId_r17).Nr_CellId_r17).Cgi_Info_r17.Decode(d); err != nil {
						return err
					}
				case VisitedPSCellInfo_r17_VisitedCellId_r17_Nr_CellId_r17_Pci_Arfcn_r17:
					(*(*ie.VisitedCellId_r17).Nr_CellId_r17).Pci_Arfcn_r17 = new(PCI_ARFCN_NR_r16)
					if err := (*(*ie.VisitedCellId_r17).Nr_CellId_r17).Pci_Arfcn_r17.Decode(d); err != nil {
						return err
					}
				}
			case VisitedPSCellInfo_r17_VisitedCellId_r17_Eutra_CellId_r17:
				(*ie.VisitedCellId_r17).Eutra_CellId_r17 = &struct {
					Choice           int
					CellGlobalId_r17 *CGI_InfoEUTRALogging
					Pci_Arfcn_r17    *PCI_ARFCN_EUTRA_r16
				}{}
				choiceDec := d.NewChoiceDecoder(visitedPSCellInfoR17VisitedCellIdR17EutraCellIdR17Constraints)
				index, _, _, err := choiceDec.DecodeChoice()
				if err != nil {
					return err
				}
				(*(*ie.VisitedCellId_r17).Eutra_CellId_r17).Choice = int(index)
				switch index {
				case VisitedPSCellInfo_r17_VisitedCellId_r17_Eutra_CellId_r17_CellGlobalId_r17:
					(*(*ie.VisitedCellId_r17).Eutra_CellId_r17).CellGlobalId_r17 = new(CGI_InfoEUTRALogging)
					if err := (*(*ie.VisitedCellId_r17).Eutra_CellId_r17).CellGlobalId_r17.Decode(d); err != nil {
						return err
					}
				case VisitedPSCellInfo_r17_VisitedCellId_r17_Eutra_CellId_r17_Pci_Arfcn_r17:
					(*(*ie.VisitedCellId_r17).Eutra_CellId_r17).Pci_Arfcn_r17 = new(PCI_ARFCN_EUTRA_r16)
					if err := (*(*ie.VisitedCellId_r17).Eutra_CellId_r17).Pci_Arfcn_r17.Decode(d); err != nil {
						return err
					}
				}
			}
		}
	}

	// 4. timeSpent-r17: integer
	{
		v1, err := d.DecodeInteger(visitedPSCellInfoR17TimeSpentR17Constraints)
		if err != nil {
			return err
		}
		ie.TimeSpent_r17 = v1
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
				{Name: "scgActiveDuration-r19", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeInteger(visitedPSCellInfoR17ScgActiveDurationR19Constraints)
			if err != nil {
				return err
			}
			ie.ScgActiveDuration_r19 = &v
		}
	}

	return nil
}
