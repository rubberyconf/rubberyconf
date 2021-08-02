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
	ctx               context.Context
)

type MongoFeature struct {
	Id         primitive.ObjectID        `json:"id,omitempty" bson:"_id,omitempty"`
	FeatureKey string                    `json:"featureKey" bson:"featureKey"`
	FeatureDef feature.FeatureDefinition `json:"featureDef" bson:"featureDef"`
	CreatedAt  time.Time                 `json:"createdAt" bson:"createdAt"`
	UpdatedAt  time.Time                 `json:"updatedAt" bson:"updatedAt"`
}

func NewDataSourceMongoDB() *DataSourceMongoDB {

	onceMongodb.Do(func() {
		mongodbDataSource = new(DataSourceMongoDB)
		mongodbDataSource.connect()
	})
	return mongodbDataSource
}
func (source *DataSourceMongoDB) GetFeature(feature *Feature) (bool, error) {

	document := MongoFeature{}
	filter := bson.D{{"FeatureKey", feature.Key}}
	err := source.featuresCollection.FindOne(ctx, filter).Decode(&document)
	if err == mongo.ErrNoDocuments {
		feature.Value = nil
		return false, nil
	} else if err != nil {
		logs.GetLogs().WriteMessage("error", fmt.Sprintf("mongodb doesn't asnwered properly when running FindOne feature %s", feature.Key), err)
		return false, err
	}
	feature.Value = &document.FeatureDef
	return true, nil
}

func (source *DataSourceMongoDB) DeleteFeature(feature Feature) bool {
	_, err := source.featuresCollection.DeleteMany(ctx, bson.D{{"FeatureKey", feature.Key}})
	if err != nil {
		logs.GetLogs().WriteMessage("error", fmt.Sprintf("mongodb didn't delete feature %s", feature.Key), err)
		return false
	}
	return true
}

func (source *DataSourceMongoDB) CreateFeature(feature Feature) bool {

	newDoc := MongoFeature{
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		FeatureKey: feature.Key,
		FeatureDef: *feature.Value,
	}
	_, err := source.featuresCollection.InsertOne(ctx, newDoc)
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

func (source *DataSourceMongoDB) connect() {

	conf := config.GetConfiguration()
	clientOptions := options.Client().ApplyURI(conf.Database.Url)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		logs.GetLogs().WriteMessage("error", "unable to cerate monogo client", err)
		os.Exit(2)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		logs.GetLogs().WriteMessage("error", "unable to connect monogo client", err)
		os.Exit(2)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		logs.GetLogs().WriteMessage("error", "mongodb doesn't answer ping", err)
		os.Exit(2)
	}
	//defer client.Disconnect(ctx)

	source.client = client
	database := client.Database(conf.Database.DatabaseName)
	source.featuresCollection = database.Collection(conf.Database.Collections.Features)

}
