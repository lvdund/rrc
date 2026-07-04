// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: Phy-ParametersFR1 (line 23198).

var phyParametersFR1Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "pdcch-MonitoringSingleOccasion", Optional: true},
		{Name: "scs-60kHz", Optional: true},
		{Name: "pdsch-256QAM-FR1", Optional: true},
		{Name: "pdsch-RE-MappingFR1-PerSymbol", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
		{Name: "extension-addition-2", Optional: true},
	},
}

const (
	Phy_ParametersFR1_Pdcch_MonitoringSingleOccasion_Supported = 0
)

var phyParametersFR1PdcchMonitoringSingleOccasionConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFR1_Pdcch_MonitoringSingleOccasion_Supported},
}

const (
	Phy_ParametersFR1_Scs_60kHz_Supported = 0
)

var phyParametersFR1Scs60kHzConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFR1_Scs_60kHz_Supported},
}

const (
	Phy_ParametersFR1_Pdsch_256QAM_FR1_Supported = 0
)

var phyParametersFR1Pdsch256QAMFR1Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFR1_Pdsch_256QAM_FR1_Supported},
}

const (
	Phy_ParametersFR1_Pdsch_RE_MappingFR1_PerSymbol_N10 = 0
	Phy_ParametersFR1_Pdsch_RE_MappingFR1_PerSymbol_N20 = 1
)

var phyParametersFR1PdschREMappingFR1PerSymbolConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFR1_Pdsch_RE_MappingFR1_PerSymbol_N10, Phy_ParametersFR1_Pdsch_RE_MappingFR1_PerSymbol_N20},
}

const (
	Phy_ParametersFR1_Ext_Pdsch_RE_MappingFR1_PerSlot_N16  = 0
	Phy_ParametersFR1_Ext_Pdsch_RE_MappingFR1_PerSlot_N32  = 1
	Phy_ParametersFR1_Ext_Pdsch_RE_MappingFR1_PerSlot_N48  = 2
	Phy_ParametersFR1_Ext_Pdsch_RE_MappingFR1_PerSlot_N64  = 3
	Phy_ParametersFR1_Ext_Pdsch_RE_MappingFR1_PerSlot_N80  = 4
	Phy_ParametersFR1_Ext_Pdsch_RE_MappingFR1_PerSlot_N96  = 5
	Phy_ParametersFR1_Ext_Pdsch_RE_MappingFR1_PerSlot_N112 = 6
	Phy_ParametersFR1_Ext_Pdsch_RE_MappingFR1_PerSlot_N128 = 7
	Phy_ParametersFR1_Ext_Pdsch_RE_MappingFR1_PerSlot_N144 = 8
	Phy_ParametersFR1_Ext_Pdsch_RE_MappingFR1_PerSlot_N160 = 9
	Phy_ParametersFR1_Ext_Pdsch_RE_MappingFR1_PerSlot_N176 = 10
	Phy_ParametersFR1_Ext_Pdsch_RE_MappingFR1_PerSlot_N192 = 11
	Phy_ParametersFR1_Ext_Pdsch_RE_MappingFR1_PerSlot_N208 = 12
	Phy_ParametersFR1_Ext_Pdsch_RE_MappingFR1_PerSlot_N224 = 13
	Phy_ParametersFR1_Ext_Pdsch_RE_MappingFR1_PerSlot_N240 = 14
	Phy_ParametersFR1_Ext_Pdsch_RE_MappingFR1_PerSlot_N256 = 15
)

