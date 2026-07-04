// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RF-ParametersMRDC (line 24691).

var rFParametersMRDCConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
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

const (
	RF_ParametersMRDC_Ext_Srs_SwitchingTimeRequested_True = 0
)

var rFParametersMRDCExtSrsSwitchingTimeRequestedConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RF_ParametersMRDC_Ext_Srs_SwitchingTimeRequested_True},
}

var rFParametersMRDCExtSupportedBandCombinationListNEDCOnlyV15a0Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "supportedBandCombinationList-v1540", Optional: true},
		{Name: "supportedBandCombinationList-v1560", Optional: true},
		{Name: "supportedBandCombinationList-v1570", Optional: true},
		{Name: "supportedBandCombinationList-v1580", Optional: true},
		{Name: "supportedBandCombinationList-v1590", Optional: true},
	},
}

var rFParametersMRDCExtSupportedBandCombinationListNEDCOnlyV1720Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "supportedBandCombinationList-v1700", Optional: true},
		{Name: "supportedBandCombinationList-v1720", Optional: true},
	},
}

type RF_ParametersMRDC struct {
	SupportedBandCombinationList                *BandCombinationList
	AppliedFreqBandListFilter                   *FreqBandList
	Srs_SwitchingTimeRequested                  *int64
	SupportedBandCombinationList_v1540          *BandCombinationList_v1540
	SupportedBandCombinationList_v1550          *BandCombinationList_v1550
	SupportedBandCombinationList_v1560          *BandCombinationList_v1560
	SupportedBandCombinationListNEDC_Only       *BandCombinationList
	SupportedBandCombinationList_v1570          *BandCombinationList_v1570
	SupportedBandCombinationList_v1580          *BandCombinationList_v1580
	SupportedBandCombinationList_v1590          *BandCombinationList_v1590
	SupportedBandCombinationListNEDC_Only_V15a0 *struct {
		SupportedBandCombinationList_v1540 *BandCombinationList_v1540
		SupportedBandCombinationList_v1560 *BandCombinationList_v1560
		SupportedBandCombinationList_v1570 *BandCombinationList_v1570
		SupportedBandCombinationList_v1580 *BandCombinationList_v1580
		SupportedBandCombinationList_v1590 *BandCombinationList_v1590
	}
	SupportedBandCombinationList_v1610                *BandCombinationList_v1610
	SupportedBandCombinationListNEDC_Only_v1610       *BandCombinationList_v1610
	SupportedBandCombinationList_UplinkTxSwitch_r16   *BandCombinationList_UplinkTxSwitch_r16
	SupportedBandCombinationList_v1630                *BandCombinationList_v1630
	SupportedBandCombinationListNEDC_Only_v1630       *BandCombinationList_v1630
	SupportedBandCombinationList_UplinkTxSwitch_v1630 *BandCombinationList_UplinkTxSwitch_v1630
	SupportedBandCombinationList_v1640                *BandCombinationList_v1640
	SupportedBandCombinationListNEDC_Only_v1640       *BandCombinationList_v1640
	SupportedBandCombinationList_UplinkTxSwitch_v1640 *BandCombinationList_UplinkTxSwitch_v1640
	SupportedBandCombinationList_UplinkTxSwitch_v1670 *BandCombinationList_UplinkTxSwitch_v1670
	SupportedBandCombinationList_v1700                *BandCombinationList_v1700
	SupportedBandCombinationList_UplinkTxSwitch_v1700 *BandCombinationList_UplinkTxSwitch_v1700
	SupportedBandCombinationList_v1720                *BandCombinationList_v1720
	SupportedBandCombinationListNEDC_Only_v1720       *struct {
		SupportedBandCombinationList_v1700 *BandCombinationList_v1700
		SupportedBandCombinationList_v1720 *BandCombinationList_v1720
	}
	SupportedBandCombinationList_UplinkTxSwitch_v1720 *BandCombinationList_UplinkTxSwitch_v1720
	SupportedBandCombinationList_v1730                *BandCombinationList_v1730
	SupportedBandCombinationListNEDC_Only_v1730       *BandCombinationList_v1730
	SupportedBandCombinationList_UplinkTxSwitch_v1730 *BandCombinationList_UplinkTxSwitch_v1730
	SupportedBandCombinationList_v1740                *BandCombinationList_v1740
	SupportedBandCombinationListNEDC_Only_v1740       *BandCombinationList_v1740
	SupportedBandCombinationList_UplinkTxSwitch_v1740 *BandCombinationList_UplinkTxSwitch_v1740
	Dummy1                                            *BandCombinationList_v1770
	Dummy2                                            *BandCombinationList_UplinkTxSwitch_v1770
	SupportedBandCombinationList_v1780                *BandCombinationList_v1780
	SupportedBandCombinationListNEDC_Only_v1780       *BandCombinationList_v1780
	SupportedBandCombinationList_UplinkTxSwitch_v1780 *BandCombinationList_UplinkTxSwitch_v1780
	SupportedBandCombinationList_v1790                *BandCombinationList_v1790
	SupportedBandCombinationList_UplinkTxSwitch_v1790 *BandCombinationList_UplinkTxSwitch_v1790
	SupportedBandCombinationList_v1800                *BandCombinationList_v1800
	SupportedBandCombinationList_UplinkTxSwitch_v1800 *BandCombinationList_UplinkTxSwitch_v1800
	SupportedBandCombinationList_v1830                *BandCombinationList_v1830
	SupportedBandCombinationList_UplinkTxSwitch_v1830 *BandCombinationList_UplinkTxSwitch_v1830
	SupportedBandCombinationList_v1840                *BandCombinationList_v1840
	SupportedBandCombinationList_UplinkTxSwitch_v1840 *BandCombinationList_UplinkTxSwitch_v1840
	SupportedBandCombinationList_v1860                *BandCombinationList_v1860
	SupportedBandCombinationList_UplinkTxSwitch_v1860 *BandCombinationList_UplinkTxSwitch_v1860
	SupportedBandCombinationList_v1900                *BandCombinationList_v1900
	SupportedBandCombinationList_UplinkTxSwitch_v1900 *BandCombinationList_UplinkTxSwitch_v1900
}

