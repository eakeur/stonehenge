package postgrestest

import (
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/ory/dockertest/v3"
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

func GetDB() *pgxpool.Pool {
	return db
}

func SetupTest(m *testing.M) int {
	teardown, err := createContainer()
	if err != nil {
		log.Fatalf("an error occurred and it was not possible to create database container: %v", err)
	}

	defer teardown()

	//err = fakes.PopulateDatabase(db)
	//if err != nil {
	//	teardown()
	//	log.Fatalf("could not populate database: %e", err)
	//}

	return m.Run()
}


func createContainer() (func(), error) {

	pool, res, err := createResource(testDatabase, testPassword)
	if err != nil {
		return nil, err
	}

	teardown := func() {
		err = pool.Purge(res)
	}

	port = res.GetPort("5432/tcp")

	if err := pool.Retry(func() error {
		db, err = connect(testUser, testPassword, testHost, port, testDatabase)
		return err
	}); err != nil {
		teardown()
		return nil, err
	}

	return teardown, nil

}

func createResource(databaseName, userPassword string) (*dockertest.Pool, *dockertest.Resource, error) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		return pool, nil, err
	}

	resource, err := pool.Run("postgres", "latest", []string{
		fmt.Sprintf("POSTGRES_PASSWORD=%s", userPassword),
		fmt.Sprintf("POSTGRES_DB=%s", databaseName),
	})
	if err != nil {
		return pool, resource, err
	}

	return pool, resource, nil
}
