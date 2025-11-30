package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	ControlResourceSet_cce_REG_MappingType_Choice_nothing uint64 = iota
	ControlResourceSet_cce_REG_MappingType_Choice_Interleaved
	ControlResourceSet_cce_REG_MappingType_Choice_NonInterleaved
)

type ControlResourceSet_cce_REG_MappingType struct {
	Choice         uint64
	Interleaved    *ControlResourceSet_cce_REG_MappingType_interleaved
	NonInterleaved aper.NULL `madatory`
}

func (ie *ControlResourceSet_cce_REG_MappingType) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case ControlResourceSet_cce_REG_MappingType_Choice_Interleaved:
		if err = ie.Interleaved.Encode(w); err != nil {
			err = utils.WrapError("Encode Interleaved", err)
		}
	case ControlResourceSet_cce_REG_MappingType_Choice_NonInterleaved:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode NonInterleaved", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *ControlResourceSet_cce_REG_MappingType) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case ControlResourceSet_cce_REG_MappingType_Choice_Interleaved:
		ie.Interleaved = new(ControlResourceSet_cce_REG_MappingType_interleaved)
		if err = ie.Interleaved.Decode(r); err != nil {
			return utils.WrapError("Decode Interleaved", err)
		}
	case ControlResourceSet_cce_REG_MappingType_Choice_NonInterleaved:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode NonInterleaved", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
