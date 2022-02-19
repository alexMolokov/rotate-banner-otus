package approtator_test

import (
	"context"
	"errors"
	"testing"

	"github.com/alexMolokov/rotate-banner-otus/internal/app/rotator"
	"github.com/alexMolokov/rotate-banner-otus/internal/app/rotator/mock"
	rotatorstorage "github.com/alexMolokov/rotate-banner-otus/internal/storage/rotator"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type AppSuite struct {
	suite.Suite
	mockCtl     *gomock.Controller
	mockStorage *mock.MockStorage
	mockLogger  *mock.MockLogger
	app         *approtator.App
}

func (s *AppSuite) SetupTest() {
	s.mockCtl = gomock.NewController(s.T())
	s.mockStorage = mock.NewMockStorage(s.mockCtl)
	s.mockLogger = mock.NewMockLogger(s.mockCtl)
	s.app = approtator.NewAppRotator(s.mockLogger, s.mockStorage)
}

func (s *AppSuite) TearDownTest() {
	s.mockCtl.Finish()
}

func (s *AppSuite) Test_AddBannerToSlot() {
	bannerID, slotID := int64(1), int64(1)
	s.mockStorage.EXPECT().AddBannerToSlot(gomock.Any(), bannerID, slotID).Times(1).Return(nil)
	err := s.app.AddBannerToSlot(context.Background(), bannerID, slotID)
	s.Require().NoError(err)

	errAdd := errors.New("can't add")
	s.mockStorage.EXPECT().AddBannerToSlot(gomock.Any(), bannerID, slotID).Times(1).Return(errAdd)
	err = s.app.AddBannerToSlot(context.Background(), bannerID, slotID)
	s.Require().ErrorIs(err, errAdd)
	s.Require().Equal(errAdd.Error(), err.Error())
}

func (s *AppSuite) Test_RemoveBannerFromSlot() {
	bannerID, slotID := int64(1), int64(1)
	s.mockStorage.EXPECT().RemoveBannerFromSlot(gomock.Any(), bannerID, slotID).Times(1).Return(nil)
	err := s.app.RemoveBannerFromSlot(context.Background(), bannerID, slotID)
	s.Require().NoError(err)

	errRemove := errors.New("can't remove")
	s.mockStorage.EXPECT().RemoveBannerFromSlot(gomock.Any(), bannerID, slotID).Times(1).Return(errRemove)
	err = s.app.RemoveBannerFromSlot(context.Background(), bannerID, slotID)
	s.Require().ErrorIs(err, errRemove)
	s.Require().Equal(errRemove.Error(), err.Error())
}

func (s *AppSuite) Test_CountTransition() {
	bannerID, slotID, sgID := int64(1), int64(1), int64(1)
	s.mockStorage.EXPECT().CountTransition(gomock.Any(), bannerID, slotID, sgID).Times(1).Return(nil)
	err := s.app.CountTransition(context.Background(), bannerID, slotID, sgID)
	s.Require().NoError(err)

	errTransition := errors.New("can't count transition")
	s.mockStorage.EXPECT().CountTransition(gomock.Any(), bannerID, slotID, sgID).Times(1).Return(errTransition)
	err = s.app.CountTransition(context.Background(), bannerID, slotID, sgID)
	s.Require().ErrorIs(err, errTransition)
	s.Require().Equal(errTransition.Error(), err.Error())
}

func (s *AppSuite) Test_ChooseBanner() {
	slotID, sgID := int64(1), int64(1)

	totalDisplay := 100
	bannerChooseID := int64(3)
	bannerStatChoose := []rotatorstorage.BannerStat{
		{ID: 1, Display: 40, Click: 10},
		{ID: 2, Display: 30, Click: 10},
		{ID: bannerChooseID, Display: 30, Click: 20},
	}

	s.mockStorage.EXPECT().CountDisplay(gomock.Any(), gomock.Any(), slotID, sgID).Times(1).Return(nil)
	s.mockStorage.EXPECT().GetBannersStat(gomock.Any(), slotID, sgID).Times(1).Return(bannerStatChoose, totalDisplay, nil)
	bannerID, err := s.app.ChooseBanner(context.Background(), slotID, sgID)
	s.Require().NoError(err)
	s.Equal(bannerChooseID, bannerID)

	errChoose := errors.New("can't get stat")
	bannerStat := make([]rotatorstorage.BannerStat, 0)
	s.mockStorage.EXPECT().GetBannersStat(gomock.Any(), slotID, sgID).Times(1).Return(bannerStat, 1, errChoose)
	_, err = s.app.ChooseBanner(context.Background(), slotID, sgID)
	s.Require().ErrorIs(err, errChoose)
	s.Require().Equal(errChoose.Error(), err.Error())
}

func TestAppSuite(t *testing.T) {
	suite.Run(t, new(AppSuite))
}
