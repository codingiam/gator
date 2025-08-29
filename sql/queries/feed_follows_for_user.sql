-- name: FeedFollowsForUser :many
SELECT feeds.name AS feed_name, feeds.url, fu.name AS user_name
FROM feed_follows
INNER JOIN users ffu ON feed_follows.user_id = ffu.id
INNER JOIN feeds ON feed_follows.feed_id = feeds.id
INNER JOIN users fu ON feeds.user_id = fu.id
-- WHERE ffu.name = $1;
