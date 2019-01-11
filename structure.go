package nest

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type structureWheres struct {
	WhereID string `json:"where_id"`
	Name    string `json:"name"`
}

// Structure contains all the data for an individual Nest structure
type Structure struct {
	StructureID         string          `json:"structure_id"`
	Thermostats         []string        `json:"thermostats"`
	SmokeCOAlarms       []string        `json:"smoke_co_alarms"`
	Cameras             []string        `json:"cameras"`
	Away                string          `json:"away"`
	Name                string          `json:"name"`
	CountryCode         string          `json:"country_code"`
	PostalCode          string          `json:"postal_code"`
	PeakPeriodStartTime string          `json:"peak_period_start_time"`
	PeakPeriodEndTime   string          `json:"peak_period_end_time"`
	TimeZone            string          `json:"time_zone"`
	ETA                 []string        `json:"eta"`
	ETABegin            string          `json:"eta_begin"`
	RHREnrollment       bool            `json:"rhr_enrollment"`
	WWNSecurityState    string          `json:"wwn_security_state"`
	Wheres              structureWheres `json:"wheres"`
	COAlarmState        string          `json:"co_alarm_state"`
	SmokeAlarmState     string          `json:"smoke_alarm_state"`
}

// GetStructures returns all Nest structures along with all their data
func (n *Connection) GetStructures() ([]Structure, error) {
	url := n.setURL("structures")

	data, err := n.execute(url, "GET", nil)
	if err != nil {
		return []Structure{}, err
	}

	// No structures found
	if len(data) == 0 {
		return []Structure{}, nil
	}

	// Create map to store JSON response
	fullResponse := make(map[string]interface{})

	err = json.Unmarshal(data, &fullResponse)
	if err != nil {
		return []Structure{}, err
	}

	allStructures := []Structure{}

	for _, val := range fullResponse {
		structure := Structure{}

		s, err := json.Marshal(val)
		if err != nil {
			return []Structure{}, err
		}

		err = json.Unmarshal(s, &structure)
		if err != nil {
			return []Structure{}, err
		}

		allStructures = append(allStructures, structure)
	}

	return allStructures, nil
}

// GetStructure returns the specified Nest structure by structure id along with its data
func (n *Connection) GetStructure(structureID string) (Structure, error) {
	// Error checking
	if strings.Trim(structureID, " ") == "" {
		return Structure{}, errors.New("Structure ID must not be empty")
	}

	url := n.setURL(fmt.Sprintf("structures/%s", structureID))

	data, err := n.execute(url, "GET", nil)
	if err != nil {
		return Structure{}, err
	}

	// Structure not found
	if len(data) == 0 {
		return Structure{}, fmt.Errorf("Structure %s not found", structureID)
	}

	structure := Structure{}

	err = json.Unmarshal(data, &structure)
	if err != nil {
		return Structure{}, err
	}

	return structure, nil
}

// GetStructureThermostats returns a list of termostats in the specified structure
func (n *Connection) GetStructureThermostats(structureID string) (string, error) {
	return n.getValue("structures", structureID, "thermostats")
}

// GetStructureSmokeCOAlarms returns a list of smoke/co alarms in the specified structure
func (n *Connection) GetStructureSmokeCOAlarms(structureID string) (string, error) {
	return n.getValue("structures", structureID, "smoke_co_alarms")
}

// GetStructureCameras returns a list of cameras in the specified structure
func (n *Connection) GetStructureCameras(structureID string) (string, error) {
	return n.getValue("structures", structureID, "cameras")
}

// GetStructureOccupancyState returns the occupancy state (home or away) for the specified structure
func (n *Connection) GetStructureOccupancyState(structureID string) (string, error) {
	return n.getValue("structures", structureID, "away")
}

// GetStructureName returns the name of the specified structure
func (n *Connection) GetStructureName(structureID string) (string, error) {
	return n.getValue("structures", structureID, "name")
}
