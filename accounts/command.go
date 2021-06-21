package accounts

type ICommand interface {
	Process() int64
}
