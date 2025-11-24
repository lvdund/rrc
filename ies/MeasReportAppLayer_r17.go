package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasReportAppLayer_r17 struct {
	MeasConfigAppLayerId_r17        MeasConfigAppLayerId_r17                          `madatory`
	MeasReportAppLayerContainer_r17 *[]byte                                           `optional`
	AppLayerSessionStatus_r17       *MeasReportAppLayer_r17_appLayerSessionStatus_r17 `optional`
	Ran_VisibleMeasurements_r17     *RAN_VisibleMeasurements_r17                      `optional`
}

func (ie *MeasReportAppLayer_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.MeasReportAppLayerContainer_r17 != nil, ie.AppLayerSessionStatus_r17 != nil, ie.Ran_VisibleMeasurements_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.MeasConfigAppLayerId_r17.Encode(w); err != nil {
		return utils.WrapError("Encode MeasConfigAppLayerId_r17", err)
	}
	if ie.MeasReportAppLayerContainer_r17 != nil {
		if err = w.WriteOctetString(*ie.MeasReportAppLayerContainer_r17, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode MeasReportAppLayerContainer_r17", err)
		}
	}
	if ie.AppLayerSessionStatus_r17 != nil {
		if err = ie.AppLayerSessionStatus_r17.Encode(w); err != nil {
			return utils.WrapError("Encode AppLayerSessionStatus_r17", err)
		}
	}
	if ie.Ran_VisibleMeasurements_r17 != nil {
		if err = ie.Ran_VisibleMeasurements_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Ran_VisibleMeasurements_r17", err)
		}
	}
	return nil
}

func (ie *MeasReportAppLayer_r17) Decode(r *uper.UperReader) error {
	var err error
	var MeasReportAppLayerContainer_r17Present bool
	if MeasReportAppLayerContainer_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var AppLayerSessionStatus_r17Present bool
	if AppLayerSessionStatus_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ran_VisibleMeasurements_r17Present bool
	if Ran_VisibleMeasurements_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.MeasConfigAppLayerId_r17.Decode(r); err != nil {
		return utils.WrapError("Decode MeasConfigAppLayerId_r17", err)
	}
	if MeasReportAppLayerContainer_r17Present {
		var tmp_os_MeasReportAppLayerContainer_r17 []byte
		if tmp_os_MeasReportAppLayerContainer_r17, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode MeasReportAppLayerContainer_r17", err)
		}
		ie.MeasReportAppLayerContainer_r17 = &tmp_os_MeasReportAppLayerContainer_r17
	}
	if AppLayerSessionStatus_r17Present {
		ie.AppLayerSessionStatus_r17 = new(MeasReportAppLayer_r17_appLayerSessionStatus_r17)
		if err = ie.AppLayerSessionStatus_r17.Decode(r); err != nil {
			return utils.WrapError("Decode AppLayerSessionStatus_r17", err)
		}
	}
	if Ran_VisibleMeasurements_r17Present {
		ie.Ran_VisibleMeasurements_r17 = new(RAN_VisibleMeasurements_r17)
		if err = ie.Ran_VisibleMeasurements_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Ran_VisibleMeasurements_r17", err)
		}
	}
	return nil
}
