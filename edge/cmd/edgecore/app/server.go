package app

import (
	"errors"
	"fmt"
	"net"
	"os"

	"github.com/mitchellh/go-ps"
	"github.com/spf13/cobra"
	netutil "k8s.io/apimachinery/pkg/util/net"
	cliflag "k8s.io/component-base/cli/flag"
	"k8s.io/component-base/cli/globalflag"
	"k8s.io/component-base/term"
	"k8s.io/klog/v2"

	"github.com/kubeedge/beehive/pkg/core"
	"github.com/kubeedge/kubeedge/edge/cmd/edgecore/app/options"
	"github.com/kubeedge/kubeedge/edge/pkg/common/dbm"
	"github.com/kubeedge/kubeedge/edge/pkg/devicetwin"
	"github.com/kubeedge/kubeedge/edge/pkg/edged"
	"github.com/kubeedge/kubeedge/edge/pkg/edgehub"
	"github.com/kubeedge/kubeedge/edge/pkg/edgestream"
	"github.com/kubeedge/kubeedge/edge/pkg/eventbus"
	"github.com/kubeedge/kubeedge/edge/pkg/metamanager"
	"github.com/kubeedge/kubeedge/edge/pkg/servicebus"
	"github.com/kubeedge/kubeedge/edge/test"
	"github.com/kubeedge/kubeedge/pkg/apis/componentconfig/edgecore/v1alpha1"
	"github.com/kubeedge/kubeedge/pkg/apis/componentconfig/edgecore/v1alpha1/validation"
	"github.com/kubeedge/kubeedge/pkg/features"
	"github.com/kubeedge/kubeedge/pkg/util"
	"github.com/kubeedge/kubeedge/pkg/util/flag"
	"github.com/kubeedge/kubeedge/pkg/version"
	"github.com/kubeedge/kubeedge/pkg/version/verflag"
)

