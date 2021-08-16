package metrics

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/rubberyconf/rubberyconf/internal/config"
	"github.com/rubberyconf/rubberyconf/internal/logs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Metrics struct {
	client            *mongo.Client
	metricsCollection *mongo.Collection
}

type MongoMetrics struct {
	Id        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Feature   string             `json:"feature" bson:"feature"`
	CreatedAt time.Time          `json:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt"`
	Counter   int64              `json:"counter"`
}

var (
	metrics   *Metrics
	mongoOnce sync.Once
	ctx       context.Context
)

func GetMetrics() *Metrics {
	mongoOnce.Do(func() {
		metrics = new(Metrics)
		//metrics.connect()
	})
	return metrics
}

func (metric *Metrics) fetchMetrics(feature string) (*MongoMetrics, error) {

	newdocument := false
	var metricRegister MongoMetrics
	filter := bson.D{primitive.E{Key: "feature", Value: feature}}
	err := metric.metricsCollection.FindOne(ctx, filter).Decode(&metricRegister)
	if err == mongo.ErrNoDocuments {
		newdocument = true
	} else if err != nil {
		logs.GetLogs().WriteMessage("error", fmt.Sprintf("mongodb doesn't asnwered properly when running FindOne feature %s", feature), err)
	}

	if newdocument {

		newDoc := MongoMetrics{
			Feature:   feature,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Counter:   0,
		}
		_, err := metric.metricsCollection.InsertOne(ctx, newDoc)
		if err != nil {
			logs.GetLogs().WriteMessage("error", fmt.Sprintf("mongodb didn't create a new metric document for feature: %s", feature), err)
			return nil, err
		}
		return &newDoc, nil
	} else {
		return &metricRegister, nil
	}
}
func (metric *Metrics) storeMetrics(metricRegister *MongoMetrics) (bool, error) {

	_, err := metric.metricsCollection.UpdateOne(ctx,
		bson.D{primitive.E{Key: "feature", Value: metricRegister.Feature}},
		bson.D{primitive.E{Key: "$set", Value: bson.D{
			primitive.E{Key: "counter", Value: metricRegister.Counter},
			primitive.E{Key: "UpdatedAt", Value: time.Now()}}}})

	if err == mongo.ErrNoDocuments {
		log.Fatal("It should be create earlier")
		return false, err
	} else if err != nil {
		log.Fatal(err)
		return false, err
	}
	return true, nil
}

func (metricRegister *MongoMetrics) Update() {

	metricRegister.Counter += 1
	metricRegister.UpdatedAt = time.Now()
}

func (metric *Metrics) Update(feature string) (*MongoMetrics, error) {

	metric.connect()
	defer metric.disconnect()
	metricRegister, err := metric.fetchMetrics(feature)
	if err != nil {
		return nil, err
	}

	metricRegister.Update()

	_, err = metric.storeMetrics(metricRegister)
	if err != nil {
		return nil, err
	}

	return metricRegister, nil
}

func (metric *Metrics) Remove(feature string) (bool, error) {

	metric.connect()
	defer metric.disconnect()
	_, err := metric.metricsCollection.DeleteMany(ctx, bson.D{primitive.E{Key: "feature", Value: feature}})
	if err != nil {
		return false, err
	}
	return true, nil
}

func (metric *Metrics) connect() {

	conf := config.GetConfiguration()
	clientOptions := options.Client().ApplyURI(conf.Database.Url)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		logs.GetLogs().WriteMessage("error", "unable to cerate monogo client", err)
		os.Exit(2)
	}
	metrics.client = client
	err = client.Connect(context.Background())
	if err != nil {
		logs.GetLogs().WriteMessage("error", "unable to connect monogo client", err)
		os.Exit(2)
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		logs.GetLogs().WriteMessage("error", "mongodb doesn't answer ping", err)
		os.Exit(2)
	}
	//defer client.Disconnect(ctx)

	database := client.Database(conf.Database.DatabaseName)
	metricsCollection := database.Collection(conf.Database.Collections.Metrics)

	metrics.metricsCollection = metricsCollection

}

func (metric *Metrics) disconnect() {
	metric.client.Disconnect(context.Background())
}
