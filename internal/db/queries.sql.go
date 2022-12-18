// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: queries.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const addEvent = `-- name: AddEvent :exec
INSERT INTO events (tag, name, organization, initial_labs, max_labs, frontend, status, exercises, started_at, finish_expected, finished_at, createdby, secretKey) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
`

type AddEventParams struct {
	Tag            string
	Name           string
	Organization   string
	InitialLabs    int32
	MaxLabs        int32
	Frontend       string
	Status         sql.NullInt32
	Exercises      string
	StartedAt      time.Time
	FinishExpected time.Time
	FinishedAt     sql.NullTime
	Createdby      string
	Secretkey      string
}

func (q *Queries) AddEvent(ctx context.Context, arg AddEventParams) error {
	_, err := q.db.ExecContext(ctx, addEvent,
		arg.Tag,
		arg.Name,
		arg.Organization,
		arg.InitialLabs,
		arg.MaxLabs,
		arg.Frontend,
		arg.Status,
		arg.Exercises,
		arg.StartedAt,
		arg.FinishExpected,
		arg.FinishedAt,
		arg.Createdby,
		arg.Secretkey,
	)
	return err
}

const addOrganization = `-- name: AddOrganization :exec
INSERT INTO organizations (name, owner_user, owner_email) VALUES ($1, $2, $3)
`

type AddOrganizationParams struct {
	Org           string
	Ownerusername string
	Owneremail    string
}

func (q *Queries) AddOrganization(ctx context.Context, arg AddOrganizationParams) error {
	_, err := q.db.ExecContext(ctx, addOrganization, arg.Org, arg.Ownerusername, arg.Owneremail)
	return err
}

const addProfile = `-- name: AddProfile :exec
INSERT INTO profiles (name, secret, challenges) VALUES ($1, $2, $3)
`

type AddProfileParams struct {
	Name       string
	Secret     bool
	Challenges string
}

func (q *Queries) AddProfile(ctx context.Context, arg AddProfileParams) error {
	_, err := q.db.ExecContext(ctx, addProfile, arg.Name, arg.Secret, arg.Challenges)
	return err
}

const addTeam = `-- name: AddTeam :exec
INSERT INTO teams (tag, event_id, email, username, password, created_at, last_access, solved_challenges) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
`

type AddTeamParams struct {
	Tag              string
	EventID          int32
	Email            string
	Username         string
	Password         string
	CreatedAt        time.Time
	LastAccess       sql.NullTime
	SolvedChallenges sql.NullString
}

func (q *Queries) AddTeam(ctx context.Context, arg AddTeamParams) error {
	_, err := q.db.ExecContext(ctx, addTeam,
		arg.Tag,
		arg.EventID,
		arg.Email,
		arg.Username,
		arg.Password,
		arg.CreatedAt,
		arg.LastAccess,
		arg.SolvedChallenges,
	)
	return err
}

const checkIfAgentExists = `-- name: CheckIfAgentExists :one
SELECT EXISTS( SELECT 1 FROM agents WHERE lower(name) = lower($1) )
`

func (q *Queries) CheckIfAgentExists(ctx context.Context, agentname string) (bool, error) {
	row := q.db.QueryRowContext(ctx, checkIfAgentExists, agentname)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const checkIfEventExist = `-- name: CheckIfEventExist :one
SELECT EXISTS (select tag from events where tag=$1)
`

func (q *Queries) CheckIfEventExist(ctx context.Context, tag string) (bool, error) {
	row := q.db.QueryRowContext(ctx, checkIfEventExist, tag)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const checkIfOrgExists = `-- name: CheckIfOrgExists :one
SELECT EXISTS( SELECT 1 FROM organizations WHERE lower(name) = lower($1) )
`

func (q *Queries) CheckIfOrgExists(ctx context.Context, orgname string) (bool, error) {
	row := q.db.QueryRowContext(ctx, checkIfOrgExists, orgname)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const checkIfUserExists = `-- name: CheckIfUserExists :one
