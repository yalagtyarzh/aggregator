package user_api

import (
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/yalagtyarzh/aggregator/pkg/logger"
	"github.com/yalagtyarzh/aggregator/pkg/models"
	"github.com/yalagtyarzh/aggregator/pkg/repo"
	"golang.org/x/crypto/bcrypt"
	"time"
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
	p, err := l.repo.DB.GetProductById(productId)
	if err != nil {
		return models.Product{}, nil
	}

	return *p, nil
}

func (l *UserAPILogic) GetProducts(after, limit, year int, genre string) ([]models.Product, error) {
	s, err := l.repo.DB.GetProducts(after, limit, year, genre)
	if err != nil {
		return []models.Product{}, err
	}

	return s, nil
}

func (l *UserAPILogic) CreateReview(rc models.ReviewCreate, userID uuid.UUID) error {
	return l.repo.DB.InsertReview(rc, userID)
}

func (l *UserAPILogic) UpdateReview(rc models.ReviewUpdate, userID uuid.UUID) error {
	r, err := l.repo.DB.GetReviewByID(rc.ID)
	if err != nil {
		return err
	}

	if r.User.ID.String() != userID.String() {
		perms, err := l.repo.DB.GetPermissionsByRole(userID)
		if err != nil {
			return err
		}

		var flag bool
		for _, v := range perms {
			if v.Permission == editForeignReviews {
				flag = true
				break
			}
		}

		if !flag {
			return errNoPermissions
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
		return models.UserResponse{}, errInvalidRole
	}

	if err == repo.ErrUniqueViolation {
		return models.UserResponse{}, errUserAlreadyCreated
	}

	access, refresh, err := l.generateTokens(id, req.Email)
	if err != nil {
		return models.UserResponse{}, err
	}

	err = l.saveToken(id, refresh)
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

	if err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return models.UserResponse{}, errInvalidPassword
	}

	access, refresh, err := l.generateTokens(u.ID, u.Email)
	if err != nil {
		return models.UserResponse{}, err
	}

	err = l.saveToken(u.ID, refresh)
	if err != nil {
		return models.UserResponse{}, err
	}

	return models.UserResponse{
		AccessToken: access, RefreshToken: refresh, UserID: u.ID.String(), Email: u.Email,
	}, err
}

func (l *UserAPILogic) generateTokens(userId uuid.UUID, email string) (string, string, error) {
	accessPayload := models.TokenPayload{
		UserID: userId,
		Email:  email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(30 * time.Minute).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	refreshPayload := models.TokenPayload{
		UserID: userId,
		Email:  email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(30 * 24 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessPayload)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshPayload)

	aTokenStr, err := accessToken.SignedString([]byte(l.repo.JWT.SigningKey))
	if err != nil {
		return "", "", err
	}

	rTokenStr, err := refreshToken.SignedString([]byte(l.repo.JWT.SigningKey))
	if err != nil {
		return "", "", err
	}

	return aTokenStr, rTokenStr, nil
}

func (l *UserAPILogic) saveToken(userId uuid.UUID, refreshToken string) error {
	t, err := l.repo.DB.GetToken(userId)
	if err != nil {
		return err
	}

	if t != nil {
		return l.repo.DB.UpdateToken(userId, refreshToken)
	}

	return l.repo.DB.InsertToken(userId, refreshToken)
}
