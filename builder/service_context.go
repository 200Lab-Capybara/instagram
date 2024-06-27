package builder

import (
	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
	"instagram/common"
	"instagram/components/hasher"
	"instagram/components/logger"
)

type serviceContext struct {
	db         common.SQLDatabase
	hasher     hasher.Hasher
	natsCon    *nats.Conn
	logger     logger.Logger
	v1         *gin.RouterGroup
	v1Internal *gin.RouterGroup
}

func NewServiceContext(db common.SQLDatabase, hasher hasher.Hasher, natsCon *nats.Conn, logger logger.Logger, v1 *gin.RouterGroup, v1Internal *gin.RouterGroup) ServiceContext {
	return &serviceContext{
		db:         db,
		hasher:     hasher,
		natsCon:    natsCon,
		logger:     logger,
		v1:         v1,
		v1Internal: v1Internal,
	}
}

func (s *serviceContext) GetDB() common.SQLDatabase {
	return s.db
}

func (s *serviceContext) GetHasher() hasher.Hasher {
	return s.hasher
}

func (s *serviceContext) GetNatsConn() *nats.Conn {
	return s.natsCon
}

func (s *serviceContext) GetLogger() logger.Logger {
	return s.logger
}

func (s *serviceContext) GetV1() *gin.RouterGroup {
	return s.v1
}

func (s *serviceContext) GetV1Internal() *gin.RouterGroup {
	return s.v1Internal
}

type ServiceContext interface {
	GetDB() common.SQLDatabase
	GetHasher() hasher.Hasher
	GetNatsConn() *nats.Conn
	GetLogger() logger.Logger
	GetV1() *gin.RouterGroup
	GetV1Internal() *gin.RouterGroup
}
