package isteam_user

const (
	//社区可见状态
	CommunityVisibilityState = iota
	CommunityVisibilityStatePrivateOrFriendOnly
	CommunityVisibilityStateOther
	CommunityVisibilityStatePublic
)

const (
	PersonaStateOffline = iota
	PersonaStateOnline
	PersonaStateBusy
	PersonaStateAway
	PersonaStateSnooze
	PersonaStateLookingToTrade
	PersonaStateLookingToPlay
)
