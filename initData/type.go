package initData

type noteList []*Note

type Note struct {
	Name               string
	path               string
	statusMap          map[status]bool
	unrecognizedStatus []string
	rowYaml            string
	yamlData           yamlData
}

type status int

type yamlData struct {
	Uuid       string   `yaml:"uuid"`
	TagList    yamlList `yaml:"tags"`
	StatusList yamlList `yaml:"status"`
}

type yamlList []string

const (
	DELETE status = iota
	TEMP
)

var statusName = map[status]string{
	DELETE: "DELETE",
	TEMP:   "TEMP",
}

func (s status) String() string {
	return statusName[s]
}
