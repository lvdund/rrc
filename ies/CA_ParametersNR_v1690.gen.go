// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CA-ParametersNR-v1690 (line 17417).

var cAParametersNRV1690Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "csi-ReportingCrossPUCCH-Grp-r16", Optional: true},
	},
}

var cAParametersNRV1690CsiReportingCrossPUCCHGrpR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "computationTimeForA-CSI-r16"},
		{Name: "additionalSymbols-r16", Optional: true},
		{Name: "sp-CSI-ReportingOnPUCCH-r16", Optional: true},
		{Name: "sp-CSI-ReportingOnPUSCH-r16", Optional: true},
		{Name: "carrierTypePairList-r16"},
	},
}

const (
	CA_ParametersNR_v1690_Csi_ReportingCrossPUCCH_Grp_r16_ComputationTimeForA_CSI_r16_SameAsNoCross = 0
	CA_ParametersNR_v1690_Csi_ReportingCrossPUCCH_Grp_r16_ComputationTimeForA_CSI_r16_Relaxed       = 1
)

var cAParametersNRV1690CsiReportingCrossPUCCHGrpR16ComputationTimeForACSIR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1690_Csi_ReportingCrossPUCCH_Grp_r16_ComputationTimeForA_CSI_r16_SameAsNoCross, CA_ParametersNR_v1690_Csi_ReportingCrossPUCCH_Grp_r16_ComputationTimeForA_CSI_r16_Relaxed},
}

var cAParametersNRV1690CsiReportingCrossPUCCHGrpR16AdditionalSymbolsR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-15kHz-additionalSymbols-r16", Optional: true},
		{Name: "scs-30kHz-additionalSymbols-r16", Optional: true},
		{Name: "scs-60kHz-additionalSymbols-r16", Optional: true},
		{Name: "scs-120kHz-additionalSymbols-r16", Optional: true},
	},
}

const (
	CA_ParametersNR_v1690_Csi_ReportingCrossPUCCH_Grp_r16_AdditionalSymbols_r16_Scs_15kHz_AdditionalSymbols_r16_S14 = 0
	CA_ParametersNR_v1690_Csi_ReportingCrossPUCCH_Grp_r16_AdditionalSymbols_r16_Scs_15kHz_AdditionalSymbols_r16_S28 = 1
)

var cAParametersNRV1690CsiReportingCrossPUCCHGrpR16AdditionalSymbolsR16Scs15kHzAdditionalSymbolsR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1690_Csi_ReportingCrossPUCCH_Grp_r16_AdditionalSymbols_r16_Scs_15kHz_AdditionalSymbols_r16_S14, CA_ParametersNR_v1690_Csi_ReportingCrossPUCCH_Grp_r16_AdditionalSymbols_r16_Scs_15kHz_AdditionalSymbols_r16_S28},
}

const (
	CA_ParametersNR_v1690_Csi_ReportingCrossPUCCH_Grp_r16_AdditionalSymbols_r16_Scs_30kHz_AdditionalSymbols_r16_S14 = 0
	CA_ParametersNR_v1690_Csi_ReportingCrossPUCCH_Grp_r16_AdditionalSymbols_r16_Scs_30kHz_AdditionalSymbols_r16_S28 = 1
)

var cAParametersNRV1690CsiReportingCrossPUCCHGrpR16AdditionalSymbolsR16Scs30kHzAdditionalSymbolsR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1690_Csi_ReportingCrossPUCCH_Grp_r16_AdditionalSymbols_r16_Scs_30kHz_AdditionalSymbols_r16_S14, CA_ParametersNR_v1690_Csi_ReportingCrossPUCCH_Grp_r16_AdditionalSymbols_r16_Scs_30kHz_AdditionalSymbols_r16_S28},
}

const (
	CA_ParametersNR_v1690_Csi_ReportingCrossPUCCH_Grp_r16_AdditionalSymbols_r16_Scs_60kHz_AdditionalSymbols_r16_S14 = 0
	CA_ParametersNR_v1690_Csi_ReportingCrossPUCCH_Grp_r16_AdditionalSymbols_r16_Scs_60kHz_AdditionalSymbols_r16_S28 = 1
	CA_ParametersNR_v1690_Csi_ReportingCrossPUCCH_Grp_r16_AdditionalSymbols_r16_Scs_60kHz_AdditionalSymbols_r16_S56 = 2
)

var cAParametersNRV1690CsiReportingCrossPUCCHGrpR16AdditionalSymbolsR16Scs60kHzAdditionalSymbolsR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1690_Csi_ReportingCrossPUCCH_Grp_r16_AdditionalSymbols_r16_Scs_60kHz_AdditionalSymbols_r16_S14, CA_ParametersNR_v1690_Csi_ReportingCrossPUCCH_Grp_r16_AdditionalSymbols_r16_Scs_60kHz_AdditionalSymbols_r16_S28, CA_ParametersNR_v1690_Csi_ReportingCrossPUCCH_Grp_r16_AdditionalSymbols_r16_Scs_60kHz_AdditionalSymbols_r16_S56},
}

