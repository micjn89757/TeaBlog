-- SET datestyle ''

-- ----------------------------
-- Create Role
-- ----------------------------

CREATE ROLE blog_dev 
LOGIN CREATEDB CREATEROLE REPLICATION 
PASSWORD 'asdjkl' VALID UNTIL '2035-01-01';


-- ----------------------------
-- Create Database
-- ----------------------------
DROP DATABASE IF EXISTS blog;
CREATE DATABASE IF NOT EXIST blog 
OWNER blog_dev;


-- ----------------------------
-- Create Schema
-- ----------------------------
CREATE SCHEMA "dev" AUTHORIZATION blog_dev;

-- ----------------------------
-- Table structure for user
-- ----------------------------
CREATE TABLE "dev"."user" (
    "id" uuid NOT NULL,
    "created_at" timestamp(3) NOT NULL DEFAULT localtimestamp,
    "deleted_at" timestamp(3),
    "updated_at" timestamp(3),
    "password" varchar(20),
    "role" smallint NOT NULL,
    CONSTRAINT "pk_usr_id" PRIMARY KEY (id)
);



-- ----------------------------
-- Table structure for article
-- ----------------------------
CREATE TABLE "dev"."article" (
    "id" uuid NOT NULL,
    "created_at" timestamp(3) NOT NULL DEFAULT localtimestamp,
    "deleted_at"  timestamp(3),
    "updated_at" timestamp(3),
    "title" varchar(50) NOT NULL,
    "content" text NOT NULL,
    "html_content" text,
    "desc" varchar(200),
    "author_id" uuid,
    "seq_id" bigserial,
    "comment_id" uuid,
    "cid" uuid, 
    "read_count" bigint DEFAULT 0,
    CONSTRAINT "pk_art_id" PRIMARY KEY (id)
);

CREATE INDEX "idx_seq_id" ON "dev"."article" USING BTREE ("seq_id");
CREATE INDEX "idx_category_id" ON "dev"."article" USING BTREE ("cid");


-- ----------------------------
-- Table structure for category
-- ----------------------------
CREATE TABLE "dev"."category" (
    "id" uuid NOT NULL,
    "created_at" timestamp(3) NOT NULL DEFAULT localtimestamp,
    "deleted_at"  timestamp(3),
    "name" varchar(20) NOT NULL,
    CONSTRAINT "pk_cid" PRIMARY KEY ("id")
);

-- ----------------------------
-- Table structure for profile
-- ----------------------------
