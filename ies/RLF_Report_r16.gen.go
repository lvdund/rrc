// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RLF-Report-r16 (line 3215).

var rLFReportR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "nr-RLF-Report-r16"},
		{Name: "eutra-RLF-Report-r16"},
	},
}

const (
	RLF_Report_r16_Nr_RLF_Report_r16    = 0
	RLF_Report_r16_Eutra_RLF_Report_r16 = 1
)

var rLFReportR16NrRLFReportR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "measResultLastServCell-r16"},
		{Name: "measResultNeighCells-r16", Optional: true},
		{Name: "c-RNTI-r16"},
		{Name: "previousPCellId-r16", Optional: true},
		{Name: "failedPCellId-r16"},
		{Name: "reconnectCellId-r16", Optional: true},
		{Name: "timeUntilReconnection-r16", Optional: true},
		{Name: "reestablishmentCellId-r16", Optional: true},
		{Name: "timeConnFailure-r16", Optional: true},
		{Name: "timeSinceFailure-r16"},
		{Name: "connectionFailureType-r16"},
		{Name: "rlf-Cause-r16"},
		{Name: "locationInfo-r16", Optional: true},
		{Name: "noSuitableCellFound-r16", Optional: true},
		{Name: "ra-InformationCommon-r16", Optional: true},
	},
}

var rLFReportR16NrRLFReportR16MeasResultNeighCellsR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "measResultListNR-r16", Optional: true},
		{Name: "measResultListEUTRA-r16", Optional: true},
	},
}

var rLFReportR16NrRLFReportR16PreviousPCellIdR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "nrPreviousCell-r16"},
		{Name: "eutraPreviousCell-r16"},
	},
}

const (
	RLF_Report_r16_Nr_RLF_Report_r16_PreviousPCellId_r16_NrPreviousCell_r16    = 0
	RLF_Report_r16_Nr_RLF_Report_r16_PreviousPCellId_r16_EutraPreviousCell_r16 = 1
)

var rLFReportR16NrRLFReportR16FailedPCellIdR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "nrFailedPCellId-r16"},
		{Name: "eutraFailedPCellId-r16"},
	},
}

const (
	RLF_Report_r16_Nr_RLF_Report_r16_FailedPCellId_r16_NrFailedPCellId_r16    = 0
	RLF_Report_r16_Nr_RLF_Report_r16_FailedPCellId_r16_EutraFailedPCellId_r16 = 1
)

var rLFReportR16NrRLFReportR16FailedPCellIdR16NrFailedPCellIdR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "cellGlobalId-r16"},
		{Name: "pci-arfcn-r16"},
	},
}

const (
	RLF_Report_r16_Nr_RLF_Report_r16_FailedPCellId_r16_NrFailedPCellId_r16_CellGlobalId_r16 = 0
	RLF_Report_r16_Nr_RLF_Report_r16_FailedPCellId_r16_NrFailedPCellId_r16_Pci_Arfcn_r16    = 1
)

var rLFReportR16NrRLFReportR16FailedPCellIdR16EutraFailedPCellIdR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "cellGlobalId-r16"},
		{Name: "pci-arfcn-r16"},
	},
}

const (
	RLF_Report_r16_Nr_RLF_Report_r16_FailedPCellId_r16_EutraFailedPCellId_r16_CellGlobalId_r16 = 0
	RLF_Report_r16_Nr_RLF_Report_r16_FailedPCellId_r16_EutraFailedPCellId_r16_Pci_Arfcn_r16    = 1
)

var rLFReportR16NrRLFReportR16ReconnectCellIdR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "nrReconnectCellId-r16"},
		{Name: "eutraReconnectCellId-r16"},
	},
}

const (
	RLF_Report_r16_Nr_RLF_Report_r16_ReconnectCellId_r16_NrReconnectCellId_r16    = 0
	RLF_Report_r16_Nr_RLF_Report_r16_ReconnectCellId_r16_EutraReconnectCellId_r16 = 1
)

