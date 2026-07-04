// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PHR-Config (line 11554).

var pHRConfigConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "phr-PeriodicTimer"},
		{Name: "phr-ProhibitTimer"},
		{Name: "phr-Tx-PowerFactorChange"},
		{Name: "multiplePHR"},
		{Name: "dummy"},
		{Name: "phr-Type2OtherCell"},
		{Name: "phr-ModeOtherCG"},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
		{Name: "extension-addition-2", Optional: true},
	},
}

const (
	PHR_Config_Phr_PeriodicTimer_Sf10     = 0
	PHR_Config_Phr_PeriodicTimer_Sf20     = 1
	PHR_Config_Phr_PeriodicTimer_Sf50     = 2
	PHR_Config_Phr_PeriodicTimer_Sf100    = 3
	PHR_Config_Phr_PeriodicTimer_Sf200    = 4
	PHR_Config_Phr_PeriodicTimer_Sf500    = 5
	PHR_Config_Phr_PeriodicTimer_Sf1000   = 6
	PHR_Config_Phr_PeriodicTimer_Infinity = 7
)

var pHRConfigPhrPeriodicTimerConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PHR_Config_Phr_PeriodicTimer_Sf10, PHR_Config_Phr_PeriodicTimer_Sf20, PHR_Config_Phr_PeriodicTimer_Sf50, PHR_Config_Phr_PeriodicTimer_Sf100, PHR_Config_Phr_PeriodicTimer_Sf200, PHR_Config_Phr_PeriodicTimer_Sf500, PHR_Config_Phr_PeriodicTimer_Sf1000, PHR_Config_Phr_PeriodicTimer_Infinity},
}

const (
	PHR_Config_Phr_ProhibitTimer_Sf0    = 0
	PHR_Config_Phr_ProhibitTimer_Sf10   = 1
	PHR_Config_Phr_ProhibitTimer_Sf20   = 2
	PHR_Config_Phr_ProhibitTimer_Sf50   = 3
	PHR_Config_Phr_ProhibitTimer_Sf100  = 4
	PHR_Config_Phr_ProhibitTimer_Sf200  = 5
	PHR_Config_Phr_ProhibitTimer_Sf500  = 6
	PHR_Config_Phr_ProhibitTimer_Sf1000 = 7
)

var pHRConfigPhrProhibitTimerConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PHR_Config_Phr_ProhibitTimer_Sf0, PHR_Config_Phr_ProhibitTimer_Sf10, PHR_Config_Phr_ProhibitTimer_Sf20, PHR_Config_Phr_ProhibitTimer_Sf50, PHR_Config_Phr_ProhibitTimer_Sf100, PHR_Config_Phr_ProhibitTimer_Sf200, PHR_Config_Phr_ProhibitTimer_Sf500, PHR_Config_Phr_ProhibitTimer_Sf1000},
}

const (
	PHR_Config_Phr_Tx_PowerFactorChange_DB1      = 0
	PHR_Config_Phr_Tx_PowerFactorChange_DB3      = 1
	PHR_Config_Phr_Tx_PowerFactorChange_DB6      = 2
	PHR_Config_Phr_Tx_PowerFactorChange_Infinity = 3
)

var pHRConfigPhrTxPowerFactorChangeConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PHR_Config_Phr_Tx_PowerFactorChange_DB1, PHR_Config_Phr_Tx_PowerFactorChange_DB3, PHR_Config_Phr_Tx_PowerFactorChange_DB6, PHR_Config_Phr_Tx_PowerFactorChange_Infinity},
}

const (
	PHR_Config_Phr_ModeOtherCG_Real    = 0
	PHR_Config_Phr_ModeOtherCG_Virtual = 1
)

var pHRConfigPhrModeOtherCGConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PHR_Config_Phr_ModeOtherCG_Real, PHR_Config_Phr_ModeOtherCG_Virtual},
}

var pHRConfigExtMpeReportingFR2R16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PHR_Config_Ext_Mpe_Reporting_FR2_r16_Release = 0
	PHR_Config_Ext_Mpe_Reporting_FR2_r16_Setup   = 1
)

var pHRConfigExtMpeReportingFR2R17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PHR_Config_Ext_Mpe_Reporting_FR2_r17_Release = 0
	PHR_Config_Ext_Mpe_Reporting_FR2_r17_Setup   = 1
)

const (
	PHR_Config_Ext_TwoPHRMode_r17_Enabled = 0
)

var pHRConfigExtTwoPHRModeR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PHR_Config_Ext_TwoPHRMode_r17_Enabled},
}

const (
	PHR_Config_Ext_Phr_AssumedPUSCH_Reporting_r18_Enabled = 0
)

