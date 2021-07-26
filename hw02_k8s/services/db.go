package services

import "context"

type dbService interface {
	Connect(context.Context, string) error
	Close() error
}

func NewDbService() {

}