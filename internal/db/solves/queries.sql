-- name: GetSolvesForEvent :many
--SELECT tag, COUNT(tag) FROM solves WHERE event_id = @eventId GROUP BY tag;