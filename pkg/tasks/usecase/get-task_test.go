package tasks

// func TestGetUserUsecase(t *testing.T) {
// 	var mockUser models.User
// 	gofakeit.Struct(&mockUser)

// 	mockUserRepo := new(mocks.UserRepository)
// 	mockUserUsecase := new(mocks.UserUsecase)

// 	t.Run("success", func(t *testing.T) {
// 		mockUserRepo.On("GetUser", mock.AnythingOfType("string")).Return(mockUser, nil).Once()

// 		usecase := users.NewUserUsecase(mockUserRepo)

// 		result, err := usecase.GetUser(mockUser.ID)

// 		assert.NoError(t, err)
// 		assert.Equal(t, mockUser, result)

// 		mockUserUsecase.AssertExpectations(t)
// 	})
// 	t.Run("error-failed", func(t *testing.T) {
// 		mockUserRepo.On("GetUser", mock.AnythingOfType("string")).Return(models.User{}, errors.New("call error")).Once()

// 		usecase := users.NewUserUsecase(mockUserRepo)

// 		result, err := usecase.GetUser(mockUser.ID)

// 		assert.Error(t, err)
// 		assert.Equal(t, models.User{}, result)

// 		mockUserUsecase.AssertExpectations(t)
// 	})

// }
