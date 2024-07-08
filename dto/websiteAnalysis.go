package dto

type AnalysisResponseDTO struct {
	Host         string                 `json:"host"`
	Port         int                    `json:"port"`
	Protocol     string                 `json:"protocol"`
	Status       string                 `json:"status"`
	Endpoints    []AnalysisEndpointsDTO `json:"endpoints"`
	ErrorMessage string
}

type AnalysisEndpointsDTO struct {
	IpAddress     string `json:"ipAddress"`
	ServerName    string `json:"serverName"`
	StatusMessage string `json:"statusMessage"`
	Progress      int    `json:"progress"`
}
