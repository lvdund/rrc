// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RRCRelease-IEs (line 1229).

var rRCReleaseIEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "redirectedCarrierInfo", Optional: true},
		{Name: "cellReselectionPriorities", Optional: true},
		{Name: "suspendConfig", Optional: true},
		{Name: "deprioritisationReq", Optional: true},
		{Name: "lateNonCriticalExtension", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

var rRCReleaseIEsLateNonCriticalExtensionConstraints = per.SizeConstraints{}

const (
	RRCRelease_IEs_DeprioritisationReq_DeprioritisationType_Frequency = 0
	RRCRelease_IEs_DeprioritisationReq_DeprioritisationType_Nr        = 1
)

var rRCReleaseIEsDeprioritisationReqDeprioritisationTypeConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RRCRelease_IEs_DeprioritisationReq_DeprioritisationType_Frequency, RRCRelease_IEs_DeprioritisationReq_DeprioritisationType_Nr},
}

const (
	RRCRelease_IEs_DeprioritisationReq_DeprioritisationTimer_Min5  = 0
	RRCRelease_IEs_DeprioritisationReq_DeprioritisationTimer_Min10 = 1
	RRCRelease_IEs_DeprioritisationReq_DeprioritisationTimer_Min15 = 2
	RRCRelease_IEs_DeprioritisationReq_DeprioritisationTimer_Min30 = 3
)

var rRCReleaseIEsDeprioritisationReqDeprioritisationTimerConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RRCRelease_IEs_DeprioritisationReq_DeprioritisationTimer_Min5, RRCRelease_IEs_DeprioritisationReq_DeprioritisationTimer_Min10, RRCRelease_IEs_DeprioritisationReq_DeprioritisationTimer_Min15, RRCRelease_IEs_DeprioritisationReq_DeprioritisationTimer_Min30},
}

type RRCRelease_IEs struct {
	RedirectedCarrierInfo     *RedirectedCarrierInfo
	CellReselectionPriorities *CellReselectionPriorities
	SuspendConfig             *SuspendConfig
	DeprioritisationReq       *struct {
		DeprioritisationType  int64
		DeprioritisationTimer int64
	}
	LateNonCriticalExtension []byte
	NonCriticalExtension     *RRCRelease_v1540_IEs
}

func (ie *RRCRelease_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rRCReleaseIEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.RedirectedCarrierInfo != nil, ie.CellReselectionPriorities != nil, ie.SuspendConfig != nil, ie.DeprioritisationReq != nil, ie.LateNonCriticalExtension != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. redirectedCarrierInfo: ref
	{
		if ie.RedirectedCarrierInfo != nil {
			if err := ie.RedirectedCarrierInfo.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. cellReselectionPriorities: ref
	{
		if ie.CellReselectionPriorities != nil {
			if err := ie.CellReselectionPriorities.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. suspendConfig: ref
	{
		if ie.SuspendConfig != nil {
			if err := ie.SuspendConfig.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. deprioritisationReq: sequence
	{
		if ie.DeprioritisationReq != nil {
			c := ie.DeprioritisationReq
			if err := e.EncodeEnumerated(c.DeprioritisationType, rRCReleaseIEsDeprioritisationReqDeprioritisationTypeConstraints); err != nil {
				return err
			}
			if err := e.EncodeEnumerated(c.DeprioritisationTimer, rRCReleaseIEsDeprioritisationReqDeprioritisationTimerConstraints); err != nil {
				return err
			}
		}
	}

	// 6. lateNonCriticalExtension: octet-string
	{
		if ie.LateNonCriticalExtension != nil {
			if err := e.EncodeOctetString(ie.LateNonCriticalExtension, rRCReleaseIEsLateNonCriticalExtensionConstraints); err != nil {
				return err
			}
		}
	}

	// 7. nonCriticalExtension: ref
	{
		if ie.NonCriticalExtension != nil {
			if err := ie.NonCriticalExtension.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *RRCRelease_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rRCReleaseIEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. redirectedCarrierInfo: ref
	{
		if seq.IsComponentPresent(0) {
			ie.RedirectedCarrierInfo = new(RedirectedCarrierInfo)
			if err := ie.RedirectedCarrierInfo.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. cellReselectionPriorities: ref
	{
		if seq.IsComponentPresent(1) {
			ie.CellReselectionPriorities = new(CellReselectionPriorities)
			if err := ie.CellReselectionPriorities.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. suspendConfig: ref
	{
		if seq.IsComponentPresent(2) {
			ie.SuspendConfig = new(SuspendConfig)
			if err := ie.SuspendConfig.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. deprioritisationReq: sequence
	{
		if seq.IsComponentPresent(3) {
			ie.DeprioritisationReq = &struct {
				DeprioritisationType  int64
				DeprioritisationTimer int64
			}{}
			c := ie.DeprioritisationReq
			{
				v, err := d.DecodeEnumerated(rRCReleaseIEsDeprioritisationReqDeprioritisationTypeConstraints)
				if err != nil {
					return err
				}
				c.DeprioritisationType = v
			}
			{
				v, err := d.DecodeEnumerated(rRCReleaseIEsDeprioritisationReqDeprioritisationTimerConstraints)
				if err != nil {
					return err
				}
				c.DeprioritisationTimer = v
			}
		}
	}

	// 6. lateNonCriticalExtension: octet-string
	{
		if seq.IsComponentPresent(4) {
			v, err := d.DecodeOctetString(rRCReleaseIEsLateNonCriticalExtensionConstraints)
			if err != nil {
				return err
			}
			ie.LateNonCriticalExtension = v
		}
	}

	// 7. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(5) {
			ie.NonCriticalExtension = new(RRCRelease_v1540_IEs)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
