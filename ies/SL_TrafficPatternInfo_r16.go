package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SL_TrafficPatternInfo_r16 struct {
	TrafficPeriodicity_r16  SL_TrafficPatternInfo_r16_trafficPeriodicity_r16 `madatory`
	TimingOffset_r16        int64                                            `lb:0,ub:10239,madatory`
	MessageSize_r16         aper.BitString                                   `lb:8,ub:8,madatory`
	Sl_QoS_FlowIdentity_r16 SL_QoS_FlowIdentity_r16                          `madatory`
}

func (ie *SL_TrafficPatternInfo_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.TrafficPeriodicity_r16.Encode(w); err != nil {
		return utils.WrapError("Encode TrafficPeriodicity_r16", err)
	}
	if err = w.WriteInteger(ie.TimingOffset_r16, &aper.Constraint{Lb: 0, Ub: 10239}, false); err != nil {
		return utils.WrapError("WriteInteger TimingOffset_r16", err)
	}
	if err = w.WriteBitString(ie.MessageSize_r16.Bytes, uint(ie.MessageSize_r16.NumBits), &aper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
		return utils.WrapError("WriteBitString MessageSize_r16", err)
	}
	if err = ie.Sl_QoS_FlowIdentity_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Sl_QoS_FlowIdentity_r16", err)
	}
	return nil
}

func (ie *SL_TrafficPatternInfo_r16) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.TrafficPeriodicity_r16.Decode(r); err != nil {
		return utils.WrapError("Decode TrafficPeriodicity_r16", err)
	}
	var tmp_int_TimingOffset_r16 int64
	if tmp_int_TimingOffset_r16, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 10239}, false); err != nil {
		return utils.WrapError("ReadInteger TimingOffset_r16", err)
	}
	ie.TimingOffset_r16 = tmp_int_TimingOffset_r16
	var tmp_bs_MessageSize_r16 []byte
	var n_MessageSize_r16 uint
	if tmp_bs_MessageSize_r16, n_MessageSize_r16, err = r.ReadBitString(&aper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
		return utils.WrapError("ReadBitString MessageSize_r16", err)
	}
	ie.MessageSize_r16 = aper.BitString{
		Bytes:   tmp_bs_MessageSize_r16,
		NumBits: uint64(n_MessageSize_r16),
	}
	if err = ie.Sl_QoS_FlowIdentity_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Sl_QoS_FlowIdentity_r16", err)
	}
	return nil
}
