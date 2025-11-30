package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	T_PollRetransmit_Enum_ms5       aper.Enumerated = 0
	T_PollRetransmit_Enum_ms10      aper.Enumerated = 1
	T_PollRetransmit_Enum_ms15      aper.Enumerated = 2
	T_PollRetransmit_Enum_ms20      aper.Enumerated = 3
	T_PollRetransmit_Enum_ms25      aper.Enumerated = 4
	T_PollRetransmit_Enum_ms30      aper.Enumerated = 5
	T_PollRetransmit_Enum_ms35      aper.Enumerated = 6
	T_PollRetransmit_Enum_ms40      aper.Enumerated = 7
	T_PollRetransmit_Enum_ms45      aper.Enumerated = 8
	T_PollRetransmit_Enum_ms50      aper.Enumerated = 9
	T_PollRetransmit_Enum_ms55      aper.Enumerated = 10
	T_PollRetransmit_Enum_ms60      aper.Enumerated = 11
	T_PollRetransmit_Enum_ms65      aper.Enumerated = 12
	T_PollRetransmit_Enum_ms70      aper.Enumerated = 13
	T_PollRetransmit_Enum_ms75      aper.Enumerated = 14
	T_PollRetransmit_Enum_ms80      aper.Enumerated = 15
	T_PollRetransmit_Enum_ms85      aper.Enumerated = 16
	T_PollRetransmit_Enum_ms90      aper.Enumerated = 17
	T_PollRetransmit_Enum_ms95      aper.Enumerated = 18
	T_PollRetransmit_Enum_ms100     aper.Enumerated = 19
	T_PollRetransmit_Enum_ms105     aper.Enumerated = 20
	T_PollRetransmit_Enum_ms110     aper.Enumerated = 21
	T_PollRetransmit_Enum_ms115     aper.Enumerated = 22
	T_PollRetransmit_Enum_ms120     aper.Enumerated = 23
	T_PollRetransmit_Enum_ms125     aper.Enumerated = 24
	T_PollRetransmit_Enum_ms130     aper.Enumerated = 25
	T_PollRetransmit_Enum_ms135     aper.Enumerated = 26
	T_PollRetransmit_Enum_ms140     aper.Enumerated = 27
	T_PollRetransmit_Enum_ms145     aper.Enumerated = 28
	T_PollRetransmit_Enum_ms150     aper.Enumerated = 29
	T_PollRetransmit_Enum_ms155     aper.Enumerated = 30
	T_PollRetransmit_Enum_ms160     aper.Enumerated = 31
	T_PollRetransmit_Enum_ms165     aper.Enumerated = 32
	T_PollRetransmit_Enum_ms170     aper.Enumerated = 33
	T_PollRetransmit_Enum_ms175     aper.Enumerated = 34
	T_PollRetransmit_Enum_ms180     aper.Enumerated = 35
	T_PollRetransmit_Enum_ms185     aper.Enumerated = 36
	T_PollRetransmit_Enum_ms190     aper.Enumerated = 37
	T_PollRetransmit_Enum_ms195     aper.Enumerated = 38
	T_PollRetransmit_Enum_ms200     aper.Enumerated = 39
	T_PollRetransmit_Enum_ms205     aper.Enumerated = 40
	T_PollRetransmit_Enum_ms210     aper.Enumerated = 41
	T_PollRetransmit_Enum_ms215     aper.Enumerated = 42
	T_PollRetransmit_Enum_ms220     aper.Enumerated = 43
	T_PollRetransmit_Enum_ms225     aper.Enumerated = 44
	T_PollRetransmit_Enum_ms230     aper.Enumerated = 45
	T_PollRetransmit_Enum_ms235     aper.Enumerated = 46
	T_PollRetransmit_Enum_ms240     aper.Enumerated = 47
	T_PollRetransmit_Enum_ms245     aper.Enumerated = 48
	T_PollRetransmit_Enum_ms250     aper.Enumerated = 49
	T_PollRetransmit_Enum_ms300     aper.Enumerated = 50
	T_PollRetransmit_Enum_ms350     aper.Enumerated = 51
	T_PollRetransmit_Enum_ms400     aper.Enumerated = 52
	T_PollRetransmit_Enum_ms450     aper.Enumerated = 53
	T_PollRetransmit_Enum_ms500     aper.Enumerated = 54
	T_PollRetransmit_Enum_ms800     aper.Enumerated = 55
	T_PollRetransmit_Enum_ms1000    aper.Enumerated = 56
	T_PollRetransmit_Enum_ms2000    aper.Enumerated = 57
	T_PollRetransmit_Enum_ms4000    aper.Enumerated = 58
	T_PollRetransmit_Enum_ms1_v1610 aper.Enumerated = 59
	T_PollRetransmit_Enum_ms2_v1610 aper.Enumerated = 60
	T_PollRetransmit_Enum_ms3_v1610 aper.Enumerated = 61
	T_PollRetransmit_Enum_ms4_v1610 aper.Enumerated = 62
	T_PollRetransmit_Enum_spare1    aper.Enumerated = 63
)

type T_PollRetransmit struct {
	Value aper.Enumerated `lb:0,ub:63,madatory`
}

func (ie *T_PollRetransmit) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
		return utils.WrapError("Encode T_PollRetransmit", err)
	}
	return nil
}

func (ie *T_PollRetransmit) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
		return utils.WrapError("Decode T_PollRetransmit", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
