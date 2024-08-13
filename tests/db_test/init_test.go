package db_test

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/arynpd/home-mgmt-service/db"
	"github.com/joho/godotenv"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

var dbPool = &db.Db{}

func TestMain(m *testing.M) {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	docker_pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not construct pool: %s", err)
	}

	err = docker_pool.Client.Ping()
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	resource, err := docker_pool.RunWithOptions(&dockertest.RunOptions{
		Repository: os.Getenv("TEST_REPO"),
		Tag:        os.Getenv("TEST_TAG"),
		Env: []string{
			"POSTGRES_USER=" + os.Getenv("TEST_USER"),
			"POSTGRES_PASSWORD=" + os.Getenv("TEST_PASSWORD"),
			"POSTGRES_DB=" + os.Getenv("TEST_DB"),
		},
	}, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{
			Name: "no",
		}
	})

	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	if err := docker_pool.Retry(func() error {
		connString := fmt.Sprintf(os.Getenv("TEST_DB_URL"), resource.GetPort("5432/tcp"))
		return dbPool.Init(connString)
	}); err != nil {
		log.Fatalf("Could not connect to database; %s", err)
	}

	defer func() {
		if err := docker_pool.Purge(resource); err != nil {
			log.Fatalf("Could not purge resource: %s", err)
		}
	}()

	err = setupSchema()
	if err != nil {
		log.Fatalf("Could not setup schema: %s", err)
	}

	m.Run()
}

func setupDb() *db.Db {
	return dbPool
}

func setupSchema() error {
	path := filepath.Join("..", "..", "utils", "schema.sql")
	return dbPool.ExecFile(path)
}
