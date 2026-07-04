// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: RF-Parameters (line 23541).

var rFParametersConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "supportedBandListNR"},
		{Name: "supportedBandCombinationList", Optional: true},
		{Name: "appliedFreqBandListFilter", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
		{Name: "extension-addition-2", Optional: true},
		{Name: "extension-addition-3", Optional: true},
		{Name: "extension-addition-4", Optional: true},
		{Name: "extension-addition-5", Optional: true},
		{Name: "extension-addition-6", Optional: true},
		{Name: "extension-addition-7", Optional: true},
		{Name: "extension-addition-8", Optional: true},
		{Name: "extension-addition-9", Optional: true},
		{Name: "extension-addition-10", Optional: true},
		{Name: "extension-addition-11", Optional: true},
		{Name: "extension-addition-12", Optional: true},
		{Name: "extension-addition-13", Optional: true},
		{Name: "extension-addition-14", Optional: true},
		{Name: "extension-addition-15", Optional: true},
		{Name: "extension-addition-16", Optional: true},
		{Name: "extension-addition-17", Optional: true},
		{Name: "extension-addition-18", Optional: true},
		{Name: "extension-addition-19", Optional: true},
		{Name: "extension-addition-20", Optional: true},
		{Name: "extension-addition-21", Optional: true},
		{Name: "extension-addition-22", Optional: true},
	},
}

var rFParametersSupportedBandListNRConstraints = per.SizeRange(1, common.MaxBands)

const (
	RF_Parameters_Ext_Srs_SwitchingTimeRequested_True = 0
)

var rFParametersExtSrsSwitchingTimeRequestedConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RF_Parameters_Ext_Srs_SwitchingTimeRequested_True},
}

const (
	RF_Parameters_Ext_ExtendedBand_N77_r16_Supported = 0
)

var rFParametersExtExtendedBandN77R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RF_Parameters_Ext_ExtendedBand_N77_r16_Supported},
}

var rFParametersSupportedBandCombinationListSLRelayDiscoveryR17Constraints = per.SizeConstraints{}

var rFParametersSupportedBandCombinationListSLNonRelayDiscoveryR17Constraints = per.SizeConstraints{}

const (
	RF_Parameters_Ext_SidelinkRequested_r17_True = 0
)

var rFParametersExtSidelinkRequestedR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RF_Parameters_Ext_SidelinkRequested_r17_True},
}

const (
	RF_Parameters_Ext_ExtendedBand_N77_2_r17_Supported = 0
)

var rFParametersExtExtendedBandN772R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RF_Parameters_Ext_ExtendedBand_N77_2_r17_Supported},
}

const (
	RF_Parameters_Ext_Cjtc_DdReportHighAccuracy_r19_Supported = 0
)

var rFParametersExtCjtcDdReportHighAccuracyR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RF_Parameters_Ext_Cjtc_DdReportHighAccuracy_r19_Supported},
}

const (
	RF_Parameters_Ext_Cjtc_FO_ReportHighAccuracy_r19_Supported = 0
)

var rFParametersExtCjtcFOReportHighAccuracyR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RF_Parameters_Ext_Cjtc_FO_ReportHighAccuracy_r19_Supported},
}

var rFParametersExtSupportedBandCombinationListSLU2URelayR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "supportedBandCombinationListSL-U2U-RelayDiscovery-r18", Optional: true},
		{Name: "supportedBandCombinationListSL-U2U-DiscoveryExt", Optional: true},
	},
}

type RF_Parameters struct {
	SupportedBandListNR                                    []BandNR
	SupportedBandCombinationList                           *BandCombinationList
	AppliedFreqBandListFilter                              *FreqBandList
	SupportedBandCombinationList_v1540                     *BandCombinationList_v1540
	Srs_SwitchingTimeRequested                             *int64
	SupportedBandCombinationList_v1550                     *BandCombinationList_v1550
	SupportedBandCombinationList_v1560                     *BandCombinationList_v1560
	SupportedBandCombinationList_v1610                     *BandCombinationList_v1610
	SupportedBandCombinationListSidelinkEUTRA_NR_r16       *BandCombinationListSidelinkEUTRA_NR_r16
	SupportedBandCombinationList_UplinkTxSwitch_r16        *BandCombinationList_UplinkTxSwitch_r16
	SupportedBandCombinationList_v1630                     *BandCombinationList_v1630
	SupportedBandCombinationListSidelinkEUTRA_NR_v1630     *BandCombinationListSidelinkEUTRA_NR_v1630
	SupportedBandCombinationList_UplinkTxSwitch_v1630      *BandCombinationList_UplinkTxSwitch_v1630
	SupportedBandCombinationList_v1640                     *BandCombinationList_v1640
	SupportedBandCombinationList_UplinkTxSwitch_v1640      *BandCombinationList_UplinkTxSwitch_v1640
	SupportedBandCombinationList_v1650                     *BandCombinationList_v1650
	SupportedBandCombinationList_UplinkTxSwitch_v1650      *BandCombinationList_UplinkTxSwitch_v1650
	ExtendedBand_N77_r16                                   *int64
	SupportedBandCombinationList_UplinkTxSwitch_v1670      *BandCombinationList_UplinkTxSwitch_v1670
	SupportedBandCombinationList_v1680                     *BandCombinationList_v1680
	SupportedBandCombinationList_v1690                     *BandCombinationList_v1690
	SupportedBandCombinationList_UplinkTxSwitch_v1690      *BandCombinationList_UplinkTxSwitch_v1690
	SupportedBandCombinationList_v1700                     *BandCombinationList_v1700
	SupportedBandCombinationList_UplinkTxSwitch_v1700      *BandCombinationList_UplinkTxSwitch_v1700
	SupportedBandCombinationListSL_RelayDiscovery_r17      []byte
	SupportedBandCombinationListSL_NonRelayDiscovery_r17   []byte
	SupportedBandCombinationListSidelinkEUTRA_NR_v1710     *BandCombinationListSidelinkEUTRA_NR_v1710
	SidelinkRequested_r17                                  *int64
	ExtendedBand_N77_2_r17                                 *int64
	SupportedBandCombinationList_v1720                     *BandCombinationList_v1720
	SupportedBandCombinationList_UplinkTxSwitch_v1720      *BandCombinationList_UplinkTxSwitch_v1720
	SupportedBandCombinationList_v1730                     *BandCombinationList_v1730
	SupportedBandCombinationList_UplinkTxSwitch_v1730      *BandCombinationList_UplinkTxSwitch_v1730
	SupportedBandCombinationListSL_RelayDiscovery_v1730    *BandCombinationListSL_Discovery_r17
	SupportedBandCombinationListSL_NonRelayDiscovery_v1730 *BandCombinationListSL_Discovery_r17
	SupportedBandCombinationList_v1740                     *BandCombinationList_v1740
	SupportedBandCombinationList_UplinkTxSwitch_v1740      *BandCombinationList_UplinkTxSwitch_v1740
	SupportedBandCombinationList_v1760                     *BandCombinationList_v1760
	SupportedBandCombinationList_UplinkTxSwitch_v1760      *BandCombinationList_UplinkTxSwitch_v1760
	Dummy1                                                 *BandCombinationList_v1770
	Dummy2                                                 *BandCombinationList_UplinkTxSwitch_v1770
	SupportedBandCombinationList_v1780                     *BandCombinationList_v1780
	SupportedBandCombinationList_UplinkTxSwitch_v1780      *BandCombinationList_UplinkTxSwitch_v1780
	SupportedBandCombinationList_v1800                     *BandCombinationList_v1800
	SupportedBandCombinationList_UplinkTxSwitch_v1800      *BandCombinationList_UplinkTxSwitch_v1800
	SupportedBandCombinationListSL_U2U_Relay_r18           *struct {
		SupportedBandCombinationListSL_U2U_RelayDiscovery_r18 []byte
		SupportedBandCombinationListSL_U2U_DiscoveryExt       *BandCombinationListSL_Discovery_r17
	}
	SupportedBandCombinationList_v1830                *BandCombinationList_v1830
	SupportedBandCombinationList_UplinkTxSwitch_v1830 *BandCombinationList_UplinkTxSwitch_v1830
	SupportedBandCombinationList_v1840                *BandCombinationList_v1840
	SupportedBandCombinationList_UplinkTxSwitch_v1840 *BandCombinationList_UplinkTxSwitch_v1840
	SupportedBandCombinationList_v1860                *BandCombinationList_v1860
	SupportedBandCombinationList_UplinkTxSwitch_v1860 *BandCombinationList_UplinkTxSwitch_v1860
	SupportedBandCombinationList_v1900                *BandCombinationList_v1900
	SupportedBandCombinationList_UplinkTxSwitch_v1900 *BandCombinationList_UplinkTxSwitch_v1900
	Cjtc_DdReportHighAccuracy_r19                     *int64
	Cjtc_FO_ReportHighAccuracy_r19                    *int64
}

