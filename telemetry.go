package pubgo

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// unmarshalEvent takes in the raw json and determines which event to unmarshal
// the json into. Additionally, it adds a pointer to the event to some helper
// slices to allow users to quickly parse specific events. Everything is kept
// in chronological order.
func (tr *TelemetryResponse) unmarshalEvent(js []byte, t string) {
	switch strings.ToLower(t) {
	case playerLogin:
		v := PlayerLoginEvent{}
		json.Unmarshal(js, &v)
		tr.Events = append(tr.Events, v)
		tr.PlayerLoginEvents = append(tr.PlayerLoginEvents, &v)
	case playerCreate:
		v := PlayerCreateEvent{}
		json.Unmarshal(js, &v)
		tr.Events = append(tr.Events, v)
		tr.PlayerCreateEvents = append(tr.PlayerCreateEvents, &v)
	case playerPosition:
		v := PlayerPositionEvent{}
		json.Unmarshal(js, &v)
		tr.Events = append(tr.Events, v)
		tr.PlayerPositionEvents = append(tr.PlayerPositionEvents, &v)
	case playerAttack:
		v := PlayerAttackEvent{}
		json.Unmarshal(js, &v)
		tr.Events = append(tr.Events, v)
		tr.PlayerAttackEvents = append(tr.PlayerAttackEvents, &v)
	case itemPickup:
		v := ItemPickupEvent{}
		json.Unmarshal(js, &v)
		tr.Events = append(tr.Events, v)
		tr.ItemPickupEvents = append(tr.ItemPickupEvents, &v)
	case itemEquip:
		v := ItemEquipEvent{}
		json.Unmarshal(js, &v)
		tr.Events = append(tr.Events, v)
		tr.ItemEquipEvent = append(tr.ItemEquipEvent, &v)
	case itemUnequip:
		v := ItemUnequipEvent{}
		json.Unmarshal(js, &v)
		tr.Events = append(tr.Events, v)
		tr.ItemUnequipEvents = append(tr.ItemUnequipEvents, &v)
	case vehicleRide:
		v := VehicleRideEvent{}
		json.Unmarshal(js, &v)
		tr.Events = append(tr.Events, v)
		tr.VehicleRideEvents = append(tr.VehicleRideEvents, &v)
	case matchDefinition:
		v := MatchDefinitionEvent{}
		json.Unmarshal(js, &v)
		tr.Events = append(tr.Events, v)
		tr.MatchDefinitionEvents = append(tr.MatchDefinitionEvents, &v)
	case matchStart:
		v := MatchStartEvent{}
		json.Unmarshal(js, &v)
		tr.Events = append(tr.Events, v)
		tr.MatchStartEvents = append(tr.MatchStartEvents, &v)
	case gameStatePeriodic:
		v := GameStatePeriodicEvent{}
		json.Unmarshal(js, &v)
		tr.Events = append(tr.Events, v)
		tr.GameStatePeriodicEvents = append(tr.GameStatePeriodicEvents, &v)
	case vehicleLeave:
		v := VehicleLeaveEvent{}
		json.Unmarshal(js, &v)
		tr.Events = append(tr.Events, v)
		tr.VehicleLeaveEvents = append(tr.VehicleLeaveEvents, &v)
	case playerTakeDamage:
		v := PlayerTakeDamageEvent{}
		json.Unmarshal(js, &v)
		tr.Events = append(tr.Events, v)
		tr.PlayerTakeDamageEvents = append(tr.PlayerTakeDamageEvents, &v)
	case playerLogout:
		v := PlayerLogoutEvent{}
		json.Unmarshal(js, &v)
		tr.Events = append(tr.Events, v)
		tr.PlayerLogoutEvents = append(tr.PlayerLogoutEvents, &v)
	case itemAttach:
		v := ItemAttachEvent{}
		json.Unmarshal(js, &v)
		tr.Events = append(tr.Events, v)
		tr.ItemAttachEvents = append(tr.ItemAttachEvents, &v)
	case itemDrop:
		v := ItemDropEvent{}
		json.Unmarshal(js, &v)
		tr.Events = append(tr.Events, v)
		tr.ItemDropEvents = append(tr.ItemDropEvents, &v)
	case playerKill:
		v := PlayerKillEvent{}
		json.Unmarshal(js, &v)
		tr.Events = append(tr.Events, v)
		tr.PlayerKillEvents = append(tr.PlayerKillEvents, &v)
	case itemDetach:
		v := ItemDetachEvent{}
		json.Unmarshal(js, &v)
		tr.Events = append(tr.Events, v)
		tr.ItemDetachEvents = append(tr.ItemDetachEvents, &v)
	case itemUse:
		v := ItemUseEvent{}
		json.Unmarshal(js, &v)
		tr.Events = append(tr.Events, v)
		tr.ItemUseEvents = append(tr.ItemUseEvents, &v)
	case carePackageSpawn:
		v := CarePackageSpawnEvent{}
		json.Unmarshal(js, &v)
		tr.Events = append(tr.Events, v)
		tr.CarePackageSpawnEvents = append(tr.CarePackageSpawnEvents, &v)
	case vehicleDestroy:
		v := VehicleDestroyEvent{}
		json.Unmarshal(js, &v)
		tr.Events = append(tr.Events, v)
		tr.VehicleDestroyEvents = append(tr.VehicleDestroyEvents, &v)
	case carePackageLand:
		v := CarePackageLandEvent{}
		json.Unmarshal(js, &v)
		tr.Events = append(tr.Events, v)
		tr.CarePackageLandEvents = append(tr.CarePackageLandEvents, &v)
	case matchEnd:
		v := MatchEndEvent{}
		json.Unmarshal(js, &v)
		tr.Events = append(tr.Events, v)
		tr.MatchEndEvents = append(tr.MatchEndEvents, &v)
	}
	return
}

// ToFile will save a TelemetryResponse to the file at a specified location.
// These data are always static and so it makes sense to cache/save this somewhere
// locally to prevent from having to request the large file multiple times
func (tr *TelemetryResponse) ToFile(path string) (err error) {
	var b []byte
	b, err = json.Marshal(tr.Events)
	if err != nil {
		return
	}
	var pretty bytes.Buffer
	err = json.Indent(&pretty, b, "", "\t")
	if err != nil {
		return
	}
	f, err := os.Create(path)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	w := bufio.NewWriter(f)
	_, err = w.Write(pretty.Bytes())
	if err != nil {
		fmt.Println(err.Error())
	}
	pretty.Reset()
	err = w.Flush()
	if err != nil {
		fmt.Println(err.Error())
	}
	f.Close()
	return
}
