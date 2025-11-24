package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FreqBandInformationNR struct {
	BandNR                  FreqBandIndicatorNR  `madatory`
	MaxBandwidthRequestedDL *AggregatedBandwidth `optional`
	MaxBandwidthRequestedUL *AggregatedBandwidth `optional`
	MaxCarriersRequestedDL  *int64               `lb:1,ub:maxNrofServingCells,optional`
	MaxCarriersRequestedUL  *int64               `lb:1,ub:maxNrofServingCells,optional`
}

func (ie *FreqBandInformationNR) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.MaxBandwidthRequestedDL != nil, ie.MaxBandwidthRequestedUL != nil, ie.MaxCarriersRequestedDL != nil, ie.MaxCarriersRequestedUL != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.BandNR.Encode(w); err != nil {
		return utils.WrapError("Encode BandNR", err)
	}
	if ie.MaxBandwidthRequestedDL != nil {
		if err = ie.MaxBandwidthRequestedDL.Encode(w); err != nil {
			return utils.WrapError("Encode MaxBandwidthRequestedDL", err)
		}
	}
	if ie.MaxBandwidthRequestedUL != nil {
		if err = ie.MaxBandwidthRequestedUL.Encode(w); err != nil {
			return utils.WrapError("Encode MaxBandwidthRequestedUL", err)
		}
	}
	if ie.MaxCarriersRequestedDL != nil {
		if err = w.WriteInteger(*ie.MaxCarriersRequestedDL, &uper.Constraint{Lb: 1, Ub: maxNrofServingCells}, false); err != nil {
			return utils.WrapError("Encode MaxCarriersRequestedDL", err)
		}
	}
	if ie.MaxCarriersRequestedUL != nil {
		if err = w.WriteInteger(*ie.MaxCarriersRequestedUL, &uper.Constraint{Lb: 1, Ub: maxNrofServingCells}, false); err != nil {
			return utils.WrapError("Encode MaxCarriersRequestedUL", err)
		}
	}
	return nil
}

func (ie *FreqBandInformationNR) Decode(r *uper.UperReader) error {
	var err error
	var MaxBandwidthRequestedDLPresent bool
	if MaxBandwidthRequestedDLPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var MaxBandwidthRequestedULPresent bool
	if MaxBandwidthRequestedULPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var MaxCarriersRequestedDLPresent bool
	if MaxCarriersRequestedDLPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var MaxCarriersRequestedULPresent bool
	if MaxCarriersRequestedULPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.BandNR.Decode(r); err != nil {
		return utils.WrapError("Decode BandNR", err)
	}
	if MaxBandwidthRequestedDLPresent {
		ie.MaxBandwidthRequestedDL = new(AggregatedBandwidth)
		if err = ie.MaxBandwidthRequestedDL.Decode(r); err != nil {
			return utils.WrapError("Decode MaxBandwidthRequestedDL", err)
		}
	}
	if MaxBandwidthRequestedULPresent {
		ie.MaxBandwidthRequestedUL = new(AggregatedBandwidth)
		if err = ie.MaxBandwidthRequestedUL.Decode(r); err != nil {
			return utils.WrapError("Decode MaxBandwidthRequestedUL", err)
		}
	}
	if MaxCarriersRequestedDLPresent {
		var tmp_int_MaxCarriersRequestedDL int64
		if tmp_int_MaxCarriersRequestedDL, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: maxNrofServingCells}, false); err != nil {
			return utils.WrapError("Decode MaxCarriersRequestedDL", err)
		}
		ie.MaxCarriersRequestedDL = &tmp_int_MaxCarriersRequestedDL
	}
	if MaxCarriersRequestedULPresent {
		var tmp_int_MaxCarriersRequestedUL int64
		if tmp_int_MaxCarriersRequestedUL, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: maxNrofServingCells}, false); err != nil {
			return utils.WrapError("Decode MaxCarriersRequestedUL", err)
		}
		ie.MaxCarriersRequestedUL = &tmp_int_MaxCarriersRequestedUL
	}
	return nil
}
