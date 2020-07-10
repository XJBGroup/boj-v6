package sphinxcore

import (
	"context"
	"encoding/binary"
	"encoding/xml"
	"fmt"
	"github.com/Myriad-Dreamin/boj-v6/model"
	"github.com/Myriad-Dreamin/boj-v6/plugin"
	"github.com/Myriad-Dreamin/boj-v6/server/router"
	submissionservice "github.com/Myriad-Dreamin/boj-v6/service/submission"
	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type Plugin struct {
	logger              plugin.Logger
	cfg                 *plugin.ServerConfig
	consumer            *cluster.Consumer
	submissionService   *submissionservice.Service
	submissionDB        *model.SubmissionDB
	contestSubmissionDB *model.ContestSubmissionDB
	contestDB           *model.ContestDB
	userDB              *model.UserDB

	clusterHosts []string
}

func New() *Plugin {
	return new(Plugin)
}

const (
	contestOffset = 1 << 60
)

func (plg *Plugin) Submit(code string, submission *model.Submission, problem *model.Problem) {
	config := sarama.NewConfig()
	// 等待服务器所有副本都保存成功后的响应
	config.Producer.RequiredAcks = sarama.WaitForAll
	// 随机的分区类型：返回一个分区器，该分区器每次选择一个随机分区
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	// 是否等待成功和失败后的响应
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true

	config.Version = sarama.V2_0_0_0

	// 使用给定代理地址和配置创建一个同步生产者
	producer, err := sarama.NewSyncProducer(plg.clusterHosts, config)
	if err != nil {
		panic(err)
	}

	defer producer.Close()

	//构建发送的消息，
	msg := &sarama.ProducerMessage{
		//Topic: "test",//包含了消息的主题
		Partition: int32(10),                   //
		Key:       sarama.StringEncoder("key"), //
	}

	msg.Topic = "in"
	var e = make([]byte, 8)
	binary.BigEndian.PutUint64(e, uint64(submission.ID))
	msg.Headers = []sarama.RecordHeader{
		{[]byte("problem"), []byte(strconv.Itoa(int(problem.ID)))},
		{[]byte("lang"), []byte{submission.Language}},
		{[]byte("uid"), e},
	}

	fmt.Println(strconv.Itoa(int(problem.ID)), []byte{submission.Language}, e, submission.ID)

	msg.Value = sarama.ByteEncoder([]byte(code))
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		plg.logger.Debug("post submission failed", "error", err)
	} else {
		plg.logger.Debug("submission post success", "partition", partition, "offset", offset)
	}
}

func (plg *Plugin) ContestSubmit(code string, submission *model.ContestSubmission, problem *model.Problem) {
	config := sarama.NewConfig()
	// 等待服务器所有副本都保存成功后的响应
	config.Producer.RequiredAcks = sarama.WaitForAll
	// 随机的分区类型：返回一个分区器，该分区器每次选择一个随机分区
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	// 是否等待成功和失败后的响应
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true

	config.Version = sarama.V2_0_0_0

	// 使用给定代理地址和配置创建一个同步生产者
	producer, err := sarama.NewSyncProducer(plg.clusterHosts, config)
	if err != nil {
		panic(err)
	}

	defer producer.Close()

	//构建发送的消息，
	msg := &sarama.ProducerMessage{
		//Topic: "test",//包含了消息的主题
		Partition: int32(10),                   //
		Key:       sarama.StringEncoder("key"), //
	}

	msg.Topic = "in"
	var e = make([]byte, 8)
	binary.BigEndian.PutUint64(e, uint64(submission.ID))
	msg.Headers = []sarama.RecordHeader{
		{[]byte("problem"), []byte(strconv.Itoa(int(problem.ID)))},
		{[]byte("lang"), []byte{submission.Language}},
		{[]byte("uid"), e},
	}

	fmt.Println(strconv.Itoa(int(problem.ID)), []byte{submission.Language}, e, submission.ID)

	msg.Value = sarama.ByteEncoder([]byte(code))
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		plg.logger.Debug("post submission failed", "error", err)
	} else {
		plg.logger.Debug("submission post success", "partition", partition, "offset", offset)
	}
}

func (plg *Plugin) SubmissionPostMiddleware(c *gin.Context) {
	c.Next()

	if c.IsAborted() {
		return
	}

	submission := c.MustGet("s").(*model.Submission)
	problem := c.MustGet("p").(*model.Problem)
	code := c.MustGet("c").(string)

	plg.Submit(code, submission, problem)
}

func (plg *Plugin) ContestSubmissionPostMiddleware(c *gin.Context) {
	c.Next()

	if c.IsAborted() {
		return
	}

	submission := c.MustGet("s").(*model.ContestSubmission)
	problem := c.MustGet("p").(*model.Problem)
	code := c.MustGet("c").(string)

	plg.ContestSubmit(code, submission, problem)
}

type ExtSphinxCoreConfig struct {
	ClusterHosts []string `json:"cluster-host" yaml:"cluster-host" toml:"cluster-host" xml:"cluster-host"`
}

type Config struct {
	LoadType            string              `json:"-" yaml:"-" toml:"-" xml:"-"`
	Name                xml.Name            `json:"-" yaml:"-" toml:"-" xml:"server-config"`
	ExtSphinxCoreConfig ExtSphinxCoreConfig `json:"sphinx-core" yaml:"sphinx-core" toml:"sphinx-core" xml:"sphinx-core"`
}

