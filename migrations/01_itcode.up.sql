drop table promo_condition_item;

drop table promo_exclusions;

drop table promo_gift_items;

drop table promo_item_selector;

drop table items;

drop table promos;

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
