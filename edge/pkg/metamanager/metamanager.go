package metamanager

import (
	"time"

	"github.com/kubeedge/kubeedge/edge/pkg/common/dbm"

	"github.com/astaxie/beego/orm"

	"k8s.io/klog"

	"github.com/kubeedge/beehive/pkg/core"
	beehiveContext "github.com/kubeedge/beehive/pkg/core/context"
	"github.com/kubeedge/beehive/pkg/core/model"
	"github.com/kubeedge/kubeedge/edge/pkg/common/modules"
	metamanagerconfig "github.com/kubeedge/kubeedge/edge/pkg/metamanager/config"
	"github.com/kubeedge/kubeedge/edge/pkg/metamanager/dao"
	"github.com/kubeedge/kubeedge/pkg/apis/edgecore/v1alpha1"
)

//constant metamanager module name
const (
	MetaManagerModuleName = "metaManager"
)

type metaManager struct {
	enable bool
}

func newMetaManager(enable bool) *metaManager {
	return &metaManager{enable: enable}
}

// Register register metamanager
func Register(m *v1alpha1.MetaManager, db *v1alpha1.DataBase) {
	metamanagerconfig.InitConfigure(m)
	dbm.InitDBConfig(db.DriverName, db.AliasName, db.DataSource)
	meta := newMetaManager(m.Enable)
	initDBTable(meta)
	core.Register(meta)
}

// initDBTable create table
func initDBTable(m core.Module) {
	klog.Infof("Begin to register %v db model", m.Name())
	if !m.Enable() {
		klog.Infof("Module %s is disabled, DB meta for it will not be registered", m.Name())
		return
	}
	orm.RegisterModel(new(dao.Meta))
}

func (*metaManager) Name() string {
	return MetaManagerModuleName
}

func (*metaManager) Group() string {
	return modules.MetaGroup
}

func (m *metaManager) Enable() bool {
	return m.enable
}

func (m *metaManager) Start() {

	go func() {
		period := getSyncInterval()
		timer := time.NewTimer(period)
		for {
			select {
			case <-beehiveContext.Done():
				klog.Warning("MetaManager stop")
				return
			case <-timer.C:
				timer.Reset(period)
				msg := model.NewMessage("").BuildRouter(MetaManagerModuleName, GroupResource, model.ResourceTypePodStatus, OperationMetaSync)
				beehiveContext.Send(MetaManagerModuleName, *msg)
			}
		}
	}()

	m.runMetaManager()
}

func getSyncInterval() time.Duration {
	return time.Duration(metamanagerconfig.Get().PodStatusSyncInterval) * time.Second
}
