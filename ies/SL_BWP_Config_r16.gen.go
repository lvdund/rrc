// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-BWP-Config-r16 (line 26746).

var sLBWPConfigR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-BWP-Id"},
		{Name: "sl-BWP-Generic-r16", Optional: true},
		{Name: "sl-BWP-PoolConfig-r16", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
	},
}

var sLBWPConfigR16ExtSlBWPPoolConfigPSR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	SL_BWP_Config_r16_Ext_Sl_BWP_PoolConfigPS_r17_Release = 0
	SL_BWP_Config_r16_Ext_Sl_BWP_PoolConfigPS_r17_Setup   = 1
)

var sLBWPConfigR16ExtSlBWPDiscPoolConfigR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	SL_BWP_Config_r16_Ext_Sl_BWP_DiscPoolConfig_r17_Release = 0
	SL_BWP_Config_r16_Ext_Sl_BWP_DiscPoolConfig_r17_Setup   = 1
)

var sLBWPConfigR16ExtSlBWPPoolConfigA2XR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	SL_BWP_Config_r16_Ext_Sl_BWP_PoolConfigA2X_r18_Release = 0
	SL_BWP_Config_r16_Ext_Sl_BWP_PoolConfigA2X_r18_Setup   = 1
)

var sLBWPConfigR16ExtSlBWPPRSPoolConfigR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	SL_BWP_Config_r16_Ext_Sl_BWP_PRS_PoolConfig_r18_Release = 0
	SL_BWP_Config_r16_Ext_Sl_BWP_PRS_PoolConfig_r18_Setup   = 1
)

type SL_BWP_Config_r16 struct {
	Sl_BWP_Id               BWP_Id
	Sl_BWP_Generic_r16      *SL_BWP_Generic_r16
	Sl_BWP_PoolConfig_r16   *SL_BWP_PoolConfig_r16
	Sl_BWP_PoolConfigPS_r17 *struct {
		Choice  int
		Release *struct{}
		Setup   *SL_BWP_PoolConfig_r16
	}
	Sl_BWP_DiscPoolConfig_r17 *struct {
		Choice  int
		Release *struct{}
		Setup   *SL_BWP_DiscPoolConfig_r17
	}
	Sl_BWP_PoolConfigA2X_r18 *struct {
		Choice  int
		Release *struct{}
		Setup   *SL_BWP_PoolConfig_r16
	}
	Sl_BWP_PRS_PoolConfig_r18 *struct {
		Choice  int
		Release *struct{}
		Setup   *SL_BWP_PRS_PoolConfig_r18
	}
}

