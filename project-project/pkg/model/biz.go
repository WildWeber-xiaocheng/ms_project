package model

const (
	Normal         = 1
	Personal int32 = 1
	AESKey         = "abcdefgehjhijkmlkjjwwoew"
)

const (
	NoDeleted = iota
	Deleted
)

const (
	NoArchive = iota
	Archived
)

const (
	Open = iota
	Private
	Custom
)

const (
	Default = "default"
	Simple  = "simple"
)

const (
	NoCollected = iota
	Collected
)

const (
	NoOwner = iota
	Owner
)

const (
	NoExecutor = iota
	Executor
)

const (
	NoCanRead = iota
	CanRead
)
