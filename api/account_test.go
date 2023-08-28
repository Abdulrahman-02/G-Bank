package api

// import (
// 	"bytes"
// 	"database/sql"
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"
// 	"time"

// 	mockdb "github.com/abdulrahman-02/G-Bank/db/mock"
// 	db "github.com/abdulrahman-02/G-Bank/db/sqlc"
// 	"github.com/abdulrahman-02/G-Bank/token"
// 	"github.com/abdulrahman-02/G-Bank/util"
// 	"github.com/stretchr/testify/require"
// 	"go.uber.org/mock/gomock"
// )

// func TestGetAccountAPI(t *testing.T) {
// 	randomUser, _ := randomUser(t)
// 	account := randomAccount(randomUser.Username)

// 	testCases := []struct {
// 		name          string
// 		accountID     int64
// 		setupAuth     func(t *testing.T, request *http.Request, token token.Maker)
// 		buildStubs    func(store *mockdb.MockStore)
// 		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
// 	}{
// 		{
// 			name:      "OK",
// 			accountID: account.ID,
// 			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
// 				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, "users", time.Minute)
// 			},
// 			buildStubs: func(store *mockdb.MockStore) {
// 				store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account.ID)).Times(1).Return(account, nil)
// 			},
// 			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusOK, recorder.Code)
// 				requireBodyMatchAccount(t, recorder.Body, account)
// 			},
// 		},
// 		{
// 			name:      "NotFound",
// 			accountID: account.ID,
// 			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
// 				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, "users", time.Minute)
// 			},
// 			buildStubs: func(store *mockdb.MockStore) {
// 				store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account.ID)).Times(1).Return(db.Accounts{}, sql.ErrNoRows)
// 			},
// 			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusNotFound, recorder.Code)
// 			},
// 		},
// 		{
// 			name:      "InternalError",
// 			accountID: account.ID,
// 			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
// 				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, "users", time.Minute)
// 			},
// 			buildStubs: func(store *mockdb.MockStore) {
// 				store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account.ID)).Times(1).Return(db.Accounts{}, sql.ErrConnDone)
// 			},
// 			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusInternalServerError, recorder.Code)
// 			},
// 		},
// 		{
// 			name:      "InvalidID",
// 			accountID: 0,
// 			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
// 				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, "users", time.Minute)
// 			},
// 			buildStubs: func(store *mockdb.MockStore) {
// 				store.EXPECT().GetAccount(gomock.Any(), gomock.Any()).Times(0)
// 			},
// 			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusBadRequest, recorder.Code)
// 			},
// 		},
// 	}

// 	for i := range testCases {
// 		tc := testCases[i]

// 		t.Run(tc.name, func(t *testing.T) {
// 			// create mock controller
// 			ctrl := gomock.NewController(t)
// 			defer ctrl.Finish()

// 			// create mock store
// 			store := mockdb.NewMockStore(ctrl)
// 			tc.buildStubs(store)

// 			// create test server
// 			server := newTestServer(t, store)
// 			recorder := httptest.NewRecorder()

// 			// create test request
// 			url := fmt.Sprintf("/accounts/%d", tc.accountID)
// 			request, err := http.NewRequest(http.MethodGet, url, nil)

// 			require.NoError(t, err)
// 			server.router.ServeHTTP(recorder, request)

// 			// check response
// 			tc.checkResponse(t, recorder)
// 		})
// 	}

// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	store := mockdb.NewMockStore(ctrl)

// 	// build stubs
// 	store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account.ID)).Times(1).Return(account, nil).Return(account, nil)

// 	// start test server and send request
// 	server := newTestServer(t, store)
// 	recorder := httptest.NewRecorder()

// 	url := fmt.Sprintf("/accounts/%d", account.ID)
// 	request, err := http.NewRequest(http.MethodGet, url, nil)
// 	require.NoError(t, err)

// 	server.router.ServeHTTP(recorder, request)

// 	// check response
// 	require.Equal(t, http.StatusOK, recorder.Code)
// 	requireBodyMatchAccount(t, recorder.Body, account)
// }

// func randomAccount(owner string) db.Accounts {
// 	return db.Accounts{
// 		ID:       util.RandomInt(1, 1000),
// 		Owner:    owner,
// 		Balance:  util.RandomMoney(),
// 		Currency: util.RandomCurrency(),
// 	}
// }

// func requireBodyMatchAccount(t *testing.T, body *bytes.Buffer, account db.Accounts) {
// 	data, err := io.ReadAll(body)
// 	require.NoError(t, err)

// 	var gotAccount db.Accounts
// 	err = json.Unmarshal(data, &gotAccount)
// 	require.NoError(t, err)
// 	require.Equal(t, account, gotAccount)
// }
