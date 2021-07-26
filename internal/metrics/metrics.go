package metrics

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/rubberyconf/rubberyconf/internal/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Metrics struct {
	client            *mongo.Client
	ctx               context.Context
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
)

func CreateMetrics() *Metrics {
	metrics = new(Metrics)
	metrics.connect()
	return metrics
}

func GetMetrics() *Metrics {
	return metrics
}

func (metric *Metrics) fetchMetrics(feature string) (*MongoMetrics, error) {

	/*
		conf := config.GetConfiguration()

		clientOptions := options.Client().ApplyURI(conf.Database.Url)
		client, err := mongo.NewClient(clientOptions)
		if err != nil {
			log.Fatal(err)
		}
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		err = client.Connect(ctx)
		if err != nil {
			log.Fatal(err)
		}
		defer client.Disconnect(ctx)
		quickstartDatabase := client.Database(conf.Database.DatabaseName)
		metricsCollection := quickstartDatabase.Collection(conf.Database.Collections.Metrics)
	*/

	newdocument := false
	var metricRegister MongoMetrics
	filter := bson.D{{"feature", feature}}
	err := metric.metricsCollection.FindOne(metrics.ctx, filter).Decode(&metricRegister)
	if err == mongo.ErrNoDocuments {
		newdocument = true
	} else if err != nil {
		log.Fatal(err)
		//return false, err
	}

	if newdocument {

		newDoc := MongoMetrics{
			Feature:   feature,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Counter:   0,
		}
		_, err := metric.metricsCollection.InsertOne(metric.ctx, newDoc)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		//err = metric.metricsCollection.FindOne(metrics.ctx, bson.M{"_id": feature}).Decode(&metricRegister)
		//if err != nil {
		//	log.Fatal(err)
		//	return nil, err
		//}
		return &newDoc, nil
	} else {
		//defer client.Disconnect(ctx)
		return &metricRegister, nil
	}
}
func (metric *Metrics) storeMetrics(metricRegister *MongoMetrics) (bool, error) {

	_, err := metric.metricsCollection.UpdateOne(metric.ctx,
		bson.D{{"feature", metricRegister.Feature}},
		bson.D{{"$set", bson.D{{"counter", metricRegister.Counter}}}})

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

	_, err := metric.metricsCollection.DeleteMany(metrics.ctx, bson.D{{"feature", feature}})
	if err != nil {
		return false, err
	}
	return true, nil
}

func (metric *Metrics) connect() {

	mongoOnce.Do(func() {

		conf := config.GetConfiguration()
		clientOptions := options.Client().ApplyURI(conf.Database.Url)
		client, err := mongo.NewClient(clientOptions)
		if err != nil {
			log.Fatal(err)
		}
		metrics.client = client
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		err = client.Connect(ctx)
		if err != nil {
			log.Fatal(err)
		}
		err = client.Ping(context.TODO(), nil)
		if err != nil {
			log.Fatal(err)
		}
		metrics.ctx = ctx
		//defer client.Disconnect(ctx)

		database := client.Database(conf.Database.DatabaseName)
		metricsCollection := database.Collection(conf.Database.Collections.Metrics)

		metrics.metricsCollection = metricsCollection
		//res, err := metricsCollection.InsertOne(ctx, bson.D{{"name", "pi"}, {"value", 3.14159}})
		//id := res.InsertedID
		//log.Printf("%d", id)
	})

}
