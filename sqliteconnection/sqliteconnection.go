package sqliteconnection

import (
	"github.com/owenlilly/progorm-connection/connection"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type sqliteConnectionManager struct {
	connection.Manager
}

// MustNewConnectionManager creates an instance of the SQLite implementation of the Manager interface.
func MustNewConnectionManager(dbname string, config *gorm.Config) connection.Manager {
	dialector := sqlite.Open(dbname)
	return connection.MustNewBaseConnectionManager(dbname, dialector, config)
}

// NewConnectionManager creates an instance of the SQLite implementation of the Manager interface.
func NewConnectionManager(dbname string, config *gorm.Config) (connection.Manager, error) {
	dialector := sqlite.Open(dbname)
	conn, err := connection.NewBaseConnectionManager(dbname, dialector, config)
	m := &sqliteConnectionManager{
		Manager: conn,
	}

	return m, err
}
