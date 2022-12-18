// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import (
	"database/sql"
	"time"
)

type AdminUser struct {
	ID           int32
	Username     string
	Password     string
	FullName     string
	Email        string
	Role         string
	Organization string
}

type Agent struct {
	ID        int32
	Name      string
	Url       string
	SignKey   string
	AuthKey   string
	Tls       bool
	Statelock bool
}

type Event struct {
	ID             int32
	Tag            string
	Organization   string
	Name           string
	InitialLabs    int32
	MaxLabs        int32
	Status         sql.NullInt32
	Frontend       string
	Exercises      string
	StartedAt      time.Time
	FinishExpected time.Time
	FinishedAt     sql.NullTime
	Createdby      string
	Secretkey      string
}

type Frontend struct {
	ID       int32
	Name     string
	Image    string
	Memorymb sql.NullInt32
}

type Organization struct {
	ID         int32
	Name       string
	OwnerUser  string
	OwnerEmail string
}

type Profile struct {
	ID           int32
	Name         string
	Secret       bool
	Organization string
	Challenges   string
}

type Team struct {
	ID               int32
	Tag              string
	EventID          int32
	Email            string
	Username         string
	Password         string
	CreatedAt        time.Time
	LastAccess       sql.NullTime
	SolvedChallenges sql.NullString
}
