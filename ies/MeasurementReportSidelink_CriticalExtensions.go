package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	MeasurementReportSidelink_CriticalExtensions_Choice_nothing uint64 = iota
	MeasurementReportSidelink_CriticalExtensions_Choice_MeasurementReportSidelink_r16
	MeasurementReportSidelink_CriticalExtensions_Choice_CriticalExtensionsFuture
)

type MeasurementReportSidelink_CriticalExtensions struct {
	Choice                        uint64
	MeasurementReportSidelink_r16 *MeasurementReportSidelink_r16_IEs
	CriticalExtensionsFuture      interface{} `madatory`
}

func (ie *MeasurementReportSidelink_CriticalExtensions) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MeasurementReportSidelink_CriticalExtensions_Choice_MeasurementReportSidelink_r16:
		if err = ie.MeasurementReportSidelink_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode MeasurementReportSidelink_r16", err)
		}
	case MeasurementReportSidelink_CriticalExtensions_Choice_CriticalExtensionsFuture:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *MeasurementReportSidelink_CriticalExtensions) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MeasurementReportSidelink_CriticalExtensions_Choice_MeasurementReportSidelink_r16:
		ie.MeasurementReportSidelink_r16 = new(MeasurementReportSidelink_r16_IEs)
		if err = ie.MeasurementReportSidelink_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MeasurementReportSidelink_r16", err)
		}
	case MeasurementReportSidelink_CriticalExtensions_Choice_CriticalExtensionsFuture:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
