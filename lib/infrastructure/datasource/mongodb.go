package datasource

import (
	"context"
	"fmt"
	"os"
	"time"

	feature "github.com/rubberyconf/language/lib"
	config "github.com/rubberyconf/rubberyconf/lib/core/configuration"
	"github.com/rubberyconf/rubberyconf/lib/core/logs"
	"github.com/rubberyconf/rubberyconf/lib/core/ports/output"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DataSourceMongoDB struct {
	//client             *mongo.Client
	//featuresCollection *mongo.Collection
}

/*var (
	mongodbDataSource *DataSourceMongoDB
	onceMongodb       sync.Once
)*/

const (
	FEATUREKEY string = "FeatureKey"
)

type MongoFeature struct {
	Id         primitive.ObjectID        `json:"id,omitempty" bson:"_id,omitempty"`
	FeatureKey string                    `json:"featureKey" bson:"featureKey"`
	FeatureDef feature.FeatureDefinition `json:"featureDef" bson:"featureDef"`
	CreatedAt  time.Time                 `json:"createdAt" bson:"createdAt"`
	UpdatedAt  time.Time                 `json:"updatedAt" bson:"updatedAt"`
}

func NewDataSourceMongoDB() *DataSourceMongoDB {

	//onceMongodb.Do(func() {
	mongodbDataSource := new(DataSourceMongoDB)
	//})
	return mongodbDataSource
}

func (source *DataSourceMongoDB) timeOut() time.Duration {
	timeout, err := time.ParseDuration(config.GetConfiguration().Database.TimeOut)
	if err != nil {
		timeout = 1 * time.Second
	}
	return timeout
}

func (source *DataSourceMongoDB) GetFeature(ctx context.Context, feature output.FeatureKeyValue) (bool, error) {

	conf := config.GetConfiguration()

	document := MongoFeature{}
	filter := bson.D{primitive.E{Key: FEATUREKEY, Value: feature.Key}}

	ctxMongo, cancel := context.WithTimeout(ctx, source.timeOut())

	client := source.connect(ctxMongo)
	defer source.disconnect(ctxMongo, client)

	database := client.Database(conf.Database.DatabaseName)
	featuresCollection := database.Collection(conf.Database.Collections.Features)

	err := featuresCollection.FindOne(ctxMongo, filter).Decode(&document)
	if err == mongo.ErrNoDocuments {
		feature.Value = nil
		cancel()
		return false, nil
	} else if err != nil {
		logs.GetLogs().WriteMessage(logs.ERROR, fmt.Sprintf("mongodb doesn't asnwered properly when running FindOne feature %s", feature.Key), err)
		cancel()
		return false, err
	}
	cancel()
	feature.Value = &document.FeatureDef
	return true, nil
}

func (source *DataSourceMongoDB) DeleteFeature(ctx context.Context, feature output.FeatureKeyValue) bool {

	conf := config.GetConfiguration()
	ctxMongo, cancel := context.WithTimeout(ctx, source.timeOut())
	client := source.connect(ctxMongo)
	defer source.disconnect(ctxMongo, client)

	database := client.Database(conf.Database.DatabaseName)
	featuresCollection := database.Collection(conf.Database.Collections.Features)

	_, err := featuresCollection.DeleteMany(ctxMongo, bson.D{primitive.E{Key: FEATUREKEY, Value: feature.Key}})
	cancel()
	if err != nil {
		logs.GetLogs().WriteMessage(logs.ERROR, fmt.Sprintf("mongodb didn't delete feature %s", feature.Key), err)
		return false
	}
	return true
}

func (source *DataSourceMongoDB) CreateFeature(ctx context.Context, feature output.FeatureKeyValue) bool {

	conf := config.GetConfiguration()
	newDoc := MongoFeature{
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		FeatureKey: feature.Key,
		FeatureDef: *feature.Value,
	}
	ctxMongo, cancel := context.WithTimeout(ctx, source.timeOut())
	client := source.connect(ctxMongo)
	defer source.disconnect(ctxMongo, client)

	database := client.Database(conf.Database.DatabaseName)
	featuresCollection := database.Collection(conf.Database.Collections.Features)

	_, err := featuresCollection.InsertOne(ctxMongo, newDoc)
	cancel()
	if err != nil {
		logs.GetLogs().WriteMessage(logs.ERROR, fmt.Sprintf("mongodb didn't create a new feature document for feature %s", feature.Key), err)
		return false
	}
	return true
}

func (source *DataSourceMongoDB) EnableFeature(keys map[string]string) (output.FeatureKeyValue, bool) {
	return enableFeature(keys)
}

func (source *DataSourceMongoDB) ReviewDependencies(conf *config.Config) {
	if conf.Api.Source == MONGODB &&
		conf.Database.Url == "" {
		logs.GetLogs().WriteMessage(logs.ERROR, "database mongodb server dependency enabled but not configured, check config yml file.", nil)
		os.Exit(2)
	}
}

func (source *DataSourceMongoDB) connect(ctx context.Context) *mongo.Client {

	conf := config.GetConfiguration()
	clientOptions := options.Client().ApplyURI(conf.Database.Url)
	client, err := mongo.NewClient(clientOptions)
	ctxMongo, cancel := context.WithTimeout(ctx, source.timeOut())
	if err != nil {
		logs.GetLogs().WriteMessage(logs.ERROR, "unable to cerate monogo client", err)
		os.Exit(2)
	}
	//ctx, _ := context.WithTimeout(ctx, 10*time.Second)
	err = client.Connect(ctxMongo)
	if err != nil {
		logs.GetLogs().WriteMessage(logs.ERROR, "unable to connect monogo client", err)
		os.Exit(2)
	}
	err = client.Ping(ctxMongo, nil)
	if err != nil {
		logs.GetLogs().WriteMessage(logs.ERROR, "mongodb doesn't answer ping", err)
		os.Exit(2)
	}
	//defer client.Disconnect(ctx)
	cancel()

	//source.client = client

	//database := client.Database(conf.Database.DatabaseName)
	//source.featuresCollection = database.Collection(conf.Database.Collections.Features)
	return client

}

func (source *DataSourceMongoDB) disconnect(ctx context.Context, client *mongo.Client) {
	ctxMongo, cancel := context.WithTimeout(ctx, source.timeOut())
	client.Disconnect(ctxMongo)
	cancel()
}
