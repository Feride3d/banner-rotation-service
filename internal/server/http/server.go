package server

import (
	"context"
	"errors"
	"net"

	"github.com/Feride3d/banner-rotation-service/internal/pb"
	"github.com/Feride3d/banner-rotation-service/internal/queues"
	"github.com/Feride3d/banner-rotation-service/internal/service"
	"github.com/Feride3d/banner-rotation-service/internal/storage/models"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// GRPCServer rotates banners.
type GRPCServer struct {
	pb.UnimplementedRotatorServer
	domain                string
	bannerRotationService service.BannerRotationService
	publisher             queues.Publisher
	logger                *logrus.Logger
}

// NewGRPCServer returns server.
func NewGRPCServer(
	domain string,
	BannerRotationService service.BannerRotationService,
	publisher queues.Publisher,
	logger *logrus.Logger,
) *GRPCServer {
	return &GRPCServer{
		domain:                domain,
		bannerRotationService: BannerRotationService,
		publisher:             publisher,
		logger:                logger,
	}
}

// Start GRPCServer.
func (s *GRPCServer) Start() error {
	serv := grpc.NewServer()
	reflection.Register(serv)

	l, err := net.Listen("tcp", s.domain)
	if err != nil {
		return err
	}

	pb.RegisterRotatorServer(serv, s)

	return serv.Serve(l)
}

// AddBanner adds the banner to the slot.
func (s *GRPCServer) AddBanner(ctx context.Context, req *pb.AddBannerRequest) (*pb.AddBannerResponse, error) {
	if ctx.Err() != nil {
		return nil, errors.New("Request is cancelled.")
	}

	if req.BannerId == 0 || req.SlotId == 0 {
		return nil, errors.New("Process request: no banners or slots to add.")
	}

	rotator := models.Rotation{
		BannerID: int(req.GetBannerId()),
		SlotID:   int(req.GetSlotId()),
	}

	newRotator, err := s.bannerRotationService.AddBanner(ctx, rotator)
	if err != nil {
		return nil, errors.New("AddBanner request is failed.")
	}

	return &pb.AddBannerResponse{
		Id:       int32(newRotator.ID),
		BannerId: int32(newRotator.BannerID),
		SlotId:   int32(newRotator.SlotID),
	}, nil
}

// DeleteBanner deletes the banner from the slot.
func (s *GRPCServer) DeleteBanner(ctx context.Context, req *pb.DeleteBannerRequest) (*pb.DeleteBannerResponse, error) {
	if ctx.Err() != nil {
		return nil, errors.New("Request is cancelled.")
	}

	if req.BannerId == 0 || req.SlotId == 0 {
		return nil, errors.New("Process request: no banners to delete.")
	}

	bannerID := int(req.GetBannerId())

	err := s.bannerRotationService.DeleteBanner(ctx, bannerID)
	if err != nil {
		return nil, errors.New("AddBanner request is failed.")
	}

	return &pb.DeleteBannerResponse{
		Id: int32(bannerID),
	}, nil
}

// AddClick adds a click on the banner.
func (s *GRPCServer) AddClick(ctx context.Context, req *pb.AddClickRequest) (*pb.AddClickResponse, error) {
	if ctx.Err() != nil {
		return nil, errors.New("Request is cancelled.")
	}

	if req.BannerId == 0 || req.SlotId == 0 || req.GroupId == 0 {
		return nil, errors.New("No banners or slots or groups to click.")
	}

	bannerID := int(req.GetBannerId())
	groupID := int(req.GetGroupId())

	rotation, err := s.bannerRotationService.RotationStorage.GetBanner(ctx, bannerID)
	if err != nil {
		return nil, errors.New("AddClick request is cancelled.")
	}

	events, err := s.bannerRotationService.AddClick(ctx, *rotation, groupID)
	if err != nil {
		return nil, err
	}
	err = s.publisher.Publish(ctx, *events)
	if err != nil {
		s.logger.Error("Process request: publish message to queue.")
	}

	return &pb.AddClickResponse{
		Id: int32(bannerID),
	}, nil
}

// AddBannerDisplay adds a banner to display.
func (s *GRPCServer) AddBannerDisplay(ctx context.Context, req *pb.AddBannerDisplayRequest) (*pb.AddBannerDisplayResponse, error) {
	if ctx.Err() != nil {
		return nil, errors.New("Request is cancelled.")
	}

	if req.SlotId == 0 || req.GroupId == 0 {
		return nil, errors.New("Process request: no slot or group.")
	}

	slotID := int(req.GetSlotId())
	groupID := int(req.GetGroupId())

	bannerID, events, err := s.bannerRotationService.AddBannerDisplay(ctx, slotID, groupID)
	if err != nil {
		return nil, err
	}

	err = s.publisher.Publish(ctx, *events)
	if err != nil {
		s.logger.Error("Process request: send message to queue error.")
	}
	return &pb.AddBannerDisplayResponse{
		Id: int32(bannerID),
	}, nil
}

/* Привести в соответствие сервису
// CreateBanner creates a new banner.
func (s *GRPCServer) CreateBanner(ctx context.Context, req *pb.CreateBannerRequest) (*pb.CreateBannerResponse, error) {
	if ctx.Err() != nil {
		return nil, errors.New("Request is cancelled.")
	}

	ID := req.Id

	if ID == "" {
		ID = uuid.NewString()
	}

	ID, err := s.bannerRotationService.CreateBanner(ctx, ID, req.Description)
	if err != nil {
		return nil, errors.New("CreateBanner request is failed.")
	}

	return &pb.CreateBannerResponse{
		Id: ID,
	}, nil
}

// CreateSlot creates a new slot.
func (s *GRPCServer) CreateSlot(ctx context.Context, req *pb.CreateSlotRequest) (*pb.CreateSlotResponse, error) {
	if ctx.Err() != nil {
		return nil, errors.New("Request is cancelled.")
	}

	ID := req.Id

	if ID == "" {
		ID = uuid.NewString()
	}

	ID, err := s.bannerRotationService.CreateSlot(ctx, ID, req.Description)
	if err != nil {
		return nil, errors.New("CreateSlot request is failed.")
	}

	return &pb.CreateSlotResponse{
		Id: ID,
	}, nil
}

// CreateGroup creates a new group.
func (s *GRPCServer) CreateGroup(ctx context.Context, req *pb.CreateGroupRequest) (*pb.CreateGroupResponse, error) {
	if ctx.Err() != nil {
		return nil, errors.New("Request is cancelled.")
	}

	ID := req.Id

	if ID == "" {
		ID = uuid.NewString()
	}

	ID, err := s.bannerRotationService.CreateGroup(ctx, ID, req.Description)
	if err != nil {
		return nil, errors.New("CreateGroup request is failed.")
	}

	return &pb.CreateGroupResponse{
		Id: ID,
	}, nil
}
*/
