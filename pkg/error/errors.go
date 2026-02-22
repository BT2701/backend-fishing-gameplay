package apperr

import "errors"

type Code string

type Error struct {
	Code    Code
	Message string
}

func (e *Error) Error() string {
	return e.Message
}

func (e *Error) Is(target error) bool {
	other, ok := target.(*Error)
	if !ok {
		return false
	}
	return e.Code == other.Code
}

func New(code Code, message string) *Error {
	return &Error{Code: code, Message: message}
}

func CodeOf(err error) Code {
	if err == nil {
		return ""
	}
	var appErr *Error
	if errors.As(err, &appErr) {
		return appErr.Code
	}
	return ""
}

const (
	CodeNotFound              Code = "NOT_FOUND"
	CodeRoomAlreadyExists     Code = "ROOM_ALREADY_EXISTS"
	CodeRoomNotFound          Code = "ROOM_NOT_FOUND"
	CodeRoomFull              Code = "ROOM_FULL"
	CodeSeatTaken             Code = "SEAT_TAKEN"
	CodeInvalidBalance        Code = "INVALID_BALANCE"
	CodeInvalidMaxPlayers     Code = "INVALID_MAX_PLAYERS"
	CodeInvalidSeat           Code = "INVALID_SEAT"
	CodePlayerInOtherRoom     Code = "PLAYER_IN_OTHER_ROOM"
	CodePlayerNotInRoom       Code = "PLAYER_NOT_IN_ROOM"
	CodePlayerAlreadyInRoom   Code = "PLAYER_ALREADY_IN_ROOM"
	CodeInvalidRoomID         Code = "INVALID_ROOM_ID"
	CodeFishTypeNotFound      Code = "FISH_TYPE_NOT_FOUND"
	CodeFishUIDExists         Code = "FISH_UID_EXISTS"
	CodeInvalidFishID         Code = "INVALID_FISH_ID"
	CodeInvalidFishUID        Code = "INVALID_FISH_UID"
	CodeFishNotFound          Code = "FISH_NOT_FOUND"
	CodeFishAlreadyDead       Code = "FISH_ALREADY_DEAD"
	CodeGunNotFound           Code = "GUN_NOT_FOUND"
	CodeInsufficientBalance   Code = "INSUFFICIENT_BALANCE"
	CodeInvalidPlayerID       Code = "INVALID_PLAYER_ID"
	CodePlayerNotFound        Code = "PLAYER_NOT_FOUND"
	CodeInvalidRTPDelta       Code = "INVALID_RTP_DELTA"
	CodeBulletConfigNotFound  Code = "BULLET_CONFIG_NOT_FOUND"
	CodeGameConfigNotFound    Code = "GAME_CONFIG_NOT_FOUND"
	CodeGameFeaturesNotFound  Code = "GAME_FEATURES_NOT_FOUND"
	CodeGamePathsNotFound     Code = "GAME_PATHS_NOT_FOUND"
	CodeGameRTPNotFound       Code = "GAME_RTP_NOT_FOUND"
	CodeGameFishTypesNotFound Code = "GAME_FISH_TYPES_NOT_FOUND"
)

var (
	ErrNotFound              = New(CodeNotFound, "not found")
	ErrRoomAlreadyExists     = New(CodeRoomAlreadyExists, "room already exists")
	ErrRoomNotFound          = New(CodeRoomNotFound, "room not found")
	ErrRoomFull              = New(CodeRoomFull, "room is full")
	ErrSeatTaken             = New(CodeSeatTaken, "seat is taken")
	ErrInvalidBalance        = New(CodeInvalidBalance, "balance must be >= 0")
	ErrInvalidMaxPlayers     = New(CodeInvalidMaxPlayers, "max players must be > 0")
	ErrInvalidSeat           = New(CodeInvalidSeat, "seat id must be >= 0")
	ErrPlayerInOtherRoom     = New(CodePlayerInOtherRoom, "player is in another room")
	ErrPlayerNotInRoom       = New(CodePlayerNotInRoom, "player is not in room")
	ErrPlayerAlreadyIn       = New(CodePlayerAlreadyInRoom, "player already in room")
	ErrInvalidRoomID         = New(CodeInvalidRoomID, "room id is required")
	ErrFishTypeNotFound      = New(CodeFishTypeNotFound, "fish type not found")
	ErrFishUIDExists         = New(CodeFishUIDExists, "fish uid already exists")
	ErrInvalidFishID         = New(CodeInvalidFishID, "fish id must be > 0")
	ErrInvalidFishUID        = New(CodeInvalidFishUID, "fish uid is required")
	ErrFishNotFound          = New(CodeFishNotFound, "fish not found")
	ErrFishAlreadyDead       = New(CodeFishAlreadyDead, "fish already dead")
	ErrGunNotFound           = New(CodeGunNotFound, "gun not found")
	ErrInsufficientBalance   = New(CodeInsufficientBalance, "insufficient balance")
	ErrInvalidPlayerID       = New(CodeInvalidPlayerID, "player id is required")
	ErrPlayerNotFound        = New(CodePlayerNotFound, "player not found")
	ErrInvalidRTPDelta       = New(CodeInvalidRTPDelta, "rtp delta must not be negative")
	ErrBulletConfigNotFound  = New(CodeBulletConfigNotFound, "bullet config not found")
	ErrGameConfigNotFound    = New(CodeGameConfigNotFound, "game config not found")
	ErrGameFeaturesNotFound  = New(CodeGameFeaturesNotFound, "game features not found")
	ErrGamePathsNotFound     = New(CodeGamePathsNotFound, "game paths not found")
	ErrGameRTPNotFound       = New(CodeGameRTPNotFound, "game rtp not found")
	ErrGameFishTypesNotFound = New(CodeGameFishTypesNotFound, "game fish types not found")
)
