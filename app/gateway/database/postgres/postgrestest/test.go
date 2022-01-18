package postgrestest

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/ory/dockertest/v3"
	"github.com/pkg/errors"
	"log"
	"testing"
)

const (
	testUser     = "postgres"
	testPassword = "password"
	testHost     = "localhost"
	testDatabase = "stonehenge"
)

var port string
var db *pgxpool.Pool

func NewCleanDatabase() (*pgxpool.Pool, error) {
	ctx := context.Background()
	err := RecycleDatabase(ctx)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func SetupTest(m *testing.M) int {
	teardown, err := createContainer()
	if err != nil {
		log.Fatal(err)
	}

	defer teardown()

	return m.Run()
}

func createContainer() (func(), error) {

	pool, res, err := createResource(testUser, testDatabase, testPassword)
	if err != nil {
		return nil, err
	}
	teardown := func() {
		err = pool.Purge(res)
	}

	port = res.GetPort("5432/tcp")

	if err := pool.Retry(func() error {
		db, err = connect(testUser, testPassword, testHost, port, testDatabase)
		log.Printf("Error connecting to the database: %v", err)
		return err
	}); err != nil {
		teardown()
		return nil, errors.Wrap(err, "an error occurred when setting up the database")
	}

	return teardown, nil

}

func createResource(userName, databaseName, userPassword string) (*dockertest.Pool, *dockertest.Resource, error) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		return pool, nil, errors.Wrap(err, "the docker pool connection could not be established")
	}

	if err := pool.Client.Ping(); err != nil {
		return pool, nil, errors.Wrap(err, "could not contact docker pool")
	}

	resource, err := pool.Run("postgres", "latest", []string{
		fmt.Sprintf("POSTGRES_USER=%s", userName),
		fmt.Sprintf("POSTGRES_PASSWORD=%s", userPassword),
		fmt.Sprintf("POSTGRES_DB=%s", databaseName),
	})

	log.Printf("Docker test container status: %v", resource.Container.State.Status)

	if err != nil {
		return pool, resource, errors.Wrap(err, "the docker container could not be created")
	}

	resource.Expire(120)

	return pool, resource, nil
}
