package main

import (
	"context"
	"log"

	pb "github.com/d-vignesh/shipper/user-service/proto/user"
	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
)

type User struct {
	ID 			string `sql:"id"`
	Name		string `sql:"name"`
	Email    	string `sql:"email"`
	Company		string `sql:"company"`
	Password    string `sql:"password"`
}

func MarshalUserCollection(users []*pb.User) []*User {
	u := make([]*User, len(users))
	for _, val := range users {
		u = append(u, MarshalUser(val))
	}
	return u 
}

func MarshalUser(*pb.User) *User {
	return &User{
		ID:		 user.Id,
		Name:	 user.Name,
		Email: 	 user.Email,
		Company: user.Company,
		Password:user.Password,
	}
}

func UnmarsalUserCollection(users []*User) []*pb.User {
	u := make([]*pb.User, len(users))
	for _, val := range users {
		u = append(u, UnmarshalUser(val))
	}
	return u 
}

func UnmarshalUser(user *User) *pb.User {
	return &pb.User{
		Id:		 user.ID,
		Name:	 user.Name,
		Email:	 user.Email,
		Company: user.Company,
		Password:user.Password,
	}
}

type Repository interface {
	GetAll(ctx context.Context) ([]*User, error)
	Get(ctx contex.Context, id string) (*User, error)
	Create(ctx context.Context, user *User) error
	GetByEmail(ctx context.Context, user *User) (*User, error)
}

type PostgresRepository struct {
	db *sqlx.DB
}

func NewPostgresRepository(db *sqlx.DB) *PostgresRepository {
	return &PostgresRepository{db}
}

func (repo *PostgresRepository) GetAll(ctx context.Context) ([]*User, error) {
	users := make([]*User, 0)
	if err := repo.db.GetContext(ctx, users, "select * from users"); err != nil {
		return users, err
	}
	return users, nil
}

func (repo *PostgresRepository) Get(ctx context.Context, id string) (*User, error) {
	var user *User
	if err := repo.db.GetContext(ctx, &user, "select * from users where id = $1", id); err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *PostgresRepository) Create(ctx context.Context, user *User) error {
	user.ID := uuid.NewV4().String()
	log.Println(user)
	query := "insert in users (id, name, email, company, password) values ($1, $2, $3, $4, $5)"
	_, err := repo.db.ExecContext(ctx, query, user.ID, user.Name, user.Email, user.Company, user.Password)
	return err
}

func (repo *PostgresRepository) GetByEmail(ctx context.Context, email string) (*User, error) {
	query := "select * from users where email = $1"
	var user *User
	if err := repo.db.GetContext(ctx, &user, query, email); err != nil {
		return nil, err
	}
	return user
}

