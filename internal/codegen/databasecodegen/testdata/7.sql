CREATE TABLE shippers (
    shipper_id STRING(63) NOT NULL,
    revision_id STRING(8),
    create_time TIMESTAMP NOT NULL OPTIONS (allow_commit_timestamp=true),
    update_time TIMESTAMP NOT NULL OPTIONS (allow_commit_timestamp=true),
    delete_time TIMESTAMP OPTIONS (allow_commit_timestamp=true),
) PRIMARY KEY(shipper_id, revision_id);
