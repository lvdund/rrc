package ies

// NotificationMessageSidelink-r17-criticalExtensions ::= CHOICE
const (
	NotificationmessagesidelinkR17CriticalextensionsChoiceNothing = iota
	NotificationmessagesidelinkR17CriticalextensionsChoiceNotificationmessagesidelinkR17
	NotificationmessagesidelinkR17CriticalextensionsChoiceCriticalextensionsfuture
)

type NotificationmessagesidelinkR17Criticalextensions struct {
	Choice                         uint64
	NotificationmessagesidelinkR17 *NotificationmessagesidelinkR17
	Criticalextensionsfuture       *NotificationmessagesidelinkR17CriticalextensionsCriticalextensionsfuture
}
