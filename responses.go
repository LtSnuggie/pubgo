package pubgo

import (
	"encoding/json"
	"time"
)

// StatusResponse is the response payload for the status end point
type StatusResponse struct {
	Data StatusData `json:"data"`
}

// StatusData contains all of the data returned in the StatusResponse
type StatusData struct {
	typeIDPair
	Attributes StatusAttributes `json:"attributes"`
}

// StatusAttributes contains all of the attributes returned in the StatusResponse
type StatusAttributes struct {
	Released string `json:"releasedAt"`
	Version  string `json:"version"`
}

// PlayerResponse is the response payload for the player end point
type PlayerResponse struct {
	Data []PlayerResponseData `json:"data"`
}

// PlayerResponseData contains all of the data returned in the PlayerResponse
type PlayerResponseData struct {
	typeIDPair
	Attributes    PlayerAttributes `json:"attributes"`
	Relationships Relationships    `json:"relationships"`
}

// PlayerAttributes contains all of the player attributes returned in the PlayerResponse
type PlayerAttributes struct {
	CreatedAt    string          `json:"createdAt"`
	Name         string          `json:"name"`
	PatchVersion string          `json:"patchVersion"`
	ShardID      string          `json:"shardId"`
	Stats        json.RawMessage `json:"stats"`
	TitleID      string          `json:"titleId"`
	Updated      string          `json:"updatedAt"`
}

// Relationships contains all of the relationships returned in the PlayerResponse
type Relationships struct {
	Matches Matches `json:"matches"`
}

// Matches contains a slice of all the matches returned in the PlayerResponse
type Matches struct {
	Data []MatchData `json:"data"`
}

// MatchData contains all the match data returned in the PlayerResponse
type MatchData struct {
	typeIDPair
}

// MatchResponse is the response payload for the match end point
type MatchResponse struct {
	Data struct {
		Type       string `json:"type"`
		ID         string `json:"id"`
		Attributes struct {
			CreatedAt    time.Time `json:"createdAt"`
			Duration     int       `json:"duration"`
			GameMode     string    `json:"gameMode"`
			MapName      string    `json:"mapName"`
			PatchVersion string    `json:"patchVersion"`
			ShardID      string    `json:"shardId"`
			Stats        string    `json:"stats"`
			Tags         string    `json:"tags"`
			TitleID      string    `json:"titleId"`
			Description  string    `json:"description"`
			Name         string    `json:"name"`
			URL          string    `json:"URL"`
		} `json:"attributes"`
		Relationships struct {
			Assets struct {
				Data []typeIDPair `json:"data"`
			} `json:"assets"`
			Rosters struct {
				Data []typeIDPair `json:"data"`
			} `json:"rosters"`
		} `json:"relationships"`
		Links struct {
			Schema string `json:"schema"`
			Self   string `json:"self"`
		} `json:"links"`
	} `json:"data"`
	Included     []json.RawMessage `json:"included"`
	Participants []MatchParticipant
	Rosters      []MatchRoster
	Assets       []MatchAsset
	Links        struct {
		Self string `json:"self"`
	} `json:"links"`
	Meta struct{} `json:"meta"`
}

// MatchParticipant contains all the participants returned in the MatchResponse
type MatchParticipant struct {
	typeIDPair
	Attributes struct {
		Actor   string     `json:"actor"`
		ShardID string     `json:"shardId"`
		Stats   MatchStats `json:"stats"`
	} `json:"attributes"`
}

// MatchRoster contains all the rosters returned in the MatchResponse
type MatchRoster struct {
	typeIDPair
	Attributes struct {
		ShardIDn string `json:"shardId"`
		Stats    struct {
			Rank   int `json:"rank"`
			TeamID int `json:"teamId"`
		} `json:"stats"`
		Won bool `json:"won"`
	} `json:"attributes"`
	Relationships struct {
		Participants struct {
			Data []typeIDPair
		} `json:"participants"`
		Team struct {
			Data string `json:"data"`
		} `json:"team"`
	} `json:"relationships"`
}

// MatchAsset contains all the assets returned in the MatchResponse
type MatchAsset struct {
	typeIDPair
	Attributes struct {
		URL         string    `json:"URL"`
		CreatedAt   time.Time `json:"createdAt"`
		Description string    `json:"description"`
		Name        string    `json:"name"`
	} `json:"attributes"`
}

