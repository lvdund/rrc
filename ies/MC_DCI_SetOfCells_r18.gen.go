// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MC-DCI-SetOfCells-r18 (line 14832).

var mCDCISetOfCellsR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "setOfCellsId-r18"},
		{Name: "nCI-Value-r18"},
		{Name: "scheduledCellListDCI-1-3-r18", Optional: true},
		{Name: "scheduledCellListDCI-0-3-r18", Optional: true},
		{Name: "scheduledCellComboListDCI-1-3-r18", Optional: true},
		{Name: "scheduledCellComboListDCI-0-3-r18", Optional: true},
		{Name: "antennaPortsDCI1-3-r18", Optional: true},
		{Name: "antennaPortsDCI0-3-r18", Optional: true},
		{Name: "tpmi-DCI0-3-r18", Optional: true},
		{Name: "sri-DCI0-3-r18", Optional: true},
		{Name: "priorityIndicatorDCI-1-3-r18", Optional: true},
		{Name: "priorityIndicatorDCI-0-3-r18", Optional: true},
		{Name: "dormancyDCI-1-3-r18", Optional: true},
		{Name: "dormancyDCI-0-3-r18", Optional: true},
		{Name: "pdcchMonAdaptDCI-1-3-r18", Optional: true},
		{Name: "pdcchMonAdaptDCI-0-3-r18", Optional: true},
		{Name: "minimumSchedulingOffsetK0DCI-1-3-r18", Optional: true},
		{Name: "minimumSchedulingOffsetK0DCI-0-3-r18", Optional: true},
		{Name: "pdsch-HARQ-ACK-OneShotFeedbackDCI-1-3-r18", Optional: true},
		{Name: "pdsch-HARQ-ACK-enhType3DCI-1-3-r18", Optional: true},
		{Name: "pdsch-HARQ-ACK-enhType3DCIfieldDCI-1-3-r18", Optional: true},
		{Name: "pdsch-HARQ-ACK-retxDCI-1-3-r18", Optional: true},
		{Name: "pucch-sSCellDynDCI-1-3-r18", Optional: true},
		{Name: "tdra-FieldIndexListDCI-1-3-r18", Optional: true},
		{Name: "tdra-FieldIndexListDCI-0-3-r18", Optional: true},
		{Name: "rateMatchListDCI-1-3-r18", Optional: true},
		{Name: "zp-CSI-RSListDCI-1-3-r18", Optional: true},
		{Name: "tci-ListDCI-1-3-r18", Optional: true},
		{Name: "srs-RequestListDCI-1-3-r18", Optional: true},
		{Name: "srs-OffsetListDCI-1-3-r18", Optional: true},
		{Name: "srs-RequestListDCI-0-3-r18", Optional: true},
		{Name: "srs-OffsetListDCI-0-3-r18", Optional: true},
	},
}

var mCDCISetOfCellsR18NCIValueR18Constraints = per.Constrained(0, 7)

var mCDCISetOfCellsR18ScheduledCellListDCI13R18Constraints = per.SizeRange(2, common.MaxNrofCellsInSet_r18)

var mCDCISetOfCellsR18ScheduledCellListDCI03R18Constraints = per.SizeRange(2, common.MaxNrofCellsInSet_r18)

var mCDCISetOfCellsR18ScheduledCellComboListDCI13R18Constraints = per.SizeRange(1, common.MaxNrofCellCombos_r18)

var mCDCISetOfCellsR18ScheduledCellComboListDCI03R18Constraints = per.SizeRange(1, common.MaxNrofCellCombos_r18)

const (
	MC_DCI_SetOfCells_r18_AntennaPortsDCI1_3_r18_Type1a = 0
	MC_DCI_SetOfCells_r18_AntennaPortsDCI1_3_r18_Type2  = 1
)

var mCDCISetOfCellsR18AntennaPortsDCI13R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MC_DCI_SetOfCells_r18_AntennaPortsDCI1_3_r18_Type1a, MC_DCI_SetOfCells_r18_AntennaPortsDCI1_3_r18_Type2},
}

const (
	MC_DCI_SetOfCells_r18_AntennaPortsDCI0_3_r18_Type1a = 0
	MC_DCI_SetOfCells_r18_AntennaPortsDCI0_3_r18_Type2  = 1
)

var mCDCISetOfCellsR18AntennaPortsDCI03R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MC_DCI_SetOfCells_r18_AntennaPortsDCI0_3_r18_Type1a, MC_DCI_SetOfCells_r18_AntennaPortsDCI0_3_r18_Type2},
}

const (
	MC_DCI_SetOfCells_r18_Tpmi_DCI0_3_r18_Type1a = 0
	MC_DCI_SetOfCells_r18_Tpmi_DCI0_3_r18_Type2  = 1
)

var mCDCISetOfCellsR18TpmiDCI03R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MC_DCI_SetOfCells_r18_Tpmi_DCI0_3_r18_Type1a, MC_DCI_SetOfCells_r18_Tpmi_DCI0_3_r18_Type2},
}

