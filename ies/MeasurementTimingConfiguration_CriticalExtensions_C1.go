package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	MeasurementTimingConfiguration_CriticalExtensions_C1_Choice_nothing uint64 = iota
	MeasurementTimingConfiguration_CriticalExtensions_C1_Choice_MeasTimingConf
	MeasurementTimingConfiguration_CriticalExtensions_C1_Choice_Spare3
	MeasurementTimingConfiguration_CriticalExtensions_C1_Choice_Spare2
	MeasurementTimingConfiguration_CriticalExtensions_C1_Choice_Spare1
)

type MeasurementTimingConfiguration_CriticalExtensions_C1 struct {
	Choice         uint64
	MeasTimingConf *MeasurementTimingConfiguration_IEs
	Spare3         aper.NULL `madatory`
	Spare2         aper.NULL `madatory`
	Spare1         aper.NULL `madatory`
}

func (ie *MeasurementTimingConfiguration_CriticalExtensions_C1) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MeasurementTimingConfiguration_CriticalExtensions_C1_Choice_MeasTimingConf:
		if err = ie.MeasTimingConf.Encode(w); err != nil {
			err = utils.WrapError("Encode MeasTimingConf", err)
		}
	case MeasurementTimingConfiguration_CriticalExtensions_C1_Choice_Spare3:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode Spare3", err)
		}
	case MeasurementTimingConfiguration_CriticalExtensions_C1_Choice_Spare2:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode Spare2", err)
		}
	case MeasurementTimingConfiguration_CriticalExtensions_C1_Choice_Spare1:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode Spare1", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *MeasurementTimingConfiguration_CriticalExtensions_C1) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MeasurementTimingConfiguration_CriticalExtensions_C1_Choice_MeasTimingConf:
		ie.MeasTimingConf = new(MeasurementTimingConfiguration_IEs)
		if err = ie.MeasTimingConf.Decode(r); err != nil {
			return utils.WrapError("Decode MeasTimingConf", err)
		}
	case MeasurementTimingConfiguration_CriticalExtensions_C1_Choice_Spare3:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode Spare3", err)
		}
	case MeasurementTimingConfiguration_CriticalExtensions_C1_Choice_Spare2:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode Spare2", err)
		}
	case MeasurementTimingConfiguration_CriticalExtensions_C1_Choice_Spare1:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode Spare1", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
