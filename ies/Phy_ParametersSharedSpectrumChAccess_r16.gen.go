// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: Phy-ParametersSharedSpectrumChAccess-r16 (line 23282).

var phyParametersSharedSpectrumChAccessR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "ss-SINR-Meas-r16", Optional: true},
		{Name: "sp-CSI-ReportPUCCH-r16", Optional: true},
		{Name: "sp-CSI-ReportPUSCH-r16", Optional: true},
		{Name: "dynamicSFI-r16", Optional: true},
		{Name: "mux-SR-HARQ-ACK-CSI-PUCCH-OncePerSlot-r16", Optional: true},
		{Name: "mux-SR-HARQ-ACK-PUCCH-r16", Optional: true},
		{Name: "mux-SR-HARQ-ACK-CSI-PUCCH-MultiPerSlot-r16", Optional: true},
		{Name: "mux-HARQ-ACK-PUSCH-DiffSymbol-r16", Optional: true},
		{Name: "pucch-Repetition-F1-3-4-r16", Optional: true},
		{Name: "type1-PUSCH-RepetitionMultiSlots-r16", Optional: true},
		{Name: "type2-PUSCH-RepetitionMultiSlots-r16", Optional: true},
		{Name: "pusch-RepetitionMultiSlots-r16", Optional: true},
		{Name: "pdsch-RepetitionMultiSlots-r16", Optional: true},
		{Name: "downlinkSPS-r16", Optional: true},
		{Name: "configuredUL-GrantType1-r16", Optional: true},
		{Name: "configuredUL-GrantType2-r16", Optional: true},
		{Name: "pre-EmptIndication-DL-r16", Optional: true},
	},
}

const (
	Phy_ParametersSharedSpectrumChAccess_r16_Ss_SINR_Meas_r16_Supported = 0
)

var phyParametersSharedSpectrumChAccessR16SsSINRMeasR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersSharedSpectrumChAccess_r16_Ss_SINR_Meas_r16_Supported},
}

const (
	Phy_ParametersSharedSpectrumChAccess_r16_Sp_CSI_ReportPUCCH_r16_Supported = 0
)

var phyParametersSharedSpectrumChAccessR16SpCSIReportPUCCHR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersSharedSpectrumChAccess_r16_Sp_CSI_ReportPUCCH_r16_Supported},
}

const (
	Phy_ParametersSharedSpectrumChAccess_r16_Sp_CSI_ReportPUSCH_r16_Supported = 0
)

var phyParametersSharedSpectrumChAccessR16SpCSIReportPUSCHR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersSharedSpectrumChAccess_r16_Sp_CSI_ReportPUSCH_r16_Supported},
}

const (
	Phy_ParametersSharedSpectrumChAccess_r16_DynamicSFI_r16_Supported = 0
)

var phyParametersSharedSpectrumChAccessR16DynamicSFIR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersSharedSpectrumChAccess_r16_DynamicSFI_r16_Supported},
}

const (
	Phy_ParametersSharedSpectrumChAccess_r16_Mux_SR_HARQ_ACK_PUCCH_r16_Supported = 0
)

var phyParametersSharedSpectrumChAccessR16MuxSRHARQACKPUCCHR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersSharedSpectrumChAccess_r16_Mux_SR_HARQ_ACK_PUCCH_r16_Supported},
}

const (
	Phy_ParametersSharedSpectrumChAccess_r16_Mux_SR_HARQ_ACK_CSI_PUCCH_MultiPerSlot_r16_Supported = 0
)

var phyParametersSharedSpectrumChAccessR16MuxSRHARQACKCSIPUCCHMultiPerSlotR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersSharedSpectrumChAccess_r16_Mux_SR_HARQ_ACK_CSI_PUCCH_MultiPerSlot_r16_Supported},
}

const (
	Phy_ParametersSharedSpectrumChAccess_r16_Mux_HARQ_ACK_PUSCH_DiffSymbol_r16_Supported = 0
)

var phyParametersSharedSpectrumChAccessR16MuxHARQACKPUSCHDiffSymbolR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersSharedSpectrumChAccess_r16_Mux_HARQ_ACK_PUSCH_DiffSymbol_r16_Supported},
}

const (
	Phy_ParametersSharedSpectrumChAccess_r16_Pucch_Repetition_F1_3_4_r16_Supported = 0
)