// MatchStats are all the stats returned in the MatchResponse
type MatchStats struct {
	DBNOs           int     `json:"DBNOs"`
	Assists         int     `json:"assists"`
	Boosts          int     `json:"boosts"`
	DamageDealt     float32 `json:"damageDealt"`
	DeathType       string  `json:"deathType"`
	HeadshotKills   int     `json:"headshotKills"`
	Heals           int     `json:"heals"`
	KillPlace       int     `json:"killPlace"`
	KillPoints      int     `json:"killPoints"`
	KillPointsDelta int     `json:"killPointsDelta"`
	KillStreaks     int     `json:"killStreaks"`
	Kills           int     `json:"kills"`
	LastKillPoints  int     `json:"lastKillPoints"`
	LastWinPoints   int     `json:"lastWinPoints"`
	LongestKill     int     `json:"longestKill"`
	MostDamage      int     `json:"mostDamage"`
	Name            string  `json:"name"`
	PlayerID        string  `json:"playerId"`
	Revives         int     `json:"revives"`
	RideDistance    float32 `json:"rideDistance"`
	RoadKills       int     `json:"roadKills"`
	TeamKills       int     `json:"teamKills"`
	TimeSurvived    int     `json:"timeSurvived"`
	VehicleDestroys int     `json:"vehicleDestroys"`
	WalkDistance    float32 `json:"walkDistance"`
	WeaponsAcquired int     `json:"weaponsAcquired"`
	WinPlace        int     `json:"winPlace"`
	WinPoints       int     `json:"winPoints"`
	WinPointsDelta  int     `json:"winPointsDelta"`
}

// TelemetryResponse is the response payload for the telemetry end point
type TelemetryResponse struct {
	Events                  []TelemetryEvent          //All telemetry events in chronoligic order
	PlayerLoginEvents       []*PlayerLoginEvent       //All PlayerLoginEvent's in chronoligic order
	PlayerCreateEvents      []*PlayerCreateEvent      //All PlayerCreateEvent's in chronoligic order
	PlayerPositionEvents    []*PlayerPositionEvent    //All PlayerPositionEvent's in chronoligic order
	PlayerAttackEvents      []*PlayerAttackEvent      //All PlayerAttackEvent's in chronoligic order
	ItemPickupEvents        []*ItemPickupEvent        //All ItemPickupEvent's in chronoligic order
	ItemEquipEvent          []*ItemEquipEvent         //All ItemEquipEvent's in chronoligic order
	ItemUnequipEvents       []*ItemUnequipEvent       //All ItemUnequipEvent's in chronoligic order
	VehicleRideEvents       []*VehicleRideEvent       //All VehicleRideEvent's in chronoligic order
	MatchDefinitionEvents   []*MatchDefinitionEvent   //All MatchDefinitionEvent's in chronoligic order
	MatchStartEvents        []*MatchStartEvent        //All MatchStartEvent's in chronoligic order
	GameStatePeriodicEvents []*GameStatePeriodicEvent //All GameStatePeriodicEvent's in chronoligic order
	VehicleLeaveEvents      []*VehicleLeaveEvent      //All VehicleLeaveEvent's in chronoligic order
	PlayerTakeDamageEvents  []*PlayerTakeDamageEvent  //All PlayerTakeDamageEvent's in chronoligic order
	PlayerLogoutEvents      []*PlayerLogoutEvent      //All PlayerLogoutEvent's in chronoligic order
	ItemAttachEvents        []*ItemAttachEvent        //All ItemAttachEvent's in chronoligic order
	ItemDropEvents          []*ItemDropEvent          //All ItemDropEvent's in chronoligic order
	PlayerKillEvents        []*PlayerKillEvent        //All PlayerKillEvent's in chronoligic order
	ItemDetachEvents        []*ItemDetachEvent        //All ItemDetachEvent's in chronoligic order
	ItemUseEvents           []*ItemUseEvent           //All ItemUseEvent's in chronoligic order
	CarePackageSpawnEvents  []*CarePackageSpawnEvent  //All CarePackageSpawnEvent's in chronoligic order
	VehicleDestroyEvents    []*VehicleDestroyEvent    //All VehicleDestroyEvent's in chronoligic order
	CarePackageLandEvents   []*CarePackageLandEvent   //All CarePackageLandEvent's in chronoligic order
	MatchEndEvents          []*MatchEndEvent          //All MatchEndEvent's in chronoligic order
}

// typeIDPair is a common pattern used throughout all responses
type typeIDPair struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}
