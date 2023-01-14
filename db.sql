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

CREATE TABLE member_assets(
	asset_id SERIAL PRIMARY KEY,
	member_id int8 NOT NULL,
	asset_name VARCHAR(100) NOT NULL,
	created_at TIMESTAMP DEFAULT NOW(),
	updated_at TIMESTAMP,
	deleted_at TIMESTAMP,
	CONSTRAINT fk_member_id FOREIGN KEY (member_id) REFERENCES members(member_id)
);

INSERT INTO member_assets(member_id, asset_name)
VALUES
	((SELECT m.member_id FROM members m WHERE m.member_name = 'Budi'), 'Samsung Universe 9'),
	((SELECT m.member_id FROM members m WHERE m.member_name = 'Budi'), 'Samsung Galaxy Book'),
	((SELECT m.member_id FROM members m WHERE m.member_name = 'Hari'), 'iPhone 9'),
	((SELECT m.member_id FROM members m WHERE m.member_name = 'Siti'), 'iPhone X'),
	((SELECT m.member_id FROM members m WHERE m.member_name = 'Nida'), 'Huawei P90'),
	((SELECT m.member_id FROM members m WHERE m.member_name = 'Bila'), 'Samsung Universe 9'),
	((SELECT m.member_id FROM members m WHERE m.member_name = 'Lesti'), 'Huawei P90'),
	((SELECT m.member_id FROM members m WHERE m.member_name = 'Lesti'), 'Huawei iPhone X'),
	((SELECT m.member_id FROM members m WHERE m.member_name = 'Andi'), 'Samsung Universe 9'),
	((SELECT m.member_id FROM members m WHERE m.member_name = 'Diki'), 'Samsung Galaxy Book'),
	((SELECT m.member_id FROM members m WHERE m.member_name = 'Sigit'), 'Huawei P90'),
	((SELECT m.member_id FROM members m WHERE m.member_name = 'Doni'), 'iPhone X');

