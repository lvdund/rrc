// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SuspendConfig (line 1289).

var suspendConfigConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "fullI-RNTI"},
		{Name: "shortI-RNTI"},
		{Name: "ran-PagingCycle"},
		{Name: "ran-NotificationAreaInfo", Optional: true},
		{Name: "t380", Optional: true},
		{Name: "nextHopChainingCount"},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
		{Name: "extension-addition-2", Optional: true},
	},
}

var suspendConfigExtSdtConfigR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	SuspendConfig_Ext_Sdt_Config_r17_Release = 0
	SuspendConfig_Ext_Sdt_Config_r17_Setup   = 1
)

var suspendConfigExtSrsPosRRCInactiveR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	SuspendConfig_Ext_Srs_PosRRC_Inactive_r17_Release = 0
	SuspendConfig_Ext_Srs_PosRRC_Inactive_r17_Setup   = 1
)

var suspendConfigExtNcdSSBRedCapInitialBWPSDTR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	SuspendConfig_Ext_Ncd_SSB_RedCapInitialBWP_SDT_r17_Release = 0
	SuspendConfig_Ext_Ncd_SSB_RedCapInitialBWP_SDT_r17_Setup   = 1
)

const (
	SuspendConfig_Ext_ResumeIndication_r18_True = 0
)

var suspendConfigExtResumeIndicationR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SuspendConfig_Ext_ResumeIndication_r18_True},
}

var suspendConfigExtSrsPosRRCInactiveEnhancedR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	SuspendConfig_Ext_Srs_PosRRC_InactiveEnhanced_r18_Release = 0
	SuspendConfig_Ext_Srs_PosRRC_InactiveEnhanced_r18_Setup   = 1
)

var suspendConfigExtMulticastConfigInactiveR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	SuspendConfig_Ext_MulticastConfigInactive_r18_Release = 0
	SuspendConfig_Ext_MulticastConfigInactive_r18_Setup   = 1
)

type SuspendConfig struct {
	FullI_RNTI               I_RNTI_Value
	ShortI_RNTI              ShortI_RNTI_Value
	Ran_PagingCycle          PagingCycle
	Ran_NotificationAreaInfo *RAN_NotificationAreaInfo
	T380                     *PeriodicRNAU_TimerValue
	NextHopChainingCount     NextHopChainingCount
	Sl_UEIdentityRemote_r17  *RNTI_Value
	Sdt_Config_r17           *struct {
		Choice  int
		Release *struct{}
		Setup   *SDT_Config_r17
	}
	Srs_PosRRC_Inactive_r17 *struct {
		Choice  int
		Release *struct{}
		Setup   *SRS_PosRRC_Inactive_r17
	}
	Ran_ExtendedPagingCycle_r17      *ExtendedPagingCycle_r17
	Ncd_SSB_RedCapInitialBWP_SDT_r17 *struct {
		Choice  int
		Release *struct{}
		Setup   *NonCellDefiningSSB_r17
	}
	ResumeIndication_r18            *int64
	Srs_PosRRC_InactiveEnhanced_r18 *struct {
		Choice  int
		Release *struct{}
		Setup   *SRS_PosRRC_InactiveEnhanced_r18
	}
	Ran_ExtendedPagingCycleConfig_r18 *ExtendedPagingCycleConfig_r18
	MulticastConfigInactive_r18       *struct {
		Choice  int
		Release *struct{}
		Setup   *MulticastConfigInactive_r18
	}
}

