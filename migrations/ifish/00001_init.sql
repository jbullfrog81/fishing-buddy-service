CREATE TABLE catches (
  id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  coordinates point SPATIAL KEY NOT NULL,
  caught_at DATETIME(3) NOT NULL,
  user_id INT NOT NULL,
  fish_species_id INT NOT NULL,
  created_at DATETIME(3) NOT NULL DEFAULT NOW(3),
  updated_at DATETIME(3) NOT NULL DEFAULT NOW(3) ON UPDATE NOW(3)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- support querying by user_id
CREATE INDEX idx_catches__user_id ON catches(user_id);

-- support querying by fish_species_id
CREATE INDEX idx_catches__fish_species_id ON catches(fish_species_id);

-- support querying by caught_at
CREATE INDEX idx_catches__caught_at ON catches(caught_at);

-- support querying by created_at
CREATE INDEX idx_catches__created_at ON catches(created_at);

-- support querying by created_at
CREATE INDEX idx_catches__updated_at ON catches(updated_at);

CREATE TABLE fish_species (
  id INT PRIMARY KEY NOT NULL,
  name VARCHAR(255) NOT NULL,
  created_at DATETIME(3) NOT NULL DEFAULT NOW(3),
  updated_at DATETIME(3) NOT NULL DEFAULT NOW(3) ON UPDATE NOW(3)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ensure fish_species does not duplicates
CREATE UNIQUE INDEX uidx_fish_species__name ON fish_species(name);

INSERT INTO fish_species(id, name) VALUES
(1, "Walleye"),
(2, "Crappie"),
(3, "Largemouth Bass"),
(4, "White Bass"),
(5, "Wiper"),
(6, "Striped Bass"),
(7, "Rainbow Trout");

CREATE TABLE users (
  id INT PRIMARY KEY NOT NULL,
  username VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL,
  created_at DATETIME(3) NOT NULL DEFAULT NOW(3),
  updated_at DATETIME(3) NOT NULL DEFAULT NOW(3) ON UPDATE NOW(3)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;