package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PDSCH_ConfigPTM_r17 struct {
	DataScramblingIdentityPDSCH_r17 *int64                                           `lb:0,ub:1023,optional`
	Dmrs_ScramblingID0_r17          *int64                                           `lb:0,ub:65535,optional`
	Pdsch_AggregationFactor_r17     *PDSCH_ConfigPTM_r17_pdsch_AggregationFactor_r17 `optional`
}

func (ie *PDSCH_ConfigPTM_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.DataScramblingIdentityPDSCH_r17 != nil, ie.Dmrs_ScramblingID0_r17 != nil, ie.Pdsch_AggregationFactor_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.DataScramblingIdentityPDSCH_r17 != nil {
		if err = w.WriteInteger(*ie.DataScramblingIdentityPDSCH_r17, &aper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
			return utils.WrapError("Encode DataScramblingIdentityPDSCH_r17", err)
		}
	}
	if ie.Dmrs_ScramblingID0_r17 != nil {
		if err = w.WriteInteger(*ie.Dmrs_ScramblingID0_r17, &aper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
			return utils.WrapError("Encode Dmrs_ScramblingID0_r17", err)
		}
	}
	if ie.Pdsch_AggregationFactor_r17 != nil {
		if err = ie.Pdsch_AggregationFactor_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Pdsch_AggregationFactor_r17", err)
		}
	}
	return nil
}

func (ie *PDSCH_ConfigPTM_r17) Decode(r *aper.AperReader) error {
	var err error
	var DataScramblingIdentityPDSCH_r17Present bool
	if DataScramblingIdentityPDSCH_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Dmrs_ScramblingID0_r17Present bool
	if Dmrs_ScramblingID0_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Pdsch_AggregationFactor_r17Present bool
	if Pdsch_AggregationFactor_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if DataScramblingIdentityPDSCH_r17Present {
		var tmp_int_DataScramblingIdentityPDSCH_r17 int64
		if tmp_int_DataScramblingIdentityPDSCH_r17, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
			return utils.WrapError("Decode DataScramblingIdentityPDSCH_r17", err)
		}
		ie.DataScramblingIdentityPDSCH_r17 = &tmp_int_DataScramblingIdentityPDSCH_r17
	}
	if Dmrs_ScramblingID0_r17Present {
		var tmp_int_Dmrs_ScramblingID0_r17 int64
		if tmp_int_Dmrs_ScramblingID0_r17, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
			return utils.WrapError("Decode Dmrs_ScramblingID0_r17", err)
		}
		ie.Dmrs_ScramblingID0_r17 = &tmp_int_Dmrs_ScramblingID0_r17
	}
	if Pdsch_AggregationFactor_r17Present {
		ie.Pdsch_AggregationFactor_r17 = new(PDSCH_ConfigPTM_r17_pdsch_AggregationFactor_r17)
		if err = ie.Pdsch_AggregationFactor_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Pdsch_AggregationFactor_r17", err)
		}
	}
	return nil
}
