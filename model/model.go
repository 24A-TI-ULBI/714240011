package model

// AppInfo represents basic application info
type AppInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Status  string `json:"status"`
}

// DBInfo holds MongoDB connection info
type DBInfo struct {
	DBString string
	DBName   string
}
