package ies

// SliceInfoDedicated-r17 ::= SEQUENCE
type SliceinfodedicatedR17 struct {
	NsagIdentityinfoR17               NsagIdentityinfoR17
	NsagCellreselectionpriorityR17    *Cellreselectionpriority
	NsagCellreselectionsubpriorityR17 *Cellreselectionsubpriority
}
