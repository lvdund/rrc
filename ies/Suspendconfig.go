package ies

import "rrc/utils"

// SuspendConfig ::= SEQUENCE
// Extensible
type Suspendconfig struct {
	FulliRnti                 IRntiValue
	ShortiRnti                ShortiRntiValue
	RanPagingcycle            Pagingcycle
	RanNotificationareainfo   *RanNotificationareainfo
	T380                      *PeriodicrnauTimervalue
	Nexthopchainingcount      Nexthopchainingcount
	SlUeidentityremoteR17     *RntiValue                                `ext`
	SdtConfigR17              *utils.Setuprelease[SdtConfigR17]         `ext`
	SrsPosrrcInactiveR17      *utils.Setuprelease[SrsPosrrcInactiveR17] `ext`
	RanExtendedpagingcycleR17 *ExtendedpagingcycleR17                   `ext`
}