const (
	MC_DCI_SetOfCells_r18_Sri_DCI0_3_r18_Type1a = 0
	MC_DCI_SetOfCells_r18_Sri_DCI0_3_r18_Type2  = 1
)

var mCDCISetOfCellsR18SriDCI03R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MC_DCI_SetOfCells_r18_Sri_DCI0_3_r18_Type1a, MC_DCI_SetOfCells_r18_Sri_DCI0_3_r18_Type2},
}

const (
	MC_DCI_SetOfCells_r18_PriorityIndicatorDCI_1_3_r18_Enabled = 0
)

var mCDCISetOfCellsR18PriorityIndicatorDCI13R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MC_DCI_SetOfCells_r18_PriorityIndicatorDCI_1_3_r18_Enabled},
}

const (
	MC_DCI_SetOfCells_r18_PriorityIndicatorDCI_0_3_r18_Enabled = 0
)

var mCDCISetOfCellsR18PriorityIndicatorDCI03R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MC_DCI_SetOfCells_r18_PriorityIndicatorDCI_0_3_r18_Enabled},
}

const (
	MC_DCI_SetOfCells_r18_DormancyDCI_1_3_r18_Enabled = 0
)

var mCDCISetOfCellsR18DormancyDCI13R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MC_DCI_SetOfCells_r18_DormancyDCI_1_3_r18_Enabled},
}

const (
	MC_DCI_SetOfCells_r18_DormancyDCI_0_3_r18_Enabled = 0
)

var mCDCISetOfCellsR18DormancyDCI03R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MC_DCI_SetOfCells_r18_DormancyDCI_0_3_r18_Enabled},
}

const (
	MC_DCI_SetOfCells_r18_PdcchMonAdaptDCI_1_3_r18_Enabled = 0
)

var mCDCISetOfCellsR18PdcchMonAdaptDCI13R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MC_DCI_SetOfCells_r18_PdcchMonAdaptDCI_1_3_r18_Enabled},
}

const (
	MC_DCI_SetOfCells_r18_PdcchMonAdaptDCI_0_3_r18_Enabled = 0
)

var mCDCISetOfCellsR18PdcchMonAdaptDCI03R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MC_DCI_SetOfCells_r18_PdcchMonAdaptDCI_0_3_r18_Enabled},
}

const (
	MC_DCI_SetOfCells_r18_MinimumSchedulingOffsetK0DCI_1_3_r18_Enabled = 0
)

var mCDCISetOfCellsR18MinimumSchedulingOffsetK0DCI13R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MC_DCI_SetOfCells_r18_MinimumSchedulingOffsetK0DCI_1_3_r18_Enabled},
}

const (
	MC_DCI_SetOfCells_r18_MinimumSchedulingOffsetK0DCI_0_3_r18_Enabled = 0
)

var mCDCISetOfCellsR18MinimumSchedulingOffsetK0DCI03R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MC_DCI_SetOfCells_r18_MinimumSchedulingOffsetK0DCI_0_3_r18_Enabled},
}

const (
	MC_DCI_SetOfCells_r18_Pdsch_HARQ_ACK_OneShotFeedbackDCI_1_3_r18_Enabled = 0
)

var mCDCISetOfCellsR18PdschHARQACKOneShotFeedbackDCI13R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MC_DCI_SetOfCells_r18_Pdsch_HARQ_ACK_OneShotFeedbackDCI_1_3_r18_Enabled},
}

const (
	MC_DCI_SetOfCells_r18_Pdsch_HARQ_ACK_EnhType3DCI_1_3_r18_Enabled = 0
)

var mCDCISetOfCellsR18PdschHARQACKEnhType3DCI13R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MC_DCI_SetOfCells_r18_Pdsch_HARQ_ACK_EnhType3DCI_1_3_r18_Enabled},
}

const (
	MC_DCI_SetOfCells_r18_Pdsch_HARQ_ACK_EnhType3DCIfieldDCI_1_3_r18_Enabled = 0
)

var mCDCISetOfCellsR18PdschHARQACKEnhType3DCIfieldDCI13R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MC_DCI_SetOfCells_r18_Pdsch_HARQ_ACK_EnhType3DCIfieldDCI_1_3_r18_Enabled},
}

const (
	MC_DCI_SetOfCells_r18_Pdsch_HARQ_ACK_RetxDCI_1_3_r18_Enabled = 0
)

var mCDCISetOfCellsR18PdschHARQACKRetxDCI13R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MC_DCI_SetOfCells_r18_Pdsch_HARQ_ACK_RetxDCI_1_3_r18_Enabled},
}

const (
	MC_DCI_SetOfCells_r18_Pucch_SSCellDynDCI_1_3_r18_Enabled = 0
)

var mCDCISetOfCellsR18PucchSSCellDynDCI13R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MC_DCI_SetOfCells_r18_Pucch_SSCellDynDCI_1_3_r18_Enabled},
}

var mCDCISetOfCellsR18TdraFieldIndexListDCI13R18Constraints = per.SizeRange(1, 32)

var mCDCISetOfCellsR18TdraFieldIndexListDCI03R18Constraints = per.SizeRange(1, 64)

