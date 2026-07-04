// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: Dummy-PDCCH-RACH-DL-Info-r18 (line 19859).

var dummyPDCCHRACHDLInfoR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "notSupported"},
		{Name: "supported"},
	},
}

const (
	Dummy_PDCCH_RACH_DL_Info_r18_NotSupported = 0
	Dummy_PDCCH_RACH_DL_Info_r18_Supported    = 1
)

var dummyPDCCHRACHDLInfoR18SupportedConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "pdcch-RACH-AffectedBands-r18"},
		{Name: "pdcch-RACH-SwitchingTimeList-r18", Optional: true},
		{Name: "pdcch-RACH-PrepTime-r18", Optional: true},
	},
}

const (
	Dummy_PDCCH_RACH_DL_Info_r18_Supported_Pdcch_RACH_AffectedBands_r18_NoIntrruption = 0
	Dummy_PDCCH_RACH_DL_Info_r18_Supported_Pdcch_RACH_AffectedBands_r18_Interruption  = 1
)

var dummyPDCCHRACHDLInfoR18SupportedPdcchRACHAffectedBandsR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Dummy_PDCCH_RACH_DL_Info_r18_Supported_Pdcch_RACH_AffectedBands_r18_NoIntrruption, Dummy_PDCCH_RACH_DL_Info_r18_Supported_Pdcch_RACH_AffectedBands_r18_Interruption},
}

const (
	Dummy_PDCCH_RACH_DL_Info_r18_Supported_Pdcch_RACH_SwitchingTimeList_r18_Ms0      = 0
	Dummy_PDCCH_RACH_DL_Info_r18_Supported_Pdcch_RACH_SwitchingTimeList_r18_Ms0dot25 = 1
	Dummy_PDCCH_RACH_DL_Info_r18_Supported_Pdcch_RACH_SwitchingTimeList_r18_Ms0dot5  = 2
	Dummy_PDCCH_RACH_DL_Info_r18_Supported_Pdcch_RACH_SwitchingTimeList_r18_Ms1      = 3
	Dummy_PDCCH_RACH_DL_Info_r18_Supported_Pdcch_RACH_SwitchingTimeList_r18_Ms2      = 4
)

var dummyPDCCHRACHDLInfoR18SupportedPdcchRACHSwitchingTimeListR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Dummy_PDCCH_RACH_DL_Info_r18_Supported_Pdcch_RACH_SwitchingTimeList_r18_Ms0, Dummy_PDCCH_RACH_DL_Info_r18_Supported_Pdcch_RACH_SwitchingTimeList_r18_Ms0dot25, Dummy_PDCCH_RACH_DL_Info_r18_Supported_Pdcch_RACH_SwitchingTimeList_r18_Ms0dot5, Dummy_PDCCH_RACH_DL_Info_r18_Supported_Pdcch_RACH_SwitchingTimeList_r18_Ms1, Dummy_PDCCH_RACH_DL_Info_r18_Supported_Pdcch_RACH_SwitchingTimeList_r18_Ms2},
}

const (
	Dummy_PDCCH_RACH_DL_Info_r18_Supported_Pdcch_RACH_PrepTime_r18_Ms1  = 0
	Dummy_PDCCH_RACH_DL_Info_r18_Supported_Pdcch_RACH_PrepTime_r18_Ms3  = 1
	Dummy_PDCCH_RACH_DL_Info_r18_Supported_Pdcch_RACH_PrepTime_r18_Ms5  = 2
	Dummy_PDCCH_RACH_DL_Info_r18_Supported_Pdcch_RACH_PrepTime_r18_Ms10 = 3
)

var dummyPDCCHRACHDLInfoR18SupportedPdcchRACHPrepTimeR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Dummy_PDCCH_RACH_DL_Info_r18_Supported_Pdcch_RACH_PrepTime_r18_Ms1, Dummy_PDCCH_RACH_DL_Info_r18_Supported_Pdcch_RACH_PrepTime_r18_Ms3, Dummy_PDCCH_RACH_DL_Info_r18_Supported_Pdcch_RACH_PrepTime_r18_Ms5, Dummy_PDCCH_RACH_DL_Info_r18_Supported_Pdcch_RACH_PrepTime_r18_Ms10},
}

