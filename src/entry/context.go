package entry

type Context struct {
	host       string
	port       int
	configfile string
	logfile    string
}

func (ctx *Context) SetHost(host string) {
	ctx.host = host
}

func (ctx *Context) GetHost() string {
	return ctx.host
}

func (ctx *Context) SetPort(port int) {
	ctx.port = port
}

func (ctx *Context) GetPort() int {
	return ctx.port
}

func (ctx *Context) SetConfigFile(configfile string) {
	ctx.configfile = configfile
}

func (ctx *Context) GetConfigFile() string {
	return ctx.configfile
}

func (ctx *Context) SetLogFile(logfile string) {
	ctx.logfile = logfile
}

func (ctx *Context) GetLogFile() string {
	return ctx.logfile
}
