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
DROP TABLE IF EXISTS "dev"."user" CASCADE;
CREATE TABLE "dev"."user" (
    "id" uuid NOT NULL,
    "created_at" timestamp(3) NOT NULL DEFAULT localtimestamp,
    "deleted_at" timestamp(3),
    "updated_at" timestamp(3),
    "username" varchar(20) NOT NULL,
    "password" varchar(20) NOT NULL,
    "role" smallint NOT NULL,
    CONSTRAINT "pk_usr_id" PRIMARY KEY ("id")
);
CREATE INDEX "idx_usr_del_at" ON "dev"."user" USING BTREE ("deleted_at");


-- ----------------------------
-- Table structure for article
-- ----------------------------
DROP TABLE IF EXISTS "dev"."article"  CASCADE;
CREATE TABLE "dev"."article" (
    "id" uuid NOT NULL,
    "created_at" timestamp(3) NOT NULL DEFAULT localtimestamp,
    "deleted_at"  timestamp(3),
    "updated_at" timestamp(3),
    "title" varchar(100) NOT NULL,
    "content" text NOT NULL,
    "html_content" text,
    "desc" varchar(200),
    "user_id" uuid,
    "seq_id" bigserial,
    "comment_id" uuid,
    "cid" uuid, 
    "read_count" bigint DEFAULT 0,
    CONSTRAINT "pk_art_id" PRIMARY KEY ("id")
);

CREATE INDEX "idx_seq_id" ON "dev"."article" USING BTREE ("seq_id");
CREATE INDEX "idx_category_id" ON "dev"."article" USING BTREE ("cid");
CREATE INDEX "idx_art_del_at" ON "dev"."article" USING BTREE ("deleted_at");

-- ----------------------------
-- Table structure for category
-- ----------------------------
DROP TABLE IF EXISTS "dev"."category" CASCADE;
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
DROP TABLE IF EXISTS "dev"."profile" CASCADE;
CREATE TABLE "dev"."profile" (
    "id" uuid NOT NULL,
    "created_at" timestamp(3) NOT NULL DEFAULT localtimestamp,
    "deleted_at" timestamp(3),
    "desc" varchar(200),
    "qq" varchar(20),
    "wechat" varchar(20),
    "weibo" varchar(20),
    "github" varchar(100),
    "bili" varchar(100),
    "email" varchar(100),
    "img" bytea,
    "avatar" bytea,

    CONSTRAINT "pk_pid" PRIMARY KEY ("id")
);
CREATE INDEX "idx_prof_del_at" ON "dev"."profile" USING BTREE ("deleted_at");

-- ----------------------------
-- Table structure for comment
-- ----------------------------
DROP TABLE IF EXISTS "dev"."comment" CASCADE;
CREATE TABLE "dev"."comment" (
    "id" uuid NOT NULL,
    "created_at" timestamp(3) NOT NULL DEFAULT localtimestamp,
    "deleted_at" timestamp(3),
    "user_id" uuid NOT NULL,
    "content" text NOT NULL,
    "article_id" uuid NOT NULL,
    "article_title" varchar(50),
    "status" smallint NOT NULL DEFAULT 2,
    "username" varchar(20),
    "title" varchar(100),

    CONSTRAINT "pk_comment_id" PRIMARY KEY ("id")
);

CREATE INDEX "idx_comment_del_at" ON "dev"."comment" USING BTREE ("deleted_at");