package accounts

type Service struct {

}

func NewService() *Service {
	return &Service{}
}

func (Service) Process(command ICommand) int64 {
	return command.Process()
}