func (ie *RF_Parameters) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rFParametersConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.SupportedBandCombinationList_v1540 != nil || ie.Srs_SwitchingTimeRequested != nil
	hasExtGroup1 := ie.SupportedBandCombinationList_v1550 != nil
	hasExtGroup2 := ie.SupportedBandCombinationList_v1560 != nil
	hasExtGroup3 := ie.SupportedBandCombinationList_v1610 != nil || ie.SupportedBandCombinationListSidelinkEUTRA_NR_r16 != nil || ie.SupportedBandCombinationList_UplinkTxSwitch_r16 != nil
	hasExtGroup4 := ie.SupportedBandCombinationList_v1630 != nil || ie.SupportedBandCombinationListSidelinkEUTRA_NR_v1630 != nil || ie.SupportedBandCombinationList_UplinkTxSwitch_v1630 != nil
	hasExtGroup5 := ie.SupportedBandCombinationList_v1640 != nil || ie.SupportedBandCombinationList_UplinkTxSwitch_v1640 != nil
	hasExtGroup6 := ie.SupportedBandCombinationList_v1650 != nil || ie.SupportedBandCombinationList_UplinkTxSwitch_v1650 != nil
	hasExtGroup7 := ie.ExtendedBand_N77_r16 != nil
	hasExtGroup8 := ie.SupportedBandCombinationList_UplinkTxSwitch_v1670 != nil
	hasExtGroup9 := ie.SupportedBandCombinationList_v1680 != nil
	hasExtGroup10 := ie.SupportedBandCombinationList_v1690 != nil || ie.SupportedBandCombinationList_UplinkTxSwitch_v1690 != nil
	hasExtGroup11 := ie.SupportedBandCombinationList_v1700 != nil || ie.SupportedBandCombinationList_UplinkTxSwitch_v1700 != nil || ie.SupportedBandCombinationListSL_RelayDiscovery_r17 != nil || ie.SupportedBandCombinationListSL_NonRelayDiscovery_r17 != nil || ie.SupportedBandCombinationListSidelinkEUTRA_NR_v1710 != nil || ie.SidelinkRequested_r17 != nil || ie.ExtendedBand_N77_2_r17 != nil
	hasExtGroup12 := ie.SupportedBandCombinationList_v1720 != nil || ie.SupportedBandCombinationList_UplinkTxSwitch_v1720 != nil
	hasExtGroup13 := ie.SupportedBandCombinationList_v1730 != nil || ie.SupportedBandCombinationList_UplinkTxSwitch_v1730 != nil || ie.SupportedBandCombinationListSL_RelayDiscovery_v1730 != nil || ie.SupportedBandCombinationListSL_NonRelayDiscovery_v1730 != nil
	hasExtGroup14 := ie.SupportedBandCombinationList_v1740 != nil || ie.SupportedBandCombinationList_UplinkTxSwitch_v1740 != nil
	hasExtGroup15 := ie.SupportedBandCombinationList_v1760 != nil || ie.SupportedBandCombinationList_UplinkTxSwitch_v1760 != nil
	hasExtGroup16 := ie.Dummy1 != nil || ie.Dummy2 != nil
	hasExtGroup17 := ie.SupportedBandCombinationList_v1780 != nil || ie.SupportedBandCombinationList_UplinkTxSwitch_v1780 != nil
	hasExtGroup18 := ie.SupportedBandCombinationList_v1800 != nil || ie.SupportedBandCombinationList_UplinkTxSwitch_v1800 != nil || ie.SupportedBandCombinationListSL_U2U_Relay_r18 != nil
	hasExtGroup19 := ie.SupportedBandCombinationList_v1830 != nil || ie.SupportedBandCombinationList_UplinkTxSwitch_v1830 != nil
	hasExtGroup20 := ie.SupportedBandCombinationList_v1840 != nil || ie.SupportedBandCombinationList_UplinkTxSwitch_v1840 != nil
	hasExtGroup21 := ie.SupportedBandCombinationList_v1860 != nil || ie.SupportedBandCombinationList_UplinkTxSwitch_v1860 != nil
	hasExtGroup22 := ie.SupportedBandCombinationList_v1900 != nil || ie.SupportedBandCombinationList_UplinkTxSwitch_v1900 != nil || ie.Cjtc_DdReportHighAccuracy_r19 != nil || ie.Cjtc_FO_ReportHighAccuracy_r19 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2 || hasExtGroup3 || hasExtGroup4 || hasExtGroup5 || hasExtGroup6 || hasExtGroup7 || hasExtGroup8 || hasExtGroup9 || hasExtGroup10 || hasExtGroup11 || hasExtGroup12 || hasExtGroup13 || hasExtGroup14 || hasExtGroup15 || hasExtGroup16 || hasExtGroup17 || hasExtGroup18 || hasExtGroup19 || hasExtGroup20 || hasExtGroup21 || hasExtGroup22

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.SupportedBandCombinationList != nil, ie.AppliedFreqBandListFilter != nil}); err != nil {
		return err
	}

	// 3. supportedBandListNR: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(rFParametersSupportedBandListNRConstraints)
		if err := seqOf.EncodeLength(int64(len(ie.SupportedBandListNR))); err != nil {
			return err
		}
		for i := range ie.SupportedBandListNR {
			if err := ie.SupportedBandListNR[i].Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. supportedBandCombinationList: ref
	{
		if ie.SupportedBandCombinationList != nil {
			if err := ie.SupportedBandCombinationList.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. appliedFreqBandListFilter: ref
	{
		if ie.AppliedFreqBandListFilter != nil {
			if err := ie.AppliedFreqBandListFilter.Encode(e); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1, hasExtGroup2, hasExtGroup3, hasExtGroup4, hasExtGroup5, hasExtGroup6, hasExtGroup7, hasExtGroup8, hasExtGroup9, hasExtGroup10, hasExtGroup11, hasExtGroup12, hasExtGroup13, hasExtGroup14, hasExtGroup15, hasExtGroup16, hasExtGroup17, hasExtGroup18, hasExtGroup19, hasExtGroup20, hasExtGroup21, hasExtGroup22}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "supportedBandCombinationList-v1540", Optional: true},
					{Name: "srs-SwitchingTimeRequested", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.SupportedBandCombinationList_v1540 != nil, ie.Srs_SwitchingTimeRequested != nil}); err != nil {
				return err
			}

			if ie.SupportedBandCombinationList_v1540 != nil {
				if err := ie.SupportedBandCombinationList_v1540.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Srs_SwitchingTimeRequested != nil {
				if err := ex.EncodeEnumerated(*ie.Srs_SwitchingTimeRequested, rFParametersExtSrsSwitchingTimeRequestedConstraints); err != nil {
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
					{Name: "supportedBandCombinationList-v1550", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.SupportedBandCombinationList_v1550 != nil}); err != nil {
				return err
			}

			if ie.SupportedBandCombinationList_v1550 != nil {
				if err := ie.SupportedBandCombinationList_v1550.Encode(ex); err != nil {
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
					{Name: "supportedBandCombinationList-v1560", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.SupportedBandCombinationList_v1560 != nil}); err != nil {
				return err
			}

			if ie.SupportedBandCombinationList_v1560 != nil {
				if err := ie.SupportedBandCombinationList_v1560.Encode(ex); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 3:
		if hasExtGroup3 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "supportedBandCombinationList-v1610", Optional: true},
					{Name: "supportedBandCombinationListSidelinkEUTRA-NR-r16", Optional: true},
					{Name: "supportedBandCombinationList-UplinkTxSwitch-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.SupportedBandCombinationList_v1610 != nil, ie.SupportedBandCombinationListSidelinkEUTRA_NR_r16 != nil, ie.SupportedBandCombinationList_UplinkTxSwitch_r16 != nil}); err != nil {
				return err
			}

			if ie.SupportedBandCombinationList_v1610 != nil {
				if err := ie.SupportedBandCombinationList_v1610.Encode(ex); err != nil {
					return err
				}
			}

			if ie.SupportedBandCombinationListSidelinkEUTRA_NR_r16 != nil {
				if err := ie.SupportedBandCombinationListSidelinkEUTRA_NR_r16.Encode(ex); err != nil {
					return err
				}
			}

			if ie.SupportedBandCombinationList_UplinkTxSwitch_r16 != nil {
				if err := ie.SupportedBandCombinationList_UplinkTxSwitch_r16.Encode(ex); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 4:
		if hasExtGroup4 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "supportedBandCombinationList-v1630", Optional: true},
					{Name: "supportedBandCombinationListSidelinkEUTRA-NR-v1630", Optional: true},
					{Name: "supportedBandCombinationList-UplinkTxSwitch-v1630", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.SupportedBandCombinationList_v1630 != nil, ie.SupportedBandCombinationListSidelinkEUTRA_NR_v1630 != nil, ie.SupportedBandCombinationList_UplinkTxSwitch_v1630 != nil}); err != nil {
				return err
			}

			if ie.SupportedBandCombinationList_v1630 != nil {
				if err := ie.SupportedBandCombinationList_v1630.Encode(ex); err != nil {
					return err
				}
			}

			if ie.SupportedBandCombinationListSidelinkEUTRA_NR_v1630 != nil {
				if err := ie.SupportedBandCombinationListSidelinkEUTRA_NR_v1630.Encode(ex); err != nil {
					return err
				}
			}

			if ie.SupportedBandCombinationList_UplinkTxSwitch_v1630 != nil {
				if err := ie.SupportedBandCombinationList_UplinkTxSwitch_v1630.Encode(ex); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 5:
		if hasExtGroup5 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "supportedBandCombinationList-v1640", Optional: true},
					{Name: "supportedBandCombinationList-UplinkTxSwitch-v1640", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.SupportedBandCombinationList_v1640 != nil, ie.SupportedBandCombinationList_UplinkTxSwitch_v1640 != nil}); err != nil {
				return err
			}

			if ie.SupportedBandCombinationList_v1640 != nil {
				if err := ie.SupportedBandCombinationList_v1640.Encode(ex); err != nil {
					return err
				}
			}

			if ie.SupportedBandCombinationList_UplinkTxSwitch_v1640 != nil {
				if err := ie.SupportedBandCombinationList_UplinkTxSwitch_v1640.Encode(ex); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 6:
		if hasExtGroup6 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "supportedBandCombinationList-v1650", Optional: true},
					{Name: "supportedBandCombinationList-UplinkTxSwitch-v1650", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.SupportedBandCombinationList_v1650 != nil, ie.SupportedBandCombinationList_UplinkTxSwitch_v1650 != nil}); err != nil {
				return err
			}

			if ie.SupportedBandCombinationList_v1650 != nil {
				if err := ie.SupportedBandCombinationList_v1650.Encode(ex); err != nil {
					return err
				}
			}

			if ie.SupportedBandCombinationList_UplinkTxSwitch_v1650 != nil {
				if err := ie.SupportedBandCombinationList_UplinkTxSwitch_v1650.Encode(ex); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 7:
		if hasExtGroup7 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "extendedBand-n77-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.ExtendedBand_N77_r16 != nil}); err != nil {
				return err
			}

			if ie.ExtendedBand_N77_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.ExtendedBand_N77_r16, rFParametersExtExtendedBandN77R16Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 8:
		if hasExtGroup8 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "supportedBandCombinationList-UplinkTxSwitch-v1670", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.SupportedBandCombinationList_UplinkTxSwitch_v1670 != nil}); err != nil {
				return err
			}

			if ie.SupportedBandCombinationList_UplinkTxSwitch_v1670 != nil {
				if err := ie.SupportedBandCombinationList_UplinkTxSwitch_v1670.Encode(ex); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 9:
		if hasExtGroup9 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "supportedBandCombinationList-v1680", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.SupportedBandCombinationList_v1680 != nil}); err != nil {
				return err
			}

			if ie.SupportedBandCombinationList_v1680 != nil {
				if err := ie.SupportedBandCombinationList_v1680.Encode(ex); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 10:
		if hasExtGroup10 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "supportedBandCombinationList-v1690", Optional: true},
					{Name: "supportedBandCombinationList-UplinkTxSwitch-v1690", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.SupportedBandCombinationList_v1690 != nil, ie.SupportedBandCombinationList_UplinkTxSwitch_v1690 != nil}); err != nil {
				return err
			}

			if ie.SupportedBandCombinationList_v1690 != nil {
				if err := ie.SupportedBandCombinationList_v1690.Encode(ex); err != nil {
					return err
				}
			}

			if ie.SupportedBandCombinationList_UplinkTxSwitch_v1690 != nil {
				if err := ie.SupportedBandCombinationList_UplinkTxSwitch_v1690.Encode(ex); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 11:
		if hasExtGroup11 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "supportedBandCombinationList-v1700", Optional: true},
					{Name: "supportedBandCombinationList-UplinkTxSwitch-v1700", Optional: true},
					{Name: "supportedBandCombinationListSL-RelayDiscovery-r17", Optional: true},
					{Name: "supportedBandCombinationListSL-NonRelayDiscovery-r17", Optional: true},
					{Name: "supportedBandCombinationListSidelinkEUTRA-NR-v1710", Optional: true},
					{Name: "sidelinkRequested-r17", Optional: true},
					{Name: "extendedBand-n77-2-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.SupportedBandCombinationList_v1700 != nil, ie.SupportedBandCombinationList_UplinkTxSwitch_v1700 != nil, ie.SupportedBandCombinationListSL_RelayDiscovery_r17 != nil, ie.SupportedBandCombinationListSL_NonRelayDiscovery_r17 != nil, ie.SupportedBandCombinationListSidelinkEUTRA_NR_v1710 != nil, ie.SidelinkRequested_r17 != nil, ie.ExtendedBand_N77_2_r17 != nil}); err != nil {
				return err
			}

			if ie.SupportedBandCombinationList_v1700 != nil {
				if err := ie.SupportedBandCombinationList_v1700.Encode(ex); err != nil {
					return err
				}
			}

			if ie.SupportedBandCombinationList_UplinkTxSwitch_v1700 != nil {
				if err := ie.SupportedBandCombinationList_UplinkTxSwitch_v1700.Encode(ex); err != nil {
					return err
				}
			}

			if ie.SupportedBandCombinationListSL_RelayDiscovery_r17 != nil {
				if err := ex.EncodeOctetString(ie.SupportedBandCombinationListSL_RelayDiscovery_r17, rFParametersSupportedBandCombinationListSLRelayDiscoveryR17Constraints); err != nil {
					return err
				}
			}

			if ie.SupportedBandCombinationListSL_NonRelayDiscovery_r17 != nil {
				if err := ex.EncodeOctetString(ie.SupportedBandCombinationListSL_NonRelayDiscovery_r17, rFParametersSupportedBandCombinationListSLNonRelayDiscoveryR17Constraints); err != nil {
					return err
				}
			}

			if ie.SupportedBandCombinationListSidelinkEUTRA_NR_v1710 != nil {
				if err := ie.SupportedBandCombinationListSidelinkEUTRA_NR_v1710.Encode(ex); err != nil {
					return err
				}
			}

			if ie.SidelinkRequested_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.SidelinkRequested_r17, rFParametersExtSidelinkRequestedR17Constraints); err != nil {
					return err
				}
			}

			if ie.ExtendedBand_N77_2_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.ExtendedBand_N77_2_r17, rFParametersExtExtendedBandN772R17Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 12:
		if hasExtGroup12 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "supportedBandCombinationList-v1720", Optional: true},
					{Name: "supportedBandCombinationList-UplinkTxSwitch-v1720", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.SupportedBandCombinationList_v1720 != nil, ie.SupportedBandCombinationList_UplinkTxSwitch_v1720 != nil}); err != nil {
				return err
			}

			if ie.SupportedBandCombinationList_v1720 != nil {
				if err := ie.SupportedBandCombinationList_v1720.Encode(ex); err != nil {
					return err
				}
			}

			if ie.SupportedBandCombinationList_UplinkTxSwitch_v1720 != nil {
				if err := ie.SupportedBandCombinationList_UplinkTxSwitch_v1720.Encode(ex); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 13:
		if hasExtGroup13 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "supportedBandCombinationList-v1730", Optional: true},
					{Name: "supportedBandCombinationList-UplinkTxSwitch-v1730", Optional: true},
					{Name: "supportedBandCombinationListSL-RelayDiscovery-v1730", Optional: true},
					{Name: "supportedBandCombinationListSL-NonRelayDiscovery-v1730", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.SupportedBandCombinationList_v1730 != nil, ie.SupportedBandCombinationList_UplinkTxSwitch_v1730 != nil, ie.SupportedBandCombinationListSL_RelayDiscovery_v1730 != nil, ie.SupportedBandCombinationListSL_NonRelayDiscovery_v1730 != nil}); err != nil {
				return err
			}

			if ie.SupportedBandCombinationList_v1730 != nil {
				if err := ie.SupportedBandCombinationList_v1730.Encode(ex); err != nil {
					return err
				}
			}

			if ie.SupportedBandCombinationList_UplinkTxSwitch_v1730 != nil {
				if err := ie.SupportedBandCombinationList_UplinkTxSwitch_v1730.Encode(ex); err != nil {
					return err
				}
			}

			if ie.SupportedBandCombinationListSL_RelayDiscovery_v1730 != nil {
				if err := ie.SupportedBandCombinationListSL_RelayDiscovery_v1730.Encode(ex); err != nil {
					return err
				}
			}

			if ie.SupportedBandCombinationListSL_NonRelayDiscovery_v1730 != nil {
				if err := ie.SupportedBandCombinationListSL_NonRelayDiscovery_v1730.Encode(ex); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 14:
		if hasExtGroup14 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "supportedBandCombinationList-v1740", Optional: true},
					{Name: "supportedBandCombinationList-UplinkTxSwitch-v1740", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.SupportedBandCombinationList_v1740 != nil, ie.SupportedBandCombinationList_UplinkTxSwitch_v1740 != nil}); err != nil {
				return err
			}

			if ie.SupportedBandCombinationList_v1740 != nil {
				if err := ie.SupportedBandCombinationList_v1740.Encode(ex); err != nil {
					return err
				}
			}

			if ie.SupportedBandCombinationList_UplinkTxSwitch_v1740 != nil {
				if err := ie.SupportedBandCombinationList_UplinkTxSwitch_v1740.Encode(ex); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 15:
		if hasExtGroup15 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "supportedBandCombinationList-v1760", Optional: true},
					{Name: "supportedBandCombinationList-UplinkTxSwitch-v1760", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.SupportedBandCombinationList_v1760 != nil, ie.SupportedBandCombinationList_UplinkTxSwitch_v1760 != nil}); err != nil {
				return err
			}

			if ie.SupportedBandCombinationList_v1760 != nil {
				if err := ie.SupportedBandCombinationList_v1760.Encode(ex); err != nil {
					return err
				}
			}

			if ie.SupportedBandCombinationList_UplinkTxSwitch_v1760 != nil {
				if err := ie.SupportedBandCombinationList_UplinkTxSwitch_v1760.Encode(ex); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 16:
		if hasExtGroup16 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "dummy1", Optional: true},
					{Name: "dummy2", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Dummy1 != nil, ie.Dummy2 != nil}); err != nil {
				return err
			}

			if ie.Dummy1 != nil {
				if err := ie.Dummy1.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Dummy2 != nil {
				if err := ie.Dummy2.Encode(ex); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 17:
		if hasExtGroup17 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "supportedBandCombinationList-v1780", Optional: true},
					{Name: "supportedBandCombinationList-UplinkTxSwitch-v1780", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.SupportedBandCombinationList_v1780 != nil, ie.SupportedBandCombinationList_UplinkTxSwitch_v1780 != nil}); err != nil {
				return err
			}

			if ie.SupportedBandCombinationList_v1780 != nil {
				if err := ie.SupportedBandCombinationList_v1780.Encode(ex); err != nil {
					return err
				}
			}

			if ie.SupportedBandCombinationList_UplinkTxSwitch_v1780 != nil {
				if err := ie.SupportedBandCombinationList_UplinkTxSwitch_v1780.Encode(ex); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 18:
		if hasExtGroup18 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "supportedBandCombinationList-v1800", Optional: true},
					{Name: "supportedBandCombinationList-UplinkTxSwitch-v1800", Optional: true},
					{Name: "supportedBandCombinationListSL-U2U-Relay-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.SupportedBandCombinationList_v1800 != nil, ie.SupportedBandCombinationList_UplinkTxSwitch_v1800 != nil, ie.SupportedBandCombinationListSL_U2U_Relay_r18 != nil}); err != nil {
				return err
			}

			if ie.SupportedBandCombinationList_v1800 != nil {
				if err := ie.SupportedBandCombinationList_v1800.Encode(ex); err != nil {
					return err
				}
			}

			if ie.SupportedBandCombinationList_UplinkTxSwitch_v1800 != nil {
				if err := ie.SupportedBandCombinationList_UplinkTxSwitch_v1800.Encode(ex); err != nil {
					return err
				}
			}

			if ie.SupportedBandCombinationListSL_U2U_Relay_r18 != nil {
				c := ie.SupportedBandCombinationListSL_U2U_Relay_r18
				rFParametersExtSupportedBandCombinationListSLU2URelayR18Seq := ex.NewSequenceEncoder(rFParametersExtSupportedBandCombinationListSLU2URelayR18Constraints)
				if err := rFParametersExtSupportedBandCombinationListSLU2URelayR18Seq.EncodePreamble([]bool{c.SupportedBandCombinationListSL_U2U_RelayDiscovery_r18 != nil, c.SupportedBandCombinationListSL_U2U_DiscoveryExt != nil}); err != nil {
					return err
				}
				if c.SupportedBandCombinationListSL_U2U_RelayDiscovery_r18 != nil {
					if err := ex.EncodeOctetString(c.SupportedBandCombinationListSL_U2U_RelayDiscovery_r18, per.SizeConstraints{}); err != nil {
						return err
					}
				}
				if c.SupportedBandCombinationListSL_U2U_DiscoveryExt != nil {
					if err := c.SupportedBandCombinationListSL_U2U_DiscoveryExt.Encode(ex); err != nil {
						return err
					}
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 19:
		if hasExtGroup19 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "supportedBandCombinationList-v1830", Optional: true},
					{Name: "supportedBandCombinationList-UplinkTxSwitch-v1830", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.SupportedBandCombinationList_v1830 != nil, ie.SupportedBandCombinationList_UplinkTxSwitch_v1830 != nil}); err != nil {
				return err
			}

			if ie.SupportedBandCombinationList_v1830 != nil {
				if err := ie.SupportedBandCombinationList_v1830.Encode(ex); err != nil {
					return err
				}
			}

			if ie.SupportedBandCombinationList_UplinkTxSwitch_v1830 != nil {
				if err := ie.SupportedBandCombinationList_UplinkTxSwitch_v1830.Encode(ex); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 20:
		if hasExtGroup20 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "supportedBandCombinationList-v1840", Optional: true},
					{Name: "supportedBandCombinationList-UplinkTxSwitch-v1840", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.SupportedBandCombinationList_v1840 != nil, ie.SupportedBandCombinationList_UplinkTxSwitch_v1840 != nil}); err != nil {
				return err
			}

			if ie.SupportedBandCombinationList_v1840 != nil {
				if err := ie.SupportedBandCombinationList_v1840.Encode(ex); err != nil {
					return err
				}
			}

			if ie.SupportedBandCombinationList_UplinkTxSwitch_v1840 != nil {
				if err := ie.SupportedBandCombinationList_UplinkTxSwitch_v1840.Encode(ex); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 21:
		if hasExtGroup21 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "supportedBandCombinationList-v1860", Optional: true},
					{Name: "supportedBandCombinationList-UplinkTxSwitch-v1860", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.SupportedBandCombinationList_v1860 != nil, ie.SupportedBandCombinationList_UplinkTxSwitch_v1860 != nil}); err != nil {
				return err
			}

			if ie.SupportedBandCombinationList_v1860 != nil {
				if err := ie.SupportedBandCombinationList_v1860.Encode(ex); err != nil {
					return err
				}
			}

			if ie.SupportedBandCombinationList_UplinkTxSwitch_v1860 != nil {
				if err := ie.SupportedBandCombinationList_UplinkTxSwitch_v1860.Encode(ex); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 22:
		if hasExtGroup22 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "supportedBandCombinationList-v1900", Optional: true},
					{Name: "supportedBandCombinationList-UplinkTxSwitch-v1900", Optional: true},
					{Name: "cjtc-DdReportHighAccuracy-r19", Optional: true},
					{Name: "cjtc-FO-ReportHighAccuracy-r19", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.SupportedBandCombinationList_v1900 != nil, ie.SupportedBandCombinationList_UplinkTxSwitch_v1900 != nil, ie.Cjtc_DdReportHighAccuracy_r19 != nil, ie.Cjtc_FO_ReportHighAccuracy_r19 != nil}); err != nil {
				return err
			}

			if ie.SupportedBandCombinationList_v1900 != nil {
				if err := ie.SupportedBandCombinationList_v1900.Encode(ex); err != nil {
					return err
				}
			}

			if ie.SupportedBandCombinationList_UplinkTxSwitch_v1900 != nil {
				if err := ie.SupportedBandCombinationList_UplinkTxSwitch_v1900.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Cjtc_DdReportHighAccuracy_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Cjtc_DdReportHighAccuracy_r19, rFParametersExtCjtcDdReportHighAccuracyR19Constraints); err != nil {
					return err
				}
			}

			if ie.Cjtc_FO_ReportHighAccuracy_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Cjtc_FO_ReportHighAccuracy_r19, rFParametersExtCjtcFOReportHighAccuracyR19Constraints); err != nil {
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

func (ie *RF_Parameters) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rFParametersConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. supportedBandListNR: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(rFParametersSupportedBandListNRConstraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.SupportedBandListNR = make([]BandNR, n)
		for i := int64(0); i < n; i++ {
			if err := ie.SupportedBandListNR[i].Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. supportedBandCombinationList: ref
	{
		if seq.IsComponentPresent(1) {
			ie.SupportedBandCombinationList = new(BandCombinationList)
			if err := ie.SupportedBandCombinationList.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. appliedFreqBandListFilter: ref
	{
		if seq.IsComponentPresent(2) {
			ie.AppliedFreqBandListFilter = new(FreqBandList)
			if err := ie.AppliedFreqBandListFilter.Decode(d); err != nil {
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
				{Name: "supportedBandCombinationList-v1540", Optional: true},
				{Name: "srs-SwitchingTimeRequested", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.SupportedBandCombinationList_v1540 = new(BandCombinationList_v1540)
			if err := ie.SupportedBandCombinationList_v1540.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(rFParametersExtSrsSwitchingTimeRequestedConstraints)
			if err != nil {
				return err
			}
			ie.Srs_SwitchingTimeRequested = &v
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "supportedBandCombinationList-v1550", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.SupportedBandCombinationList_v1550 = new(BandCombinationList_v1550)
			if err := ie.SupportedBandCombinationList_v1550.Decode(dx); err != nil {
				return err
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
				{Name: "supportedBandCombinationList-v1560", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.SupportedBandCombinationList_v1560 = new(BandCombinationList_v1560)
			if err := ie.SupportedBandCombinationList_v1560.Decode(dx); err != nil {
				return err
			}
		}
	}

	// Extension group 3:
	if seq.IsExtensionPresent(3) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "supportedBandCombinationList-v1610", Optional: true},
				{Name: "supportedBandCombinationListSidelinkEUTRA-NR-r16", Optional: true},
				{Name: "supportedBandCombinationList-UplinkTxSwitch-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.SupportedBandCombinationList_v1610 = new(BandCombinationList_v1610)
			if err := ie.SupportedBandCombinationList_v1610.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.SupportedBandCombinationListSidelinkEUTRA_NR_r16 = new(BandCombinationListSidelinkEUTRA_NR_r16)
			if err := ie.SupportedBandCombinationListSidelinkEUTRA_NR_r16.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(2) {
			ie.SupportedBandCombinationList_UplinkTxSwitch_r16 = new(BandCombinationList_UplinkTxSwitch_r16)
			if err := ie.SupportedBandCombinationList_UplinkTxSwitch_r16.Decode(dx); err != nil {
				return err
			}
		}
	}

	// Extension group 4:
	if seq.IsExtensionPresent(4) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "supportedBandCombinationList-v1630", Optional: true},
				{Name: "supportedBandCombinationListSidelinkEUTRA-NR-v1630", Optional: true},
				{Name: "supportedBandCombinationList-UplinkTxSwitch-v1630", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.SupportedBandCombinationList_v1630 = new(BandCombinationList_v1630)
			if err := ie.SupportedBandCombinationList_v1630.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.SupportedBandCombinationListSidelinkEUTRA_NR_v1630 = new(BandCombinationListSidelinkEUTRA_NR_v1630)
			if err := ie.SupportedBandCombinationListSidelinkEUTRA_NR_v1630.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(2) {
			ie.SupportedBandCombinationList_UplinkTxSwitch_v1630 = new(BandCombinationList_UplinkTxSwitch_v1630)
			if err := ie.SupportedBandCombinationList_UplinkTxSwitch_v1630.Decode(dx); err != nil {
				return err
			}
		}
	}

	// Extension group 5:
	if seq.IsExtensionPresent(5) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "supportedBandCombinationList-v1640", Optional: true},
				{Name: "supportedBandCombinationList-UplinkTxSwitch-v1640", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.SupportedBandCombinationList_v1640 = new(BandCombinationList_v1640)
			if err := ie.SupportedBandCombinationList_v1640.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.SupportedBandCombinationList_UplinkTxSwitch_v1640 = new(BandCombinationList_UplinkTxSwitch_v1640)
			if err := ie.SupportedBandCombinationList_UplinkTxSwitch_v1640.Decode(dx); err != nil {
				return err
			}
		}
	}

	// Extension group 6:
	if seq.IsExtensionPresent(6) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "supportedBandCombinationList-v1650", Optional: true},
				{Name: "supportedBandCombinationList-UplinkTxSwitch-v1650", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.SupportedBandCombinationList_v1650 = new(BandCombinationList_v1650)
			if err := ie.SupportedBandCombinationList_v1650.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.SupportedBandCombinationList_UplinkTxSwitch_v1650 = new(BandCombinationList_UplinkTxSwitch_v1650)
			if err := ie.SupportedBandCombinationList_UplinkTxSwitch_v1650.Decode(dx); err != nil {
				return err
			}
		}
	}

	// Extension group 7:
	if seq.IsExtensionPresent(7) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "extendedBand-n77-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(rFParametersExtExtendedBandN77R16Constraints)
			if err != nil {
				return err
			}
			ie.ExtendedBand_N77_r16 = &v
		}
	}

	// Extension group 8:
	if seq.IsExtensionPresent(8) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "supportedBandCombinationList-UplinkTxSwitch-v1670", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.SupportedBandCombinationList_UplinkTxSwitch_v1670 = new(BandCombinationList_UplinkTxSwitch_v1670)
			if err := ie.SupportedBandCombinationList_UplinkTxSwitch_v1670.Decode(dx); err != nil {
				return err
			}
		}
	}

	// Extension group 9:
	if seq.IsExtensionPresent(9) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "supportedBandCombinationList-v1680", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.SupportedBandCombinationList_v1680 = new(BandCombinationList_v1680)
			if err := ie.SupportedBandCombinationList_v1680.Decode(dx); err != nil {
				return err
			}
		}
	}

	// Extension group 10:
	if seq.IsExtensionPresent(10) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "supportedBandCombinationList-v1690", Optional: true},
				{Name: "supportedBandCombinationList-UplinkTxSwitch-v1690", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.SupportedBandCombinationList_v1690 = new(BandCombinationList_v1690)
			if err := ie.SupportedBandCombinationList_v1690.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.SupportedBandCombinationList_UplinkTxSwitch_v1690 = new(BandCombinationList_UplinkTxSwitch_v1690)
			if err := ie.SupportedBandCombinationList_UplinkTxSwitch_v1690.Decode(dx); err != nil {
				return err
			}
		}
	}

	// Extension group 11:
	if seq.IsExtensionPresent(11) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "supportedBandCombinationList-v1700", Optional: true},
				{Name: "supportedBandCombinationList-UplinkTxSwitch-v1700", Optional: true},
				{Name: "supportedBandCombinationListSL-RelayDiscovery-r17", Optional: true},
				{Name: "supportedBandCombinationListSL-NonRelayDiscovery-r17", Optional: true},
				{Name: "supportedBandCombinationListSidelinkEUTRA-NR-v1710", Optional: true},
				{Name: "sidelinkRequested-r17", Optional: true},
				{Name: "extendedBand-n77-2-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.SupportedBandCombinationList_v1700 = new(BandCombinationList_v1700)
			if err := ie.SupportedBandCombinationList_v1700.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.SupportedBandCombinationList_UplinkTxSwitch_v1700 = new(BandCombinationList_UplinkTxSwitch_v1700)
			if err := ie.SupportedBandCombinationList_UplinkTxSwitch_v1700.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeOctetString(rFParametersSupportedBandCombinationListSLRelayDiscoveryR17Constraints)
			if err != nil {
				return err
			}
			ie.SupportedBandCombinationListSL_RelayDiscovery_r17 = v
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeOctetString(rFParametersSupportedBandCombinationListSLNonRelayDiscoveryR17Constraints)
			if err != nil {
				return err
			}
			ie.SupportedBandCombinationListSL_NonRelayDiscovery_r17 = v
		}

		if groupSeq.IsComponentPresent(4) {
			ie.SupportedBandCombinationListSidelinkEUTRA_NR_v1710 = new(BandCombinationListSidelinkEUTRA_NR_v1710)
			if err := ie.SupportedBandCombinationListSidelinkEUTRA_NR_v1710.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(5) {
			v, err := dx.DecodeEnumerated(rFParametersExtSidelinkRequestedR17Constraints)
			if err != nil {
				return err
			}
			ie.SidelinkRequested_r17 = &v
		}

		if groupSeq.IsComponentPresent(6) {
			v, err := dx.DecodeEnumerated(rFParametersExtExtendedBandN772R17Constraints)
			if err != nil {
				return err
			}
			ie.ExtendedBand_N77_2_r17 = &v
		}
	}

	// Extension group 12:
	if seq.IsExtensionPresent(12) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "supportedBandCombinationList-v1720", Optional: true},
				{Name: "supportedBandCombinationList-UplinkTxSwitch-v1720", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.SupportedBandCombinationList_v1720 = new(BandCombinationList_v1720)
			if err := ie.SupportedBandCombinationList_v1720.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.SupportedBandCombinationList_UplinkTxSwitch_v1720 = new(BandCombinationList_UplinkTxSwitch_v1720)
			if err := ie.SupportedBandCombinationList_UplinkTxSwitch_v1720.Decode(dx); err != nil {
				return err
			}
		}
	}

	// Extension group 13:
	if seq.IsExtensionPresent(13) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "supportedBandCombinationList-v1730", Optional: true},
				{Name: "supportedBandCombinationList-UplinkTxSwitch-v1730", Optional: true},
				{Name: "supportedBandCombinationListSL-RelayDiscovery-v1730", Optional: true},
				{Name: "supportedBandCombinationListSL-NonRelayDiscovery-v1730", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.SupportedBandCombinationList_v1730 = new(BandCombinationList_v1730)
			if err := ie.SupportedBandCombinationList_v1730.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.SupportedBandCombinationList_UplinkTxSwitch_v1730 = new(BandCombinationList_UplinkTxSwitch_v1730)
			if err := ie.SupportedBandCombinationList_UplinkTxSwitch_v1730.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(2) {
			ie.SupportedBandCombinationListSL_RelayDiscovery_v1730 = new(BandCombinationListSL_Discovery_r17)
			if err := ie.SupportedBandCombinationListSL_RelayDiscovery_v1730.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(3) {
			ie.SupportedBandCombinationListSL_NonRelayDiscovery_v1730 = new(BandCombinationListSL_Discovery_r17)
			if err := ie.SupportedBandCombinationListSL_NonRelayDiscovery_v1730.Decode(dx); err != nil {
				return err
			}
		}
	}

	// Extension group 14:
	if seq.IsExtensionPresent(14) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "supportedBandCombinationList-v1740", Optional: true},
				{Name: "supportedBandCombinationList-UplinkTxSwitch-v1740", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.SupportedBandCombinationList_v1740 = new(BandCombinationList_v1740)
			if err := ie.SupportedBandCombinationList_v1740.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.SupportedBandCombinationList_UplinkTxSwitch_v1740 = new(BandCombinationList_UplinkTxSwitch_v1740)
			if err := ie.SupportedBandCombinationList_UplinkTxSwitch_v1740.Decode(dx); err != nil {
				return err
			}
		}
	}

	// Extension group 15:
	if seq.IsExtensionPresent(15) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "supportedBandCombinationList-v1760", Optional: true},
				{Name: "supportedBandCombinationList-UplinkTxSwitch-v1760", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.SupportedBandCombinationList_v1760 = new(BandCombinationList_v1760)
			if err := ie.SupportedBandCombinationList_v1760.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.SupportedBandCombinationList_UplinkTxSwitch_v1760 = new(BandCombinationList_UplinkTxSwitch_v1760)
			if err := ie.SupportedBandCombinationList_UplinkTxSwitch_v1760.Decode(dx); err != nil {
				return err
			}
		}
	}

	// Extension group 16:
	if seq.IsExtensionPresent(16) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "dummy1", Optional: true},
				{Name: "dummy2", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Dummy1 = new(BandCombinationList_v1770)
			if err := ie.Dummy1.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.Dummy2 = new(BandCombinationList_UplinkTxSwitch_v1770)
			if err := ie.Dummy2.Decode(dx); err != nil {
				return err
			}
		}
	}

	// Extension group 17:
	if seq.IsExtensionPresent(17) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "supportedBandCombinationList-v1780", Optional: true},
				{Name: "supportedBandCombinationList-UplinkTxSwitch-v1780", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.SupportedBandCombinationList_v1780 = new(BandCombinationList_v1780)
			if err := ie.SupportedBandCombinationList_v1780.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.SupportedBandCombinationList_UplinkTxSwitch_v1780 = new(BandCombinationList_UplinkTxSwitch_v1780)
			if err := ie.SupportedBandCombinationList_UplinkTxSwitch_v1780.Decode(dx); err != nil {
				return err
			}
		}
	}

	// Extension group 18:
	if seq.IsExtensionPresent(18) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "supportedBandCombinationList-v1800", Optional: true},
				{Name: "supportedBandCombinationList-UplinkTxSwitch-v1800", Optional: true},
				{Name: "supportedBandCombinationListSL-U2U-Relay-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.SupportedBandCombinationList_v1800 = new(BandCombinationList_v1800)
			if err := ie.SupportedBandCombinationList_v1800.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.SupportedBandCombinationList_UplinkTxSwitch_v1800 = new(BandCombinationList_UplinkTxSwitch_v1800)
			if err := ie.SupportedBandCombinationList_UplinkTxSwitch_v1800.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(2) {
			ie.SupportedBandCombinationListSL_U2U_Relay_r18 = &struct {
				SupportedBandCombinationListSL_U2U_RelayDiscovery_r18 []byte
				SupportedBandCombinationListSL_U2U_DiscoveryExt       *BandCombinationListSL_Discovery_r17
			}{}
			c := ie.SupportedBandCombinationListSL_U2U_Relay_r18
			rFParametersExtSupportedBandCombinationListSLU2URelayR18Seq := dx.NewSequenceDecoder(rFParametersExtSupportedBandCombinationListSLU2URelayR18Constraints)
			if err := rFParametersExtSupportedBandCombinationListSLU2URelayR18Seq.DecodePreamble(); err != nil {
				return err
			}
			if rFParametersExtSupportedBandCombinationListSLU2URelayR18Seq.IsComponentPresent(0) {
				v, err := dx.DecodeOctetString(per.SizeConstraints{})
				if err != nil {
					return err
				}
				c.SupportedBandCombinationListSL_U2U_RelayDiscovery_r18 = v
			}
			if rFParametersExtSupportedBandCombinationListSLU2URelayR18Seq.IsComponentPresent(1) {
				c.SupportedBandCombinationListSL_U2U_DiscoveryExt = new(BandCombinationListSL_Discovery_r17)
				if err := (*c.SupportedBandCombinationListSL_U2U_DiscoveryExt).Decode(dx); err != nil {
					return err
				}
			}
		}
	}

	// Extension group 19:
	if seq.IsExtensionPresent(19) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "supportedBandCombinationList-v1830", Optional: true},
				{Name: "supportedBandCombinationList-UplinkTxSwitch-v1830", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.SupportedBandCombinationList_v1830 = new(BandCombinationList_v1830)
			if err := ie.SupportedBandCombinationList_v1830.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.SupportedBandCombinationList_UplinkTxSwitch_v1830 = new(BandCombinationList_UplinkTxSwitch_v1830)
			if err := ie.SupportedBandCombinationList_UplinkTxSwitch_v1830.Decode(dx); err != nil {
				return err
			}
		}
	}

	// Extension group 20:
	if seq.IsExtensionPresent(20) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "supportedBandCombinationList-v1840", Optional: true},
				{Name: "supportedBandCombinationList-UplinkTxSwitch-v1840", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.SupportedBandCombinationList_v1840 = new(BandCombinationList_v1840)
			if err := ie.SupportedBandCombinationList_v1840.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.SupportedBandCombinationList_UplinkTxSwitch_v1840 = new(BandCombinationList_UplinkTxSwitch_v1840)
			if err := ie.SupportedBandCombinationList_UplinkTxSwitch_v1840.Decode(dx); err != nil {
				return err
			}
		}
	}

	// Extension group 21:
	if seq.IsExtensionPresent(21) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "supportedBandCombinationList-v1860", Optional: true},
				{Name: "supportedBandCombinationList-UplinkTxSwitch-v1860", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.SupportedBandCombinationList_v1860 = new(BandCombinationList_v1860)
			if err := ie.SupportedBandCombinationList_v1860.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.SupportedBandCombinationList_UplinkTxSwitch_v1860 = new(BandCombinationList_UplinkTxSwitch_v1860)
			if err := ie.SupportedBandCombinationList_UplinkTxSwitch_v1860.Decode(dx); err != nil {
				return err
			}
		}
	}

	// Extension group 22:
	if seq.IsExtensionPresent(22) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "supportedBandCombinationList-v1900", Optional: true},
				{Name: "supportedBandCombinationList-UplinkTxSwitch-v1900", Optional: true},
				{Name: "cjtc-DdReportHighAccuracy-r19", Optional: true},
				{Name: "cjtc-FO-ReportHighAccuracy-r19", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.SupportedBandCombinationList_v1900 = new(BandCombinationList_v1900)
			if err := ie.SupportedBandCombinationList_v1900.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.SupportedBandCombinationList_UplinkTxSwitch_v1900 = new(BandCombinationList_UplinkTxSwitch_v1900)
			if err := ie.SupportedBandCombinationList_UplinkTxSwitch_v1900.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeEnumerated(rFParametersExtCjtcDdReportHighAccuracyR19Constraints)
			if err != nil {
				return err
			}
			ie.Cjtc_DdReportHighAccuracy_r19 = &v
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeEnumerated(rFParametersExtCjtcFOReportHighAccuracyR19Constraints)
			if err != nil {
				return err
			}
			ie.Cjtc_FO_ReportHighAccuracy_r19 = &v
		}
	}

	return nil
}
