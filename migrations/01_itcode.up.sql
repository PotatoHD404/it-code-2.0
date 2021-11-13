drop table items

drop table promos

drop table promo_condition_item

drop table promo_exclusions

drop table promo_gift_items

drop table promo_item_selector

create table items
(
    id    int unsigned auto_increment
        primary key,
    title varchar(100)  not null,
    price decimal(7, 2) not null
);

create table promos
(
    id            int unsigned auto_increment
        primary key,
    min_order_sum decimal(7, 2)                                       null,
    promocode     varchar(100)                                        not null,
    priority      int unsigned                                        not null,
    action        enum ('percent_discount', 'price_discount', 'gift') not null,
    discount      decimal(7, 2)                                       null,
    title         varchar(100)                                        not null,
    scope         enum ('order', 'item')                              not null
);

create table promo_condition_item
(
    id       int unsigned auto_increment
        primary key,
    promo_id int unsigned not null,
    item_id  int unsigned not null,
    constraint promo_condition_item_FK
        foreign key (item_id) references items (id),
    constraint promo_condition_item_FK_1
        foreign key (promo_id) references promos (id)
);

create table promo_exclusions
(
    id                 int unsigned auto_increment
        primary key,
    promo_id           int unsigned not null,
    exclusion_promo_id int unsigned not null,
    constraint promo_exclusions_FK
        foreign key (promo_id) references promos (id),
    constraint promo_exclusions_FK_1
        foreign key (exclusion_promo_id) references promos (id)
);

create table promo_gift_items
(
    id       int unsigned auto_increment
        primary key,
    promo_id int unsigned not null,
    item_id  int unsigned not null,
    constraint promo_gift_items_FK
        foreign key (promo_id) references promos (id),
    constraint promo_gift_items_FK_1
        foreign key (item_id) references items (id)
);

create table promo_item_selector
(
    id       int unsigned auto_increment
        primary key,
    item_id  int unsigned not null,
    promo_id int unsigned not null,
    constraint promo_item_selector_FK
        foreign key (item_id) references items (id),
    constraint promo_item_selector_FK_1
        foreign key (promo_id) references promos (id)
);

create table orders
(
    id        int unsigned auto_increment
        primary key,
    cart_id   varchar(10)  not null,
    promocode varchar(100) not null
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
