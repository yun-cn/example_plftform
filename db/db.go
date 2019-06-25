package db

import (
	"database/sql"
	"fmt"

	"github.com/yanshiyason/noonde_platform/utils"

	"github.com/spf13/viper"

	// initialize postgres
	_ "github.com/lib/pq"
)

// Service Database service
type Service struct {
	Conn *sql.DB
}

// NewService Init Service with database connection
func NewService() *Service {

	connStr := fmt.Sprintf(
		"role=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		viper.GetString("database.username"),
		viper.GetString("database.username"),
		viper.GetString("database.password"),
		viper.GetString("database.name"),
		viper.GetInt("database.port"))

	conn, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(fmt.Errorf("Couldn't open DB connection: %s", err))
	}

	return &Service{Conn: conn}
}

// RoomIds get room ids for spaces which are for day time usage
func (s *Service) RoomIds() ([]string, error) {
	rows, err := s.Conn.Query(`SELECT uid FROM app.rooms_search_results_rooms_ids`)

	if err != nil {
		return nil, err
	}

	uids := []string{}

	for rows.Next() {
		var uid string
		err = rows.Scan(&uid)
		utils.MaybePanic("", err)

		uids = append(uids, uid)
	}

	return uids, nil
}

// StayRoomIds get room ids for spaces which are for night time usage
func (s *Service) StayRoomIds() ([]string, error) {
	rows, err := s.Conn.Query(`SELECT uid FROM app.stay_rooms_search_results_rooms_ids`)

	if err != nil {
		return nil, err
	}

	uids := []string{}

	for rows.Next() {
		var uid string
		err = rows.Scan(&uid)
		utils.MaybePanic("", err)

		uids = append(uids, uid)
	}

	return uids, nil
}

// RoomIds returns ids of day rooms
func RoomIds() ([]string, error) {
	s := NewService()
	defer s.Conn.Close()

	ids, err := s.RoomIds()

	if err != nil {
		return nil, err
	}

	return ids, nil
}

// StayRoomIds returns ids of night rooms
func StayRoomIds() ([]string, error) {
	s := NewService()
	defer s.Conn.Close()

	ids, err := s.StayRoomIds()

	if err != nil {
		return nil, err
	}

	return ids, nil
}
