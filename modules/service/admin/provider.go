// Copyright (c) 2021 Terminus, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package admin

import (
	"os"

	"github.com/sirupsen/logrus"

	"github.com/erda-project/erda-infra/base/logs"
	"github.com/erda-project/erda-infra/base/servicehub"
	hs "github.com/erda-project/erda-infra/providers/httpserver"
	"github.com/erda-project/erda-infra/providers/i18n"
	"github.com/erda-project/erda/modules/dop/conf"
	"github.com/erda-project/erda/modules/service/admin/dao"
	"github.com/erda-project/erda/modules/service/admin/manager"
	"github.com/erda-project/erda/pkg/http/httpserver"
	"github.com/erda-project/erda/pkg/jsonstore/etcd"
)

type Config struct {
	Debug bool `default:"false" env:"DEBUG" desc:"enable debug logging"`
}

type provider struct {
	Config Config

	Log    logs.Logger
	Router hs.Router       `autowired:"http-router"`
	CPTran i18n.I18n       `autowired:"i18n@cp"`
	Tran   i18n.Translator `translator:"common"`
}

func (p *provider) Init(ctx servicehub.Context) error {
	p.Log.Info("init admin")
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05.000000000",
	})
	logrus.SetOutput(os.Stdout)

	if p.Config.Debug {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}

	var (
		dbClient *dao.DBClient
		err      error
	)
	if dbClient, err = dao.Open(); err != nil {
		logrus.Fatal(err)
	}
	defer dbClient.Close()

	etcdStore, err := etcd.New()
	if err != nil {
		return err
	}
	admin := manager.NewAdminManager(
		manager.WithDB(dbClient),
		manager.WithBundle(manager.NewBundle()),
		manager.WithETCDStore(etcdStore),
	)
	server := httpserver.New(conf.ListenAddr())
	server.Router().UseEncodedPath()
	server.RegisterEndpoint(admin.Routers())
	p.Router.Any("/**", server.Router())
	return nil
}

func init() {
	servicehub.Register("service.admin", &servicehub.Spec{
		Services:     []string{"admin"},
		Dependencies: []string{"http-server"},
		Description:  "erda platform admin",
		Creator: func() servicehub.Provider {
			return &provider{}
		},
	})
}
