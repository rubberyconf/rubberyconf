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
)

func GetMetrics() *Metrics {
	mongoOnce.Do(func() {
		metrics = new(Metrics)
	})
	return metrics
}

func (metric *Metrics) timeOut() time.Duration {
	timeout, err := time.ParseDuration(config.GetConfiguration().Database.TimeOut)
	if err != nil {
		timeout = 1 * time.Second
	}
	return timeout
}

func (metric *Metrics) fetchMetrics(ctx context.Context, client *mongo.Client, feature string) (*MongoMetrics, error) {

	newdocument := false
	var metricRegister MongoMetrics
	filter := bson.D{primitive.E{Key: "feature", Value: feature}}

	ctxMongo, cancel := context.WithTimeout(ctx, metric.timeOut())

	metricsCollection := metric.getMetricsCollection(client)
	err := metricsCollection.FindOne(ctxMongo, filter).Decode(&metricRegister)
	if err == mongo.ErrNoDocuments {
		newdocument = true
	} else if err != nil {
		logs.GetLogs().WriteMessage("error", fmt.Sprintf("mongodb doesn't asnwered properly when running FindOne feature %s", feature), err)
	}
	cancel()

	if newdocument {

		newDoc := MongoMetrics{
			Feature:   feature,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Counter:   0,
		}
		ctxMongo, cancel = context.WithTimeout(ctx, metric.timeOut())
		_, err := metricsCollection.InsertOne(ctxMongo, newDoc)
		if err != nil {
			logs.GetLogs().WriteMessage("error", fmt.Sprintf("mongodb didn't create a new metric document for feature: %s", feature), err)
			cancel()
			return nil, err
		}
		cancel()
		return &newDoc, nil
	} else {
		return &metricRegister, nil
	}
}
func (metric *Metrics) storeMetrics(ctx context.Context, client *mongo.Client, metricRegister *MongoMetrics) (bool, error) {

	ctxMongo, cancel := context.WithTimeout(ctx, metric.timeOut())
	metricsCollection := metric.getMetricsCollection(client)
	_, err := metricsCollection.UpdateOne(ctxMongo,
		bson.D{primitive.E{Key: "feature", Value: metricRegister.Feature}},
		bson.D{primitive.E{Key: "$set", Value: bson.D{
			primitive.E{Key: "counter", Value: metricRegister.Counter},
			primitive.E{Key: "UpdatedAt", Value: time.Now()}}}})
	cancel()
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

func (metric *Metrics) Update(ctx context.Context, feature string) (*MongoMetrics, error) {

	client := metric.connect(ctx)
	defer metric.disconnect(ctx, client)
	metricRegister, err := metric.fetchMetrics(ctx, client, feature)
	if err != nil {
		return nil, err
	}

	metricRegister.Update()

	_, err = metric.storeMetrics(ctx, client, metricRegister)
	if err != nil {
		return nil, err
	}

	return metricRegister, nil
}

func (metric *Metrics) getMetricsCollection(client *mongo.Client) *mongo.Collection {

	conf := config.GetConfiguration()
	database := client.Database(conf.Database.DatabaseName)
	metricsCollection := database.Collection(conf.Database.Collections.Metrics)
	return metricsCollection
}

func (metric *Metrics) Remove(ctx context.Context, feature string) (bool, error) {

	client := metric.connect(ctx)
	defer metric.disconnect(ctx, client)
	ctxMongo, cancel := context.WithTimeout(ctx, metric.timeOut())
	metricsCollection := metric.getMetricsCollection(client)
	_, err := metricsCollection.DeleteMany(ctxMongo, bson.D{primitive.E{Key: "feature", Value: feature}})
	cancel()
	if err != nil {
		return false, err
	}
	return true, nil
}

func (metric *Metrics) connect(ctx context.Context) *mongo.Client {

	conf := config.GetConfiguration()
	clientOptions := options.Client().ApplyURI(conf.Database.Url)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		logs.GetLogs().WriteMessage("error", "unable to cerate monogo client", err)
		os.Exit(2)
	}

	ctxMongo, cancel := context.WithTimeout(ctx, metric.timeOut())
	err = client.Connect(ctxMongo)
	if err != nil {
		logs.GetLogs().WriteMessage("error", "unable to connect monogo client", err)
		os.Exit(2)
	}
	cancel()
	err = client.Ping(ctxMongo, nil)
	if err != nil {
		logs.GetLogs().WriteMessage("error", "mongodb doesn't answer ping", err)
		os.Exit(2)
	}
	cancel()
	return client
}

func (metric *Metrics) disconnect(ctx context.Context, client *mongo.Client) {
	ctxMongo, cancel := context.WithTimeout(ctx, metric.timeOut())
	client.Disconnect(ctxMongo)
	cancel()
}
