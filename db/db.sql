#----------------------------------------------#
CREATE TABLE `atoms` (
                         `id` int NOT NULL AUTO_INCREMENT,
                         `sid` bigint default 0 not null,
                         `content` text,
                         `link` varchar(200),
                         `type` tinyint comment '图片、音频等',
                         `version` int default 0 not null,
                         `doc_id` bigint default 0 NOT NULL,
                         `prev_id` bigint default 0 NOT NULL,
                         `created_at` timestamp NULL DEFAULT NULL,
                         `updated_at` timestamp NULL DEFAULT NULL,
                         `deleted_at` timestamp NULL DEFAULT NULL,
                         constraint atom_pk
                             primary key (`id`)
) DEFAULT CHARSET=utf8mb4 COLLATE utf8mb4_unicode_ci;

create index sid__index
    on atoms (sid);

create index doc_id__index
    on atoms (doc_id);

create index prev_id__index
    on atoms (prev_id);

create index atoms_sid_version_index
    on atoms (sid, version);

#----------------------------------------------#
CREATE TABLE `docs` (
                         `id` int NOT NULL AUTO_INCREMENT,
                         `sid` bigint default 0 not null,
                         `collection_id` bigint default 0 NOT NULL,
                         `title` varchar(200),
                         `version` int default 0 NOT NULL,
                         `created_at` timestamp NULL DEFAULT NULL,
                         `updated_at` timestamp NULL DEFAULT NULL,
                         `deleted_at` timestamp NULL DEFAULT NULL,
                         constraint doc_pk
                             primary key (`id`)
) DEFAULT CHARSET=utf8mb4 COLLATE utf8mb4_unicode_ci;

create index sid__index
    on docs (sid);

#----------------------------------------------#

CREATE TABLE `collections` (
                        `id` int NOT NULL AUTO_INCREMENT,
                        `sid` bigint default 0 not null,
                        `name` varchar(200),
                        `created_at` timestamp NULL DEFAULT NULL,
                        `updated_at` timestamp NULL DEFAULT NULL,
                        `deleted_at` timestamp NULL DEFAULT NULL,
                        constraint collection_pk
                            primary key (`id`)
) DEFAULT CHARSET=utf8mb4 COLLATE utf8mb4_unicode_ci;

create index sid__index
    on collections (sid);

