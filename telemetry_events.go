package pubgo

import "time"

const (
	playerLogin       = "logplayerlogin"
	playerCreate      = "logplayercreate"
	playerPosition    = "logplayerposition"
	playerAttack      = "logplayerattack"
	itemPickup        = "logitempickup"
	itemEquip         = "logitemequip"
	itemUnequip       = "logitemunequip"
	vehicleRide       = "logvehicleride"
	matchDefinition   = "logmatchdefinition"
	matchStart        = "logmatchstart"
	gameStatePeriodic = "loggamestateperiodic"
	vehicleLeave      = "logvehicleleave"
	playerTakeDamage  = "logplayertakedamage"
	playerLogout      = "logplayerlogout"
	itemAttach        = "logitemattach"
	itemDrop          = "logitemdrop"
	playerKill        = "logplayerkill"
	itemDetach        = "logitemdetach"
	itemUse           = "logitemuse"
	carePackageSpawn  = "logcarepackagespawn"
	vehicleDestroy    = "logvehicledestroy"
	carePackageLand   = "logcarepackageland"
	matchEnd          = "logmatchend"
)

// TelemetryEvent is an interface for TelemetryEvent's.
type TelemetryEvent interface {
	GetType() string         // returns the event type.
	GetTimestamp() time.Time // returns the Timestamp of the event.
	GetVersion() int         // returns the version of the event.
	GetCommon() Common       // returns the Common object of the event.
}

// Base is the base of all telemetery event types.
// This information will always be returned.
type Base struct {
	// Version   int       `json:"_V"`
	Timestamp time.Time `json:"_D"`
	Type      string    `json:"_T"`
	Common    Common    `json:"Common"`
}

// GetType returns the event type.
func (b Base) GetType() string {
	return b.Type
}

// GetTimestamp returns the event timestamp.
func (b Base) GetTimestamp() time.Time {
	return b.Timestamp
}

// GetVersion returns the version of the event.
func (b Base) GetVersion() int {
	return 0
}

// GetCommon returns the common object of the event.
func (b Base) GetCommon() Common {
	return b.Common
}

// PlayerLoginEvent contains data from the player loging in
type PlayerLoginEvent struct {
	Base
	// Result bool `json:"Result"`
	// ErrorMessage string `json:"ErrorMessage"`
	AccountID string `json:"AccountId"`
}

// PlayerCreateEvent contains data from the player being created on the server
type PlayerCreateEvent struct {
	Base
	Character Character `json:"Character"`
}

// PlayerPositionEvent contains data about the current position of the player
type PlayerPositionEvent struct {
	Base
	Character       Character `json:"Character"`
	ElapsedTime     float32   `json:"ElapsedTime"`
	NumAlivePlayers int       `json:"NumAlivePlayers"`
}

// PlayerAttackEvent contains data about the player attacking another player
type PlayerAttackEvent struct {
	Base
	AttackID   int       `json:"AttackId"`
	Attacker   Character `json:"Attacker"`
	AttackType string    `json:"AttackType"`
	Weapon     Item      `json:"Weapon"`
	Vehicle    Vehicle   `json:"Vehicle"`
}

// ItemPickupEvent contains data from the player picking up an item
type ItemPickupEvent struct {
	Base
	Character Character `json:"Character"`
	Item      Item      `json:"Item"`
}

// ItemEquipEvent contains data from the player equipping an item
type ItemEquipEvent struct {
	Base
	Character Character `json:"Character"`
	Item      Item      `json:"Item"`
}

// ItemUnequipEvent contains data from the player unequipping an item
type ItemUnequipEvent struct {
	Base
	Character Character `json:"Character"`
	Item      Item      `json:"Item"`
}

// VehicleRideEvent contains data from the player riding in a vehicle
type VehicleRideEvent struct {
	Base
	Character Character `json:"Character"`
	Vehicle   Vehicle   `json:"Vehicle"`
}

// MatchDefinitionEvent contains data about the current match
type MatchDefinitionEvent struct {
	Base
	MatchID     string `json:"MatchId"`
	PingQuality string `json:"PingQuality"`
}

// MatchStartEvent contains data about the start of the match
type MatchStartEvent struct {
	Base
	Characters []Character `json:"Characters"`
}

// GameStatePeriodicEvent contains periodic data about the current state of the match
type GameStatePeriodicEvent struct {
	Base
	GameState GameState `json:"GameState"`
}

// VehicleLeaveEvent contains data from the player leaving a vehicle
type VehicleLeaveEvent struct {
	Base
	Character Character `json:"Character"`
	Vehicle   Vehicle   `json:"Vehicle"`
}

// PlayerTakeDamageEvent contains data from the player taking damage
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

// PlayerLogoutEvent contains data about the player logging out
type PlayerLogoutEvent struct {
	Base
	AccountID string `json:"AccountId"`
}

// ItemAttachEvent contains data from the player attaching an item
type ItemAttachEvent struct {
	Base
	Character  Character `json:"Character"`
	ParentItem Item      `json:"ParentItem"`
	ChildItem  Item      `json:"ChildItem"`
}

// ItemDropEvent contains data from the player dropping an item
type ItemDropEvent struct {
	Base
	Character Character `json:"Character"`
	Item      Item      `json:"Item"`
}

// PlayerKillEvent contains data from the player killing another player
type PlayerKillEvent struct {
	Base
	AttackID           int       `json:"AttackId"`
	Killer             Character `json:"Killer"`
	Victim             Character `json:"Victim"`
	DamageTypeCategory string    `json:"DamageTypeCategory"`
	DamageCauserName   string    `json:"DamageCauserName"`
	Distance           float32   `json:"Distance"`
}

// ItemDetachEvent contains data from the player detaching an item
type ItemDetachEvent struct {
	Base
	Character  Character `json:"Character"`
	ParentItem Item      `json:"ParentItem"`
	ChildItem  Item      `json:"ChildItem"`
}

// ItemUseEvent contains data from the player using an item
type ItemUseEvent struct {
	Base
	Character Character `json:"Character"`
	Item      Item      `json:"Item"`
}

// CarePackageSpawnEvent contains data about a care package spawning
type CarePackageSpawnEvent struct {
	Base
	ItemPackage ItemPackage `json:"ItemPackage"`
}

// VehicleDestroyEvent contains data from a vehicle being destroyed
type VehicleDestroyEvent struct {
	Base
	AttackID           int       `json:"AttackId"`
	Attacker           Character `json:"Attacker"`
	Vehicle            Vehicle   `json:"Vehicle"`
	DamageTypeCategory string    `json:"DamageTypeCategory"`
	DamageCauserName   string    `json:"DamageCauserName"`
	Distance           float32   `json:"Distance"`
}

// CarePackageLandEvent contains data about a care package landing on the map
type CarePackageLandEvent struct {
	Base
	ItemPackage ItemPackage `json:"ItemPackage"`
}

// MatchEndEvent contains data about the match ending
type MatchEndEvent struct {
	Base
	Characters []Character `json:"Characters"`
}
