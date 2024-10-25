package models

type UserState struct {
	ID    int64
	State string
	Data  map[string]interface{}
}

var userStates = make(map[int64]*UserState)

func getUserState(userID int64) *UserState {
	if state, exists := userStates[userID]; exists {
		return state
	}
	return &UserState{ID: userID, State: "initial", Data: make(map[string]interface{})}
}
