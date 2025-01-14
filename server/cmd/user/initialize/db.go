package initialize

import (
	"fmt"
	"time"

	"github.com/CyanAsterisk/TikGok/server/cmd/user/config"
	"github.com/CyanAsterisk/TikGok/server/shared/consts"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/opentelemetry/logging/logrus"
	"gorm.io/plugin/opentelemetry/tracing"
)

// InitDB to init database
func InitDB() *gorm.DB {
	// config from nacos
	c := config.GlobalServerConfig.MysqlInfo
	dsn := fmt.Sprintf(consts.MySqlDSN, c.User, c.Password, c.Host, c.Port, c.Name)
	newLogger := logger.New(
		logrus.NewWriter(), // io writer write to otel (from otel)
		logger.Config{
			SlowThreshold: time.Second,   // Slow SQL Threshold
			LogLevel:      logger.Silent, // Log level
			Colorful:      true,          // Disable color printing
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		klog.Fatalf("init gorm failed: %s", err)
	}
	if err := db.Use(tracing.NewPlugin()); err != nil { // otel trace (from gorm)
		klog.Fatalf("use tracing plugin failed: %s", err)
	}
	return db
}
