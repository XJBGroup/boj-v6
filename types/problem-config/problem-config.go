package problemconfig

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"github.com/Myriad-Dreamin/boj-v6/lib/unit/memory"
	"github.com/Myriad-Dreamin/boj-v6/lib/unit/time"
	"github.com/Myriad-Dreamin/boj-v6/types"
	"strconv"
	"strings"
)

type kiloByte = int

type TaskConfig struct {
	Name        string `json:"name" yaml:"name" toml:"name" xml:"name"`
	CaseCount   int    `json:"case-count" yaml:"case-count" toml:"case-count" xml:"case-count"`
	TimeLimit   int64  `json:"time-limit" yaml:"time-limit" toml:"time-limit" xml:"time-limit"`
	MemoryLimit int64  `json:"memory-limit" yaml:"memory-limit" toml:"memory-limit" xml:"memory-limit"`
	Score       int    `json:"score" yaml:"score" toml:"score" xml:"score"`
	InputPath   string `json:"input-path" yaml:"input-path" toml:"input-path" xml:"input-path"`
	OutputPath  string `json:"output-path" yaml:"output-path" toml:"output-path" xml:"output-path"`
}

type JudgeConfig struct {
	Type  string       `json:"judge-type" yaml:"judge-type" toml:"judge-type" xml:"judge-type"`
	Tasks []TaskConfig `json:"tasks" yaml:"tasks" toml:"tasks" xml:"tasks"`
}

type SpecialJudgeConfig struct {
	SpecialJudge uint8  `json:"enable" yaml:"enable" toml:"enable" xml:"enable"`
	LanguageType uint8  `json:"language-type" yaml:"language-type" toml:"language-type" xml:"language-type"`
	FilePath     string `json:"file-path" yaml:"file-path" toml:"file-path" xml:"file-path"`
}

type ProblemConfig struct {
	LoadType           string             `json:"-" toml:"-" yaml:"-" xml:"-"`
	Name               xml.Name           `json:"-" xml:"Problem-Config" yaml:"-" toml:"-"`
	JudgeConfig        JudgeConfig        `json:"judge" yaml:"judge" toml:"judge" xml:"judge"`
	SpecialJudgeConfig SpecialJudgeConfig `json:"special-judge" yaml:"special-judge" toml:"special-judge" xml:"special-judge"`
}

func DefaultProblemConfig() *ProblemConfig {
	return &ProblemConfig{
		LoadType: ".toml",
		Name:     xml.Name{},
		JudgeConfig: JudgeConfig{
			Type: "acm",
			Tasks: []TaskConfig{
				{
					Name:        "default",
					CaseCount:   0,
					TimeLimit:   (time.Second).Millisecond(),
					MemoryLimit: (memory.KiloByte * 64).Byte(),
					Score:       100,
					InputPath:   "/1",
					OutputPath:  "/1",
				},
			},
		},
		SpecialJudgeConfig: SpecialJudgeConfig{
			SpecialJudge: 0,
			LanguageType: types.LanguageGCCCpp11,
			FilePath:     "",
		},
	}
}

func (c *ProblemConfig) Modify(path string, val json.RawMessage) error {
	if path == "" {
		err := json.Unmarshal(val, c)
		if err != nil {
			return err
		}
		return nil
	}
	paths := strings.Split(path, ".")
	switch paths[0] {
	case "judge":
		return c.JudgeConfig.Modifys(paths[1:], val)
	case "special-judge":
		return c.SpecialJudgeConfig.Modifys(paths[1:], val)
	default:
		return errors.New("property missing")
	}
}

func (c *ProblemConfig) Modifys(paths []string, val json.RawMessage) error {
	if len(paths) == 0 {
		err := json.Unmarshal(val, c)
		if err != nil {
			return err
		}
		return nil
	}
	switch paths[0] {
	case "judge-type":
		return c.JudgeConfig.Modifys(paths[1:], val)
	case "special-judge":
		return c.SpecialJudgeConfig.Modifys(paths[1:], val)
	default:
		return errors.New("property missing")
	}
}

func (c *JudgeConfig) Modifys(paths []string, val json.RawMessage) error {
	if len(paths) == 0 {
		err := json.Unmarshal(val, c)
		if err != nil {
			return err
		}
		return nil
	}
	switch paths[0] {
	case "judge-type":
		err := json.Unmarshal(val, &c.Type)
		if err != nil {
			return err
		}
		return nil
	case "tasks":
		if len(paths) < 2 {
			return errors.New("path consumed")
		}
		index, err := strconv.Atoi(paths[1])
		if err != nil {
			return err
		}
		if len(c.Tasks) <= index {
			return errors.New("index overflow")
		}
		return c.Tasks[index].Modifys(paths[2:], val)
	default:
		return errors.New("property missing")
	}
}

func (c *SpecialJudgeConfig) Modifys(paths []string, val json.RawMessage) error {
	if len(paths) == 0 {
		err := json.Unmarshal(val, c)
		if err != nil {
			return err
		}
		return nil
	}
	switch paths[0] {
	case "enable":
		err := json.Unmarshal(val, &c.SpecialJudge)
		if err != nil {
			return err
		}
		return nil
	case "language-type":
		err := json.Unmarshal(val, &c.LanguageType)
		if err != nil {
			return err
		}
		return nil
	case "file-path":
		err := json.Unmarshal(val, &c.FilePath)
		if err != nil {
			return err
		}
		return nil
	default:
		return errors.New("property missing")
	}
}

func (c *TaskConfig) Modifys(paths []string, val json.RawMessage) error {
	if len(paths) == 0 {
		err := json.Unmarshal(val, c)
		if err != nil {
			return err
		}
		return nil
	}
	switch paths[0] {
	case "name":
		err := json.Unmarshal(val, &c.Name)
		if err != nil {
			return err
		}
		return nil
	case "score":
		err := json.Unmarshal(val, &c.Score)
		if err != nil {
			return err
		}
		return nil
	case "time-limit":
		err := json.Unmarshal(val, &c.TimeLimit)
		if err != nil {
			return err
		}
		return nil
	case "memory-limit":
		err := json.Unmarshal(val, &c.MemoryLimit)
		if err != nil {
			return err
		}
		return nil
	case "input-path":
		err := json.Unmarshal(val, &c.InputPath)
		if err != nil {
			return err
		}
		return nil
	case "output-path":
		err := json.Unmarshal(val, &c.OutputPath)
		if err != nil {
			return err
		}
		return nil
	default:
		return errors.New("property missing")
	}
}
