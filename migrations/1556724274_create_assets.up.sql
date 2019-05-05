CREATE TABLE IF NOT EXISTS assets (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    type text NOT NULL,
    url text NOT NULL,
    is_private boolean,
    tags text, 
    created_at timestamp with time zone NOT NULL default current_timestamp,
    deleted_at timestamp with time zone,
    user_id uuid NOT NULL   
);