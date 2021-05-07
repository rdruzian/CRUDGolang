use News;

create table Table_News(
    id integer not null auto_increment primary key,
    title varchar(150) not null,
    category varchar(20),
    text varchar(65535) not null
)