var phyParametersFR1ExtPdschREMappingFR1PerSlotConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFR1_Ext_Pdsch_RE_MappingFR1_PerSlot_N16, Phy_ParametersFR1_Ext_Pdsch_RE_MappingFR1_PerSlot_N32, Phy_ParametersFR1_Ext_Pdsch_RE_MappingFR1_PerSlot_N48, Phy_ParametersFR1_Ext_Pdsch_RE_MappingFR1_PerSlot_N64, Phy_ParametersFR1_Ext_Pdsch_RE_MappingFR1_PerSlot_N80, Phy_ParametersFR1_Ext_Pdsch_RE_MappingFR1_PerSlot_N96, Phy_ParametersFR1_Ext_Pdsch_RE_MappingFR1_PerSlot_N112, Phy_ParametersFR1_Ext_Pdsch_RE_MappingFR1_PerSlot_N128, Phy_ParametersFR1_Ext_Pdsch_RE_MappingFR1_PerSlot_N144, Phy_ParametersFR1_Ext_Pdsch_RE_MappingFR1_PerSlot_N160, Phy_ParametersFR1_Ext_Pdsch_RE_MappingFR1_PerSlot_N176, Phy_ParametersFR1_Ext_Pdsch_RE_MappingFR1_PerSlot_N192, Phy_ParametersFR1_Ext_Pdsch_RE_MappingFR1_PerSlot_N208, Phy_ParametersFR1_Ext_Pdsch_RE_MappingFR1_PerSlot_N224, Phy_ParametersFR1_Ext_Pdsch_RE_MappingFR1_PerSlot_N240, Phy_ParametersFR1_Ext_Pdsch_RE_MappingFR1_PerSlot_N256},
}

const (
	Phy_ParametersFR1_Ext_Pdcch_MonitoringSingleSpanFirst4Sym_r16_Supported = 0
)

var phyParametersFR1ExtPdcchMonitoringSingleSpanFirst4SymR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFR1_Ext_Pdcch_MonitoringSingleSpanFirst4Sym_r16_Supported},
}

const (
	Phy_ParametersFR1_Ext_K1_RangeExtensionATG_r18_Supported = 0
)

var phyParametersFR1ExtK1RangeExtensionATGR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFR1_Ext_K1_RangeExtensionATG_r18_Supported},
}

const (
	Phy_ParametersFR1_Ext_MaxHARQ_ProcessNumberATG_r18_U16d32 = 0
	Phy_ParametersFR1_Ext_MaxHARQ_ProcessNumberATG_r18_U32d16 = 1
	Phy_ParametersFR1_Ext_MaxHARQ_ProcessNumberATG_r18_U32d32 = 2
)

var phyParametersFR1ExtMaxHARQProcessNumberATGR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFR1_Ext_MaxHARQ_ProcessNumberATG_r18_U16d32, Phy_ParametersFR1_Ext_MaxHARQ_ProcessNumberATG_r18_U32d16, Phy_ParametersFR1_Ext_MaxHARQ_ProcessNumberATG_r18_U32d32},
}

const (
	Phy_ParametersFR1_Ext_UplinkPreCompensationATG_r18_Supported = 0
)

var phyParametersFR1ExtUplinkPreCompensationATGR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFR1_Ext_UplinkPreCompensationATG_r18_Supported},
}

const (
	Phy_ParametersFR1_Ext_UplinkTA_ReportingATG_r18_Supported = 0
)

var phyParametersFR1ExtUplinkTAReportingATGR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFR1_Ext_UplinkTA_ReportingATG_r18_Supported},
}

const (
	Phy_ParametersFR1_Ext_AdvReceiver_MU_MIMO_r18_Supported = 0
)

var phyParametersFR1ExtAdvReceiverMUMIMOR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFR1_Ext_AdvReceiver_MU_MIMO_r18_Supported},
}

const (
	Phy_ParametersFR1_Ext_DeltaPowerClassReporting_r18_Type1 = 0
	Phy_ParametersFR1_Ext_DeltaPowerClassReporting_r18_Type2 = 1
)

var phyParametersFR1ExtDeltaPowerClassReportingR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFR1_Ext_DeltaPowerClassReporting_r18_Type1, Phy_ParametersFR1_Ext_DeltaPowerClassReporting_r18_Type2},
}

const (
	Phy_ParametersFR1_Ext_Support12PRB_CORESET0_GSCN_41637_r18_Supported = 0
)

