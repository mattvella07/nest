package nest

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type cameraLastEvent struct {
	HasSound         bool     `json:"has_sound"`
	HasMotion        bool     `json:"has_motion"`
	HasPerson        bool     `json:"has_person"`
	StartTime        string   `json:"start_time"`
	EndTime          string   `json:"end_time"`
	UrlsExpireTime   string   `json:"urls_expire_time"`
	WebURL           string   `json:"web_url"`
	AppURL           string   `json:"app_url"`
	ImageURL         string   `json:"image_url"`
	AnimatedImageURL string   `json:"animated_image_url"`
	ActivityZoneIDs  []string `json:"activity_zone_ids"`
}

type cameraActivityZone struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

// Camera contains all the data for an individual Nest camera
type Camera struct {
	DeviceID              string               `json:"device_id"`
	SoftwareVersion       string               `json:"software_version"`
	StructureID           string               `json:"structure_id"`
	WhereID               string               `json:"where_id"`
	WhereName             string               `json:"where_name"`
	Name                  string               `json:"name"`
	NameLong              string               `json:"name_long"`
	IsOnline              bool                 `json:"is_online"`
	IsStreaming           bool                 `json:"is_streaming"`
	IsAudioInputEnabled   bool                 `json:"is_audio_input_enabled"`
	LastIsOnlineChange    string               `json:"last_is_online_change"`
	IsVideoHistoryEnabled bool                 `json:"is_video_history_enabled"`
	WebURL                string               `json:"web_url"`
	AppURL                string               `json:"app_url"`
	IsPublicShareEnabled  bool                 `json:"is_public_share_enabled"`
	ActivityZones         []cameraActivityZone `json:"activity_zones"`
	PublicShareURL        string               `json:"public_share_url"`
	SnapshotURL           string               `json:"snapshot_url"`
	LastEvent             []cameraLastEvent    `json:"last_event"`
}

// GetCameras returns all Nest cameras along with all their data
func (n *Connection) GetCameras() ([]Camera, error) {
	url := n.setURL("cameras")

	data, err := n.execute(url, "GET", nil)
	if err != nil {
		return []Camera{}, err
	}

	// No cameras found
	if len(data) == 0 {
		return []Camera{}, nil
	}

	// Create map to store JSON response
	fullResponse := make(map[string]interface{})

	err = json.Unmarshal(data, &fullResponse)
	if err != nil {
		return []Camera{}, err
	}

	allCameras := []Camera{}

	for _, val := range fullResponse {
		camera := Camera{}

		c, err := json.Marshal(val)
		if err != nil {
			return []Camera{}, err
		}

		err = json.Unmarshal(c, &camera)
		if err != nil {
			return []Camera{}, err
		}

		allCameras = append(allCameras, camera)
	}

	return allCameras, nil
}

// GetCamera returns the specified Nest camera by device id along with its data
func (n *Connection) GetCamera(deviceID string) (Camera, error) {
	// Error checking
	if strings.Trim(deviceID, " ") == "" {
		return Camera{}, errors.New("Device ID must not be empty")
	}

	url := n.setURL(fmt.Sprintf("cameras/%s", deviceID))

	data, err := n.execute(url, "GET", nil)
	if err != nil {
		return Camera{}, err
	}

	// Camera not found
	if len(data) == 0 {
		return Camera{}, fmt.Errorf("Camera %s not found", deviceID)
	}

	camera := Camera{}

	err = json.Unmarshal(data, &camera)
	if err != nil {
		return Camera{}, err
	}

	return camera, nil
}

// GetCameraSoftwareVersion returns the software version of the specified camera
func (n *Connection) GetCameraSoftwareVersion(deviceID string) (string, error) {
	return n.getValue("cameras", deviceID, "software_version")
}

// GetCameraName returns the name of the specified camera
func (n *Connection) GetCameraName(deviceID string) (string, error) {
	return n.getValue("cameras", deviceID, "name")
}

// IsCameraOnline returns true if the specified camera is online, false if it isn't
func (n *Connection) IsCameraOnline(deviceID string) (string, error) {
	return n.getValue("cameras", deviceID, "is_online")
}

// IsCameraStreaming returns true if the specified camera is actively streaming video, false if it isn't
func (n *Connection) IsCameraStreaming(deviceID string) (string, error) {
	return n.getValue("cameras", deviceID, "is_streaming")
}

// IsCameraAudioInputEnabled returns true if the specified camera mic is listening, false if it isn't
func (n *Connection) IsCameraAudioInputEnabled(deviceID string) (string, error) {
	return n.getValue("cameras", deviceID, "is_audio_input_enabled")
}

// IsCameraVideoHistoryEnabled returns true if the Nest Aware subscription is active, false if it isn't
func (n *Connection) IsCameraVideoHistoryEnabled(deviceID string) (string, error) {
	return n.getValue("cameras", deviceID, "is_video_history_enabled")
}

// GetCameraWebURL returns the web URL of the specified camera
func (n *Connection) GetCameraWebURL(deviceID string) (string, error) {
	return n.getValue("cameras", deviceID, "web_url")
}

// GetCameraAppURL returns the app URL of the specified camera
func (n *Connection) GetCameraAppURL(deviceID string) (string, error) {
	return n.getValue("cameras", deviceID, "app_url")
}

// GetCameraLastEvent returns info for the last event that triggered a notification for the specified camera
func (n *Connection) GetCameraLastEvent(deviceID string) (string, error) {
	return n.getValue("cameras", deviceID, "last_event")
}

// TurnOnStreaming turns on streaming for the specified camera
func (n *Connection) TurnOnStreaming(deviceID string) error {
	vals := make(map[string]interface{})
	vals["is_streaming"] = true

	return n.setValue("cameras", deviceID, vals)
}

// TurnOffStreaming turns off streaming for the specified camera
func (n *Connection) TurnOffStreaming(deviceID string) error {
	vals := make(map[string]interface{})
	vals["is_streaming"] = false

	return n.setValue("cameras", deviceID, vals)
}
