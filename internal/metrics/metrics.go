package metrics

import (
	"context"
	"log"
	"time"

	"github.com/rubberyconf/rubberyconf/internal/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Metrics struct {
	client            *mongo.Client
	ctx               context.Context
	metricsCollection *mongo.Collection
}

type MongoMetrics struct {
	_id       string    `json:"_id"`
	createdAt time.Time `json:"createdAt"`
	updatedAt time.Time `json:"updatedAt"`
	counter   int64     `json:"counter"`
}

var (
	metrics *Metrics
)

func CreateMetrics() *Metrics {
	metrics = new(Metrics)
	conf := config.GetConfiguration()
	clientOptions := options.Client().ApplyURI(conf.Database.Url)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	metrics.client = client
	ctx, _ := context.WithTimeout(context.TODO(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	metrics.ctx = ctx
	defer client.Disconnect(ctx)

	quickstartDatabase := client.Database(conf.Database.DatabaseName)
	metricsCollection := quickstartDatabase.Collection(conf.Database.Collections.Metrics)

	metrics.metricsCollection = metricsCollection

	return metrics
}

func GetMetrics() *Metrics {
	return metrics
}

func (metric *Metrics) fetchMetrics(feature string) (*MongoMetrics, error) {
	newdocument := false
	var metricRegister MongoMetrics
	err := metric.metricsCollection.FindOne(metrics.ctx, bson.M{"_id": feature}).Decode(&metricRegister)
	if err != nil {
		//log.Fatal(err)
		//return false, err
		newdocument = true
	}

	if newdocument {

		newDoc := MongoMetrics{
			_id:       feature,
			createdAt: time.Now(),
			updatedAt: time.Now(),
			counter:   0,
		}
		_, err := metric.metricsCollection.InsertOne(metric.ctx, newDoc)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		err = metric.metricsCollection.FindOne(metrics.ctx, bson.M{"_id": feature}).Decode(&metricRegister)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
	}
	return &metricRegister, nil
}
func (metric *Metrics) storeMetrics(metricRegister *MongoMetrics) (bool, error) {

	_, err := metric.metricsCollection.UpdateOne(metric.ctx, bson.M{"_id": metricRegister._id}, metricRegister)

	if err != nil {
		log.Fatal(err)
		return false, err
	}
	return true, nil
}

func (metricRegister *MongoMetrics) Update() {

	metricRegister.counter += 1
	metricRegister.updatedAt = time.Now()
}

func (metric *Metrics) Update(feature string) (*MongoMetrics, error) {

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

	_, err := metric.metricsCollection.DeleteOne(metrics.ctx, bson.M{"_id": feature})
	if err != nil {
		return false, err
	}
	return true, nil
}
