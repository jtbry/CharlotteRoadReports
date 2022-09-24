CREATE INDEX active_idx ON incidents (active) USING BTREE;
CREATE INDEX start_timestamp_idx ON incidents (start_timestamp) USING BTREE;