SELECT EXISTS( SELECT 1 FROM admin_users WHERE lower(username) = lower($1) )
`

func (q *Queries) CheckIfUserExists(ctx context.Context, username string) (bool, error) {
	row := q.db.QueryRowContext(ctx, checkIfUserExists, username)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const checkIfUserExistsInOrg = `-- name: CheckIfUserExistsInOrg :one
SELECT EXISTS( SELECT 1 FROM admin_users WHERE lower(username) = lower($1) AND lower(organization) = lower($2))
`

type CheckIfUserExistsInOrgParams struct {
	Username     string
	Organization string
}

func (q *Queries) CheckIfUserExistsInOrg(ctx context.Context, arg CheckIfUserExistsInOrgParams) (bool, error) {
	row := q.db.QueryRowContext(ctx, checkIfUserExistsInOrg, arg.Username, arg.Organization)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const checkIfUserOwnsOrg = `-- name: CheckIfUserOwnsOrg :one
SELECT EXISTS( SELECT 1 FROM organizations WHERE lower(owner_user) = lower($1))
`

func (q *Queries) CheckIfUserOwnsOrg(ctx context.Context, ownerusername string) (bool, error) {
	row := q.db.QueryRowContext(ctx, checkIfUserOwnsOrg, ownerusername)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const checkProfileExists = `-- name: CheckProfileExists :one
SELECT EXISTS(SELECT 1 FROM profiles WHERE name = $1)
`

func (q *Queries) CheckProfileExists(ctx context.Context, name string) (bool, error) {
	row := q.db.QueryRowContext(ctx, checkProfileExists, name)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const createAdminUser = `-- name: CreateAdminUser :exec
