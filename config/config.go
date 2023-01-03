package config

type Conf struct {
	Server   Server   `json:"server"`
	Jwt      Jwt      `json:"jwt"`
	Database Database `json:"database"`
}
