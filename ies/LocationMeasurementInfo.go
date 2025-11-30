package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	LocationMeasurementInfo_Choice_nothing uint64 = iota
	LocationMeasurementInfo_Choice_Eutra_RSTD
	LocationMeasurementInfo_Choice_Eutra_FineTimingDetection
	LocationMeasurementInfo_Choice_Nr_PRS_Measurement_r16
)

type LocationMeasurementInfo struct {
	Choice                    uint64
	Eutra_RSTD                *EUTRA_RSTD_InfoList
	Eutra_FineTimingDetection aper.NULL `madatory,ext`
	Nr_PRS_Measurement_r16    *NR_PRS_MeasurementInfoList_r16
}

func (ie *LocationMeasurementInfo) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case LocationMeasurementInfo_Choice_Eutra_RSTD:
		if err = ie.Eutra_RSTD.Encode(w); err != nil {
			err = utils.WrapError("Encode Eutra_RSTD", err)
		}
	case LocationMeasurementInfo_Choice_Eutra_FineTimingDetection:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode Eutra_FineTimingDetection", err)
		}
	case LocationMeasurementInfo_Choice_Nr_PRS_Measurement_r16:
		if err = ie.Nr_PRS_Measurement_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode Nr_PRS_Measurement_r16", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *LocationMeasurementInfo) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case LocationMeasurementInfo_Choice_Eutra_RSTD:
		ie.Eutra_RSTD = new(EUTRA_RSTD_InfoList)
		if err = ie.Eutra_RSTD.Decode(r); err != nil {
			return utils.WrapError("Decode Eutra_RSTD", err)
		}
	case LocationMeasurementInfo_Choice_Eutra_FineTimingDetection:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode Eutra_FineTimingDetection", err)
		}
	case LocationMeasurementInfo_Choice_Nr_PRS_Measurement_r16:
		ie.Nr_PRS_Measurement_r16 = new(NR_PRS_MeasurementInfoList_r16)
		if err = ie.Nr_PRS_Measurement_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Nr_PRS_Measurement_r16", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
