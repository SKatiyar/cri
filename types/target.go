package types

type Target_TargetID string
type Target_SessionID string
type Target_BrowserContextID string
type Target_TargetInfo struct {
	TargetId Target_TargetID  `json:"targetId"`
	Type     string           `json:"type"`
	Title    string           `json:"title"`
	Url      string           `json:"url"`
	Attached bool             `json:"attached"`
	OpenerId *Target_TargetID `json:"openerId,omitempty"`
}
type Target_RemoteLocation struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}
