package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	CSI_RS_ResourceMapping_density_Choice_nothing uint64 = iota
	CSI_RS_ResourceMapping_density_Choice_Dot5
	CSI_RS_ResourceMapping_density_Choice_One
	CSI_RS_ResourceMapping_density_Choice_Three
	CSI_RS_ResourceMapping_density_Choice_Spare
)

type CSI_RS_ResourceMapping_density struct {
	Choice uint64
	Dot5   *CSI_RS_ResourceMapping_density_dot5
	One    aper.NULL `madatory`
	Three  aper.NULL `madatory`
	Spare  aper.NULL `madatory`
}

func (ie *CSI_RS_ResourceMapping_density) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CSI_RS_ResourceMapping_density_Choice_Dot5:
		if err = ie.Dot5.Encode(w); err != nil {
			err = utils.WrapError("Encode Dot5", err)
		}
	case CSI_RS_ResourceMapping_density_Choice_One:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode One", err)
		}
	case CSI_RS_ResourceMapping_density_Choice_Three:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode Three", err)
		}
	case CSI_RS_ResourceMapping_density_Choice_Spare:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode Spare", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *CSI_RS_ResourceMapping_density) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CSI_RS_ResourceMapping_density_Choice_Dot5:
		ie.Dot5 = new(CSI_RS_ResourceMapping_density_dot5)
		if err = ie.Dot5.Decode(r); err != nil {
			return utils.WrapError("Decode Dot5", err)
		}
	case CSI_RS_ResourceMapping_density_Choice_One:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode One", err)
		}
	case CSI_RS_ResourceMapping_density_Choice_Three:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode Three", err)
		}
	case CSI_RS_ResourceMapping_density_Choice_Spare:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode Spare", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
