package couchbase

import (
	"fmt"
	"time"

	"github.com/couchbase/gocb/v2"
	"go.uber.org/zap"
)

type Couchbaser interface {
	Collection(name string) *gocb.Collection
}

type Couchbase struct {
	Bucket *gocb.Bucket
}

func NewCouchbase(logger *zap.Logger, connectionStr string, username string, password string, bucketName string) (*Couchbase, error) {
	cluster, err := gocb.Connect("couchbase://"+connectionStr, gocb.ClusterOptions{
		Authenticator: gocb.PasswordAuthenticator{
			Username: username,
			Password: password,
		},
	})
	if err != nil {
		return nil, err
	}

	bucket := cluster.Bucket(bucketName)
	err = bucket.WaitUntilReady(15*time.Second, nil)
	if err != nil {
		return nil, err
	}
	logger.Info(fmt.Sprintf("Successfully connected to %s, with username %s", bucketName, username))
	return &Couchbase{
		Bucket: bucket,
	}, nil
}

func (c *Couchbase) Collection(name string) *gocb.Collection {
	return c.Bucket.Collection(name)
}
