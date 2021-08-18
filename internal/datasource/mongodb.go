package datasource

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/rubberyconf/rubberyconf/internal/config"
	"github.com/rubberyconf/rubberyconf/internal/feature"
	"github.com/rubberyconf/rubberyconf/internal/logs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DataSourceMongoDB struct {
	client             *mongo.Client
	featuresCollection *mongo.Collection
}

var (
	mongodbDataSource *DataSourceMongoDB
	onceMongodb       sync.Once
)

type MongoFeature struct {
	Id         primitive.ObjectID        `json:"id,omitempty" bson:"_id,omitempty"`
	FeatureKey string                    `json:"featureKey" bson:"featureKey"`
	FeatureDef feature.FeatureDefinition `json:"featureDef" bson:"featureDef"`
	CreatedAt  time.Time                 `json:"createdAt" bson:"createdAt"`
	UpdatedAt  time.Time                 `json:"updatedAt" bson:"updatedAt"`
}

func NewDataSourceMongoDB(ctx context.Context) *DataSourceMongoDB {

	onceMongodb.Do(func() {
		mongodbDataSource = new(DataSourceMongoDB)
		mongodbDataSource.connect(ctx)
	})
	return mongodbDataSource
}

func (source *DataSourceMongoDB) timeOut() time.Duration {
	timeout, err := time.ParseDuration(config.GetConfiguration().Database.TimeOut)
	if err != nil {
		timeout = 1 * time.Second
	}
	return timeout
}

func (source *DataSourceMongoDB) GetFeature(ctx context.Context, feature *Feature) (bool, error) {

	document := MongoFeature{}
	filter := bson.D{primitive.E{Key: "FeatureKey", Value: feature.Key}}

	ctxMongo, cancel := context.WithTimeout(ctx, source.timeOut())
	err := source.featuresCollection.FindOne(ctxMongo, filter).Decode(&document)
	if err == mongo.ErrNoDocuments {
		feature.Value = nil
		cancel()
		return false, nil
	} else if err != nil {
		logs.GetLogs().WriteMessage("error", fmt.Sprintf("mongodb doesn't asnwered properly when running FindOne feature %s", feature.Key), err)
		cancel()
		return false, err
	}
	cancel()
	feature.Value = &document.FeatureDef
	return true, nil
}

func (source *DataSourceMongoDB) DeleteFeature(ctx context.Context, feature Feature) bool {

	ctxMongo, cancel := context.WithTimeout(ctx, source.timeOut())
	_, err := source.featuresCollection.DeleteMany(ctxMongo, bson.D{primitive.E{Key: "FeatureKey", Value: feature.Key}})
	cancel()
	if err != nil {
		logs.GetLogs().WriteMessage("error", fmt.Sprintf("mongodb didn't delete feature %s", feature.Key), err)
		return false
	}
	return true
}

func (source *DataSourceMongoDB) CreateFeature(ctx context.Context, feature Feature) bool {

	newDoc := MongoFeature{
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		FeatureKey: feature.Key,
		FeatureDef: *feature.Value,
	}
	ctxMongo, cancel := context.WithTimeout(ctx, source.timeOut())
	_, err := source.featuresCollection.InsertOne(ctxMongo, newDoc)
	cancel()
	if err != nil {
		logs.GetLogs().WriteMessage("error", fmt.Sprintf("mongodb didn't create a new feature document for feature %s", feature.Key), err)
		return false
	}
	return true
}

func (source *DataSourceMongoDB) EnableFeature(keys map[string]string) (Feature, bool) {
	return enableFeature(keys)
}

func (source *DataSourceMongoDB) reviewDependencies(conf *config.Config) {
	if conf.Api.Source == MONGODB &&
		conf.Database.Url == "" {
		logs.GetLogs().WriteMessage("error", "database mongodb server dependency enabled but not configured, check config yml file.", nil)
		os.Exit(2)
	}
}

func (source *DataSourceMongoDB) connect(ctx context.Context) {

	conf := config.GetConfiguration()
	clientOptions := options.Client().ApplyURI(conf.Database.Url)
	client, err := mongo.NewClient(clientOptions)
	ctxMongo, cancel := context.WithTimeout(ctx, source.timeOut())
	if err != nil {
		logs.GetLogs().WriteMessage("error", "unable to cerate monogo client", err)
		os.Exit(2)
	}
	//ctx, _ := context.WithTimeout(ctx, 10*time.Second)
	err = client.Connect(ctxMongo)
	if err != nil {
		logs.GetLogs().WriteMessage("error", "unable to connect monogo client", err)
		os.Exit(2)
	}
	err = client.Ping(ctxMongo, nil)
	if err != nil {
		logs.GetLogs().WriteMessage("error", "mongodb doesn't answer ping", err)
		os.Exit(2)
	}
	//defer client.Disconnect(ctx)
	cancel()

	source.client = client
	database := client.Database(conf.Database.DatabaseName)
	source.featuresCollection = database.Collection(conf.Database.Collections.Features)

}
