package entry

type params struct {
	M    string
	Args []interface{}
}

type Request struct {
	Id     string
	Action string
	Params params
}
