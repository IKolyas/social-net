package usecases

import (
	"social_network/internal/interfaces/repositories"
)

type FriendUseCase struct {
	friendRepo repositories.FriendRepository
}

func NewFriendUseCase(friendRepo repositories.FriendRepository) *FriendUseCase {
	return &FriendUseCase{friendRepo: friendRepo}
}

func (uc *FriendUseCase) AddFriend(userID, friendID string) error {
	return uc.friendRepo.AddFriend(userID, friendID)
}

func (uc *FriendUseCase) DeleteFriend(userID, friendID string) error {
	return uc.friendRepo.DeleteFriend(userID, friendID)
}
