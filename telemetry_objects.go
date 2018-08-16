package pubgo

// Location data for telemetry events
type Location struct {
	X float32 `json:"X"`
	Y float32 `json:"Y"`
	Z float32 `json:"Z"`
}

// GameState data for telemetry events
type GameState struct {
	ElapsedTime              int      `json:"ElapsedTime"`
	NumAliveTeams            int      `json:"NumAliveTeams"`
	NumJoinPlayers           int      `json:"NumJoinPlayers"`
	NumStartPlayers          int      `json:"NumStartPlayers"`
	NumAlivePlayers          int      `json:"NumAlivePlayers"`
	SafetyZonePosition       Location `json:"SafetyZonePosition"`
	SafetyZoneRadius         float32  `json:"SafetyZoneRadius"`
	PoisonGasWarningPosition Location `json:"PoisonGasWarningPosition"`
	PoisonGasWarningRadius   float32  `json:"PoisonGasWarningRadius"`
	RedZonePosition          Location `json:"RedZonePosition"`
	RedZoneRadius            float32  `json:"RedZoneRadius"`
}

// Vehicle data for telemetry events
type Vehicle struct {
	VehicleType   string  `json:"VehicleType"`
	VehicleID     string  `json:"VehicleId"`
	HealthPercent float32 `json:"HealthPercent"`
	FuelPercent   float32 `json:"FeulPercent"`
}

// Character data for telemetry events
type Character struct {
	Name      string   `json:"Name"`
	TeamID    int      `json:"TeamId"`
	Health    float32  `json:"Health"`
	Location  Location `json:"Location"`
	Ranking   int      `json:"Ranking"`
	AccountID string   `json:"AccountId"`
}

// Item data for telemetry events
type Item struct {
	ItemID        string   `json:"ItemId"`
	StackCount    int      `json:"StackCount"`
	Category      string   `json:"Category"`
	SubCategory   string   `json:"SubCategory"`
	AttachedItems []string `json:"AttachedItems"`
}

// ItemPackage data for telemetry events
type ItemPackage struct {
	ItemPackageID string   `json:"ItemPackageId"`
	Location      Location `json:"Location"`
	Items         []Item   `json:"Items"`
}

// Common data for telemetry events
type Common struct {
	MatchID string  `json:"matchId"`
	MapName string  `json:"mapName"`
	IsGame  float32 `json:"isGame"`
}
