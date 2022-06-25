package main

import (
	"flag"
	"log"

	"github.com/go-playground/validator/v10"
	"gitlab.com/jacob-ernst/mets/pkg/models"
)

type CreateInput struct {
	Name string  `validate:"required"`
	DSN  string  `validate:"file"`
	MET  float64 `validate:"required,gte=0.5"`
}

func NewCreateCommand() *CreateCommand {
	cc := &CreateCommand{
		fs: flag.NewFlagSet("create", flag.ContinueOnError),
	}

	cc.fs.StringVar(&cc.dsn, "dsn", "data/dev.db", "dsn for the DB to migrate")
	cc.fs.StringVar(&cc.name, "name", "", "name of the activity")
	cc.fs.StringVar(&cc.desc, "desc", "", "description for the activity")
	cc.fs.Float64Var(&cc.met, "met", -1, "MET value for the activity")

	return cc
}

type CreateCommand struct {
	fs *flag.FlagSet

	name string `validate:"required"`
	dsn  string `validate:"file"`
	desc string
	met  float64 `validate:"required,gte=0.5"`
}

func (g *CreateCommand) Name() string {
	return g.fs.Name()
}

func (g *CreateCommand) Init(args []string) error {
	return g.fs.Parse(args)
}

func (g *CreateCommand) Validate() error {
	validator := validator.New()
	input := CreateInput{Name: g.name, DSN: g.dsn, MET: g.met}

	err := validator.Struct(input)
	if err != nil {
		return err
	}

	return nil
}

func (g *CreateCommand) Run() error {
	db, err := models.OpenDB(g.dsn)
	if err != nil {
		return err
	}

	activity := models.Activity{Name: g.name, Description: g.desc, Effort: g.met}

	tx := db.Create(&activity)
	if tx.Error != nil {
		return tx.Error
	}

	log.Printf("[id: %v, name: %v] Successfully created activity", activity.ID, activity.Name)

	return nil
}
