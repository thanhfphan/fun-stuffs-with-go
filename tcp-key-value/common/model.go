package common

// Put
type PutArgs struct {
	Key   string
	Value string
}

type PutReply struct{}

// Get
type GetArgs struct {
	Key string
}

type GetReply struct {
	Value string
}