var phyParametersFR1ExtSupport12PRBCORESET0GSCN41637R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFR1_Ext_Support12PRB_CORESET0_GSCN_41637_r18_Supported},
}

const (
	Phy_ParametersFR1_Ext_Support5MHz_ChannelBW_20PRB_CORESET0_r18_Supported = 0
)

var phyParametersFR1ExtSupport5MHzChannelBW20PRBCORESET0R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFR1_Ext_Support5MHz_ChannelBW_20PRB_CORESET0_r18_Supported},
}

type Phy_ParametersFR1 struct {
	Pdcch_MonitoringSingleOccasion           *int64
	Scs_60kHz                                *int64
	Pdsch_256QAM_FR1                         *int64
	Pdsch_RE_MappingFR1_PerSymbol            *int64
	Pdsch_RE_MappingFR1_PerSlot              *int64
	Pdcch_MonitoringSingleSpanFirst4Sym_r16  *int64
	K1_RangeExtensionATG_r18                 *int64
	MaxHARQ_ProcessNumberATG_r18             *int64
	UplinkPreCompensationATG_r18             *int64
	UplinkTA_ReportingATG_r18                *int64
	AdvReceiver_MU_MIMO_r18                  *int64
	DeltaPowerClassReporting_r18             *int64
	Support12PRB_CORESET0_GSCN_41637_r18     *int64
	Support5MHz_ChannelBW_20PRB_CORESET0_r18 *int64
}

