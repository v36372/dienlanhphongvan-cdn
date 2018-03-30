package webservice

import (
	"dienlanhphongvan-cdn/util"
	"fmt"
	"net/http"
	"utilities/ulog"

	"github.com/facebookgo/grace/gracehttp"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type Config struct {
	Host      string
	Port      string
	DebugMode bool
}

func (conf Config) Address() string {
	return fmt.Sprintf("%s:%s", conf.Host, conf.Port)
}

type Webservice struct {
	address string
	engine  *gin.Engine
}

func NewWebservice(conf Config, log *ulog.Ulogger) *Webservice {
	binding.Validator = util.Validator()
	var engine *gin.Engine
	if conf.DebugMode {
		gin.SetMode(gin.DebugMode)
		engine = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		engine = gin.New()
		engine.Use(gin.Recovery())
	}

	return &Webservice{
		address: conf.Address(),
		engine:  engine,
	}
}

func (w Webservice) Address() string {
	return w.address
}

func (w *Webservice) Use(middleware ...gin.HandlerFunc) {
	w.engine.Use(middleware...)
}

func (w *Webservice) Serve() {
	server := http.Server{
		Addr:    w.address,
		Handler: w.engine,
	}
	if err := gracehttp.Serve(&server); err != nil {
		panic(err)
	}
}

func (w *Webservice) Group(path string) *gin.RouterGroup {
	return w.engine.Group(path)
}

func (Webservice) Get(group *gin.RouterGroup, relativePath string, f func(*gin.Context)) {
	route(group, "GET", relativePath, f)
}

func (Webservice) Post(group *gin.RouterGroup, relativePath string, f func(*gin.Context)) {
	route(group, "POST", relativePath, f)
}

func (Webservice) Put(group *gin.RouterGroup, relativePath string, f func(*gin.Context)) {
	route(group, "PUT", relativePath, f)
}

func (Webservice) Delete(group *gin.RouterGroup, relativePath string, f func(*gin.Context)) {
	route(group, "DELETE", relativePath, f)
}

func route(group *gin.RouterGroup, method string, relativePath string, f func(*gin.Context)) {
	switch method {
	case "POST":
		group.POST(relativePath, f)
	case "GET":
		group.GET(relativePath, f)
	case "PUT":
		group.PUT(relativePath, f)
	case "DELETE":
		group.DELETE(relativePath, f)
	}
}
