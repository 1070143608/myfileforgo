package house

import "github.com/bigwhite/functrace"

type Friend struct {
	Where string
	Sex   int
	Age   int
}

type FriendOptions func(friend *Friend)

func WithWhere(where string) FriendOptions {
	defer functrace.Trace()()
	return func(friend *Friend) {
		friend.Where = where
	}
}

func WithSex(sex int) FriendOptions {
	defer functrace.Trace()()
	return func(friend *Friend) {
		friend.Sex = sex
	}
}

func WithAge(age int) FriendOptions {
	defer functrace.Trace()()
	return func(friend *Friend) {
		friend.Age = age
	}
}

func FindFriend(ops ...FriendOptions) *Friend {
	defer functrace.Trace()()
	friend := &Friend{}
	for _, op := range ops {
		op(friend)
	}
	return friend
}
