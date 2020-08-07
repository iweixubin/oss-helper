CREATE DATABASE `storage_service` CHARACTER SET 'ascii' COLLATE 'ascii_bin';

CREATE TABLE `storage_service`.`video`  (
  `Length` bigint NOT NULL COMMENT '文件的长度',
  `Tail` bigint NOT NULL COMMENT '取文件最后 8个byte 转成一个 int64 作为尾部值',
  `MD5` char(32) NOT NULL COMMENT '文件的 MD5 哈希值',
  `SHA512` char(128) NOT NULL COMMENT '文件的 SHA512 哈希值',
  `Signature` varchar(32) NOT NULL COMMENT '文件标识',
  `BucketName` varchar(63) NOT NULL COMMENT '存储空间名称，阿里云限制63个字节',
  `ObjectName` varchar(255) NOT NULL COMMENT '存储对象名称',
  `CreatedAt` datetime NOT NULL COMMENT '创建时间'
);

CREATE TABLE `storage_service`.`image_9`  (
  `Length` bigint NOT NULL COMMENT '文件的长度',
  `Tail` bigint NOT NULL COMMENT '取文件最后 8个byte 转成一个 int64 作为尾部值',
  `MD5` char(32) NOT NULL COMMENT '文件的 MD5 哈希值',
  `SHA512` char(128) NOT NULL COMMENT '文件的 SHA512 哈希值',
  `Signature` varchar(32) NOT NULL COMMENT '文件标识',
  `BucketName` varchar(63) NOT NULL COMMENT '存储空间名称，阿里云限制63个字节',
  `ObjectName` varchar(255) NOT NULL COMMENT '存储对象名称',
  `CreatedAt` datetime NOT NULL COMMENT '创建时间'
);



CREATE TABLE `storage_service`.`temp`  (
  `Id` bigint NOT NULL COMMENT '主键值为过期时间的数字化',
  `BucketName` varchar(63) NOT NULL COMMENT '存储空间名称，阿里云限制63个字节',
  `ObjectName` varchar(255) NOT NULL COMMENT '存储对象名称',
  PRIMARY KEY (`Id`)
);