var pHRConfigExtPhrAssumedPUSCHReportingR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PHR_Config_Ext_Phr_AssumedPUSCH_Reporting_r18_Enabled},
}

const (
	PHR_Config_Ext_Dpc_Reporting_FR1_r18_Enabled = 0
)

var pHRConfigExtDpcReportingFR1R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PHR_Config_Ext_Dpc_Reporting_FR1_r18_Enabled},
}

type PHR_Config struct {
	Phr_PeriodicTimer        int64
	Phr_ProhibitTimer        int64
	Phr_Tx_PowerFactorChange int64
	MultiplePHR              bool
	Dummy                    bool
	Phr_Type2OtherCell       bool
	Phr_ModeOtherCG          int64
	Mpe_Reporting_FR2_r16    *struct {
		Choice  int
		Release *struct{}
		Setup   *MPE_Config_FR2_r16
	}
	Mpe_Reporting_FR2_r17 *struct {
		Choice  int
		Release *struct{}
		Setup   *MPE_Config_FR2_r17
	}
	TwoPHRMode_r17                 *int64
	Phr_AssumedPUSCH_Reporting_r18 *int64
	Dpc_Reporting_FR1_r18          *int64
}

func (ie *PHR_Config) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pHRConfigConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Mpe_Reporting_FR2_r16 != nil
	hasExtGroup1 := ie.Mpe_Reporting_FR2_r17 != nil || ie.TwoPHRMode_r17 != nil
	hasExtGroup2 := ie.Phr_AssumedPUSCH_Reporting_r18 != nil || ie.Dpc_Reporting_FR1_r18 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. phr-PeriodicTimer: enumerated
	{
		if err := e.EncodeEnumerated(ie.Phr_PeriodicTimer, pHRConfigPhrPeriodicTimerConstraints); err != nil {
			return err
		}
	}

	// 3. phr-ProhibitTimer: enumerated
	{
		if err := e.EncodeEnumerated(ie.Phr_ProhibitTimer, pHRConfigPhrProhibitTimerConstraints); err != nil {
			return err
		}
	}

	// 4. phr-Tx-PowerFactorChange: enumerated
	{
		if err := e.EncodeEnumerated(ie.Phr_Tx_PowerFactorChange, pHRConfigPhrTxPowerFactorChangeConstraints); err != nil {
			return err
		}
	}

	// 5. multiplePHR: boolean
	{
		if err := e.EncodeBoolean(ie.MultiplePHR); err != nil {
			return err
		}
	}

	// 6. dummy: boolean
	{
		if err := e.EncodeBoolean(ie.Dummy); err != nil {
			return err
		}
	}

	// 7. phr-Type2OtherCell: boolean
	{
		if err := e.EncodeBoolean(ie.Phr_Type2OtherCell); err != nil {
			return err
		}
	}

	// 8. phr-ModeOtherCG: enumerated
	{
		if err := e.EncodeEnumerated(ie.Phr_ModeOtherCG, pHRConfigPhrModeOtherCGConstraints); err != nil {
			return err
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
					{Name: "mpe-Reporting-FR2-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Mpe_Reporting_FR2_r16 != nil}); err != nil {
				return err
			}

			if ie.Mpe_Reporting_FR2_r16 != nil {
				choiceEnc := ex.NewChoiceEncoder(pHRConfigExtMpeReportingFR2R16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Mpe_Reporting_FR2_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Mpe_Reporting_FR2_r16).Choice {
				case PHR_Config_Ext_Mpe_Reporting_FR2_r16_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PHR_Config_Ext_Mpe_Reporting_FR2_r16_Setup:
					if err := (*ie.Mpe_Reporting_FR2_r16).Setup.Encode(ex); err != nil {
						return err
					}
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
					{Name: "mpe-Reporting-FR2-r17", Optional: true},
					{Name: "twoPHRMode-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Mpe_Reporting_FR2_r17 != nil, ie.TwoPHRMode_r17 != nil}); err != nil {
				return err
			}

			if ie.Mpe_Reporting_FR2_r17 != nil {
				choiceEnc := ex.NewChoiceEncoder(pHRConfigExtMpeReportingFR2R17Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Mpe_Reporting_FR2_r17).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Mpe_Reporting_FR2_r17).Choice {
				case PHR_Config_Ext_Mpe_Reporting_FR2_r17_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PHR_Config_Ext_Mpe_Reporting_FR2_r17_Setup:
					if err := (*ie.Mpe_Reporting_FR2_r17).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.TwoPHRMode_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.TwoPHRMode_r17, pHRConfigExtTwoPHRModeR17Constraints); err != nil {
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
					{Name: "phr-AssumedPUSCH-Reporting-r18", Optional: true},
					{Name: "dpc-Reporting-FR1-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Phr_AssumedPUSCH_Reporting_r18 != nil, ie.Dpc_Reporting_FR1_r18 != nil}); err != nil {
				return err
			}

			if ie.Phr_AssumedPUSCH_Reporting_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Phr_AssumedPUSCH_Reporting_r18, pHRConfigExtPhrAssumedPUSCHReportingR18Constraints); err != nil {
					return err
				}
			}

			if ie.Dpc_Reporting_FR1_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Dpc_Reporting_FR1_r18, pHRConfigExtDpcReportingFR1R18Constraints); err != nil {
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

func (ie *PHR_Config) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pHRConfigConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. phr-PeriodicTimer: enumerated
	{
		v0, err := d.DecodeEnumerated(pHRConfigPhrPeriodicTimerConstraints)
		if err != nil {
			return err
		}
		ie.Phr_PeriodicTimer = v0
	}

	// 3. phr-ProhibitTimer: enumerated
	{
		v1, err := d.DecodeEnumerated(pHRConfigPhrProhibitTimerConstraints)
		if err != nil {
			return err
		}
		ie.Phr_ProhibitTimer = v1
	}

	// 4. phr-Tx-PowerFactorChange: enumerated
	{
		v2, err := d.DecodeEnumerated(pHRConfigPhrTxPowerFactorChangeConstraints)
		if err != nil {
			return err
		}
		ie.Phr_Tx_PowerFactorChange = v2
	}

	// 5. multiplePHR: boolean
	{
		v3, err := d.DecodeBoolean()
		if err != nil {
			return err
		}
		ie.MultiplePHR = v3
	}

	// 6. dummy: boolean
	{
		v4, err := d.DecodeBoolean()
		if err != nil {
			return err
		}
		ie.Dummy = v4
	}

	// 7. phr-Type2OtherCell: boolean
	{
		v5, err := d.DecodeBoolean()
		if err != nil {
			return err
		}
		ie.Phr_Type2OtherCell = v5
	}

	// 8. phr-ModeOtherCG: enumerated
	{
		v6, err := d.DecodeEnumerated(pHRConfigPhrModeOtherCGConstraints)
		if err != nil {
			return err
		}
		ie.Phr_ModeOtherCG = v6
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
				{Name: "mpe-Reporting-FR2-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Mpe_Reporting_FR2_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *MPE_Config_FR2_r16
			}{}
			choiceDec := dx.NewChoiceDecoder(pHRConfigExtMpeReportingFR2R16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Mpe_Reporting_FR2_r16).Choice = int(index)
			switch index {
			case PHR_Config_Ext_Mpe_Reporting_FR2_r16_Release:
				(*ie.Mpe_Reporting_FR2_r16).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PHR_Config_Ext_Mpe_Reporting_FR2_r16_Setup:
				(*ie.Mpe_Reporting_FR2_r16).Setup = new(MPE_Config_FR2_r16)
				if err := (*ie.Mpe_Reporting_FR2_r16).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "mpe-Reporting-FR2-r17", Optional: true},
				{Name: "twoPHRMode-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Mpe_Reporting_FR2_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *MPE_Config_FR2_r17
			}{}
			choiceDec := dx.NewChoiceDecoder(pHRConfigExtMpeReportingFR2R17Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Mpe_Reporting_FR2_r17).Choice = int(index)
			switch index {
			case PHR_Config_Ext_Mpe_Reporting_FR2_r17_Release:
				(*ie.Mpe_Reporting_FR2_r17).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PHR_Config_Ext_Mpe_Reporting_FR2_r17_Setup:
				(*ie.Mpe_Reporting_FR2_r17).Setup = new(MPE_Config_FR2_r17)
				if err := (*ie.Mpe_Reporting_FR2_r17).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(pHRConfigExtTwoPHRModeR17Constraints)
			if err != nil {
				return err
			}
			ie.TwoPHRMode_r17 = &v
		}
	}

	// Extension group 2:
	if seq.IsExtensionPresent(2) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "phr-AssumedPUSCH-Reporting-r18", Optional: true},
				{Name: "dpc-Reporting-FR1-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(pHRConfigExtPhrAssumedPUSCHReportingR18Constraints)
			if err != nil {
				return err
			}
			ie.Phr_AssumedPUSCH_Reporting_r18 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(pHRConfigExtDpcReportingFR1R18Constraints)
			if err != nil {
				return err
			}
			ie.Dpc_Reporting_FR1_r18 = &v
		}
	}

	return nil
}
