package problem_desc_test

import (
	"bytes"
	"fmt"
	problem_desc2 "github.com/Myriad-Dreamin/boj-v6/abstract/problem-desc"
	problem_desc "github.com/Myriad-Dreamin/boj-v6/app/problem-desc"
	"github.com/Myriad-Dreamin/boj-v6/config"
	"github.com/Myriad-Dreamin/boj-v6/deployment/database"
	"github.com/Myriad-Dreamin/boj-v6/deployment/oss"
	"github.com/Myriad-Dreamin/boj-v6/external"
	"github.com/Myriad-Dreamin/minimum-lib/logger"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zapcore"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

var db *problem_desc.DBImpl

func TestMain(m *testing.M) {
	engine, err := oss.NewLevelDB("./test", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err := engine.Close(); err != nil {
			panic(err)
		}
	}()

	g := database.NewModule()

	mo := make(module.Module)
	l, err := logger.NewZapLogger(logger.NewZapDevelopmentSugarOption(), zapcore.DebugLevel)
	if err != nil {
		panic(err)
	}
	mo.Provide(config.ModulePath.Global.Logger, l)
	cfg := config.Default()
	mo.Provide(config.ModulePath.Global.Configuration, cfg)

	if ok := g.InstallMock(mo); !ok {
		panic(ok)
	}

	sugar.HandlerError0(mo.ProvideImpl(new(*external.OSSEngine), engine))
	db, err = problem_desc.NewDB(mo)
	if err != nil {
		panic(err)
	}

	sugar.HandlerError0(db.Migrate())
	m.Run()
}

const mdContent = `
# Description
this is a problem description
# Sample Input
` + "```" + `
1
` + "```" + `
# Sample Output
`

func TestT(t *testing.T) {
	var obj = new(problem_desc2.ProblemDesc)
	obj.Name = "233"
	i, err := db.Create(obj)
	if i == 0 || err != nil {
		t.Fatal(i, err)
	}

	pd := db.NewProblemDesc(obj.ProblemID, obj.Name, []byte(mdContent))
	assert.NoError(t, db.SaveDesc(pd))
	assert.NoError(t, db.ReleaseDesc(pd))

	pd.Content = nil
	assert.NoError(t, db.LoadDesc(pd))
	assert.True(t, bytes.Equal(pd.Content, []byte(mdContent)))
	assert.NoError(t, db.ReleaseDesc(pd))

	assert.True(t, bytes.Equal(pd.Key, problem_desc.CreateProblemDescKey(pd.ProblemID, pd.Name)))

	pd.Key = nil
	pd.Content = nil
	assert.NoError(t, db.LoadDesc(pd))
	assert.True(t, bytes.Equal(pd.Content, []byte(mdContent)))
	assert.NoError(t, db.ReleaseDesc(pd))

	pd.Key = nil
	pd.Content = nil
	assert.NoError(t, db.SaveDesc(pd))
	assert.NoError(t, db.ReleaseDesc(pd))

	pd.Key = nil
	pd.Content = nil
	assert.NoError(t, db.LoadDesc(pd))
	assert.True(t, bytes.Equal(pd.Content, []byte{}))
	assert.NoError(t, db.ReleaseDesc(pd))
}
