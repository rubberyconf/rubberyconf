package datasource

import (
	"log"
	"strings"
	"sync"

	"github.com/rubberyconf/rubberyconf/internal/config"
)

type IDataSource interface {
	GetFeature() (interface{}, bool)
}

type DataSource struct {
	Url string
}

var (
	source   *DataSource
	onceGogs sync.Once
)

const (
	GOGS string = "Gogs"
)

func SelectSource() *DataSource {

	conf := config.GetConfiguration()
	var res *DataSource
	typeSource := conf.Api.SourceType
	if typeSource == GOGS {
		res = NewDataSourceGogs()
	} else {
		log.Fatal("no data source selected")
	}

	return res
}

func GetSource() *DataSource {
	return source
}
func NewDataSourceGogs() *DataSource {

	onceGogs.Do(func() {
		conf := config.GetConfiguration()
		source = new(DataSource)
		source.Url = strings.Join([]string{conf.GitServer.Url, "/raw/"}, "")
	})
	return source
}
