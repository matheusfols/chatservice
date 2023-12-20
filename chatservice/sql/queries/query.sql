-- name: CreateChat :exec
INSERT INTO chats 
    (id, user_id, initial_message_id, status, token_usage, model, model_max_tokens,temperature, top_p, n, stop, max_tokens, presence_penalty, frequency_penalty, created_at, updated_at)
    VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16);

-- name: AddMessage :exec
INSERT INTO messages (id, chat_id, role, content, tokens, model, erased, order_msg, created_at) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9);

-- name: FindChatByID :one
SELECT * FROM chats WHERE id = $1;

-- name: FindMessagesByChatID :many
SELECT * FROM messages WHERE erased=0 and chat_id = $1 order by order_msg asc;

-- name: FindErasedMessagesByChatID :many
SELECT * FROM messages WHERE erased=1 and chat_id = $1 order by order_msg asc;

-- name: SaveChat :exec
UPDATE chats SET user_id = $1, initial_message_id = $2, status = $3, token_usage = $4, model = $5, model_max_tokens=$6, temperature = $7, top_p = $8, n = $9, stop = $10, max_tokens = $11, presence_penalty = $12, frequency_penalty = $13, updated_at = $14 WHERE id = $15;

-- name: DeleteChatMessages :exec
DELETE FROM messages WHERE chat_id = $1;

-- name: DeleteErasedChatMessages :exec
DELETE FROM messages WHERE erased=1 and chat_id = $1;
