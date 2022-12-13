#----------------------------------------------#
CREATE TABLE `atoms` (
                         `id` int NOT NULL AUTO_INCREMENT,
                         `sid` bigint default 0 not null,
                         `content` text,
                         `link` varchar(200),
                         `type` tinyint comment '图片、音频等',
                         `doc_id` bigint default 0 NOT NULL,
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

#----------------------------------------------#
CREATE TABLE `docs` (
                         `id` int NOT NULL AUTO_INCREMENT,
                         `sid` bigint default 0 not null,
                         `title` varchar(200),
                         `created_at` timestamp NULL DEFAULT NULL,
                         `updated_at` timestamp NULL DEFAULT NULL,
                         `deleted_at` timestamp NULL DEFAULT NULL,
                         constraint doc_pk
                             primary key (`id`)
) DEFAULT CHARSET=utf8mb4 COLLATE utf8mb4_unicode_ci;

create index sid__index
    on docs (sid);
