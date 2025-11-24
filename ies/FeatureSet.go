package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	FeatureSet_Choice_nothing uint64 = iota
	FeatureSet_Choice_Eutra
	FeatureSet_Choice_Nr
)

type FeatureSet struct {
	Choice uint64
	Eutra  *FeatureSet_eutra
	Nr     *FeatureSet_nr
}

func (ie *FeatureSet) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case FeatureSet_Choice_Eutra:
		if err = ie.Eutra.Encode(w); err != nil {
			err = utils.WrapError("Encode Eutra", err)
		}
	case FeatureSet_Choice_Nr:
		if err = ie.Nr.Encode(w); err != nil {
			err = utils.WrapError("Encode Nr", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *FeatureSet) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case FeatureSet_Choice_Eutra:
		ie.Eutra = new(FeatureSet_eutra)
		if err = ie.Eutra.Decode(r); err != nil {
			return utils.WrapError("Decode Eutra", err)
		}
	case FeatureSet_Choice_Nr:
		ie.Nr = new(FeatureSet_nr)
		if err = ie.Nr.Decode(r); err != nil {
			return utils.WrapError("Decode Nr", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
