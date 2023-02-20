package core

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/viper"
	"os"
	"strconv"
	"strings"
)

type NacosConfig struct {
	ServerConfigs []constant.ServerConfig
	ClientConfig  *constant.ClientConfig
	ConfigId      string
	Group         string
	SharedConfigs []SharedConfig
}

type SharedConfig struct {
	DataId  string
	Group   string
	Refresh bool
}

func initNacos() (*NacosConfig, error) {
	nac := new(NacosConfig)
	vs := viper.New()
	vs.SetConfigName("nacos.yaml")
	vs.AddConfigPath("./config")
	vs.SetConfigType("yaml")
	if err := vs.ReadInConfig(); err != nil {
		Error(err.Error())
		return nil, err
	}

	if err := vs.Unmarshal(nac); err != nil {
		Error(err.Error())
		return nil, err
	}
	//fmt.Println(nac)
	return nac, nil
}

func (nac *NacosConfig) initNacosCfg() (*Config, error) {
	config := new(Config)
	client, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  nac.ClientConfig,
			ServerConfigs: nac.ServerConfigs,
		},
	)
	if err != nil {
		Error(err.Error())
		return config, err
	}

	content, err := client.GetConfig(vo.ConfigParam{
		DataId: nac.ConfigId,
		Group:  nac.Group,
	})
	if len(nac.SharedConfigs) > 0 {
		for _, v := range nac.SharedConfigs {
			cont, err := client.GetConfig(vo.ConfigParam{
				DataId: v.DataId,
				Group:  v.Group,
			})
			if err != nil {
				Error(err.Error())
				os.Exit(1)
			}
			content += "\n" + cont
		}
	}
	if content != "" && strings.Index(content, "AppName") == 0 {
		return dealContent2Config(content, config)
	}

	return config, errors.New("读取远程配置失败")
}

// 仅做测试，实际使用k8s服务名，不需要nacos做服务发现
// 仅供本机测试环境参考

func dealContent2Config(content string, config *Config) (*Config, error) {
	vs := viper.New()
	vs.SetConfigType("yaml")
	vs.ReadConfig(bytes.NewReader([]byte(content)))
	if err := vs.Unmarshal(config); err != nil {
		Error(err.Error())
		return config, err
	}
	return config, nil
}

func (nac *NacosConfig) AddNacosListener() {

	client, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  nac.ClientConfig,
			ServerConfigs: nac.ServerConfigs,
		},
	)
	if err != nil {
		Error(err.Error())

	}

	err = client.ListenConfig(vo.ConfigParam{
		DataId: nac.ConfigId,
		Group:  nac.Group,
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("ListenConfig group:" + group + ", dataId:" + dataId + ", data:" + data)
			if len(data) == 0 || strings.Index(data, "AppName") != 0 {
				Error("getNacosConfigData - LoadErr , len is 0")
			} else {
				Cfg, err = nac.initNacosCfg()
				if err != nil {
					Error((err.Error()))
				}
			}
		},
	})

	if len(nac.SharedConfigs) > 0 {
		for _, v := range nac.SharedConfigs {
			if v.Refresh {
				err = client.ListenConfig(vo.ConfigParam{
					DataId: v.DataId,
					Group:  v.Group,
					OnChange: func(namespace, group, dataId, data string) {
						fmt.Println("ListenConfig group:" + group + ", dataId:" + dataId + ", data:" + data)
						if len(data) == 0 {
							Error("getNacosConfigData - LoadErr , len is 0")
						} else {
							Cfg, err = nac.initNacosCfg()
							if err != nil {
								Error((err.Error()))
							}
						}
					},
				})
			}
		}
	}

}

func (nac *NacosConfig) initNamingClient() error {
	namingClient, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  nac.ClientConfig,
			ServerConfigs: nac.ServerConfigs,
		},
	)
	if err != nil {
		fmt.Errorf("初始化nacos 失败: %s", err.Error())
		return err
	}
	port, _ := strconv.Atoi(Cfg.GrpcSettings.Port)
	success, err := namingClient.RegisterInstance(vo.RegisterInstanceParam{
		//Ip:        "serverb-srv.test.svc.cluster.local"
		Ip:          "127.0.0.1",
		Port:        uint64(port),
		ServiceName: Cfg.AppName,
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
		//Metadata:    map[string]string{"name": "test"},
		ClusterName: "DEFAULT", // 默认值DEFAULT
		GroupName:   nac.Group, // 默认值DEFAULT_GROUP
	})
	if err != nil {
		fmt.Errorf("注册服务失败: %s", err.Error())
		return err
	}
	if !success {
		Error("注册服务失败")
		return errors.New("注册服务失败")
	}
	return nil
}
