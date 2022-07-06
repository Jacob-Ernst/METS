package main

import (
	"flag"
	"fmt"

	"github.com/go-playground/validator/v10"
	"gitlab.com/jacob-ernst/mets/pkg/models"
)

type ListInput struct {
	DSN   string `validate:"file"`
	Limit int    `validate:"gte=1"`
}

func NewListCommand() *ListCommand {
	cc := &ListCommand{
		fs: flag.NewFlagSet("list", flag.ContinueOnError),
	}

	cc.fs.StringVar(&cc.dsn, "dsn", "data/dev.db", "dsn for the DB to migrate")
	cc.fs.IntVar(&cc.limit, "limit", 3, "optional limit for number of results")
	cc.fs.StringVar(&cc.query, "query", "", "optional search query to filter by")
	return cc
}

type ListCommand struct {
	fs *flag.FlagSet

	dsn   string
	limit int
	query string
}

func (l *ListCommand) Name() string {
	return l.fs.Name()
}

func (l *ListCommand) Init(args []string) error {
	return l.fs.Parse(args)
}

func (l *ListCommand) Validate() error {
	validator := validator.New()
	input := ListInput{DSN: l.dsn, Limit: l.limit}

	err := validator.Struct(input)
	if err != nil {
		return err
	}

	return nil
}

func (l *ListCommand) Run() error {
	var activities []models.Activity
	db, err := models.OpenDB(l.dsn)
	if err != nil {
		return err
	}

	tx := db

	if l.query != "" {
		query := fmt.Sprintf("%%%v%%", l.query)
		tx = db.Where("name LIKE ?", query)
	}

	tx = tx.Select("name", "effort", "description").Order("effort desc").Limit(l.limit).Find(&activities)

	if tx.Error != nil {
		return tx.Error
	}

	for _, a := range activities {
		fmt.Printf("[name: %v, desc: %v, mets: %v]\n", a.Name, a.Description, a.Effort)
	}

	return nil
}
