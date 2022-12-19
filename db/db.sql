create table atoms
(
    id         int auto_increment
        primary key,
    sid        bigint default 0 not null,
    content    text             null,
    link       varchar(200)     null,
    type       tinyint          null comment '图片、音频等',
    version    int    default 0 not null,
    doc_id     bigint default 0 not null,
    prev_id    bigint default 0 not null,
    created_at timestamp        null,
    updated_at timestamp        null,
    deleted_at timestamp        null
)
    collate = utf8mb4_unicode_ci;

create index atoms_sid_version_index
    on atoms (sid, version);

create index doc_id__index
    on atoms (doc_id);

create index prev_id__index
    on atoms (prev_id);

create index sid__index
    on atoms (sid);

create table collections
(
    id         int auto_increment
        primary key,
    sid        bigint default 0 not null,
    name       varchar(200)     null,
    created_at timestamp        null,
    updated_at timestamp        null,
    deleted_at timestamp        null
)
    collate = utf8mb4_unicode_ci;

create index sid__index
    on collections (sid);

create table docs
(
    id            int auto_increment
        primary key,
    sid           bigint default 0 not null,
    collection_id bigint default 0 not null,
    title         varchar(200)     null,
    version       int    default 0 not null,
    created_at    timestamp        null,
    updated_at    timestamp        null,
    deleted_at    timestamp        null
)
    collate = utf8mb4_unicode_ci;

create index sid__index
    on docs (sid);