const (
	RLF_Report_r16_Nr_RLF_Report_r16_ConnectionFailureType_r16_Rlf = 0
	RLF_Report_r16_Nr_RLF_Report_r16_ConnectionFailureType_r16_Hof = 1
)

var rLFReportR16NrRLFReportR16ConnectionFailureTypeR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RLF_Report_r16_Nr_RLF_Report_r16_ConnectionFailureType_r16_Rlf, RLF_Report_r16_Nr_RLF_Report_r16_ConnectionFailureType_r16_Hof},
}

const (
	RLF_Report_r16_Nr_RLF_Report_r16_Rlf_Cause_r16_T310_Expiry                = 0
	RLF_Report_r16_Nr_RLF_Report_r16_Rlf_Cause_r16_RandomAccessProblem        = 1
	RLF_Report_r16_Nr_RLF_Report_r16_Rlf_Cause_r16_Rlc_MaxNumRetx             = 2
	RLF_Report_r16_Nr_RLF_Report_r16_Rlf_Cause_r16_BeamFailureRecoveryFailure = 3
	RLF_Report_r16_Nr_RLF_Report_r16_Rlf_Cause_r16_LbtFailure_r16             = 4
	RLF_Report_r16_Nr_RLF_Report_r16_Rlf_Cause_r16_Bh_RlfRecoveryFailure      = 5
	RLF_Report_r16_Nr_RLF_Report_r16_Rlf_Cause_r16_T312_Expiry_r17            = 6
	RLF_Report_r16_Nr_RLF_Report_r16_Rlf_Cause_r16_Spare1                     = 7
)

var rLFReportR16NrRLFReportR16RlfCauseR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RLF_Report_r16_Nr_RLF_Report_r16_Rlf_Cause_r16_T310_Expiry, RLF_Report_r16_Nr_RLF_Report_r16_Rlf_Cause_r16_RandomAccessProblem, RLF_Report_r16_Nr_RLF_Report_r16_Rlf_Cause_r16_Rlc_MaxNumRetx, RLF_Report_r16_Nr_RLF_Report_r16_Rlf_Cause_r16_BeamFailureRecoveryFailure, RLF_Report_r16_Nr_RLF_Report_r16_Rlf_Cause_r16_LbtFailure_r16, RLF_Report_r16_Nr_RLF_Report_r16_Rlf_Cause_r16_Bh_RlfRecoveryFailure, RLF_Report_r16_Nr_RLF_Report_r16_Rlf_Cause_r16_T312_Expiry_r17, RLF_Report_r16_Nr_RLF_Report_r16_Rlf_Cause_r16_Spare1},
}

const (
	RLF_Report_r16_Nr_RLF_Report_r16_NoSuitableCellFound_r16_True = 0
)

var rLFReportR16NrRLFReportR16NoSuitableCellFoundR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RLF_Report_r16_Nr_RLF_Report_r16_NoSuitableCellFound_r16_True},
}

var rLFReportR16EutraRLFReportR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "failedPCellId-EUTRA"},
		{Name: "measResult-RLF-Report-EUTRA-r16"},
	},
}