var mCDCISetOfCellsR18RateMatchListDCI13R18Constraints = per.SizeRange(1, 16)

var mCDCISetOfCellsR18ZpCSIRSListDCI13R18Constraints = per.SizeRange(1, 8)

var mCDCISetOfCellsR18TciListDCI13R18Constraints = per.SizeRange(1, 16)

var mCDCISetOfCellsR18SrsRequestListDCI13R18Constraints = per.SizeRange(1, 16)

var mCDCISetOfCellsR18SrsOffsetListDCI13R18Constraints = per.SizeRange(1, 8)

var mCDCISetOfCellsR18SrsRequestListDCI03R18Constraints = per.SizeRange(1, 16)

var mCDCISetOfCellsR18SrsOffsetListDCI03R18Constraints = per.SizeRange(1, 8)

type MC_DCI_SetOfCells_r18 struct {
	SetOfCellsId_r18                           SetOfCellsId_r18
	NCI_Value_r18                              int64
	ScheduledCellListDCI_1_3_r18               []ServCellIndex
	ScheduledCellListDCI_0_3_r18               []ServCellIndex
	ScheduledCellComboListDCI_1_3_r18          []ScheduledCellCombo_r18
	ScheduledCellComboListDCI_0_3_r18          []ScheduledCellCombo_r18
	AntennaPortsDCI1_3_r18                     *int64
	AntennaPortsDCI0_3_r18                     *int64
	Tpmi_DCI0_3_r18                            *int64
	Sri_DCI0_3_r18                             *int64
	PriorityIndicatorDCI_1_3_r18               *int64
	PriorityIndicatorDCI_0_3_r18               *int64
	DormancyDCI_1_3_r18                        *int64
	DormancyDCI_0_3_r18                        *int64
	PdcchMonAdaptDCI_1_3_r18                   *int64
	PdcchMonAdaptDCI_0_3_r18                   *int64
	MinimumSchedulingOffsetK0DCI_1_3_r18       *int64
	MinimumSchedulingOffsetK0DCI_0_3_r18       *int64
	Pdsch_HARQ_ACK_OneShotFeedbackDCI_1_3_r18  *int64
	Pdsch_HARQ_ACK_EnhType3DCI_1_3_r18         *int64
	Pdsch_HARQ_ACK_EnhType3DCIfieldDCI_1_3_r18 *int64
	Pdsch_HARQ_ACK_RetxDCI_1_3_r18             *int64
	Pucch_SSCellDynDCI_1_3_r18                 *int64
	Tdra_FieldIndexListDCI_1_3_r18             []TDRA_FieldIndexDCI_1_3_r18
	Tdra_FieldIndexListDCI_0_3_r18             []TDRA_FieldIndexDCI_0_3_r18
	RateMatchListDCI_1_3_r18                   []RateMatchDCI_1_3_r18
	Zp_CSI_RSListDCI_1_3_r18                   []ZP_CSI_DCI_1_3_r18
	Tci_ListDCI_1_3_r18                        []TCI_DCI_1_3_r18
	Srs_RequestListDCI_1_3_r18                 []SRS_RequestCombo_r18
	Srs_OffsetListDCI_1_3_r18                  []SRS_OffsetCombo_r18
	Srs_RequestListDCI_0_3_r18                 []SRS_RequestCombo_r18
	Srs_OffsetListDCI_0_3_r18                  []SRS_OffsetCombo_r18
}

