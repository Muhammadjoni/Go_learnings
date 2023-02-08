CREATE TABLE "users" (
    "id" serial NOT NULL unique,
    "username" varchar PRIMARY KEY,
    "email" varchar UNIQUE NOT NULL,
    "age" int NOT NULL
);