const (
	CA_ParametersNR_v1690_Csi_ReportingCrossPUCCH_Grp_r16_AdditionalSymbols_r16_Scs_120kHz_AdditionalSymbols_r16_S14 = 0
	CA_ParametersNR_v1690_Csi_ReportingCrossPUCCH_Grp_r16_AdditionalSymbols_r16_Scs_120kHz_AdditionalSymbols_r16_S28 = 1
	CA_ParametersNR_v1690_Csi_ReportingCrossPUCCH_Grp_r16_AdditionalSymbols_r16_Scs_120kHz_AdditionalSymbols_r16_S56 = 2
)

var cAParametersNRV1690CsiReportingCrossPUCCHGrpR16AdditionalSymbolsR16Scs120kHzAdditionalSymbolsR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1690_Csi_ReportingCrossPUCCH_Grp_r16_AdditionalSymbols_r16_Scs_120kHz_AdditionalSymbols_r16_S14, CA_ParametersNR_v1690_Csi_ReportingCrossPUCCH_Grp_r16_AdditionalSymbols_r16_Scs_120kHz_AdditionalSymbols_r16_S28, CA_ParametersNR_v1690_Csi_ReportingCrossPUCCH_Grp_r16_AdditionalSymbols_r16_Scs_120kHz_AdditionalSymbols_r16_S56},
}

const (
	CA_ParametersNR_v1690_Csi_ReportingCrossPUCCH_Grp_r16_Sp_CSI_ReportingOnPUCCH_r16_Supported = 0
)

var cAParametersNRV1690CsiReportingCrossPUCCHGrpR16SpCSIReportingOnPUCCHR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1690_Csi_ReportingCrossPUCCH_Grp_r16_Sp_CSI_ReportingOnPUCCH_r16_Supported},
}

const (
	CA_ParametersNR_v1690_Csi_ReportingCrossPUCCH_Grp_r16_Sp_CSI_ReportingOnPUSCH_r16_Supported = 0
)

var cAParametersNRV1690CsiReportingCrossPUCCHGrpR16SpCSIReportingOnPUSCHR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1690_Csi_ReportingCrossPUCCH_Grp_r16_Sp_CSI_ReportingOnPUSCH_r16_Supported},
}

var cAParametersNRV1690CsiReportingCrossPUCCHGrpR16CarrierTypePairListR16Constraints = per.SizeRange(1, common.MaxCarrierTypePairList_r16)

type CA_ParametersNR_v1690 struct {
	Csi_ReportingCrossPUCCH_Grp_r16 *struct {
		ComputationTimeForA_CSI_r16 int64
		AdditionalSymbols_r16       *struct {
			Scs_15kHz_AdditionalSymbols_r16  *int64
			Scs_30kHz_AdditionalSymbols_r16  *int64
			Scs_60kHz_AdditionalSymbols_r16  *int64
			Scs_120kHz_AdditionalSymbols_r16 *int64
		}
		Sp_CSI_ReportingOnPUCCH_r16 *int64
		Sp_CSI_ReportingOnPUSCH_r16 *int64
		CarrierTypePairList_r16     []CarrierTypePair_r16
	}
}