type RLF_Report_r16 struct {
	Choice            int
	Nr_RLF_Report_r16 *struct {
		MeasResultLastServCell_r16 MeasResultRLFNR_r16
		MeasResultNeighCells_r16   *struct {
			MeasResultListNR_r16    *MeasResultList2NR_r16
			MeasResultListEUTRA_r16 *MeasResultList2EUTRA_r16
		}
		C_RNTI_r16          RNTI_Value
		PreviousPCellId_r16 *struct {
			Choice                int
			NrPreviousCell_r16    *CGI_Info_Logging_r16
			EutraPreviousCell_r16 *CGI_InfoEUTRALogging
		}
		FailedPCellId_r16 struct {
			Choice              int
			NrFailedPCellId_r16 *struct {
				Choice           int
				CellGlobalId_r16 *CGI_Info_Logging_r16
				Pci_Arfcn_r16    *PCI_ARFCN_NR_r16
			}
			EutraFailedPCellId_r16 *struct {
				Choice           int
				CellGlobalId_r16 *CGI_InfoEUTRALogging
				Pci_Arfcn_r16    *PCI_ARFCN_EUTRA_r16
			}
		}
		ReconnectCellId_r16 *struct {
			Choice                   int
			NrReconnectCellId_r16    *CGI_Info_Logging_r16
			EutraReconnectCellId_r16 *CGI_InfoEUTRALogging
		}
		TimeUntilReconnection_r16 *TimeUntilReconnection_r16
		ReestablishmentCellId_r16 *CGI_Info_Logging_r16
		TimeConnFailure_r16       *int64
		TimeSinceFailure_r16      TimeSinceFailure_r16
		ConnectionFailureType_r16 int64
		Rlf_Cause_r16             int64
		LocationInfo_r16          *LocationInfo_r16
		NoSuitableCellFound_r16   *int64
		Ra_InformationCommon_r16  *RA_InformationCommon_r16
	}
	Eutra_RLF_Report_r16 *struct {
		FailedPCellId_EUTRA             CGI_InfoEUTRALogging
		MeasResult_RLF_Report_EUTRA_r16 []byte
	}
}