func (plg *Plugin) Configuration(logger plugin.Logger, loader plugin.ConfigLoader, cfg *plugin.ServerConfig) plugin.Plugin {
	plg.logger = logger
	plg.cfg = cfg

	var scc Config
	if !loader(&scc) {
		return nil
	}

	plg.clusterHosts = scc.ExtSphinxCoreConfig.ClusterHosts
	plg.logger.Debug("sphinx core configuration", "cluster host", plg.clusterHosts)
	return plg
}

func (plg *Plugin) Inject(services *plugin.ServiceProvider, dbs *plugin.DatabaseProvider, module plugin.Module) plugin.Plugin {
	plg.submissionService = services.SubmissionService()
	_ = dbs

	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = true
	config.Consumer.Offsets.CommitInterval = 1 * time.Second
	config.Consumer.Offsets.Initial = sarama.OffsetNewest
	config.Version = sarama.V2_0_0_0

	var err error
	plg.consumer, err = cluster.NewConsumer(plg.clusterHosts, "Q2", []string{"result"}, config)

	if err != nil {
		plg.logger.Info("build consumer failed", "error", err)
		return nil
	}

	r := module.Require("/router/root").(*router.RootRouter)
	plg.submissionDB = dbs.SubmissionDB()
	plg.contestSubmissionDB = dbs.ContestSubmissionDB()
	plg.contestDB = dbs.ContestDB()

	r.Problem.Submission.SubmissionOnProblem.Post.Use(plg.SubmissionPostMiddleware)
	r.Contest.IDRouter.ContestSubmission.ContestSubmissionOnProblem.Post.Use(plg.ContestSubmissionPostMiddleware)

	return plg
}

type JudgeInfo interface {
}

func strintify(rcs []*sarama.RecordHeader) interface{} {
	var kvs = make([]struct {
		Key   string
		Value string
		Raw   *sarama.RecordHeader
	}, len(rcs))
	for i, rc := range rcs {
		kvs[i].Key = string(rc.Key)
		kvs[i].Value = string(rc.Value)
		kvs[i].Raw = rc
	}
	return kvs
}

func (plg *Plugin) Work(ctx context.Context) {

	go func() {
		for err := range plg.consumer.Errors() {
			plg.logger.Error("error occurs", "error", err)
		}
	}()

	go func() {
		for notify := range plg.consumer.Notifications() {
			plg.logger.Error("notification: ", "notify", notify)
		}
	}()

	for {
		fmt.Println("im here")
		select {
		case msg := <-plg.consumer.Messages():
			if len(msg.Headers) != 6 {
				plg.logger.Info("bad msg received", "message reply", "topic", msg.Topic,
					"timestamp", msg.Timestamp, "partition", msg.Partition, "offset", msg.Offset,
					"header", strintify(msg.Headers), "key", string(msg.Key),
					"value", string(msg.Value))
				plg.consumer.MarkOffset(msg, "")
				continue
			}

			mem, err := strconv.ParseInt(string(msg.Headers[0].Value), 10, 64)
			if err != nil {
				plg.logger.Info("cannot unmarshal mem", "data", string(msg.Headers[0].Value), "error", err)
			}
			timeX, err := strconv.ParseInt(string(msg.Headers[1].Value), 10, 64)
			if err != nil {
				plg.logger.Info("cannot unmarshal time", "data", string(msg.Headers[1].Value), "error", err)
			}
			sid, err := strconv.ParseUint(string(msg.Headers[2].Value), 10, 64)
			if err != nil {
				plg.logger.Info("cannot unmarshal sid", "data", string(msg.Headers[2].Value), "error", err)
			}
			lastTestCase, err := strconv.ParseInt(string(msg.Headers[3].Value), 10, 64)
			if err != nil {
				plg.logger.Info("cannot unmarshal last test case", "data", string(msg.Headers[3].Value), "error", err)
			}
			score, err := strconv.ParseInt(string(msg.Headers[4].Value), 10, 64)
			if err != nil {
				plg.logger.Info("cannot unmarshal score", "data", string(msg.Headers[4].Value), "error", err)
			}

			if sid >= contestOffset {
				sid -= contestOffset
				submission, err := plg.contestSubmissionDB.ID(uint(sid))
				if err != nil || submission == nil {
					plg.logger.Error("fetch contest submission error", "error", err)
					plg.consumer.MarkOffset(msg, "")
					continue
				}
				submission.RunMemory = mem
				submission.RunTime = timeX
				submission.LastTestCase = lastTestCase
				submission.Score = score
				submission.Information = string(msg.Headers[5].Value)

				if aff, err := submission.Update(); err != nil || aff == 0 {
					plg.logger.Error("update submission error", "error", err)
					continue
				}
				plg.consumer.MarkOffset(msg, "")

				if err := plg.contestDB.Rate(submission); err != nil {
					plg.logger.Error("rate submission error", "error", err)
					continue
				}
			} else {
				submission, err := plg.submissionDB.ID(uint(sid))
				if err != nil || submission == nil {
					plg.logger.Error("fetch submission error", "error", err)
					plg.consumer.MarkOffset(msg, "")
					continue
				}
				submission.RunMemory = mem
				submission.RunTime = timeX
				submission.LastTestCase = lastTestCase
				submission.Score = score
				submission.Information = string(msg.Headers[5].Value)

				if aff, err := submission.Update(); err != nil || aff == 0 {
					plg.logger.Error("update submission error", "error", err)
					continue
				}
				plg.consumer.MarkOffset(msg, "")
			}
		case <-ctx.Done():
			fmt.Println("gua le")
			err := plg.consumer.Close()
			if err != nil {
				plg.logger.Error("plugin Sphinx-Core close error", "error", err)
			}
			return
		}
	}
}