var phyParametersSharedSpectrumChAccessR16PucchRepetitionF134R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersSharedSpectrumChAccess_r16_Pucch_Repetition_F1_3_4_r16_Supported},
}

const (
	Phy_ParametersSharedSpectrumChAccess_r16_Type1_PUSCH_RepetitionMultiSlots_r16_Supported = 0
)

var phyParametersSharedSpectrumChAccessR16Type1PUSCHRepetitionMultiSlotsR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersSharedSpectrumChAccess_r16_Type1_PUSCH_RepetitionMultiSlots_r16_Supported},
}

const (
	Phy_ParametersSharedSpectrumChAccess_r16_Type2_PUSCH_RepetitionMultiSlots_r16_Supported = 0
)

var phyParametersSharedSpectrumChAccessR16Type2PUSCHRepetitionMultiSlotsR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersSharedSpectrumChAccess_r16_Type2_PUSCH_RepetitionMultiSlots_r16_Supported},
}

const (
	Phy_ParametersSharedSpectrumChAccess_r16_Pusch_RepetitionMultiSlots_r16_Supported = 0
)

var phyParametersSharedSpectrumChAccessR16PuschRepetitionMultiSlotsR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersSharedSpectrumChAccess_r16_Pusch_RepetitionMultiSlots_r16_Supported},
}

const (
	Phy_ParametersSharedSpectrumChAccess_r16_Pdsch_RepetitionMultiSlots_r16_Supported = 0
)

var phyParametersSharedSpectrumChAccessR16PdschRepetitionMultiSlotsR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersSharedSpectrumChAccess_r16_Pdsch_RepetitionMultiSlots_r16_Supported},
}

const (
	Phy_ParametersSharedSpectrumChAccess_r16_DownlinkSPS_r16_Supported = 0
)

var phyParametersSharedSpectrumChAccessR16DownlinkSPSR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersSharedSpectrumChAccess_r16_DownlinkSPS_r16_Supported},
}

const (
	Phy_ParametersSharedSpectrumChAccess_r16_ConfiguredUL_GrantType1_r16_Supported = 0
)

var phyParametersSharedSpectrumChAccessR16ConfiguredULGrantType1R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersSharedSpectrumChAccess_r16_ConfiguredUL_GrantType1_r16_Supported},
}

const (
	Phy_ParametersSharedSpectrumChAccess_r16_ConfiguredUL_GrantType2_r16_Supported = 0
)

var phyParametersSharedSpectrumChAccessR16ConfiguredULGrantType2R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersSharedSpectrumChAccess_r16_ConfiguredUL_GrantType2_r16_Supported},
}

const (
	Phy_ParametersSharedSpectrumChAccess_r16_Pre_EmptIndication_DL_r16_Supported = 0
)

var phyParametersSharedSpectrumChAccessR16PreEmptIndicationDLR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersSharedSpectrumChAccess_r16_Pre_EmptIndication_DL_r16_Supported},
}

var phyParametersSharedSpectrumChAccessR16MuxSRHARQACKCSIPUCCHOncePerSlotR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sameSymbol-r16", Optional: true},
		{Name: "diffSymbol-r16", Optional: true},
	},
}

const (
	Phy_ParametersSharedSpectrumChAccess_r16_Mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot_r16_SameSymbol_r16_Supported = 0
)

var phyParametersSharedSpectrumChAccessR16MuxSRHARQACKCSIPUCCHOncePerSlotR16SameSymbolR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersSharedSpectrumChAccess_r16_Mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot_r16_SameSymbol_r16_Supported},
}

const (
	Phy_ParametersSharedSpectrumChAccess_r16_Mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot_r16_DiffSymbol_r16_Supported = 0
)

var phyParametersSharedSpectrumChAccessR16MuxSRHARQACKCSIPUCCHOncePerSlotR16DiffSymbolR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersSharedSpectrumChAccess_r16_Mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot_r16_DiffSymbol_r16_Supported},
}

