package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MeasurementReport_CriticalExtensions_Choice_nothing uint64 = iota
	MeasurementReport_CriticalExtensions_Choice_MeasurementReport
	MeasurementReport_CriticalExtensions_Choice_CriticalExtensionsFuture
)

type MeasurementReport_CriticalExtensions struct {
	Choice                   uint64
	MeasurementReport        *MeasurementReport_IEs
	CriticalExtensionsFuture interface{} `madatory`
}

func (ie *MeasurementReport_CriticalExtensions) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MeasurementReport_CriticalExtensions_Choice_MeasurementReport:
		if err = ie.MeasurementReport.Encode(w); err != nil {
			err = utils.WrapError("Encode MeasurementReport", err)
		}
	case MeasurementReport_CriticalExtensions_Choice_CriticalExtensionsFuture:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *MeasurementReport_CriticalExtensions) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MeasurementReport_CriticalExtensions_Choice_MeasurementReport:
		ie.MeasurementReport = new(MeasurementReport_IEs)
		if err = ie.MeasurementReport.Decode(r); err != nil {
			return utils.WrapError("Decode MeasurementReport", err)
		}
	case MeasurementReport_CriticalExtensions_Choice_CriticalExtensionsFuture:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