INSERT INTO admin_users (username, password, full_name, email, role, organization) VALUES ($1, $2, $3, $4, $5, $6)
`

type CreateAdminUserParams struct {
	Username     string
	Password     string
	FullName     string
	Email        string
	Role         string
	Organization string
}

func (q *Queries) CreateAdminUser(ctx context.Context, arg CreateAdminUserParams) error {
	_, err := q.db.ExecContext(ctx, createAdminUser,
		arg.Username,
		arg.Password,
		arg.FullName,
		arg.Email,
		arg.Role,
		arg.Organization,
	)
	return err
}

const deleteAdminUserByUsername = `-- name: DeleteAdminUserByUsername :exec
DELETE FROM admin_users WHERE LOWER(username)=LOWER($1)
`

func (q *Queries) DeleteAdminUserByUsername(ctx context.Context, lower string) error {
	_, err := q.db.ExecContext(ctx, deleteAdminUserByUsername, lower)
	return err
}

const deleteAgentByName = `-- name: DeleteAgentByName :exec
DELETE FROM agents WHERE lower(name) = lower($1)
`

func (q *Queries) DeleteAgentByName(ctx context.Context, name string) error {
	_, err := q.db.ExecContext(ctx, deleteAgentByName, name)
	return err
}

const deleteEventByTagAndStatus = `-- name: DeleteEventByTagAndStatus :exec
DELETE FROM events WHERE tag=$1 and status=$2
`

type DeleteEventByTagAndStatusParams struct {
	Tag    string
	Status sql.NullInt32
}

func (q *Queries) DeleteEventByTagAndStatus(ctx context.Context, arg DeleteEventByTagAndStatusParams) error {
	_, err := q.db.ExecContext(ctx, deleteEventByTagAndStatus, arg.Tag, arg.Status)
	return err
}

const deleteEventOlderThan = `-- name: DeleteEventOlderThan :exec
DELETE FROM events WHERE finished_at < GETDATE() - $1 and status = 2
`

func (q *Queries) DeleteEventOlderThan(ctx context.Context, numberofdays interface{}) error {
	_, err := q.db.ExecContext(ctx, deleteEventOlderThan, numberofdays)
	return err
}

const deleteOrganization = `-- name: DeleteOrganization :exec
DELETE FROM organizations WHERE lower(name) = lower($1)
`

func (q *Queries) DeleteOrganization(ctx context.Context, orgname string) error {
	_, err := q.db.ExecContext(ctx, deleteOrganization, orgname)
	return err
}

const deleteProfile = `-- name: DeleteProfile :exec
DELETE FROM profiles WHERE name = $1
`

func (q *Queries) DeleteProfile(ctx context.Context, name string) error {
	_, err := q.db.ExecContext(ctx, deleteProfile, name)
	return err
}

const deleteTeam = `-- name: DeleteTeam :exec
DELETE FROM teams WHERE tag=$1 and event_id = $2
`

type DeleteTeamParams struct {
	Tag     string
	EventID int32
}

func (q *Queries) DeleteTeam(ctx context.Context, arg DeleteTeamParams) error {
	_, err := q.db.ExecContext(ctx, deleteTeam, arg.Tag, arg.EventID)
	return err
}

const getAdminUserByUsername = `-- name: GetAdminUserByUsername :one
SELECT id, username, password, full_name, email, role, organization FROM admin_users WHERE LOWER(username)=LOWER($1)
`

func (q *Queries) GetAdminUserByUsername(ctx context.Context, username string) (AdminUser, error) {
	row := q.db.QueryRowContext(ctx, getAdminUserByUsername, username)
	var i AdminUser
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.FullName,
		&i.Email,
		&i.Role,
		&i.Organization,
	)
	return i, err
}

const getAdminUserNoPwByUsername = `-- name: GetAdminUserNoPwByUsername :one
SELECT username, full_name, email, role, organization FROM admin_users WHERE LOWER(username)=LOWER($1)
`

type GetAdminUserNoPwByUsernameRow struct {
	Username     string
	FullName     string
	Email        string
	Role         string
	Organization string
}

func (q *Queries) GetAdminUserNoPwByUsername(ctx context.Context, lower string) (GetAdminUserNoPwByUsernameRow, error) {
	row := q.db.QueryRowContext(ctx, getAdminUserNoPwByUsername, lower)
	var i GetAdminUserNoPwByUsernameRow
	err := row.Scan(
		&i.Username,
		&i.FullName,
		&i.Email,
		&i.Role,
		&i.Organization,
	)
	return i, err
}

const getAdminUsers = `-- name: GetAdminUsers :many
SELECT username, full_name, email, role, organization FROM admin_users WHERE LOWER(organization) = CASE WHEN $1='' THEN LOWER(organization) ELSE LOWER($1) END
`

type GetAdminUsersRow struct {
	Username     string
	FullName     string
	Email        string
	Role         string
	Organization string
}

func (q *Queries) GetAdminUsers(ctx context.Context, organization interface{}) ([]GetAdminUsersRow, error) {
	rows, err := q.db.QueryContext(ctx, getAdminUsers, organization)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAdminUsersRow
	for rows.Next() {
		var i GetAdminUsersRow
		if err := rows.Scan(
			&i.Username,
			&i.FullName,
			&i.Email,
			&i.Role,
			&i.Organization,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAgentByName = `-- name: GetAgentByName :one
SELECT id, name, url, sign_key, auth_key, tls, statelock FROM agents WHERE lower(name) = lower($1)
`

func (q *Queries) GetAgentByName(ctx context.Context, name string) (Agent, error) {
	row := q.db.QueryRowContext(ctx, getAgentByName, name)
	var i Agent
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Url,
		&i.SignKey,
		&i.AuthKey,
		&i.Tls,
		&i.Statelock,
	)
	return i, err
}

const getAgents = `-- name: GetAgents :many
SELECT id, name, url, sign_key, auth_key, tls, statelock FROM agents
`

func (q *Queries) GetAgents(ctx context.Context) ([]Agent, error) {
	rows, err := q.db.QueryContext(ctx, getAgents)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Agent
	for rows.Next() {
		var i Agent
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Url,
			&i.SignKey,
			&i.AuthKey,
			&i.Tls,
			&i.Statelock,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllEvents = `-- name: GetAllEvents :many
