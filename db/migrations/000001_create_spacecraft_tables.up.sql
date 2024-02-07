CREATE TABLE IF NOT EXISTS spacecraft (
    id BIGINT NOT NULL AUTO_INCREMENT,
    name VARCHAR(20),
    class VARCHAR(10),
    crew INTEGER,
    image VARCHAR(80),
    value DECIMAL(11, 2),
    status VARCHAR(16),
    created_at DATETIME,
    updated_at DATETIME,
    PRIMARY KEY (id),
    INDEX sc_name_idx (name),
    INDEX sc_class_idx (class),
    INDEX sc_status_idx (status),
    INDEX sc_name_class_idx (name, class),
    INDEX sc_name_status_idx (name, status),
    INDEX sc_class_status_idx (class, status),
    INDEX sc_name_class_status_idx (name, class, status)
);

CREATE TABLE IF NOT EXISTS armament (
    id BIGINT NOT NULL AUTO_INCREMENT,
    title VARCHAR(32),
    created_at DATETIME,
    updated_at DATETIME,
    PRIMARY KEY (id),
    INDEX ar_title_idx (title)
);

CREATE TABLE IF NOT EXISTS armament_detail (
    id BIGINT NOT NULL AUTO_INCREMENT,
    armament_id BIGINT NOT NULL,
    spacecraft_id BIGINT NOT NULL,
    qty INTEGER,
    PRIMARY KEY (id),
    INDEX ad_ar_idx (armament_id),
    INDEX ad_sc_idx (spacecraft_id),
    INDEX ad_sc_ar_idx (spacecraft_id, armament_id),
    FOREIGN KEY (spacecraft_id) REFERENCES spacecraft(id),
    FOREIGN KEY (armament_id) REFERENCES armament(id)
);