func (ie *CA_ParametersNR_v1690) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cAParametersNRV1690Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Csi_ReportingCrossPUCCH_Grp_r16 != nil}); err != nil {
		return err
	}

	// 2. csi-ReportingCrossPUCCH-Grp-r16: sequence
	{
		if ie.Csi_ReportingCrossPUCCH_Grp_r16 != nil {
			c := ie.Csi_ReportingCrossPUCCH_Grp_r16
			cAParametersNRV1690CsiReportingCrossPUCCHGrpR16Seq := e.NewSequenceEncoder(cAParametersNRV1690CsiReportingCrossPUCCHGrpR16Constraints)
			if err := cAParametersNRV1690CsiReportingCrossPUCCHGrpR16Seq.EncodePreamble([]bool{c.AdditionalSymbols_r16 != nil, c.Sp_CSI_ReportingOnPUCCH_r16 != nil, c.Sp_CSI_ReportingOnPUSCH_r16 != nil}); err != nil {
				return err
			}
			if err := e.EncodeEnumerated(c.ComputationTimeForA_CSI_r16, cAParametersNRV1690CsiReportingCrossPUCCHGrpR16ComputationTimeForACSIR16Constraints); err != nil {
				return err
			}
			if c.AdditionalSymbols_r16 != nil {
				c := (*c.AdditionalSymbols_r16)
				cAParametersNRV1690CsiReportingCrossPUCCHGrpR16AdditionalSymbolsR16Seq := e.NewSequenceEncoder(cAParametersNRV1690CsiReportingCrossPUCCHGrpR16AdditionalSymbolsR16Constraints)
				if err := cAParametersNRV1690CsiReportingCrossPUCCHGrpR16AdditionalSymbolsR16Seq.EncodePreamble([]bool{c.Scs_15kHz_AdditionalSymbols_r16 != nil, c.Scs_30kHz_AdditionalSymbols_r16 != nil, c.Scs_60kHz_AdditionalSymbols_r16 != nil, c.Scs_120kHz_AdditionalSymbols_r16 != nil}); err != nil {
					return err
				}
				if c.Scs_15kHz_AdditionalSymbols_r16 != nil {
					if err := e.EncodeEnumerated((*c.Scs_15kHz_AdditionalSymbols_r16), cAParametersNRV1690CsiReportingCrossPUCCHGrpR16AdditionalSymbolsR16Scs15kHzAdditionalSymbolsR16Constraints); err != nil {
						return err
					}
				}
				if c.Scs_30kHz_AdditionalSymbols_r16 != nil {
					if err := e.EncodeEnumerated((*c.Scs_30kHz_AdditionalSymbols_r16), cAParametersNRV1690CsiReportingCrossPUCCHGrpR16AdditionalSymbolsR16Scs30kHzAdditionalSymbolsR16Constraints); err != nil {
						return err
					}
				}
				if c.Scs_60kHz_AdditionalSymbols_r16 != nil {
					if err := e.EncodeEnumerated((*c.Scs_60kHz_AdditionalSymbols_r16), cAParametersNRV1690CsiReportingCrossPUCCHGrpR16AdditionalSymbolsR16Scs60kHzAdditionalSymbolsR16Constraints); err != nil {
						return err
					}
				}
				if c.Scs_120kHz_AdditionalSymbols_r16 != nil {
					if err := e.EncodeEnumerated((*c.Scs_120kHz_AdditionalSymbols_r16), cAParametersNRV1690CsiReportingCrossPUCCHGrpR16AdditionalSymbolsR16Scs120kHzAdditionalSymbolsR16Constraints); err != nil {
						return err
					}
				}
			}
			if c.Sp_CSI_ReportingOnPUCCH_r16 != nil {
				if err := e.EncodeEnumerated((*c.Sp_CSI_ReportingOnPUCCH_r16), cAParametersNRV1690CsiReportingCrossPUCCHGrpR16SpCSIReportingOnPUCCHR16Constraints); err != nil {
					return err
				}
			}
			if c.Sp_CSI_ReportingOnPUSCH_r16 != nil {
				if err := e.EncodeEnumerated((*c.Sp_CSI_ReportingOnPUSCH_r16), cAParametersNRV1690CsiReportingCrossPUCCHGrpR16SpCSIReportingOnPUSCHR16Constraints); err != nil {
					return err
				}
			}
			{
				seqOf := e.NewSequenceOfEncoder(cAParametersNRV1690CsiReportingCrossPUCCHGrpR16CarrierTypePairListR16Constraints)
				if err := seqOf.EncodeLength(int64(len(c.CarrierTypePairList_r16))); err != nil {
					return err
				}
				for i := range c.CarrierTypePairList_r16 {
					if err := c.CarrierTypePairList_r16[i].Encode(e); err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}

func (ie *CA_ParametersNR_v1690) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cAParametersNRV1690Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. csi-ReportingCrossPUCCH-Grp-r16: sequence
	{
		if seq.IsComponentPresent(0) {
			ie.Csi_ReportingCrossPUCCH_Grp_r16 = &struct {
				ComputationTimeForA_CSI_r16 int64
				AdditionalSymbols_r16       *struct {
					Scs_15kHz_AdditionalSymbols_r16  *int64
					Scs_30kHz_AdditionalSymbols_r16  *int64
					Scs_60kHz_AdditionalSymbols_r16  *int64
					Scs_120kHz_AdditionalSymbols_r16 *int64
				}
				Sp_CSI_ReportingOnPUCCH_r16 *int64
				Sp_CSI_ReportingOnPUSCH_r16 *int64
				CarrierTypePairList_r16     []CarrierTypePair_r16
			}{}
			c := ie.Csi_ReportingCrossPUCCH_Grp_r16
			cAParametersNRV1690CsiReportingCrossPUCCHGrpR16Seq := d.NewSequenceDecoder(cAParametersNRV1690CsiReportingCrossPUCCHGrpR16Constraints)
			if err := cAParametersNRV1690CsiReportingCrossPUCCHGrpR16Seq.DecodePreamble(); err != nil {
				return err
			}
			{
				v, err := d.DecodeEnumerated(cAParametersNRV1690CsiReportingCrossPUCCHGrpR16ComputationTimeForACSIR16Constraints)
				if err != nil {
					return err
				}
				c.ComputationTimeForA_CSI_r16 = v
			}
			if cAParametersNRV1690CsiReportingCrossPUCCHGrpR16Seq.IsComponentPresent(1) {
				c.AdditionalSymbols_r16 = &struct {
					Scs_15kHz_AdditionalSymbols_r16  *int64
					Scs_30kHz_AdditionalSymbols_r16  *int64
					Scs_60kHz_AdditionalSymbols_r16  *int64
					Scs_120kHz_AdditionalSymbols_r16 *int64
				}{}
				c := (*c.AdditionalSymbols_r16)
				cAParametersNRV1690CsiReportingCrossPUCCHGrpR16AdditionalSymbolsR16Seq := d.NewSequenceDecoder(cAParametersNRV1690CsiReportingCrossPUCCHGrpR16AdditionalSymbolsR16Constraints)
				if err := cAParametersNRV1690CsiReportingCrossPUCCHGrpR16AdditionalSymbolsR16Seq.DecodePreamble(); err != nil {
					return err
				}
				if cAParametersNRV1690CsiReportingCrossPUCCHGrpR16AdditionalSymbolsR16Seq.IsComponentPresent(0) {
					c.Scs_15kHz_AdditionalSymbols_r16 = new(int64)
					v, err := d.DecodeEnumerated(cAParametersNRV1690CsiReportingCrossPUCCHGrpR16AdditionalSymbolsR16Scs15kHzAdditionalSymbolsR16Constraints)
					if err != nil {
						return err
					}
					(*c.Scs_15kHz_AdditionalSymbols_r16) = v
				}
				if cAParametersNRV1690CsiReportingCrossPUCCHGrpR16AdditionalSymbolsR16Seq.IsComponentPresent(1) {
					c.Scs_30kHz_AdditionalSymbols_r16 = new(int64)
					v, err := d.DecodeEnumerated(cAParametersNRV1690CsiReportingCrossPUCCHGrpR16AdditionalSymbolsR16Scs30kHzAdditionalSymbolsR16Constraints)
					if err != nil {
						return err
					}
					(*c.Scs_30kHz_AdditionalSymbols_r16) = v
				}
				if cAParametersNRV1690CsiReportingCrossPUCCHGrpR16AdditionalSymbolsR16Seq.IsComponentPresent(2) {
					c.Scs_60kHz_AdditionalSymbols_r16 = new(int64)
					v, err := d.DecodeEnumerated(cAParametersNRV1690CsiReportingCrossPUCCHGrpR16AdditionalSymbolsR16Scs60kHzAdditionalSymbolsR16Constraints)
					if err != nil {
						return err
					}
					(*c.Scs_60kHz_AdditionalSymbols_r16) = v
				}
				if cAParametersNRV1690CsiReportingCrossPUCCHGrpR16AdditionalSymbolsR16Seq.IsComponentPresent(3) {
					c.Scs_120kHz_AdditionalSymbols_r16 = new(int64)
					v, err := d.DecodeEnumerated(cAParametersNRV1690CsiReportingCrossPUCCHGrpR16AdditionalSymbolsR16Scs120kHzAdditionalSymbolsR16Constraints)
					if err != nil {
						return err
					}
					(*c.Scs_120kHz_AdditionalSymbols_r16) = v
				}
			}
			if cAParametersNRV1690CsiReportingCrossPUCCHGrpR16Seq.IsComponentPresent(2) {
				c.Sp_CSI_ReportingOnPUCCH_r16 = new(int64)
				v, err := d.DecodeEnumerated(cAParametersNRV1690CsiReportingCrossPUCCHGrpR16SpCSIReportingOnPUCCHR16Constraints)
				if err != nil {
					return err
				}
				(*c.Sp_CSI_ReportingOnPUCCH_r16) = v
			}
			if cAParametersNRV1690CsiReportingCrossPUCCHGrpR16Seq.IsComponentPresent(3) {
				c.Sp_CSI_ReportingOnPUSCH_r16 = new(int64)
				v, err := d.DecodeEnumerated(cAParametersNRV1690CsiReportingCrossPUCCHGrpR16SpCSIReportingOnPUSCHR16Constraints)
				if err != nil {
					return err
				}
				(*c.Sp_CSI_ReportingOnPUSCH_r16) = v
			}
			{
				seqOf := d.NewSequenceOfDecoder(cAParametersNRV1690CsiReportingCrossPUCCHGrpR16CarrierTypePairListR16Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.CarrierTypePairList_r16 = make([]CarrierTypePair_r16, n)
				for i := int64(0); i < n; i++ {
					if err := c.CarrierTypePairList_r16[i].Decode(d); err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}