func (ie *MC_DCI_SetOfCells_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mCDCISetOfCellsR18Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ScheduledCellListDCI_1_3_r18 != nil, ie.ScheduledCellListDCI_0_3_r18 != nil, ie.ScheduledCellComboListDCI_1_3_r18 != nil, ie.ScheduledCellComboListDCI_0_3_r18 != nil, ie.AntennaPortsDCI1_3_r18 != nil, ie.AntennaPortsDCI0_3_r18 != nil, ie.Tpmi_DCI0_3_r18 != nil, ie.Sri_DCI0_3_r18 != nil, ie.PriorityIndicatorDCI_1_3_r18 != nil, ie.PriorityIndicatorDCI_0_3_r18 != nil, ie.DormancyDCI_1_3_r18 != nil, ie.DormancyDCI_0_3_r18 != nil, ie.PdcchMonAdaptDCI_1_3_r18 != nil, ie.PdcchMonAdaptDCI_0_3_r18 != nil, ie.MinimumSchedulingOffsetK0DCI_1_3_r18 != nil, ie.MinimumSchedulingOffsetK0DCI_0_3_r18 != nil, ie.Pdsch_HARQ_ACK_OneShotFeedbackDCI_1_3_r18 != nil, ie.Pdsch_HARQ_ACK_EnhType3DCI_1_3_r18 != nil, ie.Pdsch_HARQ_ACK_EnhType3DCIfieldDCI_1_3_r18 != nil, ie.Pdsch_HARQ_ACK_RetxDCI_1_3_r18 != nil, ie.Pucch_SSCellDynDCI_1_3_r18 != nil, ie.Tdra_FieldIndexListDCI_1_3_r18 != nil, ie.Tdra_FieldIndexListDCI_0_3_r18 != nil, ie.RateMatchListDCI_1_3_r18 != nil, ie.Zp_CSI_RSListDCI_1_3_r18 != nil, ie.Tci_ListDCI_1_3_r18 != nil, ie.Srs_RequestListDCI_1_3_r18 != nil, ie.Srs_OffsetListDCI_1_3_r18 != nil, ie.Srs_RequestListDCI_0_3_r18 != nil, ie.Srs_OffsetListDCI_0_3_r18 != nil}); err != nil {
		return err
	}

	// 2. setOfCellsId-r18: ref
	{
		if err := ie.SetOfCellsId_r18.Encode(e); err != nil {
			return err
		}
	}

	// 3. nCI-Value-r18: integer
	{
		if err := e.EncodeInteger(ie.NCI_Value_r18, mCDCISetOfCellsR18NCIValueR18Constraints); err != nil {
			return err
		}
	}

	// 4. scheduledCellListDCI-1-3-r18: sequence-of
	{
		if ie.ScheduledCellListDCI_1_3_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(mCDCISetOfCellsR18ScheduledCellListDCI13R18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.ScheduledCellListDCI_1_3_r18))); err != nil {
				return err
			}
			for i := range ie.ScheduledCellListDCI_1_3_r18 {
				if err := ie.ScheduledCellListDCI_1_3_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 5. scheduledCellListDCI-0-3-r18: sequence-of
	{
		if ie.ScheduledCellListDCI_0_3_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(mCDCISetOfCellsR18ScheduledCellListDCI03R18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.ScheduledCellListDCI_0_3_r18))); err != nil {
				return err
			}
			for i := range ie.ScheduledCellListDCI_0_3_r18 {
				if err := ie.ScheduledCellListDCI_0_3_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 6. scheduledCellComboListDCI-1-3-r18: sequence-of
	{
		if ie.ScheduledCellComboListDCI_1_3_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(mCDCISetOfCellsR18ScheduledCellComboListDCI13R18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.ScheduledCellComboListDCI_1_3_r18))); err != nil {
				return err
			}
			for i := range ie.ScheduledCellComboListDCI_1_3_r18 {
				if err := ie.ScheduledCellComboListDCI_1_3_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 7. scheduledCellComboListDCI-0-3-r18: sequence-of
	{
		if ie.ScheduledCellComboListDCI_0_3_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(mCDCISetOfCellsR18ScheduledCellComboListDCI03R18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.ScheduledCellComboListDCI_0_3_r18))); err != nil {
				return err
			}
			for i := range ie.ScheduledCellComboListDCI_0_3_r18 {
				if err := ie.ScheduledCellComboListDCI_0_3_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 8. antennaPortsDCI1-3-r18: enumerated
	{
		if ie.AntennaPortsDCI1_3_r18 != nil {
			if err := e.EncodeEnumerated(*ie.AntennaPortsDCI1_3_r18, mCDCISetOfCellsR18AntennaPortsDCI13R18Constraints); err != nil {
				return err
			}
		}
	}

	// 9. antennaPortsDCI0-3-r18: enumerated
	{
		if ie.AntennaPortsDCI0_3_r18 != nil {
			if err := e.EncodeEnumerated(*ie.AntennaPortsDCI0_3_r18, mCDCISetOfCellsR18AntennaPortsDCI03R18Constraints); err != nil {
				return err
			}
		}
	}

	// 10. tpmi-DCI0-3-r18: enumerated
	{
		if ie.Tpmi_DCI0_3_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Tpmi_DCI0_3_r18, mCDCISetOfCellsR18TpmiDCI03R18Constraints); err != nil {
				return err
			}
		}
	}

	// 11. sri-DCI0-3-r18: enumerated
	{
		if ie.Sri_DCI0_3_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Sri_DCI0_3_r18, mCDCISetOfCellsR18SriDCI03R18Constraints); err != nil {
				return err
			}
		}
	}

	// 12. priorityIndicatorDCI-1-3-r18: enumerated
	{
		if ie.PriorityIndicatorDCI_1_3_r18 != nil {
			if err := e.EncodeEnumerated(*ie.PriorityIndicatorDCI_1_3_r18, mCDCISetOfCellsR18PriorityIndicatorDCI13R18Constraints); err != nil {
				return err
			}
		}
	}

	// 13. priorityIndicatorDCI-0-3-r18: enumerated
	{
		if ie.PriorityIndicatorDCI_0_3_r18 != nil {
			if err := e.EncodeEnumerated(*ie.PriorityIndicatorDCI_0_3_r18, mCDCISetOfCellsR18PriorityIndicatorDCI03R18Constraints); err != nil {
				return err
			}
		}
	}

	// 14. dormancyDCI-1-3-r18: enumerated
	{
		if ie.DormancyDCI_1_3_r18 != nil {
			if err := e.EncodeEnumerated(*ie.DormancyDCI_1_3_r18, mCDCISetOfCellsR18DormancyDCI13R18Constraints); err != nil {
				return err
			}
		}
	}

	// 15. dormancyDCI-0-3-r18: enumerated
	{
		if ie.DormancyDCI_0_3_r18 != nil {
			if err := e.EncodeEnumerated(*ie.DormancyDCI_0_3_r18, mCDCISetOfCellsR18DormancyDCI03R18Constraints); err != nil {
				return err
			}
		}
	}

	// 16. pdcchMonAdaptDCI-1-3-r18: enumerated
	{
		if ie.PdcchMonAdaptDCI_1_3_r18 != nil {
			if err := e.EncodeEnumerated(*ie.PdcchMonAdaptDCI_1_3_r18, mCDCISetOfCellsR18PdcchMonAdaptDCI13R18Constraints); err != nil {
				return err
			}
		}
	}

	// 17. pdcchMonAdaptDCI-0-3-r18: enumerated
	{
		if ie.PdcchMonAdaptDCI_0_3_r18 != nil {
			if err := e.EncodeEnumerated(*ie.PdcchMonAdaptDCI_0_3_r18, mCDCISetOfCellsR18PdcchMonAdaptDCI03R18Constraints); err != nil {
				return err
			}
		}
	}

	// 18. minimumSchedulingOffsetK0DCI-1-3-r18: enumerated
	{
		if ie.MinimumSchedulingOffsetK0DCI_1_3_r18 != nil {
			if err := e.EncodeEnumerated(*ie.MinimumSchedulingOffsetK0DCI_1_3_r18, mCDCISetOfCellsR18MinimumSchedulingOffsetK0DCI13R18Constraints); err != nil {
				return err
			}
		}
	}

	// 19. minimumSchedulingOffsetK0DCI-0-3-r18: enumerated
	{
		if ie.MinimumSchedulingOffsetK0DCI_0_3_r18 != nil {
			if err := e.EncodeEnumerated(*ie.MinimumSchedulingOffsetK0DCI_0_3_r18, mCDCISetOfCellsR18MinimumSchedulingOffsetK0DCI03R18Constraints); err != nil {
				return err
			}
		}
	}

	// 20. pdsch-HARQ-ACK-OneShotFeedbackDCI-1-3-r18: enumerated
	{
		if ie.Pdsch_HARQ_ACK_OneShotFeedbackDCI_1_3_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Pdsch_HARQ_ACK_OneShotFeedbackDCI_1_3_r18, mCDCISetOfCellsR18PdschHARQACKOneShotFeedbackDCI13R18Constraints); err != nil {
				return err
			}
		}
	}

	// 21. pdsch-HARQ-ACK-enhType3DCI-1-3-r18: enumerated
	{
		if ie.Pdsch_HARQ_ACK_EnhType3DCI_1_3_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Pdsch_HARQ_ACK_EnhType3DCI_1_3_r18, mCDCISetOfCellsR18PdschHARQACKEnhType3DCI13R18Constraints); err != nil {
				return err
			}
		}
	}

	// 22. pdsch-HARQ-ACK-enhType3DCIfieldDCI-1-3-r18: enumerated
	{
		if ie.Pdsch_HARQ_ACK_EnhType3DCIfieldDCI_1_3_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Pdsch_HARQ_ACK_EnhType3DCIfieldDCI_1_3_r18, mCDCISetOfCellsR18PdschHARQACKEnhType3DCIfieldDCI13R18Constraints); err != nil {
				return err
			}
		}
	}

	// 23. pdsch-HARQ-ACK-retxDCI-1-3-r18: enumerated
	{
		if ie.Pdsch_HARQ_ACK_RetxDCI_1_3_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Pdsch_HARQ_ACK_RetxDCI_1_3_r18, mCDCISetOfCellsR18PdschHARQACKRetxDCI13R18Constraints); err != nil {
				return err
			}
		}
	}

	// 24. pucch-sSCellDynDCI-1-3-r18: enumerated
	{
		if ie.Pucch_SSCellDynDCI_1_3_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Pucch_SSCellDynDCI_1_3_r18, mCDCISetOfCellsR18PucchSSCellDynDCI13R18Constraints); err != nil {
				return err
			}
		}
	}

	// 25. tdra-FieldIndexListDCI-1-3-r18: sequence-of
	{
		if ie.Tdra_FieldIndexListDCI_1_3_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(mCDCISetOfCellsR18TdraFieldIndexListDCI13R18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Tdra_FieldIndexListDCI_1_3_r18))); err != nil {
				return err
			}
			for i := range ie.Tdra_FieldIndexListDCI_1_3_r18 {
				if err := ie.Tdra_FieldIndexListDCI_1_3_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 26. tdra-FieldIndexListDCI-0-3-r18: sequence-of
	{
		if ie.Tdra_FieldIndexListDCI_0_3_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(mCDCISetOfCellsR18TdraFieldIndexListDCI03R18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Tdra_FieldIndexListDCI_0_3_r18))); err != nil {
				return err
			}
			for i := range ie.Tdra_FieldIndexListDCI_0_3_r18 {
				if err := ie.Tdra_FieldIndexListDCI_0_3_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 27. rateMatchListDCI-1-3-r18: sequence-of
	{
		if ie.RateMatchListDCI_1_3_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(mCDCISetOfCellsR18RateMatchListDCI13R18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.RateMatchListDCI_1_3_r18))); err != nil {
				return err
			}
			for i := range ie.RateMatchListDCI_1_3_r18 {
				if err := ie.RateMatchListDCI_1_3_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 28. zp-CSI-RSListDCI-1-3-r18: sequence-of
	{
		if ie.Zp_CSI_RSListDCI_1_3_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(mCDCISetOfCellsR18ZpCSIRSListDCI13R18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Zp_CSI_RSListDCI_1_3_r18))); err != nil {
				return err
			}
			for i := range ie.Zp_CSI_RSListDCI_1_3_r18 {
				if err := ie.Zp_CSI_RSListDCI_1_3_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 29. tci-ListDCI-1-3-r18: sequence-of
	{
		if ie.Tci_ListDCI_1_3_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(mCDCISetOfCellsR18TciListDCI13R18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Tci_ListDCI_1_3_r18))); err != nil {
				return err
			}
			for i := range ie.Tci_ListDCI_1_3_r18 {
				if err := ie.Tci_ListDCI_1_3_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 30. srs-RequestListDCI-1-3-r18: sequence-of
	{
		if ie.Srs_RequestListDCI_1_3_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(mCDCISetOfCellsR18SrsRequestListDCI13R18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Srs_RequestListDCI_1_3_r18))); err != nil {
				return err
			}
			for i := range ie.Srs_RequestListDCI_1_3_r18 {
				if err := ie.Srs_RequestListDCI_1_3_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 31. srs-OffsetListDCI-1-3-r18: sequence-of
	{
		if ie.Srs_OffsetListDCI_1_3_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(mCDCISetOfCellsR18SrsOffsetListDCI13R18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Srs_OffsetListDCI_1_3_r18))); err != nil {
				return err
			}
			for i := range ie.Srs_OffsetListDCI_1_3_r18 {
				if err := ie.Srs_OffsetListDCI_1_3_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 32. srs-RequestListDCI-0-3-r18: sequence-of
	{
		if ie.Srs_RequestListDCI_0_3_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(mCDCISetOfCellsR18SrsRequestListDCI03R18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Srs_RequestListDCI_0_3_r18))); err != nil {
				return err
			}
			for i := range ie.Srs_RequestListDCI_0_3_r18 {
				if err := ie.Srs_RequestListDCI_0_3_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 33. srs-OffsetListDCI-0-3-r18: sequence-of
	{
		if ie.Srs_OffsetListDCI_0_3_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(mCDCISetOfCellsR18SrsOffsetListDCI03R18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Srs_OffsetListDCI_0_3_r18))); err != nil {
				return err
			}
			for i := range ie.Srs_OffsetListDCI_0_3_r18 {
				if err := ie.Srs_OffsetListDCI_0_3_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *MC_DCI_SetOfCells_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mCDCISetOfCellsR18Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. setOfCellsId-r18: ref
	{
		if err := ie.SetOfCellsId_r18.Decode(d); err != nil {
			return err
		}
	}

	// 3. nCI-Value-r18: integer
	{
		v1, err := d.DecodeInteger(mCDCISetOfCellsR18NCIValueR18Constraints)
		if err != nil {
			return err
		}
		ie.NCI_Value_r18 = v1
	}

	// 4. scheduledCellListDCI-1-3-r18: sequence-of
	{
		if seq.IsComponentPresent(2) {
			seqOf := d.NewSequenceOfDecoder(mCDCISetOfCellsR18ScheduledCellListDCI13R18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.ScheduledCellListDCI_1_3_r18 = make([]ServCellIndex, n)
			for i := int64(0); i < n; i++ {
				if err := ie.ScheduledCellListDCI_1_3_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 5. scheduledCellListDCI-0-3-r18: sequence-of
	{
		if seq.IsComponentPresent(3) {
			seqOf := d.NewSequenceOfDecoder(mCDCISetOfCellsR18ScheduledCellListDCI03R18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.ScheduledCellListDCI_0_3_r18 = make([]ServCellIndex, n)
			for i := int64(0); i < n; i++ {
				if err := ie.ScheduledCellListDCI_0_3_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 6. scheduledCellComboListDCI-1-3-r18: sequence-of
	{
		if seq.IsComponentPresent(4) {
			seqOf := d.NewSequenceOfDecoder(mCDCISetOfCellsR18ScheduledCellComboListDCI13R18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.ScheduledCellComboListDCI_1_3_r18 = make([]ScheduledCellCombo_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.ScheduledCellComboListDCI_1_3_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 7. scheduledCellComboListDCI-0-3-r18: sequence-of
	{
		if seq.IsComponentPresent(5) {
			seqOf := d.NewSequenceOfDecoder(mCDCISetOfCellsR18ScheduledCellComboListDCI03R18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.ScheduledCellComboListDCI_0_3_r18 = make([]ScheduledCellCombo_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.ScheduledCellComboListDCI_0_3_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 8. antennaPortsDCI1-3-r18: enumerated
	{
		if seq.IsComponentPresent(6) {
			idx, err := d.DecodeEnumerated(mCDCISetOfCellsR18AntennaPortsDCI13R18Constraints)
			if err != nil {
				return err
			}
			ie.AntennaPortsDCI1_3_r18 = &idx
		}
	}

	// 9. antennaPortsDCI0-3-r18: enumerated
	{
		if seq.IsComponentPresent(7) {
			idx, err := d.DecodeEnumerated(mCDCISetOfCellsR18AntennaPortsDCI03R18Constraints)
			if err != nil {
				return err
			}
			ie.AntennaPortsDCI0_3_r18 = &idx
		}
	}

	// 10. tpmi-DCI0-3-r18: enumerated
	{
		if seq.IsComponentPresent(8) {
			idx, err := d.DecodeEnumerated(mCDCISetOfCellsR18TpmiDCI03R18Constraints)
			if err != nil {
				return err
			}
			ie.Tpmi_DCI0_3_r18 = &idx
		}
	}

	// 11. sri-DCI0-3-r18: enumerated
	{
		if seq.IsComponentPresent(9) {
			idx, err := d.DecodeEnumerated(mCDCISetOfCellsR18SriDCI03R18Constraints)
			if err != nil {
				return err
			}
			ie.Sri_DCI0_3_r18 = &idx
		}
	}

	// 12. priorityIndicatorDCI-1-3-r18: enumerated
	{
		if seq.IsComponentPresent(10) {
			idx, err := d.DecodeEnumerated(mCDCISetOfCellsR18PriorityIndicatorDCI13R18Constraints)
			if err != nil {
				return err
			}
			ie.PriorityIndicatorDCI_1_3_r18 = &idx
		}
	}

	// 13. priorityIndicatorDCI-0-3-r18: enumerated
	{
		if seq.IsComponentPresent(11) {
			idx, err := d.DecodeEnumerated(mCDCISetOfCellsR18PriorityIndicatorDCI03R18Constraints)
			if err != nil {
				return err
			}
			ie.PriorityIndicatorDCI_0_3_r18 = &idx
		}
	}

	// 14. dormancyDCI-1-3-r18: enumerated
	{
		if seq.IsComponentPresent(12) {
			idx, err := d.DecodeEnumerated(mCDCISetOfCellsR18DormancyDCI13R18Constraints)
			if err != nil {
				return err
			}
			ie.DormancyDCI_1_3_r18 = &idx
		}
	}

	// 15. dormancyDCI-0-3-r18: enumerated
	{
		if seq.IsComponentPresent(13) {
			idx, err := d.DecodeEnumerated(mCDCISetOfCellsR18DormancyDCI03R18Constraints)
			if err != nil {
				return err
			}
			ie.DormancyDCI_0_3_r18 = &idx
		}
	}

	// 16. pdcchMonAdaptDCI-1-3-r18: enumerated
	{
		if seq.IsComponentPresent(14) {
			idx, err := d.DecodeEnumerated(mCDCISetOfCellsR18PdcchMonAdaptDCI13R18Constraints)
			if err != nil {
				return err
			}
			ie.PdcchMonAdaptDCI_1_3_r18 = &idx
		}
	}

	// 17. pdcchMonAdaptDCI-0-3-r18: enumerated
	{
		if seq.IsComponentPresent(15) {
			idx, err := d.DecodeEnumerated(mCDCISetOfCellsR18PdcchMonAdaptDCI03R18Constraints)
			if err != nil {
				return err
			}
			ie.PdcchMonAdaptDCI_0_3_r18 = &idx
		}
	}

	// 18. minimumSchedulingOffsetK0DCI-1-3-r18: enumerated
	{
		if seq.IsComponentPresent(16) {
			idx, err := d.DecodeEnumerated(mCDCISetOfCellsR18MinimumSchedulingOffsetK0DCI13R18Constraints)
			if err != nil {
				return err
			}
			ie.MinimumSchedulingOffsetK0DCI_1_3_r18 = &idx
		}
	}

	// 19. minimumSchedulingOffsetK0DCI-0-3-r18: enumerated
	{
		if seq.IsComponentPresent(17) {
			idx, err := d.DecodeEnumerated(mCDCISetOfCellsR18MinimumSchedulingOffsetK0DCI03R18Constraints)
			if err != nil {
				return err
			}
			ie.MinimumSchedulingOffsetK0DCI_0_3_r18 = &idx
		}
	}

	// 20. pdsch-HARQ-ACK-OneShotFeedbackDCI-1-3-r18: enumerated
	{
		if seq.IsComponentPresent(18) {
			idx, err := d.DecodeEnumerated(mCDCISetOfCellsR18PdschHARQACKOneShotFeedbackDCI13R18Constraints)
			if err != nil {
				return err
			}
			ie.Pdsch_HARQ_ACK_OneShotFeedbackDCI_1_3_r18 = &idx
		}
	}

	// 21. pdsch-HARQ-ACK-enhType3DCI-1-3-r18: enumerated
	{
		if seq.IsComponentPresent(19) {
			idx, err := d.DecodeEnumerated(mCDCISetOfCellsR18PdschHARQACKEnhType3DCI13R18Constraints)
			if err != nil {
				return err
			}
			ie.Pdsch_HARQ_ACK_EnhType3DCI_1_3_r18 = &idx
		}
	}

	// 22. pdsch-HARQ-ACK-enhType3DCIfieldDCI-1-3-r18: enumerated
	{
		if seq.IsComponentPresent(20) {
			idx, err := d.DecodeEnumerated(mCDCISetOfCellsR18PdschHARQACKEnhType3DCIfieldDCI13R18Constraints)
			if err != nil {
				return err
			}
			ie.Pdsch_HARQ_ACK_EnhType3DCIfieldDCI_1_3_r18 = &idx
		}
	}

	// 23. pdsch-HARQ-ACK-retxDCI-1-3-r18: enumerated
	{
		if seq.IsComponentPresent(21) {
			idx, err := d.DecodeEnumerated(mCDCISetOfCellsR18PdschHARQACKRetxDCI13R18Constraints)
			if err != nil {
				return err
			}
			ie.Pdsch_HARQ_ACK_RetxDCI_1_3_r18 = &idx
		}
	}

	// 24. pucch-sSCellDynDCI-1-3-r18: enumerated
	{
		if seq.IsComponentPresent(22) {
			idx, err := d.DecodeEnumerated(mCDCISetOfCellsR18PucchSSCellDynDCI13R18Constraints)
			if err != nil {
				return err
			}
			ie.Pucch_SSCellDynDCI_1_3_r18 = &idx
		}
	}

	// 25. tdra-FieldIndexListDCI-1-3-r18: sequence-of
	{
		if seq.IsComponentPresent(23) {
			seqOf := d.NewSequenceOfDecoder(mCDCISetOfCellsR18TdraFieldIndexListDCI13R18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Tdra_FieldIndexListDCI_1_3_r18 = make([]TDRA_FieldIndexDCI_1_3_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Tdra_FieldIndexListDCI_1_3_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 26. tdra-FieldIndexListDCI-0-3-r18: sequence-of
	{
		if seq.IsComponentPresent(24) {
			seqOf := d.NewSequenceOfDecoder(mCDCISetOfCellsR18TdraFieldIndexListDCI03R18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Tdra_FieldIndexListDCI_0_3_r18 = make([]TDRA_FieldIndexDCI_0_3_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Tdra_FieldIndexListDCI_0_3_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 27. rateMatchListDCI-1-3-r18: sequence-of
	{
		if seq.IsComponentPresent(25) {
			seqOf := d.NewSequenceOfDecoder(mCDCISetOfCellsR18RateMatchListDCI13R18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.RateMatchListDCI_1_3_r18 = make([]RateMatchDCI_1_3_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.RateMatchListDCI_1_3_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 28. zp-CSI-RSListDCI-1-3-r18: sequence-of
	{
		if seq.IsComponentPresent(26) {
			seqOf := d.NewSequenceOfDecoder(mCDCISetOfCellsR18ZpCSIRSListDCI13R18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Zp_CSI_RSListDCI_1_3_r18 = make([]ZP_CSI_DCI_1_3_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Zp_CSI_RSListDCI_1_3_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 29. tci-ListDCI-1-3-r18: sequence-of
	{
		if seq.IsComponentPresent(27) {
			seqOf := d.NewSequenceOfDecoder(mCDCISetOfCellsR18TciListDCI13R18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Tci_ListDCI_1_3_r18 = make([]TCI_DCI_1_3_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Tci_ListDCI_1_3_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 30. srs-RequestListDCI-1-3-r18: sequence-of
	{
		if seq.IsComponentPresent(28) {
			seqOf := d.NewSequenceOfDecoder(mCDCISetOfCellsR18SrsRequestListDCI13R18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Srs_RequestListDCI_1_3_r18 = make([]SRS_RequestCombo_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Srs_RequestListDCI_1_3_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 31. srs-OffsetListDCI-1-3-r18: sequence-of
	{
		if seq.IsComponentPresent(29) {
			seqOf := d.NewSequenceOfDecoder(mCDCISetOfCellsR18SrsOffsetListDCI13R18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Srs_OffsetListDCI_1_3_r18 = make([]SRS_OffsetCombo_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Srs_OffsetListDCI_1_3_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 32. srs-RequestListDCI-0-3-r18: sequence-of
	{
		if seq.IsComponentPresent(30) {
			seqOf := d.NewSequenceOfDecoder(mCDCISetOfCellsR18SrsRequestListDCI03R18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Srs_RequestListDCI_0_3_r18 = make([]SRS_RequestCombo_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Srs_RequestListDCI_0_3_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 33. srs-OffsetListDCI-0-3-r18: sequence-of
	{
		if seq.IsComponentPresent(31) {
			seqOf := d.NewSequenceOfDecoder(mCDCISetOfCellsR18SrsOffsetListDCI03R18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Srs_OffsetListDCI_0_3_r18 = make([]SRS_OffsetCombo_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Srs_OffsetListDCI_0_3_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
