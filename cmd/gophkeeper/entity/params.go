package entity

// Params Contains information about application launch options.
type Params struct {
	ServerAddress string `json:"server_address,omitempty"`
	FileStorePath string `json:"file_store_path,omitempty"`
	ArchiveName   string `json:"archive_name,omitempty"`
	DataFileName  string `json:"data_file_name,omitempty"`
	Session       string `json:"session,omitempty"`
}
