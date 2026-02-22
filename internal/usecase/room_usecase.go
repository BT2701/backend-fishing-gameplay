package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/BT2701/backend-fishing-gameplay/internal/domain/entity"
	"github.com/BT2701/backend-fishing-gameplay/internal/domain/port"
)

var (
	ErrRoomAlreadyExists = errors.New("room already exists")
	ErrRoomNotFound      = errors.New("room not found")
	ErrRoomFull          = errors.New("room is full")
	ErrSeatTaken         = errors.New("seat is taken")
	ErrInvalidBalance    = errors.New("balance must be >= 0")
	ErrInvalidMaxPlayers = errors.New("max players must be > 0")
	ErrInvalidSeat       = errors.New("seat id must be >= 0")
	ErrPlayerInOtherRoom = errors.New("player is in another room")
	ErrPlayerNotInRoom   = errors.New("player is not in room")
	ErrPlayerAlreadyIn   = errors.New("player already in room")
)

type RoomUsecase struct {
	roomRepo   port.RoomRepository
	playerRepo port.PlayerRepository
	now        func() time.Time
}

func NewRoomUsecase(roomRepo port.RoomRepository, playerRepo port.PlayerRepository) *RoomUsecase {
	return &RoomUsecase{
		roomRepo:   roomRepo,
		playerRepo: playerRepo,
		now:        time.Now,
	}
}

func (uc *RoomUsecase) CreateRoom(ctx context.Context, roomID string, maxPlayers int) (*entity.Room, error) {
	if maxPlayers <= 0 {
		return nil, ErrInvalidMaxPlayers
	}
	if roomID == "" {
		return nil, errors.New("room id is required")
	}

	existing, err := uc.roomRepo.GetByID(ctx, roomID)
	if err == nil && existing != nil {
		return nil, ErrRoomAlreadyExists
	}
	if err != nil && !errors.Is(err, port.ErrNotFound) {
		return nil, err
	}

	room := &entity.Room{
		RoomID:  roomID,
		Status:  "open",
		Players: map[string]*entity.Player{},
		FishMap: map[string]*entity.FishInstance{},
		Config: entity.RoomConfig{
			MaxPlayers: maxPlayers,
		},
	}

	if err := uc.roomRepo.Save(ctx, room); err != nil {
		return nil, err
	}

	return room, nil
}

func (uc *RoomUsecase) JoinRoom(ctx context.Context, roomID, playerID string, seatID int, initialBalance int64) (*entity.Room, *entity.Player, error) {
	if seatID < 0 {
		return nil, nil, ErrInvalidSeat
	}
	if initialBalance < 0 {
		return nil, nil, ErrInvalidBalance
	}

	room, err := uc.roomRepo.GetByID(ctx, roomID)
	if err != nil {
		if errors.Is(err, port.ErrNotFound) {
			return nil, nil, ErrRoomNotFound
		}
		return nil, nil, err
	}

	if room.Players == nil {
		room.Players = map[string]*entity.Player{}
	}

	if room.Config.MaxPlayers > 0 && len(room.Players) >= room.Config.MaxPlayers {
		return nil, nil, ErrRoomFull
	}

	for _, p := range room.Players {
		if p.SeatID == seatID {
			return nil, nil, ErrSeatTaken
		}
	}

	player, err := uc.playerRepo.GetByID(ctx, playerID)
	if err != nil {
		if !errors.Is(err, port.ErrNotFound) {
			return nil, nil, err
		}
		player = &entity.Player{
			PlayerID: playerID,
			Balance:  initialBalance,
		}
	}

	if player.RoomID != "" && player.RoomID != roomID {
		return nil, nil, ErrPlayerInOtherRoom
	}
	if player.Balance < 0 {
		return nil, nil, ErrInvalidBalance
	}

	if _, exists := room.Players[playerID]; exists {
		return nil, nil, ErrPlayerAlreadyIn
	}

	player.SeatID = seatID
	player.RoomID = roomID
	player.IsOnline = true
	player.LastActionAt = uc.now().Unix()

	room.Players[playerID] = player

	if err := uc.roomRepo.Save(ctx, room); err != nil {
		return nil, nil, err
	}
	if err := uc.playerRepo.Save(ctx, player); err != nil {
		return nil, nil, err
	}

	return room, player, nil
}

func (uc *RoomUsecase) LeaveRoom(ctx context.Context, roomID, playerID string) (*entity.Room, *entity.Player, error) {
	room, err := uc.roomRepo.GetByID(ctx, roomID)
	if err != nil {
		if errors.Is(err, port.ErrNotFound) {
			return nil, nil, ErrRoomNotFound
		}
		return nil, nil, err
	}

	player, ok := room.Players[playerID]
	if !ok {
		return nil, nil, ErrPlayerNotInRoom
	}

	delete(room.Players, playerID)

	player.RoomID = ""
	player.IsOnline = false
	player.SeatID = 0
	player.LastActionAt = uc.now().Unix()

	if err := uc.roomRepo.Save(ctx, room); err != nil {
		return nil, nil, err
	}
	if err := uc.playerRepo.Save(ctx, player); err != nil {
		return nil, nil, err
	}

	return room, player, nil
}
