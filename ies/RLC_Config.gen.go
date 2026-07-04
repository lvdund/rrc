// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RLC-Config (line 14048).

var rLCConfigConstraints = per.ChoiceConstraints{
	Extensible: true,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "am"},
		{Name: "um-Bi-Directional"},
		{Name: "um-Uni-Directional-UL"},
		{Name: "um-Uni-Directional-DL"},
	},
}

const (
	RLC_Config_Am                    = 0
	RLC_Config_Um_Bi_Directional     = 1
	RLC_Config_Um_Uni_Directional_UL = 2
	RLC_Config_Um_Uni_Directional_DL = 3
)

type RLC_Config struct {
	Choice int
	Am     *struct {
		Ul_AM_RLC UL_AM_RLC
		Dl_AM_RLC DL_AM_RLC
	}
	Um_Bi_Directional *struct {
		Ul_UM_RLC UL_UM_RLC
		Dl_UM_RLC DL_UM_RLC
	}
	Um_Uni_Directional_UL *struct{ Ul_UM_RLC UL_UM_RLC }
	Um_Uni_Directional_DL *struct{ Dl_UM_RLC DL_UM_RLC }
}

func (ie *RLC_Config) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(rLCConfigConstraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case RLC_Config_Am:
		c := (*ie.Am)
		if err := c.Ul_AM_RLC.Encode(e); err != nil {
			return err
		}
		if err := c.Dl_AM_RLC.Encode(e); err != nil {
			return err
		}
	case RLC_Config_Um_Bi_Directional:
		c := (*ie.Um_Bi_Directional)
		if err := c.Ul_UM_RLC.Encode(e); err != nil {
			return err
		}
		if err := c.Dl_UM_RLC.Encode(e); err != nil {
			return err
		}
	case RLC_Config_Um_Uni_Directional_UL:
		c := (*ie.Um_Uni_Directional_UL)
		if err := c.Ul_UM_RLC.Encode(e); err != nil {
			return err
		}
	case RLC_Config_Um_Uni_Directional_DL:
		c := (*ie.Um_Uni_Directional_DL)
		if err := c.Dl_UM_RLC.Encode(e); err != nil {
			return err
		}
	default:
		// Extension alternative: not modeled; bail out.
		return &per.ConstraintViolationError{
			TypeName:   "RLC-Config",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *RLC_Config) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(rLCConfigConstraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "RLC-Config", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case RLC_Config_Am:
		ie.Am = &struct {
			Ul_AM_RLC UL_AM_RLC
			Dl_AM_RLC DL_AM_RLC
		}{}
		c := (*ie.Am)
		{
			if err := c.Ul_AM_RLC.Decode(d); err != nil {
				return err
			}
		}
		{
			if err := c.Dl_AM_RLC.Decode(d); err != nil {
				return err
			}
		}
	case RLC_Config_Um_Bi_Directional:
		ie.Um_Bi_Directional = &struct {
			Ul_UM_RLC UL_UM_RLC
			Dl_UM_RLC DL_UM_RLC
		}{}
		c := (*ie.Um_Bi_Directional)
		{
			if err := c.Ul_UM_RLC.Decode(d); err != nil {
				return err
			}
		}
		{
			if err := c.Dl_UM_RLC.Decode(d); err != nil {
				return err
			}
		}
	case RLC_Config_Um_Uni_Directional_UL:
		ie.Um_Uni_Directional_UL = &struct{ Ul_UM_RLC UL_UM_RLC }{}
		c := (*ie.Um_Uni_Directional_UL)
		{
			if err := c.Ul_UM_RLC.Decode(d); err != nil {
				return err
			}
		}
	case RLC_Config_Um_Uni_Directional_DL:
		ie.Um_Uni_Directional_DL = &struct{ Dl_UM_RLC DL_UM_RLC }{}
		c := (*ie.Um_Uni_Directional_DL)
		{
			if err := c.Dl_UM_RLC.Decode(d); err != nil {
				return err
			}
		}
	default:
		return &per.DecodeError{TypeName: "RLC-Config", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
