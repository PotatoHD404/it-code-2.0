create table orders
(
    id            int unsigned auto_increment
        primary key,
    cart_id       varchar(10)  not null,
    promocode     varchar(100) not null,
    cart_sum      float        not null,
    cart_discount float        not null
);

create table cart_items
(
    id           int unsigned auto_increment
        primary key,
    price        float        null,
    item_id      int unsigned not null,
    cart_id      int unsigned not null,
    orig_price   float        not null,
    cart_item_id varchar(30)  not null,

    constraint cart_items_items_id_fk
        foreign key (item_id) references items (id)
            on update cascade on delete cascade,
    constraint cart_items_orders_id_fk
        foreign key (cart_id) references orders (id)
            on update cascade on delete cascade
);



create table cart_promos
(
    id       int unsigned auto_increment
        primary key,
    promo_id int unsigned not null,
    cart_id  int unsigned not null,
    constraint cart_promos_orders_id_fk
        foreign key (cart_id) references orders (id),
    constraint cart_promos_promos_id_fk
        foreign key (promo_id) references promos (id)
);
