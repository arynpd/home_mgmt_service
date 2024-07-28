package db_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/arynpd/home-mgmt-service/db"
	"github.com/jackc/pgx/v5/pgxpool"
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

	var pool *pgxpool.Pool
	if err := docker_pool.Retry(func() error {
		var err error
		pool, err = pgxpool.New(context.Background(), fmt.Sprintf(os.Getenv("TEST_DB_URL"), resource.GetPort("5432/tcp")))
		if err != nil {
			return err
		}
		return pool.QueryRow(context.Background(), "Select 1").Scan(new(int))
	}); err != nil {
		log.Fatalf("Could not connect to database; %s", err)
	}

	defer func() {
		if err := docker_pool.Purge(resource); err != nil {
			log.Fatalf("Could not purge resource: %s", err)
		}
	}()

	dbPool.Pool = pool
	err = setupSchema()
	if err != nil {
		log.Fatalf("Could not setup schema: %s", err)
	}

	m.Run()
}

func setupSchema() error {
	path := filepath.Join("..", "utils", "schema.sql")

	c, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	sql := string(c)
	_, err = dbPool.Pool.Exec(context.Background(), sql)
	if err != nil {
		return err
	}
	return nil
}

func setupTransaction() *db.Transaction {
	transaction, err := dbPool.BeginTransaction()
	if err != nil {
		log.Fatalf("Error starting transaction: %s", err)
	}
	return transaction
}
