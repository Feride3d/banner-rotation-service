// +build integration

package integration_test

import (
	"context"
	"os"
	"testing"

	"github.com/Feride3d/banner-rotation-service/internal/pb"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
)

type RotatorSuite struct {
	suite.Suite
	ctx           context.Context
	rotatorConn   *grpc.ClientConn // устанавливаем соединение для дальнейшего переиспользования
	rotatorClient pb.RotatorClient
}

func (s *RotatorSuite) SetupSuite() {
	rotatorHost := os.Getenv("ROTATOR_SERVER_HOST")
	if rotatorHost == "" { // если запускаем в локальной сети, а не в CI-CD
		rotatorHost = "127.0.0.1:9001"
	}

	rotatorConn, err := grpc.Dial(rotatorHost, grpc.WithInsecure())
	s.Require().NoError(err) // требуется, чтобы не было ошибки при подключении

	s.ctx = context.Background()
	s.rotatorClient = pb.NewRotatorClient(rotatorConn)
	s.rotatorClient.CreateBanner
}

// test CreateBanner method
func (s *RotatorSuite) TestCreateBanner() {
	req := &pb.CreateBannerRequest{
//		BannerId: "banner_id",
	}
	_, err := s.rotatorClient.CreateBanner(s.ctx, req)
	s.Require().NoError(err)

	s.Require().Equal(pb.BannerId, resp.GetBannerId() // проверить появилось ли что-то в бд
}

// test AddBanner method
func (s *RotatorSuite) TestAddBanner() {
	
	bannerID := resp.GetBannerId()
	slotID := resp.GetSlotId()
	
	sendReq := &pb.AddBannerRequest{
		BannerId: bannerID,
		SlotId: slotID,
	}
	resp, err := s.rotatorClient.AddBanner(s.ctx, req)
	s.Require().NoError(err)
	
	s.Require().Equal([]string{""}, resp.GetBannerId()) // мб проверить через слот в бд 
}



func TestRotatorSuite(t *testing.T) {
	suite.Run(t, new(RotatorSuite))
}

