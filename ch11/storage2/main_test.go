package storage2

import (
	"strings"
	"testing"
)

func TestCheckQuota(t *testing.T) {

	saved := notifyUser
	defer func() { notifyUser = saved }()

	var notifiedUser, notifiedMsg string
	notifyUser = func(username, msg string) {
		notifiedUser, notifiedMsg = username, msg
	}

	const user = "tm@example.com"
	usage[user] = 990
	CheckQuota(user)

	if notifiedUser != user {
		t.Errorf("wrong user (%s) notified, want (%s)", notifiedUser, user)
	}

	const wantSubstring = "99% of your quota"

	if !strings.Contains(notifiedMsg, wantSubstring) {
		t.Errorf("wrong message (%s) notified, want (%s) in message", notifiedMsg, wantSubstring)
	}

}
