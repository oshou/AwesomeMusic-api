//package usecase
//
//import (
// "fmt"
// "log"
//
// "github.com/pkg/errors"
// "go.uber.org/zap"
//
// "github.com/oshou/AwesomeMusic-api/api/domain/model"
// "github.com/oshou/AwesomeMusic-api/api/domain/repository"
// "github.com/oshou/AwesomeMusic-api/log"
//)
//
//type IAuthUsecase interface {
// GetByAuthToken(authToken string) (*model.User, error)
// LoggedInAuth(authToken string) (int, string, error)
//}
//
//type authUsecase struct {
// repo repository.IAuthRepository
//}
//
//var \_ IAuthUsecase = &authUsecase{}
//
//func NewAuthUsecase(repo repository.IAuthRepository) IAuthUsecase {
// return &authUsecase{
// repo: repo,
// }
//}
//
//func (au *authUsecase) GetByAuthToken(authToken string) (*model.User, error) {
// user, err := au.repo.GetByAuthToken(authToken)
// if err != nil {
// e := errors.Cause(err)
// switch e.(type) {
// case *repository.NoRowsError:
// return nil, NotFoundError{}
// default:
// log.Logger.Error("failed to get user by authToken.", zap.Error(err))
// return nil, InternalServerError{}
// }
// }
// return user, nil
//}
//
//func (au \*authUsecase) LoggedInAuth(authToken string) (int, string, error) {
// user, err := au.GetByAuthToken(authToken)
// if err != nil {
// return 0, "", err
// }
// fmt.Println(user)
// return 0, "", nil
//}
//