// NewEdgeCoreCommand create edgecore cmd
func NewEdgeCoreCommand() *cobra.Command {
	opts := options.NewEdgeCoreOptions()
	cmd := &cobra.Command{
		Use: "edgecore",
		Long: `Edgecore is the core edge part of KubeEdge, which contains six modules: devicetwin, edged,
edgehub, eventbus, metamanager, and servicebus. DeviceTwin is responsible for storing device status
and syncing device status to the cloud. It also provides query interfaces for applications. Edged is an
agent that runs on edge nodes and manages containerized applications and devices. Edgehub is a web socket
client responsible for interacting with Cloud Service for the edge computing (like Edge Controller as in the KubeEdge
Architecture). This includes syncing cloud-side resource updates to the edge, and reporting
edge-side host and device status changes to the cloud. EventBus is a MQTT client to interact with MQTT
servers (mosquito), offering publish and subscribe capabilities to other components. MetaManager
is the message processor between edged and edgehub. It is also responsible for storing/retrieving metadata
to/from a lightweight database (SQLite).ServiceBus is a HTTP client to interact with HTTP servers (REST),
offering HTTP client capabilities to components of cloud to reach HTTP servers running at edge. `,
		Run: func(cmd *cobra.Command, args []string) {
			verflag.PrintAndExitIfRequested()
			flag.PrintMinConfigAndExitIfRequested(v1alpha1.NewMinEdgeCoreConfig())
			flag.PrintDefaultConfigAndExitIfRequested(v1alpha1.NewDefaultEdgeCoreConfig())
			flag.PrintFlags(cmd.Flags())

			if errs := opts.Validate(); len(errs) > 0 {
				klog.Exit(util.SpliceErrors(errs))
			}

			config, err := opts.Config()
			if err != nil {
				klog.Exit(err)
			}
			if errs := validation.ValidateEdgeCoreConfiguration(config); len(errs) > 0 {
				klog.Exit(util.SpliceErrors(errs.ToAggregate().Errors()))
			}

			if err := features.DefaultMutableFeatureGate.SetFromMap(config.FeatureGates); err != nil {
				klog.Exit(err)
			}

			// To help debugging, immediately log version
			klog.Infof("Version: %+v", version.Get())

			// Check the running environment by default
			checkEnv := os.Getenv("CHECK_EDGECORE_ENVIRONMENT")
			// Force skip check if enable metaserver
			if config.Modules.MetaManager.MetaServer.Enable {
				checkEnv = "false"
			}
			if checkEnv != "false" {
				// Check running environment before run edge core
				if err := environmentCheck(); err != nil {
					klog.Exit(fmt.Errorf("failed to check the running environment: %v", err))
				}
			}

			// Get edge node local ip only when the custiomInterfaceName has been set.
			// Defaults to the local IP from the default interface by the default config
			if config.Modules.Edged.CustomInterfaceName != "" {
				ip, err := netutil.ChooseBindAddressForInterface(config.Modules.Edged.CustomInterfaceName)
				if err != nil {
					klog.Errorf("Failed to get IP address by custom interface %s, err: %v", config.Modules.Edged.CustomInterfaceName, err)
					os.Exit(1)
				}
				config.Modules.Edged.NodeIP = ip.String()
				klog.Infof("Get IP address by custom interface successfully, %s: %s", config.Modules.Edged.CustomInterfaceName, config.Modules.Edged.NodeIP)
			} else {
				if net.ParseIP(config.Modules.Edged.NodeIP) != nil {
					klog.Infof("Use node IP address from config: %s", config.Modules.Edged.NodeIP)
				} else if config.Modules.Edged.NodeIP != "" {
					klog.Errorf("invalid node IP address specified: %s", config.Modules.Edged.NodeIP)
					os.Exit(1)
				} else {
					nodeIP, err := util.GetLocalIP(util.GetHostname())
					if err != nil {
						klog.Errorf("Failed to get Local IP address: %v", err)
						os.Exit(1)
					}
					config.Modules.Edged.NodeIP = nodeIP
					klog.Infof("Get node local IP address successfully: %s", nodeIP)
				}
			}

			registerModules(config)
			// start all modules
			core.Run()
		},
	}
	fs := cmd.Flags()
	namedFs := opts.Flags()
	flag.AddFlags(namedFs.FlagSet("global"))
	verflag.AddFlags(namedFs.FlagSet("global"))
	globalflag.AddGlobalFlags(namedFs.FlagSet("global"), cmd.Name())
	for _, f := range namedFs.FlagSets {
		fs.AddFlagSet(f)
	}

	usageFmt := "Usage:\n  %s\n"
	cols, _, _ := term.TerminalSize(cmd.OutOrStdout())
	cmd.SetUsageFunc(func(cmd *cobra.Command) error {
		fmt.Fprintf(cmd.OutOrStderr(), usageFmt, cmd.UseLine())
		cliflag.PrintSections(cmd.OutOrStderr(), namedFs, cols)
		return nil
	})
	cmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		fmt.Fprintf(cmd.OutOrStdout(), "%s\n\n"+usageFmt, cmd.Long, cmd.UseLine())
		cliflag.PrintSections(cmd.OutOrStdout(), namedFs, cols)
	})

	return cmd
}

// environmentCheck check the environment before edgecore start
// if Check failed,  return errors
func environmentCheck() error {
	processes, err := ps.Processes()
	if err != nil {
		return err
	}

	for _, process := range processes {
		switch process.Executable() {
		case "kubelet": // if kubelet is running, return error
			return errors.New("kubelet should not running on edge node when running edgecore")
		case "kube-proxy": // if kube-proxy is running, return error
			return errors.New("kube-proxy should not running on edge node when running edgecore")
		}
	}

	return nil
}

// registerModules register all the modules started in edgecore
func registerModules(c *v1alpha1.EdgeCoreConfig) {
	devicetwin.Register(c.Modules.DeviceTwin, c.Modules.Edged.HostnameOverride)
	edged.Register(c.Modules.Edged)
	edgehub.Register(c.Modules.EdgeHub, c.Modules.Edged.HostnameOverride)
	eventbus.Register(c.Modules.EventBus, c.Modules.Edged.HostnameOverride)
	metamanager.Register(c.Modules.MetaManager)
	servicebus.Register(c.Modules.ServiceBus)
	edgestream.Register(c.Modules.EdgeStream, c.Modules.Edged.HostnameOverride, c.Modules.Edged.NodeIP)
	test.Register(c.Modules.DBTest)
	// Note: Need to put it to the end, and wait for all models to register before executing
	dbm.InitDBConfig(c.DataBase.DriverName, c.DataBase.AliasName, c.DataBase.DataSource)
}
