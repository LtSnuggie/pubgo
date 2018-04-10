package pubgo

import (
	"encoding/json"
	"time"
)

type StatusResponse struct {
	Data StatusData `json:"data"`
}

type StatusData struct {
	typeIdPair
	Attributes StatusAttributes `json:"attributes"`
}

type StatusAttributes struct {
	Released string `json:"releasedAt"`
	Version  string `json:"version"`
}

type PlayerResponse struct {
	Data []PlayerResponseData `json:"data"`
}

type PlayerResponseData struct {
	typeIdPair
	Attributes    PlayerAttributes `json:"attributes"`
	Relationships Relationships    `json:"relationships"`
}

type PlayerAttributes struct {
	CreatedAt    string          `json:"createdAt"`
	Name         string          `json:"name"`
	PatchVersion string          `json:"patchVersion"`
	ShardID      string          `json:"shardId"`
	Stats        json.RawMessage `json:"stats"`
	TitleID      string          `json:"titleId"`
	Updated      string          `json:"updatedAt"`
}

type Relationships struct {
	Matches Matches `json:"matches"`
}

type Matches struct {
	Data []MatchData `json:"data"`
}

type MatchData struct {
	typeIdPair
}

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
				Data []typeIdPair `json:"data"`
			} `json:"assets"`
			Rosters struct {
				Data []typeIdPair `json:"data"`
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

type MatchParticipant struct {
	typeIdPair
	Attributes struct {
		Actor   string     `json:"actor"`
		ShardID string     `json:"shardId"`
		Stats   MatchStats `json:"stats"`
	} `json:"attributes"`
}

type MatchRoster struct {
	typeIdPair
	Attributes struct {
		ShardID string `json:"shardId"`
		Stats   struct {
			Rank   int `json:"rank"`
			TeamID int `json:"teamId"`
		} `json:"stats"`
		Won bool `json:"won"`
	} `json:"attributes"`
	Relationships struct {
		Participants struct {
			Data []typeIdPair
		} `json:"participants"`
		Team struct {
			Data string `json:"data"`
		} `json:"team"`
	} `json:"relationships"`
}

type MatchAsset struct {
	typeIdPair
	Attributes struct {
		URL         string    `json:"URL"`
		CreatedAt   time.Time `json:"createdAt"`
		Description string    `json:"description"`
		Name        string    `json:"name"`
	} `json:"attributes"`
}

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

type TelemetryResponse struct {
	Events                  []TelemetryEvent
	PlayerLoginEvents       []*PlayerLoginEvent
	PlayerCreateEvents      []*PlayerCreateEvent
	PlayerPositionEvents    []*PlayerPositionEvent
	PlayerAttackEvents      []*PlayerAttackEvent
	ItemPickupEvents        []*ItemPickupEvent
	ItemEquipEvent          []*ItemEquipEvent
	ItemUnequipEvents       []*ItemUnequipEvent
	VehicleRideEvents       []*VehicleRideEvent
	MatchDefinitionEvents   []*MatchDefinitionEvent
	MatchStartEvents        []*MatchStartEvent
	GameStatePeriodicEvents []*GameStatePeriodicEvent
	VehicleLeaveEvents      []*VehicleLeaveEvent
	PlayerTakeDamageEvents  []*PlayerTakeDamageEvent
	PlayerLogoutEvents      []*PlayerLogoutEvent
	ItemAttachEvents        []*ItemAttachEvent
	ItemDropEvents          []*ItemDropEvent
	PlayerKillEvents        []*PlayerKillEvent
	ItemDetachEvents        []*ItemDetachEvent
	ItemUseEvents           []*ItemUseEvent
	CarePackageSpawnEvents  []*CarePackageSpawnEvent
	VehicleDestroyEvents    []*VehicleDestroyEvent
	CarePackageLandEvents   []*CarePackageLandEvent
	MatchEndEvents          []*MatchEndEvent
}

type typeIdPair struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}
