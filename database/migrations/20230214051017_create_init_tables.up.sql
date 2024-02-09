
CREATE TABLE "orders" (
     "id" serial not null,
     "order_name" varchar unique not null,
     "customer_name" varchar not null,
     "customer_company" varchar not null,
     "delivered_amount" float,
     "total_amount" float not null,
     "order_date" timestamptz not null default current_timestamp
);
