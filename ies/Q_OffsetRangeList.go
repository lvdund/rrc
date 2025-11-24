package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type Q_OffsetRangeList struct {
	RsrpOffsetSSB    Q_OffsetRange `madatory`
	RsrqOffsetSSB    Q_OffsetRange `madatory`
	SinrOffsetSSB    Q_OffsetRange `madatory`
	RsrpOffsetCSI_RS Q_OffsetRange `madatory`
	RsrqOffsetCSI_RS Q_OffsetRange `madatory`
	SinrOffsetCSI_RS Q_OffsetRange `madatory`
}

func (ie *Q_OffsetRangeList) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.RsrpOffsetSSB.Encode(w); err != nil {
		return utils.WrapError("Encode RsrpOffsetSSB", err)
	}
	if err = ie.RsrqOffsetSSB.Encode(w); err != nil {
		return utils.WrapError("Encode RsrqOffsetSSB", err)
	}
	if err = ie.SinrOffsetSSB.Encode(w); err != nil {
		return utils.WrapError("Encode SinrOffsetSSB", err)
	}
	if err = ie.RsrpOffsetCSI_RS.Encode(w); err != nil {
		return utils.WrapError("Encode RsrpOffsetCSI_RS", err)
	}
	if err = ie.RsrqOffsetCSI_RS.Encode(w); err != nil {
		return utils.WrapError("Encode RsrqOffsetCSI_RS", err)
	}
	if err = ie.SinrOffsetCSI_RS.Encode(w); err != nil {
		return utils.WrapError("Encode SinrOffsetCSI_RS", err)
	}
	return nil
}

func (ie *Q_OffsetRangeList) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.RsrpOffsetSSB.Decode(r); err != nil {
		return utils.WrapError("Decode RsrpOffsetSSB", err)
	}
	if err = ie.RsrqOffsetSSB.Decode(r); err != nil {
		return utils.WrapError("Decode RsrqOffsetSSB", err)
	}
	if err = ie.SinrOffsetSSB.Decode(r); err != nil {
		return utils.WrapError("Decode SinrOffsetSSB", err)
	}
	if err = ie.RsrpOffsetCSI_RS.Decode(r); err != nil {
		return utils.WrapError("Decode RsrpOffsetCSI_RS", err)
	}
	if err = ie.RsrqOffsetCSI_RS.Decode(r); err != nil {
		return utils.WrapError("Decode RsrqOffsetCSI_RS", err)
	}
	if err = ie.SinrOffsetCSI_RS.Decode(r); err != nil {
		return utils.WrapError("Decode SinrOffsetCSI_RS", err)
	}
	return nil
}
