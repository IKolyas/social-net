package database

import (
	"social_network/internal/entities"
)

type PostgresRepository struct {
	master *PostgresMaster
	slaves []*PostgresSlave
}

func NewPostgresRepository(masterDSN string, slaveDSNs ...string) (*PostgresRepository, error) {
	master, err := NewPostgresMaster(masterDSN)
	if err != nil {
		return nil, err
	}

	var slaves []*PostgresSlave
	for _, dsn := range slaveDSNs {
		slave, err := NewPostgresSlave(dsn)
		if err != nil {
			return nil, err
		}
		slaves = append(slaves, slave)
	}

	return &PostgresRepository{
		master: master,
		slaves: slaves,
	}, nil
}

func (r *PostgresRepository) getNextSlave() *PostgresSlave {
	if r == nil || len(r.slaves) == 0 {
		return nil
	}
	nextSlave := r.slaves[0]
	r.slaves = append(r.slaves[1:], r.slaves[0])
	return nextSlave
}

func (r *PostgresRepository) CreatePost(post *entities.Post) error {
	return r.master.CreatePost(post)
}

func (r *PostgresRepository) UpdatePost(post *entities.Post) error {
	return r.master.UpdatePost(post)
}

func (r *PostgresRepository) DeletePost(postID string) error {
	return r.master.DeletePost(postID)
}

func (r *PostgresRepository) GetFeed(userID string, offset, limit int) ([]*entities.Post, error) {
	slave := r.getNextSlave()
	return slave.GetFeed(userID, offset, limit)
}

func (r *PostgresRepository) GetPostByID(postID string) (*entities.Post, error) {
	slave := r.getNextSlave()
	return slave.GetPostByID(postID)
}

func (r *PostgresRepository) CreateUser(user *entities.User) error {
	return r.master.CreateUser(user)
}

func (r *PostgresRepository) GetUserByID(userID string) (*entities.User, error) {
	slave := r.getNextSlave()
	return slave.GetUserByID(userID)
}

func (r *PostgresRepository) SearchUsers(query string) ([]*entities.User, error) {
	slave := r.getNextSlave()
	return slave.SearchUsers(query)
}

func (r *PostgresRepository) AddFriend(userID, friendID string) error {
	return r.master.AddFriend(userID, friendID)
}

func (r *PostgresRepository) DeleteFriend(userID, friendID string) error {
	return r.master.DeleteFriend(userID, friendID)
}
