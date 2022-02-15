package tests

import (
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/ory/dockertest/v3"
	"github.com/pkg/errors"
	"log"
	"stonehenge/app/config"
	"testing"
)

var (
	port string
	db   *pgxpool.Pool
	env  = config.DatabaseConfigurations{
		User:           "postgres",
		Password:       "postgres",
		Host:           "localhost",
		Port:           "5432",
		SSLMode:        "disable",
		MigrationsPath: "../migrations",
	}
)

func SetupTest(m *testing.M) int {
	teardown, err := createContainer()
	if err != nil {
		log.Fatal(err)
	}

	defer teardown()
	return m.Run()
}

func createContainer() (func(), error) {
	dbName := createRandomName()
	pool, res, err := createResource(env.User, dbName, env.Password)
	if err != nil {
		return nil, err
	}
	teardown := func() {
		err = pool.Purge(res)
	}

	port = res.GetPort("5432/tcp")

	if err := pool.Retry(func() error {
		db, err = connect(dbName, env)
		return err
	}); err != nil {
		teardown()
		return nil, errors.Wrap(err, "an error occurred when setting up the database")
	}

	return teardown, nil

}

func createResource(userName, dbName, userPassword string) (*dockertest.Pool, *dockertest.Resource, error) {
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
		fmt.Sprintf("POSTGRES_DB=%s", dbName),
	})

	if err != nil {
		return pool, resource, errors.Wrap(err, "the docker container could not be created")
	}

	resource.Expire(120)

	return pool, resource, nil
}
