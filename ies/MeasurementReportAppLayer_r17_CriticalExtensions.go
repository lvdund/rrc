package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	MeasurementReportAppLayer_r17_CriticalExtensions_Choice_nothing uint64 = iota
	MeasurementReportAppLayer_r17_CriticalExtensions_Choice_MeasurementReportAppLayer_r17
	MeasurementReportAppLayer_r17_CriticalExtensions_Choice_CriticalExtensionsFuture
)

type MeasurementReportAppLayer_r17_CriticalExtensions struct {
	Choice                        uint64
	MeasurementReportAppLayer_r17 *MeasurementReportAppLayer_r17_IEs
	CriticalExtensionsFuture      interface{} `madatory`
}

func (ie *MeasurementReportAppLayer_r17_CriticalExtensions) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MeasurementReportAppLayer_r17_CriticalExtensions_Choice_MeasurementReportAppLayer_r17:
		if err = ie.MeasurementReportAppLayer_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode MeasurementReportAppLayer_r17", err)
		}
	case MeasurementReportAppLayer_r17_CriticalExtensions_Choice_CriticalExtensionsFuture:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *MeasurementReportAppLayer_r17_CriticalExtensions) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MeasurementReportAppLayer_r17_CriticalExtensions_Choice_MeasurementReportAppLayer_r17:
		ie.MeasurementReportAppLayer_r17 = new(MeasurementReportAppLayer_r17_IEs)
		if err = ie.MeasurementReportAppLayer_r17.Decode(r); err != nil {
			return utils.WrapError("Decode MeasurementReportAppLayer_r17", err)
		}
	case MeasurementReportAppLayer_r17_CriticalExtensions_Choice_CriticalExtensionsFuture:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
