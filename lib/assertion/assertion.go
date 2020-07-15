package assertion

type BasicSpec struct {
	Version string                 `yaml:"version"`
	Meta    map[string]interface{} `yaml:"meta,inline"`
}
type Version struct {
	Version string `yaml:"version"`
}

type SpecV1 struct {
	BasicSpec   `yaml:"basic,inline"`
	Selector    []SelectorDef          `yaml:"selector"`
	Default     map[string]interface{} `yaml:"default"`
	PackageDefs []PackageDef           `yaml:"package"`
	TestDefs    []TestDef              `yaml:"tests"`
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
	Name       string                 `yaml:"name"`
	Using      map[string]string      `yaml:"using"`
	UsingForce map[string]string      `yaml:"using-force"`
	Inherit    []string               `yaml:"inherit"`
	Cases      []TestDef              `yaml:"cases"`
	Assertion  [][]string             `yaml:"assertion"`
	Assert     []string               `yaml:"assert"`
	Meta       map[string]interface{} `yaml:"meta,inline"`
}
