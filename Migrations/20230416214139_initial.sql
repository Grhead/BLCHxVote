-- +goose Up
CREATE TABLE IF NOT EXISTS KeyLinks (
                                                Id CHAR(36) PRIMARY KEY,
                                                PublicKey CHAR(64) NOT NULL,
                                                PrivateKey CHAR(64) NOT NULL
);

CREATE TABLE IF NOT EXISTS RelationPatterns (
                                                Id CHAR(36) PRIMARY KEY,
                                                PersonIdentifier CHAR(64) NOT NULL,
                                                PrivateKeyTemplate CHAR(64) NOT NULL
);

CREATE TABLE IF NOT EXISTS PublicKeySets (
                                                Id CHAR(36) PRIMARY KEY,
                                                PublicKey CHAR(64) NOT NULL,
                                                IsUsed BOOLEAN NOT NULL,
                                                VotingAffiliation CHAR(64) NOT NULL
);

CREATE TABLE IF NOT EXISTS ElectionSubjects (
                                                Id CHAR(36) PRIMARY KEY,
                                                PublicKey CHAR(64) NOT NULL,
                                                Description TEXT NOT NULL,
                                                VotingAffiliation CHAR(64) NOT NULL
);

-- +goose Down
DROP TABLE KeyLinks;
DROP TABLE RelationPatterns;
DROP TABLE PublicKeySets;
DROP TABLE ElectionSubjects;