func (ie *Phy_ParametersFR1) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(phyParametersFR1Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Pdsch_RE_MappingFR1_PerSlot != nil
	hasExtGroup1 := ie.Pdcch_MonitoringSingleSpanFirst4Sym_r16 != nil
	hasExtGroup2 := ie.K1_RangeExtensionATG_r18 != nil || ie.MaxHARQ_ProcessNumberATG_r18 != nil || ie.UplinkPreCompensationATG_r18 != nil || ie.UplinkTA_ReportingATG_r18 != nil || ie.AdvReceiver_MU_MIMO_r18 != nil || ie.DeltaPowerClassReporting_r18 != nil || ie.Support12PRB_CORESET0_GSCN_41637_r18 != nil || ie.Support5MHz_ChannelBW_20PRB_CORESET0_r18 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Pdcch_MonitoringSingleOccasion != nil, ie.Scs_60kHz != nil, ie.Pdsch_256QAM_FR1 != nil, ie.Pdsch_RE_MappingFR1_PerSymbol != nil}); err != nil {
		return err
	}

	// 3. pdcch-MonitoringSingleOccasion: enumerated
	{
		if ie.Pdcch_MonitoringSingleOccasion != nil {
			if err := e.EncodeEnumerated(*ie.Pdcch_MonitoringSingleOccasion, phyParametersFR1PdcchMonitoringSingleOccasionConstraints); err != nil {
				return err
			}
		}
	}

	// 4. scs-60kHz: enumerated
	{
		if ie.Scs_60kHz != nil {
			if err := e.EncodeEnumerated(*ie.Scs_60kHz, phyParametersFR1Scs60kHzConstraints); err != nil {
				return err
			}
		}
	}

	// 5. pdsch-256QAM-FR1: enumerated
	{
		if ie.Pdsch_256QAM_FR1 != nil {
			if err := e.EncodeEnumerated(*ie.Pdsch_256QAM_FR1, phyParametersFR1Pdsch256QAMFR1Constraints); err != nil {
				return err
			}
		}
	}

	// 6. pdsch-RE-MappingFR1-PerSymbol: enumerated
	{
		if ie.Pdsch_RE_MappingFR1_PerSymbol != nil {
			if err := e.EncodeEnumerated(*ie.Pdsch_RE_MappingFR1_PerSymbol, phyParametersFR1PdschREMappingFR1PerSymbolConstraints); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1, hasExtGroup2}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "pdsch-RE-MappingFR1-PerSlot", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Pdsch_RE_MappingFR1_PerSlot != nil}); err != nil {
				return err
			}

			if ie.Pdsch_RE_MappingFR1_PerSlot != nil {
				if err := ex.EncodeEnumerated(*ie.Pdsch_RE_MappingFR1_PerSlot, phyParametersFR1ExtPdschREMappingFR1PerSlotConstraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 1:
		if hasExtGroup1 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "pdcch-MonitoringSingleSpanFirst4Sym-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Pdcch_MonitoringSingleSpanFirst4Sym_r16 != nil}); err != nil {
				return err
			}

			if ie.Pdcch_MonitoringSingleSpanFirst4Sym_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Pdcch_MonitoringSingleSpanFirst4Sym_r16, phyParametersFR1ExtPdcchMonitoringSingleSpanFirst4SymR16Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 2:
		if hasExtGroup2 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "k1-RangeExtensionATG-r18", Optional: true},
					{Name: "maxHARQ-ProcessNumberATG-r18", Optional: true},
					{Name: "uplinkPreCompensationATG-r18", Optional: true},
					{Name: "uplinkTA-ReportingATG-r18", Optional: true},
					{Name: "advReceiver-MU-MIMO-r18", Optional: true},
					{Name: "deltaPowerClassReporting-r18", Optional: true},
					{Name: "support12PRB-CORESET0-GSCN-41637-r18", Optional: true},
					{Name: "support5MHz-ChannelBW-20PRB-CORESET0-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.K1_RangeExtensionATG_r18 != nil, ie.MaxHARQ_ProcessNumberATG_r18 != nil, ie.UplinkPreCompensationATG_r18 != nil, ie.UplinkTA_ReportingATG_r18 != nil, ie.AdvReceiver_MU_MIMO_r18 != nil, ie.DeltaPowerClassReporting_r18 != nil, ie.Support12PRB_CORESET0_GSCN_41637_r18 != nil, ie.Support5MHz_ChannelBW_20PRB_CORESET0_r18 != nil}); err != nil {
				return err
			}

			if ie.K1_RangeExtensionATG_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.K1_RangeExtensionATG_r18, phyParametersFR1ExtK1RangeExtensionATGR18Constraints); err != nil {
					return err
				}
			}

			if ie.MaxHARQ_ProcessNumberATG_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.MaxHARQ_ProcessNumberATG_r18, phyParametersFR1ExtMaxHARQProcessNumberATGR18Constraints); err != nil {
					return err
				}
			}

			if ie.UplinkPreCompensationATG_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.UplinkPreCompensationATG_r18, phyParametersFR1ExtUplinkPreCompensationATGR18Constraints); err != nil {
					return err
				}
			}

			if ie.UplinkTA_ReportingATG_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.UplinkTA_ReportingATG_r18, phyParametersFR1ExtUplinkTAReportingATGR18Constraints); err != nil {
					return err
				}
			}

			if ie.AdvReceiver_MU_MIMO_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.AdvReceiver_MU_MIMO_r18, phyParametersFR1ExtAdvReceiverMUMIMOR18Constraints); err != nil {
					return err
				}
			}

			if ie.DeltaPowerClassReporting_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.DeltaPowerClassReporting_r18, phyParametersFR1ExtDeltaPowerClassReportingR18Constraints); err != nil {
					return err
				}
			}

			if ie.Support12PRB_CORESET0_GSCN_41637_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Support12PRB_CORESET0_GSCN_41637_r18, phyParametersFR1ExtSupport12PRBCORESET0GSCN41637R18Constraints); err != nil {
					return err
				}
			}

			if ie.Support5MHz_ChannelBW_20PRB_CORESET0_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Support5MHz_ChannelBW_20PRB_CORESET0_r18, phyParametersFR1ExtSupport5MHzChannelBW20PRBCORESET0R18Constraints); err != nil {
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

func (ie *Phy_ParametersFR1) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(phyParametersFR1Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. pdcch-MonitoringSingleOccasion: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(phyParametersFR1PdcchMonitoringSingleOccasionConstraints)
			if err != nil {
				return err
			}
			ie.Pdcch_MonitoringSingleOccasion = &idx
		}
	}

	// 4. scs-60kHz: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(phyParametersFR1Scs60kHzConstraints)
			if err != nil {
				return err
			}
			ie.Scs_60kHz = &idx
		}
	}

	// 5. pdsch-256QAM-FR1: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(phyParametersFR1Pdsch256QAMFR1Constraints)
			if err != nil {
				return err
			}
			ie.Pdsch_256QAM_FR1 = &idx
		}
	}

	// 6. pdsch-RE-MappingFR1-PerSymbol: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(phyParametersFR1PdschREMappingFR1PerSymbolConstraints)
			if err != nil {
				return err
			}
			ie.Pdsch_RE_MappingFR1_PerSymbol = &idx
		}
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
				{Name: "pdsch-RE-MappingFR1-PerSlot", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(phyParametersFR1ExtPdschREMappingFR1PerSlotConstraints)
			if err != nil {
				return err
			}
			ie.Pdsch_RE_MappingFR1_PerSlot = &v
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "pdcch-MonitoringSingleSpanFirst4Sym-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(phyParametersFR1ExtPdcchMonitoringSingleSpanFirst4SymR16Constraints)
			if err != nil {
				return err
			}
			ie.Pdcch_MonitoringSingleSpanFirst4Sym_r16 = &v
		}
	}

	// Extension group 2:
	if seq.IsExtensionPresent(2) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "k1-RangeExtensionATG-r18", Optional: true},
				{Name: "maxHARQ-ProcessNumberATG-r18", Optional: true},
				{Name: "uplinkPreCompensationATG-r18", Optional: true},
				{Name: "uplinkTA-ReportingATG-r18", Optional: true},
				{Name: "advReceiver-MU-MIMO-r18", Optional: true},
				{Name: "deltaPowerClassReporting-r18", Optional: true},
				{Name: "support12PRB-CORESET0-GSCN-41637-r18", Optional: true},
				{Name: "support5MHz-ChannelBW-20PRB-CORESET0-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(phyParametersFR1ExtK1RangeExtensionATGR18Constraints)
			if err != nil {
				return err
			}
			ie.K1_RangeExtensionATG_r18 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(phyParametersFR1ExtMaxHARQProcessNumberATGR18Constraints)
			if err != nil {
				return err
			}
			ie.MaxHARQ_ProcessNumberATG_r18 = &v
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeEnumerated(phyParametersFR1ExtUplinkPreCompensationATGR18Constraints)
			if err != nil {
				return err
			}
			ie.UplinkPreCompensationATG_r18 = &v
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeEnumerated(phyParametersFR1ExtUplinkTAReportingATGR18Constraints)
			if err != nil {
				return err
			}
			ie.UplinkTA_ReportingATG_r18 = &v
		}

		if groupSeq.IsComponentPresent(4) {
			v, err := dx.DecodeEnumerated(phyParametersFR1ExtAdvReceiverMUMIMOR18Constraints)
			if err != nil {
				return err
			}
			ie.AdvReceiver_MU_MIMO_r18 = &v
		}

		if groupSeq.IsComponentPresent(5) {
			v, err := dx.DecodeEnumerated(phyParametersFR1ExtDeltaPowerClassReportingR18Constraints)
			if err != nil {
				return err
			}
			ie.DeltaPowerClassReporting_r18 = &v
		}

		if groupSeq.IsComponentPresent(6) {
			v, err := dx.DecodeEnumerated(phyParametersFR1ExtSupport12PRBCORESET0GSCN41637R18Constraints)
			if err != nil {
				return err
			}
			ie.Support12PRB_CORESET0_GSCN_41637_r18 = &v
		}

		if groupSeq.IsComponentPresent(7) {
			v, err := dx.DecodeEnumerated(phyParametersFR1ExtSupport5MHzChannelBW20PRBCORESET0R18Constraints)
			if err != nil {
				return err
			}
			ie.Support5MHz_ChannelBW_20PRB_CORESET0_r18 = &v
		}
	}

	return nil
}
