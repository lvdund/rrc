package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type VarMobilityHistoryReport_r17 struct {
	VisitedCellInfoList_r16         VisitedCellInfoList_r16    `madatory`
	VisitedPSCellInfoListReport_r17 *VisitedPSCellInfoList_r17 `optional`
}

func (ie *VarMobilityHistoryReport_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.VisitedPSCellInfoListReport_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.VisitedCellInfoList_r16.Encode(w); err != nil {
		return utils.WrapError("Encode VisitedCellInfoList_r16", err)
	}
	if ie.VisitedPSCellInfoListReport_r17 != nil {
		if err = ie.VisitedPSCellInfoListReport_r17.Encode(w); err != nil {
			return utils.WrapError("Encode VisitedPSCellInfoListReport_r17", err)
		}
	}
	return nil
}

func (ie *VarMobilityHistoryReport_r17) Decode(r *aper.AperReader) error {
	var err error
	var VisitedPSCellInfoListReport_r17Present bool
	if VisitedPSCellInfoListReport_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.VisitedCellInfoList_r16.Decode(r); err != nil {
		return utils.WrapError("Decode VisitedCellInfoList_r16", err)
	}
	if VisitedPSCellInfoListReport_r17Present {
		ie.VisitedPSCellInfoListReport_r17 = new(VisitedPSCellInfoList_r17)
		if err = ie.VisitedPSCellInfoListReport_r17.Decode(r); err != nil {
			return utils.WrapError("Decode VisitedPSCellInfoListReport_r17", err)
		}
	}
	return nil
}
