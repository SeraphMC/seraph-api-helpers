package utils

const (
	// PostData Ability to post data to the API.
	PostData PermissionNode = "post/data"

	// PingAll All ping fom a specified player.
	PingAll PermissionNode = "ping/all"
	// PingToday Today's ping for a specified player.
	PingToday PermissionNode = "ping/today"
	// PingRecent Recent ping for a specified player.
	PingRecent PermissionNode = "ping/recent"

	// ChatAll All chat messages from a specified player.
	ChatAll PermissionNode = "chat/all"
	// ChatLobbyOne Live lobby chat messages
	ChatLobbyOne PermissionNode = "chat/lobby-one"

	// AdminGenerate Generate a new API key.
	AdminGenerate PermissionNode = "admin/generate"
	// AdminDeveloper Developer Permissions.
	AdminDeveloper PermissionNode = "admin/developer"
	// AdminDatabaseStats Allowed to fetch Database Stats.
	AdminDatabaseStats PermissionNode = "admin/database-stats"

	// ClientEssential User is allowed to fetch Essential Client
	ClientEssential PermissionNode = "client/essential"
	// ClientLaby User is allowed to fetch Labymod Client
	ClientLaby PermissionNode = "client/laby"
	// ClientLunar User is allowed to fetch Lunar Client
	ClientLunar PermissionNode = "client/lunar"
	// ClientBlc User is allowed to fetch Badlion Client
	ClientBlc PermissionNode = "client/blc"
	// ClientFeather User is allowed to fetch Feather Client
	ClientFeather PermissionNode = "client/feather"
	// ClientCosmetics User is allowed to fetch Cosmetics Client
	ClientCosmetics PermissionNode = "client/cosmetics"

	// HypixelProxy User is allowed to fetch Hypixel Proxy
	HypixelProxy PermissionNode = "hypixel/proxy"
	// HypixelLowCache User is allowed to fetch Hypixel Proxy with a lower cache time
	HypixelLowCache PermissionNode = "hypixel/low-cache"
	// HypixelLeaderboards User is allowed to fetch Hypixel Leaderboards
	HypixelLeaderboards PermissionNode = "hypixel/leaderboards"
	// PlayerLookup User is allowed to fetch player lookup data
	PlayerLookup PermissionNode = "player/lookup"

	// AccountLinking User is allowed to link players to an account
	AccountLinking PermissionNode = "account/linking"
	// AccountUser User is allowed to fetch account data
	AccountUser PermissionNode = "account/user"

	// ErrorLogging User is allowed to modify error logging data
	ErrorLogging PermissionNode = "error/logging"
)
