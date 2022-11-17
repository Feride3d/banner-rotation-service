package internalgrpc

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"

	pb "github.com/Feride3d/banner-rotation-service/internal/pb/api/rotator"
	"github.com/Feride3d/banner-rotation-service/internal/service"
	"github.com/google/uuid"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type HttpServer struct {
	service    service.Service
	httpserver *http.Server
}

type grpcServer struct {
	pb.UnimplementedRotatorServer
	service service.Service
}

func NewHttpServer(service *service.Service, addr string, port string, grpcPort string) (*HttpServer, error) {
	grpcEndpoint := net.JoinHostPort(addr, grpcPort)

	l, err := net.Listen("tcp", grpcEndpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to start server %w", err)
	}

	log := service.LoggerBegin().Entry()

	srv := grpc.NewServer(grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
		grpc_logrus.StreamServerInterceptor(log),
	)),

		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_logrus.UnaryServerInterceptor(log),
		)))

	pb.RegisterRotatorServer(srv, &grpcServer{service: *service})

	go func() {
		err := srv.Serve(l)
		if err != nil {
			service.LoggerBegin().Error(fmt.Errorf("failed to serve grpc server, %w", err).Error())
		}
	}()

	ctx, cncl := context.WithTimeout(context.Background(), 10*time.Second)
	defer cncl()

	conn, err := grpc.DialContext(
		ctx,
		grpcEndpoint,
		grpc.WithBlock(),
	)

	if err != nil {
		return nil, fmt.Errorf("failed to start grpc server, %w", err)
	}

	// Register service's handler
	mux := runtime.NewServeMux()
	err = pb.RegisterRotatorHandler(ctx, mux, conn)
	if err != nil {
		return nil, fmt.Errorf("failed to register endpoints, %w", err)
	}

	httpserver := &http.Server{
		Addr:         net.JoinHostPort(addr, port),
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	return &HttpServer{*service, httpserver}, nil
}

func (s *HttpServer) Start(ctx context.Context) error {
	err := s.httpserver.ListenAndServe()
	if err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}

		return fmt.Errorf("failed to start http server, %w", err)
	}

	return nil
}

func (s *HttpServer) Stop(ctx context.Context) error {
	if err := s.httpserver.Shutdown(ctx); err != nil {
		return fmt.Errorf("failed to stop http server, %w", err)
	}
	return nil
}

// CreateBanner creates a new banner.
func (s *grpcServer) CreateBanner(ctx context.Context, req *pb.CreateBannerRequest) (*pb.CreateBannerResponse, error) {
	if ctx.Err() != nil {
		return nil, errors.New("CreateBanner request is cancelled")
	}

	bannerID := req.Id

	if bannerID == "" {
		bannerID = uuid.NewString()
	}

	bannerID, err := s.service.CreateBanner(ctx, bannerID, req.Description)
	if err != nil {
		return nil, errors.New("CreateBanner request is failed")
	}

	return &pb.CreateBannerResponse{
		Id: bannerID}, nil
}

// CreateSlot creates a new slot.
func (s *grpcServer) CreateSlot(ctx context.Context, req *pb.CreateSlotRequest) (*pb.CreateSlotResponse, error) {
	if ctx.Err() != nil {
		return nil, errors.New("CreateSlot request is cancelled")
	}

	slotID := req.Id

	if slotID == "" {
		slotID = uuid.NewString()
	}

	slotID, err := s.service.CreateSlot(ctx, slotID, req.Description)
	if err != nil {
		return nil, errors.New("CreateSlot request is failed")
	}

	return &pb.CreateSlotResponse{
		Id: slotID}, nil
}

// CreateGroup creates a new group.
func (s *grpcServer) CreateGroup(ctx context.Context, req *pb.CreateGroupRequest) (*pb.CreateGroupResponse, error) {
	if ctx.Err() != nil {
		return nil, errors.New("CreateGroup request is cancelled")
	}

	groupID := req.Id

	if groupID == "" {
		groupID = uuid.NewString()
	}

	groupID, err := s.service.CreateGroup(ctx, groupID, req.Description)
	if err != nil {
		return nil, errors.New("CreateGroup request is failed")
	}

	return &pb.CreateGroupResponse{
		Id: groupID}, nil
}

// AddBanner adds the banner to the slot.
func (s *grpcServer) AddBanner(ctx context.Context, req *pb.AddBannerRequest) (*pb.AddBannerResponse, error) {
	if ctx.Err() != nil {
		return nil, errors.New("addBanner request is cancelled")
	}

	if req.BannerId == "" || req.SlotId == "" {
		return nil, errors.New("process request: no banners or slots to add")
	}

	err := s.service.AddBanner(ctx, req.BannerId, req.SlotId)
	if err != nil {
		return nil, errors.New("addBanner request is failed")
	}

	return &pb.AddBannerResponse{
		Status: "ok, banner was added to slot"}, nil
}

// DeleteBanner deletes the banner from the slot.
func (s *grpcServer) DeleteBanner(ctx context.Context, req *pb.DeleteBannerRequest) (*pb.DeleteBannerResponse, error) {
	if ctx.Err() != nil {
		return nil, errors.New("deleteBanner request is cancelled")
	}

	if req.BannerId == "" || req.SlotId == "" {
		return nil, errors.New("process request: no banners to delete")
	}

	err := s.service.DeleteBanner(ctx, req.BannerId, req.SlotId)
	if err != nil {
		return nil, errors.New("deleteBanner request is failed")
	}

	return &pb.DeleteBannerResponse{
		Status: "ok, banner was deleted"}, nil
}

// AddClick adds a click on the banner.
func (s *grpcServer) AddClick(ctx context.Context, req *pb.AddClickRequest) (*pb.AddClickResponse, error) {
	if ctx.Err() != nil {
		return nil, errors.New("AddClick request is cancelled")
	}

	if req.BannerId == "" || req.SlotId == "" || req.GroupId == "" {
		return nil, errors.New("no banners or slots or groups or group to click")
	}

	err := s.service.AddClick(ctx, req.BannerId, req.SlotId, req.GroupId)
	if err != nil {
		return nil, errors.New("addClick request is failed")
	}

	return &pb.AddClickResponse{
		Status: "ok, click on banner was added"}, nil
}

// AddBannerDisplay adds a banner to display.
func (s *grpcServer) AddBannerDisplay(ctx context.Context, req *pb.AddBannerDisplayRequest) (*pb.AddBannerDisplayResponse, error) {
	if ctx.Err() != nil {
		return nil, errors.New("AddBannerDisplay request is cancelled")
	}

	if req.SlotId == "" || req.GroupId == "" {
		return nil, errors.New("no slot or group for banner's display")
	}

	bannerID, err := s.service.AddBannerDisplay(ctx, req.SlotId, req.GroupId)
	if err != nil {
		return nil, errors.New("addBannerDisplay request is failed")
	}

	return &pb.AddBannerDisplayResponse{
		Id: bannerID}, nil
}
