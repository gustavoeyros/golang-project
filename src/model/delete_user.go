package model

import "github.com/gustavoeyros/golang-project/src/configurations/rest_err"

func (*UserDomain) DeleteUser(string) *rest_err.RestErr {
	return nil
}