func (ie *RLF_Report_r16) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(rLFReportR16Constraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case RLF_Report_r16_Nr_RLF_Report_r16:
		c := (*ie.Nr_RLF_Report_r16)
		rLFReportR16NrRLFReportR16Seq := e.NewSequenceEncoder(rLFReportR16NrRLFReportR16Constraints)
		if err := rLFReportR16NrRLFReportR16Seq.EncodeExtensionBit(false); err != nil {
			return err
		}
		if err := rLFReportR16NrRLFReportR16Seq.EncodePreamble([]bool{c.MeasResultNeighCells_r16 != nil, c.PreviousPCellId_r16 != nil, c.ReconnectCellId_r16 != nil, c.TimeUntilReconnection_r16 != nil, c.ReestablishmentCellId_r16 != nil, c.TimeConnFailure_r16 != nil, c.LocationInfo_r16 != nil, c.NoSuitableCellFound_r16 != nil, c.Ra_InformationCommon_r16 != nil}); err != nil {
			return err
		}
		if err := c.MeasResultLastServCell_r16.Encode(e); err != nil {
			return err
		}
		if c.MeasResultNeighCells_r16 != nil {
			c := (*c.MeasResultNeighCells_r16)
			rLFReportR16NrRLFReportR16MeasResultNeighCellsR16Seq := e.NewSequenceEncoder(rLFReportR16NrRLFReportR16MeasResultNeighCellsR16Constraints)
			if err := rLFReportR16NrRLFReportR16MeasResultNeighCellsR16Seq.EncodePreamble([]bool{c.MeasResultListNR_r16 != nil, c.MeasResultListEUTRA_r16 != nil}); err != nil {
				return err
			}
			if c.MeasResultListNR_r16 != nil {
				if err := c.MeasResultListNR_r16.Encode(e); err != nil {
					return err
				}
			}
			if c.MeasResultListEUTRA_r16 != nil {
				if err := c.MeasResultListEUTRA_r16.Encode(e); err != nil {
					return err
				}
			}
		}
		if err := c.C_RNTI_r16.Encode(e); err != nil {
			return err
		}
		if c.PreviousPCellId_r16 != nil {
			choiceEnc := e.NewChoiceEncoder(rLFReportR16NrRLFReportR16PreviousPCellIdR16Constraints)
			if err := choiceEnc.EncodeChoice(int64((*c.PreviousPCellId_r16).Choice), false, nil); err != nil {
				return err
			}
			switch (*c.PreviousPCellId_r16).Choice {
			case RLF_Report_r16_Nr_RLF_Report_r16_PreviousPCellId_r16_NrPreviousCell_r16:
				if err := (*c.PreviousPCellId_r16).NrPreviousCell_r16.Encode(e); err != nil {
					return err
				}
			case RLF_Report_r16_Nr_RLF_Report_r16_PreviousPCellId_r16_EutraPreviousCell_r16:
				if err := (*c.PreviousPCellId_r16).EutraPreviousCell_r16.Encode(e); err != nil {
					return err
				}
			}
		}
		{
			choiceEnc := e.NewChoiceEncoder(rLFReportR16NrRLFReportR16FailedPCellIdR16Constraints)
			if err := choiceEnc.EncodeChoice(int64(c.FailedPCellId_r16.Choice), false, nil); err != nil {
				return err
			}
			switch c.FailedPCellId_r16.Choice {
			case RLF_Report_r16_Nr_RLF_Report_r16_FailedPCellId_r16_NrFailedPCellId_r16:
				choiceEnc := e.NewChoiceEncoder(rLFReportR16NrRLFReportR16FailedPCellIdR16NrFailedPCellIdR16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*c.FailedPCellId_r16.NrFailedPCellId_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*c.FailedPCellId_r16.NrFailedPCellId_r16).Choice {
				case RLF_Report_r16_Nr_RLF_Report_r16_FailedPCellId_r16_NrFailedPCellId_r16_CellGlobalId_r16:
					if err := (*c.FailedPCellId_r16.NrFailedPCellId_r16).CellGlobalId_r16.Encode(e); err != nil {
						return err
					}
				case RLF_Report_r16_Nr_RLF_Report_r16_FailedPCellId_r16_NrFailedPCellId_r16_Pci_Arfcn_r16:
					if err := (*c.FailedPCellId_r16.NrFailedPCellId_r16).Pci_Arfcn_r16.Encode(e); err != nil {
						return err
					}
				}
			case RLF_Report_r16_Nr_RLF_Report_r16_FailedPCellId_r16_EutraFailedPCellId_r16:
				choiceEnc := e.NewChoiceEncoder(rLFReportR16NrRLFReportR16FailedPCellIdR16EutraFailedPCellIdR16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*c.FailedPCellId_r16.EutraFailedPCellId_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*c.FailedPCellId_r16.EutraFailedPCellId_r16).Choice {
				case RLF_Report_r16_Nr_RLF_Report_r16_FailedPCellId_r16_EutraFailedPCellId_r16_CellGlobalId_r16:
					if err := (*c.FailedPCellId_r16.EutraFailedPCellId_r16).CellGlobalId_r16.Encode(e); err != nil {
						return err
					}
				case RLF_Report_r16_Nr_RLF_Report_r16_FailedPCellId_r16_EutraFailedPCellId_r16_Pci_Arfcn_r16:
					if err := (*c.FailedPCellId_r16.EutraFailedPCellId_r16).Pci_Arfcn_r16.Encode(e); err != nil {
						return err
					}
				}
			}
		}
		if c.ReconnectCellId_r16 != nil {
			choiceEnc := e.NewChoiceEncoder(rLFReportR16NrRLFReportR16ReconnectCellIdR16Constraints)
			if err := choiceEnc.EncodeChoice(int64((*c.ReconnectCellId_r16).Choice), false, nil); err != nil {
				return err
			}
			switch (*c.ReconnectCellId_r16).Choice {
			case RLF_Report_r16_Nr_RLF_Report_r16_ReconnectCellId_r16_NrReconnectCellId_r16:
				if err := (*c.ReconnectCellId_r16).NrReconnectCellId_r16.Encode(e); err != nil {
					return err
				}
			case RLF_Report_r16_Nr_RLF_Report_r16_ReconnectCellId_r16_EutraReconnectCellId_r16:
				if err := (*c.ReconnectCellId_r16).EutraReconnectCellId_r16.Encode(e); err != nil {
					return err
				}
			}
		}
		if c.TimeUntilReconnection_r16 != nil {
			if err := c.TimeUntilReconnection_r16.Encode(e); err != nil {
				return err
			}
		}
		if c.ReestablishmentCellId_r16 != nil {
			if err := c.ReestablishmentCellId_r16.Encode(e); err != nil {
				return err
			}
		}
		if c.TimeConnFailure_r16 != nil {
			if err := e.EncodeInteger((*c.TimeConnFailure_r16), per.Constrained(0, 1023)); err != nil {
				return err
			}
		}
		if err := c.TimeSinceFailure_r16.Encode(e); err != nil {
			return err
		}
		if err := e.EncodeEnumerated(c.ConnectionFailureType_r16, rLFReportR16NrRLFReportR16ConnectionFailureTypeR16Constraints); err != nil {
			return err
		}
		if err := e.EncodeEnumerated(c.Rlf_Cause_r16, rLFReportR16NrRLFReportR16RlfCauseR16Constraints); err != nil {
			return err
		}
		if c.LocationInfo_r16 != nil {
			if err := c.LocationInfo_r16.Encode(e); err != nil {
				return err
			}
		}
		if c.NoSuitableCellFound_r16 != nil {
			if err := e.EncodeEnumerated((*c.NoSuitableCellFound_r16), rLFReportR16NrRLFReportR16NoSuitableCellFoundR16Constraints); err != nil {
				return err
			}
		}
		if c.Ra_InformationCommon_r16 != nil {
			if err := c.Ra_InformationCommon_r16.Encode(e); err != nil {
				return err
			}
		}
	case RLF_Report_r16_Eutra_RLF_Report_r16:
		c := (*ie.Eutra_RLF_Report_r16)
		rLFReportR16EutraRLFReportR16Seq := e.NewSequenceEncoder(rLFReportR16EutraRLFReportR16Constraints)
		if err := rLFReportR16EutraRLFReportR16Seq.EncodeExtensionBit(false); err != nil {
			return err
		}
		if err := c.FailedPCellId_EUTRA.Encode(e); err != nil {
			return err
		}
		if err := e.EncodeOctetString(c.MeasResult_RLF_Report_EUTRA_r16, per.SizeConstraints{}); err != nil {
			return err
		}
	default:
		// Choice index out of root range.
		return &per.ConstraintViolationError{
			TypeName:   "RLF-Report-r16",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *RLF_Report_r16) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(rLFReportR16Constraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "RLF-Report-r16", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case RLF_Report_r16_Nr_RLF_Report_r16:
		ie.Nr_RLF_Report_r16 = &struct {
			MeasResultLastServCell_r16 MeasResultRLFNR_r16
			MeasResultNeighCells_r16   *struct {
				MeasResultListNR_r16    *MeasResultList2NR_r16
				MeasResultListEUTRA_r16 *MeasResultList2EUTRA_r16
			}
			C_RNTI_r16          RNTI_Value
			PreviousPCellId_r16 *struct {
				Choice                int
				NrPreviousCell_r16    *CGI_Info_Logging_r16
				EutraPreviousCell_r16 *CGI_InfoEUTRALogging
			}
			FailedPCellId_r16 struct {
				Choice              int
				NrFailedPCellId_r16 *struct {
					Choice           int
					CellGlobalId_r16 *CGI_Info_Logging_r16
					Pci_Arfcn_r16    *PCI_ARFCN_NR_r16
				}
				EutraFailedPCellId_r16 *struct {
					Choice           int
					CellGlobalId_r16 *CGI_InfoEUTRALogging
					Pci_Arfcn_r16    *PCI_ARFCN_EUTRA_r16
				}
			}
			ReconnectCellId_r16 *struct {
				Choice                   int
				NrReconnectCellId_r16    *CGI_Info_Logging_r16
				EutraReconnectCellId_r16 *CGI_InfoEUTRALogging
			}
			TimeUntilReconnection_r16 *TimeUntilReconnection_r16
			ReestablishmentCellId_r16 *CGI_Info_Logging_r16
			TimeConnFailure_r16       *int64
			TimeSinceFailure_r16      TimeSinceFailure_r16
			ConnectionFailureType_r16 int64
			Rlf_Cause_r16             int64
			LocationInfo_r16          *LocationInfo_r16
			NoSuitableCellFound_r16   *int64
			Ra_InformationCommon_r16  *RA_InformationCommon_r16
		}{}
		c := (*ie.Nr_RLF_Report_r16)
		rLFReportR16NrRLFReportR16Seq := d.NewSequenceDecoder(rLFReportR16NrRLFReportR16Constraints)
		if err := rLFReportR16NrRLFReportR16Seq.DecodeExtensionBit(); err != nil {
			return err
		}
		if err := rLFReportR16NrRLFReportR16Seq.DecodePreamble(); err != nil {
			return err
		}
		{
			if err := c.MeasResultLastServCell_r16.Decode(d); err != nil {
				return err
			}
		}
		if rLFReportR16NrRLFReportR16Seq.IsComponentPresent(1) {
			c.MeasResultNeighCells_r16 = &struct {
				MeasResultListNR_r16    *MeasResultList2NR_r16
				MeasResultListEUTRA_r16 *MeasResultList2EUTRA_r16
			}{}
			c := (*c.MeasResultNeighCells_r16)
			rLFReportR16NrRLFReportR16MeasResultNeighCellsR16Seq := d.NewSequenceDecoder(rLFReportR16NrRLFReportR16MeasResultNeighCellsR16Constraints)
			if err := rLFReportR16NrRLFReportR16MeasResultNeighCellsR16Seq.DecodePreamble(); err != nil {
				return err
			}
			if rLFReportR16NrRLFReportR16MeasResultNeighCellsR16Seq.IsComponentPresent(0) {
				c.MeasResultListNR_r16 = new(MeasResultList2NR_r16)
				if err := (*c.MeasResultListNR_r16).Decode(d); err != nil {
					return err
				}
			}
			if rLFReportR16NrRLFReportR16MeasResultNeighCellsR16Seq.IsComponentPresent(1) {
				c.MeasResultListEUTRA_r16 = new(MeasResultList2EUTRA_r16)
				if err := (*c.MeasResultListEUTRA_r16).Decode(d); err != nil {
					return err
				}
			}
		}
		{
			if err := c.C_RNTI_r16.Decode(d); err != nil {
				return err
			}
		}
		if rLFReportR16NrRLFReportR16Seq.IsComponentPresent(3) {
			c.PreviousPCellId_r16 = &struct {
				Choice                int
				NrPreviousCell_r16    *CGI_Info_Logging_r16
				EutraPreviousCell_r16 *CGI_InfoEUTRALogging
			}{}
			choiceDec := d.NewChoiceDecoder(rLFReportR16NrRLFReportR16PreviousPCellIdR16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*c.PreviousPCellId_r16).Choice = int(index)
			switch index {
			case RLF_Report_r16_Nr_RLF_Report_r16_PreviousPCellId_r16_NrPreviousCell_r16:
				(*c.PreviousPCellId_r16).NrPreviousCell_r16 = new(CGI_Info_Logging_r16)
				if err := (*c.PreviousPCellId_r16).NrPreviousCell_r16.Decode(d); err != nil {
					return err
				}
			case RLF_Report_r16_Nr_RLF_Report_r16_PreviousPCellId_r16_EutraPreviousCell_r16:
				(*c.PreviousPCellId_r16).EutraPreviousCell_r16 = new(CGI_InfoEUTRALogging)
				if err := (*c.PreviousPCellId_r16).EutraPreviousCell_r16.Decode(d); err != nil {
					return err
				}
			}
		}
		{
			choiceDec := d.NewChoiceDecoder(rLFReportR16NrRLFReportR16FailedPCellIdR16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			c.FailedPCellId_r16.Choice = int(index)
			switch index {
			case RLF_Report_r16_Nr_RLF_Report_r16_FailedPCellId_r16_NrFailedPCellId_r16:
				c.FailedPCellId_r16.NrFailedPCellId_r16 = &struct {
					Choice           int
					CellGlobalId_r16 *CGI_Info_Logging_r16
					Pci_Arfcn_r16    *PCI_ARFCN_NR_r16
				}{}
				choiceDec := d.NewChoiceDecoder(rLFReportR16NrRLFReportR16FailedPCellIdR16NrFailedPCellIdR16Constraints)
				index, _, _, err := choiceDec.DecodeChoice()
				if err != nil {
					return err
				}
				(*c.FailedPCellId_r16.NrFailedPCellId_r16).Choice = int(index)
				switch index {
				case RLF_Report_r16_Nr_RLF_Report_r16_FailedPCellId_r16_NrFailedPCellId_r16_CellGlobalId_r16:
					(*c.FailedPCellId_r16.NrFailedPCellId_r16).CellGlobalId_r16 = new(CGI_Info_Logging_r16)
					if err := (*c.FailedPCellId_r16.NrFailedPCellId_r16).CellGlobalId_r16.Decode(d); err != nil {
						return err
					}
				case RLF_Report_r16_Nr_RLF_Report_r16_FailedPCellId_r16_NrFailedPCellId_r16_Pci_Arfcn_r16:
					(*c.FailedPCellId_r16.NrFailedPCellId_r16).Pci_Arfcn_r16 = new(PCI_ARFCN_NR_r16)
					if err := (*c.FailedPCellId_r16.NrFailedPCellId_r16).Pci_Arfcn_r16.Decode(d); err != nil {
						return err
					}
				}
			case RLF_Report_r16_Nr_RLF_Report_r16_FailedPCellId_r16_EutraFailedPCellId_r16:
				c.FailedPCellId_r16.EutraFailedPCellId_r16 = &struct {
					Choice           int
					CellGlobalId_r16 *CGI_InfoEUTRALogging
					Pci_Arfcn_r16    *PCI_ARFCN_EUTRA_r16
				}{}
				choiceDec := d.NewChoiceDecoder(rLFReportR16NrRLFReportR16FailedPCellIdR16EutraFailedPCellIdR16Constraints)
				index, _, _, err := choiceDec.DecodeChoice()
				if err != nil {
					return err
				}
				(*c.FailedPCellId_r16.EutraFailedPCellId_r16).Choice = int(index)
				switch index {
				case RLF_Report_r16_Nr_RLF_Report_r16_FailedPCellId_r16_EutraFailedPCellId_r16_CellGlobalId_r16:
					(*c.FailedPCellId_r16.EutraFailedPCellId_r16).CellGlobalId_r16 = new(CGI_InfoEUTRALogging)
					if err := (*c.FailedPCellId_r16.EutraFailedPCellId_r16).CellGlobalId_r16.Decode(d); err != nil {
						return err
					}
				case RLF_Report_r16_Nr_RLF_Report_r16_FailedPCellId_r16_EutraFailedPCellId_r16_Pci_Arfcn_r16:
					(*c.FailedPCellId_r16.EutraFailedPCellId_r16).Pci_Arfcn_r16 = new(PCI_ARFCN_EUTRA_r16)
					if err := (*c.FailedPCellId_r16.EutraFailedPCellId_r16).Pci_Arfcn_r16.Decode(d); err != nil {
						return err
					}
				}
			}
		}
		if rLFReportR16NrRLFReportR16Seq.IsComponentPresent(5) {
			c.ReconnectCellId_r16 = &struct {
				Choice                   int
				NrReconnectCellId_r16    *CGI_Info_Logging_r16
				EutraReconnectCellId_r16 *CGI_InfoEUTRALogging
			}{}
			choiceDec := d.NewChoiceDecoder(rLFReportR16NrRLFReportR16ReconnectCellIdR16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*c.ReconnectCellId_r16).Choice = int(index)
			switch index {
			case RLF_Report_r16_Nr_RLF_Report_r16_ReconnectCellId_r16_NrReconnectCellId_r16:
				(*c.ReconnectCellId_r16).NrReconnectCellId_r16 = new(CGI_Info_Logging_r16)
				if err := (*c.ReconnectCellId_r16).NrReconnectCellId_r16.Decode(d); err != nil {
					return err
				}
			case RLF_Report_r16_Nr_RLF_Report_r16_ReconnectCellId_r16_EutraReconnectCellId_r16:
				(*c.ReconnectCellId_r16).EutraReconnectCellId_r16 = new(CGI_InfoEUTRALogging)
				if err := (*c.ReconnectCellId_r16).EutraReconnectCellId_r16.Decode(d); err != nil {
					return err
				}
			}
		}
		if rLFReportR16NrRLFReportR16Seq.IsComponentPresent(6) {
			c.TimeUntilReconnection_r16 = new(TimeUntilReconnection_r16)
			if err := (*c.TimeUntilReconnection_r16).Decode(d); err != nil {
				return err
			}
		}
		if rLFReportR16NrRLFReportR16Seq.IsComponentPresent(7) {
			c.ReestablishmentCellId_r16 = new(CGI_Info_Logging_r16)
			if err := (*c.ReestablishmentCellId_r16).Decode(d); err != nil {
				return err
			}
		}
		if rLFReportR16NrRLFReportR16Seq.IsComponentPresent(8) {
			c.TimeConnFailure_r16 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 1023))
			if err != nil {
				return err
			}
			(*c.TimeConnFailure_r16) = v
		}
		{
			if err := c.TimeSinceFailure_r16.Decode(d); err != nil {
				return err
			}
		}
		{
			v, err := d.DecodeEnumerated(rLFReportR16NrRLFReportR16ConnectionFailureTypeR16Constraints)
			if err != nil {
				return err
			}
			c.ConnectionFailureType_r16 = v
		}
		{
			v, err := d.DecodeEnumerated(rLFReportR16NrRLFReportR16RlfCauseR16Constraints)
			if err != nil {
				return err
			}
			c.Rlf_Cause_r16 = v
		}
		if rLFReportR16NrRLFReportR16Seq.IsComponentPresent(12) {
			c.LocationInfo_r16 = new(LocationInfo_r16)
			if err := (*c.LocationInfo_r16).Decode(d); err != nil {
				return err
			}
		}
		if rLFReportR16NrRLFReportR16Seq.IsComponentPresent(13) {
			c.NoSuitableCellFound_r16 = new(int64)
			v, err := d.DecodeEnumerated(rLFReportR16NrRLFReportR16NoSuitableCellFoundR16Constraints)
			if err != nil {
				return err
			}
			(*c.NoSuitableCellFound_r16) = v
		}
		if rLFReportR16NrRLFReportR16Seq.IsComponentPresent(14) {
			c.Ra_InformationCommon_r16 = new(RA_InformationCommon_r16)
			if err := (*c.Ra_InformationCommon_r16).Decode(d); err != nil {
				return err
			}
		}
	case RLF_Report_r16_Eutra_RLF_Report_r16:
		ie.Eutra_RLF_Report_r16 = &struct {
			FailedPCellId_EUTRA             CGI_InfoEUTRALogging
			MeasResult_RLF_Report_EUTRA_r16 []byte
		}{}
		c := (*ie.Eutra_RLF_Report_r16)
		rLFReportR16EutraRLFReportR16Seq := d.NewSequenceDecoder(rLFReportR16EutraRLFReportR16Constraints)
		if err := rLFReportR16EutraRLFReportR16Seq.DecodeExtensionBit(); err != nil {
			return err
		}
		{
			if err := c.FailedPCellId_EUTRA.Decode(d); err != nil {
				return err
			}
		}
		{
			v, err := d.DecodeOctetString(per.SizeConstraints{})
			if err != nil {
				return err
			}
			c.MeasResult_RLF_Report_EUTRA_r16 = v
		}
	default:
		return &per.DecodeError{TypeName: "RLF-Report-r16", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
