package main

import (
	"dienlanhphongvan-cdn/app/entity"
	"dienlanhphongvan-cdn/cmd"
	"dienlanhphongvan-cdn/config"
	"dienlanhphongvan-cdn/middleware"
	"dienlanhphongvan-cdn/route"
	"dienlanhphongvan-cdn/webservice"
	"fmt"
	"utilities/ulog"
)

func init() {
	cmd.Root().Use = "rsimgx"
	cmd.Root().Short = "rsimgx: REST service used to crop, resize, compress images"
	cmd.Root().Long = "rsimgx: REST service used to crop, resize, compress images"

	cmd.SetRunFunc(run)
}

func main() {
	cmd.Execute()
}

func run() {
	conf := config.Load()

	var (
		log        = newLogger(conf.Log.Dir, conf.Log.LevelDebug)
		compressor = entity.NewCompress(conf.Compressor.Enable, conf.Compressor.Exec)
		convertor  = entity.NewConvertor(conf.Convertor.Enable, conf.Convertor.Exec)
		entieis    = entity.NewImage(convertor, compressor)
		image      = route.NewImage(entieis)
		r          = webservice.NewWebservice(webservice.Config{
			Host:      conf.Web.Host,
			Port:      conf.Web.Port,
			DebugMode: conf.Web.Debug,
		}, log)
	)

	// middleware
	r.Use(middleware.RewriteStatus())
	r.Use(middleware.InfoLogger(log))
	r.Use(middleware.ErrorLogger(log))

	// api
	api := r.Group("/v1/images")
	r.Post(api, "/compress", image.Compress)
	r.Post(api, "/resize", image.Resize)
	r.Post(api, "/crop", image.Crop)

	// start
	fmt.Println("Listen on", r.Address())
	r.Serve()
}

func newLogger(logDir string, logDebug bool) *ulog.Ulogger {
	ret := ulog.NewLogger(logDir, "", logDebug)
	return &ret
}
