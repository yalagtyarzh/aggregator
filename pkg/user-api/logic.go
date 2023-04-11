package user_api

import "github.com/yalagtyarzh/aggregator/pkg/models"

type IUserAPILogic interface {
	GetReviews(pid int) ([]models.Review, error)
}