SELECT id, tag, organization, name, initial_labs, max_labs, status, frontend, exercises, started_at, finish_expected, finished_at, createdby, secretkey FROM events
`

func (q *Queries) GetAllEvents(ctx context.Context) ([]Event, error) {
	rows, err := q.db.QueryContext(ctx, getAllEvents)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Event
	for rows.Next() {
		var i Event
		if err := rows.Scan(
			&i.ID,
			&i.Tag,
			&i.Organization,
			&i.Name,
			&i.InitialLabs,
			&i.MaxLabs,
			&i.Status,
			&i.Frontend,
			&i.Exercises,
			&i.StartedAt,
			&i.FinishExpected,
			&i.FinishedAt,
			&i.Createdby,
			&i.Secretkey,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getEventStatus = `-- name: GetEventStatus :many
SELECT status FROM events WHERE tag=$1
`

func (q *Queries) GetEventStatus(ctx context.Context, tag string) ([]sql.NullInt32, error) {
	rows, err := q.db.QueryContext(ctx, getEventStatus, tag)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []sql.NullInt32
	for rows.Next() {
		var status sql.NullInt32
		if err := rows.Scan(&status); err != nil {
			return nil, err
		}
		items = append(items, status)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getEventsByStatus = `-- name: GetEventsByStatus :many
SELECT id, tag, organization, name, initial_labs, max_labs, status, frontend, exercises, started_at, finish_expected, finished_at, createdby, secretkey FROM events WHERE status=$1
`

func (q *Queries) GetEventsByStatus(ctx context.Context, status sql.NullInt32) ([]Event, error) {
	rows, err := q.db.QueryContext(ctx, getEventsByStatus, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Event
	for rows.Next() {
		var i Event
		if err := rows.Scan(
			&i.ID,
			&i.Tag,
			&i.Organization,
			&i.Name,
			&i.InitialLabs,
			&i.MaxLabs,
			&i.Status,
			&i.Frontend,
			&i.Exercises,
			&i.StartedAt,
			&i.FinishExpected,
			&i.FinishedAt,
			&i.Createdby,
			&i.Secretkey,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getEventsByUser = `-- name: GetEventsByUser :many
SELECT id, tag, organization, name, initial_labs, max_labs, status, frontend, exercises, started_at, finish_expected, finished_at, createdby, secretkey FROM events WHERE createdBy=$1
`

func (q *Queries) GetEventsByUser(ctx context.Context, createdby string) ([]Event, error) {
	rows, err := q.db.QueryContext(ctx, getEventsByUser, createdby)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Event
	for rows.Next() {
		var i Event
		if err := rows.Scan(
			&i.ID,
			&i.Tag,
			&i.Organization,
			&i.Name,
			&i.InitialLabs,
			&i.MaxLabs,
			&i.Status,
			&i.Frontend,
			&i.Exercises,
			&i.StartedAt,
			&i.FinishExpected,
			&i.FinishedAt,
			&i.Createdby,
			&i.Secretkey,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getEventsExeptClosed = `-- name: GetEventsExeptClosed :many
SELECT id, tag, organization, name, initial_labs, max_labs, status, frontend, exercises, started_at, finish_expected, finished_at, createdby, secretkey FROM events WHERE status!=2
`

func (q *Queries) GetEventsExeptClosed(ctx context.Context) ([]Event, error) {
	rows, err := q.db.QueryContext(ctx, getEventsExeptClosed)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Event
	for rows.Next() {
		var i Event
		if err := rows.Scan(
			&i.ID,
			&i.Tag,
			&i.Organization,
			&i.Name,
			&i.InitialLabs,
			&i.MaxLabs,
			&i.Status,
			&i.Frontend,
			&i.Exercises,
			&i.StartedAt,
			&i.FinishExpected,
			&i.FinishedAt,
			&i.Createdby,
			&i.Secretkey,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getExpectedFinishDate = `-- name: GetExpectedFinishDate :one
