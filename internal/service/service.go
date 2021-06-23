package service

import (
	"context"
	"encoding/base64"
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/ecode"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/wire"
	pb "github.com/little-bit-shy/go-xgz/api"
	"github.com/little-bit-shy/go-xgz/internal/cli"
	"github.com/little-bit-shy/go-xgz/internal/dao"
	"github.com/little-bit-shy/go-xgz/internal/data/db"
	"github.com/little-bit-shy/go-xgz/internal/data/es"
	"github.com/little-bit-shy/go-xgz/internal/data/hbase"
	"github.com/little-bit-shy/go-xgz/internal/data/jrpc"
	"github.com/little-bit-shy/go-xgz/internal/data/redis"
	"github.com/little-bit-shy/go-xgz/pkg/api"
	"github.com/little-bit-shy/go-xgz/pkg/config"
	db2 "github.com/little-bit-shy/go-xgz/pkg/dao/db"
	"github.com/little-bit-shy/go-xgz/pkg/dao/jsonRpc"
	"github.com/little-bit-shy/go-xgz/pkg/database/sql"
	"github.com/little-bit-shy/go-xgz/pkg/elastic"
	"github.com/mitchellh/mapstructure"
	"github.com/tsuna/gohbase/hrpc"
)

var Provider = wire.NewSet(New, wire.Bind(new(pb.AppServer), new(*Service)))

type cfg struct {
}

// Service service.
type Service struct {
	ac  *paladin.Map
	dao *dao.Dao
	cfg *cfg
}

// New new a service and return.
func New(c *cli.Cli) (s *Service, cf func(), err error) {
	s = &Service{
		ac:  &paladin.Map{},
		dao: c.Dao,
		cfg: &cfg{},
	}
	cf = s.Close
	if err = paladin.Get("application.toml").Unmarshal(s.ac); err != nil {
		return
	}
	if err = config.Env(s.ac, s.cfg, "App"); err != nil {
		return
	}
	return
}

// CallMyself grpc func.
func (s *Service) CallSelf(ctx context.Context, req *pb.HelloReq) (reply *pb.HelloResp, err error) {
	if req.Name == "233" {
		err = ecode.AccessDenied
		return
	}
	reply = &pb.HelloResp{
		Content: "who is " + req.Name,
	}
	return
}

// Say bm func.
func (s *Service) Say(ctx context.Context, req *pb.HelloReq) (reply *pb.SayResp, err error) {
	var HelloResp *api.HelloResp
	// TODO: the actual business needs to be open
	//if HelloResp, err = client.CallSelf(s.dao, ctx, req.Name); err != nil {
	//	err = ecode.Int(-10000)
	//	return
	//}
	callSelfResp := pb.SayResp_CallSelfMsg{}
	if err = mapstructure.Decode(HelloResp, &callSelfResp); err != nil {
		return
	}

	var tx *sql.Tx
	if tx, err = s.dao.Db.Connect.Begin(ctx); err != nil {
		err = ecode.Int(-1000)
		return
	}
	var book *db.Book
	var metadata *db2.Metatada
	if book, metadata, err = db.GetBook(s.dao, tx, ctx); err != nil {
		err = ecode.Int(-1000)
		return
	}
	if !metadata.Exist {
		if _, err = db.CreateBook(s.dao, tx, ctx, req.Name); err != nil {
			err = ecode.Int(-1000)
			return
		}
	} else {
		if _, err = db.UpdateBook(s.dao, tx, ctx, req.Name); err != nil {
			err = ecode.Int(-1000)
			return
		}
	}
	if book, _, err = db.GetBook(s.dao, tx, ctx); err != nil {
		err = ecode.Int(-1000)
		return
	}
	if err = tx.Commit(); err != nil {
		err = ecode.Int(-1000)
		return
	}
	db := pb.SayResp_DbMsg{}
	if err = mapstructure.Decode(book, &db); err != nil {
		return
	}

	// do hbase
	if _, err = hbase.PutSomething(s.dao, ctx); err != nil {
		err = ecode.Int(-1001)
		return
	}
	var resultGet *hrpc.Result
	if resultGet, err = hbase.GetSomething(s.dao, ctx); err != nil {
		err = ecode.Int(-1001)
		return
	}
	hbaseResults := [][]byte{}
	for _, cell := range resultGet.Cells {
		var value []byte
		if value, err = base64.StdEncoding.DecodeString(string(cell.Value)); err != nil {
			err = ecode.Int(-1001)
			return
		}
		hbaseResults = append(hbaseResults, value)
	}
	// do redis
	if _, err = redis.SetTest(s.dao, ctx); err != nil {
		err = ecode.Int(-1002)
		return
	}
	var redisResult interface{}
	if redisResult, err = redis.GetTest(s.dao, ctx); err != nil {
		err = ecode.Int(-1002)
		return
	}
	redisGet := string(redisResult.([]uint8))

	// do es
	var health *elastic.ClusterHealthResponse
	if health, err = es.GetHealth(s.dao, ctx); err != nil {
		err = ecode.Int(-1003)
		return
	}
	es := pb.SayResp_EsMsg{}
	if err = mapstructure.Decode(health, &es); err != nil {
		return
	}

	// do jrpc
	var result jsonRpc.Result
	if result, err = jrpc.Dict(s.dao, ctx); err != nil {
		err = ecode.Int(-1004)
		return
	}
	jrpc := pb.SayResp_JrpcMsg{}
	if err = mapstructure.Decode(result, &jrpc); err != nil {
		return
	}

	reply = &pb.SayResp{
		Db:    &db,
		Es:    &es,
		Redis: redisGet,
		Jrpc:  &jrpc,
		//Client: &callSelfResp,
		Hbase: hbaseResults,
	}
	return
}

// Ping ping the resource.
func (s *Service) Ping(ctx context.Context, e *empty.Empty) (*empty.Empty, error) {
	return &empty.Empty{}, s.dao.Ping(ctx)
}

// Close close the resource.
func (s *Service) Close() {
}