type Dummy_PDCCH_RACH_DL_Info_r18 struct {
	Choice       int
	NotSupported *struct{}
	Supported    *struct {
		Pdcch_RACH_AffectedBands_r18     int64
		Pdcch_RACH_SwitchingTimeList_r18 *int64
		Pdcch_RACH_PrepTime_r18          *int64
	}
}

func (ie *Dummy_PDCCH_RACH_DL_Info_r18) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(dummyPDCCHRACHDLInfoR18Constraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case Dummy_PDCCH_RACH_DL_Info_r18_NotSupported:
		if err := e.EncodeNull(); err != nil {
			return err
		}
	case Dummy_PDCCH_RACH_DL_Info_r18_Supported:
		c := (*ie.Supported)
		dummyPDCCHRACHDLInfoR18SupportedSeq := e.NewSequenceEncoder(dummyPDCCHRACHDLInfoR18SupportedConstraints)
		if err := dummyPDCCHRACHDLInfoR18SupportedSeq.EncodePreamble([]bool{c.Pdcch_RACH_SwitchingTimeList_r18 != nil, c.Pdcch_RACH_PrepTime_r18 != nil}); err != nil {
			return err
		}
		if err := e.EncodeEnumerated(c.Pdcch_RACH_AffectedBands_r18, dummyPDCCHRACHDLInfoR18SupportedPdcchRACHAffectedBandsR18Constraints); err != nil {
			return err
		}
		if c.Pdcch_RACH_SwitchingTimeList_r18 != nil {
			if err := e.EncodeEnumerated((*c.Pdcch_RACH_SwitchingTimeList_r18), dummyPDCCHRACHDLInfoR18SupportedPdcchRACHSwitchingTimeListR18Constraints); err != nil {
				return err
			}
		}
		if c.Pdcch_RACH_PrepTime_r18 != nil {
			if err := e.EncodeEnumerated((*c.Pdcch_RACH_PrepTime_r18), dummyPDCCHRACHDLInfoR18SupportedPdcchRACHPrepTimeR18Constraints); err != nil {
				return err
			}
		}
	default:
		// Choice index out of root range.
		return &per.ConstraintViolationError{
			TypeName:   "Dummy-PDCCH-RACH-DL-Info-r18",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *Dummy_PDCCH_RACH_DL_Info_r18) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(dummyPDCCHRACHDLInfoR18Constraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "Dummy-PDCCH-RACH-DL-Info-r18", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case Dummy_PDCCH_RACH_DL_Info_r18_NotSupported:
		ie.NotSupported = &struct{}{}
		if err := d.DecodeNull(); err != nil {
			return err
		}
	case Dummy_PDCCH_RACH_DL_Info_r18_Supported:
		ie.Supported = &struct {
			Pdcch_RACH_AffectedBands_r18     int64
			Pdcch_RACH_SwitchingTimeList_r18 *int64
			Pdcch_RACH_PrepTime_r18          *int64
		}{}
		c := (*ie.Supported)
		dummyPDCCHRACHDLInfoR18SupportedSeq := d.NewSequenceDecoder(dummyPDCCHRACHDLInfoR18SupportedConstraints)
		if err := dummyPDCCHRACHDLInfoR18SupportedSeq.DecodePreamble(); err != nil {
			return err
		}
		{
			v, err := d.DecodeEnumerated(dummyPDCCHRACHDLInfoR18SupportedPdcchRACHAffectedBandsR18Constraints)
			if err != nil {
				return err
			}
			c.Pdcch_RACH_AffectedBands_r18 = v
		}
		if dummyPDCCHRACHDLInfoR18SupportedSeq.IsComponentPresent(1) {
			c.Pdcch_RACH_SwitchingTimeList_r18 = new(int64)
			v, err := d.DecodeEnumerated(dummyPDCCHRACHDLInfoR18SupportedPdcchRACHSwitchingTimeListR18Constraints)
			if err != nil {
				return err
			}
			(*c.Pdcch_RACH_SwitchingTimeList_r18) = v
		}
		if dummyPDCCHRACHDLInfoR18SupportedSeq.IsComponentPresent(2) {
			c.Pdcch_RACH_PrepTime_r18 = new(int64)
			v, err := d.DecodeEnumerated(dummyPDCCHRACHDLInfoR18SupportedPdcchRACHPrepTimeR18Constraints)
			if err != nil {
				return err
			}
			(*c.Pdcch_RACH_PrepTime_r18) = v
		}
	default:
		return &per.DecodeError{TypeName: "Dummy-PDCCH-RACH-DL-Info-r18", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
