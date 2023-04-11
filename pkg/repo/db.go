package repo

import "github.com/yalagtyarzh/aggregator/pkg/models"

type IDB interface {
	GetReviewsByProductID(pid int) ([]models.Review, error)
}
