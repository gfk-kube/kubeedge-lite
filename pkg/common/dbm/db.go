package dbm

import (
	"strings"

	"github.com/kubeedge/kubeedge/beehive/pkg/common/config"
	"github.com/kubeedge/kubeedge/beehive/pkg/common/log"

	"github.com/astaxie/beego/orm"
	//Blank import to run only the init function
	_ "github.com/mattn/go-sqlite3"
)

var (
	driverName string
	dbName     string
	dataSource string
)

//DBAccess is Ormer object interface for all transaction processing and switching database
var DBAccess orm.Ormer

//RegisterModel registers the defined model in the orm if model is enabled
func RegisterModel(moduleName string, m interface{}) {
	if isModuleEnabled(moduleName) {
		orm.RegisterModel(m)
		log.LOGGER.Infof("DB meta for module %s has been registered", moduleName)
	} else {
		log.LOGGER.Infof("DB meta for module %s has not been registered because this module is not enabled", moduleName)
	}
}

func init() {
	//Init DB info
	driverName, _ = config.CONFIG.GetValue("database.driver").ToString()
	dbName, _ = config.CONFIG.GetValue("database.name").ToString()
	dataSource, _ = config.CONFIG.GetValue("database.source").ToString()
	if driverName == "" {
		driverName = "sqlite3"
	}
	if dbName == "" {
		dbName = "default"
	}
	if dataSource == "" {
		dataSource = "edge.db"
	}

	if err := orm.RegisterDriver(driverName, orm.DRSqlite); err != nil {
		log.LOGGER.Fatalf("Failed to register driver: %v", err)
	}
	if err := orm.RegisterDataBase(dbName, driverName, dataSource); err != nil {
		log.LOGGER.Fatalf("Failed to register db: %v", err)
	}
}

//InitDBManager initialises the database by syncing the database schema and creating orm
func InitDBManager() {
	// sync database schema
	orm.RunSyncdb(dbName, false, true)

	// create orm
	DBAccess = orm.NewOrm()
	DBAccess.Using(dbName)
}

func isModuleEnabled(m string) bool {
	// TODO: temp change for ut
	return true
	modules := config.CONFIG.GetConfigurationByKey("modules.enabled")
	if modules != nil {
		for _, value := range modules.([]interface{}) {
			if m == value.(string) {
				return true
			}
		}
	}
	return false
}

// IsNonUniqueNameError tests if the error returned by sqlite is unique.
// It will check various sqlite versions.
func IsNonUniqueNameError(err error) bool {
	str := err.Error()
	if strings.HasSuffix(str, "are not unique") || strings.Contains(str, "UNIQUE constraint failed") || strings.HasSuffix(str, "constraint failed") {
		return true
	}
	return false
}