func (ie *RF_ParametersMRDC) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rFParametersMRDCConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Srs_SwitchingTimeRequested != nil || ie.SupportedBandCombinationList_v1540 != nil
	hasExtGroup1 := ie.SupportedBandCombinationList_v1550 != nil
	hasExtGroup2 := ie.SupportedBandCombinationList_v1560 != nil || ie.SupportedBandCombinationListNEDC_Only != nil
	hasExtGroup3 := ie.SupportedBandCombinationList_v1570 != nil
	hasExtGroup4 := ie.SupportedBandCombinationList_v1580 != nil
	hasExtGroup5 := ie.SupportedBandCombinationList_v1590 != nil
	hasExtGroup6 := ie.SupportedBandCombinationListNEDC_Only_V15a0 != nil
	hasExtGroup7 := ie.SupportedBandCombinationList_v1610 != nil || ie.SupportedBandCombinationListNEDC_Only_v1610 != nil || ie.SupportedBandCombinationList_UplinkTxSwitch_r16 != nil
	hasExtGroup8 := ie.SupportedBandCombinationList_v1630 != nil || ie.SupportedBandCombinationListNEDC_Only_v1630 != nil || ie.SupportedBandCombinationList_UplinkTxSwitch_v1630 != nil
	hasExtGroup9 := ie.SupportedBandCombinationList_v1640 != nil || ie.SupportedBandCombinationListNEDC_Only_v1640 != nil || ie.SupportedBandCombinationList_UplinkTxSwitch_v1640 != nil
	hasExtGroup10 := ie.SupportedBandCombinationList_UplinkTxSwitch_v1670 != nil
	hasExtGroup11 := ie.SupportedBandCombinationList_v1700 != nil || ie.SupportedBandCombinationList_UplinkTxSwitch_v1700 != nil
	hasExtGroup12 := ie.SupportedBandCombinationList_v1720 != nil || ie.SupportedBandCombinationListNEDC_Only_v1720 != nil || ie.SupportedBandCombinationList_UplinkTxSwitch_v1720 != nil
	hasExtGroup13 := ie.SupportedBandCombinationList_v1730 != nil || ie.SupportedBandCombinationListNEDC_Only_v1730 != nil || ie.SupportedBandCombinationList_UplinkTxSwitch_v1730 != nil
	hasExtGroup14 := ie.SupportedBandCombinationList_v1740 != nil || ie.SupportedBandCombinationListNEDC_Only_v1740 != nil || ie.SupportedBandCombinationList_UplinkTxSwitch_v1740 != nil
	hasExtGroup15 := ie.Dummy1 != nil || ie.Dummy2 != nil
	hasExtGroup16 := ie.SupportedBandCombinationList_v1780 != nil || ie.SupportedBandCombinationListNEDC_Only_v1780 != nil || ie.SupportedBandCombinationList_UplinkTxSwitch_v1780 != nil
	hasExtGroup17 := ie.SupportedBandCombinationList_v1790 != nil || ie.SupportedBandCombinationList_UplinkTxSwitch_v1790 != nil
	hasExtGroup18 := ie.SupportedBandCombinationList_v1800 != nil || ie.SupportedBandCombinationList_UplinkTxSwitch_v1800 != nil
	hasExtGroup19 := ie.SupportedBandCombinationList_v1830 != nil || ie.SupportedBandCombinationList_UplinkTxSwitch_v1830 != nil
	hasExtGroup20 := ie.SupportedBandCombinationList_v1840 != nil || ie.SupportedBandCombinationList_UplinkTxSwitch_v1840 != nil
	hasExtGroup21 := ie.SupportedBandCombinationList_v1860 != nil || ie.SupportedBandCombinationList_UplinkTxSwitch_v1860 != nil
	hasExtGroup22 := ie.SupportedBandCombinationList_v1900 != nil || ie.SupportedBandCombinationList_UplinkTxSwitch_v1900 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2 || hasExtGroup3 || hasExtGroup4 || hasExtGroup5 || hasExtGroup6 || hasExtGroup7 || hasExtGroup8 || hasExtGroup9 || hasExtGroup10 || hasExtGroup11 || hasExtGroup12 || hasExtGroup13 || hasExtGroup14 || hasExtGroup15 || hasExtGroup16 || hasExtGroup17 || hasExtGroup18 || hasExtGroup19 || hasExtGroup20 || hasExtGroup21 || hasExtGroup22

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.SupportedBandCombinationList != nil, ie.AppliedFreqBandListFilter != nil}); err != nil {
		return err
	}

	// 3. supportedBandCombinationList: ref
	{
		if ie.SupportedBandCombinationList != nil {
			if err := ie.SupportedBandCombinationList.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. appliedFreqBandListFilter: ref
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
					{Name: "srs-SwitchingTimeRequested", Optional: true},
					{Name: "supportedBandCombinationList-v1540", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Srs_SwitchingTimeRequested != nil, ie.SupportedBandCombinationList_v1540 != nil}); err != nil {
				return err
			}

			if ie.Srs_SwitchingTimeRequested != nil {
				if err := ex.EncodeEnumerated(*ie.Srs_SwitchingTimeRequested, rFParametersMRDCExtSrsSwitchingTimeRequestedConstraints); err != nil {
					return err
				}
			}

			if ie.SupportedBandCombinationList_v1540 != nil {
				if err := ie.SupportedBandCombinationList_v1540.Encode(ex); err != nil {
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
					{Name: "supportedBandCombinationListNEDC-Only", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.SupportedBandCombinationList_v1560 != nil, ie.SupportedBandCombinationListNEDC_Only != nil}); err != nil {
				return err
			}

			if ie.SupportedBandCombinationList_v1560 != nil {
				if err := ie.SupportedBandCombinationList_v1560.Encode(ex); err != nil {
					return err
				}
			}

			if ie.SupportedBandCombinationListNEDC_Only != nil {
				if err := ie.SupportedBandCombinationListNEDC_Only.Encode(ex); err != nil {
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
					{Name: "supportedBandCombinationList-v1570", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.SupportedBandCombinationList_v1570 != nil}); err != nil {
				return err
			}

			if ie.SupportedBandCombinationList_v1570 != nil {
				if err := ie.SupportedBandCombinationList_v1570.Encode(ex); err != nil {
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
					{Name: "supportedBandCombinationList-v1580", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.SupportedBandCombinationList_v1580 != nil}); err != nil {
				return err
			}

			if ie.SupportedBandCombinationList_v1580 != nil {
				if err := ie.SupportedBandCombinationList_v1580.Encode(ex); err != nil {
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
					{Name: "supportedBandCombinationList-v1590", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.SupportedBandCombinationList_v1590 != nil}); err != nil {
				return err
			}

			if ie.SupportedBandCombinationList_v1590 != nil {
				if err := ie.SupportedBandCombinationList_v1590.Encode(ex); err != nil {
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
					{Name: "supportedBandCombinationListNEDC-Only-v15a0", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.SupportedBandCombinationListNEDC_Only_V15a0 != nil}); err != nil {
				return err
			}

			if ie.SupportedBandCombinationListNEDC_Only_V15a0 != nil {
				c := ie.SupportedBandCombinationListNEDC_Only_V15a0
				rFParametersMRDCExtSupportedBandCombinationListNEDCOnlyV15a0Seq := ex.NewSequenceEncoder(rFParametersMRDCExtSupportedBandCombinationListNEDCOnlyV15a0Constraints)
				if err := rFParametersMRDCExtSupportedBandCombinationListNEDCOnlyV15a0Seq.EncodePreamble([]bool{c.SupportedBandCombinationList_v1540 != nil, c.SupportedBandCombinationList_v1560 != nil, c.SupportedBandCombinationList_v1570 != nil, c.SupportedBandCombinationList_v1580 != nil, c.SupportedBandCombinationList_v1590 != nil}); err != nil {
					return err
				}
				if c.SupportedBandCombinationList_v1540 != nil {
					if err := c.SupportedBandCombinationList_v1540.Encode(ex); err != nil {
						return err
					}
				}
				if c.SupportedBandCombinationList_v1560 != nil {
					if err := c.SupportedBandCombinationList_v1560.Encode(ex); err != nil {
						return err
					}
				}
				if c.SupportedBandCombinationList_v1570 != nil {
					if err := c.SupportedBandCombinationList_v1570.Encode(ex); err != nil {
						return err
					}
				}
				if c.SupportedBandCombinationList_v1580 != nil {
					if err := c.SupportedBandCombinationList_v1580.Encode(ex); err != nil {
						return err
					}
				}
				if c.SupportedBandCombinationList_v1590 != nil {
					if err := c.SupportedBandCombinationList_v1590.Encode(ex); err != nil {
						return err
					}
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
					{Name: "supportedBandCombinationList-v1610", Optional: true},
					{Name: "supportedBandCombinationListNEDC-Only-v1610", Optional: true},
					{Name: "supportedBandCombinationList-UplinkTxSwitch-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.SupportedBandCombinationList_v1610 != nil, ie.SupportedBandCombinationListNEDC_Only_v1610 != nil, ie.SupportedBandCombinationList_UplinkTxSwitch_r16 != nil}); err != nil {
				return err
			}

			if ie.SupportedBandCombinationList_v1610 != nil {
				if err := ie.SupportedBandCombinationList_v1610.Encode(ex); err != nil {
					return err
				}
			}

			if ie.SupportedBandCombinationListNEDC_Only_v1610 != nil {
				if err := ie.SupportedBandCombinationListNEDC_Only_v1610.Encode(ex); err != nil {
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

		// Extension group 8:
		if hasExtGroup8 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "supportedBandCombinationList-v1630", Optional: true},
					{Name: "supportedBandCombinationListNEDC-Only-v1630", Optional: true},
					{Name: "supportedBandCombinationList-UplinkTxSwitch-v1630", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.SupportedBandCombinationList_v1630 != nil, ie.SupportedBandCombinationListNEDC_Only_v1630 != nil, ie.SupportedBandCombinationList_UplinkTxSwitch_v1630 != nil}); err != nil {
				return err
			}

			if ie.SupportedBandCombinationList_v1630 != nil {
				if err := ie.SupportedBandCombinationList_v1630.Encode(ex); err != nil {
					return err
				}
			}

			if ie.SupportedBandCombinationListNEDC_Only_v1630 != nil {
				if err := ie.SupportedBandCombinationListNEDC_Only_v1630.Encode(ex); err != nil {
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

		// Extension group 9:
		if hasExtGroup9 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "supportedBandCombinationList-v1640", Optional: true},
					{Name: "supportedBandCombinationListNEDC-Only-v1640", Optional: true},
					{Name: "supportedBandCombinationList-UplinkTxSwitch-v1640", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.SupportedBandCombinationList_v1640 != nil, ie.SupportedBandCombinationListNEDC_Only_v1640 != nil, ie.SupportedBandCombinationList_UplinkTxSwitch_v1640 != nil}); err != nil {
				return err
			}

			if ie.SupportedBandCombinationList_v1640 != nil {
				if err := ie.SupportedBandCombinationList_v1640.Encode(ex); err != nil {
					return err
				}
			}

			if ie.SupportedBandCombinationListNEDC_Only_v1640 != nil {
				if err := ie.SupportedBandCombinationListNEDC_Only_v1640.Encode(ex); err != nil {
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

		// Extension group 10:
		if hasExtGroup10 {
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

		// Extension group 11:
		if hasExtGroup11 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "supportedBandCombinationList-v1700", Optional: true},
					{Name: "supportedBandCombinationList-UplinkTxSwitch-v1700", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.SupportedBandCombinationList_v1700 != nil, ie.SupportedBandCombinationList_UplinkTxSwitch_v1700 != nil}); err != nil {
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
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 12:
		if hasExtGroup12 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "supportedBandCombinationList-v1720", Optional: true},
					{Name: "supportedBandCombinationListNEDC-Only-v1720", Optional: true},
					{Name: "supportedBandCombinationList-UplinkTxSwitch-v1720", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.SupportedBandCombinationList_v1720 != nil, ie.SupportedBandCombinationListNEDC_Only_v1720 != nil, ie.SupportedBandCombinationList_UplinkTxSwitch_v1720 != nil}); err != nil {
				return err
			}

			if ie.SupportedBandCombinationList_v1720 != nil {
				if err := ie.SupportedBandCombinationList_v1720.Encode(ex); err != nil {
					return err
				}
			}

			if ie.SupportedBandCombinationListNEDC_Only_v1720 != nil {
				c := ie.SupportedBandCombinationListNEDC_Only_v1720
				rFParametersMRDCExtSupportedBandCombinationListNEDCOnlyV1720Seq := ex.NewSequenceEncoder(rFParametersMRDCExtSupportedBandCombinationListNEDCOnlyV1720Constraints)
				if err := rFParametersMRDCExtSupportedBandCombinationListNEDCOnlyV1720Seq.EncodePreamble([]bool{c.SupportedBandCombinationList_v1700 != nil, c.SupportedBandCombinationList_v1720 != nil}); err != nil {
					return err
				}
				if c.SupportedBandCombinationList_v1700 != nil {
					if err := c.SupportedBandCombinationList_v1700.Encode(ex); err != nil {
						return err
					}
				}
				if c.SupportedBandCombinationList_v1720 != nil {
					if err := c.SupportedBandCombinationList_v1720.Encode(ex); err != nil {
						return err
					}
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
					{Name: "supportedBandCombinationListNEDC-Only-v1730", Optional: true},
					{Name: "supportedBandCombinationList-UplinkTxSwitch-v1730", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.SupportedBandCombinationList_v1730 != nil, ie.SupportedBandCombinationListNEDC_Only_v1730 != nil, ie.SupportedBandCombinationList_UplinkTxSwitch_v1730 != nil}); err != nil {
				return err
			}

			if ie.SupportedBandCombinationList_v1730 != nil {
				if err := ie.SupportedBandCombinationList_v1730.Encode(ex); err != nil {
					return err
				}
			}

			if ie.SupportedBandCombinationListNEDC_Only_v1730 != nil {
				if err := ie.SupportedBandCombinationListNEDC_Only_v1730.Encode(ex); err != nil {
					return err
				}
			}

			if ie.SupportedBandCombinationList_UplinkTxSwitch_v1730 != nil {
				if err := ie.SupportedBandCombinationList_UplinkTxSwitch_v1730.Encode(ex); err != nil {
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
					{Name: "supportedBandCombinationListNEDC-Only-v1740", Optional: true},
					{Name: "supportedBandCombinationList-UplinkTxSwitch-v1740", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.SupportedBandCombinationList_v1740 != nil, ie.SupportedBandCombinationListNEDC_Only_v1740 != nil, ie.SupportedBandCombinationList_UplinkTxSwitch_v1740 != nil}); err != nil {
				return err
			}

			if ie.SupportedBandCombinationList_v1740 != nil {
				if err := ie.SupportedBandCombinationList_v1740.Encode(ex); err != nil {
					return err
				}
			}

			if ie.SupportedBandCombinationListNEDC_Only_v1740 != nil {
				if err := ie.SupportedBandCombinationListNEDC_Only_v1740.Encode(ex); err != nil {
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

		// Extension group 16:
		if hasExtGroup16 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "supportedBandCombinationList-v1780", Optional: true},
					{Name: "supportedBandCombinationListNEDC-Only-v1780", Optional: true},
					{Name: "supportedBandCombinationList-UplinkTxSwitch-v1780", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.SupportedBandCombinationList_v1780 != nil, ie.SupportedBandCombinationListNEDC_Only_v1780 != nil, ie.SupportedBandCombinationList_UplinkTxSwitch_v1780 != nil}); err != nil {
				return err
			}

			if ie.SupportedBandCombinationList_v1780 != nil {
				if err := ie.SupportedBandCombinationList_v1780.Encode(ex); err != nil {
					return err
				}
			}

			if ie.SupportedBandCombinationListNEDC_Only_v1780 != nil {
				if err := ie.SupportedBandCombinationListNEDC_Only_v1780.Encode(ex); err != nil {
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

		// Extension group 17:
		if hasExtGroup17 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "supportedBandCombinationList-v1790", Optional: true},
					{Name: "supportedBandCombinationList-UplinkTxSwitch-v1790", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.SupportedBandCombinationList_v1790 != nil, ie.SupportedBandCombinationList_UplinkTxSwitch_v1790 != nil}); err != nil {
				return err
			}

			if ie.SupportedBandCombinationList_v1790 != nil {
				if err := ie.SupportedBandCombinationList_v1790.Encode(ex); err != nil {
					return err
				}
			}

			if ie.SupportedBandCombinationList_UplinkTxSwitch_v1790 != nil {
				if err := ie.SupportedBandCombinationList_UplinkTxSwitch_v1790.Encode(ex); err != nil {
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
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.SupportedBandCombinationList_v1800 != nil, ie.SupportedBandCombinationList_UplinkTxSwitch_v1800 != nil}); err != nil {
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
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.SupportedBandCombinationList_v1900 != nil, ie.SupportedBandCombinationList_UplinkTxSwitch_v1900 != nil}); err != nil {
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
			extValues = append(extValues, ex.Bytes())
		}

		if err := seq.EncodeExtensionAdditions(extPresence, extValues); err != nil {
			return err
		}
	}

	return nil
}

func (ie *RF_ParametersMRDC) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rFParametersMRDCConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. supportedBandCombinationList: ref
	{
		if seq.IsComponentPresent(0) {
			ie.SupportedBandCombinationList = new(BandCombinationList)
			if err := ie.SupportedBandCombinationList.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. appliedFreqBandListFilter: ref
	{
		if seq.IsComponentPresent(1) {
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
				{Name: "srs-SwitchingTimeRequested", Optional: true},
				{Name: "supportedBandCombinationList-v1540", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(rFParametersMRDCExtSrsSwitchingTimeRequestedConstraints)
			if err != nil {
				return err
			}
			ie.Srs_SwitchingTimeRequested = &v
		}

		if groupSeq.IsComponentPresent(1) {
			ie.SupportedBandCombinationList_v1540 = new(BandCombinationList_v1540)
			if err := ie.SupportedBandCombinationList_v1540.Decode(dx); err != nil {
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
				{Name: "supportedBandCombinationListNEDC-Only", Optional: true},
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

		if groupSeq.IsComponentPresent(1) {
			ie.SupportedBandCombinationListNEDC_Only = new(BandCombinationList)
			if err := ie.SupportedBandCombinationListNEDC_Only.Decode(dx); err != nil {
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
				{Name: "supportedBandCombinationList-v1570", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.SupportedBandCombinationList_v1570 = new(BandCombinationList_v1570)
			if err := ie.SupportedBandCombinationList_v1570.Decode(dx); err != nil {
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
				{Name: "supportedBandCombinationList-v1580", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.SupportedBandCombinationList_v1580 = new(BandCombinationList_v1580)
			if err := ie.SupportedBandCombinationList_v1580.Decode(dx); err != nil {
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
				{Name: "supportedBandCombinationList-v1590", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.SupportedBandCombinationList_v1590 = new(BandCombinationList_v1590)
			if err := ie.SupportedBandCombinationList_v1590.Decode(dx); err != nil {
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
				{Name: "supportedBandCombinationListNEDC-Only-v15a0", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.SupportedBandCombinationListNEDC_Only_V15a0 = &struct {
				SupportedBandCombinationList_v1540 *BandCombinationList_v1540
				SupportedBandCombinationList_v1560 *BandCombinationList_v1560
				SupportedBandCombinationList_v1570 *BandCombinationList_v1570
				SupportedBandCombinationList_v1580 *BandCombinationList_v1580
				SupportedBandCombinationList_v1590 *BandCombinationList_v1590
			}{}
			c := ie.SupportedBandCombinationListNEDC_Only_V15a0
			rFParametersMRDCExtSupportedBandCombinationListNEDCOnlyV15a0Seq := dx.NewSequenceDecoder(rFParametersMRDCExtSupportedBandCombinationListNEDCOnlyV15a0Constraints)
			if err := rFParametersMRDCExtSupportedBandCombinationListNEDCOnlyV15a0Seq.DecodePreamble(); err != nil {
				return err
			}
			if rFParametersMRDCExtSupportedBandCombinationListNEDCOnlyV15a0Seq.IsComponentPresent(0) {
				c.SupportedBandCombinationList_v1540 = new(BandCombinationList_v1540)
				if err := (*c.SupportedBandCombinationList_v1540).Decode(dx); err != nil {
					return err
				}
			}
			if rFParametersMRDCExtSupportedBandCombinationListNEDCOnlyV15a0Seq.IsComponentPresent(1) {
				c.SupportedBandCombinationList_v1560 = new(BandCombinationList_v1560)
				if err := (*c.SupportedBandCombinationList_v1560).Decode(dx); err != nil {
					return err
				}
			}
			if rFParametersMRDCExtSupportedBandCombinationListNEDCOnlyV15a0Seq.IsComponentPresent(2) {
				c.SupportedBandCombinationList_v1570 = new(BandCombinationList_v1570)
				if err := (*c.SupportedBandCombinationList_v1570).Decode(dx); err != nil {
					return err
				}
			}
			if rFParametersMRDCExtSupportedBandCombinationListNEDCOnlyV15a0Seq.IsComponentPresent(3) {
				c.SupportedBandCombinationList_v1580 = new(BandCombinationList_v1580)
				if err := (*c.SupportedBandCombinationList_v1580).Decode(dx); err != nil {
					return err
				}
			}
			if rFParametersMRDCExtSupportedBandCombinationListNEDCOnlyV15a0Seq.IsComponentPresent(4) {
				c.SupportedBandCombinationList_v1590 = new(BandCombinationList_v1590)
				if err := (*c.SupportedBandCombinationList_v1590).Decode(dx); err != nil {
					return err
				}
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
				{Name: "supportedBandCombinationList-v1610", Optional: true},
				{Name: "supportedBandCombinationListNEDC-Only-v1610", Optional: true},
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
			ie.SupportedBandCombinationListNEDC_Only_v1610 = new(BandCombinationList_v1610)
			if err := ie.SupportedBandCombinationListNEDC_Only_v1610.Decode(dx); err != nil {
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

	// Extension group 8:
	if seq.IsExtensionPresent(8) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "supportedBandCombinationList-v1630", Optional: true},
				{Name: "supportedBandCombinationListNEDC-Only-v1630", Optional: true},
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
			ie.SupportedBandCombinationListNEDC_Only_v1630 = new(BandCombinationList_v1630)
			if err := ie.SupportedBandCombinationListNEDC_Only_v1630.Decode(dx); err != nil {
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

	// Extension group 9:
	if seq.IsExtensionPresent(9) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "supportedBandCombinationList-v1640", Optional: true},
				{Name: "supportedBandCombinationListNEDC-Only-v1640", Optional: true},
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
			ie.SupportedBandCombinationListNEDC_Only_v1640 = new(BandCombinationList_v1640)
			if err := ie.SupportedBandCombinationListNEDC_Only_v1640.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(2) {
			ie.SupportedBandCombinationList_UplinkTxSwitch_v1640 = new(BandCombinationList_UplinkTxSwitch_v1640)
			if err := ie.SupportedBandCombinationList_UplinkTxSwitch_v1640.Decode(dx); err != nil {
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

	// Extension group 11:
	if seq.IsExtensionPresent(11) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "supportedBandCombinationList-v1700", Optional: true},
				{Name: "supportedBandCombinationList-UplinkTxSwitch-v1700", Optional: true},
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
	}

	// Extension group 12:
	if seq.IsExtensionPresent(12) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "supportedBandCombinationList-v1720", Optional: true},
				{Name: "supportedBandCombinationListNEDC-Only-v1720", Optional: true},
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
			ie.SupportedBandCombinationListNEDC_Only_v1720 = &struct {
				SupportedBandCombinationList_v1700 *BandCombinationList_v1700
				SupportedBandCombinationList_v1720 *BandCombinationList_v1720
			}{}
			c := ie.SupportedBandCombinationListNEDC_Only_v1720
			rFParametersMRDCExtSupportedBandCombinationListNEDCOnlyV1720Seq := dx.NewSequenceDecoder(rFParametersMRDCExtSupportedBandCombinationListNEDCOnlyV1720Constraints)
			if err := rFParametersMRDCExtSupportedBandCombinationListNEDCOnlyV1720Seq.DecodePreamble(); err != nil {
				return err
			}
			if rFParametersMRDCExtSupportedBandCombinationListNEDCOnlyV1720Seq.IsComponentPresent(0) {
				c.SupportedBandCombinationList_v1700 = new(BandCombinationList_v1700)
				if err := (*c.SupportedBandCombinationList_v1700).Decode(dx); err != nil {
					return err
				}
			}
			if rFParametersMRDCExtSupportedBandCombinationListNEDCOnlyV1720Seq.IsComponentPresent(1) {
				c.SupportedBandCombinationList_v1720 = new(BandCombinationList_v1720)
				if err := (*c.SupportedBandCombinationList_v1720).Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(2) {
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
				{Name: "supportedBandCombinationListNEDC-Only-v1730", Optional: true},
				{Name: "supportedBandCombinationList-UplinkTxSwitch-v1730", Optional: true},
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
			ie.SupportedBandCombinationListNEDC_Only_v1730 = new(BandCombinationList_v1730)
			if err := ie.SupportedBandCombinationListNEDC_Only_v1730.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(2) {
			ie.SupportedBandCombinationList_UplinkTxSwitch_v1730 = new(BandCombinationList_UplinkTxSwitch_v1730)
			if err := ie.SupportedBandCombinationList_UplinkTxSwitch_v1730.Decode(dx); err != nil {
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
				{Name: "supportedBandCombinationListNEDC-Only-v1740", Optional: true},
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
			ie.SupportedBandCombinationListNEDC_Only_v1740 = new(BandCombinationList_v1740)
			if err := ie.SupportedBandCombinationListNEDC_Only_v1740.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(2) {
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

	// Extension group 16:
	if seq.IsExtensionPresent(16) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "supportedBandCombinationList-v1780", Optional: true},
				{Name: "supportedBandCombinationListNEDC-Only-v1780", Optional: true},
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
			ie.SupportedBandCombinationListNEDC_Only_v1780 = new(BandCombinationList_v1780)
			if err := ie.SupportedBandCombinationListNEDC_Only_v1780.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(2) {
			ie.SupportedBandCombinationList_UplinkTxSwitch_v1780 = new(BandCombinationList_UplinkTxSwitch_v1780)
			if err := ie.SupportedBandCombinationList_UplinkTxSwitch_v1780.Decode(dx); err != nil {
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
				{Name: "supportedBandCombinationList-v1790", Optional: true},
				{Name: "supportedBandCombinationList-UplinkTxSwitch-v1790", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.SupportedBandCombinationList_v1790 = new(BandCombinationList_v1790)
			if err := ie.SupportedBandCombinationList_v1790.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.SupportedBandCombinationList_UplinkTxSwitch_v1790 = new(BandCombinationList_UplinkTxSwitch_v1790)
			if err := ie.SupportedBandCombinationList_UplinkTxSwitch_v1790.Decode(dx); err != nil {
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
	}

	return nil
}