type Phy_ParametersSharedSpectrumChAccess_r16 struct {
	Ss_SINR_Meas_r16                          *int64
	Sp_CSI_ReportPUCCH_r16                    *int64
	Sp_CSI_ReportPUSCH_r16                    *int64
	DynamicSFI_r16                            *int64
	Mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot_r16 *struct {
		SameSymbol_r16 *int64
		DiffSymbol_r16 *int64
	}
	Mux_SR_HARQ_ACK_PUCCH_r16                  *int64
	Mux_SR_HARQ_ACK_CSI_PUCCH_MultiPerSlot_r16 *int64
	Mux_HARQ_ACK_PUSCH_DiffSymbol_r16          *int64
	Pucch_Repetition_F1_3_4_r16                *int64
	Type1_PUSCH_RepetitionMultiSlots_r16       *int64
	Type2_PUSCH_RepetitionMultiSlots_r16       *int64
	Pusch_RepetitionMultiSlots_r16             *int64
	Pdsch_RepetitionMultiSlots_r16             *int64
	DownlinkSPS_r16                            *int64
	ConfiguredUL_GrantType1_r16                *int64
	ConfiguredUL_GrantType2_r16                *int64
	Pre_EmptIndication_DL_r16                  *int64
}

func (ie *Phy_ParametersSharedSpectrumChAccess_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(phyParametersSharedSpectrumChAccessR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ss_SINR_Meas_r16 != nil, ie.Sp_CSI_ReportPUCCH_r16 != nil, ie.Sp_CSI_ReportPUSCH_r16 != nil, ie.DynamicSFI_r16 != nil, ie.Mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot_r16 != nil, ie.Mux_SR_HARQ_ACK_PUCCH_r16 != nil, ie.Mux_SR_HARQ_ACK_CSI_PUCCH_MultiPerSlot_r16 != nil, ie.Mux_HARQ_ACK_PUSCH_DiffSymbol_r16 != nil, ie.Pucch_Repetition_F1_3_4_r16 != nil, ie.Type1_PUSCH_RepetitionMultiSlots_r16 != nil, ie.Type2_PUSCH_RepetitionMultiSlots_r16 != nil, ie.Pusch_RepetitionMultiSlots_r16 != nil, ie.Pdsch_RepetitionMultiSlots_r16 != nil, ie.DownlinkSPS_r16 != nil, ie.ConfiguredUL_GrantType1_r16 != nil, ie.ConfiguredUL_GrantType2_r16 != nil, ie.Pre_EmptIndication_DL_r16 != nil}); err != nil {
		return err
	}

	// 3. ss-SINR-Meas-r16: enumerated
	{
		if ie.Ss_SINR_Meas_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Ss_SINR_Meas_r16, phyParametersSharedSpectrumChAccessR16SsSINRMeasR16Constraints); err != nil {
				return err
			}
		}
	}

	// 4. sp-CSI-ReportPUCCH-r16: enumerated
	{
		if ie.Sp_CSI_ReportPUCCH_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Sp_CSI_ReportPUCCH_r16, phyParametersSharedSpectrumChAccessR16SpCSIReportPUCCHR16Constraints); err != nil {
				return err
			}
		}
	}

	// 5. sp-CSI-ReportPUSCH-r16: enumerated
	{
		if ie.Sp_CSI_ReportPUSCH_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Sp_CSI_ReportPUSCH_r16, phyParametersSharedSpectrumChAccessR16SpCSIReportPUSCHR16Constraints); err != nil {
				return err
			}
		}
	}

	// 6. dynamicSFI-r16: enumerated
	{
		if ie.DynamicSFI_r16 != nil {
			if err := e.EncodeEnumerated(*ie.DynamicSFI_r16, phyParametersSharedSpectrumChAccessR16DynamicSFIR16Constraints); err != nil {
				return err
			}
		}
	}

	// 7. mux-SR-HARQ-ACK-CSI-PUCCH-OncePerSlot-r16: sequence
	{
		if ie.Mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot_r16 != nil {
			c := ie.Mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot_r16
			phyParametersSharedSpectrumChAccessR16MuxSRHARQACKCSIPUCCHOncePerSlotR16Seq := e.NewSequenceEncoder(phyParametersSharedSpectrumChAccessR16MuxSRHARQACKCSIPUCCHOncePerSlotR16Constraints)
			if err := phyParametersSharedSpectrumChAccessR16MuxSRHARQACKCSIPUCCHOncePerSlotR16Seq.EncodePreamble([]bool{c.SameSymbol_r16 != nil, c.DiffSymbol_r16 != nil}); err != nil {
				return err
			}
			if c.SameSymbol_r16 != nil {
				if err := e.EncodeEnumerated((*c.SameSymbol_r16), phyParametersSharedSpectrumChAccessR16MuxSRHARQACKCSIPUCCHOncePerSlotR16SameSymbolR16Constraints); err != nil {
					return err
				}
			}
			if c.DiffSymbol_r16 != nil {
				if err := e.EncodeEnumerated((*c.DiffSymbol_r16), phyParametersSharedSpectrumChAccessR16MuxSRHARQACKCSIPUCCHOncePerSlotR16DiffSymbolR16Constraints); err != nil {
					return err
				}
			}
		}
	}

	// 8. mux-SR-HARQ-ACK-PUCCH-r16: enumerated
	{
		if ie.Mux_SR_HARQ_ACK_PUCCH_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Mux_SR_HARQ_ACK_PUCCH_r16, phyParametersSharedSpectrumChAccessR16MuxSRHARQACKPUCCHR16Constraints); err != nil {
				return err
			}
		}
	}

	// 9. mux-SR-HARQ-ACK-CSI-PUCCH-MultiPerSlot-r16: enumerated
	{
		if ie.Mux_SR_HARQ_ACK_CSI_PUCCH_MultiPerSlot_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Mux_SR_HARQ_ACK_CSI_PUCCH_MultiPerSlot_r16, phyParametersSharedSpectrumChAccessR16MuxSRHARQACKCSIPUCCHMultiPerSlotR16Constraints); err != nil {
				return err
			}
		}
	}

	// 10. mux-HARQ-ACK-PUSCH-DiffSymbol-r16: enumerated
	{
		if ie.Mux_HARQ_ACK_PUSCH_DiffSymbol_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Mux_HARQ_ACK_PUSCH_DiffSymbol_r16, phyParametersSharedSpectrumChAccessR16MuxHARQACKPUSCHDiffSymbolR16Constraints); err != nil {
				return err
			}
		}
	}

	// 11. pucch-Repetition-F1-3-4-r16: enumerated
	{
		if ie.Pucch_Repetition_F1_3_4_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Pucch_Repetition_F1_3_4_r16, phyParametersSharedSpectrumChAccessR16PucchRepetitionF134R16Constraints); err != nil {
				return err
			}
		}
	}

	// 12. type1-PUSCH-RepetitionMultiSlots-r16: enumerated
	{
		if ie.Type1_PUSCH_RepetitionMultiSlots_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Type1_PUSCH_RepetitionMultiSlots_r16, phyParametersSharedSpectrumChAccessR16Type1PUSCHRepetitionMultiSlotsR16Constraints); err != nil {
				return err
			}
		}
	}

	// 13. type2-PUSCH-RepetitionMultiSlots-r16: enumerated
	{
		if ie.Type2_PUSCH_RepetitionMultiSlots_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Type2_PUSCH_RepetitionMultiSlots_r16, phyParametersSharedSpectrumChAccessR16Type2PUSCHRepetitionMultiSlotsR16Constraints); err != nil {
				return err
			}
		}
	}

	// 14. pusch-RepetitionMultiSlots-r16: enumerated
	{
		if ie.Pusch_RepetitionMultiSlots_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Pusch_RepetitionMultiSlots_r16, phyParametersSharedSpectrumChAccessR16PuschRepetitionMultiSlotsR16Constraints); err != nil {
				return err
			}
		}
	}

	// 15. pdsch-RepetitionMultiSlots-r16: enumerated
	{
		if ie.Pdsch_RepetitionMultiSlots_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Pdsch_RepetitionMultiSlots_r16, phyParametersSharedSpectrumChAccessR16PdschRepetitionMultiSlotsR16Constraints); err != nil {
				return err
			}
		}
	}

	// 16. downlinkSPS-r16: enumerated
	{
		if ie.DownlinkSPS_r16 != nil {
			if err := e.EncodeEnumerated(*ie.DownlinkSPS_r16, phyParametersSharedSpectrumChAccessR16DownlinkSPSR16Constraints); err != nil {
				return err
			}
		}
	}

	// 17. configuredUL-GrantType1-r16: enumerated
	{
		if ie.ConfiguredUL_GrantType1_r16 != nil {
			if err := e.EncodeEnumerated(*ie.ConfiguredUL_GrantType1_r16, phyParametersSharedSpectrumChAccessR16ConfiguredULGrantType1R16Constraints); err != nil {
				return err
			}
		}
	}

	// 18. configuredUL-GrantType2-r16: enumerated
	{
		if ie.ConfiguredUL_GrantType2_r16 != nil {
			if err := e.EncodeEnumerated(*ie.ConfiguredUL_GrantType2_r16, phyParametersSharedSpectrumChAccessR16ConfiguredULGrantType2R16Constraints); err != nil {
				return err
			}
		}
	}

	// 19. pre-EmptIndication-DL-r16: enumerated
	{
		if ie.Pre_EmptIndication_DL_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Pre_EmptIndication_DL_r16, phyParametersSharedSpectrumChAccessR16PreEmptIndicationDLR16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *Phy_ParametersSharedSpectrumChAccess_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(phyParametersSharedSpectrumChAccessR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. ss-SINR-Meas-r16: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(phyParametersSharedSpectrumChAccessR16SsSINRMeasR16Constraints)
			if err != nil {
				return err
			}
			ie.Ss_SINR_Meas_r16 = &idx
		}
	}

	// 4. sp-CSI-ReportPUCCH-r16: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(phyParametersSharedSpectrumChAccessR16SpCSIReportPUCCHR16Constraints)
			if err != nil {
				return err
			}
			ie.Sp_CSI_ReportPUCCH_r16 = &idx
		}
	}

	// 5. sp-CSI-ReportPUSCH-r16: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(phyParametersSharedSpectrumChAccessR16SpCSIReportPUSCHR16Constraints)
			if err != nil {
				return err
			}
			ie.Sp_CSI_ReportPUSCH_r16 = &idx
		}
	}

	// 6. dynamicSFI-r16: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(phyParametersSharedSpectrumChAccessR16DynamicSFIR16Constraints)
			if err != nil {
				return err
			}
			ie.DynamicSFI_r16 = &idx
		}
	}

	// 7. mux-SR-HARQ-ACK-CSI-PUCCH-OncePerSlot-r16: sequence
	{
		if seq.IsComponentPresent(4) {
			ie.Mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot_r16 = &struct {
				SameSymbol_r16 *int64
				DiffSymbol_r16 *int64
			}{}
			c := ie.Mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot_r16
			phyParametersSharedSpectrumChAccessR16MuxSRHARQACKCSIPUCCHOncePerSlotR16Seq := d.NewSequenceDecoder(phyParametersSharedSpectrumChAccessR16MuxSRHARQACKCSIPUCCHOncePerSlotR16Constraints)
			if err := phyParametersSharedSpectrumChAccessR16MuxSRHARQACKCSIPUCCHOncePerSlotR16Seq.DecodePreamble(); err != nil {
				return err
			}
			if phyParametersSharedSpectrumChAccessR16MuxSRHARQACKCSIPUCCHOncePerSlotR16Seq.IsComponentPresent(0) {
				c.SameSymbol_r16 = new(int64)
				v, err := d.DecodeEnumerated(phyParametersSharedSpectrumChAccessR16MuxSRHARQACKCSIPUCCHOncePerSlotR16SameSymbolR16Constraints)
				if err != nil {
					return err
				}
				(*c.SameSymbol_r16) = v
			}
			if phyParametersSharedSpectrumChAccessR16MuxSRHARQACKCSIPUCCHOncePerSlotR16Seq.IsComponentPresent(1) {
				c.DiffSymbol_r16 = new(int64)
				v, err := d.DecodeEnumerated(phyParametersSharedSpectrumChAccessR16MuxSRHARQACKCSIPUCCHOncePerSlotR16DiffSymbolR16Constraints)
				if err != nil {
					return err
				}
				(*c.DiffSymbol_r16) = v
			}
		}
	}

	// 8. mux-SR-HARQ-ACK-PUCCH-r16: enumerated
	{
		if seq.IsComponentPresent(5) {
			idx, err := d.DecodeEnumerated(phyParametersSharedSpectrumChAccessR16MuxSRHARQACKPUCCHR16Constraints)
			if err != nil {
				return err
			}
			ie.Mux_SR_HARQ_ACK_PUCCH_r16 = &idx
		}
	}

	// 9. mux-SR-HARQ-ACK-CSI-PUCCH-MultiPerSlot-r16: enumerated
	{
		if seq.IsComponentPresent(6) {
			idx, err := d.DecodeEnumerated(phyParametersSharedSpectrumChAccessR16MuxSRHARQACKCSIPUCCHMultiPerSlotR16Constraints)
			if err != nil {
				return err
			}
			ie.Mux_SR_HARQ_ACK_CSI_PUCCH_MultiPerSlot_r16 = &idx
		}
	}

	// 10. mux-HARQ-ACK-PUSCH-DiffSymbol-r16: enumerated
	{
		if seq.IsComponentPresent(7) {
			idx, err := d.DecodeEnumerated(phyParametersSharedSpectrumChAccessR16MuxHARQACKPUSCHDiffSymbolR16Constraints)
			if err != nil {
				return err
			}
			ie.Mux_HARQ_ACK_PUSCH_DiffSymbol_r16 = &idx
		}
	}

	// 11. pucch-Repetition-F1-3-4-r16: enumerated
	{
		if seq.IsComponentPresent(8) {
			idx, err := d.DecodeEnumerated(phyParametersSharedSpectrumChAccessR16PucchRepetitionF134R16Constraints)
			if err != nil {
				return err
			}
			ie.Pucch_Repetition_F1_3_4_r16 = &idx
		}
	}

	// 12. type1-PUSCH-RepetitionMultiSlots-r16: enumerated
	{
		if seq.IsComponentPresent(9) {
			idx, err := d.DecodeEnumerated(phyParametersSharedSpectrumChAccessR16Type1PUSCHRepetitionMultiSlotsR16Constraints)
			if err != nil {
				return err
			}
			ie.Type1_PUSCH_RepetitionMultiSlots_r16 = &idx
		}
	}

	// 13. type2-PUSCH-RepetitionMultiSlots-r16: enumerated
	{
		if seq.IsComponentPresent(10) {
			idx, err := d.DecodeEnumerated(phyParametersSharedSpectrumChAccessR16Type2PUSCHRepetitionMultiSlotsR16Constraints)
			if err != nil {
				return err
			}
			ie.Type2_PUSCH_RepetitionMultiSlots_r16 = &idx
		}
	}

	// 14. pusch-RepetitionMultiSlots-r16: enumerated
	{
		if seq.IsComponentPresent(11) {
			idx, err := d.DecodeEnumerated(phyParametersSharedSpectrumChAccessR16PuschRepetitionMultiSlotsR16Constraints)
			if err != nil {
				return err
			}
			ie.Pusch_RepetitionMultiSlots_r16 = &idx
		}
	}

	// 15. pdsch-RepetitionMultiSlots-r16: enumerated
	{
		if seq.IsComponentPresent(12) {
			idx, err := d.DecodeEnumerated(phyParametersSharedSpectrumChAccessR16PdschRepetitionMultiSlotsR16Constraints)
			if err != nil {
				return err
			}
			ie.Pdsch_RepetitionMultiSlots_r16 = &idx
		}
	}

	// 16. downlinkSPS-r16: enumerated
	{
		if seq.IsComponentPresent(13) {
			idx, err := d.DecodeEnumerated(phyParametersSharedSpectrumChAccessR16DownlinkSPSR16Constraints)
			if err != nil {
				return err
			}
			ie.DownlinkSPS_r16 = &idx
		}
	}

	// 17. configuredUL-GrantType1-r16: enumerated
	{
		if seq.IsComponentPresent(14) {
			idx, err := d.DecodeEnumerated(phyParametersSharedSpectrumChAccessR16ConfiguredULGrantType1R16Constraints)
			if err != nil {
				return err
			}
			ie.ConfiguredUL_GrantType1_r16 = &idx
		}
	}

	// 18. configuredUL-GrantType2-r16: enumerated
	{
		if seq.IsComponentPresent(15) {
			idx, err := d.DecodeEnumerated(phyParametersSharedSpectrumChAccessR16ConfiguredULGrantType2R16Constraints)
			if err != nil {
				return err
			}
			ie.ConfiguredUL_GrantType2_r16 = &idx
		}
	}

	// 19. pre-EmptIndication-DL-r16: enumerated
	{
		if seq.IsComponentPresent(16) {
			idx, err := d.DecodeEnumerated(phyParametersSharedSpectrumChAccessR16PreEmptIndicationDLR16Constraints)
			if err != nil {
				return err
			}
			ie.Pre_EmptIndication_DL_r16 = &idx
		}
	}

	return nil
}
