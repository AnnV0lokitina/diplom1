package service

//func TestAddCredentials(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	repo := mock.NewMockRepo(ctrl)
//	conn := mock.NewMockExtConnection(ctrl)
//	s := mock.NewMockSession(ctrl)
//	ctx := context.Background()
//
//	service := NewService(repo, conn, s)
//
//	getSessionOdd := false
//	s.EXPECT().Get().DoAndReturn(func() error {
//		if getSessionOdd {
//			getSessionOdd = false
//			return nil
//		}
//		getSessionOdd = true
//		return errors.New("error")
//	}).AnyTimes()
//
//	getInfoOdd := false
//	repo.EXPECT().GetInfo().DoAndReturn(func() error {
//		if getInfoOdd {
//			getInfoOdd = false
//			return errors.New("error")
//		}
//		getInfoOdd = true
//		return nil
//	}).AnyTimes()
//
//	//repo.EXPECT().GetInfo().DoAndReturn(func() error {
//	//	if getInfoOdd {
//	//		getInfoOdd = false
//	//		return errors.New("error")
//	//	}
//	//	getInfoOdd = true
//	//	return nil
//	//}).AnyTimes()
//
//	err := service.AddCredentials(ctx, "login", "password", "meta")
//
//}
