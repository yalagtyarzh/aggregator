package user_api

import (
	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
	"github.com/yalagtyarzh/aggregator/pkg/errors"
	"github.com/yalagtyarzh/aggregator/pkg/logger"
	"github.com/yalagtyarzh/aggregator/pkg/models"
	"github.com/yalagtyarzh/aggregator/pkg/repo"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type UserAPILogic struct {
	log  logger.ILogger
	repo *repo.UserAPIRepository
}

func NewUserAPILogic(repositoryPool *repo.UserAPIRepository, log logger.ILogger) IUserAPILogic {
	return &UserAPILogic{
		log:  log,
		repo: repositoryPool,
	}
}

func (l *UserAPILogic) GetReviews(productId int) ([]models.Review, error) {
	r, err := l.repo.DB.GetReviewsByProductID(productId)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (l *UserAPILogic) GetProduct(productId int) (models.Product, error) {
	p, err := l.repo.DB.GetProductWithDeleted(productId, false)
	if err != nil {
		return models.Product{}, nil
	}

	if p == nil {
		return models.Product{}, errors.ErrNoProduct
	}

	return *p, nil
}

func (l *UserAPILogic) GetProducts(after, limit, year int, genre string) ([]models.Product, error) {
	s, err := l.repo.DB.GetProductsWithFilter(after, limit, year, genre, false)
	if err != nil {
		return []models.Product{}, err
	}

	return s, nil
}

func (l *UserAPILogic) CreateReview(rc models.ReviewCreate, userID uuid.UUID) error {
	u, err := l.repo.DB.GetUserByID(userID)
	if err != nil {
		return err
	}

	if u == nil {
		return errors.ErrNoUser
	}

	r, err := l.repo.DB.GetReviewByUserAndProduct(rc.ProductID, userID)
	if err != nil {
		return err
	}

	if r != nil {
		return errors.ErrTooManyReviews
	}

	return l.repo.DB.InsertReview(rc, userID)
}

func (l *UserAPILogic) UpdateReview(rc models.ReviewUpdate, userID uuid.UUID) error {
	u, err := l.repo.DB.GetUserByID(userID)
	if err != nil {
		return err
	}

	if u == nil {
		return errors.ErrNoUser
	}

	r, err := l.repo.DB.GetReviewByID(rc.ID)
	if err != nil {
		return err
	}

	if r == nil {
		return errors.ErrNoReview
	}

	if r.UserID.String() != userID.String() {
		if u.Role != "Moderator" && u.Role != "Admin" {
			return errors.ErrNoPermissions
		}
	}

	if rc.Delete {
		return l.repo.DB.DeleteReview(rc.ID)
	}

	return l.repo.DB.UpdateReview(rc)
}

func (l *UserAPILogic) CreateUser(req models.CreateUser) (models.UserResponse, error) {
	hashedP, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.MinCost)
	if err != nil {
		return models.UserResponse{}, err
	}

	req.Password = string(hashedP)
	id, err := l.repo.DB.InsertUser(req)
	if err == repo.ErrForeignKeyViolation {
		return models.UserResponse{}, errors.ErrInvalidRole
	}

	if err == repo.ErrUniqueViolation {
		return models.UserResponse{}, errors.ErrUserAlreadyCreated
	}

	access, refresh, err := l.repo.JWT.GenerateTokens(id, "Registered", req.Email)
	if err != nil {
		return models.UserResponse{}, err
	}

	err = l.repo.JWT.SaveToken(id, refresh)
	if err != nil {
		return models.UserResponse{}, err
	}

	return models.UserResponse{
		AccessToken: access, RefreshToken: refresh, UserID: id.String(), Email: req.Email,
	}, err
}

func (l *UserAPILogic) Login(username, password string) (models.UserResponse, error) {
	u, err := l.repo.DB.GetUserByUsername(username)
	if err != nil {
		return models.UserResponse{}, err
	}

	if u == nil {
		return models.UserResponse{}, errors.ErrNoUser
	}

	if err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return models.UserResponse{}, errors.ErrInvalidPassword
	}

	access, refresh, err := l.repo.JWT.GenerateTokens(u.ID, u.Role, u.Email)
	if err != nil {
		return models.UserResponse{}, err
	}

	err = l.repo.JWT.SaveToken(u.ID, refresh)
	if err != nil {
		return models.UserResponse{}, err
	}

	return models.UserResponse{
		AccessToken: access, RefreshToken: refresh, UserID: u.ID.String(), Email: u.Email,
	}, err
}

func (l *UserAPILogic) Logout(token string) error {
	return l.repo.DB.DeleteToken(token)
}

func (l *UserAPILogic) Refresh(token string) (models.UserResponse, error) {
	resp, err := l.repo.JWT.ValidateRefreshToken(token)
	if err != nil {
		return models.UserResponse{}, err
	}

	t, err := l.repo.DB.FindToken(token)
	if t == nil {
		return models.UserResponse{}, errors.ErrNoToken
	}

	u, err := l.repo.DB.GetUserByID(resp.UserID)
	if err != nil {
		return models.UserResponse{}, err
	}

	if u == nil {
		return models.UserResponse{}, errors.ErrNoUser
	}

	access, refresh, err := l.repo.JWT.GenerateTokens(u.ID, u.Role, u.Email)
	if err != nil {
		return models.UserResponse{}, err
	}

	err = l.repo.JWT.SaveToken(u.ID, refresh)
	if err != nil {
		return models.UserResponse{}, err
	}

	return models.UserResponse{
		AccessToken: access, RefreshToken: refresh, UserID: u.ID.String(), Email: u.Email,
	}, err
}

func (l *UserAPILogic) ListGenres() ([]models.Genre, error) {
	genres, err := l.repo.DB.SelectGenres()
	if err != nil {
		return nil, err
	}

	return genres, nil
}

func (l *UserAPILogic) GraphqlList(query string) (*graphql.Result, error) {
	p, err := l.repo.DB.GetProducts()
	if err != nil {
		return nil, err
	}

	// graphql schema
	fields := GetFields(p)
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		return nil, err
	}

	params := graphql.Params{Schema: schema, RequestString: query}
	resp := graphql.Do(params)
	if len(resp.Errors) > 0 {
		l.log.Errorf("%+v", resp.Errors)
		return nil, errors.ErrInvalidGraphql
	}

	return resp, nil
}

func GetFields(products []models.Product) graphql.Fields {
	return graphql.Fields{
		"product": &graphql.Field{
			Type:        productType,
			Description: "Get product by id",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id, ok := p.Args["id"].(int)
				if ok {
					for _, v := range products {
						if v.ID == id {
							return v, nil
						}
					}
				}
				return nil, nil
			},
		},
		"list": &graphql.Field{
			Type:        graphql.NewList(productType),
			Description: "Get all products",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return products, nil
			},
		},
		"search": &graphql.Field{
			Type:        graphql.NewList(productType),
			Description: "Search products by title",
			Args: graphql.FieldConfigArgument{
				"titleContains": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				var list []models.Product
				search, ok := params.Args["titleContains"].(string)
				if ok {
					for _, v := range products {
						if strings.Contains(strings.ToLower(v.Title), strings.ToLower(search)) {
							list = append(list, v)
						}
					}
				}
				return list, nil
			},
		},
	}
}
