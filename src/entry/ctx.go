package entry

type Ctx struct {
	host       string
	port       int
	configfile string
	logfile    string
}

func (ctx *Ctx) SetHost(host string) {
	ctx.host = host
}

func (ctx *Ctx) GetHost() string {
	return ctx.host
}

func (ctx *Ctx) SetPort(port int) {
	ctx.port = port
}

func (ctx *Ctx) GetPort() int {
	return ctx.port
}

func (ctx *Ctx) SetConfigFile(configfile string) {
	ctx.configfile = configfile
}

func (ctx *Ctx) GetConfigFile() string {
	return ctx.configfile
}

func (ctx *Ctx) SetLogFile(logfile string) {
	ctx.logfile = logfile
}

func (ctx *Ctx) GetLogFile() string {
	return ctx.logfile
}
