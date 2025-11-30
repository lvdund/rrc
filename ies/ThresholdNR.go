package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type ThresholdNR struct {
	ThresholdRSRP *RSRP_Range `optional`
	ThresholdRSRQ *RSRQ_Range `optional`
	ThresholdSINR *SINR_Range `optional`
}

func (ie *ThresholdNR) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.ThresholdRSRP != nil, ie.ThresholdRSRQ != nil, ie.ThresholdSINR != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ThresholdRSRP != nil {
		if err = ie.ThresholdRSRP.Encode(w); err != nil {
			return utils.WrapError("Encode ThresholdRSRP", err)
		}
	}
	if ie.ThresholdRSRQ != nil {
		if err = ie.ThresholdRSRQ.Encode(w); err != nil {
			return utils.WrapError("Encode ThresholdRSRQ", err)
		}
	}
	if ie.ThresholdSINR != nil {
		if err = ie.ThresholdSINR.Encode(w); err != nil {
			return utils.WrapError("Encode ThresholdSINR", err)
		}
	}
	return nil
}

func (ie *ThresholdNR) Decode(r *aper.AperReader) error {
	var err error
	var ThresholdRSRPPresent bool
	if ThresholdRSRPPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ThresholdRSRQPresent bool
	if ThresholdRSRQPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ThresholdSINRPresent bool
	if ThresholdSINRPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if ThresholdRSRPPresent {
		ie.ThresholdRSRP = new(RSRP_Range)
		if err = ie.ThresholdRSRP.Decode(r); err != nil {
			return utils.WrapError("Decode ThresholdRSRP", err)
		}
	}
	if ThresholdRSRQPresent {
		ie.ThresholdRSRQ = new(RSRQ_Range)
		if err = ie.ThresholdRSRQ.Decode(r); err != nil {
			return utils.WrapError("Decode ThresholdRSRQ", err)
		}
	}
	if ThresholdSINRPresent {
		ie.ThresholdSINR = new(SINR_Range)
		if err = ie.ThresholdSINR.Decode(r); err != nil {
			return utils.WrapError("Decode ThresholdSINR", err)
		}
	}
	return nil
}
