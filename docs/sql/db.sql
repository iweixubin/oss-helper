CREATE DATABASE `storage_service` CHARACTER SET 'ascii' COLLATE 'ascii_bin';

CREATE TABLE `storage_service`.`video`  (
  `Length` bigint NOT NULL COMMENT '文件的长度',
  `Tail` bigint NOT NULL COMMENT '取文件最后 8个byte 转成一个 int64 作为尾部值',
  `MD5` char(32) NOT NULL COMMENT '文件的 MD5 哈希值',
  `SHA512` char(128) NOT NULL COMMENT '文件的 SHA512 哈希值',
  `Signature` varchar(32) NOT NULL COMMENT '文件标识',
  `FilePath` varchar(255) NOT NULL COMMENT '文件路径',
  `FileName` varchar(255) NOT NULL COMMENT '文件名'
);