SELECT finish_expected FROM events WHERE tag=$1
`

func (q *Queries) GetExpectedFinishDate(ctx context.Context, tag string) (time.Time, error) {
	row := q.db.QueryRowContext(ctx, getExpectedFinishDate, tag)
	var finish_expected time.Time
	err := row.Scan(&finish_expected)
	return finish_expected, err
}

const getOrgByName = `-- name: GetOrgByName :one
SELECT id, name, owner_user, owner_email FROM organizations WHERE LOWER(name)=LOWER($1)
`

func (q *Queries) GetOrgByName(ctx context.Context, orgname string) (Organization, error) {
	row := q.db.QueryRowContext(ctx, getOrgByName, orgname)
	var i Organization
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.OwnerUser,
		&i.OwnerEmail,
	)
	return i, err
}

const getOrganizations = `-- name: GetOrganizations :many
SELECT id, name, owner_user, owner_email FROM organizations
`

func (q *Queries) GetOrganizations(ctx context.Context) ([]Organization, error) {
	rows, err := q.db.QueryContext(ctx, getOrganizations)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Organization
	for rows.Next() {
		var i Organization
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.OwnerUser,
			&i.OwnerEmail,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getProfiles = `-- name: GetProfiles :many
SELECT id, name, secret, organization, challenges FROM profiles ORDER BY id asc
`

func (q *Queries) GetProfiles(ctx context.Context) ([]Profile, error) {
	rows, err := q.db.QueryContext(ctx, getProfiles)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Profile
	for rows.Next() {
		var i Profile
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Secret,
			&i.Organization,
			&i.Challenges,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTeamCount = `-- name: GetTeamCount :many
SELECT count(teams.id) FROM teams WHERE teams.event_id=$1
`

func (q *Queries) GetTeamCount(ctx context.Context, eventID int32) ([]int64, error) {
	rows, err := q.db.QueryContext(ctx, getTeamCount, eventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []int64
	for rows.Next() {
		var count int64
		if err := rows.Scan(&count); err != nil {
			return nil, err
		}
		items = append(items, count)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTeamsForEvent = `-- name: GetTeamsForEvent :many
SELECT id, tag, event_id, email, username, password, created_at, last_access, solved_challenges FROM teams WHERE event_id=$1
`

func (q *Queries) GetTeamsForEvent(ctx context.Context, eventID int32) ([]Team, error) {
	rows, err := q.db.QueryContext(ctx, getTeamsForEvent, eventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Team
	for rows.Next() {
		var i Team
		if err := rows.Scan(
			&i.ID,
			&i.Tag,
			&i.EventID,
			&i.Email,
			&i.Username,
			&i.Password,
			&i.CreatedAt,
			&i.LastAccess,
			&i.SolvedChallenges,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertNewAgent = `-- name: InsertNewAgent :exec
INSERT INTO agents (name, url, sign_key, auth_key, tls, statelock) VALUES ($1, $2, $3, $4, $5, false)
`

type InsertNewAgentParams struct {
	Name    string
	Url     string
	Signkey string
	Authkey string
	Tls     bool
}

func (q *Queries) InsertNewAgent(ctx context.Context, arg InsertNewAgentParams) error {
	_, err := q.db.ExecContext(ctx, insertNewAgent,
		arg.Name,
		arg.Url,
		arg.Signkey,
		arg.Authkey,
		arg.Tls,
	)
	return err
}

const teamSolvedChls = `-- name: TeamSolvedChls :many
SELECT solved_challenges FROM teams WHERE tag=$1
`

func (q *Queries) TeamSolvedChls(ctx context.Context, tag string) ([]sql.NullString, error) {
	rows, err := q.db.QueryContext(ctx, teamSolvedChls, tag)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []sql.NullString
	for rows.Next() {
		var solved_challenges sql.NullString
		if err := rows.Scan(&solved_challenges); err != nil {
			return nil, err
		}
		items = append(items, solved_challenges)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateAdminEmail = `-- name: UpdateAdminEmail :exec
