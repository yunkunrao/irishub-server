package services

import (
	"testing"

	"github.com/irisnet/irishub-server/modules/logger"
	"github.com/irisnet/irishub-server/rpc/vo"
	"github.com/irisnet/irishub-server/utils/helper"
)

func TestValidatorService_List(t *testing.T) {
	type args struct {
		listVo vo.ValidatorListReqVO
	}
	tests := []struct {
		name string
		s    ValidatorService
		args args
	}{
		{
			name: "test get validator list",
			s:    ValidatorService{},
			args: args{
				listVo: vo.ValidatorListReqVO{
					Address: "461B0D58301072D68EB95C54DDFCFBFF7D67DA7C",
					Page:    1,
					PerPage: 50,
					Sort:    "-voting_power",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := ValidatorService{}
			res, err := s.List(tt.args.listVo)
			if err.IsNotNull() {
				logger.Error.Fatalln(err)
			}
			logger.Info.Println(helper.ToJson(res))
		})
	}
}

func TestValidatorService_GetValidatorExRate(t *testing.T) {
	type args struct {
		reqVO vo.ValidatorExRateReqVO
	}
	tests := []struct {
		name  string
		args  args
	}{
		{
			name: "test get validatot ex_rate",
			args: args{
				reqVO: vo.ValidatorExRateReqVO{
					ValidatorAddress: "46834815791157A845A3FA5912A695D30F56CF4D",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := ValidatorService{}
			res, err := s.GetValidatorExRate(tt.args.reqVO)
			if err.IsNotNull() {
				logger.Error.Fatalln(err)
			}
			logger.Info.Println(helper.ToJson(res))
		})
	}
}

func TestValidatorService_GetTotalShares(t *testing.T) {
	tests := []struct {
		name    string
	}{
		{
			name: "test get validator total shares",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := ValidatorService{}
			res, err := s.GetTotalShares()
			if err != nil {
				logger.Error.Fatalln(err)
			}
			logger.Info.Println(helper.ToJson(res))
		})
	}
}