package internalgrpc

import (
	"context"
	"errors"
	"fmt"
	"net"
	"testing"
	"time"

	"github.com/alexMolokov/rotate-banner-otus/api/pb"
	"github.com/alexMolokov/rotate-banner-otus/internal/server/grpc/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GRPCServerSuite struct {
	suite.Suite
	mockCtl *gomock.Controller
	conn    *grpc.ClientConn
	client  pb.RotatorClient
	app     *mock.MockApplication
}

func (s *GRPCServerSuite) setLoggerStub(ml *mock.MockLogger) {
	ml.EXPECT().Info(gomock.Any(), gomock.Any()).Do(func(_ string, _ ...interface{}) {
	}).AnyTimes()
	ml.EXPECT().Error(gomock.Any(), gomock.Any()).Do(func(_ string, _ ...interface{}) {
	}).AnyTimes()
	ml.EXPECT().Debug(gomock.Any(), gomock.Any()).Do(func(_ string, _ ...interface{}) {
	}).AnyTimes()
	ml.EXPECT().Warning(gomock.Any(), gomock.Any()).Do(func(_ string, _ ...interface{}) {
	}).AnyTimes()
}

func (s *GRPCServerSuite) SetupSuite() {
	s.mockCtl = gomock.NewController(s.T())
	ml := mock.NewMockLogger(s.mockCtl)
	s.setLoggerStub(ml)

	s.app = mock.NewMockApplication(s.mockCtl)

	bufSize := 1024 * 1024
	lsn := bufconn.Listen(bufSize)
	grpcS := NewServer(ml, s.app, "")
	pb.RegisterRotatorServer(grpcS.grpcServer, grpcS.service)
	go func() {
		if err := grpcS.grpcServer.Serve(lsn); err != nil {
			fmt.Printf("Server exited with error: %v", err)
		}
	}()
	bufDialer := func(context.Context, string) (net.Conn, error) {
		return lsn.Dial()
	}

	ctx := context.Background()
	conn, _ := grpc.DialContext(
		ctx,
		"bufnet",
		grpc.WithContextDialer(bufDialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	s.conn = conn
	s.client = pb.NewRotatorClient(conn)
}

func (s *GRPCServerSuite) Test_AddBannerToSlot() { //nolint
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	bannerID, slotID := int64(1), int64(1)
	errAddBanner := errors.New("can't add banner")
	s.app.EXPECT().AddBannerToSlot(gomock.Any(), bannerID, slotID).Return(errAddBanner)

	result, err := s.client.AddBannerToSlot(ctx, &pb.AddBannerToSlotRequest{BannerId: bannerID, SlotId: slotID})
	s.Require().Nil(result)
	s.Require().Error(err)

	s.app.EXPECT().AddBannerToSlot(gomock.Any(), bannerID, slotID).Return(nil)
	result, err = s.client.AddBannerToSlot(ctx, &pb.AddBannerToSlotRequest{BannerId: bannerID, SlotId: slotID})
	s.Require().IsType(&emptypb.Empty{}, result)
	s.Require().Nil(err)

	// не указан обязательный параметр bannerId
	_, err = s.client.AddBannerToSlot(ctx, &pb.AddBannerToSlotRequest{SlotId: slotID})
	s.Require().NotNil(err)

	// не указан обязательный параметр slotId
	_, err = s.client.AddBannerToSlot(ctx, &pb.AddBannerToSlotRequest{BannerId: bannerID})
	s.Require().NotNil(err)
}

func (s *GRPCServerSuite) Test_RemoveBannerFromSlot() { //nolint
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	bannerID, slotID := int64(1), int64(1)
	errRemoveBanner := errors.New("can't remove banner")
	s.app.EXPECT().RemoveBannerFromSlot(gomock.Any(), bannerID, slotID).Return(errRemoveBanner)

	result, err := s.client.RemoveBannerFromSlot(ctx, &pb.RemoveBannerFromSlotRequest{BannerId: bannerID, SlotId: slotID})
	s.Require().Nil(result)
	s.Require().Error(err)

	s.app.EXPECT().RemoveBannerFromSlot(gomock.Any(), bannerID, slotID).Return(nil)
	result, err = s.client.RemoveBannerFromSlot(ctx, &pb.RemoveBannerFromSlotRequest{BannerId: bannerID, SlotId: slotID})
	s.Require().IsType(&emptypb.Empty{}, result)
	s.Require().Nil(err)

	// не указан обязательный параметр slotId
	_, err = s.client.RemoveBannerFromSlot(ctx, &pb.RemoveBannerFromSlotRequest{BannerId: bannerID})
	s.Require().NotNil(err)

	// не указан обязательный параметр bannerId
	_, err = s.client.RemoveBannerFromSlot(ctx, &pb.RemoveBannerFromSlotRequest{SlotId: slotID})
	s.Require().NotNil(err)
}

func (s *GRPCServerSuite) Test_CountTransition() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	bannerID, slotID, sgID := int64(1), int64(1), int64(1)
	errCountTransition := errors.New("can't count transition")
	s.app.EXPECT().CountTransition(gomock.Any(), bannerID, slotID, sgID).Return(errCountTransition)

	result, err := s.client.CountTransition(ctx,
		&pb.CountTransitionRequest{BannerId: bannerID, SlotId: slotID, SgId: sgID},
	)
	s.Require().Nil(result)
	s.Require().Error(err)

	s.app.EXPECT().CountTransition(gomock.Any(), bannerID, slotID, sgID).Return(nil)
	result, err = s.client.CountTransition(ctx,
		&pb.CountTransitionRequest{BannerId: bannerID, SlotId: slotID, SgId: sgID},
	)
	s.Require().IsType(&emptypb.Empty{}, result)
	s.Require().Nil(err)

	// не указан обязательный параметр bannerId
	_, err = s.client.CountTransition(ctx, &pb.CountTransitionRequest{SlotId: slotID, SgId: sgID})
	s.Require().NotNil(err)

	// не указан обязательный параметр slotId
	_, err = s.client.CountTransition(ctx, &pb.CountTransitionRequest{BannerId: bannerID, SgId: sgID})
	s.Require().NotNil(err)

	// не указан обязательный параметр sgId
	_, err = s.client.CountTransition(ctx, &pb.CountTransitionRequest{BannerId: bannerID, SlotId: slotID})
	s.Require().NotNil(err)

	// указаны некорректные значения
	_, err = s.client.CountTransition(ctx, &pb.CountTransitionRequest{BannerId: 0, SlotId: 0, SgId: 0})
	s.Require().NotNil(err)
}

func (s *GRPCServerSuite) Test_ChooseBanner() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	bannerID, slotID, sgID := int64(13), int64(1), int64(1)
	errChooseBanner := errors.New("can't choose banner")
	s.app.EXPECT().ChooseBanner(gomock.Any(), slotID, sgID).Return(int64(0), errChooseBanner)

	result, err := s.client.ChooseBanner(ctx, &pb.ChooseBannerRequest{SlotId: slotID, SgId: sgID})
	s.Require().Nil(result)
	s.Require().Error(err)

	s.app.EXPECT().ChooseBanner(gomock.Any(), slotID, sgID).Return(bannerID, nil)
	result, err = s.client.ChooseBanner(ctx, &pb.ChooseBannerRequest{SlotId: slotID, SgId: sgID})
	s.Require().IsType(&pb.ChooseBannerResponse{}, result)
	s.Require().Nil(err)
	s.Require().Equal(bannerID, result.BannerId)

	// не указан обязательный параметр SgId
	_, err = s.client.ChooseBanner(ctx, &pb.ChooseBannerRequest{SlotId: slotID})
	s.Require().NotNil(err)

	// не указан обязательный параметр slotId
	_, err = s.client.ChooseBanner(ctx, &pb.ChooseBannerRequest{SgId: sgID})
	s.Require().NotNil(err)

	// указаны некорректные значения
	_, err = s.client.ChooseBanner(ctx, &pb.ChooseBannerRequest{SgId: 0, SlotId: 0})
	s.Require().NotNil(err)
}

func (s *GRPCServerSuite) TearDownSuite() {
	s.mockCtl.Finish()
}

func TestGRPCServerSuite(t *testing.T) {
	suite.Run(t, new(GRPCServerSuite))
}
