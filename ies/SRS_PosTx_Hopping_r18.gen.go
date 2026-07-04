// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SRS-PosTx-Hopping-r18 (line 15754).

var sRSPosTxHoppingR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "srs-PosConfig-r18"},
		{Name: "bwp-r18", Optional: true},
		{Name: "inactivePosSRS-TimeAlignmentTimer-r18", Optional: true},
		{Name: "inactivePosSRS-RSRP-ChangeThreshold-r18", Optional: true},
		{Name: "srs-PosUplinkTransmissionWindowConfig-r18", Optional: true},
	},
}

var sRS_PosTx_Hopping_r18SrsPosUplinkTransmissionWindowConfigR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	SRS_PosTx_Hopping_r18_Srs_PosUplinkTransmissionWindowConfig_r18_Release = 0
	SRS_PosTx_Hopping_r18_Srs_PosUplinkTransmissionWindowConfig_r18_Setup   = 1
)

type SRS_PosTx_Hopping_r18 struct {
	Srs_PosConfig_r18                         SRS_PosConfig_r17
	Bwp_r18                                   *BWP
	InactivePosSRS_TimeAlignmentTimer_r18     *TimeAlignmentTimer
	InactivePosSRS_RSRP_ChangeThreshold_r18   *RSRP_ChangeThreshold_r17
	Srs_PosUplinkTransmissionWindowConfig_r18 *struct {
		Choice  int
		Release *struct{}
		Setup   *SRS_PosUplinkTransmissionWindowConfig_r18
	}
}

func (ie *SRS_PosTx_Hopping_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sRSPosTxHoppingR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Bwp_r18 != nil, ie.InactivePosSRS_TimeAlignmentTimer_r18 != nil, ie.InactivePosSRS_RSRP_ChangeThreshold_r18 != nil, ie.Srs_PosUplinkTransmissionWindowConfig_r18 != nil}); err != nil {
		return err
	}

	// 3. srs-PosConfig-r18: ref
	{
		if err := ie.Srs_PosConfig_r18.Encode(e); err != nil {
			return err
		}
	}

	// 4. bwp-r18: ref
	{
		if ie.Bwp_r18 != nil {
			if err := ie.Bwp_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. inactivePosSRS-TimeAlignmentTimer-r18: ref
	{
		if ie.InactivePosSRS_TimeAlignmentTimer_r18 != nil {
			if err := ie.InactivePosSRS_TimeAlignmentTimer_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. inactivePosSRS-RSRP-ChangeThreshold-r18: ref
	{
		if ie.InactivePosSRS_RSRP_ChangeThreshold_r18 != nil {
			if err := ie.InactivePosSRS_RSRP_ChangeThreshold_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 7. srs-PosUplinkTransmissionWindowConfig-r18: choice
	{
		if ie.Srs_PosUplinkTransmissionWindowConfig_r18 != nil {
			choiceEnc := e.NewChoiceEncoder(sRS_PosTx_Hopping_r18SrsPosUplinkTransmissionWindowConfigR18Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Srs_PosUplinkTransmissionWindowConfig_r18).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Srs_PosUplinkTransmissionWindowConfig_r18).Choice {
			case SRS_PosTx_Hopping_r18_Srs_PosUplinkTransmissionWindowConfig_r18_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case SRS_PosTx_Hopping_r18_Srs_PosUplinkTransmissionWindowConfig_r18_Setup:
				if err := (*ie.Srs_PosUplinkTransmissionWindowConfig_r18).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Srs_PosUplinkTransmissionWindowConfig_r18).Choice), Constraint: "index out of range"}
			}
		}
	}

	return nil
}

func (ie *SRS_PosTx_Hopping_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sRSPosTxHoppingR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. srs-PosConfig-r18: ref
	{
		if err := ie.Srs_PosConfig_r18.Decode(d); err != nil {
			return err
		}
	}

	// 4. bwp-r18: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Bwp_r18 = new(BWP)
			if err := ie.Bwp_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. inactivePosSRS-TimeAlignmentTimer-r18: ref
	{
		if seq.IsComponentPresent(2) {
			ie.InactivePosSRS_TimeAlignmentTimer_r18 = new(TimeAlignmentTimer)
			if err := ie.InactivePosSRS_TimeAlignmentTimer_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. inactivePosSRS-RSRP-ChangeThreshold-r18: ref
	{
		if seq.IsComponentPresent(3) {
			ie.InactivePosSRS_RSRP_ChangeThreshold_r18 = new(RSRP_ChangeThreshold_r17)
			if err := ie.InactivePosSRS_RSRP_ChangeThreshold_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 7. srs-PosUplinkTransmissionWindowConfig-r18: choice
	{
		if seq.IsComponentPresent(4) {
			ie.Srs_PosUplinkTransmissionWindowConfig_r18 = &struct {
				Choice  int
				Release *struct{}
				Setup   *SRS_PosUplinkTransmissionWindowConfig_r18
			}{}
			choiceDec := d.NewChoiceDecoder(sRS_PosTx_Hopping_r18SrsPosUplinkTransmissionWindowConfigR18Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Srs_PosUplinkTransmissionWindowConfig_r18).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case SRS_PosTx_Hopping_r18_Srs_PosUplinkTransmissionWindowConfig_r18_Release:
				(*ie.Srs_PosUplinkTransmissionWindowConfig_r18).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case SRS_PosTx_Hopping_r18_Srs_PosUplinkTransmissionWindowConfig_r18_Setup:
				(*ie.Srs_PosUplinkTransmissionWindowConfig_r18).Setup = new(SRS_PosUplinkTransmissionWindowConfig_r18)
				if err := (*ie.Srs_PosUplinkTransmissionWindowConfig_r18).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
