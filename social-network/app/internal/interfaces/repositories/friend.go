package repositories

import "social_network/internal/infrastructure/database"

type FriendRepository interface {
	AddFriend(userID, friendID string) error
	DeleteFriend(userID, friendID string) error
}

type FriendRepositoryImpl struct {
	db *database.PostgresRepository
}

func NewFriendRepository(db *database.PostgresRepository) *FriendRepositoryImpl {
	return &FriendRepositoryImpl{db: db}
}

func (r *FriendRepositoryImpl) AddFriend(userID, friendID string) error {
	return r.db.AddFriend(userID, friendID)
}

func (r *FriendRepositoryImpl) DeleteFriend(userID, friendID string) error {
	return r.db.DeleteFriend(userID, friendID)
}
