package models

import (
	"sync"
)

type UserState struct {
	ID    int64
	State string
	Data  map[string]interface{}
}

var (
	userStates = make(map[int64]*UserState)
	mu         sync.Mutex
)

func GetUserState(userID int64) *UserState {
	mu.Lock()
	defer mu.Unlock()
	if state, exists := userStates[userID]; exists {
		return state
	}
	userState := &UserState{ID: userID, State: "root", Data: make(map[string]interface{})}
	userStates[userID] = userState
	return userState
}

func UpdateUserState(userID int64, state string, data map[string]interface{}) {
	mu.Lock()
	defer mu.Unlock()
	if userState, exists := userStates[userID]; exists {
		userState.State = state
		userState.Data = data
	} else {
		userStates[userID] = &UserState{ID: userID, State: state, Data: data}
	}
}
