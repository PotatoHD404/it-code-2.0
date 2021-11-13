INSERT INTO itcode.promos (id, min_order_sum, promocode, priority, action, discount, title, scope)
VALUES (1, null, '10percent', 1, 'percent_discount', 10.00, 'Скидка 10% за заказ по промокоду', 'order');
INSERT INTO itcode.promos (id, min_order_sum, promocode, priority, action, discount, title, scope)
VALUES (2, null, 'noaction', 2, 'percent_discount', 0.00, 'Скидка 0 %', 'order');
INSERT INTO itcode.promos (id, min_order_sum, promocode, priority, action, discount, title, scope)
VALUES (3, null, 'discount1000', 3, 'price_discount', 1000.00, 'Скидка 1000 за заказ по промокоду', 'order');
INSERT INTO itcode.promos (id, min_order_sum, promocode, priority, action, discount, title, scope)
VALUES (4, null, '1peperoni', 4, 'gift', null, 'Подарок 1 пеперони за заказ с промокодом', 'order');
INSERT INTO itcode.promos (id, min_order_sum, promocode, priority, action, discount, title, scope)
VALUES (5, null, '2peperoni', 5, 'gift', null, 'Подарок 2 пеперони за заказ с промокодом', 'order');
INSERT INTO itcode.promos (id, min_order_sum, promocode, priority, action, discount, title, scope)
VALUES (6, null, 'nogift', 6, 'gift', null, 'В подарок ничего по промокоду', 'order');
INSERT INTO itcode.promos (id, min_order_sum, promocode, priority, action, discount, title, scope)
VALUES (7, null, '', 7, 'gift', null, 'В подарок ничего за Наггетсы', 'item');
INSERT INTO itcode.promos (id, min_order_sum, promocode, priority, action, discount, title, scope)
VALUES (8, null, '3xpepsi', 8, 'price_discount', 50.00,
        'Каждая третья пепси в заказе со скидкой 50 рублей по промокоду 3xpepsi', 'item');
INSERT INTO itcode.promos (id, min_order_sum, promocode, priority, action, discount, title, scope)
VALUES (9, 1000.00, 'pepsi', 9, 'gift', null, 'Пепси 1л в подарок при сумме заказа от 1000р по промокоду pepsi',
        'order');
INSERT INTO itcode.promos (id, min_order_sum, promocode, priority, action, discount, title, scope)
VALUES (10, null, 'pepsi10', 10, 'percent_discount', null,
        'Пепси 1л со скидкой 10% при покупке Пепси 1л и двух пицц Пепперони по промокоду pepsi10', 'item');
INSERT INTO itcode.promos (id, min_order_sum, promocode, priority, action, discount, title, scope)
VALUES (99, 400.00, 'test', 99, 'gift', 30.00, 'Сложная скидка', 'item');