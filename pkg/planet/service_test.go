package planet

import (
	"strconv"
	"testing"
	"time"

	"github.com/Ocelani/swapi-planets/gen"
	"github.com/stretchr/testify/assert"
)

var (
	Error   = assert.Error
	NoError = assert.NoError
)

type ErrorAssertion assert.ErrorAssertionFunc

func TestNewService(t *testing.T) {
	s := NewDefaultService(NewCacheRepository([]*gen.Planet{}))
	assert.NotNil(t, s)
}

func Test_service_Create(t *testing.T) {
	tests := []struct {
		name string
		have []*gen.Planet                  // Database saved data
		args map[*gen.Planet]ErrorAssertion // Planets to be created and desired error value
	}{
		{
			name: "create planets with nil error",
			have: []*gen.Planet{},
			args: map[*gen.Planet]ErrorAssertion{
				{Name: "test"}: NoError,
			},
		},
	}
	for _, tt := range tests {
		repo := &CacheRepository{
			Data: tt.have,
		}
		s := NewDefaultService(repo)

		created := len(tt.have)

		for planet, wantErr := range tt.args {
			t.Run(tt.name, func(t *testing.T) {
				err := s.Create(planet)

				wantErr(t, err)
				if err != nil {
					return
				}
				created++

				assert.Len(t, repo.Data, created)
				assert.Equal(t, planet, repo.Data[created-1])
				assert.NotEmpty(t, planet.Id)
			})
		}
	}
}

func Test_service_ReadAll(t *testing.T) {
	tests := []struct {
		name    string
		want    []*gen.Planet // Database saved Data to be retrieved
		wantErr ErrorAssertion
	}{
		{
			wantErr: NoError,
			want: []*gen.Planet{
				{Id: "1"}, {Id: "2"}, {Id: "3"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &CacheRepository{
				Data: tt.want,
			}
			s := NewDefaultService(repo)

			got, err := s.ReadAll()

			assert.NoError(t, err)
			assert.NotNil(t, got)
			assert.NotEmpty(t, got)
			assert.Len(t, got, len(tt.want))
		})
	}
}

func Test_service_ReadOne(t *testing.T) {
	tests := []struct {
		name    string
		have    []*gen.Planet // Database saved data
		wantID  int           // gen.Planet  Id to be retrieved
		wantErr ErrorAssertion
	}{
		{
			name: "read planet with nil error",
			have: []*gen.Planet{
				{Id: "1", Name: "test-1"},
				{Id: "2", Name: "test-2"},
				{Id: "3", Name: "test-3"},
			},
			wantID:  2,
			wantErr: NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &CacheRepository{
				Data: tt.have,
			}
			s := NewDefaultService(repo)

			id := strconv.Itoa(tt.wantID)
			got, err := s.ReadOne(id)

			tt.wantErr(t, err)
			if err != nil {
				return
			}
			assert.NotNil(t, got)
			assert.NotEmpty(t, got)
			assert.Equal(t, tt.have[tt.wantID-1], got)
		})
	}
}

func Test_service_Update(t *testing.T) {
	tests := []struct {
		name    string
		have    []*gen.Planet // Database saved data
		updt    *gen.Planet   // gen.Planet  to be updated
		wantErr ErrorAssertion
	}{
		{
			have: []*gen.Planet{
				{Id: "1", Name: "test-1"},
				{Id: "2", Name: "test-2"},
				{Id: "3", Name: "test-3"},
			},
			updt: &gen.Planet{
				Id: "2", Name: "update",
			},
			wantErr: NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &CacheRepository{
				Data: tt.have,
			}
			s := NewDefaultService(repo)

			if planet, _ := s.ReadOne(tt.updt.Id); planet != nil {
				assert.Eventually(t, func() bool {
					for _, p := range tt.have {
						if p.Id == planet.Id {
							return p == planet
						}
					}
					return false
				}, 3*time.Second, 100*time.Millisecond)
			}
			assert.NotSubset(t, repo.Data, []*gen.Planet{tt.updt})

			err := s.Update(tt.updt)

			tt.wantErr(t, err)
			if err != nil {
				return
			}
			assert.Subset(t, repo.Data, []*gen.Planet{tt.updt})
		})
	}
}

func Test_service_Delete(t *testing.T) {
	tests := []struct {
		name    string
		have    []*gen.Planet // Database saved data
		delID   []string      // gen.Planet  Id to be deleted
		wantErr ErrorAssertion
	}{
		{
			name:  "delete planet with nil error",
			delID: []string{"2", "1", "3"},
			have: []*gen.Planet{
				{Id: "1", Name: "test-1"},
				{Id: "2", Name: "test-2"},
				{Id: "3", Name: "test-3"},
			},
			wantErr: NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &CacheRepository{
				Data: tt.have,
			}
			s := NewDefaultService(repo)

			for _, id := range tt.delID {
				err := s.Delete(id)

				tt.wantErr(t, err)
				if err != nil {
					return
				}
				before := len(tt.have)
				after := len(repo.Data)
				assert.Less(t, after, before)

				_, err = s.ReadOne(id)
				assert.Error(t, err)
			}
		})
	}
}
