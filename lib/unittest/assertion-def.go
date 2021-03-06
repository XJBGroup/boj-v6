package unittest

type BasicSpec struct {
	Version string                 `yaml:"version"`
	Meta    map[string]interface{} `yaml:"meta,inline"`
}
type Version struct {
	Version string `yaml:"version"`
}

type SpecV1 struct {
	BasicSpec   `yaml:"basic,inline"`
	Selector    []SelectorDef            `yaml:"selector"`
	Default     []map[string]interface{} `yaml:"default"`
	PackageDefs []PackageDef             `yaml:"package"`
	TestDefs    []TestDef                `yaml:"tests"`
}

type SelectorDef struct {
	Name string                 `yaml:"name"`
	Case map[string]interface{} `yaml:"selector"`
}

type PackageDef struct {
	Namespace string `yaml:"namespace"`
	Path      string `yaml:"path"`
}

type TestDef struct {
	Name     string                 `yaml:"name"`
	Abstract bool                   `yaml:"abstract"`
	Using    map[string]string      `yaml:"using"`
	Inherit  []string               `yaml:"inherit"`
	Cases    []TestDef              `yaml:"cases"`
	Script   [][]interface{}        `yaml:"script"`
	Assert   []interface{}          `yaml:"assert"`
	Meta     map[string]interface{} `yaml:"meta,inline"`
}
