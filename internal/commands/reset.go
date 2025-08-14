package commands

import (
	"codingiam/gator/internal/state"
	"context"
)

func handlerReset(st *state.State, _ command) error {
	return st.Db.TruncateUsers(context.Background())
}