func (ie *SL_BWP_Config_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLBWPConfigR16Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Sl_BWP_PoolConfigPS_r17 != nil || ie.Sl_BWP_DiscPoolConfig_r17 != nil
	hasExtGroup1 := ie.Sl_BWP_PoolConfigA2X_r18 != nil || ie.Sl_BWP_PRS_PoolConfig_r18 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_BWP_Generic_r16 != nil, ie.Sl_BWP_PoolConfig_r16 != nil}); err != nil {
		return err
	}

	// 3. sl-BWP-Id: ref
	{
		if err := ie.Sl_BWP_Id.Encode(e); err != nil {
			return err
		}
	}

	// 4. sl-BWP-Generic-r16: ref
	{
		if ie.Sl_BWP_Generic_r16 != nil {
			if err := ie.Sl_BWP_Generic_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. sl-BWP-PoolConfig-r16: ref
	{
		if ie.Sl_BWP_PoolConfig_r16 != nil {
			if err := ie.Sl_BWP_PoolConfig_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "sl-BWP-PoolConfigPS-r17", Optional: true},
					{Name: "sl-BWP-DiscPoolConfig-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Sl_BWP_PoolConfigPS_r17 != nil, ie.Sl_BWP_DiscPoolConfig_r17 != nil}); err != nil {
				return err
			}

			if ie.Sl_BWP_PoolConfigPS_r17 != nil {
				choiceEnc := ex.NewChoiceEncoder(sLBWPConfigR16ExtSlBWPPoolConfigPSR17Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Sl_BWP_PoolConfigPS_r17).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Sl_BWP_PoolConfigPS_r17).Choice {
				case SL_BWP_Config_r16_Ext_Sl_BWP_PoolConfigPS_r17_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case SL_BWP_Config_r16_Ext_Sl_BWP_PoolConfigPS_r17_Setup:
					if err := (*ie.Sl_BWP_PoolConfigPS_r17).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Sl_BWP_DiscPoolConfig_r17 != nil {
				choiceEnc := ex.NewChoiceEncoder(sLBWPConfigR16ExtSlBWPDiscPoolConfigR17Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Sl_BWP_DiscPoolConfig_r17).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Sl_BWP_DiscPoolConfig_r17).Choice {
				case SL_BWP_Config_r16_Ext_Sl_BWP_DiscPoolConfig_r17_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case SL_BWP_Config_r16_Ext_Sl_BWP_DiscPoolConfig_r17_Setup:
					if err := (*ie.Sl_BWP_DiscPoolConfig_r17).Setup.Encode(ex); err != nil {
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
					{Name: "sl-BWP-PoolConfigA2X-r18", Optional: true},
					{Name: "sl-BWP-PRS-PoolConfig-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Sl_BWP_PoolConfigA2X_r18 != nil, ie.Sl_BWP_PRS_PoolConfig_r18 != nil}); err != nil {
				return err
			}

			if ie.Sl_BWP_PoolConfigA2X_r18 != nil {
				choiceEnc := ex.NewChoiceEncoder(sLBWPConfigR16ExtSlBWPPoolConfigA2XR18Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Sl_BWP_PoolConfigA2X_r18).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Sl_BWP_PoolConfigA2X_r18).Choice {
				case SL_BWP_Config_r16_Ext_Sl_BWP_PoolConfigA2X_r18_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case SL_BWP_Config_r16_Ext_Sl_BWP_PoolConfigA2X_r18_Setup:
					if err := (*ie.Sl_BWP_PoolConfigA2X_r18).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Sl_BWP_PRS_PoolConfig_r18 != nil {
				choiceEnc := ex.NewChoiceEncoder(sLBWPConfigR16ExtSlBWPPRSPoolConfigR18Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Sl_BWP_PRS_PoolConfig_r18).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Sl_BWP_PRS_PoolConfig_r18).Choice {
				case SL_BWP_Config_r16_Ext_Sl_BWP_PRS_PoolConfig_r18_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case SL_BWP_Config_r16_Ext_Sl_BWP_PRS_PoolConfig_r18_Setup:
					if err := (*ie.Sl_BWP_PRS_PoolConfig_r18).Setup.Encode(ex); err != nil {
						return err
					}
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

func (ie *SL_BWP_Config_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLBWPConfigR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. sl-BWP-Id: ref
	{
		if err := ie.Sl_BWP_Id.Decode(d); err != nil {
			return err
		}
	}

	// 4. sl-BWP-Generic-r16: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Sl_BWP_Generic_r16 = new(SL_BWP_Generic_r16)
			if err := ie.Sl_BWP_Generic_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. sl-BWP-PoolConfig-r16: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Sl_BWP_PoolConfig_r16 = new(SL_BWP_PoolConfig_r16)
			if err := ie.Sl_BWP_PoolConfig_r16.Decode(d); err != nil {
				return err
			}
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
				{Name: "sl-BWP-PoolConfigPS-r17", Optional: true},
				{Name: "sl-BWP-DiscPoolConfig-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Sl_BWP_PoolConfigPS_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *SL_BWP_PoolConfig_r16
			}{}
			choiceDec := dx.NewChoiceDecoder(sLBWPConfigR16ExtSlBWPPoolConfigPSR17Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Sl_BWP_PoolConfigPS_r17).Choice = int(index)
			switch index {
			case SL_BWP_Config_r16_Ext_Sl_BWP_PoolConfigPS_r17_Release:
				(*ie.Sl_BWP_PoolConfigPS_r17).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case SL_BWP_Config_r16_Ext_Sl_BWP_PoolConfigPS_r17_Setup:
				(*ie.Sl_BWP_PoolConfigPS_r17).Setup = new(SL_BWP_PoolConfig_r16)
				if err := (*ie.Sl_BWP_PoolConfigPS_r17).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.Sl_BWP_DiscPoolConfig_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *SL_BWP_DiscPoolConfig_r17
			}{}
			choiceDec := dx.NewChoiceDecoder(sLBWPConfigR16ExtSlBWPDiscPoolConfigR17Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Sl_BWP_DiscPoolConfig_r17).Choice = int(index)
			switch index {
			case SL_BWP_Config_r16_Ext_Sl_BWP_DiscPoolConfig_r17_Release:
				(*ie.Sl_BWP_DiscPoolConfig_r17).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case SL_BWP_Config_r16_Ext_Sl_BWP_DiscPoolConfig_r17_Setup:
				(*ie.Sl_BWP_DiscPoolConfig_r17).Setup = new(SL_BWP_DiscPoolConfig_r17)
				if err := (*ie.Sl_BWP_DiscPoolConfig_r17).Setup.Decode(dx); err != nil {
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
				{Name: "sl-BWP-PoolConfigA2X-r18", Optional: true},
				{Name: "sl-BWP-PRS-PoolConfig-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Sl_BWP_PoolConfigA2X_r18 = &struct {
				Choice  int
				Release *struct{}
				Setup   *SL_BWP_PoolConfig_r16
			}{}
			choiceDec := dx.NewChoiceDecoder(sLBWPConfigR16ExtSlBWPPoolConfigA2XR18Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Sl_BWP_PoolConfigA2X_r18).Choice = int(index)
			switch index {
			case SL_BWP_Config_r16_Ext_Sl_BWP_PoolConfigA2X_r18_Release:
				(*ie.Sl_BWP_PoolConfigA2X_r18).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case SL_BWP_Config_r16_Ext_Sl_BWP_PoolConfigA2X_r18_Setup:
				(*ie.Sl_BWP_PoolConfigA2X_r18).Setup = new(SL_BWP_PoolConfig_r16)
				if err := (*ie.Sl_BWP_PoolConfigA2X_r18).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.Sl_BWP_PRS_PoolConfig_r18 = &struct {
				Choice  int
				Release *struct{}
				Setup   *SL_BWP_PRS_PoolConfig_r18
			}{}
			choiceDec := dx.NewChoiceDecoder(sLBWPConfigR16ExtSlBWPPRSPoolConfigR18Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Sl_BWP_PRS_PoolConfig_r18).Choice = int(index)
			switch index {
			case SL_BWP_Config_r16_Ext_Sl_BWP_PRS_PoolConfig_r18_Release:
				(*ie.Sl_BWP_PRS_PoolConfig_r18).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case SL_BWP_Config_r16_Ext_Sl_BWP_PRS_PoolConfig_r18_Setup:
				(*ie.Sl_BWP_PRS_PoolConfig_r18).Setup = new(SL_BWP_PRS_PoolConfig_r18)
				if err := (*ie.Sl_BWP_PRS_PoolConfig_r18).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
