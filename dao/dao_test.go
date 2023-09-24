package dao

import (
	"fmt"
	"reflect"
	"testing"
	"url-shortener/internal"
	"url-shortener/model"
)

var insertDBData = &model.URL{
	ShortURL:  "fromdbShortURL",
	LongURL:   "fromdbLongURL",
	ExpiresAt: nil,
}

var insertCacheData = &model.URL{
	ShortURL:  "fromCacheShortURL",
	LongURL:   "fromCacheLongURL",
	ExpiresAt: nil,
}

func TestDaoImpl_QueryURLRecord(t *testing.T) {
	type tesetCase struct {
		name            string
		query           string
		want            *model.URL
		wantErr         bool
		additionalSetup func(dao *DaoImpl, t *testing.T, testCase tesetCase)
	}

	tests := []tesetCase{
		{
			name:  "test query from cache",
			query: "fromCacheShortURL",
			want: &model.URL{
				ShortURL:  "fromCacheShortURL",
				LongURL:   "fromCacheLongURL",
				ExpiresAt: nil,
			},
			wantErr: false,
		},
		{
			name:    "test query from db",
			query:   "fromdbShortURL",
			want:    insertDBData,
			wantErr: false,
		},
		{
			name:  "test query from cache after from db",
			query: "fromdbShortURL",
			want: &model.URL{
				ShortURL:  "fromdbShortURL",
				LongURL:   "fromdbLongURL",
				ExpiresAt: nil,
			},
			wantErr: false,
			additionalSetup: func(dao *DaoImpl, t *testing.T, tt tesetCase) {
				got, err := dao.cache.QueryCache("fromdbShortURL")
				if (err != nil) != tt.wantErr {
					t.Errorf("DaoImpl.QueryURLRecord() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("DaoImpl.QueryURLRecord() = %v, want %v", got, tt.want)
				}
			},
		},
		{
			name:    "test not exists data",
			query:   "notExists",
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, cache := makeMockStorage()
			dao := &DaoImpl{
				db:    db,
				cache: cache,
			}
			got, err := dao.QueryURLRecord(tt.query)
			fmt.Println(internal.DumpStruct(got))
			if (err != nil) != tt.wantErr {
				t.Errorf("DaoImpl.QueryURLRecord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DaoImpl.QueryURLRecord() = %v, want %v", got, tt.want)
			}
			if tt.additionalSetup != nil {
				tt.additionalSetup(dao, t, tt)
			}
		})
	}
}

func makeMockStorage() (DBinterface, CacheInterface) {
	db := newDatabaseMock()
	db.WriteDB(insertDBData)
	cache := newCacheMock()
	cache.WriteCache(insertCacheData)
	return db, cache
}
