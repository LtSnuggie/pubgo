package pubgo

import "time"

const (
	playerLogin       = "LogPlayerLogin"
	playerCreate      = "LogPlayerCreate"
	playerPosition    = "LogPlayerPosition"
	playerAttack      = "LogPlayerAttack"
	itemPickup        = "LogItemPickup"
	itemEquip         = "LogItemEquip"
	itemUnequip       = "LogItemUnequip"
	vehicleRide       = "LogVehicleRide"
	matchDefinition   = "LogMatchDefinition"
	matchStart        = "LogMatchStart"
	gameStatePeriodic = "LogGameStatePeriodic"
	vehicleLeave      = "LogVehicleLeave"
	playerTakeDamage  = "LogPlayerTakeDamage"
	playerLogout      = "LogPlayerLogout"
	itemAttach        = "LogItemAttach"
	itemDrop          = "LogItemDrop"
	playerKill        = "LogPlayerKill"
	itemDetach        = "LogItemDetach"
	itemUse           = "LogItemUse"
	carePackageSpawn  = "LogCarePackageSpawn"
	vehicleDestroy    = "LogVehicleDestroy"
	carePackageLand   = "LogCarePackageLand"
	matchEnd          = "LogMatchEnd"
)

type TelemetryEvent interface {
	GetType() string
	GetTimestamp() time.Time
	GetVersion() int
	GetCommon() Common
}

type Base struct {
	Version   int       `json:"_V"`
	Timestamp time.Time `json:"_D"`
	Type      string    `json:"_T"`
	Common    Common    `json:"Common"`
}

func (b Base) GetType() string {
	return b.Type
}

func (b Base) GetTimestamp() time.Time {
	return b.Timestamp
}

func (b Base) GetVersion() int {
	return b.Version
}

func (b Base) GetCommon() Common {
	return b.Common
}

type PlayerLoginEvent struct {
	Base
	Result       bool   `json:"Result"`
	ErrorMessage string `json:"ErrorMessage"`
	AccountID    string `json:"AccountId"`
}

type PlayerCreateEvent struct {
	Base
	Character Character `json:"Character"`
}

type PlayerPositionEvent struct {
	Base
	Character       Character `json:"Character"`
	ElapsedTime     float32   `json:"ElapsedTime"`
	NumAlivePlayers int       `json:"NumAlivePlayers"`
}

type PlayerAttackEvent struct {
	Base
	AttackID   int       `json:"AttackId"`
	Attacker   Character `json:"Attacker"`
	AttackType string    `json:"AttackType"`
	Weapon     Item      `json:"Weapon"`
	Vehicle    Vehicle   `json:"Vehicle"`
}

type ItemPickupEvent struct {
	Base
	Character Character `json:"Character"`
	Item      Item      `json:"Item"`
}

type ItemEquipEvent struct {
	Base
	Character Character `json:"Character"`
	Item      Item      `json:"Item"`
}

type ItemUnequipEvent struct {
	Base
	Character Character `json:"Character"`
	Item      Item      `json:"Item"`
}

type VehicleRideEvent struct {
	Base
	Character Character `json:"Character"`
	Vehicle   Vehicle   `json:"Vehicle"`
}

type MatchDefinitionEvent struct {
	Base
	MatchId     string `json:"MatchId"`
	PingQuality string `json:"PingQuality"`
}

type MatchStartEvent struct {
	Base
	Characters []Character `json:"Characters"`
}

type GameStatePeriodicEvent struct {
	Base
	GameState GameState `json:"GameState"`
}

type VehicleLeaveEvent struct {
	Base
	Character Character `json:"Character"`
	Vehicle   Vehicle   `json:"Vehicle"`
}

type PlayerTakeDamageEvent struct {
	Base
	AttackID           int       `json:"AttackId"`
	Attacker           Character `json:"Attacker"`
	Victim             Character `json:"Victim"`
	DamageTypeCategory string    `json:"DamageTypeCategory"`
	DamageReason       string    `json:"DamageReason"`
	Damage             float32   `json:"Damage"`
	DamageCauserName   string    `json:"DamageCauserName"`
}

type PlayerLogoutEvent struct {
	Base
	AccountID string `json:"AccountId"`
}

type ItemAttachEvent struct {
	Base
	Character  Character `json:"Character"`
	ParentItem Item      `json:"ParentItem"`
	ChildItem  Item      `json:"ChildItem"`
}

type ItemDropEvent struct {
	Base
	Character Character `json:"Character"`
	Item      Item      `json:"Item"`
}

type PlayerKillEvent struct {
	Base
	AttackID           int       `json:"AttackId"`
	Killer             Character `json:"Killer"`
	Victim             Character `json:"Victim"`
	DamageTypeCategory string    `json:"DamageTypeCategory"`
	DamageCauserName   string    `json:"DamageCauserName"`
	Distance           float32   `json:"Distance"`
}

type ItemDetachEvent struct {
	Base
	Character  Character `json:"Character"`
	ParentItem Item      `json:"ParentItem"`
	ChildItem  Item      `json:"ChildItem"`
}

type ItemUseEvent struct {
	Base
	Character Character `json:"Character"`
	Item      Item      `json:"Item"`
}

type CarePackageSpawnEvent struct {
	Base
	ItemPackage ItemPackage `json:"ItemPackage"`
}

type VehicleDestroyEvent struct {
	Base
	AttackID           int       `json:"AttackId"`
	Attacker           Character `json:"Attacker"`
	Vehicle            Vehicle   `json:"Vehicle"`
	DamageTypeCategory string    `json:"DamageTypeCategory"`
	DamageCauserName   string    `json:"DamageCauserName"`
	Distance           float32   `json:"Distance"`
}

type CarePackageLandEvent struct {
	Base
	ItemPackage ItemPackage `json:"ItemPackage"`
}

type MatchEndEvent struct {
	Base
	Characters []Character `json:"Characters"`
}
