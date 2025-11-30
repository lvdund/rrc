package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CMRGroupingAndPairing_r17 struct {
	NrofResourcesGroup1_r17 int64                   `lb:1,ub:7,madatory`
	Pair1OfNZP_CSI_RS_r17   *NZP_CSI_RS_Pairing_r17 `optional`
	Pair2OfNZP_CSI_RS_r17   *NZP_CSI_RS_Pairing_r17 `optional`
}

func (ie *CMRGroupingAndPairing_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Pair1OfNZP_CSI_RS_r17 != nil, ie.Pair2OfNZP_CSI_RS_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteInteger(ie.NrofResourcesGroup1_r17, &aper.Constraint{Lb: 1, Ub: 7}, false); err != nil {
		return utils.WrapError("WriteInteger NrofResourcesGroup1_r17", err)
	}
	if ie.Pair1OfNZP_CSI_RS_r17 != nil {
		if err = ie.Pair1OfNZP_CSI_RS_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Pair1OfNZP_CSI_RS_r17", err)
		}
	}
	if ie.Pair2OfNZP_CSI_RS_r17 != nil {
		if err = ie.Pair2OfNZP_CSI_RS_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Pair2OfNZP_CSI_RS_r17", err)
		}
	}
	return nil
}

func (ie *CMRGroupingAndPairing_r17) Decode(r *aper.AperReader) error {
	var err error
	var Pair1OfNZP_CSI_RS_r17Present bool
	if Pair1OfNZP_CSI_RS_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Pair2OfNZP_CSI_RS_r17Present bool
	if Pair2OfNZP_CSI_RS_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_int_NrofResourcesGroup1_r17 int64
	if tmp_int_NrofResourcesGroup1_r17, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 7}, false); err != nil {
		return utils.WrapError("ReadInteger NrofResourcesGroup1_r17", err)
	}
	ie.NrofResourcesGroup1_r17 = tmp_int_NrofResourcesGroup1_r17
	if Pair1OfNZP_CSI_RS_r17Present {
		ie.Pair1OfNZP_CSI_RS_r17 = new(NZP_CSI_RS_Pairing_r17)
		if err = ie.Pair1OfNZP_CSI_RS_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Pair1OfNZP_CSI_RS_r17", err)
		}
	}
	if Pair2OfNZP_CSI_RS_r17Present {
		ie.Pair2OfNZP_CSI_RS_r17 = new(NZP_CSI_RS_Pairing_r17)
		if err = ie.Pair2OfNZP_CSI_RS_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Pair2OfNZP_CSI_RS_r17", err)
		}
	}
	return nil
}
