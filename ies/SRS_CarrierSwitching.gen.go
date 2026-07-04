// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SRS-CarrierSwitching (line 15287).

var sRSCarrierSwitchingConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "srs-SwitchFromServCellIndex", Optional: true},
		{Name: "srs-SwitchFromCarrier"},
		{Name: "srs-TPC-PDCCH-Group", Optional: true},
		{Name: "monitoringCells", Optional: true},
	},
}

var sRSCarrierSwitchingSrsSwitchFromServCellIndexConstraints = per.Constrained(0, 31)

const (
	SRS_CarrierSwitching_Srs_SwitchFromCarrier_SUL = 0
	SRS_CarrierSwitching_Srs_SwitchFromCarrier_NUL = 1
)

var sRSCarrierSwitchingSrsSwitchFromCarrierConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SRS_CarrierSwitching_Srs_SwitchFromCarrier_SUL, SRS_CarrierSwitching_Srs_SwitchFromCarrier_NUL},
}

var sRS_CarrierSwitchingSrsTPCPDCCHGroupConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "typeA"},
		{Name: "typeB"},
	},
}

const (
	SRS_CarrierSwitching_Srs_TPC_PDCCH_Group_TypeA = 0
	SRS_CarrierSwitching_Srs_TPC_PDCCH_Group_TypeB = 1
)

var sRSCarrierSwitchingMonitoringCellsConstraints = per.SizeRange(1, common.MaxNrofServingCells)

var sRSCarrierSwitchingSrsTPCPDCCHGroupTypeAConstraints = per.SizeRange(1, 32)

type SRS_CarrierSwitching struct {
	Srs_SwitchFromServCellIndex *int64
	Srs_SwitchFromCarrier       int64
	Srs_TPC_PDCCH_Group         *struct {
		Choice int
		TypeA  *[]SRS_TPC_PDCCH_Config
		TypeB  *SRS_TPC_PDCCH_Config
	}
	MonitoringCells []ServCellIndex
}

func (ie *SRS_CarrierSwitching) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sRSCarrierSwitchingConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Srs_SwitchFromServCellIndex != nil, ie.Srs_TPC_PDCCH_Group != nil, ie.MonitoringCells != nil}); err != nil {
		return err
	}

	// 3. srs-SwitchFromServCellIndex: integer
	{
		if ie.Srs_SwitchFromServCellIndex != nil {
			if err := e.EncodeInteger(*ie.Srs_SwitchFromServCellIndex, sRSCarrierSwitchingSrsSwitchFromServCellIndexConstraints); err != nil {
				return err
			}
		}
	}

	// 4. srs-SwitchFromCarrier: enumerated
	{
		if err := e.EncodeEnumerated(ie.Srs_SwitchFromCarrier, sRSCarrierSwitchingSrsSwitchFromCarrierConstraints); err != nil {
			return err
		}
	}

	// 5. srs-TPC-PDCCH-Group: choice
	{
		if ie.Srs_TPC_PDCCH_Group != nil {
			choiceEnc := e.NewChoiceEncoder(sRS_CarrierSwitchingSrsTPCPDCCHGroupConstraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Srs_TPC_PDCCH_Group).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Srs_TPC_PDCCH_Group).Choice {
			case SRS_CarrierSwitching_Srs_TPC_PDCCH_Group_TypeA:
				seqOf := e.NewSequenceOfEncoder(sRSCarrierSwitchingSrsTPCPDCCHGroupTypeAConstraints)
				if err := seqOf.EncodeLength(int64(len((*(*ie.Srs_TPC_PDCCH_Group).TypeA)))); err != nil {
					return err
				}
				for i := range *(*ie.Srs_TPC_PDCCH_Group).TypeA {
					if err := (*(*ie.Srs_TPC_PDCCH_Group).TypeA)[i].Encode(e); err != nil {
						return err
					}
				}
			case SRS_CarrierSwitching_Srs_TPC_PDCCH_Group_TypeB:
				if err := (*ie.Srs_TPC_PDCCH_Group).TypeB.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Srs_TPC_PDCCH_Group).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 6. monitoringCells: sequence-of
	{
		if ie.MonitoringCells != nil {
			seqOf := e.NewSequenceOfEncoder(sRSCarrierSwitchingMonitoringCellsConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.MonitoringCells))); err != nil {
				return err
			}
			for i := range ie.MonitoringCells {
				if err := ie.MonitoringCells[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *SRS_CarrierSwitching) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sRSCarrierSwitchingConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. srs-SwitchFromServCellIndex: integer
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeInteger(sRSCarrierSwitchingSrsSwitchFromServCellIndexConstraints)
			if err != nil {
				return err
			}
			ie.Srs_SwitchFromServCellIndex = &v
		}
	}

	// 4. srs-SwitchFromCarrier: enumerated
	{
		v1, err := d.DecodeEnumerated(sRSCarrierSwitchingSrsSwitchFromCarrierConstraints)
		if err != nil {
			return err
		}
		ie.Srs_SwitchFromCarrier = v1
	}

	// 5. srs-TPC-PDCCH-Group: choice
	{
		if seq.IsComponentPresent(2) {
			ie.Srs_TPC_PDCCH_Group = &struct {
				Choice int
				TypeA  *[]SRS_TPC_PDCCH_Config
				TypeB  *SRS_TPC_PDCCH_Config
			}{}
			choiceDec := d.NewChoiceDecoder(sRS_CarrierSwitchingSrsTPCPDCCHGroupConstraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Srs_TPC_PDCCH_Group).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case SRS_CarrierSwitching_Srs_TPC_PDCCH_Group_TypeA:
				(*ie.Srs_TPC_PDCCH_Group).TypeA = new([]SRS_TPC_PDCCH_Config)
				seqOf := d.NewSequenceOfDecoder(sRSCarrierSwitchingSrsTPCPDCCHGroupTypeAConstraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				(*(*ie.Srs_TPC_PDCCH_Group).TypeA) = make([]SRS_TPC_PDCCH_Config, n)
				for i := int64(0); i < n; i++ {
					if err := (*(*ie.Srs_TPC_PDCCH_Group).TypeA)[i].Decode(d); err != nil {
						return err
					}
				}
			case SRS_CarrierSwitching_Srs_TPC_PDCCH_Group_TypeB:
				(*ie.Srs_TPC_PDCCH_Group).TypeB = new(SRS_TPC_PDCCH_Config)
				if err := (*ie.Srs_TPC_PDCCH_Group).TypeB.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 6. monitoringCells: sequence-of
	{
		if seq.IsComponentPresent(3) {
			seqOf := d.NewSequenceOfDecoder(sRSCarrierSwitchingMonitoringCellsConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.MonitoringCells = make([]ServCellIndex, n)
			for i := int64(0); i < n; i++ {
				if err := ie.MonitoringCells[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
