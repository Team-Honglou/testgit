package sqlstore

import (
	"github.com/logdisplayplatform/logdisplayplatform/pkg/bus"
	m "github.com/logdisplayplatform/logdisplayplatform/pkg/models"
)

func init() {
	bus.AddHandler("sql", GetDBHealthQuery)
}

func GetDBHealthQuery(query *m.GetDBHealthQuery) error {
	return x.Ping()
}
