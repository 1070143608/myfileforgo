package house

type Friend struct {
	Where string
	Sex int
	Age int
}

type FriendOptions func(friend *Friend)

func WithWhere(where string) FriendOptions {
	return func(friend *Friend) {
		friend.Where = where
	}
}

func WithSex(sex int) FriendOptions {
	return func(friend *Friend) {
		friend.Sex = sex
	}
}

func WithAge(age int) FriendOptions {
	return func(friend *Friend) {
		friend.Age = age
	}
}

func FindFriend(ops ...FriendOptions) *Friend {
	friend := &Friend{}
	for _, op := range ops {
		op(friend)
	}
	return friend
}