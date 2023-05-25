package sqliteconnection

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const testDbName = "test.db"

type TestTable struct {
	ID   string
	Name string
}

func TestMustNewConnectionManager(t *testing.T) {
	connMan := MustNewConnectionManager(testDbName, &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold: time.Second, // Slow SQL threshold
				LogLevel:      logger.Info, // Log level
				Colorful:      true,        // Disable color
			},
		),
	})

	assert.NotNil(t, connMan)
}

func TestInsert(t *testing.T) {
	connMan := MustNewConnectionManager(testDbName, &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold: time.Second, // Slow SQL threshold
				LogLevel:      logger.Info, // Log level
				Colorful:      true,        // Disable color
			},
		),
	})
	assert.NotNil(t, connMan)

	defer func() {
		_ = os.Remove(testDbName)
	}()

	connMan.AutoMigrateOrWarn(&TestTable{})

	db, err := connMan.GetConnection()
	assert.NoError(t, err)

	tbl := TestTable{
		ID:   "id",
		Name: "UnitTest",
	}

	result := db.Model(TestTable{}).Create(&tbl)
	assert.NoError(t, result.Error)
	assert.EqualValues(t, 1, result.RowsAffected)
}