func (ie *SuspendConfig) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(suspendConfigConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Sl_UEIdentityRemote_r17 != nil || ie.Sdt_Config_r17 != nil || ie.Srs_PosRRC_Inactive_r17 != nil || ie.Ran_ExtendedPagingCycle_r17 != nil
	hasExtGroup1 := ie.Ncd_SSB_RedCapInitialBWP_SDT_r17 != nil
	hasExtGroup2 := ie.ResumeIndication_r18 != nil || ie.Srs_PosRRC_InactiveEnhanced_r18 != nil || ie.Ran_ExtendedPagingCycleConfig_r18 != nil || ie.MulticastConfigInactive_r18 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ran_NotificationAreaInfo != nil, ie.T380 != nil}); err != nil {
		return err
	}

	// 3. fullI-RNTI: ref
	{
		if err := ie.FullI_RNTI.Encode(e); err != nil {
			return err
		}
	}

	// 4. shortI-RNTI: ref
	{
		if err := ie.ShortI_RNTI.Encode(e); err != nil {
			return err
		}
	}

	// 5. ran-PagingCycle: ref
	{
		if err := ie.Ran_PagingCycle.Encode(e); err != nil {
			return err
		}
	}

	// 6. ran-NotificationAreaInfo: ref
	{
		if ie.Ran_NotificationAreaInfo != nil {
			if err := ie.Ran_NotificationAreaInfo.Encode(e); err != nil {
				return err
			}
		}
	}

	// 7. t380: ref
	{
		if ie.T380 != nil {
			if err := ie.T380.Encode(e); err != nil {
				return err
			}
		}
	}

	// 8. nextHopChainingCount: ref
	{
		if err := ie.NextHopChainingCount.Encode(e); err != nil {
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
					{Name: "sl-UEIdentityRemote-r17", Optional: true},
					{Name: "sdt-Config-r17", Optional: true},
					{Name: "srs-PosRRC-Inactive-r17", Optional: true},
					{Name: "ran-ExtendedPagingCycle-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Sl_UEIdentityRemote_r17 != nil, ie.Sdt_Config_r17 != nil, ie.Srs_PosRRC_Inactive_r17 != nil, ie.Ran_ExtendedPagingCycle_r17 != nil}); err != nil {
				return err
			}

			if ie.Sl_UEIdentityRemote_r17 != nil {
				if err := ie.Sl_UEIdentityRemote_r17.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Sdt_Config_r17 != nil {
				choiceEnc := ex.NewChoiceEncoder(suspendConfigExtSdtConfigR17Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Sdt_Config_r17).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Sdt_Config_r17).Choice {
				case SuspendConfig_Ext_Sdt_Config_r17_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case SuspendConfig_Ext_Sdt_Config_r17_Setup:
					if err := (*ie.Sdt_Config_r17).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Srs_PosRRC_Inactive_r17 != nil {
				choiceEnc := ex.NewChoiceEncoder(suspendConfigExtSrsPosRRCInactiveR17Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Srs_PosRRC_Inactive_r17).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Srs_PosRRC_Inactive_r17).Choice {
				case SuspendConfig_Ext_Srs_PosRRC_Inactive_r17_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case SuspendConfig_Ext_Srs_PosRRC_Inactive_r17_Setup:
					if err := (*ie.Srs_PosRRC_Inactive_r17).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Ran_ExtendedPagingCycle_r17 != nil {
				if err := ie.Ran_ExtendedPagingCycle_r17.Encode(ex); err != nil {
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
					{Name: "ncd-SSB-RedCapInitialBWP-SDT-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Ncd_SSB_RedCapInitialBWP_SDT_r17 != nil}); err != nil {
				return err
			}

			if ie.Ncd_SSB_RedCapInitialBWP_SDT_r17 != nil {
				choiceEnc := ex.NewChoiceEncoder(suspendConfigExtNcdSSBRedCapInitialBWPSDTR17Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Ncd_SSB_RedCapInitialBWP_SDT_r17).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Ncd_SSB_RedCapInitialBWP_SDT_r17).Choice {
				case SuspendConfig_Ext_Ncd_SSB_RedCapInitialBWP_SDT_r17_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case SuspendConfig_Ext_Ncd_SSB_RedCapInitialBWP_SDT_r17_Setup:
					if err := (*ie.Ncd_SSB_RedCapInitialBWP_SDT_r17).Setup.Encode(ex); err != nil {
						return err
					}
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
					{Name: "resumeIndication-r18", Optional: true},
					{Name: "srs-PosRRC-InactiveEnhanced-r18", Optional: true},
					{Name: "ran-ExtendedPagingCycleConfig-r18", Optional: true},
					{Name: "multicastConfigInactive-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.ResumeIndication_r18 != nil, ie.Srs_PosRRC_InactiveEnhanced_r18 != nil, ie.Ran_ExtendedPagingCycleConfig_r18 != nil, ie.MulticastConfigInactive_r18 != nil}); err != nil {
				return err
			}

			if ie.ResumeIndication_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.ResumeIndication_r18, suspendConfigExtResumeIndicationR18Constraints); err != nil {
					return err
				}
			}

			if ie.Srs_PosRRC_InactiveEnhanced_r18 != nil {
				choiceEnc := ex.NewChoiceEncoder(suspendConfigExtSrsPosRRCInactiveEnhancedR18Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Srs_PosRRC_InactiveEnhanced_r18).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Srs_PosRRC_InactiveEnhanced_r18).Choice {
				case SuspendConfig_Ext_Srs_PosRRC_InactiveEnhanced_r18_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case SuspendConfig_Ext_Srs_PosRRC_InactiveEnhanced_r18_Setup:
					if err := (*ie.Srs_PosRRC_InactiveEnhanced_r18).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Ran_ExtendedPagingCycleConfig_r18 != nil {
				if err := ie.Ran_ExtendedPagingCycleConfig_r18.Encode(ex); err != nil {
					return err
				}
			}

			if ie.MulticastConfigInactive_r18 != nil {
				choiceEnc := ex.NewChoiceEncoder(suspendConfigExtMulticastConfigInactiveR18Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.MulticastConfigInactive_r18).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.MulticastConfigInactive_r18).Choice {
				case SuspendConfig_Ext_MulticastConfigInactive_r18_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case SuspendConfig_Ext_MulticastConfigInactive_r18_Setup:
					if err := (*ie.MulticastConfigInactive_r18).Setup.Encode(ex); err != nil {
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

func (ie *SuspendConfig) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(suspendConfigConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. fullI-RNTI: ref
	{
		if err := ie.FullI_RNTI.Decode(d); err != nil {
			return err
		}
	}

	// 4. shortI-RNTI: ref
	{
		if err := ie.ShortI_RNTI.Decode(d); err != nil {
			return err
		}
	}

	// 5. ran-PagingCycle: ref
	{
		if err := ie.Ran_PagingCycle.Decode(d); err != nil {
			return err
		}
	}

	// 6. ran-NotificationAreaInfo: ref
	{
		if seq.IsComponentPresent(3) {
			ie.Ran_NotificationAreaInfo = new(RAN_NotificationAreaInfo)
			if err := ie.Ran_NotificationAreaInfo.Decode(d); err != nil {
				return err
			}
		}
	}

	// 7. t380: ref
	{
		if seq.IsComponentPresent(4) {
			ie.T380 = new(PeriodicRNAU_TimerValue)
			if err := ie.T380.Decode(d); err != nil {
				return err
			}
		}
	}

	// 8. nextHopChainingCount: ref
	{
		if err := ie.NextHopChainingCount.Decode(d); err != nil {
			return err
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
				{Name: "sl-UEIdentityRemote-r17", Optional: true},
				{Name: "sdt-Config-r17", Optional: true},
				{Name: "srs-PosRRC-Inactive-r17", Optional: true},
				{Name: "ran-ExtendedPagingCycle-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Sl_UEIdentityRemote_r17 = new(RNTI_Value)
			if err := ie.Sl_UEIdentityRemote_r17.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.Sdt_Config_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *SDT_Config_r17
			}{}
			choiceDec := dx.NewChoiceDecoder(suspendConfigExtSdtConfigR17Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Sdt_Config_r17).Choice = int(index)
			switch index {
			case SuspendConfig_Ext_Sdt_Config_r17_Release:
				(*ie.Sdt_Config_r17).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case SuspendConfig_Ext_Sdt_Config_r17_Setup:
				(*ie.Sdt_Config_r17).Setup = new(SDT_Config_r17)
				if err := (*ie.Sdt_Config_r17).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(2) {
			ie.Srs_PosRRC_Inactive_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *SRS_PosRRC_Inactive_r17
			}{}
			choiceDec := dx.NewChoiceDecoder(suspendConfigExtSrsPosRRCInactiveR17Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Srs_PosRRC_Inactive_r17).Choice = int(index)
			switch index {
			case SuspendConfig_Ext_Srs_PosRRC_Inactive_r17_Release:
				(*ie.Srs_PosRRC_Inactive_r17).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case SuspendConfig_Ext_Srs_PosRRC_Inactive_r17_Setup:
				(*ie.Srs_PosRRC_Inactive_r17).Setup = new(SRS_PosRRC_Inactive_r17)
				if err := (*ie.Srs_PosRRC_Inactive_r17).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(3) {
			ie.Ran_ExtendedPagingCycle_r17 = new(ExtendedPagingCycle_r17)
			if err := ie.Ran_ExtendedPagingCycle_r17.Decode(dx); err != nil {
				return err
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
				{Name: "ncd-SSB-RedCapInitialBWP-SDT-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Ncd_SSB_RedCapInitialBWP_SDT_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *NonCellDefiningSSB_r17
			}{}
			choiceDec := dx.NewChoiceDecoder(suspendConfigExtNcdSSBRedCapInitialBWPSDTR17Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Ncd_SSB_RedCapInitialBWP_SDT_r17).Choice = int(index)
			switch index {
			case SuspendConfig_Ext_Ncd_SSB_RedCapInitialBWP_SDT_r17_Release:
				(*ie.Ncd_SSB_RedCapInitialBWP_SDT_r17).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case SuspendConfig_Ext_Ncd_SSB_RedCapInitialBWP_SDT_r17_Setup:
				(*ie.Ncd_SSB_RedCapInitialBWP_SDT_r17).Setup = new(NonCellDefiningSSB_r17)
				if err := (*ie.Ncd_SSB_RedCapInitialBWP_SDT_r17).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}
	}

	// Extension group 2:
	if seq.IsExtensionPresent(2) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "resumeIndication-r18", Optional: true},
				{Name: "srs-PosRRC-InactiveEnhanced-r18", Optional: true},
				{Name: "ran-ExtendedPagingCycleConfig-r18", Optional: true},
				{Name: "multicastConfigInactive-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(suspendConfigExtResumeIndicationR18Constraints)
			if err != nil {
				return err
			}
			ie.ResumeIndication_r18 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			ie.Srs_PosRRC_InactiveEnhanced_r18 = &struct {
				Choice  int
				Release *struct{}
				Setup   *SRS_PosRRC_InactiveEnhanced_r18
			}{}
			choiceDec := dx.NewChoiceDecoder(suspendConfigExtSrsPosRRCInactiveEnhancedR18Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Srs_PosRRC_InactiveEnhanced_r18).Choice = int(index)
			switch index {
			case SuspendConfig_Ext_Srs_PosRRC_InactiveEnhanced_r18_Release:
				(*ie.Srs_PosRRC_InactiveEnhanced_r18).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case SuspendConfig_Ext_Srs_PosRRC_InactiveEnhanced_r18_Setup:
				(*ie.Srs_PosRRC_InactiveEnhanced_r18).Setup = new(SRS_PosRRC_InactiveEnhanced_r18)
				if err := (*ie.Srs_PosRRC_InactiveEnhanced_r18).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(2) {
			ie.Ran_ExtendedPagingCycleConfig_r18 = new(ExtendedPagingCycleConfig_r18)
			if err := ie.Ran_ExtendedPagingCycleConfig_r18.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(3) {
			ie.MulticastConfigInactive_r18 = &struct {
				Choice  int
				Release *struct{}
				Setup   *MulticastConfigInactive_r18
			}{}
			choiceDec := dx.NewChoiceDecoder(suspendConfigExtMulticastConfigInactiveR18Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.MulticastConfigInactive_r18).Choice = int(index)
			switch index {
			case SuspendConfig_Ext_MulticastConfigInactive_r18_Release:
				(*ie.MulticastConfigInactive_r18).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case SuspendConfig_Ext_MulticastConfigInactive_r18_Setup:
				(*ie.MulticastConfigInactive_r18).Setup = new(MulticastConfigInactive_r18)
				if err := (*ie.MulticastConfigInactive_r18).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
