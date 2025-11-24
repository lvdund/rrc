package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type AppLayerMeasParameters_r17 struct {
	Qoe_Streaming_MeasReport_r17            *AppLayerMeasParameters_r17_qoe_Streaming_MeasReport_r17            `optional`
	Qoe_MTSI_MeasReport_r17                 *AppLayerMeasParameters_r17_qoe_MTSI_MeasReport_r17                 `optional`
	Qoe_VR_MeasReport_r17                   *AppLayerMeasParameters_r17_qoe_VR_MeasReport_r17                   `optional`
	Ran_VisibleQoE_Streaming_MeasReport_r17 *AppLayerMeasParameters_r17_ran_VisibleQoE_Streaming_MeasReport_r17 `optional`
	Ran_VisibleQoE_VR_MeasReport_r17        *AppLayerMeasParameters_r17_ran_VisibleQoE_VR_MeasReport_r17        `optional`
	Ul_MeasurementReportAppLayer_Seg_r17    *AppLayerMeasParameters_r17_ul_MeasurementReportAppLayer_Seg_r17    `optional`
}

func (ie *AppLayerMeasParameters_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Qoe_Streaming_MeasReport_r17 != nil, ie.Qoe_MTSI_MeasReport_r17 != nil, ie.Qoe_VR_MeasReport_r17 != nil, ie.Ran_VisibleQoE_Streaming_MeasReport_r17 != nil, ie.Ran_VisibleQoE_VR_MeasReport_r17 != nil, ie.Ul_MeasurementReportAppLayer_Seg_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Qoe_Streaming_MeasReport_r17 != nil {
		if err = ie.Qoe_Streaming_MeasReport_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Qoe_Streaming_MeasReport_r17", err)
		}
	}
	if ie.Qoe_MTSI_MeasReport_r17 != nil {
		if err = ie.Qoe_MTSI_MeasReport_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Qoe_MTSI_MeasReport_r17", err)
		}
	}
	if ie.Qoe_VR_MeasReport_r17 != nil {
		if err = ie.Qoe_VR_MeasReport_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Qoe_VR_MeasReport_r17", err)
		}
	}
	if ie.Ran_VisibleQoE_Streaming_MeasReport_r17 != nil {
		if err = ie.Ran_VisibleQoE_Streaming_MeasReport_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Ran_VisibleQoE_Streaming_MeasReport_r17", err)
		}
	}
	if ie.Ran_VisibleQoE_VR_MeasReport_r17 != nil {
		if err = ie.Ran_VisibleQoE_VR_MeasReport_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Ran_VisibleQoE_VR_MeasReport_r17", err)
		}
	}
	if ie.Ul_MeasurementReportAppLayer_Seg_r17 != nil {
		if err = ie.Ul_MeasurementReportAppLayer_Seg_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Ul_MeasurementReportAppLayer_Seg_r17", err)
		}
	}
	return nil
}

func (ie *AppLayerMeasParameters_r17) Decode(r *uper.UperReader) error {
	var err error
	var Qoe_Streaming_MeasReport_r17Present bool
	if Qoe_Streaming_MeasReport_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Qoe_MTSI_MeasReport_r17Present bool
	if Qoe_MTSI_MeasReport_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Qoe_VR_MeasReport_r17Present bool
	if Qoe_VR_MeasReport_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ran_VisibleQoE_Streaming_MeasReport_r17Present bool
	if Ran_VisibleQoE_Streaming_MeasReport_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ran_VisibleQoE_VR_MeasReport_r17Present bool
	if Ran_VisibleQoE_VR_MeasReport_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ul_MeasurementReportAppLayer_Seg_r17Present bool
	if Ul_MeasurementReportAppLayer_Seg_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Qoe_Streaming_MeasReport_r17Present {
		ie.Qoe_Streaming_MeasReport_r17 = new(AppLayerMeasParameters_r17_qoe_Streaming_MeasReport_r17)
		if err = ie.Qoe_Streaming_MeasReport_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Qoe_Streaming_MeasReport_r17", err)
		}
	}
	if Qoe_MTSI_MeasReport_r17Present {
		ie.Qoe_MTSI_MeasReport_r17 = new(AppLayerMeasParameters_r17_qoe_MTSI_MeasReport_r17)
		if err = ie.Qoe_MTSI_MeasReport_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Qoe_MTSI_MeasReport_r17", err)
		}
	}
	if Qoe_VR_MeasReport_r17Present {
		ie.Qoe_VR_MeasReport_r17 = new(AppLayerMeasParameters_r17_qoe_VR_MeasReport_r17)
		if err = ie.Qoe_VR_MeasReport_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Qoe_VR_MeasReport_r17", err)
		}
	}
	if Ran_VisibleQoE_Streaming_MeasReport_r17Present {
		ie.Ran_VisibleQoE_Streaming_MeasReport_r17 = new(AppLayerMeasParameters_r17_ran_VisibleQoE_Streaming_MeasReport_r17)
		if err = ie.Ran_VisibleQoE_Streaming_MeasReport_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Ran_VisibleQoE_Streaming_MeasReport_r17", err)
		}
	}
	if Ran_VisibleQoE_VR_MeasReport_r17Present {
		ie.Ran_VisibleQoE_VR_MeasReport_r17 = new(AppLayerMeasParameters_r17_ran_VisibleQoE_VR_MeasReport_r17)
		if err = ie.Ran_VisibleQoE_VR_MeasReport_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Ran_VisibleQoE_VR_MeasReport_r17", err)
		}
	}
	if Ul_MeasurementReportAppLayer_Seg_r17Present {
		ie.Ul_MeasurementReportAppLayer_Seg_r17 = new(AppLayerMeasParameters_r17_ul_MeasurementReportAppLayer_Seg_r17)
		if err = ie.Ul_MeasurementReportAppLayer_Seg_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Ul_MeasurementReportAppLayer_Seg_r17", err)
		}
	}
	return nil
}