UPDATE admin_users SET email = $1 WHERE username = $2
`

type UpdateAdminEmailParams struct {
	Email    string
	Username string
}

func (q *Queries) UpdateAdminEmail(ctx context.Context, arg UpdateAdminEmailParams) error {
	_, err := q.db.ExecContext(ctx, updateAdminEmail, arg.Email, arg.Username)
	return err
}

const updateAdminPassword = `-- name: UpdateAdminPassword :exec
UPDATE admin_users SET password = $1 WHERE username = $2
`

type UpdateAdminPasswordParams struct {
	Password string
	Username string
}

func (q *Queries) UpdateAdminPassword(ctx context.Context, arg UpdateAdminPasswordParams) error {
	_, err := q.db.ExecContext(ctx, updateAdminPassword, arg.Password, arg.Username)
	return err
}

const updateCloseEvent = `-- name: UpdateCloseEvent :exec
UPDATE events SET tag = $2, finished_at = $3 WHERE tag = $1
`

type UpdateCloseEventParams struct {
	Tag        string
	Tag_2      string
	FinishedAt sql.NullTime
}

func (q *Queries) UpdateCloseEvent(ctx context.Context, arg UpdateCloseEventParams) error {
	_, err := q.db.ExecContext(ctx, updateCloseEvent, arg.Tag, arg.Tag_2, arg.FinishedAt)
	return err
}

const updateEventStatus = `-- name: UpdateEventStatus :exec
UPDATE events SET status = $2 WHERE tag = $1
`

type UpdateEventStatusParams struct {
	Tag    string
	Status sql.NullInt32
}

func (q *Queries) UpdateEventStatus(ctx context.Context, arg UpdateEventStatusParams) error {
	_, err := q.db.ExecContext(ctx, updateEventStatus, arg.Tag, arg.Status)
	return err
}

const updateExercises = `-- name: UpdateExercises :exec

UPDATE teams SET last_access = $2 WHERE tag = $1
`

type UpdateExercisesParams struct {
	Tag        string
	LastAccess sql.NullTime
}

// UPDATE event SET exercises = (SELECT (SELECT exercises FROM event WHERE id = $1) || $2) WHERE id=$1;
func (q *Queries) UpdateExercises(ctx context.Context, arg UpdateExercisesParams) error {
	_, err := q.db.ExecContext(ctx, updateExercises, arg.Tag, arg.LastAccess)
	return err
}

const updateOrganization = `-- name: UpdateOrganization :exec
UPDATE organizations SET owner_user = $1, owner_email = $2 WHERE lower(name) = lower($3)
`

type UpdateOrganizationParams struct {
	Ownerusername string
	Owneremail    string
	Orgname       string
}

func (q *Queries) UpdateOrganization(ctx context.Context, arg UpdateOrganizationParams) error {
	_, err := q.db.ExecContext(ctx, updateOrganization, arg.Ownerusername, arg.Owneremail, arg.Orgname)
	return err
}

const updateProfile = `-- name: UpdateProfile :exec
UPDATE profiles SET secret = $1, challenges = $2 WHERE name = $3
`

type UpdateProfileParams struct {
	Secret     bool
	Challenges string
	Name       string
}

func (q *Queries) UpdateProfile(ctx context.Context, arg UpdateProfileParams) error {
	_, err := q.db.ExecContext(ctx, updateProfile, arg.Secret, arg.Challenges, arg.Name)
	return err
}

const updateTeamPassword = `-- name: UpdateTeamPassword :exec
UPDATE teams SET password = $1 WHERE tag = $2 and event_id = $3
`

type UpdateTeamPasswordParams struct {
	Password string
	Tag      string
	EventID  int32
}

func (q *Queries) UpdateTeamPassword(ctx context.Context, arg UpdateTeamPasswordParams) error {
	_, err := q.db.ExecContext(ctx, updateTeamPassword, arg.Password, arg.Tag, arg.EventID)
	return err
}

const updateTeamSolvedChl = `-- name: UpdateTeamSolvedChl :exec
UPDATE teams SET solved_challenges = $2 WHERE tag = $1
`

type UpdateTeamSolvedChlParams struct {
	Tag              string
	SolvedChallenges sql.NullString
}

func (q *Queries) UpdateTeamSolvedChl(ctx context.Context, arg UpdateTeamSolvedChlParams) error {
	_, err := q.db.ExecContext(ctx, updateTeamSolvedChl, arg.Tag, arg.SolvedChallenges)
	return err
}
