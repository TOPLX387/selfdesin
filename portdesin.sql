/*
 Navicat Premium Data Transfer

 Source Server         : 郭百川
 Source Server Type    : PostgreSQL
 Source Server Version : 140003
 Source Host           : 8.130.104.141:5432
 Source Catalog        : govue3
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 140003
 File Encoding         : 65001

 Date: 16/06/2022 12:12:02
*/


-- ----------------------------
-- Table structure for Vessel
-- ----------------------------
DROP TABLE IF EXISTS "public"."Vessel";
CREATE TABLE "public"."Vessel" (
  "ladingnum" varchar(50) COLLATE "pg_catalog"."default",
  "cname" varchar(50) COLLATE "pg_catalog"."default",
  "destination" varchar(50) COLLATE "pg_catalog"."default"
)
;

-- ----------------------------
-- Records of Vessel
-- ----------------------------
INSERT INTO "public"."Vessel" VALUES ('前进', '水泥', '刚果');
INSERT INTO "public"."Vessel" VALUES ('振兴', '煤炭', '俄罗斯');
INSERT INTO "public"."Vessel" VALUES ('阿萨德', '石斑鱼', '日本');
INSERT INTO "public"."Vessel" VALUES ('胜利', '金属垃圾', '印度');
INSERT INTO "public"."Vessel" VALUES ('复兴', '精钢建材', '南非');

-- ----------------------------
-- Table structure for loginto
-- ----------------------------
DROP TABLE IF EXISTS "public"."loginto";
CREATE TABLE "public"."loginto" (
  "loginto_id" int4 NOT NULL,
  "loginto_name" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "loginto_password" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "loginto_level" int4 NOT NULL,
  "loginto_remarks" varchar(255) COLLATE "pg_catalog"."default",
  "update_time" timestamptz(0) NOT NULL
)
;

-- ----------------------------
-- Records of loginto
-- ----------------------------
INSERT INTO "public"."loginto" VALUES (1, 'test', '123456', 2, '管理员', '2022-06-01 11:04:59+08');

-- ----------------------------
-- Table structure for menu
-- ----------------------------
DROP TABLE IF EXISTS "public"."menu";
CREATE TABLE "public"."menu" (
  "menu_id" int4 NOT NULL,
  "menu_address" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "menu_level" int4 NOT NULL,
  "menu_name" varchar(255) COLLATE "pg_catalog"."default" NOT NULL
)
;

-- ----------------------------
-- Records of menu
-- ----------------------------
INSERT INTO "public"."menu" VALUES (1, '/Newstockin', 2, '入库管理');
INSERT INTO "public"."menu" VALUES (2, '/Newstockout', 2, '出库管理');
INSERT INTO "public"."menu" VALUES (3, '/view', 2, '费用可视化');
INSERT INTO "public"."menu" VALUES (4, '/view2', 2, '储量可视化');
INSERT INTO "public"."menu" VALUES (5, '/vessel', 2, '货物管理');

-- ----------------------------
-- Table structure for newstockins
-- ----------------------------
DROP TABLE IF EXISTS "public"."newstockins";
CREATE TABLE "public"."newstockins" (
  "created_at" timestamptz(6),
  "whtime" char(8) COLLATE "pg_catalog"."default",
  "lpnum" varchar(20) COLLATE "pg_catalog"."default",
  "ibton" int8,
  "ibpieces" int8
)
;

-- ----------------------------
-- Records of newstockins
-- ----------------------------
INSERT INTO "public"."newstockins" VALUES ('2022-06-15 21:16:42.298906+08', '20210801', '黑EL9882', 30, 600);
INSERT INTO "public"."newstockins" VALUES ('2022-06-15 21:17:02.577942+08', '20210801', '黑ME0003', 33, 660);
INSERT INTO "public"."newstockins" VALUES ('2022-06-15 21:17:42.38923+08', '20210818', '辽CD5123', 33, 660);
INSERT INTO "public"."newstockins" VALUES ('2022-06-15 21:25:24.419037+08', '20210818', '辽BM7685', 33, 660);
INSERT INTO "public"."newstockins" VALUES ('2022-06-16 07:42:25.714258+08', '20210819', '黑MJ2996', 33, 660);
INSERT INTO "public"."newstockins" VALUES ('2022-06-16 07:42:53.493619+08', '20210820', '黑EK5309 ', 33, 660);
INSERT INTO "public"."newstockins" VALUES ('2022-06-16 07:43:22.571143+08', '20210821', '黑EK0575', 5, 100);

-- ----------------------------
-- Table structure for newstockouts
-- ----------------------------
DROP TABLE IF EXISTS "public"."newstockouts";
CREATE TABLE "public"."newstockouts" (
  "created_at" timestamptz(6),
  "detime" char(8) COLLATE "pg_catalog"."default",
  "ordernum" varchar(20) COLLATE "pg_catalog"."default",
  "ladingnum" varchar(50) COLLATE "pg_catalog"."default",
  "obton" int8,
  "outpieces" int8,
  "storsgefee" float8
)
;

-- ----------------------------
-- Records of newstockouts
-- ----------------------------
INSERT INTO "public"."newstockouts" VALUES ('2022-06-15 21:37:42.877915+08', '20210821', '6212751', '前进', 36, 720, 126);
INSERT INTO "public"."newstockouts" VALUES ('2022-06-15 21:37:55.381539+08', '20210821', '6212752', '振兴', 30, 600, 94.5);
INSERT INTO "public"."newstockouts" VALUES ('2022-06-15 21:47:23.217156+08', '20210901', '6212815', '阿萨德', 36, 720, 18);
INSERT INTO "public"."newstockouts" VALUES ('2022-06-16 07:45:12.888046+08', '20210913', '6212862', '胜利', 36, 720, 229.5);
INSERT INTO "public"."newstockouts" VALUES ('2022-06-16 07:46:41.353768+08', '20210908', '6212863', '复兴', 36, 720, 120);

-- ----------------------------
-- Function structure for count_ibton
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."count_ibton"();
CREATE OR REPLACE FUNCTION "public"."count_ibton"()
  RETURNS "pg_catalog"."int4" AS $BODY$

declare
counts integer;

begin select sum(ibton) into counts from "public"."newstockins";
return counts;
end;

$BODY$
  LANGUAGE plpgsql VOLATILE
  COST 100;

-- ----------------------------
-- Function structure for count_obton
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."count_obton"();
CREATE OR REPLACE FUNCTION "public"."count_obton"()
  RETURNS "pg_catalog"."int4" AS $BODY$
declare
counts integer;

begin select sum(obton) into counts from "public"."newstockouts";
return counts;
end;

$BODY$
  LANGUAGE plpgsql VOLATILE
  COST 100;

-- ----------------------------
-- Function structure for upstoragefee
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."upstoragefee"();
CREATE OR REPLACE FUNCTION "public"."upstoragefee"()
  RETURNS "pg_catalog"."trigger" AS $BODY$
BEGIN
IF NEW.storsgefee IS NULL
THEN
    SET NEW.storsgefee =-1;
END IF;
RETURN NULL;
END;
$BODY$
  LANGUAGE plpgsql VOLATILE
  COST 100;

-- ----------------------------
-- Primary Key structure for table loginto
-- ----------------------------
ALTER TABLE "public"."loginto" ADD CONSTRAINT "staff_pkey" PRIMARY KEY ("loginto_id");

-- ----------------------------
-- Triggers structure for table newstockouts
-- ----------------------------

