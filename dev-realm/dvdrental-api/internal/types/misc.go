package types

type HeartbeatResponse struct {
	Heartbeat int64 `json:"heartbeat"`
}

type MessageResponse struct {
	Message string `json:"message,omitempty"`
}
