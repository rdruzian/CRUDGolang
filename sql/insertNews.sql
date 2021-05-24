create procedure News.insertNews(
    title varchar(150),
    category varchar(20),
    text TEXT
)

BEGIN
    set @title = title;
    set @category = category;
    set @text = text;

    insert into Table_News (title, category, text) values (@title, @category, @text);
end;