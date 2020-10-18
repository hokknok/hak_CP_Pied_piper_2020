CREATE TABLE Poligon (
	    id    serial PRIMARY KEY,
	    erlid  varchar(30) NOT NULL,
	    p1x float  NOT NULL,
		p1y float  NOT NULL,
	    p2x float  NOT NULL,
		p2y float  NOT NULL,
	    p3x float  NOT NULL,
		p3y float  NOT NULL,
	    p4x float  NOT NULL,
		p4y float  NOT NULL
);


Create TABLE dateP (
        id    serial PRIMARY KEY,
	    zid varchar(30) NOT NULL,
        ts   timestamp NOT NULL,
        customers_cnt_long_sum integer,
        customers_cnt_work_sum integer,
        customers_cnt_loc_sum integer,
        customers_cnt_total_sum integer
);


CREATE TABLE PointEx (
        id    serial PRIMARY KEY,
	    px float  NOT NULL,
		py float  NOT NULL
);
