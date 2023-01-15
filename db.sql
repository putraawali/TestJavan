CREATE TABLE members(
	member_id SERIAL PRIMARY KEY,
	member_name VARCHAR(20) NOT NULL,
	gender VARCHAR(1) NOT NULL,
	created_at TIMESTAMP DEFAULT NOW(),
	updated_at TIMESTAMP,
	deleted_at TIMESTAMP
);

INSERT INTO members(member_name, gender) 
VALUES
	('Bani', 'M'),
	('Budi', 'M'),
	('Nida', 'F'),
	('Andi', 'M'),
	('Sigit', 'M'),
	('Hari', 'M'),
	('Siti', 'F'),
	('Bila', 'F'),
	('Lesti', 'F'),
	('Diki', 'M'),
	('Doni', 'M'),
	('Toni', 'M');

CREATE TABLE assets(
	asset_id SERIAL PRIMARY KEY,
	asset_name VARCHAR(100) NOT NULL,
	created_at TIMESTAMP DEFAULT NOW(),
	updated_at TIMESTAMP,
	deleted_at TIMESTAMP
);

CREATE TABLE member_assets(
	member_asset_id SERIAL PRIMARY KEY,
	member_id int8 NOT NULL,
	asset_id int8 NOT NULL,
	created_at TIMESTAMP DEFAULT NOW(),
	updated_at TIMESTAMP,
	deleted_at TIMESTAMP,
	CONSTRAINT fk_member_id FOREIGN KEY(member_id) REFERENCES members(member_id),
	CONSTRAINT fk_asset_id FOREIGN KEY (asset_id) REFERENCES assets(asset_id)
);

INSERT INTO member_assets(member_id, asset_id)
VALUES
	((SELECT m.member_id FROM members m WHERE m.member_name = 'Budi'), (SELECT a.asset_id FROM assets a WHERE a.asset_name = 'Samsung Universe 9')),
	((SELECT m.member_id FROM members m WHERE m.member_name = 'Budi'), (SELECT a.asset_id FROM assets a WHERE a.asset_name = 'Samsung Galaxy Book')),
	((SELECT m.member_id FROM members m WHERE m.member_name = 'Hari'), (SELECT a.asset_id FROM assets a WHERE a.asset_name = 'iPhone 9')),
	((SELECT m.member_id FROM members m WHERE m.member_name = 'Siti'), (SELECT a.asset_id FROM assets a WHERE a.asset_name = 'iPhone X')),
	((SELECT m.member_id FROM members m WHERE m.member_name = 'Nida'), (SELECT a.asset_id FROM assets a WHERE a.asset_name = 'Huawei P30')),
	((SELECT m.member_id FROM members m WHERE m.member_name = 'Bila'), (SELECT a.asset_id FROM assets a WHERE a.asset_name = 'Samsung Universe 9')),
	((SELECT m.member_id FROM members m WHERE m.member_name = 'Lesti'), (SELECT a.asset_id FROM assets a WHERE a.asset_name = 'Huawei P30')),
	((SELECT m.member_id FROM members m WHERE m.member_name = 'Lesti'), (SELECT a.asset_id FROM assets a WHERE a.asset_name = 'Huawei P30')),
	((SELECT m.member_id FROM members m WHERE m.member_name = 'Andi'), (SELECT a.asset_id FROM assets a WHERE a.asset_name = 'Samsung Universe 9')),
	((SELECT m.member_id FROM members m WHERE m.member_name = 'Diki'), (SELECT a.asset_id FROM assets a WHERE a.asset_name = 'Samsung Galaxy Book')),
	((SELECT m.member_id FROM members m WHERE m.member_name = 'Sigit'), (SELECT a.asset_id FROM assets a WHERE a.asset_name = 'Huawei P30')),
	((SELECT m.member_id FROM members m WHERE m.member_name = 'Doni'), (SELECT a.asset_id FROM assets a WHERE a.asset_name = 'iPhone X'));
